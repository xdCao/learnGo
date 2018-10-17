package main

import "fmt"

func main() {
	dict := NewDictionary()

	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("hah", 24)

	dict.Visit(func(k, v interface{}) bool {
		return true
	})

}

type Dictionary struct {
	data map[interface{}]interface{}
}

func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {
	if callback == nil {
		return
	}

	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
		fmt.Println(k, v)
	}
}

func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}
