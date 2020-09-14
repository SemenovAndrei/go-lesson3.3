package export

import (
	"encoding/json"
	"github.com/i-hit/go-lesson3.3.git/pkg/save"
	"io/ioutil"
	"log"
)

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

