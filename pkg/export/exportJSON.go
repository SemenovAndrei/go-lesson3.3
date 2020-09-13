package export

import (
	"encoding/json"
	"github.com/i-hit/go-lesson3.3.git/pkg/save"
	"io/ioutil"
	"log"
)

type Currency struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type Currencies struct {
	Currency []Currency
}

func Export(storage []save.Currency, filename string) error {

	encoded, err := json.MarshalIndent(storage, "", " ")
	if err != nil {
		log.Println(err)
		return err
	}

	err = ioutil.WriteFile(filename, encoded, 0777)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

