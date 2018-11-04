package kafkaclusters

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Cluster struct {
	Error    bool              `json:"error"`
	Message  string            `json:"message"`
	Clusters []string          `json:"clusters"`
	Request  map[string]string `json:"request"`
}

const (
	url string = "http://10.131.7.31:8000/v3/kafka/"
)

func GetCluster(url string) (clusters []string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("requet http failed:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read http respone faild:", err)
		return
	}
	//fmt.Println(string(body))
	var c *Cluster
	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println("json data faiild:", err)
		return
	}
	clusters = c.Clusters
	return
}
