package lib01

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	factor int
)

func init() {
	factor = 40
}

func AddWithFactor(x int, y int) int {
	return (x + y) * factor
}

type local_api struct {
	port int
	root string
}

func (o *local_api) init(port int) {
	o.port = port
	o.root = fmt.Sprintf("http://localhost:%d", port)
}

func NewLocalApi(port int) *local_api {
	o := new(local_api)
	o.init(port)
	return o
}

func (o *local_api) get_url(method_name string) string {
	return fmt.Sprintf("%s/%s", o.root, method_name)
}

func (o *local_api) Call(method_name string, args any) any {
	url := o.get_url(method_name)
	jsonBytes, err := json.Marshal(
		map[string]any{"args": args})
	if err != nil {
		panic(err)
	}
	fmt.Printf("string(jsonBytes): %v\n", string(jsonBytes))
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	fmt.Println(resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var rawJson map[string]any
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&rawJson)
		if err != nil {
			panic(err)
		}
		fmt.Println(rawJson)
		return rawJson["result"]
	} else {
		panic(fmt.Sprintf("Get failed with error: %v", resp.Status))
	}
}
