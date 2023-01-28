package k8s

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func LeadConfigMap() {
	// k8sクラスタに接続するためのconfigを作成
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatalf("Unable to get our client configuration: %v", err)
	}

	// k8sクラスタに接続
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Unable to create our clientset: %v", err)
	}

	// configmapを取得
	configMap, err := clientset.CoreV1().ConfigMaps("go-app").Get(context.TODO(), "sample-app", metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Failed to get configmap: %v", err)
	}

	fmt.Println(configMap)

	// configmapからクラスターIDを取得
	clusterID := configMap.Data["clusterID"]
	fmt.Println(clusterID)
}
