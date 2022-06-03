package container

import (
	"os"
	"syscall"

	log "github.com/sirupsen/logrus"
)

func RunContainerInitProcess(cmd string, args []string) error {
	log.Infof("cmd: %s", cmd)

	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NODEV | syscall.MS_NOSUID
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	argv := []string{cmd}
	if err := syscall.Exec(cmd, argv, os.Environ()); err != nil {
		log.Errorf(err.Error())
	}

	return nil
}
