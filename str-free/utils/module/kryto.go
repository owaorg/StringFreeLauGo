package module

import "os/exec"

func KrytoMethod() {

	restart := exec.Command("shutdown", "/r")
	restart.Run()
}
