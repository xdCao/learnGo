package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	l := CreateLogger()
	l.Log("hello")
}

type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writerList []LogWriter
}

func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

func (l *Logger) Log(data interface{}) {
	for _, writer := range l.writerList {
		writer.Write(data)
	}
}

func NewLogger() *Logger {
	return &Logger{}
}

type FileWriter struct {
	file *os.File
}

func (f *FileWriter) SetFile(filename string) (err error) {
	if f.file != nil {
		f.file.Close()
	}
	f.file, err = os.Create(filename)
	return err
}

func (f *FileWriter) Write(data interface{}) error {
	if f.file == nil {
		return errors.New("file not created")
	}
	str := fmt.Sprintf("%v\n", data)
	_, err := f.file.Write([]byte(str))
	return err
}

func NewFileWriter() *FileWriter {
	return &FileWriter{}
}

type ConsoleWriter struct {
}

func (f *ConsoleWriter) Write(data interface{}) error {
	str := fmt.Sprintf("%v\n", data)
	_, err := os.Stdout.Write([]byte(str))
	return err
}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{}
}

func CreateLogger() *Logger {
	l := NewLogger()
	cw := NewConsoleWriter()
	l.RegisterWriter(cw)
	fw := NewFileWriter()
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}
	l.RegisterWriter(fw)
	return l
}
