package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	"github.com/ygorazambuja/azure-task-gen/config"
)

type OutputTask struct {
	Title       string
	Description string
}

type Output struct {
	Tasks        []OutputTask
	FullFilePath string
}

func GetOpenAiResponse(filename string, diffContent string) (Output, error) {
	client := openai.NewClient(config.AppConfig.OpenAIAPIKey)

	var result Output

	schema, err := jsonschema.GenerateSchemaForType(result)

	if err != nil {
		log.Fatalf("GenerateSchemaForType error: %v", err)
	}

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `Você receberá um arquivo, e um diff, crie quantas tarefas forem necessarias para explicar tudo o que ocorreu nas alterações dos arquivos, crie um titulo e uma descrição. 
              Leve em conta a extensão e nome do arquivo que foram enviados, não retorne o markdown, retorne apenas o texto puro. Retorne sempre o texto em portugues brasileiro, se a task for muito grande divida em mais de uma task.
              Evite palavras como 'Inutil', 'Refatoração' troque elas por Reprocessamento, ou algo que seja menos agressivo e mais agradavel ao cliente
              `,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: fmt.Sprintf("Arquivo: %s\n\nDiff: %s", filename, diffContent),
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:   "commit_message",
					Schema: schema,
					Strict: true,
				},
			},
		},
	)
	if err != nil {
		log.Fatalf("CreateChatCompletion error: %v", err)
	}

	err = schema.Unmarshal(resp.Choices[0].Message.Content, &result)
	if err != nil {
		log.Fatalf("Unmarshal schema error: %v", err)
	}

	return result, nil
}
