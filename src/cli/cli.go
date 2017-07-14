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

func readFile(fileName string) ([]byte, error) {
	contents, err := ioutil.ReadFile(fileName)
	return contents, err
}

func Cli() {
	var filename string
	var tasknum int

	var body []byte
	arg := flag.String("file", "",
		"Usage: -file=name.ext")

	arg2 := flag.Int("task", 0, "Usage: -task=1")

	flag.Parse()
	filename, tasknum = *arg, *arg2

	for {
		if filename == "" {
			fmt.Println("fail :(")
			var input string
			fmt.Print("Please enter filename: ")
			fmt.Scanln(&input)
			filename = input
		} else {
			fmt.Println("success :)")
			break
		}
	}

	if tasknum == 0 {
		body, _ = readFile(filename)
		fmt.Println(string(body))
		resp, _ := http.Post("http://localhost:1111/tasks", "application/json", bytes.NewBuffer(body))
		r, _ := ioutil.ReadAll(resp.Body)
		var a []Answer
		json.Unmarshal(r, &a)
		fmt.Println(a[1])
	} else {
		file, _ := readFile(filename)
		m := make(map[int]interface{})
		_ = json.Unmarshal(file, &m)
		body, _ = json.Marshal(m[tasknum])
		fmt.Println(string(body))
		target := fmt.Sprintf("http://localhost:1111/task/%d", tasknum)
		resp, _ := http.Post(target, "application/json", bytes.NewBuffer(body))
		r, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(r))
	}

	fmt.Println(filename, tasknum)
}
