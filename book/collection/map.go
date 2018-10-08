package main

import (
	"fmt"
	"sync"
)

func main() {

	//scene := make(map[string]int)
	//
	//scene["route"] = 6
	//
	//fmt.Println(scene["route"])
	//
	//v := scene["route2"]
	//
	//fmt.Println(v)
	//
	//scene["brazil"]=4
	//scene["china"]=960
	//
	//var sceneList []string
	//
	//for k := range scene  {
	//	sceneList = append(sceneList,k)
	//}
	//
	//sort.Strings(sceneList)
	//
	//fmt.Println(sceneList)

	//m := make(map[int]int)
	//
	//go func() {
	//	for  {
	//		m[1] = 1
	//	}
	//}()
	//
	//go func() {
	//	for  {
	//		_=m[1]
	//	}
	//}()
	//
	//select {}

	var scene sync.Map

	scene.Store("Greece", 97)
	scene.Store("London", 100)
	scene.Store("egypt", 200)

	fmt.Println(scene.Load("London"))

	scene.Delete("London")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("iterate: ", key, value)
		return true
	})

}
