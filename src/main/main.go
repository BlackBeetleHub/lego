package main


func main(){
	sp := SimpleConnector{"deniskozlov2012@mail.ru", "7572836d"}
	sp.Connect()
	sp.GetPageOfDictionary(1)
	return
}