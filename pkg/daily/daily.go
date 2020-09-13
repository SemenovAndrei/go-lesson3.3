package daily

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
)

type Service struct {
	url     string
	client  *http.Client
}

func NewService(
	url string,
	client *http.Client,
) *Service {
	return &Service{
		url: url, client: client,
	}
}

func (s *Service) Extract() (data []byte, err error) {
	ctx := context.Background()
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		s.url,
		nil,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := s.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer func() {
		if cerr := resp.Body.Close(); cerr != nil {
			if err == nil {
				log.Println(err)
				err = cerr
			}
		}
	}()

	return respBody, err
}

