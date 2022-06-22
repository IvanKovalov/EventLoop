package main

import (
	"reflect"
	"testing"

	"github.com/lab4/engine"
	"github.com/stretchr/testify/assert"
)

func TestParserPalindrom(t *testing.T) {
	commandString := "palindrom shox"
	command := Parser(commandString)
	exampleP := engine.NewPolindrom("shox")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(exampleP), reflect.TypeOf(command))
	}
}

func TestParserPrint(t *testing.T) {
	commandString := "print Hello world"
	command := Parser(commandString)
	examplePrint := engine.NewPrintCommand("Hello world")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(examplePrint), reflect.TypeOf(command))
	}
}

func TestParserDefault(t *testing.T) {
	commandString := ""
	command := Parser(commandString)
	examplePrint := engine.NewPrintCommand("Hello world")

	if assert.NotNil(t, command) {
		assert.IsType(t, reflect.TypeOf(examplePrint), reflect.TypeOf(command))
	}
}

func TestLoopPosting(t *testing.T) {
	command := engine.NewPrintCommand("Hello!")

	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	eventLoop.AwaitFinish()

	err := eventLoop.Post(command)
	assert.Error(t, err)
}
