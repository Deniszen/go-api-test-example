package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func ExecuteGet(url string) (resp *http.Response, bodyBytes []byte) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ExecutePost(url string, contentType string, body io.Reader) (resp *http.Response, bodyBytes []byte) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func ExecuteDelete(url string) (resp *http.Response) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	resp, err = client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	return
}
