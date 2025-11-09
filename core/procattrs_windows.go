package core

import (
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

func applyProcAttrs(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: windows.CREATE_NO_WINDOW,
		HideWindow:    true,
	}
}
