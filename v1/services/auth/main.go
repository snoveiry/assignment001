// Package auth will have all functions needed for auth
package auth

import "github.com/snoveiry/assignment001/config"

type Service struct {
	Assignment001 config.Assignment001
}

func New(config config.Assignment001) *Service {
	return &Service{
		Assignment001: config,
	}
}
