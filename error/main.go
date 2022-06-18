// Package error handles error formats and responses for the assignment001-sina application
package error

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

const ReferenceContextKey = "Reference"

type Default struct {
	Errors []*E `json:"errors,omitempty"`
}

type E struct {
	Type        string `json:"type,omitempty" example:"BAD_VALUE"`
	Description string `json:"error,omitempty" example:"A description explaining  the error may be shown here."`
	Field       string `json:"field,omitempty" example:"FirstName"`
	Rule        string `json:"rule,omitempty"  example:"name_validator"`
	Docs        string `json:"docs,omitempty"`
	Status      int    `json:"-"` // Status is option and should only be used to override the default status when needed
}

func JSON(ctx *gin.Context, status int, errors ...*E) {
	ctx.JSON(determineStatus(status, errors), Default{
		Errors: errors,
	})
}

func determineStatus(s int, errors []*E) int {
	for _, e := range errors {
		if e != nil {
			if e.Status != 0 {
				return e.Status
			}
		}
	}
	return s
}

func (e *E) Error() string {
	if e != nil {
		errorStr, err := json.Marshal(e)
		if err != nil {
			return string(errorStr)
		}
		return e.Description
	}
	return ""
}
