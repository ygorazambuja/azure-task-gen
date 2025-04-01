package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	minInputLength = 1
	maxInputLength = 10
)

type Model struct {
	inputs     []string
	cursor     int
	values     []string
	inputNames []string
	err        string
}

type inputField struct {
	name        string
	description string
}

func InitialModel() Model {
	fields := []inputField{
		{name: "sprintId", description: "ID da Sprint"},
		{name: "areaPathId", description: "ID do Caminho da Área"},
	}

	inputs := make([]string, len(fields))
	values := make([]string, len(fields))
	inputNames := make([]string, len(fields))

	for i, field := range fields {
		inputNames[i] = field.name
	}

	return Model{
		inputs:     inputs,
		cursor:     0,
		values:     values,
		inputNames: inputNames,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			if m.cursor < len(m.inputs)-1 {
				if m.validateCurrentInput() {
					m.cursor++
					m.err = ""
				}
			} else if m.validateCurrentInput() {
				return m, tea.Quit
			}

		case "up", "k":
			if m.cursor > 0 {
				if m.validateCurrentInput() {
					m.cursor--
					m.err = ""
				}
			}

		case "down", "j":
			if m.cursor < len(m.inputs)-1 {
				if m.validateCurrentInput() {
					m.cursor++
					m.err = ""
				}
			}

		case "backspace":
			if len(m.values[m.cursor]) > 0 {
				m.values[m.cursor] = m.values[m.cursor][:len(m.values[m.cursor])-1]
				m.err = ""
			}

		default:
			if len(msg.String()) == 1 && msg.String() >= "0" && msg.String() <= "9" {
				if len(m.values[m.cursor]) < maxInputLength {
					m.values[m.cursor] += msg.String()
					m.err = ""
				}
			}
		}
	}
	return m, nil
}

func (m Model) validateCurrentInput() bool {
	value := m.values[m.cursor]
	if len(value) < minInputLength {
		m.err = fmt.Sprintf("O valor deve ter pelo menos %d dígitos", minInputLength)
		return false
	}
	return true
}

func (m Model) View() string {
	s := "Digite os seguintes valores (apenas números):\n\n"

	for i := 0; i < len(m.inputs); i++ {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		s += fmt.Sprintf("%s %s: %s\n", cursor, m.inputNames[i], m.values[i])
	}

	if m.err != "" {
		s += fmt.Sprintf("\nErro: %s\n", m.err)
	}

	s += "\nPressione Enter para mover para o próximo campo ou enviar\n"
	s += "Pressione q para sair\n"

	return s
}

func (m Model) GetValues() []string {
	return m.values
}
