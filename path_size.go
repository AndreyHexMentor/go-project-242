package pathsize

import (
	"fmt"
	"os"
)

// GetPathSize возвращает размер файлов или суммарный размер файлов в директории
func GetPathSize(
	path string,
	recursive bool,
	human bool,
	all bool,
) (string, error) {
	//заглушки для линтера
	_ = recursive
	_ = human
	_ = all

	info, err := os.Lstat(path)
	if err != nil {
		return "", fmt.Errorf("cannot acces path %q: %w", path, err)
	}

	if !info.IsDir() {
		return fmt.Sprintf("%s\t%s", humanizeSize(info.Size()), path), nil
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("cannot read directory %q: %w", path, err)
	}

	var total int64
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			info, err = entry.Info()
			if err != nil {
				return "", fmt.Errorf("cannot get info for file %q: %w", entry.Name(), err)
			}
			total += info.Size()
		}
	}

	return fmt.Sprintf("%s\t%s", humanizeSize(total), path), nil
}

func humanizeSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%dKB", size/1024)
	} else if size < 1024*1024*1024 {
		return fmt.Sprintf("%dMB", size/(1024*1024))
	}
	return fmt.Sprintf("%dGB", size/(1024*1024*1024))
}
