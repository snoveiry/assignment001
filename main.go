// Package assignment001 is our main API backend
package assignment001

import (
	"github.com/snoveiry/assignment001/config"
)

type Assignment001 struct {
	Config config.Assignment001
}

func New() *Assignment001 {
	return &Assignment001{}
}

func (a *Assignment001) Run() {
	// Start the router
	a.Rout()
}
