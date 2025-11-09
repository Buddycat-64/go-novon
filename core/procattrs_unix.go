package core

import (
	"os/exec"
	"syscall"
)

func applyProcAttrs(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{}
}
