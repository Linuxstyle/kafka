package main

import (
	"KafkaLog/src/KafkaLag/GetCluster"
	"KafkaLog/src/KafkaLag/GetConsumers"
	getlag "KafkaLog/src/KafkaLag/GetLag"
	"fmt"
)

const (
	url string = "xxxxx:8000/v3/kafka/"
)

func main() {
	c, err := kafkaclusters.GetCluster(url)
	if err != nil {
		fmt.Println("get kafka cluster faild:", err)
	}
	cluster := c[0]
	data, err := getconsumers.GetConsumer(url, cluster)
	if err != nil {
		fmt.Println("get all consumers faild:", err)
	}
	//fmt.Println(data)
	for _, v := range data {
		lag, err := getlag.GetLag(url, cluster, v)
		if err != nil {
			fmt.Println("get lag faild:", err)
		}
		fmt.Println(lag)
	}
}
