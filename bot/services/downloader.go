package services

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type Downloader struct {
	ScriptsFolder  string
	SongsDirectory string
}

func (d *Downloader) DownloadSong(url string) error {
	scriptName := filepath.Join(d.ScriptsFolder, "download")

	env := os.Environ()
	env = append(env, fmt.Sprintf("DOWNLOAD_DIR=%s", d.SongsDirectory))

	cmd := exec.Command(scriptName, url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env

	if err := cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}

func (d *Downloader) DownloadArchive(archivePath string) error {
	scriptName := filepath.Join(d.ScriptsFolder, "download_archive")

	env := os.Environ()
	env = append(env, fmt.Sprintf("DOWNLOAD_DIR=%s", d.SongsDirectory))

	archive, err := os.Open(archivePath)
	if err != nil {
		return err
	}
	defer archive.Close()

	cmd := exec.Command(scriptName)
	cmd.Stdin = archive
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = env

	if err = cmd.Start(); err != nil {
		return err
	}

	return cmd.Wait()
}
