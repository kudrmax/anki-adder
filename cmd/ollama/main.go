package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ollama/ollama/api"
)

func main() {
	ctx := context.Background()

	// Создаём клиента. Хост берётся из OLLAMA_HOST или по умолчанию (localhost:11434).
	client, err := api.ClientFromEnvironment()
	if err != nil {
		log.Fatalf("create client: %v", err)
	}

	model := "llama2" // поменяй на свою модель, например "llama3.2"
	stream := true    // включаем стриминг, чтобы видеть ответ как в UI

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Ollama REPL. Вводи промпт. 'exit', 'quit' или 'выход' — чтобы завершить.")

	for {
		fmt.Print("> ")

		if !scanner.Scan() {
			// EOF (Ctrl+D) или ошибка чтения
			fmt.Println("\nEOF, выходим.")
			break
		}
		prompt := scanner.Text()

		if prompt == "exit" || prompt == "quit" || prompt == "выход" {
			break
		}
		if prompt == "" {
			continue
		}

		fmt.Println("Ответ:")

		err = client.Generate(ctx, &api.GenerateRequest{
			Model:  model,
			Prompt: prompt,
			Stream: &stream,
			Options: map[string]any{
				"num_predict": 500,
			},
		}, func(res api.GenerateResponse) error {
			// Коллбек вызывается много раз — по мере генерации токенов.
			fmt.Print(res.Response)
			return nil
		})
		if err != nil {
			log.Printf("generate error: %v\n", err)
			continue
		}

		fmt.Println() // перевод строки после ответа
	}

	if err := scanner.Err(); err != nil {
		log.Printf("stdin error: %v\n", err)
	}
}
