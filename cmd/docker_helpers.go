package main

import (
	"os"
	"path/filepath"
	"github.com/go-yaml/yaml"
	"io"
)

type Service struct {
	ContainerName string `yaml:"container_name"`
	Build         Build
}

type Build struct {
	Context    string `yaml:"context"`
	Dockerfile string `yaml:"dockerfile"`
}

type ComposeFile struct {
	Services map[string]Service
}

func (s *server) read_compose_file(project_path string) (*ComposeFile, error) {
	var c ComposeFile

	f, err := os.Open(filepath.Join(project_path, "docker-compose.yml"))

	if err != nil {
		s.LogError("read_compose_file", err)
		return nil, err
	}

	data, err := io.ReadAll(f)

	if err != nil {
		s.LogError("read_compose_file", err)
		return nil, err
	}


	err = yaml.Unmarshal(data, &c)

	if err != nil {
		s.LogError("read_compose_file", err)
		return nil, err
	}

	return &c, nil

}
