package main

import (
	"github.com/i-hit/go-lesson3.3.git/pkg/daily"
	"github.com/i-hit/go-lesson3.3.git/pkg/export"
	"github.com/i-hit/go-lesson3.3.git/pkg/save"
	"log"
	"net/http"
	"os"
)

func main() {
	const JSONFile = "export.json"
	url := "https://raw.githubusercontent.com/netology-code/bgo-homeworks/master/10_client/assets/daily.xml"

	if err := execute(url, JSONFile); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}

func execute(url string, filename string) (err error) {

	svc := daily.NewService(url, &http.Client{})

	extract, err := svc.Extract()
	if err != nil {
		log.Println(err)
		return err
	}

	storage, err := save.SaveData(extract)
	if err != nil {
		log.Println(err)
		return err
	}

	err = export.Export(storage, filename)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}