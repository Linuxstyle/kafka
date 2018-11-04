package getconsumers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Consumer struct {
	Error     bool              `json:"error"`
	Message   string            `json:"message"`
	Consumers []string          `json:"consumers"`
	Request   map[string]string `json:"request"`
}

func GetConsumer(url string, cluster string) (consumers []string, err error) {
	conurl := fmt.Sprintf(url + "/" + cluster + "/consumer")

	resp, err := http.Get(conurl)
	if err != nil {
		fmt.Println("Get consumer faild:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("request consumer body faild:", err)
		return
	}
	var c *Consumer
	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println("json consumers faild:", err)
		return
	}
	consumers = c.Consumers
	return

}
