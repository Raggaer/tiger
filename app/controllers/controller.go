package controllers

import (
	"time"

	"github.com/raggaer/tiger/app/config"
	"github.com/raggaer/tiger/app/xml"
)

// Context main controller for all actions
type Context struct {
	Start    time.Time
	Config   *config.Config
	Monsters map[string]*xml.Monster
	Items    map[int]xml.Item
}
