package main

import (
	"context"
	"net/http"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)


func (s *server) handleGetService(w http.ResponseWriter, r *http.Request) {

	namespaceName, ok := r.Context().Value("namespace_id").(string)
	if !ok {
		s.JSON(w, map[string]string{"error": "forbidden"}, 401)
		return
	}

	svcName := r.PathValue("svc")

	svc, err := s.kubernetes_get_servicee(namespaceName, svcName)
	if err != nil {
		s.LogError("handleGetService", err)
		s.JSON(w, map[string]string{"error": "internal server error"}, 500)
		return
	}

	resp := map[string]any{
		"name":      svc.Name,
		"namespace": svc.Namespace,
		"type":      svc.Spec.Type,
		"clusterIP": svc.Spec.ClusterIP,
		"ports":     svc.Spec.Ports,
		"selector":  svc.Spec.Selector,
	}

	podsInfo := []map[string]any{}
	if len(svc.Spec.Selector) > 0 {

		labelSelector := labels.SelectorFromSet(svc.Spec.Selector).String()

		pods, err := s.kclient.CoreV1().Pods(namespaceName).List(
			context.Background(),
			metav1.ListOptions{LabelSelector: labelSelector},
		)

		if err != nil {
			s.LogError("handleGetService: list pods", err)

		} else {
			for _, p := range pods.Items {
				podsInfo = append(podsInfo, map[string]any{
					"name":     p.Name,
					"phase":    p.Status.Phase,
					"podIP":    p.Status.PodIP,
					"nodeName": p.Spec.NodeName,
				})
			}
		}
	}

	resp["pods"] = podsInfo

	var deploymentName string
	if strings.HasSuffix(svcName, "-ser") {
		deploymentName = strings.TrimSuffix(svcName, "-ser") + "-dep"
	} else {
		deploymentName = svcName + "-dep" 
	}

	dep, err := s.kclient.AppsV1().Deployments(namespaceName).Get(
		context.Background(), deploymentName, metav1.GetOptions{},
	)

	if err == nil {
		containers := []map[string]string{}
		for _, c := range dep.Spec.Template.Spec.Containers {
			containers = append(containers, map[string]string{
				"name":  c.Name,
				"image": c.Image,
			})
		}

		resp["deployment"] = map[string]any{
			"name":              dep.Name,
			"replicas":          dep.Status.Replicas,
			"readyReplicas":     dep.Status.ReadyReplicas,
			"availableReplicas": dep.Status.AvailableReplicas,
			"labels":            dep.Labels,
			"containers":        containers,
		}
	} else {
		resp["deployment"] = nil
	}

	s.JSON(w, resp, 200)
}

