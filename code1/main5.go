package main

import (
	"fmt"
	"time"

	"github.com/RussellLuo/timingwheel"
)

type Task struct {
	tm   *time.Timer
	name string
}

type TimingTask struct {
	tw *timingwheel.TimingWheel
	ch chan *time.Timer
}

func (tt *TimingTask) Init() {
	tt.tw = timingwheel.NewTimingWheel(time.Millisecond, 20)
	tt.ch = make(chan *Task)
}

func (tt *TimingTask) Start() {
	tt.tw.Start()
}

func (tt *TimingTask) Stop() {
	tt.tw.Stop()

}

func (tt *TimingTask) Watch() {
	for {
		select {
		case task := <-tt.ch:
			{

			}
		default:
			time.Sleep(time.Millisecond * 20)
		}
	}
}

type Callback func()

func (tt *TimingTask) AddTask(n string, d time.Duration, f Callback) {
	t := tt.tw.AfterFunc(d, f)
	task := &Task{t, n}
	tt.ch <- task
}

func main() {
	tw := timingwheel.NewTimingWheel(time.Millisecond, 20)
	tw.Start()
	defer tw.Stop()

	t := tw.AfterFunc(time.Second, func() {
		fmt.Println("The timer fires")
	})
	fmt.Println(t.Stop())

	<-time.After(2 * time.Second)
	// Stop the timer before it fires
	fmt.Println(t.Stop())
}
