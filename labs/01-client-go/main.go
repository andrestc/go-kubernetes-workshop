package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func runClient() error {
	kubeConfigPath := flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "kube config path")
	flag.Parse()

	// CREATE CLIENT OMIT
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfigPath) // HL1
	if err != nil {
		return err
	}
	clientset, err := kubernetes.NewForConfig(config) // HL2
	if err != nil {
		return err
	}
	// CREATE CLIENT OMIT

	_, err = clientset.CoreV1().Pods("default").Get("myotherpod", metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			if _, errCreate := clientset.CoreV1().Pods("default").Create(&v1.Pod{
				ObjectMeta: metav1.ObjectMeta{Name: "myotherpod"},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{Name: "cont", Image: "nginx"},
					},
				},
			}); errCreate != nil {
				return errCreate
			}
		}
	}

	// LIST PODS OMIT
	for {
		pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{}) // HL
		if err != nil {
			return err
		}
		// LIST PODS OMIT

		for _, p := range pods.Items {
			fmt.Printf("Namespace: %v\tPod: %v\n", p.Namespace, p.Name)
		}
		time.Sleep(10 * time.Second)
	}

	return nil
}

func main() {
	err := runClient()
	if err != nil {
		panic(err)
	}
}
