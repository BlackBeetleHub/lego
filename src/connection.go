package main

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

}
