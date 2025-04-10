//go:build aix
// +build aix

package gopty

const (
	TIOCGWINSZ = 0
	TIOCSWINSZ = 0
)

func ioctl(fd, cmd, ptr uintptr) error {
	return ErrUnsupported
}
