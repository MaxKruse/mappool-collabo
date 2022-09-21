package util

import (
	"backend/models"
	"encoding/json"
	"io"
)

type convertible interface {
	models.BanchoUserResponse
}

func Convert[k convertible](r io.ReadCloser) (k, error) {
	var v k
	err := json.NewDecoder(r).Decode(&v)
	return v, err
}
