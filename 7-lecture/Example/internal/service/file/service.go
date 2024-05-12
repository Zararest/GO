package file

import (
	"errors"
	"fmt"
	"os"

	sfssErrors "gitlab.com/lgp/http-server-example/internal/errors"
)

//go:generate go run github.com/golang/mock/mockgen -source=service.go -destination=internal/mock_file_repository/file_repository.go --package=mock_file_repository

type fileReadRepository interface {
	Exists(name string) error
	Get(name string) ([]byte, error)
}

type fileWriteRepository interface {
	Upload(string, []byte) error
	Delete(string) error
}

// FileService contains operations with files in base dir.
// Use readRepository and writeRepository for operations.
type FileService struct {
	readRepository  fileReadRepository
	writeRepository fileWriteRepository
}

// NewFileService returns FileServer by readRepository and writeRepository.
func NewFileService(readRepository fileReadRepository, writeRepository fileWriteRepository) FileService {
	return FileService{
		readRepository:  readRepository,
		writeRepository: writeRepository,
	}
}

// Get checks file exists and return file data.
func (s FileService) Get(name string) ([]byte, error) {
	if err := s.readRepository.Exists(name); err != nil {
		return nil, fmt.Errorf("%w: %s: %w", sfssErrors.ErrFileNotExists, name, err)
	}

	data, err := s.readRepository.Get(name)
	if err != nil {
		return nil, fmt.Errorf("%w: %s: %w", sfssErrors.ErrInternalError, name, err)
	}

	return data, nil
}

// Upload checks file not exists and return create new file with data.
func (s FileService) Upload(name string, data []byte) error {
	err := s.readRepository.Exists(name)
	switch {
	case err == nil:
		return fmt.Errorf("%w: %s", sfssErrors.ErrFileAlreadyExists, name)
	case err != nil && !errors.Is(err, os.ErrNotExist):
		return fmt.Errorf("%w: %s: %w", sfssErrors.ErrInternalError, name, err)
	}

	if err := s.writeRepository.Upload(name, data); err != nil {
		return fmt.Errorf("%w: %s: %w", sfssErrors.ErrInternalError, name, err)
	}

	return nil
}

// Delete checks file exists and remove it from system.
func (s FileService) Delete(name string) error {
	if err := s.readRepository.Exists(name); err != nil {
		return fmt.Errorf("%w: %s: %w", sfssErrors.ErrFileNotExists, name, err)
	}

	if err := s.writeRepository.Delete(name); err != nil {
		return fmt.Errorf("%w: %s: %w", sfssErrors.ErrInternalError, name, err)
	}

	return nil
}
