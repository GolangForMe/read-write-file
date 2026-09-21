package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputFilePath := filepath.Join("data", "input.txt")
	outputFilePath := filepath.Join("data", "output.txt")

	upper := func(s string) (string, error) {
		return strings.ToUpper(s), nil
	}

	err := ReadProcessWrite(inputFilePath, outputFilePath, upper)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Готово: данные обработаны и записаны")
}

// ReadProcessWrite читает файл по пути inputPath, применяет к его содержимому
// функцию process и дописывает результат в файл outputPath.
// Возвращает ошибку, если чтение, обработка или запись завершились неудачей.
func ReadProcessWrite(
	inputPath string,
	outputPath string,
	process func(string) (string, error),
) error {
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return fmt.Errorf("не удалось открыть входной файл %q: %w", inputPath, err)
	}
	defer inputFile.Close()

	data, err := io.ReadAll(inputFile)
	if err != nil {
		return fmt.Errorf("не удалось прочитать входной файл %q: %w", inputPath, err)
	}

	result, err := process(string(data))
	if err != nil {
		return fmt.Errorf("ошибка обработки данных: %w", err)
	}

	outputFile, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("не удалось открыть/создать выходной файл %q: %w", outputPath, err)
	}
	defer outputFile.Close()

	if _, err := outputFile.WriteString(result); err != nil {
		return fmt.Errorf("не удалось записать результат в %q: %w", outputPath, err)
	}

	return nil
}
