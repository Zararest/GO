package file

import (
	"io"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteRepositoryUpload(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		tempData []byte
		prepare  func(string) error
		err      error
	}{
		{
			name:     "ok",
			filename: "ok-file.txt",
			tempData: tempData,
			prepare: func(string) error {
				return nil
			},
			err: nil,
		},
		{
			name:     "already exists",
			filename: "file.txt",
			tempData: tempData,
			prepare: func(name string) error {
				_, err := os.Create(name)
				if err != nil {
					return err
				}
				return nil
			},
			err: os.ErrExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			tmp := t.TempDir()
			cfg := &FileRepositoryConfig{
				BasePath: tmp,
			}
			r, _ := NewWriteRepository(cfg)

			if err := tt.prepare(path.Join(tmp, tt.filename)); err != nil {
				t.Fatal(err)
			}

			err := r.Upload(tt.filename, tt.tempData)
			if tt.err != nil {
				assert.ErrorIs(err, tt.err)
			} else {
				assert.NoError(err)

				file, err := os.Open(path.Join(tmp, tt.filename))
				if err != nil {
					t.Fatal(err)
				}

				data, err := io.ReadAll(file)
				if err != nil {
					t.Fatal(err)
				}

				assert.Equal(tt.tempData, data)
			}
		})
	}
}

func TestWriteRepositoryDelete(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		tempData []byte
		prepare  func(string) error
		err      error
	}{
		{
			name:     "file exists",
			filename: "exists-file.txt",
			tempData: tempData,
			prepare: func(name string) error {
				_, err := os.Create(name)
				if err != nil {
					return err
				}
				return nil
			},
			err: nil,
		},
		{
			name:     "already exists",
			filename: "not-exists-file.txt",
			tempData: tempData,
			prepare: func(string) error {
				return nil
			},
			err: os.ErrNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			tmp := t.TempDir()
			cfg := &FileRepositoryConfig{
				BasePath: tmp,
			}
			r, _ := NewWriteRepository(cfg)

			if err := tt.prepare(path.Join(tmp, tt.filename)); err != nil {
				t.Fatal(err)
			}

			err := r.Delete(tt.filename)
			if tt.err != nil {
				assert.ErrorIs(err, tt.err)
			} else {
				assert.NoError(err)

				_, statErr := os.Stat(path.Join(tmp, tt.filename))
				assert.ErrorIs(statErr, os.ErrNotExist)
			}
		})
	}
}
