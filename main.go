// Package assignment001 is our main API backend 
package assignment001

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gocql/gocql"
)

type Assignment001 struct {
	Config config.Assignment001
}

func New() *Assignment001 {
	return &Assignment001{}
}

func (a *Assignment001) Run() {
	var err error

	// StartPostgreSQL session
	/*a.Config.PostgresSession, err = database.Create(a.Config.PostgresUsername, a.Config.PostgresPassword, a.Config.PostgresDatabase, a.Config.PostgresHost)
	logger.FatalIfError(err)*/

	// Start the router
	a.router()
}
