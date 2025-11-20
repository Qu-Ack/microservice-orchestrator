package main

import (
	"net/http"
)

func (s *server) handlePostDeploy(w http.ResponseWriter, r *http.Request) {
	type Body struct {
		CloneUrl  string `json:"clone_url"`
		Branch    string `json:"branch"`
		Subdomain string `json:"subdomain"`
	}

	var b Body

	err := s.DecodeBody(r, &b)

	if err != nil {
		s.JSON(w, map[string]string{"error": "internal server error"}, 500)
		return
	}

}


