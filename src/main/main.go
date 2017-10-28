package main

import (
	"connection"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"fmt"
	"encoding/json"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(ctx, "{ name: 123 }")
}

func getAllWords(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
	email := ctx.QueryArgs().Peek("email")
	pass := ctx.QueryArgs().Peek("pass")
	sp := connection.SimpleConnector{string(email), string(pass)}
	sp.Connect()
	dict := sp.GetAllWords()
	res_json, err := json.Marshal(&dict)
	if err != nil {
		panic("marshal error")
	}
	fmt.Fprintf(ctx, string(res_json))
}

func main(){
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/get_all_words", getAllWords)

	log.Fatal(fasthttp.ListenAndServe(":4000", router.Handler))
	return
}