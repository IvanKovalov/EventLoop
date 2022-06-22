package engine

import (
	"errors"
	"fmt"
	"sync"
)

func NewPrintCommand(arg string) *printCommand {
	return &printCommand{
		arg: arg,
	}
}

type printCommand struct {
	arg string
}

func (p *printCommand) Execute(h Handler) {
	fmt.Println(p.arg)
}

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command) error
}

type EventLoop struct {
	q          *commandsQueue
	stop       bool
	stopSignal chan struct{}
}

func (l *EventLoop) Start() {
	l.q = &commandsQueue{
		notEmpty: make(chan struct{}),
	}
	l.stopSignal = make(chan struct{})
	go func() {
		for !l.stop || !l.q.empty() {
			cmd := l.q.pull()
			cmd.Execute(l)
		}
		l.stopSignal <- struct{}{}
	}()
}

type stopCommand struct{}

func (s stopCommand) Execute(h Handler) {
	h.(*EventLoop).stop = true
}

func (l *EventLoop) AwaitFinish() {
	l.Post(stopCommand{})
	<-l.stopSignal
}

func (l *EventLoop) Post(cmd Command) error {
	if l.stop == true {
		return errors.New("error! can't post command if event loop isn't running")
	} else {
		l.q.push(cmd)
		return nil
	}

	//l.q.push(cmd)
}

type commandsQueue struct {
	a        []Command
	mu       sync.Mutex
	notEmpty chan struct{}
	wait     bool
}

func (cq *commandsQueue) push(c Command) {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	cq.a = append(cq.a, c)
	if cq.wait {
		cq.notEmpty <- struct{}{}
	}
}

func (cq *commandsQueue) pull() Command {
	cq.mu.Lock()
	defer cq.mu.Unlock()

	if len(cq.a) == 0 {
		cq.wait = true
		cq.mu.Unlock()
		<-cq.notEmpty
		cq.mu.Lock()
	}

	res := cq.a[0]
	cq.a[0] = nil
	cq.a = cq.a[1:]
	return res
}

func (cq *commandsQueue) empty() bool {
	cq.mu.Lock()
	defer cq.mu.Unlock()
	return len(cq.a) == 0
}
