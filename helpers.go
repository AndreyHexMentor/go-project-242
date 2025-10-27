package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Constants для конвертации при флаге human
const (
	B  = 1         // Byte
	KB = 1024 * B  // Kilobyte
	MB = 1024 * KB // Megabyte
	GB = 1024 * MB // Gigabyte
	TB = 1024 * GB // Terabyte
	PB = 1024 * TB // Petabyte
	EB = 1024 * PB // Exabyte
)

func isVisible(name string, all bool) bool {
	return all || !strings.HasPrefix(name, ".")
}

func formatSize(size int64, human bool) string {
	if !human {
		// Если human false, возвращаем размер в байтах
		return fmt.Sprintf("%dB", size)
	}

	switch {
	case size >= EB:
		return fmt.Sprintf("%.1fEB", float64(size)/EB)
	case size >= PB:
		return fmt.Sprintf("%.1fPB", float64(size)/PB)
	case size >= TB:
		return fmt.Sprintf("%.1fTB", float64(size)/TB)
	case size >= GB:
		return fmt.Sprintf("%.1fGB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.1fMB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.1fKB", float64(size)/KB)
	default:
		return fmt.Sprintf("%dB", size)
	}
}

func calculateDirSize(path string, recursive, all bool) (int64, error) {
	// Начинаем с размера текущей директории
	dirSize, err := processDir(path, recursive, all)
	if err != nil {
		return 0, err
	}
	return dirSize, nil
}

func processFile(path string, all bool) (int64, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return 0, fmt.Errorf("cannot access file %q: %w", path, err)
	}
	if !info.IsDir() && isVisible(info.Name(), all) {
		return info.Size(), nil
	}
	return 0, nil
}

func processDir(path string, recursive, all bool) (int64, error) {
	var total int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("cannot read directory %q: %w", path, err)
	}

	for _, entry := range entries {
		if !isVisible(entry.Name(), all) {
			continue
		}

		fullPath := filepath.Join(path, entry.Name())

		// Обрабатываем только файлы
		if entry.Type().IsRegular() {
			fileSize, err := processFile(fullPath, all)
			if err != nil {
				return 0, err
			}
			total += fileSize
			continue
		}

		// Обрабатываем поддиректории, если установлен флаг recursive
		if entry.Type().IsDir() && recursive {
			subDirSize, err := calculateDirSize(fullPath, recursive, all)
			if err != nil {
				return 0, err
			}
			total += subDirSize
		}
	}

	return total, nil
}
