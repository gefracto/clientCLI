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

const alltasks int = 0
const taskscount int = 7
const target string = "http://localhost:"

func getFilename(filename string) string {
	file := filename
	for {
		if file != "" {
			fmt.Println("Читаю файл " + file)
			break
		}
		fmt.Println("Введите имя файла")
		var input string
		fmt.Scanln(&input)
		file = input
	}
	return file
}

func sortTasks(a []Answer) (result string) {
	for i := 1; i <= 7; i++ {
		for _, resp := range a {
			if resp.Task == i {
				result += fmt.Sprintf("\nTask: %d\n", i)
				if resp.Reason == "<nil>" {
					result += fmt.Sprintln(resp.Resp)
				} else {
					result += fmt.Sprintln(resp.Reason)
				}

			}
		}
	}
	return result
}

func dosingletask(file string, port string, tasknum int) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	datamap := make(map[int]interface{})
	json.Unmarshal(data, &datamap)
	body, _ := json.Marshal(datamap[tasknum])
	addr := fmt.Sprintf(target+port+"/task/%d", tasknum)
	response, err := http.Post(addr, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	result, err := ioutil.ReadAll(response.Body)
	return string(result), err
}

func doalltasks(body []byte, port string) (string, error) {
	var a []Answer
	var result string

	resp, err := http.Post(target+port+"/tasks", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return result, err
	}
	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	json.Unmarshal(response, &a)
	result = sortTasks(a)

	return result, nil
}

func Cli() (string, error) {
	var file string
	filename := flag.String("file", "", "-file=name.ext")
	tasknum := flag.Int("task", 0, "-task=1")
	port := flag.String("port", "1111", "-port=1111")

	flag.Parse()

	if *filename == "" {
		file = getFilename(*filename)
	} else {
		file = *filename
	}

	if *tasknum == alltasks {
		body, err := ioutil.ReadFile(file)
		if err != nil {
			return "", err
		}
		return doalltasks(body, *port)
	}
	return dosingletask(file, *port, *tasknum)

}
