package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const updatePath = "/profile/clientId"

type PlayUpdateService struct {
	baseUrl string
}

func NewPlayUpdateService(baseUrl string) *PlayUpdateService {
	return &PlayUpdateService{
		baseUrl: baseUrl,
	}
}

func (p *PlayUpdateService) Update(mac string) error {
	url := fmt.Sprintf("%s/%s:%s", p.baseUrl, updatePath, mac)
	body, _ := json.Marshal(NewRequestBody())
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	} else {
		defer res.Body.Close()

		fmt.Printf("Done processing mac %s\n", mac)
		return nil
	}
}
