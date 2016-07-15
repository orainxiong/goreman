package main

import (
	"bytes"
	"github.com/daviddengcn/go-colortext"
	"log"
	"os"
	"sync"
	"fmt"
)

type clogger struct {
	idx    int
	proc   string
	logger *log.Logger
}

var colors = []ct.Color{
	ct.Green,
	ct.Cyan,
	ct.Magenta,
	ct.Yellow,
	ct.Blue,
	ct.Red,
}
var ci int

var mutex = new(sync.Mutex)

// write handler of logger.
func (l *clogger) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	for {
		line, err := buf.ReadBytes('\n')
		if len(line) > 1 {
			s := string(line)

			mutex.Lock()
			ct.ChangeColor(colors[l.idx], false, ct.None, false)
			l.logger.Print(s)
			ct.ResetColor()
			mutex.Unlock()
		}
		if err != nil {
			break
		}
	}
	return len(p), nil
}

// create logger instance.
func createLogger(proc string) *clogger {
	mutex.Lock()
	defer mutex.Unlock()
	l := &clogger{ci, proc, log.New(os.Stdout, fmt.Sprintf("%s: ",proc), log.Ldate|log.Ltime|log.Lshortfile)}
	ci++
	if ci >= len(colors) {
		ci = 0
	}
	return l
}
