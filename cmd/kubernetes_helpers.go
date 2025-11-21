package main

import (
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/networking/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"errors"
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
		return nil, errors.New("error occured while creating config")
	}

	pathType := networkingv1.PathTypePrefix
	ingress, err := client.Ingresses("default").Create(context.Background(), &networkingv1.Ingress{
		metav1.TypeMeta{
			Kind: "ingress",	
			APIVersion: "networking.k8s.io/v1",
		},
		metav1.ObjectMeta{
			Name: "simple-ingress",
		},
		networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{

					Host: "foo.bar.com",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									PathType: &pathType,
									Path: "/",
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: "web-service",
											Port: networkingv1.ServiceBackendPort{
												Number: 8080,
											},
										},
									},
								},
							},

						},
					},
				},
			},
		},
		networkingv1.IngressStatus{},
	}, metav1.CreateOptions{})

	if err != nil {
		s.LogError("kubernetes_new_ingress", err)
		return nil, errors.New("error occured while creating ingress")
	}

	return ingress, nil
}

