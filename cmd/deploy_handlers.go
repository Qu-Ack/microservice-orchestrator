package main

import (
	"path/filepath"
	"net/http"
	"fmt"
)

func (s *server) handlePostDeploy(w http.ResponseWriter, r *http.Request) {
	s.LogRequest(r)

	type Body struct {
		CloneUrl  string `json:"clone_url"`
		Branch    string `json:"branch"`
		Subdomain string `json:"subdomain"`
		Type      string `json:"type"`
	}

	var b Body

	err := s.DecodeBody(r, &b)

	if err != nil {
		s.JSON(w, map[string]string{"error": "internal server error"}, 500)
		return
	}

	clone_path := make_clone_path(b.Subdomain)

	err = git_clone(b.CloneUrl, b.Branch, "1", b.Subdomain, clone_path)

	if err != nil {
		s.JSON(w, map[string]string{"error": "internal server error", "message": "couldn't clone repo"}, 500)
		return
	}

	switch b.Type {

	case "docker_compose":

		compose_file, err := s.read_compose_file(clone_path)
		if err != nil {
			s.JSON(w, map[string]string{"error": "internal server error", "message": "couldn't read compose_file"}, 500)
			return
		}

		for name, service := range compose_file.Services {
			s.LogMsg(fmt.Sprintf("reading %v service", name))

			err := s.docker_build_image(filepath.Join(clone_path, service.Build.Context))

			if err != nil {
				s.JSON(w, map[string]string{"error": "internal server error", "message": "couldn't clone repo"}, 500)
				return
			}
		}

		s.LogMsg("completed building images..")

		s.JSON(w, map[string]string{"status": "ok"}, 200)
		break
	case "dockerfile":
		break

	default:
		s.JSON(w, map[string]string{"error": "invalid type"}, 401)
	}
}
