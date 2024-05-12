package file

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type fileRepository struct {
	basePath string
}

var (
	once sync.Once
)

func prepare(dir string) error {
	var onceErr error
	once.Do(func() {
		log.Printf("prepare dir for data: %s\n", dir)
		if err := os.MkdirAll(dir, 0777); err != nil {
			onceErr = fmt.Errorf("error while prepare dir: %w", err)
		}
	})
	return onceErr
}
