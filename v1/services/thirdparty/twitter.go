package thirdparty

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	error1 "github.com/snoveiry/assignment001/error"
)

func (s *Service) GetTweetService(ctx *gin.Context, url string, bearer string) (bool, []byte) {
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		error1.JSON(ctx, http.StatusNotFound, &error1.E{
			Type:        "END POINT ERROR",
			Description: "Some description goes here.",
		})
		return false, nil
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		error1.JSON(ctx, http.StatusInternalServerError, &error1.E{
			Type:        "SERVER ERROR",
			Description: "Calling from thirdparty has error.",
		})
		return false, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		error1.JSON(ctx, http.StatusBadRequest, &error1.E{
			Type:        "READ RESPONSE ERROR",
			Description: "Could not read service response.",
		})
		return false, nil
	}

	return true, body
}

func (s *Service) GetStreamService(ctx *gin.Context, url string, bearer string) {
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		error1.JSON(ctx, http.StatusNotFound, &error1.E{
			Type:        "END POINT ERROR",
			Description: "Some description goes here.",
		})
		return
	}

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		error1.JSON(ctx, http.StatusInternalServerError, &error1.E{
			Type:        "SERVER ERROR",
			Description: "Calling from thirdparty has error.",
		})
		return
	}
	defer resp.Body.Close()

	bufferedReader := bufio.NewReader(resp.Body)

	buffer := make([]byte, 4*1024)

	var totalBytesReceived int

	// Reads the response
	for {
		len, err := bufferedReader.Read(buffer)
		if len > 2 {
			totalBytesReceived += len
			log.Println(len, "bytes received")
			// Prints received data
			log.Println(string(buffer[:len]))
		}

		if err != nil {
			if err == io.EOF {
				// Last chunk received
				log.Println(err)
			}
			break
		}
	}
	//log.Println("Total Bytes Sent:", len(data))
	log.Println("Total Bytes Received:", totalBytesReceived)
}
