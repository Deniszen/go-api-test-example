package main

import (
	"io"
	"io/ioutil"
	"net/http"
)

func executeGet(url string) (*http.Response, []byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return resp, bodyBytes, nil
}

func executePost(url string, contentType string, body io.Reader) (*http.Response, []byte, error) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	return resp, bodyBytes, nil
}

func executeDelete(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return resp, nil
}
