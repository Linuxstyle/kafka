package getlag

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StartEnd struct {
	Offset    int32 `json:"offset"`
	timestamp int64 `json:"timestamp"`
	Lag       int64 `json:"lag"`
}

type Partition struct {
	Topic      string   `json:"topic"`
	Partition  int32    `json:"partition"`
	Owner      string   `json:"owner"`
	Status     string   `json:"status"`
	Start      StartEnd `json:"start"`
	End        StartEnd `json:"end"`
	Currentlag int32    `json:"current_lag"`
	Complete   int64    `json:"complete"`
}

type Requests struct {
	Url  string `json:"url"`
	Host string `json:"host"`
}
type Sta struct {
	Cluster        string      `json:"cluster"`
	Group          string      `json:"group"`
	Statu          string      `json:"status"`
	Complete       int         `json:"complete"`
	Partitions     []Partition `json:"partitions"`
	Partitioncount int         `json:"partition_count"`
	Maxlag         Partition   `json:"maxlag"`
	Totallag       int64       `json:"totallag"`
}

type GetResp struct {
	Error   bool     `json:"error"`
	Message string   `json:"message"`
	Status  Sta      `json:"status"`
	Request Requests `json:"requests"`
}

func GetLag(url, cluster, consumer string) (Lag []Partition, err error) {
	Lurl := fmt.Sprintf(url + "/" + cluster + "/" + consumer + "/lag")
	resp, err := http.Get(Lurl)
	if err != nil {
		fmt.Println("get lag data faild:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read laog resp data faild:", err)
		return
	}
	var c *GetResp
	err = json.Unmarshal(body, c)
	if err != nil {
		fmt.Println("json lag faild:", err)
		return
	}
	Lag = c.Status.Partitions
	return
}
