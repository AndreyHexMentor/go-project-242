package code

import (
	"fmt"
	"os"
)

// GetPathSize возвращает размер файлов или суммарный размер файлов в директории
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("cannot access path %q: %w", path, err)
	}

	if !info.IsDir() {
		// Если это файл, просто возвращаем его размер
		if !isVisible(info.Name(), all) {
			return fmt.Sprintf("%s\t%s", formatSize(0, human), path), nil
		}
		return fmt.Sprintf("%s\t%s", formatSize(info.Size(), human), path), nil
	}

	// Если это директория, начинаем подсчет (результат будет зависеть от флагов)
	total, err := calculateDirSize(path, recursive, all)
	if err != nil {
		return "", fmt.Errorf("error calculating directory size for %q: %w", path, err)
	}

	// Возвращаем общий размер директории
	return formatSize(total, human), nil
}
