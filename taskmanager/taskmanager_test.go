package taskmanager_test

import (
	"taskmanager/taskmanager"
	"testing"
)

func Test1(t *testing.T) {
	tm := taskmanager.NewTaskManager()
	task := taskmanager.Task{
		Name:        "Task1",
		CronExpr:    "* * * * *",
		TaskContent: "Do Math exercise",
	}
	err := tm.Add(task)
	err = tm.Remove(task.Name)

	if err != nil {
		t.Errorf("Expect: err = nil. Return: %v", err)
	}
}

func Test2(t *testing.T) {
	tm := taskmanager.NewTaskManager()
	task := taskmanager.Task{
		Name:        "Task1",
		CronExpr:    "* * * * *",
		TaskContent: "Do Math exercise",
	}
	err := tm.Add(task)
	err = tm.Add(task)

	if err == nil {
		t.Errorf("Expect: err == 'task already existed'. Return: nil")
	}
}

func Test3(t *testing.T) {
	tm := taskmanager.NewTaskManager()
	task1 := taskmanager.Task{
		Name:        "Task1",
		CronExpr:    "* * * * *",
		TaskContent: "Do Math exercise",
	}

	task2 := taskmanager.Task{
		Name:        "Task2",
		CronExpr:    "* * * * *",
		TaskContent: "Turn the light on",
	}
	err1 := tm.Add(task1)
	err2 := tm.Add(task2)
	err3 := tm.Start(task1.Name)
	err4 := tm.Start(task2.Name)
	err5 := tm.Stop(task1.Name)
	err6 := tm.Stop(task2.Name)
	err7 := tm.Remove(task1.Name)
	err8 := tm.Remove(task2.Name)

	if !(err1 == nil && err2 == nil && err3 == nil && err4 == nil &&
		err5 == nil && err6 == nil && err7 == nil && err8 == nil) {
		t.Errorf("Expect: err == nil. Return: not nil")

	}
}
