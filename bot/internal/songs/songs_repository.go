package songs

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type SongsRepository struct {
	Destination string
}

func (r *SongsRepository) AddSong(url string) error {
	if r.Destination == "" {
		return errors.New("destination not set")
	}

	if err := r.ensurePath(r.Destination); err != nil {
		return fmt.Errorf("error creating dir for %s: %w", r.Destination, err)
	}

	file, err := os.OpenFile(r.Destination, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(url + "\n")
	if err != nil {
		return err
	}

	return nil
}

// ensurePath ensure that parent directories exists for provided file
func (r *SongsRepository) ensurePath(file string) error {
	dirPath := filepath.Dir(file)
	return os.MkdirAll(dirPath, os.ModeDir|os.ModePerm)
}
