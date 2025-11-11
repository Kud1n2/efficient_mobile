package main

import (
	"bytes"
	"io"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func handlerRequest(context *gin.Context) {
	body, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Reading body error"})
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(body)) //создание копии запроса
	var requestList ListsRequest
	if err := context.BindJSON(&requestList); err == nil && len(requestList.Links_list) > 0 {
		ListsRequests = append(ListsRequests, requestList)
		context.IndentedJSON(http.StatusOK, ListsRequests)
		return
	}

	context.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	var requestURL URLrequest
	if err := context.BindJSON(&requestURL); err == nil && len(requestURL.Links) > 0 {
		URLrequests = append(URLrequests, requestURL)
		CheckAvailable()
		context.IndentedJSON(http.StatusOK, URLresponses[len(URLresponses)-1])
		return
	}
	context.JSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
}

func getURLs(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, URLresponses)
}

func CheckAvailable() {
	dict := make(map[string]string)
	if len(URLrequests) > 0 {
		used := URLrequests[len(URLrequests)-1]
		for i := 0; i < len(used.Links); i++ {
			_, err := net.Dial("tcp", used.Links[i]+":http")
			if err != nil {
				dict[used.Links[i]] = "not available"
			} else {
				dict[used.Links[i]] = "available"
			}

		}
		URLresponses = append(URLresponses, URLresponse{Links: dict, Links_num: len(URLrequests)})
	}
}
