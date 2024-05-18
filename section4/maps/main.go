package main

import (
	"fmt"
)

func main() {

	llms := make(map[string]string)

	llms["openai"] = "gpt"
	llms["google"] = "gemini"
	llms["anthropic"] = "claude"
	llms["meta"] = "llama"

	for k, v := range llms {
		fmt.Println(k, v)
	}

	delete(llms, "meta")
	fmt.Println()

	for k, v := range llms {
		fmt.Println(k, v)
	}

	v, ok := llms["meta"]

	fmt.Println(v, ok)

}
