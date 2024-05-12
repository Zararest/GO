package file

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadRepositoryGet(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		tempData []byte
		prepare  func(string, []byte) error
		err      error
	}{
		{
			name:     "file exists",
			filename: "exists-file.txt",
			tempData: tempData,
			prepare: func(name string, b []byte) error {
				file, err := os.Create(name)
				if err != nil {
					return err
				}
				if _, err := file.Write(b); err != nil {
					return err
				}
				return nil
			},
			err: nil,
		},
		{
			name:     "file not exists",
			filename: "not-exists.txt",
			tempData: tempData,
			prepare: func(string, []byte) error {
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
			r, _ := NewReadRepository(cfg)

			if err := tt.prepare(path.Join(tmp, tt.filename), tt.tempData); err != nil {
				t.Fatal(err)
			}

			data, err := r.Get(tt.filename)
			if tt.err != nil {
				assert.ErrorIs(err, tt.err)
			} else {
				assert.NoError(err)
				assert.Equal(tt.tempData, data)
			}

		})
	}
}

func TestReadRepositoryExists(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		prepare  func(string) error
		err      error
	}{
		{
			name:     "file exists",
			filename: "exists-file.txt",
			prepare: func(name string) error {
				if _, err := os.Create(name); err != nil {
					return err
				}
				return nil
			},
			err: nil,
		},
		{
			name:     "file not exists",
			filename: "not-exists.txt",
			prepare: func(s string) error {
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
			r, _ := NewReadRepository(cfg)

			if err := tt.prepare(path.Join(tmp, tt.filename)); err != nil {
				t.Fatal(err)
			}

			err := r.Exists(tt.filename)
			if tt.err != nil {
				assert.ErrorIs(err, tt.err)
			} else {
				assert.NoError(err)
			}

		})
	}
}
