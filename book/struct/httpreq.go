package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {

	client := &http.Client{}

	request, e := http.NewRequest("POST", "http://www.163.com", strings.NewReader("key=value"))

	if e != nil {
		fmt.Println(e)
		os.Exit(1)
		return
	}

	request.Header.Add("User-Agent", "myclient")

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}

	data, err := ioutil.ReadAll(response.Body)

	fmt.Println(string(data))

	defer response.Body.Close()

}
