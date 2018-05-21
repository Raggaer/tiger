package controllers

import (
	"time"

	"github.com/raggaer/tiger/app/config"
)

// Context main controller for all actions
type Context struct {
	Start  time.Time
	Config *config.Config
}
