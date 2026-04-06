package module

import (
	"os"
	"os/exec"
)

func ClearScreenMethod() {

	clear := exec.Command("clear", "cls", "/c")
	clear.Stdout = os.Stdout
	clear.Run()

}
