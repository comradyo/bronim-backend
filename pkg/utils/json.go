package utils

import (
	log "bronim/pkg/logger"
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
		log.ErrorAtFunc(GetObjectFromRequest, err)
		return ErrJSONDecoding
	}
	return nil
}

type _body struct {
	Body interface{} `json:"body"`
}

func Marshall(body interface{}) ([]byte, error) {
	jsonBody := _body{Body: body}
	res, err := json.Marshal(jsonBody)
	if err != nil {
		return nil, ErrJSONEncoding
	}
	return res, nil
}
