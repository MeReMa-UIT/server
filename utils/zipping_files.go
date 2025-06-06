package utils

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
)

func ZipDirectory(dirPath string) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == dirPath {
			return nil
		}

		relPath, err := filepath.Rel(dirPath, path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			_, err = zipWriter.Create(relPath + "/")
			return err
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		writer, err := zipWriter.Create(relPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, f)
		return err
	})

	if err != nil {
		return nil, err
	}

	if err := zipWriter.Close(); err != nil {
		return nil, err
	}

	return buf, nil
}
