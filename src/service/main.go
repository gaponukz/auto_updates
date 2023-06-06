package service

import (
	"fmt"
	"updates/src/fileWorker"
	"updates/src/runner"
	"updates/src/versions"
)

type Service struct {
	LockalVersionExporter  versions.IVersionExporter
	CurrentVersionExporter versions.IVersionExporter
	ArchiveInstallerUrl    string
	RunProgrammArgs        []string
}

func (service *Service) Run() error {

	localVersion, err := service.LockalVersionExporter.Load()
	if err != nil {
		fmt.Println("Couldn't parse version from file", err.Error())
		return err
	}

	currentVersion, err := service.CurrentVersionExporter.Load()
	if err != nil {
		fmt.Println("Couldn't parse version from git", err.Error())
		return err
	}

	if currentVersion.IsGreaterThan(localVersion) {
		runner.RunUntilComplete("curl", "-0", service.ArchiveInstallerUrl, "-LO", "main.zip")
		err := fileWorker.Unzip("main.zip", "main")

		if err != nil {
			fmt.Println("Couldn't unpack archive", err.Error())
			return err
		}

		err = fileWorker.MoveFiles("main/romregam-updates-main", "")
		if err != nil {
			fmt.Println("Couldn't unpack archive", err.Error())
			return err
		}

		fileWorker.RemoveNotEmptyFolder("main")
		fileWorker.RemoveFile("main.zip")
	}

	runner.RunInBackground(service.RunProgrammArgs[0], service.RunProgrammArgs[1:]...)
	return nil
}
