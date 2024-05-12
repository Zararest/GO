package file

import (
	"errors"
	"os"
	"path"
	"path/filepath"
)

// fileWriteRepository is repository for modify files in system.
type fileWriteRepository struct {
	fileRepository
}

// NewWriteRepository return fileWriteRepository by cfg.
// Prepare base dir (once), invoke err if something went wrong.
func NewWriteRepository(cfg *FileRepositoryConfig) (fileWriteRepository, error) {
	if err := prepare(cfg.BasePath); err != nil {
		return fileWriteRepository{}, err
	}
	return fileWriteRepository{
		fileRepository: fileRepository{
			basePath: cfg.BasePath,
		},
	}, nil
}

// Upload creates new file in base dir and copy data in this file.
// Returns error if operation not permitted or file already exists.
func (r fileWriteRepository) Upload(name string, data []byte) error {
	_, err := os.Stat(path.Join(r.basePath, filepath.Base(name)))
	switch {
	case err == nil:
		return os.ErrExist
	case err != nil && !errors.Is(err, os.ErrNotExist):
		return err
	}

	f, err := os.Create(path.Join(r.basePath, filepath.Base(name)))
	if err != nil {
		return err
	}
	if _, err := f.Write(data); err != nil {
		return err
	}
	return nil
}

// Delete removes file from base dir.
// Returns error if operation not permitted or file not exists.
func (r fileWriteRepository) Delete(name string) error {
	_, err := os.Stat(path.Join(r.basePath, filepath.Base(name)))
	if err != nil {
		return err
	}
	return os.Remove(path.Join(r.basePath, filepath.Base(name)))
}
