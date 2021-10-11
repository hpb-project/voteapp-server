package task

import "sync"

type Task interface {
	Start() error
	Stop() error
}

var (
	once      sync.Once
	gTaskList []Task
)

func Init() error {
	var e error
	once.Do(func() {
		gTaskList = make([]Task, 0)
		if boelistTask, err := newBoeListTask(); err != nil {
			e = err
		} else {
			gTaskList = append(gTaskList, boelistTask)
		}
	})
	return e
}

func Start() error {
	for _, t := range gTaskList {
		err := t.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func Stop() error {
	for _, t := range gTaskList {
		err := t.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}
