package main

import (
	"archive/tar"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/go-yaml/yaml"
	"io"
	"os"
	"path/filepath"
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

func docker_new_client() *client.Client {
	apiClient, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)

	}
	return apiClient
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

func (s *server) docker_create_context(dirPath string) (io.Reader, error) {
	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	defer tw.Close()

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(dirPath, path)
		if err != nil {
			return err
		}

		if rel == "." || rel == ".." {
			return nil
		}

		header, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		header.Name = filepath.ToSlash(rel)

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err = io.Copy(tw, f); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		s.LogError("docker_create_context", err)
		return nil, errors.New("error occured while creating context")
	}

	return bytes.NewReader(buf.Bytes()), nil
}

func (s *server) docker_build_image(root_dir string, service_name string) error {
	build_context, err := s.docker_create_context(root_dir)

	if err != nil {
		s.LogError("docker_build_image", err)
		return err

	}

	image_build_resp, err := s.dclient.ImageBuild(context.Background(), build_context, types.ImageBuildOptions{
		Tags:        []string{fmt.Sprintf("service-%v:latest", service_name)},
		Remove:      true,
		ForceRemove: true,
	})

	if err != nil {
		s.LogError("docker_build_image", err)
		return errors.New("could not build image")
	}

	_, err = io.Copy(os.Stdout, image_build_resp.Body)
	defer image_build_resp.Body.Close()

	if err != nil {
		s.LogError("docker_build_image", err)
		return errors.New("could not write image output")

	}
	return nil
}
