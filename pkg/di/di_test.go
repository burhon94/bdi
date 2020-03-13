package di

import (
	"errors"
	"log"
	"testing"
)

type Msg string

func NewMsg() Msg {
	return "hello msg"
}

type dependency struct{
	value string
}

func NewDependency(message Msg) *dependency {
	log.Print("dependency created")
	return &dependency{string(message)}
}

type consumer struct {
	dep *dependency
}

func NewConsumer(dep *dependency) *consumer {
	if dep == nil {
		log.Print(errors.New("dependency can't be nil"))
	}
	log.Print("consumer created")
	return &consumer{dep: dep}
}

func TestNewMsg_NoDep(t *testing.T) {
	container := NewContainer()
	err := container.Provide(NewMsg)
	if err != nil {
		t.Fatalf("error just be nil: %v", err)
	}
}

func TestNewMsg_Err(t *testing.T) {
	container := NewContainer()
	err := container.Provide(NewMsg())
	if err == nil {
		t.Fatalf("just have error: %v", err)
	}
}

func TestNewDependency_Dep_Msg(t *testing.T) {
	container := NewContainer()
	err := container.Provide(NewMsg, NewDependency)
	if err != nil {
		t.Fatalf("error just be nil: %v", err)
	}
}

func TestNewConsumer(t *testing.T) {
	container := NewContainer()
	err := container.Provide(NewMsg, NewDependency, NewConsumer)
	if err != nil {
		t.Fatalf("error just be nil: %v", err)
	}
}



