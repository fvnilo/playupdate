package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nylo-andry/playupdate"
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

func (p *PlayUpdateService) Update(mac string) (*playupdate.UpdateProfile, error) {
	url := fmt.Sprintf("%s/%s:%s", p.baseUrl, updatePath, mac)
	/*
		For the sake of the "demo" the same body is always sent but
		I am assuming that the values in this body should be obtained from
		another source
	*/
	body, _ := json.Marshal(NewUpdateProfile())
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		defer res.Body.Close()
		return handleResponse(res)
	}
}

func handleResponse(res *http.Response) (*playupdate.UpdateProfile, error) {
	if res.StatusCode == http.StatusOK {
		return unmarshallResponse(res.Body), nil
	}

	return nil, unmarshallApiError(res.Body)

}

func unmarshallResponse(r io.Reader) *playupdate.UpdateProfile {
	p := &playupdate.UpdateProfile{}
	json.NewDecoder(r).Decode(p)

	return p
}

func unmarshallApiError(r io.Reader) ApiError {
	err := &ApiError{}
	json.NewDecoder(r).Decode(err)

	return *err
}
