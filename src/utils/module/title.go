package module

import (
	"os/exec"
)

func TitleMethod() {
	title := exec.Command("title", "String-Free TOOL")
	title.Run()
}
