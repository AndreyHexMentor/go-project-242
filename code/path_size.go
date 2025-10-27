package code

import (
	"fmt"
	"os"
)

// GetPathSize возвращает размер файлов или суммарный размер файлов в директории
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	// заглушки для будущей логики
	_ = recursive
	_ = all

	info, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("cannot access path %q: %w", path, err)
	}

	var total int64
	if !info.IsDir() {
		total = info.Size()
	} else {
		entries, err := os.ReadDir(path)
		if err != nil {
			return "", fmt.Errorf("cannot read directory %q: %w", path, err)
		}

		for _, entry := range entries {
			if entry.Type().IsRegular() {
				finfo, err := entry.Info()
				if err != nil {
					return "", fmt.Errorf("cannot get info for file %q: %w", entry.Name(), err)
				}
				total += finfo.Size()
			}
		}
	}

	return fmt.Sprintf("%s\t%s", FormatSize(total, human), path), nil
}
