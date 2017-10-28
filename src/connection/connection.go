package connection

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"encoding/json"
	json_response "json"
	"json/lingualeo"
	"strconv"
)

var cookieJar, _ = cookiejar.New(nil)

var client = &http.Client{
	Jar: cookieJar,
}

const string_url_api = "http://api.lingualeo.com/"
const string_url_common = "http://lingualeo.com/"

type Connector interface {
	Connect()
}

type SimpleConnector struct {
	Login string
	Pass  string
}

func (s *SimpleConnector) SetPass(pass string) {
	s.Pass = pass
}

func (s *SimpleConnector) SetLogin(login string) {
	s.Login = login
}

func (s *SimpleConnector) Connect() {
	auth_request := string_url_api + "api/login"
	auth_params := url.Values{
		"email":    {s.Login},
		"password": {s.Pass},
	}
	resp, err := client.PostForm(auth_request, auth_params)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
}

//TODO: urlBuilder for "string_url_common" and rename veriable
func (s *SimpleConnector) GetPageOfDictionary(index int) json_response.Dictionary {
	request_str := string_url_common + "userdict/json?groupId=dictionary&filter=learned&page=" + strconv.Itoa(index) // NOTE: Use only this string
	request_args := url.Values{}
	resp, err := client.PostForm(request_str, request_args)
	defer resp.Body.Close()
	var m lingualeo.LeoDictionaryImpl
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&m)
	if err != nil {
		panic(err.Error())
	}
	return &m
}

func (s *SimpleConnector) GetAllWords() []json_response.Word {
	var result []json_response.Word

	for i := 0; ; i++ {
		dict := s.GetPageOfDictionary(i)
		arrayWords := dict.GetWords()
		if len(arrayWords) == 0 {
			break
		}
		result = append(result, arrayWords...)
	}

	return result
}

func (s *SimpleConnector) AddWord(word, translate, context string) {
	request_str := string_url_api + "addword"
	request_args := url.Values{
		"word":    {word},
		"tword":   {translate},
		"context": {context},
	}
	resp, err := client.PostForm(request_str, request_args)
	defer resp.Body.Close()
	if err != nil {
		panic(err.Error())
	}
}
