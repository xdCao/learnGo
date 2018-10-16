package main

import "fmt"

func main() {
	actor := new(Actor)
	RegisterEvent("OnSkill", actor.OnEvent)
	RegisterEvent("OnSkill", GlobalEvent)
	CallEvent("OnSkill", 100)

}

var eventByName = make(map[string][]func(interface{}))

func RegisterEvent(name string, callback func(interface{})) {
	list := eventByName[name]

	list = append(list, callback)

	eventByName[name] = list
}

func CallEvent(name string, param interface{}) {
	list := eventByName[name]
	for _, callback := range list {
		callback(param)
	}
}

type Actor struct {
}

func (a *Actor) OnEvent(param interface{}) {
	fmt.Println("actor event: ", param)
}

func GlobalEvent(param interface{}) {
	fmt.Println("Global event: ", param)
}
