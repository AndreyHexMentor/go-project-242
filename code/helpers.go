package code

import "fmt"

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

// FormatSize форматирует размер файла в человекочитаемый формат
func FormatSize(size int64, human bool) string {
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
