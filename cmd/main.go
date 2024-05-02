package main

import (
	"log"
	"testB/pkg/optional"

	"testB/pkg/api"
	"testB/pkg/repository"

	"github.com/gorilla/mux"
)

func main() {
	connStr, fileDirect, fileDirectPdf, err := optional.ReadConf("conf.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}
	api := api.New(mux.NewRouter(), db)
	api.FillEndpoints()
	go optional.ParseAndCheckTcv(fileDirect, fileDirectPdf,db)
	log.Fatal(api.ListenAndServe("localhost:8090"))
}
