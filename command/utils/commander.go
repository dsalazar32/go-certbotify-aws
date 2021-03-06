package utils

import (
	"os"
	"os/exec"
)

func Commander(cwd string, eVars ...string) func(c string, output bool) ([]byte, error) {
	cmd := func(c string, output bool) ([]byte, error) {
		command := exec.Command("sh", "-c", c)
		command.Env = append(os.Environ(), eVars...)
		command.Stderr = os.Stderr
		command.Dir = cwd
		if output {
			return command.Output()
		}
		if err := command.Run(); err != nil {
			return nil, err
		}
		return nil, nil
	}

	return cmd
}
