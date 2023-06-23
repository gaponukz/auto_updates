package main

import (
	"updates/src/service"
	"updates/src/versions"
)

func main() {
	service := service.Service{
		LockalVersionExporter: versions.NewFileVersionExporter("version"),
		ArchiveInstallerUrl:   "http://51.75.76.105/versions/download",
		RunProgrammArgs:       []string{"python3", "script.py"},
	}

	service.Run()
}
