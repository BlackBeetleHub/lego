package main

//import json_response "json"

func main(){
	sp := SimpleConnector{"deniskozlov2012@mail.ru", "7572836d"}
	sp.Connect()
	dict := sp.GetPageOfDictionary(1)

	println(dict.GetCountWords())
	arrayWords := dict.GetWords()
	for _, v := range arrayWords {
		println(v.GetWord())
	}
	return
}