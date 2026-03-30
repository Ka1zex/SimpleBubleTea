package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	cursor  int
	choices []string
	message string
}

var (
	// Жирная зелёная граница
	borderStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FF00"))
	// Цвет подсветки для выбранного пункта
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5F87"))
)

func main() {
	m := model{
		cursor:  0,
		choices: []string{"Количество дней в году", "Количество часов в году", "Количество минут в году"},
		message: "",
	}

	p := tea.NewProgram(m)
	if err := p.Start(); err != nil {
		fmt.Println("Ошибка запуска программы:", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			switch m.cursor {
			case 0:
				m.message = fmt.Sprintf("Дней в году: %d", 365)
			case 1:
				m.message = fmt.Sprintf("Часов в году: %d", 365*24)
			case 2:
				m.message = fmt.Sprintf("Минут в году: %d", 365*24*60)
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := borderStyle.Render("==============================") + "\n"
	s += "Выберите информацию:\n\n"

	for i, choice := range m.choices {
		display := choice
		if m.cursor == i {
			display = cursorStyle.Render(choice)
		}
		s += fmt.Sprintf("> %s\n", display)
	}

	s += "\n(Используйте ↑/↓ для навигации, Enter для подтверждения, q для выхода)\n"
	s += borderStyle.Render("==============================") + "\n"

	if m.message != "" {
		s += "\n" + m.message + "\n"
	}

	return s
}
