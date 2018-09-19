package main

import (
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func runClient() error {
	kubeConfigPath := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	// Create client config from ~/.kube/config file

	// CREATE CLIENT OMIT
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath) // HL1
	if err != nil {
		return err
	}

	clientset, err := kubernetes.NewForConfig(config) // HL2
	if err != nil {
		return err
	}
	// CREATE CLIENT OMIT

	// LIST PODS OMIT
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{}) // HL
	if err != nil {
		return err
	}
	// LIST PODS OMIT

	for _, p := range pods.Items {
		fmt.Printf("Namespace: %v\tPod: %v\n", p.Namespace, p.Name)
	}

	return nil
}

func main() {
	err := runClient()
	if err != nil {
		panic(err)
	}
}
