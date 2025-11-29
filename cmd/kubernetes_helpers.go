package main

import (
	"context"
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	typesv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/kubernetes/typed/networking/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func kubernetes_new_config(cfg_path string) *rest.Config {
	cfg, err := clientcmd.BuildConfigFromFlags("", cfg_path)

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return cfg
}

func kubernetes_new_clientset(cfg *rest.Config) *kubernetes.Clientset {
	clientSet, err := kubernetes.NewForConfig(cfg)

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return clientSet
}

func kubernetes_new_ingress(cfg *rest.Config, namespace_name string) (*networkingv1.Ingress, error) {
	client, err := v1.NewForConfig(cfg)

	if err != nil {
		panic(err.Error())
	}

	pathType := networkingv1.PathTypePrefix
	ingress, err := client.Ingresses(namespace_name).Create(context.Background(), &networkingv1.Ingress{
		metav1.TypeMeta{
			Kind:       "ingress",
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
									Path:     "/",
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
		log.Println(err.Error())
		return nil, err
	}

	return ingress, nil
}

func (s *server) kuberentes_new_deployment(dep_name string, replicas *int32, pod_selector_labels map[string]string, port int32, image_name string, namespace_name string) (*appsv1.Deployment, error) {

	client, err := typesv1.NewForConfig(s.kconfig)

	if err != nil {
		s.LogError("kubernetes_new_deployment", err)
		return nil, err
	}

	dep, err := client.Deployments(namespace_name).Create(context.Background(), &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: dep_name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: pod_selector_labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: pod_selector_labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            dep_name,
							ImagePullPolicy: corev1.PullNever,
							Image:           image_name,
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

func (s *server) kuberentes_new_service(service_name string, selector map[string]string, port int32, target_port intstr.IntOrString, namespace_name string) (*corev1.Service, error) {
	corev1_client, err := typedcorev1.NewForConfig(s.kconfig)

	if err != nil {
		s.LogError("kubernetes_new_service", err)
		return nil, err
	}

	ser, err := corev1_client.Services(namespace_name).Create(context.Background(), &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: service_name,
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceTypeNodePort,
			Selector: selector,
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       port,
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

type PartialIngressUpdateStruct struct {
}

// make this go routine safe
func (s *server) kubernetes_ingress_update(service_name, host_name string, namespace_name string) (*networkingv1.Ingress, error) {
	client, err := v1.NewForConfig(s.kconfig)
	if err != nil {
		return nil, err
	}

	// lock the mutex
	s.mu.Lock()
	current, err := client.Ingresses(namespace_name).Get(context.Background(), "simple-ingress", metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	pathType := networkingv1.PathTypePrefix
	current.Spec.Rules = append(current.Spec.Rules, networkingv1.IngressRule{
		Host: fmt.Sprintf("%v-%v.bar.com", service_name, host_name),
		IngressRuleValue: networkingv1.IngressRuleValue{
			HTTP: &networkingv1.HTTPIngressRuleValue{
				Paths: []networkingv1.HTTPIngressPath{
					{
						Path:     "/",
						PathType: &pathType,
						Backend: networkingv1.IngressBackend{
							Service: &networkingv1.IngressServiceBackend{
								Name: service_name,
								Port: networkingv1.ServiceBackendPort{Number: 80},
							},
						},
					},
				},
			},
		},
	})

	updated, err := client.Ingresses(namespace_name).Update(context.Background(), current, metav1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	s.mu.Unlock()
	// free the mutex
	return updated, nil
}

type PartialDeployment struct {
	replicas *int32
	name     *string
}

func (s *server) kubernetes_update_deployment(deployment_name string, part_dep PartialDeployment, namespace_name string) error {

	client, err := typesv1.NewForConfig(s.kconfig)

	if err != nil {
		s.LogError("kubernetes_update_deployment", err)
		return err
	}

	s.mu.Lock()

	dep, err := client.Deployments(namespace_name).Get(context.Background(), deployment_name, metav1.GetOptions{})

	if err != nil {
		s.LogError("kubernetes_update_deployment", err)
		return err
	}

	if part_dep.name != nil {
		dep.ObjectMeta.Name = *part_dep.name
	}

	if part_dep.replicas != nil {
		dep.Spec.Replicas = part_dep.replicas
	}

	dep, err = client.Deployments(namespace_name).Update(context.Background(), dep, metav1.UpdateOptions{})

	if err != nil {
		s.LogError("kubernetes_update_deployment", err)
		return err
	}

	return nil
}

// TODO: Implement AutoScale Feature.

func (s *server) kubernetes_create_namespace(namespace_name string) (*corev1.Namespace, error) {
	namespace, err := s.kclient.CoreV1().Namespaces().Create(context.Background(), &corev1.Namespace{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Namespace",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: namespace_name,
		},
	}, metav1.CreateOptions{})

	if err != nil {
		s.LogError("kubernetes_create_namespace", err)
		return nil, err
	}

	return namespace, nil
}

func (s *server) kubernetes_list_services(namespace_name string) (*corev1.ServiceList, error) {

	services, err := s.kclient.CoreV1().Services(namespace_name).List(context.Background(), metav1.ListOptions{})

	if err != nil {
		s.LogError("kubernetes_list_services", err)
		return nil, err
	}

	return services, nil

}
func (s *server) kubernetes_delete_namespace(namespace_name string) error {
	return s.kclient.CoreV1().Namespaces().Delete(context.Background(), namespace_name, metav1.DeleteOptions{})
}

func (s *server) kubernetes_delete_ingress(namespace_name string, ingress_name string) error {
	return s.kclient.NetworkingV1().Ingresses(namespace_name).Delete(context.Background(), ingress_name, metav1.DeleteOptions{})
}
