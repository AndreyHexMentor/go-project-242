package code

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Constants для конвертации при флаге human
const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
	PB = 1024 * TB
	EB = 1024 * PB
)

func isVisible(name string, all bool) bool {
	return all || !strings.HasPrefix(name, ".")
}

func formatSize(size int64, human bool) string {
	if !human {
		return fmt.Sprintf("%dB", size)
	}

	units := []struct {
		name  string
		value int64
	}{
		{"EB", EB}, {"PB", PB}, {"TB", TB},
		{"GB", GB}, {"MB", MB}, {"KB", KB},
	}

	for _, u := range units {
		if size >= u.value {
			return fmt.Sprintf("%.1f%s", float64(size)/float64(u.value), u.name)
		}
	}
	return fmt.Sprintf("%dB", size)
}

func getFileSize(path string, info os.FileInfo) int64 {
	size := info.Size()
	if info.Mode()&os.ModeSymlink != 0 {
		if stat, err := os.Stat(path); err == nil {
			size = stat.Size()
		} else {
			log.Printf("warning: broken symlink %q: %v", path, err)
		}
	}
	return size
}

func calculateDirSize(path string, recursive, all bool) (int64, error) {
	var total int64

	entries, err := os.ReadDir(path)
	if err != nil {
		return 0, fmt.Errorf("reading directory %q: %w", path, err)
	}

	for _, entry := range entries {
		if !isVisible(entry.Name(), all) {
			continue
		}

		full := filepath.Join(path, entry.Name())
		info, err := entry.Info()
		if err != nil {
			log.Printf("warning: cannot get info for %q: %v", full, err)
			continue
		}

		// Если это директория
		if info.IsDir() {
			if recursive {
				size, err := calculateDirSize(full, recursive, all)
				if err != nil {
					return 0, err
				}
				total += size
			}
			continue
		}

		// Файл или symlink
		total += getFileSize(full, info)
	}

	return total, nil
}

// GetPathSize возвращает размер файлов или суммарный размер файлов в директории
func GetPathSize(path string, recursive, human, all bool) (string, error) {
	info, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("cannot access path %q: %w", path, err)
	}

	var total int64

	if !info.IsDir() {
		// файл или symlink на файл
		total = getFileSize(path, info)
	} else {
		// директория
		total, err = calculateDirSize(path, recursive, all)
		if err != nil {
			return "", fmt.Errorf("error calculating directory size for %q: %w", path, err)
		}
	}

	return fmt.Sprintf("%s", formatSize(total, human)), nil
}
