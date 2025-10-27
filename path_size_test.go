package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupTestDataForTest(t *testing.T) string {
	currentDir, err := os.Getwd()
	require.NoError(t, err, "Failed to get current working directory")

	testdataPath := filepath.Join(currentDir, "testdata")

	// чистим старые данные
	err = os.RemoveAll(testdataPath)
	require.NoError(t, err, "Failed to clean up previous testdata directory at %s", testdataPath)

	// создаём структуру директорий
	err = os.MkdirAll(testdataPath, 0750)
	require.NoError(t, err, "Failed to create testdata directory at %s", testdataPath)

	subdirPath := filepath.Join(testdataPath, "subdir")
	err = os.MkdirAll(subdirPath, 0750)
	require.NoError(t, err, "Failed to create subdir directory at %s", subdirPath)

	// обычные файлы
	filesToCreate := map[string]string{
		"file1.txt":        "Hello, world!",
		"file2.txt":        "This is a test file.",
		"subdir/file3.txt": "Another file in subdirectory. Really long name to test kb",
	}

	for relativePath, content := range filesToCreate {
		fullPath := filepath.Join(testdataPath, relativePath)
		err := os.WriteFile(fullPath, []byte(content), 0600)
		require.NoError(t, err, "Failed to write file: %s", fullPath)
	}

	// скрытые файлы
	hiddenFiles := map[string]string{
		".hidden1.txt":        "Secret content",
		"subdir/.hidden2.txt": "Another secret",
	}
	for relativePath, content := range hiddenFiles {
		fullPath := filepath.Join(testdataPath, relativePath)
		err := os.WriteFile(fullPath, []byte(content), 0600)
		require.NoError(t, err, "Failed to write hidden file: %s", fullPath)
	}

	// скрытые директории
	hiddenDirs := []string{
		filepath.Join(testdataPath, ".hidden_dir"),
		filepath.Join(testdataPath, "subdir/.hidden_subdir"),
	}
	for _, dir := range hiddenDirs {
		err := os.MkdirAll(dir, 0750)
		require.NoError(t, err, "Failed to create hidden directory at %s", dir)
	}

	// пустая директория
	emptyDirPath := filepath.Join(testdataPath, "empty_dir")
	err = os.Mkdir(emptyDirPath, 0750)
	require.NoError(t, err, "Failed to create empty_dir directory at %s", emptyDirPath)

	return testdataPath
}

func TestGetPathSize(t *testing.T) {
	testdata := setupTestDataForTest(t)

	tests := []struct {
		name          string
		relativePath  string
		recursive     bool
		human         bool
		includeHidden bool
		expectedSize  string
		expectErr     bool
	}{
		{"Single file", "file1.txt", false, false, false, "13B", false},
		{"Directory (non-recursive)", ".", false, false, false, "33B", false},
		{"Directory (recursive)", ".", true, true, true, "118B", false},
		{"Directory with hidden files", ".", false, false, true, "47B", false},
		{"Directory with hidden dirs", ".", false, false, true, "47B", false},
		{"Non-existent path", "no_such_file.txt", false, false, false, "", true},
		{"Empty directory", "empty_dir", false, false, false, "0B", false},
		{"Empty directory with hidden files", "empty_dir", false, false, true, "0B", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fullPath := filepath.Join(testdata, tt.relativePath)

			result, err := GetPathSize(fullPath, tt.recursive, tt.human, tt.includeHidden)

			if tt.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), fullPath)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expectedSize, result)
		})
	}
}

func TestGetPathSize_RecursiveHiddenFiles(t *testing.T) {
	testdata := setupTestDataForTest(t)

	tests := []struct {
		name          string
		relativePath  string
		recursive     bool
		human         bool
		includeHidden bool
		expectedSize  string
		expectErr     bool
	}{
		{"Directory with hidden files (recursive)", ".", true, false, true, "118B", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fullPath := filepath.Join(testdata, tt.relativePath)

			result, err := GetPathSize(fullPath, tt.recursive, tt.human, tt.includeHidden)

			if tt.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), fullPath)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expectedSize, result)
		})
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		name     string
		size     int64
		human    bool
		expected string
	}{
		{"Bytes", 123, false, "123B"},
		{"Kilobytes", 1024, true, "1.0KB"},
		{"Megabytes", 1048576, true, "1.0MB"},
		{"Gigabytes", 1073741824, true, "1.0GB"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatSize(tt.size, tt.human)
			require.Equal(t, tt.expected, result)
		})
	}
}
