package main

import (
	"fmt"
	"os"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ygorazambuja/azure-task-gen/cmd"
	commitHelper "github.com/ygorazambuja/commit-helper/pkg"
)

func main() {
	if !HasOpenaiKey() {
		fmt.Println("Erro: A chave da API do OpenAI não está definida")
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

	fmt.Printf("Files: %v\n", files)
}

func HasOpenaiKey() bool {
	return os.Getenv("OPENAI_API_KEY") != ""
}
