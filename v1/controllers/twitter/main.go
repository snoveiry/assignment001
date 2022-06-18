package twitter

import (
	"github.com/snoveiry/assignment001/config"
)

type Controller struct {
	Assignment001 config.Assignment001
}

func NewController() *Controller {
	return &Controller{}
}