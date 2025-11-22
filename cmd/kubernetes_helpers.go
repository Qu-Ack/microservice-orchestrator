package main

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/apimachinery/pkg/util/intstr"
	"log"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	corev1 "k8s.io/api/core/v1"
	typesv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/kubernetes/typed/networking/v1"
	appsv1 "k8s.io/api/apps/v1"
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


func (s *server) kuberentes_new_deployment(dep_name string, labels map[string]string, replicas *int32, pod_labels map[string]string, port int32, image_name string) (*appsv1.Deployment, error)  {

	client, err := typesv1.NewForConfig(s.kconfig)

	if err != nil {
		s.LogError("kubernetes_new_deployment", err)
		return nil, err
	}

	dep, err := client.Deployments("default").Create(context.Background(), &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind: "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: dep_name,
			Labels: labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: pod_labels, 
			},
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name: dep_name,
							Image: image_name,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: port,
								},
							},
						},
					},
				},

			}, 
		},
		
	}, metav1.CreateOptions{})

	if err != nil {
		s.LogError("kubernetes_new_deployment", err)
		return nil, err
	}

	return dep, err
}


func (s *server) kuberentes_new_service(service_name string, selector map[string]string, port int32, target_port intstr.IntOrString) (*corev1.Service, error) {
	corev1_client, err := typedcorev1.NewForConfig(s.kconfig)

	if err != nil {
		s.LogError("kubernetes_new_service", err)
		return nil, err
	}

	ser , err := corev1_client.Services("default").Create(context.Background(), &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind: "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: service_name,
		},
		Spec: corev1.ServiceSpec{
			Selector: selector,
			Ports: []corev1.ServicePort{
				{
					Protocol: corev1.ProtocolTCP,
					Port: port,
					TargetPort: target_port, 
				},
			},
		},

	}, metav1.CreateOptions{})

	if err != nil {
		s.LogError("kubernetes_new_service", err)
		return nil, err
	}

	return ser, nil
}

