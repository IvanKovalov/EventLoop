package engine

import (
	"strings"
)

type Palindrom struct {
	word string
}

func (po *Palindrom) Execute(h Handler) {
	res := make([]string, 0)
	i := len(po.word) - 1
	str := strings.Split(po.word, "")
	for i != -1 {
		res = append(res, str[i])
		i--
	}
	res = append(res, po.word)
	output := strings.Join(res, "")
	h.Post(NewPrintCommand(output))
}

func NewPolindrom(str string) *Palindrom {
	return &Palindrom{
		word: str,
	}
}
