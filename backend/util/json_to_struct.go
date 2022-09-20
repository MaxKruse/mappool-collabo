package util

import (
	"backend/models"
	"encoding/json"
	"io"
)

type convertible interface {
	models.BanchoUserResponse
}

func Convert[k convertible](r io.ReadCloser) k {
	var v k
	json.NewDecoder(r).Decode(&v)
	return v
}
