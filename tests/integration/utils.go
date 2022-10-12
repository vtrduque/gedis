package integration

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func newGedisCmd() *exec.Cmd {
	fmt.Println("starting gedis...")

	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir, err := os.Open(path.Join(dirname, "../../"))
	if err != nil {
		panic(err)
	}

	cmd := exec.Command(dir.Name() + "/dist/gedis")

	go func() {
		err := cmd.Run()

		if err != nil { //TO-DO: improve this error catch
			fmt.Println("error: ", err.Error())
		}
	}()

	return cmd
}
