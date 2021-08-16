package adapter

import "fmt"

type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

type AWSClient struct{}

func (aws *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run at cpu: %f, mem: %f\n", cpu, mem)
	return nil
}

type AWSClientAdapter struct {
	Client AWSClient
}

func (adapter *AWSClientAdapter) CreateServer(cpu, mem float64) error {
	return adapter.Client.RunInstance(cpu, mem)
}

type AliyunClient struct{}

func (aliyun *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aliyun client run at cpu: %d, mem: %d", cpu, mem)
	return nil
}

type AliyunClientAdapter struct {
	Client AliyunClient
}

func (adapter *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	return adapter.Client.CreateServer(int(cpu), int(mem))
}
