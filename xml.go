package main

import (
	"path/filepath"
	"sync"

	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/xml"
)

type xmlTaskList struct {
	Path     string
	rw       sync.Mutex
	Errors   []*xmlTaskError
	Monsters map[string]*xml.Monster
}

type xmlTaskError struct {
	Name  string
	Error error
}

func loadServerData(cfg *config.Config) (*xmlTaskList, *xmlTaskError) {
	taskList := &xmlTaskList{
		Path: cfg.Server.Path,
	}
	// Create wait group for all parsing tasks
	tasks := &sync.WaitGroup{}
	tasks.Add(1)

	// Execute tasks
	go loadServerMonsters(taskList, tasks, cfg.Server.Path)

	// Wait for all tasks to end
	tasks.Wait()

	// Check for errors
	if len(taskList.Errors) >= 1 {
		return nil, taskList.Errors[0]
	}
	return taskList, nil
}

func loadServerMonsters(taskList *xmlTaskList, wg *sync.WaitGroup, path string) {
	defer wg.Done()

	// Load monster list file
	monsterList, err := xml.LoadMonsterList(filepath.Join(path, "data", "monster", "monsters.xml"))
	if err != nil {
		taskList.rw.Lock()
		taskList.Errors = append(taskList.Errors, &xmlTaskError{
			Name:  "Monster list",
			Error: err,
		})
		taskList.rw.Unlock()
		return
	}

	// Load each monster from the main list
	monsters := make(map[string]*xml.Monster, len(monsterList.Monsters))
	for _, m := range monsterList.Monsters {
		xmlMonster, err := xml.LoadMonster(filepath.Join(taskList.Path, "data", "monster", m.File))
		if err != nil {
			taskList.rw.Lock()
			taskList.Errors = append(taskList.Errors, &xmlTaskError{
				Name:  "Load monster " + m.Name,
				Error: err,
			})
			taskList.rw.Unlock()
			return
		}

		// Append monster to the list
		monsters[xmlMonster.Name] = xmlMonster
	}

	// Set task monster list
	taskList.rw.Lock()
	taskList.Monsters = monsters
	taskList.rw.Unlock()
}
