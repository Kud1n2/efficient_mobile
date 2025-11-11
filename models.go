package main

type URLrequest struct {
	Links []string `json:"links"`
}

type URLresponse struct {
	Links     map[string]string `json:"links"`
	Links_num int               `json:"links_num"`
}

type ListsRequest struct {
	Links_list []int `json:"links_list"`
}

var URLrequests []URLrequest
var URLresponses []URLresponse
var ListsRequests []ListsRequest
