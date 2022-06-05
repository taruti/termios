package termios

import (
	"os"
	"syscall"
	"unsafe"
)

// Get current terminal settings
func CurrentTerminal() (terminal Termios, err error) {
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), _TCGETS, uintptr(unsafe.Pointer(&terminal)))
	err = os.NewSyscallError("ioctl", e)
	return
}

// Set terminal settings
func (t *Termios) Set() error {
	_, _, e := syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), _TCSETSW, uintptr(unsafe.Pointer(t)))
	return os.NewSyscallError("ioctl", e)
}

// Change echoing, but do not *change* it
func (t *Termios) Echo(echo bool) *Termios {
	if echo {
		t.Lflag |= _ECHO
	} else {
		t.Lflag &= ^uint32(_ECHO)
	}
	return t
}

// Change canonical mode, but do not *change* it
func (t *Termios) Canon(echo bool) *Termios {
	if echo {
		t.Lflag |= _CANON
	} else {
		t.Lflag &= ^uint32(_CANON)
	}
	return t
}

const (
	_TCGETS  = 0x5401
	_TCSETSW = 0x5403
	_NCCS    = 32
	_ECHO    = 010
        _CANON   = 002
)

type Termios struct {
	Iflag, Oflag, Cflag, Lflag uint32
	Line                       uint8
	Cc                         [_NCCS]uint8
	Ispeed, Ospeed             uint32
}
