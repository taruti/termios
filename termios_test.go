package termios

import "testing"
import "os"

func TestT(tt *testing.T) {
    t,_ := CurrentTerminal()
    os.Stdout.Write([]byte("echoing\n"))
    t.Echo(false).Set()
    CurrentTerminal()
    t.Echo(true)
    t.Set()
    os.Stdout.Write([]byte("echoing\n"))
	String("foo: ")
	Password("bar: ")
}

