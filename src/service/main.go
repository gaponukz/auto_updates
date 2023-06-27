package service

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"updates/src/fileWorker"
	"updates/src/runner"
	"updates/src/versions"
)

type Service struct {
	LockalVersionExporter versions.IVersionExporter
	ArchiveInstallerUrl   string
	RunProgrammArgs       []string
}

func (service *Service) Run() error {
	defer runner.RunInBackground(service.RunProgrammArgs[0], service.RunProgrammArgs[1:]...)

	localVersion, err := service.LockalVersionExporter.Load()
	if err != nil {
		fmt.Println("Couldn't parse version from file", err.Error())
		return err
	}

	downloadUrl := service.ArchiveInstallerUrl + "?current_user_version=" + localVersion.String()
	filename, err := downloadZipFileFromUrl(downloadUrl)
	strVersion := extractVersion(filename)

	err = fileWorker.Unzip(filename, "main")

	if err != nil {
		fmt.Println("Couldn't unpack archive", err.Error())
		return err
	}

	err = fileWorker.MoveFiles("main", "")
	if err != nil {
		fmt.Println("Couldn't unpack archive", err.Error())
		return err
	}

	newVersion, err := versions.NewVersion(strVersion)
	if err != nil {
		return err
	}

	service.LockalVersionExporter.Set(newVersion)
	fileWorker.RemoveNotEmptyFolder("main")
	fileWorker.RemoveFile(filename)

	return nil
}

func downloadZipFileFromUrl(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("err: %s", err)
		return "", err
	}
	defer resp.Body.Close()
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return "", err
	}

	contentDisposition := resp.Header.Get("Content-Disposition")
	filename := extractFilename(contentDisposition)

	out, err := os.Create(filename)
	if err != nil {
		fmt.Printf("err: %s", err)
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Printf("err: %s", err)
		return "", err
	}

	return filename, nil
}

func extractFilename(contentDisposition string) string {
	const prefix = "filename="
	idx := strings.Index(contentDisposition, prefix)
	if idx == -1 {
		return ""
	}
	filename := contentDisposition[idx+len(prefix):]
	filename = strings.Trim(filename, "\"'")
	return filename
}

func extractVersion(filename string) string {
	re := regexp.MustCompile(`\d+\.\d+\.\d+`)
	match := re.FindString(filename)
	return match
}
