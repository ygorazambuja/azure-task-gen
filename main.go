package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ygorazambuja/azure-task-gen/cmd"
	"github.com/ygorazambuja/azure-task-gen/config"
	commitHelper "github.com/ygorazambuja/commit-helper/pkg"
)

func main() {
	if err := config.Init(); err != nil {
		fmt.Printf("Erro ao inicializar configuração: %v\n", err)
		os.Exit(1)
	}

	if config.AppConfig.OpenAIAPIKey == "" {
		fmt.Println("Erro: A chave da API do OpenAI não está definida")
		fmt.Println("Por favor, execute o programa novamente para configurar a chave da API")
		os.Exit(1)
	}

	m := cmd.InitialModel()
	p := tea.NewProgram(m)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Erro ao executar o programa: %v\n", err)
		os.Exit(1)
	}

	values := m.GetValues()
	sprintId, err := strconv.Atoi(values[0])
	if err != nil {
		fmt.Println("Erro: ID da Sprint inválido")
		os.Exit(1)
	}

	areaPathId, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Println("Erro: ID do Caminho da Área inválido")
		os.Exit(1)
	}

	if sprintId == 0 || areaPathId == 0 {
		fmt.Println("Erro: Os valores não podem ser zero")
		os.Exit(1)
	}

	// Get file changes
	newFilesDiff, err := commitHelper.GetNewFiles()
	if err != nil {
		fmt.Println("Erro: Não foi possível obter os arquivos novos")
		os.Exit(1)
	}

	modifiedFilesDiff, err := commitHelper.GetModifiedFiles()
	if err != nil {
		fmt.Println("Erro: Não foi possível obter os arquivos modificados")
		os.Exit(1)
	}

	deletedFilesDiff, err := commitHelper.GetDeletedFiles()
	if err != nil {
		fmt.Println("Erro: Não foi possível obter os arquivos deletados")
		os.Exit(1)
	}

	files := append(newFilesDiff, append(modifiedFilesDiff, deletedFilesDiff...)...)

	if len(files) == 0 {
		fmt.Println("Nenhum arquivo modificado encontrado para processar")
		os.Exit(0)
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(files))
	results := make(chan cmd.Output, len(files))
	for _, file := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			fmt.Printf("Processando arquivo: %s\n", f)

			diff, err := commitHelper.GetFileDiff(f)
			if err != nil {
				errChan <- fmt.Errorf("erro ao obter diff do arquivo %s: %v", f, err)
				return
			}

			response, err := cmd.GetOpenAiResponse(f, diff)
			if err != nil {
				errChan <- fmt.Errorf("erro ao obter resposta do OpenAI para %s: %v", f, err)
				return
			}

			results <- response
		}(file)
	}

	go func() {
		wg.Wait()
		close(errChan)
		close(results)
	}()

	for err := range errChan {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}

	var taskList cmd.TaskList
	taskList.Tasks = make([]cmd.Task, 0)

	for result := range results {
		for _, t := range result.Tasks {
			var task cmd.Task
			task.State = "To Do"
			task.Title = t.Title
			task.Description = t.Description
			task.AreaID = areaPathId
			task.IterationID = sprintId
			task.Activity = config.AppConfig.DefaultTask.Activity
			task.AssignedTo = config.AppConfig.DefaultTask.AssignedTo
			task.WorkItemType = config.AppConfig.DefaultTask.WorkItemType
			task.ID = ""
			task.EstimateMade = 0
			task.OriginalEstimate = 0
			task.RemainingWork = 0
			task.ItemContrato = "Item 1"
			task.IDSPF = 19
			task.UST = config.AppConfig.DefaultTask.DefaultUST
			task.Complexidade = "BAIXA"
			taskList.Tasks = append(taskList.Tasks, task)
		}
	}

	if err := taskList.GenerateCSV(); err != nil {
		fmt.Println("Erro ao gerar o arquivo CSV:", err)
		os.Exit(1)
	}

	fmt.Println("Tasks gerados com sucesso!")
	os.Exit(0)
}
