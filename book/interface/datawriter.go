package main

import "fmt"

func main() {
	f := new(file)
	var writer DataWriter
	writer = f
	writer.WriteData("data")
}

type DataWriter interface {
	WriteData(data interface{}) error
	//CanWrite() bool
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	fmt.Println("Write data: ", data)
	return nil
}
