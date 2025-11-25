package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/intstr"
	"net/http"
	"path/filepath"
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

	id := random_string_from_charset(6)
	clone_path := make_clone_path(b.Subdomain, id)

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
			go func() {
				s.LogMsg(fmt.Sprintf("reading %v service", name))

				img_name := fmt.Sprintf("%v-%v-%v:latest", name, id, b.Subdomain)
				dep_name := fmt.Sprintf("%v-%v-%v-dep", name, id, b.Subdomain)
				pod_name := fmt.Sprintf("%v-%v-%v-pod", name, id, b.Subdomain)
				ser_name := fmt.Sprintf("%v-%v-%v-ser", name, id, b.Subdomain)

				pod_selector_labels := map[string]string{
					"app": pod_name,
				}

				containerPort, err := s.docker_get_container_port(filepath.Join(clone_path, service.Build.Context, service.Build.Dockerfile))

				if err != nil {
					s.LogError("go func docker get container port", err)
					return
				}

				err = s.docker_build_image(filepath.Join(clone_path, service.Build.Context), img_name)

				if err != nil {
					s.LogError("go func building image", err)
					return
				}

				var replicas int32 = 3
				_, err = s.kuberentes_new_deployment(dep_name, &replicas, pod_selector_labels, 80, img_name)

				if err != nil {
					s.LogError("go func building deployment", err)
					return
				}

				_, err = s.kuberentes_new_service(ser_name, pod_selector_labels, 80, intstr.FromInt(containerPort))

				if err != nil {
					s.LogError("go func building service", err)
					return
				}

				_, err = s.kubernetes_ingress_update(ser_name, b.Subdomain)

				if err != nil {
					s.LogError("go func updating ingress error", err)
					return
				}

			}()
		}

		s.LogMsg("completed building images..")
		s.JSON(w, map[string]string{"status": "ok"}, 200)

		break
	case "dockerfile":

		img_name := fmt.Sprintf("%v-%v", b.Subdomain, id)
		dep_name := fmt.Sprintf("%v-%v-dep", b.Subdomain, id)
		ser_name := fmt.Sprintf("%v-%v-ser", b.Subdomain, id)
		pod_name := fmt.Sprintf("%v-%v-%v-pod", id, b.Subdomain)

		pod_selector_labels := map[string]string{
			"app": pod_name,
		}

		containerPort, err := s.docker_get_container_port(filepath.Join(clone_path, "Dockerfile"))

		if err != nil {
			s.LogError("go func docker get container port", err)
			return
		}

		err = s.docker_build_image(clone_path, img_name)

		if err != nil {
			s.LogError("go func building image", err)
			return
		}

		var replicas int32 = 3
		_, err = s.kuberentes_new_deployment(dep_name, &replicas, pod_selector_labels, 80, img_name)

		if err != nil {
			s.LogError("go func building deployment", err)
			return
		}

		_, err = s.kuberentes_new_service(ser_name, pod_selector_labels, 80, intstr.FromInt(containerPort))

		if err != nil {
			s.LogError("go func building service", err)
			return
		}

		_, err = s.kubernetes_ingress_update(ser_name, b.Subdomain)

		if err != nil {
			s.LogError("go func updating ingress error", err)
			return
		}

		break

	default:
		s.JSON(w, map[string]string{"error": "invalid type"}, 401)
	}
}

type Body struct {
	DeploymentName *string `json:"deployment_name"`
	Replicas       *int32  `json:"replicas"`
	UpdatedName    *string `json:"updated_name"`
}

func (s *server) handlePutDeploy(w http.ResponseWriter, r *http.Request) {
	var b Body

	if err := s.DecodeBody(r, &b); err != nil {
		s.JSON(w, map[string]string{"error": "internal server error"}, 500)
		return
	}

	var partial PartialDeployment

	if b.Replicas != nil {
		partial.replicas = b.Replicas
	}
	if b.UpdatedName != nil {
		partial.name = b.UpdatedName
	}

	if err := s.kubernetes_update_deployment(*b.DeploymentName, partial); err != nil {
		s.LogError("handlePutDeploy", err)
		s.JSON(w, map[string]string{"error": "internal server error"}, 500)
		return
	}

	s.JSON(w, map[string]string{"status": "ok"}, 200)
}
