package fileWorker

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func RemoveNotEmptyFolder(folderPath string) error {
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			err = os.RemoveAll(path)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to remove non-empty folder: %v", err)
	}

	err = os.RemoveAll(folderPath)
	if err != nil {
		return err
	}

	return nil
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("failed to remove file: %v", err)
	}

	return nil
}

func MoveFiles(sourceDir string, destDir string) error {
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destDir, relPath)

		if info.IsDir() {
			err = os.MkdirAll(destPath, info.Mode())
			if err != nil {
				return err
			}
		} else {
			err = moveFile(path, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to move files: %v", err)
	}

	return nil
}

func moveFile(sourcePath string, destPath string) error {
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}

	return nil
}
