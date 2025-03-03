package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	currentDir string   // Текущая директория
	entries    []string // Список файлов и директорий
	cursor     int      // Текущая позиция курсора
	offset     int      // Смещение для отображения элементов
	visible    int      // Количество видимых элементов
	selected   *string  // Указатель на строку для записи выбранного элемента
	quitting   bool     // Флаг завершения программы
}

// Инициализация модели
func (m model) Init() tea.Cmd {
	return nil
}

// Загрузка содержимого директории
func loadDirContents(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var contents []string
	contents = append(contents, "..") // Возможность подняться выше

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() {
			name = name + "/" // Добавляем '/' для директорий
		}
		contents = append(contents, name)
	}

	// Сортируем содержимое для удобства
	sort.Strings(contents[1:]) // Пропускаем ".."
	return contents, nil
}

// Логика обновления состояния модели при вводе клавиш и изменении окна
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		// Управление курсором вверх
		case "up":
			if m.cursor > 0 {
				m.cursor--
				// Если курсор поднялся в самый верх, не смещаемся дальше
				if m.cursor < m.offset {
					m.offset--
				}
			}
			// Если мы на первом элементе (..), смещение всегда равно 0
			if m.cursor == 0 {
				m.offset = 0
			}
		// Управление курсором вниз
		case "down":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
				// Если курсор в нижней части экрана, увеличиваем смещение
				if m.cursor >= m.offset+m.visible {
					m.offset++
				}
			}
		// Выбор элемента клавишей Enter
		case "enter":
			selectedEntry := m.entries[m.cursor]

			if selectedEntry == ".." {
				// Переход на директорию выше, если не корень
				if m.currentDir != "/" {
					m.currentDir = filepath.Dir(m.currentDir)
					m.entries, _ = loadDirContents(m.currentDir)
					m.cursor = 0
					m.offset = 0
				}
			} else {
				fullPath := filepath.Join(m.currentDir, strings.TrimSuffix(selectedEntry, "/"))

				info, err := os.Stat(fullPath)
				if err != nil {
					break
				}

				if info.IsDir() {
					// Если это директория, открываем её содержимое
					m.currentDir = fullPath
					m.entries, _ = loadDirContents(m.currentDir)
					m.cursor = 0
					m.offset = 0
				} else {
					// Если это файл, записываем его полный путь и выходим
					*m.selected = fullPath
					m.quitting = true
					return m, tea.Quit
				}
			}
		// Переход в родительскую директорию клавишей Left
		case "left":
			if m.currentDir != "/" {
				m.currentDir = filepath.Dir(m.currentDir)
				m.entries, _ = loadDirContents(m.currentDir)
				m.cursor = 0
				m.offset = 0
			}
		// Переход в выбранную директорию клавишей Right (если выбрана директория, и это не "..")
		case "right":
			selectedEntry := m.entries[m.cursor]

			// Проверка, что выбран не ".."
			if selectedEntry != ".." {
				fullPath := filepath.Join(m.currentDir, strings.TrimSuffix(selectedEntry, "/"))

				info, err := os.Stat(fullPath)
				if err != nil {
					break
				}

				if info.IsDir() {
					// Если это директория, открываем её содержимое
					m.currentDir = fullPath
					m.entries, _ = loadDirContents(m.currentDir)
					m.cursor = 0
					m.offset = 0
				}
			}
		// Завершение программы по клавише 'q'
		case "q":
			m.quitting = true
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		// Обновляем количество видимых строк на основе нового размера окна
		m.visible = msg.Height - 1 // Минус одна строка для инструкций
	}
	return m, nil
}

// Представление модели (вывод в консоль)
func (m model) View() string {
	if m.quitting {
		// Ничего не выводим при завершении
		return ""
	}

	// Построение интерфейса с выводом списка файлов/директорий и курсора
	s := ""
	start := m.offset
	end := m.offset + m.visible
	if end > len(m.entries) {
		end = len(m.entries)
	}

	for i := start; i < end; i++ {
		cursor := " " // Пробел если элемент не выбран
		if m.cursor == i {
			cursor = ">" // Стрелка указывает на выбранный элемент
		}
		s += fmt.Sprintf("%s %s\n", cursor, m.entries[i])
	}
	s += "\n'enter' для выбора, 'q' для выхода, 'left' для перехода вверх, 'right' для перехода в директорию."
	return s
}

// Функция для выбора файла и записи его пути в переданную строку
func SelectFile(selected *string) error {
	// Начинаем с текущей директории
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}

	entries, err := loadDirContents(currentDir)
	if err != nil {
		return err
	}

	m := model{
		currentDir: currentDir,
		entries:    entries,
		cursor:     0,
		offset:     0,
		visible:    10, // По умолчанию, обновится при получении WindowSizeMsg
		selected:   selected,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	return p.Start()
}
