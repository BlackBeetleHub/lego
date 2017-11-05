package connection

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"encoding/json"
	json_response "json"
	"json/lingualeo"
	"strconv"
	"time"
)

var cookieJar, _ = cookiejar.New(nil)

var client = &http.Client{
	Jar: cookieJar,
}

const string_url_api = "http://api.lingualeo.com/"
const string_url_common = "http://lingualeo.com/"
const count_words_in_page = 100

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

func (s *SimpleConnector) GetCountWords() int {
	return s.GetPageOfDictionary(0).GetCountWords()
}

func (s *SimpleConnector) GetAllWords() []json_response.Word {
	start := time.Now()
	println(start.String())
	var result []json_response.Word
	chRes := make(chan * []json_response.Word)

	countWords := s.GetCountWords()
	println(countWords)
	countPages := (countWords / count_words_in_page) + 1;

	for i := 1; i <= countPages; i++ {
		go func(i int) {
			tmpArray := new([]json_response.Word)
			println(i)
			*tmpArray = s.GetPageOfDictionary(i).GetWords()
			chRes <- tmpArray
		}(i)
	}

	for i:= 0; i < countPages; i++ {
		tmpData := <-chRes
		result = append(result, *tmpData...);
	}
	end := time.Now()
	println(end.String())
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
