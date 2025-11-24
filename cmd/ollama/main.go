package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/teilomillet/gollm"
)

func main() {
	// Load API key from environment variable
	//apiKey := os.Getenv("OPENAI_API_KEY")
	//if apiKey == "" {
	//	log.Fatalf("OPENAI_API_KEY environment variable is not set")
	//}

	// Create a new LLM instance with custom configuration
	llm, err := gollm.NewLLM(
		gollm.SetOllamaEndpoint("http://127.0.0.1:11434"),
		gollm.SetProvider("ollama"),
		gollm.SetModel("llama2"),
		//gollm.SetAPIKey("ollama-local"),
		gollm.SetMaxTokens(200),
		gollm.SetMaxRetries(1),
		gollm.SetRetryDelay(time.Second*2),
		gollm.SetLogLevel(gollm.LogLevelInfo),
	)
	if err != nil {
		log.Fatalf("Failed to create LLM: %v", err)
	}

	ctx := context.Background()

	// Create a basic prompt
	prompt := gollm.NewPrompt("Explain the concept of 'recursion' in programming.")

	// Generate a response
	response, err := llm.Generate(ctx, prompt)
	if err != nil {
		log.Fatalf("Failed to generate text: %v", err)
	}
	fmt.Printf("Response:\n%s\n", response)
}
