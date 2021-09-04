package task

import (
	"fmt"
	"sync"
	"time"
	"tm/pkg/setting"
)

type Task struct {
	wg         sync.WaitGroup  //wait for all gorotines finished
	tasks      []TaskDetail    //tasks
	taskChan   chan TaskDetail //job scheduler
	max        int             //max count of gorotine
	cost       int64           //milliseconds of this job
	quickMode  bool            //if set to quick mode, jobs will be executed when added
	curTaskNum int
}

/*
* GoTaskDetail defines methods
 */
type TaskDetail struct {
	fn     func(map[string]string)
	params map[string]string
}

/*
* Generate GoTask manager
 */
func NewGoTask(maxConcurentNum int) *Task {
	ret := &Task{
		wg:         sync.WaitGroup{},
		tasks:      make([]TaskDetail, 0),
		taskChan:   make(chan TaskDetail, maxConcurentNum),
		max:        maxConcurentNum,
		curTaskNum: 0,
	}
	return ret
}

/*
* Record how much time spent for all jobs.
 */
func (self *Task) Cost() int64 {
	return self.cost
}

/*
* Add tasks
 */
func (self *Task) Add() {
	if Size() > 0 {
		brokers := Pop()
		for _, task := range brokers {
			self.tasks = append(self.tasks, task)
		}
	}
}

/*
* Get paramters from context
 */
func (self *Task) GetParamter(index int, params interface{}) interface{} {
	if p, ok := params.([]interface{}); ok {
		if len(p) > 0 {
			if t, ok := p[0].([]interface{}); ok {
				if len(t) > index {
					return t[index]
				}
			}
		}
	}
	return nil
}

/*
* Start concurrent tasks
 */
func (self *Task) Start() {
	for _, v := range self.tasks {
		self.wg.Add(1)
		go func(v TaskDetail) {
			defer func() {
				self.Done()
			}()
			v.fn(v.params)
		}(v)
	}
	self.tasks = make([]TaskDetail, 0)
	self.wg.Wait()
}

/*
* if set quickMode == true, you must invoke Done() to finish manually. Deprecated
 */
func (self *Task) Done() {
	// if self.quickMode == false {
	// 	return
	// }
	self.wg.Done()

}

func Listen() {
	fmt.Println("异步队列开始")
	go func() {
		task := NewGoTask(setting.MaxCpuNum)
		for {
			task.Add()
			task.Start()
			time.Sleep(10 * time.Second)
		}
	}()
}
