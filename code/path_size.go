package code

import (
	"fmt"
	"os"
)

// GetPathSize возвращает размер файлов или суммарный размер файлов в директории
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	// заглушки для будущей логики
	_ = recursive

	info, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("cannot access path %q: %w", path, err)
	}

	if !info.IsDir() {
		if !isVisible(info.Name(), all) {
			return fmt.Sprintf("%s\t%s", formatSize(0, human), path), nil
		}
		return fmt.Sprintf("%s\t%s", formatSize(info.Size(), human), path), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("cannot read directory %q: %w", path, err)
	}

	var total int64
	for _, entry := range entries {
		if !isVisible(entry.Name(), all) {
			continue
		}
		if entry.Type().IsRegular() {
			finfo, err := entry.Info()
			if err != nil {
				return "", fmt.Errorf("cannot get info for file %q: %w", entry.Name(), err)
			}
			total += finfo.Size()
		}
	}
	return fmt.Sprintf("%s\t%s", formatSize(total, human), path), nil
}
