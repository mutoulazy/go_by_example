package main

import (
	"flag"
	"github.com/astaxie/beego/logs"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// 配置 k8s 集群外 kubeconfig 配置文件
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "./config/admin.conf", "absolute path to the kubeconfig file")
	flag.Parse()

	//在 kubeconfig 中使用当前上下文环境，config 获取支持 url 和 path 方式
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		logs.Error("构建kubeconfig 失败", err)
		panic(err.Error())
	}

	// 根据指定的 config 创建一个新的 clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logs.Error("创建k8s客户端 失败", err)
		panic(err.Error())
	}

	// 通过实现 clientset 的 CoreV1Interface 接口列表中的 PodsGetter 接口方法 Pods(namespace string)返回 PodInterface
	// PodInterface 接口拥有操作 Pod 资源的方法，例如 Create、Update、Get、List 等方法
	// 注意：Pods() 方法中 namespace 不指定则获取 Cluster 所有 Pod 列表
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		logs.Error("执行查询pod列表  失败", err)
		panic(err.Error())
	}
	logs.Debug("There are %d pods in the k8s 集群", len(pods.Items))

	// 指定获取namespace下的pod
	namespace := "admin"
	pods, err = clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	if err != nil {
		logs.Error("执行查询租户下的pod列表  失败", err)
		panic(err.Error())
	}
	logs.Debug("There are %d pods in the k8s %s", len(pods.Items), namespace)
	for _, pod := range pods.Items {
		logs.Debug("Name: %s, Status: %s, CreateTime: %s", pod.ObjectMeta.Name, pod.Status.Phase,
			pod.ObjectMeta.CreationTimestamp)
	}

	// 获取指定namespace和podname的 pod详细信息
	namespace = "default"
	podName := "vtctlclient-576d4f66bf-v4l75"
	pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		logs.Error("Pod %s in namespace %s not found", podName, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		logs.Error("Error getting pod %s in namespace %s: %v",
			podName, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		logs.Error("执行查询租户下的pod详细信息 失败", err)
	} else {
		maps := map[string]interface{}{
			"Name":        pod.ObjectMeta.Name,
			"Namespaces":  pod.ObjectMeta.Namespace,
			"NodeName":    pod.Spec.NodeName,
			"Annotations": pod.ObjectMeta.Annotations,
			"Labels":      pod.ObjectMeta.Labels,
			"SelfLink":    pod.ObjectMeta.SelfLink,
			"Uid":         pod.ObjectMeta.UID,
			"Status":      pod.Status.Phase,
			"IP":          pod.Status.PodIP,
			"Image":       pod.Spec.Containers[0].Image,
		}
		prettyPrint(maps)
	}

	// 创建一个namespace
	name := "client-go-test"
	namespacesClient := clientset.CoreV1().Namespaces()
	// 构建namespace结构
	newNamespace := &apiv1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Status: apiv1.NamespaceStatus{
			Phase: apiv1.NamespaceActive,
		},
	}
	resultNamespace, err := namespacesClient.Create(newNamespace)
	if err != nil {
		logs.Error("创建namespace失败", err)
	}
	logs.Debug("Created Namespaces %s on %s", resultNamespace.ObjectMeta.Name,
		resultNamespace.ObjectMeta.CreationTimestamp)

	// 获取创建namespace的信息
	resultNamespace, err = namespacesClient.Get(name, metav1.GetOptions{})
	if err != nil {
		logs.Error("获取namespace失败", err)
	}
	logs.Debug("Name: %s, Status: %s, selfLink: %s, uid: %s", resultNamespace.ObjectMeta.Name,
		resultNamespace.Status.Phase, resultNamespace.ObjectMeta.SelfLink, resultNamespace.ObjectMeta.UID)

	// 删除namespace
	deletePolicy := metav1.DeletePropagationForeground
	if err := namespacesClient.Delete(name, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		logs.Error("删除namespace失败", err)
	}
	logs.Debug("Deleted Namespaces %s", name)
}

/**
	格式化输出
*/
func prettyPrint(maps map[string]interface{}) {
	lens := 0
	for k, _ := range maps {
		if lens <= len(k) {
			lens = len(k)
		}
	}
	for key, values := range maps {
		spaces := lens - len(key)
		v := ""
		for i := 0; i < spaces; i++ {
			v += " "
		}
		logs.Debug("%s: %s%v", key, v, values)
	}
}
