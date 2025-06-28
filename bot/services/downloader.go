package services

import (
	"os"
	"os/exec"
	"path/filepath"
)

type Downloader struct {
	ScriptsFolder string
	SongsDirectory string
}

func (d *Downloader) Download() error {
	scriptName := filepath.Join(d.ScriptsFolder, "download")
	cmd := exec.Command(scriptName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}
