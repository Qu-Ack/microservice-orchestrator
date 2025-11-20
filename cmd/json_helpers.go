package main

import (
	"net/http"
	"encoding/json"
	"io"
)


func (s *server) JSON(w http.ResponseWriter, msg any, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	data, err := json.Marshal(msg)

	if err != nil {
		s.LogError("JSON", err)
		return err
	}
	w.Write(data)
	return nil
}


func (s *server) UnmarshalBody(r *http.Request, v any) error {

	data, err := io.ReadAll(r.Body)

	if err != nil {
		s.LogError("UnmarshalBody", err)
		return err
	}

	err = json.Unmarshal(data, v)

	if err != nil {
		s.LogError("UnmarshalBody", err)
		return err
	}

	return nil
}

func (s *server) DecodeBody(r *http.Request, v any) error {
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(v)

	if err != nil {
		s.LogError("DecodeBody", err)
		return err
	}

	return nil
}
