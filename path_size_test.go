package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupTestDataForTest(t *testing.T) string {
	t.Helper()

	root := t.TempDir()

	// Основные файлы
	require.NoError(t, os.WriteFile(filepath.Join(root, "file.txt"), []byte("hello"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(root, "empty.txt"), []byte(""), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(root, "юникод.txt"), []byte("Привет 🌍"), 0600))

	// Вложенные директории
	dirPath := filepath.Join(root, "dir")
	require.NoError(t, os.MkdirAll(filepath.Join(dirPath, "nested"), 0750))
	require.NoError(t, os.WriteFile(filepath.Join(dirPath, "a.txt"), []byte("abc"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(dirPath, ".hidden"), []byte("hidden"), 0600))
	require.NoError(t, os.WriteFile(filepath.Join(dirPath, "nested", "nested.txt"), []byte("deep"), 0600))

	// Симлинки
	require.NoError(t, os.Symlink(filepath.Join(root, "file.txt"), filepath.Join(root, "symlink")))
	require.NoError(t, os.Symlink(filepath.Join(root, "nonexistent.txt"), filepath.Join(root, "broken_symlink")))

	return root
}

func TestGetPathSize_TableDriven(t *testing.T) {
	testdata := setupTestDataForTest(t)

	tests := []struct {
		name          string
		relativePath  string
		recursive     bool
		human         bool
		includeHidden bool
		expectPart    string
		expectErr     bool
	}{
		{
			name:          "Single file",
			relativePath:  "file.txt",
			recursive:     false,
			human:         false,
			includeHidden: false,
			expectPart:    "5B",
		},
		{
			name:          "File with unicode characters",
			relativePath:  "юникод.txt",
			recursive:     false,
			human:         false,
			includeHidden: false,
			expectPart:    "17B",
		},
		{
			name:          "Empty file",
			relativePath:  "empty.txt",
			recursive:     false,
			human:         false,
			includeHidden: false,
			expectPart:    "0B",
		},
		{
			name:          "Directory first level, skip hidden",
			relativePath:  "dir",
			recursive:     false,
			human:         false,
			includeHidden: false,
			expectPart:    "3B",
		},
		{
			name:          "Directory recursive with hidden",
			relativePath:  "dir",
			recursive:     true,
			human:         false,
			includeHidden: true,
			expectPart:    "13B",
		},
		{
			name:          "Directory recursive human-readable",
			relativePath:  "dir",
			recursive:     true,
			human:         true,
			includeHidden: true,
			expectPart:    "13B",
		},
		{
			name:          "Symlink to file",
			relativePath:  "symlink",
			recursive:     false,
			human:         false,
			includeHidden: false,
			expectPart:    "5B",
		},
		{
			name:          "Broken symlink (uses link size itself)",
			relativePath:  "broken_symlink",
			recursive:     false,
			human:         false,
			includeHidden: false,
			expectPart:    "B", // длина пути символической ссылки (в байтах), не гарантировано точное значение
		},
		{
			name:          "Non-existent path should return error",
			relativePath:  "no_such_file",
			recursive:     false,
			human:         false,
			includeHidden: false,
			expectErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fullPath := filepath.Join(testdata, tt.relativePath)

			result, err := GetPathSize(fullPath, tt.recursive, tt.human, tt.includeHidden)

			if tt.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tt.relativePath)
				return
			}

			require.NoError(t, err)
			require.Contains(t, result, tt.expectPart)
			require.Contains(t, result, fullPath)
		})
	}
}

func TestFormatSize_TableDriven(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		human    bool
		expected string
	}{
		{"Bytes", 123, false, "123B"},
		{"Zero bytes", 0, true, "0B"},
		{"1KB", 1024, true, "1.0KB"},
		{"1.5KB", 1536, true, "1.5KB"},
		{"1MB", 1048576, true, "1.0MB"},
		{"1GB", 1073741824, true, "1.0GB"},
		{"1TB", 1099511627776, true, "1.0TB"},
		{"1PB", 1125899906842624, true, "1.0PB"},
		{"1EB", 1152921504606846976, true, "1.0EB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatSize(tt.size, tt.human)
			require.Equal(t, tt.expected, result)
		})
	}
}
