package main

import (
	"container/list"
	"fmt"
)

func main() {

	l := list.New()
	//	尾部添加
	l.PushBack("canon")
	//头部添加
	l.PushFront(67)
	//尾部添加后保存元素句柄
	ele := l.PushBack("first")
	//	在ele之后添加
	l.InsertAfter("high", ele)
	//	在ele之前添加
	l.InsertBefore("noon", ele)
	//	删除元素
	l.Remove(ele)

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
