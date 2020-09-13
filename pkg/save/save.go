package save

import (
	"encoding/xml"
	"github.com/i-hit/go-lesson3.3.git/pkg/daily"
	"log"
)

type Currency struct {
	Code  string
	Name  string
	Value float64
}

func SaveData(extract []byte) (storage []Currency, err error) {

	var data *daily.Curses
	err = xml.Unmarshal(extract, &data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	buf := Currency{}
	for _, j := range data.Currencies {
		buf.Code = j.CharCode
		buf.Name = j.Name
		buf.Value = j.Value

		storage = append(storage, buf)
	}

	return storage, nil
}
