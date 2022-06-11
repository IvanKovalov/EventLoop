package main

import (
	//"bufio"

	//"fmt"

	"io"
	"os"
	"strings"

	"github.com/lab4/engine"
)

func main() {

	loop := new(engine.EventLoop)

	file, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)
	var str string
	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		str = string(data[:n])
	}
	text := strings.Split(str, "\n")
	loop.Start()
	for i := 0; i < len(text); i++ {
		loop.Post(Parser(text[i]))
	}
	loop.AwaitFinish()
}

func Parser(str string) engine.Command {

	if strings.Contains(str, "print") {
		text := strings.Split(str, "print ")
		return engine.NewPrintCommand(string(text[1]))
	} else if strings.Contains(str, "polindrom") {
		text := strings.Split(str, "polindrom ")
		return engine.NewPolindrom(string(text[1]))
	}

	return engine.NewPrintCommand("SYNTAX ERROR: Not Enough Parameters")

}
