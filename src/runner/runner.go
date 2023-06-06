package runner

import (
	"fmt"
	"os/exec"
)

func RunInBackground(name string, args ...string) {
	cmd := exec.Command(name, args...)
	err := cmd.Start()
	if err != nil {
		fmt.Println("Can't start commands: ", err)
	}
}

func RunUntilComplete(name string, args ...string) {
	cmd := exec.Command(name, args...)
	out, err := cmd.CombinedOutput()

	if err != nil {
		if err.Error() != "exit status 6" {
			fmt.Println("Error executing command:", err)
			return
		}
	}

	fmt.Println(string(out))
}
