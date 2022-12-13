package k8s

import (
	"bufio"
	"context"
	"io"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
)

type PodClient struct {
	config    *restclient.Config
	clientset *kubernetes.Clientset
}

func NewPodClient(config *restclient.Config, clientset *kubernetes.Clientset) *PodClient {
	return &PodClient{
		config:    config,
		clientset: clientset,
	}
}

func (cli *PodClient) Get(name, namespace string) (*corev1.Pod, error) {
	opts := metav1.GetOptions{}
	return cli.clientset.CoreV1().Pods(namespace).Get(context.Background(), name, opts)
}

func (cli *PodClient) Logs(name, namespace string, opts *corev1.PodLogOptions) *restclient.Request {
	return cli.clientset.CoreV1().Pods(namespace).GetLogs(name, opts)
}

func (cli *PodClient) LogStream(name, namespace string, opts *corev1.PodLogOptions, out io.Writer) error {
	req := cli.Logs(name, namespace, opts)
	stream, err := req.Stream(context.Background())
	if err != nil {
		return err
	}
	defer stream.Close()

	buffer := bufio.NewReader(stream)
	for {
		bytes, err := buffer.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				_, err = out.Write(bytes)
			}
			return err
		}

		_, err = out.Write(bytes)
		if err != nil {
			return err
		}
	}
}
