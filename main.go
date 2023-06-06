package main

import (
	"fmt"
	"updates/src/runner"
	"updates/src/versions"
)

func main() {
	fileExporte := versions.NewFileVersionExporter("version")
	gitExporter := versions.NewGitVersionExporter(
		"https://raw.githubusercontent.com/gaponukz/romregam-updates/main/version",
	)

	localVersion, err := fileExporte.Load()
	if err != nil {
		fmt.Println("Couldn't parse version from file", err.Error())
		return
	}

	currentVersion, err := gitExporter.Load()
	if err != nil {
		fmt.Println("Couldn't parse version from git", err.Error())
		return
	}

	if currentVersion.IsGreaterThan(localVersion) {
		runner.RunUntilComplete("curl", "-0", "https://github.com/gaponukz/romregam-updates/archive/refs/heads/main.zip", "-LO", "main.zip")
	}

	runner.RunInBackground("python3", "script.py")
}
