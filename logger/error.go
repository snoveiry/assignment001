package logger

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	error1 "github.com/snoveiry/assignment001/error"
	"log"
	"os"
	"reflect"
	"strings"
	"time"
)

const ReferenceContextKey = "Reference"

func FatalError(ctx *gin.Context, e interface{}) {
	switch v := e.(type) {
	case *error1.E:
		if v != nil {
			jsonContent, err := json.Marshal(&gin.H{
				"time":      time.Now().Format(time.RFC3339),
				"error":     v.Error(),
				"reference": GetReference(ctx),
			})
			if err == nil {
				fmt.Printf("%s\n", string(jsonContent))
			} else {
				log.Println(err)
			}
			
			if err != nil {
				log.Println(err)
			}
			panic(e)
		}
	case error:
		if v != nil {
			jsonContent, err := json.Marshal(&gin.H{
				"time":      time.Now().Format(time.RFC3339),
				"error":     v.Error(),
				"reference": GetReference(ctx),
			})
			if err == nil {
				fmt.Printf("%s\n", string(jsonContent))
			} else {
				log.Println(err)
			}
			
			if err != nil {
				log.Println(err)
			}
			panic(e)
		}
	}
}

func FatalIfError(e error) {
	if e != nil {
		jsonContent, err := json.Marshal(&gin.H{
			"time":  time.Now().Format(time.RFC3339),
			"error": e.Error(),
		})
		if err == nil {
			fmt.Printf("%s\n", string(jsonContent))
		} else {
			log.Println(err)
		}
		
		if err != nil {
			log.Println(err)
		}
		panic(e)
	}
}

func Error(ctx *gin.Context, e interface{}) bool {
	switch v := e.(type) {
	case *error1.E:
		if v != nil {
			jsonContent, err := json.Marshal(&gin.H{
				"time":      time.Now().Format(time.RFC3339),
				"error":     v.Error(),
				"reference": GetReference(ctx),
			})
			if err == nil {
				fmt.Printf("%s\n", string(jsonContent))
			} else {
				log.Println(err, e)
			}
			return true
		}
	case error:
		if v != nil {
			jsonContent, err := json.Marshal(&gin.H{
				"time":      time.Now().Format(time.RFC3339),
				"error":     v.Error(),
				"reference": GetReference(ctx),
			})
			if err == nil {
				fmt.Printf("%s\n", string(jsonContent))
			} else {
				log.Println(err, e)
			}
			return true
		}
	}
	return false
}

func IsError(e error) bool {
	if e != nil {
		jsonContent, err := json.Marshal(&gin.H{
			"time":  time.Now().Format(time.RFC3339),
			"error": e.Error(),
		})
		if err == nil {
			fmt.Printf("%s\n", string(jsonContent))
		} else {
			log.Println(err, e)
		}
		return true
	}
	return false
}

func GetReference(ctx *gin.Context) (reference string) {
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
