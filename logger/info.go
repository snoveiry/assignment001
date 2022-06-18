// Package logger defines log formatting and provides functions for writing logs
package logger

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Info(ctx *gin.Context, message string) {
	jsonContent, err := json.Marshal(&gin.H{
		"time":      time.Now().Format(time.RFC3339),
		"info":      message,
	})
	if err == nil {
		fmt.Printf("%s\n", string(jsonContent))
	} else {
		log.Println(err, message)
	}
}
