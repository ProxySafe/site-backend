package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func SetRequestDto(r *http.Request, dto any) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(reqBody, dto); err != nil {
		return err
	}

	return nil
}
