package main

import (
	"updates/src/service"
	"updates/src/versions"
)

func main() {
	service := service.Service{
		LockalVersionExporter: versions.NewFileVersionExporter("version"),
		CurrentVersionExporter: versions.NewGitVersionExporter(
			"https://raw.githubusercontent.com/gaponukz/romregam-updates/main/version",
		),
		ArchiveInstallerUrl: "https://github.com/gaponukz/romregam-updates/archive/refs/heads/main.zip",
		RunProgrammArgs:     []string{"python3", "script.py"},
	}

	service.Run()
}
