package cli

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gefracto/kostrika-go-tasks/src/server"
)

type Answer server.Answer

func Cli() {
	var filename string
	var tasknum int
	var port string
	var addr string = "http://localhost:"

	var body []byte

	arg := flag.String("file", "", "Usage: -file=name.ext")

	arg2 := flag.Int("task", 0, "Usage: -task=1")

	arg3 := flag.String("port", "1111", "Usage: -port=1111")

	flag.Parse()
	filename, tasknum, port = *arg, *arg2, *arg3
	addr += port

	for {
		if filename != "" {
			fmt.Println("success :)")
			break
		}
		fmt.Println("fail :(")
		var input string
		fmt.Print("Please enter filename: ")
		fmt.Scanln(&input)
		filename = input

	}

	if tasknum == 0 {
		body, _ = ioutil.ReadFile(filename)
		resp, _ := http.Post(addr+"/tasks", "application/json", bytes.NewBuffer(body))
		r, _ := ioutil.ReadAll(resp.Body)
		var a []Answer
		json.Unmarshal(r, &a)
		for i := 1; i <= 7; i++ {
			for _, resp := range a {
				if resp.Task == i {
					if resp.Reason == "<nil>" {
						fmt.Printf("\nTask: %d\n", i)
						fmt.Println(resp.Resp)
					} else {
						fmt.Println(resp.Reason)
					}

				}
			}
		}
	} else {
		file, _ := ioutil.ReadFile(filename)
		m := make(map[int]interface{})
		_ = json.Unmarshal(file, &m)
		body, _ = json.Marshal(m[tasknum])
		target := fmt.Sprintf(addr+"/task/%d", tasknum)
		resp, _ := http.Post(target, "application/json", bytes.NewBuffer(body))
		r, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(r))
	}

}
