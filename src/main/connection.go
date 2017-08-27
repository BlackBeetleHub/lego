package main

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"io/ioutil"
)

var cookieJar, _ = cookiejar.New(nil)

var client = &http.Client{
	Jar: cookieJar,
}

const string_url_api = "http://api.lingualeo.com/"

type Connector interface{
	Connect()
}

type SimpleConnector struct {
	login string
	pass  string
}

func (s *SimpleConnector) SetPass(pass string) {
	s.pass = pass
}

func (s *SimpleConnector) SetLogin(login string) {
	s.login = login
}

func (s *SimpleConnector) Connect() {
	auth_request := string_url_api + "api/login"
	auth_params := url.Values{
		"email" : {s.login},
		"password" : {s.pass},
	}
	resp, err := client.PostForm(auth_request,auth_params)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	println(err)
	println(string(body))
}

func (s *SimpleConnector) AddWord(word, translate, context string) {
	auth_request := string_url_api + "addword"
	request := url.Values{
		"word" : {word},
		"tword" : {translate},
		"context" : {context},
	}
	resp, err := client.PostForm(auth_request,request)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	println(err)
	println(string(body))
}