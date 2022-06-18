// Package error handles error formats and responses for the assignment001-sina application
package error

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

const ReferenceContextKey = "Reference"

type Default struct {
	Reference string `json:"reference,omitempty" example:"PROD01FXKSPD0G1Q6VCJZED41WHXEE"`
	Errors    []*E   `json:"errors,omitempty"`
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
	reference := getReference(ctx)
	for _, e := range errors {
		if e.Docs != "" && strings.HasPrefix(e.Docs, "/") {
			env := os.Getenv("ENVIRONMENT")
			if env == "production" {
				e.Docs = "https://api.zeipt.io/"
			} else if env == "staging" {
				e.Docs = "https://staging.api.zeipt.io/"
			} else {
				e.Docs = "http://localhost:8080/"
			}
		}
	}
	ctx.JSON(determineStatus(status, errors), Default{
		Reference: reference,
		Errors:    errors,
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

func getReference(ctx *gin.Context) (reference string) {
	env := os.Getenv("ENVIRONMENT")
	if ctx != nil {
		tmpReference, exists := ctx.Get(ReferenceContextKey)
		if exists && reflect.TypeOf(tmpReference).Kind() == reflect.String {
			reference = tmpReference.(string)
		}
	}
	if len(env) >= 4 {
		if env == "staging" {
			env = strings.ToUpper(env)
		} else {
			env = strings.ToUpper(env[0:4])
		}
		reference = fmt.Sprintf("%s%s", env, reference)
	}
	return reference
}
