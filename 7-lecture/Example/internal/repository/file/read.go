package file

import (
	"errors"
	"io"
	"os"
	"path"
	"path/filepath"
)

// fileReadRepository is repository for get files info and data from system.
type fileReadRepository struct {
	fileRepository
}

// NewReadRepository return fileReadRepository by cfg.
// Prepare base dir (once), invoke err if something went wrong.
func NewReadRepository(cfg *FileRepositoryConfig) (fileReadRepository, error) {
	if err := prepare(cfg.BasePath); err != nil {
		return fileReadRepository{}, err
	}
	return fileReadRepository{
		fileRepository: fileRepository{
			basePath: cfg.BasePath,
		},
	}, nil
}

// Exists invokes error if file not exists in base dir.
func (r fileReadRepository) Exists(name string) error {
	_, err := os.Stat(path.Join(r.basePath, name))
	if err != nil {
		return err
	}
	return nil
}

// Get returns file content (data).
// Returns error if operation not permitted or file not exists.
func (r fileReadRepository) Get(name string) ([]byte, error) {
	file, err := os.Open(path.Join(r.basePath, filepath.Base(name)))
	switch {
	case err != nil && errors.Is(os.ErrNotExist, err):
		return nil, err
	case err != nil:
		return nil, err
	default:
		bytes, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		return bytes, nil
	}
}
