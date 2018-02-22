//https://gist.github.com/NewbMiao/0ae9f9e7f4915d78a963980126e53f49

package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

type ActionListener interface {
	DoAction() (err error)
}

type Job struct {
	ActionListener
}

type ChanPool struct {
	maxWorker  int
	queueLen   int
	wg         sync.WaitGroup
	jobQueue   chan Job
	workerPool chan chan Job
	limiter    chan bool
}

var p ChanPool

type Worker struct {
	Id         int
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(i int) Worker {
	return Worker{
		Id:         i,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	for {
		p.workerPool <- w.JobChannel
		select {
		case job := <-w.JobChannel:
			p.limiter <- true
			fmt.Printf("====worker:%d, get job:%v====\n", w.Id, job.ActionListener)
			go func() {
				job.DoAction()
				<-p.limiter
				p.wg.Done()
			}()
		case <-w.quit:
			return
		}
	}
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func AddJob(ac ActionListener) {
	p.wg.Add(1)
	work := Job{ActionListener: ac}
	p.jobQueue <- work
}

func WaitDone() {
	p.wg.Wait()
	close(p.jobQueue)
	close(p.workerPool)
}

func init() {
	p = ChanPool{
		wg:        sync.WaitGroup{},
		maxWorker: runtime.NumCPU(),
		queueLen:  10,
		limiter:   make(chan bool, 10),
	}
	p.jobQueue = make(chan Job, p.queueLen)
	p.workerPool = make(chan chan Job, p.maxWorker)

	for i := 1; i <= p.maxWorker; i++ {
		worker := NewWorker(i)
		go worker.Start()
	}
	// 启动调度
	go dispatch()
}

func dispatch() {
	for job := range p.jobQueue {
		<-p.workerPool <- job
	}
}

type TestJob struct {
	id int
}

func (a *TestJob) DoAction() (err error) {
	time.Sleep(time.Second)
	return
}

func TestChanPool(t *testing.T) {
	for i := 1; i <= 50; i++ {
		j := &TestJob{i}
		AddJob(j)
	}
	WaitDone()
}
