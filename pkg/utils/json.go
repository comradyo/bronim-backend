package utils

import (
	"encoding/json"
	"errors"
	"io"
)

var (
	ErrJSONDecoding = errors.New("data decoding error")
	ErrJSONEncoding = errors.New("data encoding error")
)

func GetObjectFromRequest(r io.Reader, obj interface{}) error {
	err := json.NewDecoder(r).Decode(obj)
	if err != nil {
		return ErrJSONDecoding
	}
	return nil
}

func Marshall(body interface{}) ([]byte, error) {
	res, err := json.Marshal(body)
	if err != nil {
		return nil, ErrJSONEncoding
	}
	return res, nil
}
