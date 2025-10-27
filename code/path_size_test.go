package code

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func setupTestDataForTest(t *testing.T) string {
	currentDir, err := os.Getwd()
	require.NoError(t, err, "Failed to get current working directory")

	testdataPath := filepath.Join(currentDir, "testdata")

	err = os.RemoveAll(testdataPath)
	require.NoError(t, err, "Failed to clean up previous testdata directory at %s", testdataPath)

	err = os.MkdirAll(testdataPath, 0750)
	require.NoError(t, err, "Failed to create testdata directory at %s", testdataPath)

	subdirPath := filepath.Join(testdataPath, "subdir")
	err = os.MkdirAll(subdirPath, 0750)
	require.NoError(t, err, "Failed to create subdir directory at %s", subdirPath)

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

	emptyDirPath := filepath.Join(testdataPath, "empty_dir")
	err = os.Mkdir(emptyDirPath, 0750)
	require.NoError(t, err, "Failed to create empty_dir directory at %s", emptyDirPath)

	return testdataPath
}

func TestFormatSize_Bytes(t *testing.T) {
	require.Equal(t, "123B", FormatSize(123, false))
	require.Equal(t, "123B", FormatSize(123, true))
}

func TestFormatSize_Kilobytes(t *testing.T) {
	require.Equal(t, "1.0KB", FormatSize(1024, true))
	require.Equal(t, "1024B", FormatSize(1024, false))
}

func TestFormatSize_Megabytes(t *testing.T) {
	require.Equal(t, "1.0MB", FormatSize(1048576, true))
	require.Equal(t, "1048576B", FormatSize(1048576, false))
}

func TestGetPathSize_File_HumanFalse(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	filePath := filepath.Join(testdataPath, "file1.txt")
	expectedSize := int64(len("Hello, world!"))
	expectedOutput := fmt.Sprintf("%dB\t%s", expectedSize, filePath)

	result, err := GetPathSize(filePath, false, false, false)

	require.NoError(t, err)
	require.Equal(t, expectedOutput, result)
}

func TestGetPathSize_File_HumanTrue(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	filePath := filepath.Join(testdataPath, "file1.txt")
	expectedSize := int64(len("Hello, world!"))
	expectedOutput := fmt.Sprintf("%dB\t%s", expectedSize, filePath)

	result, err := GetPathSize(filePath, false, true, false) // human=true

	require.NoError(t, err)
	require.Equal(t, expectedOutput, result)
}

func TestGetPathSize_DirectorySingleLevel_HumanFalse(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	dirPath := testdataPath
	sizeFile1 := int64(len("Hello, world!"))
	sizeFile2 := int64(len("This is a test file."))
	expectedTotalSize := sizeFile1 + sizeFile2
	expectedOutput := fmt.Sprintf("%dB\t%s", expectedTotalSize, dirPath)

	result, err := GetPathSize(dirPath, false, false, false)

	require.NoError(t, err)
	require.Equal(t, expectedOutput, result)
}

func TestGetPathSize_DirectorySingleLevel_HumanTrue(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	dirPath := testdataPath
	sizeFile1 := int64(len("Hello, world!"))
	sizeFile2 := int64(len("This is a test file."))
	expectedTotalSize := sizeFile1 + sizeFile2
	expectedOutput := fmt.Sprintf("%dB\t%s", expectedTotalSize, dirPath)

	result, err := GetPathSize(dirPath, false, true, false)

	require.NoError(t, err)
	require.Equal(t, expectedOutput, result)
}

func TestGetPathSize_DirectoryWithSubdirNotRecursive_HumanFalse(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	dirPath := testdataPath
	sizeFile1 := int64(len("Hello, world!"))
	sizeFile2 := int64(len("This is a test file."))
	expectedTotalSize := sizeFile1 + sizeFile2
	expectedOutput := fmt.Sprintf("%dB\t%s", expectedTotalSize, dirPath)

	result, err := GetPathSize(dirPath, false, false, false)

	require.NoError(t, err)
	require.Equal(t, expectedOutput, result)
}

func TestGetPathSize_DirectoryWithSubdirNotRecursive_HumanTrue(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	dirPath := testdataPath
	sizeFile1 := int64(len("Hello, world!"))
	sizeFile2 := int64(len("This is a test file."))
	expectedTotalSize := sizeFile1 + sizeFile2
	expectedOutput := fmt.Sprintf("%dB\t%s", expectedTotalSize, dirPath)

	result, err := GetPathSize(dirPath, false, true, false)

	require.NoError(t, err)
	require.Equal(t, expectedOutput, result)
}

func TestGetPathSize_EmptyDirectory_HumanFalse(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	emptyDirPath := filepath.Join(testdataPath, "empty_dir")
	expectedTotalSize := int64(0)
	expectedOutput := fmt.Sprintf("%dB\t%s", expectedTotalSize, emptyDirPath)

	result, err := GetPathSize(emptyDirPath, false, false, false)

	require.NoError(t, err)
	require.Equal(t, expectedOutput, result)
}

func TestGetPathSize_EmptyDirectory_HumanTrue(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	emptyDirPath := filepath.Join(testdataPath, "empty_dir")
	expectedTotalSize := int64(0)
	expectedOutput := fmt.Sprintf("%dB\t%s", expectedTotalSize, emptyDirPath)

	result, err := GetPathSize(emptyDirPath, false, true, false)

	require.NoError(t, err)
	require.Equal(t, expectedOutput, result)
}

func TestGetPathSize_NonExistentPath_HumanFalse(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	nonExistentPath := filepath.Join(testdataPath, "non_existent_file_or_dir.txt")

	_, err := GetPathSize(nonExistentPath, false, false, false)

	require.Error(t, err)
	require.Contains(t, err.Error(), nonExistentPath)
}

func TestGetPathSize_NonExistentPath_HumanTrue(t *testing.T) {
	testdataPath := setupTestDataForTest(t)

	nonExistentPath := filepath.Join(testdataPath, "non_existent_file_or_dir.txt")

	_, err := GetPathSize(nonExistentPath, false, true, false)

	require.Error(t, err)
	require.Contains(t, err.Error(), nonExistentPath)
}
