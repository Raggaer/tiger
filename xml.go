package main

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/xml"
)

type xmlTaskList struct {
	Path      string
	rw        sync.Mutex
	Errors    []*xmlTaskError
	Monsters  map[string]*xml.Monster
	Vocations map[string]*xml.Vocation
	Items     map[int]xml.Item
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
	tasks.Add(3)

	// Execute tasks
	go loadServerMonsters(taskList, tasks, cfg.Server.Path)
	go loadServerItems(taskList, tasks, cfg.Server.Path)
	go loadServerVocations(taskList, tasks, cfg.Server.Path)

	// Wait for all tasks to end
	tasks.Wait()

	// Check for errors
	if len(taskList.Errors) >= 1 {
		return nil, taskList.Errors[0]
	}
	return taskList, nil
}

func loadServerVocations(taskList *xmlTaskList, wg *sync.WaitGroup, path string) {
	defer wg.Done()

	// Load vocation list
	vocList, err := xml.LoadVocationList(filepath.Join(path, "data", "XML", "vocations.xml"))
	if err != nil {
		taskList.rw.Lock()
		taskList.Errors = append(taskList.Errors, &xmlTaskError{
			Name:  "Vocation list",
			Error: err,
		})
		taskList.rw.Unlock()
		return
	}

	// Convert vocation slice to map
	vocs := make(map[string]*xml.Vocation, len(vocList.Vocations))
	for e, i := range vocList.Vocations {
		vocs[strings.ToLower(i.Name)] = &vocList.Vocations[e]
	}

	// Set task vocation list
	taskList.rw.Lock()
	taskList.Vocations = vocs
	taskList.rw.Unlock()
}

func loadServerItems(taskList *xmlTaskList, wg *sync.WaitGroup, path string) {
	defer wg.Done()

	// Load item list
	itemList, err := xml.LoadItemList(filepath.Join(path, "data", "items", "items.xml"))
	if err != nil {
		taskList.rw.Lock()
		taskList.Errors = append(taskList.Errors, &xmlTaskError{
			Name:  "Item list",
			Error: err,
		})
		taskList.rw.Unlock()
		return
	}

	// Convert item slice to map
	items := make(map[int]xml.Item, len(itemList.Items))
	for _, i := range itemList.Items {
		if i.FromID != 0 && i.ToID != 0 {

			// Populate items range
			for x := i.FromID; x <= i.ToID; x++ {
				items[x] = i
			}
			continue
		}

		// Populate normal item
		if i.ID != 0 {
			items[i.ID] = i
		}
	}

	// Set task item list
	taskList.rw.Lock()
	taskList.Items = items
	taskList.rw.Unlock()
}

func loadServerMonsters(taskList *xmlTaskList, wg *sync.WaitGroup, path string) {
	defer wg.Done()

	// Load monster list
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
		monsters[strings.ToLower(xmlMonster.Name)] = xmlMonster
	}

	// Set task monster list
	taskList.rw.Lock()
	taskList.Monsters = monsters
	taskList.rw.Unlock()
}
