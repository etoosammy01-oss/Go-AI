package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	request := Request{
		Prompt: "What is recursion?",
	}

	data, err := json.Marshal(request)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(data))
}