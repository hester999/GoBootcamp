package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: ex2.exe <command> [args...]")
		os.Exit(1)
	}

	// Базовая команда, переданная как аргумент
	cmdName := os.Args[1]
	initialArgs := os.Args[2:] // например, "-l"

	// Чтение путей из stdin
	scanner := bufio.NewScanner(os.Stdin)
	var builder strings.Builder

	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString(" ")
	}

	// Преобразуем в список аргументов
	stdinArgs := strings.Fields(builder.String())

	// Объединяем: флаги + пути
	finalArgs := append(initialArgs, stdinArgs...)

	// Запускаем команду
	cmd := exec.Command(cmdName, finalArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка запуска:", err)
		os.Exit(1)
	}
}
