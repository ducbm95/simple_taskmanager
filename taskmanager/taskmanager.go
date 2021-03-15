package taskmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/robfig/cron/v3"
)

const (
	ACTIVE   = 1
	INACTIVE = 0
)

type Task struct {
	Name        string
	CronExpr    string
	cronID      cron.EntryID
	state       int
	TaskContent string
}

func (task *Task) doJob() {
	fmt.Println(task.TaskContent)
}

type TaskManager struct {
	taskMap map[string]*Task
	c       *cron.Cron
	parser  cron.Parser
}

func NewTaskManager() *TaskManager {
	tm := TaskManager{}

	tm.taskMap = make(map[string]*Task)
	tm.parser = cron.NewParser(
		cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor,
	)

	tm.c = cron.New()
	tm.c.Start()

	return &tm
}

func (tm *TaskManager) Add(task Task) error {
	_, ok := tm.taskMap[task.Name]
	if ok {
		return errors.New("task already existed")
	}

	_, err := tm.parser.Parse(task.CronExpr)
	if err != nil {
		return err
	}
	tm.taskMap[task.Name] = &task
	return nil
}

func (tm *TaskManager) Start(name string) error {
	tb, ok := tm.taskMap[name]
	if !ok {
		return errors.New("task not existed")
	}
	if tb.state == ACTIVE {
		return errors.New("task already running")
	}
	cronID, err := tm.c.AddFunc(tb.CronExpr, tb.doJob)
	tb.cronID = cronID
	tb.state = ACTIVE

	return err
}

func (tm *TaskManager) Stop(name string) error {
	tb, ok := tm.taskMap[name]
	if !ok {
		return errors.New("task not existed")
	}
	tm.c.Remove(tb.cronID)
	tb.state = INACTIVE
	return nil
}

func (tm *TaskManager) Pause(name string) {
	tm.Stop(name)
}

func (tm *TaskManager) Resume(name string) error {
	return tm.Start(name)
}

func (tm *TaskManager) Remove(name string) error {
	_, ok := tm.taskMap[name]
	if !ok {
		return errors.New("task not existed")
	} else {
		err := tm.Stop(name)
		if err != nil {
			return err
		}

		delete(tm.taskMap, name)

		if len(tm.taskMap) <= 0 {
			tm.c.Stop()
		}

		return nil
	}
}

func (tm *TaskManager) Export(name string) error {
	tb, ok := tm.taskMap[name]
	if !ok {
		return errors.New("task not existed")
	}

	b, err := json.Marshal(tb)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(name, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (tm *TaskManager) Import(name string) error {
	_, ok := tm.taskMap[name]
	if ok {
		return errors.New("task already existed")
	}

	b, err := ioutil.ReadFile(name)
	if err != nil {
		return err
	}

	var task Task
	err = json.Unmarshal(b, &task)
	if err != nil {
		return err
	}

	return tm.Add(task)
}
