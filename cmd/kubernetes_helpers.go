package main

import (
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/networking/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"context"
)


func kubernetes_new_config (cfg_path string) *rest.Config {
	cfg, err := clientcmd.BuildConfigFromFlags("", cfg_path) 

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return cfg
}


func kubernetes_new_clientset (cfg *rest.Config) *kubernetes.Clientset  {

	clientSet, err := kubernetes.NewForConfig(cfg)

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return clientSet
}


func (s *server) kubernetes_new_ingress() (*networkingv1.Ingress, error) {

	client, err := v1.NewForConfig(s.kconfig)

	if err != nil {
		s.LogError("kubernetes_new_ingress", err)
		return nil, err
	}

	ingress, err := client.Ingresses("default").Create(context.Background(), &networkingv1.Ingress{}, metav1.CreateOptions{})

	if err != nil {
		s.LogError("kubernetes_new_ingress", err)
		return nil, err
	}

	return ingress, nil
}
