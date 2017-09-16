package main

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"io/ioutil"
	"encoding/json"
	json_response "json"
)

var cookieJar, _ = cookiejar.New(nil)

var client = &http.Client{
	Jar: cookieJar,
}

const string_url_api = "http://api.lingualeo.com/"
const string_url_common = "http://lingualeo.com/"

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

//TODO: urlBuilder for "string_url_common" and rename veriable
func (s* SimpleConnector) GetPageOfDictionary(index int) json_response.Dictionary {
	request_str := string_url_common + "userdict/json?groupId=dictionary&filter=learned"
	//ine := strconv.Itoa(index)
	request_args := url.Values{}
	request_args.Set("sortBy","date")
	request_args.Add("wordType","1")
	request_args.Add("filter","learned")
	request_args.Add("page","1")
	request_args.Add("groupId","dictionary")
	resp, err := client.PostForm(request_str, request_args)
	defer resp.Body.Close()
	var m json_response.LeoDictionaryImpl
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&m)
	if err != nil {
		print(err.Error())
	}
	return &m
}

func (s *SimpleConnector) AddWord(word, translate, context string) {
	request_str := string_url_api + "addword"
	request_args := url.Values{
		"word" : {word},
		"tword" : {translate},
		"context" : {context},
	}
	resp, err := client.PostForm(request_str, request_args)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	println(err)
	println(string(body))
}