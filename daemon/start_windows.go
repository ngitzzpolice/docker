package daemon

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/container"
)

// TODO Windows containerd. Needs complete implementation. Very slim at the moment.
func (daemon *Daemon) platformStart(container *container.Container) error {

	spec, err := daemon.createSpec(container)
	if err != nil {
		return err
	}

	logrus.Errorf("daemon.libcontainerdStart Not implemented on Windows yet %v", spec)
	return errors.New("daemon.libcontainerdStart Not implemented on Windows yet")
}
