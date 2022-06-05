// Terminal handling with line and password input
package termios

import "os"

func String(prompt string) string {
	os.Stderr.Write([]byte(prompt))
	var res []byte
	for {
		var c [1]byte
		_,e := os.Stdin.Read(c[:])
		if e!=nil { return "" }
		switch c[0] {
		case '\r':
		case '\n':
			return string(res)
		default: 
			res = append(res, c[0])
		}
	}
	return ""
}

func Password(prompt string) string {
	old,_ := CurrentTerminal()
	t := old
	t.Echo(false).Set()
	defer old.Set()
	s := String(prompt)
	os.Stderr.Write([]byte{'\n'})
	return s
}

func PasswordWithEcho(prompt string, echo string) string {
	old,_ := CurrentTerminal()
	t := old
	t.Echo(false).Canon(false).Set()
	defer old.Set()

        os.Stderr.Write([]byte(prompt))
        var res []byte
        for {
                var c [1]byte
                _,e := os.Stdin.Read(c[:])
                if e!=nil { return "" }
                switch c[0] {
                case '\r':
                case '\n':
	                os.Stderr.Write([]byte{'\n'})
                        return string(res)
                default:
			os.Stderr.WriteString(echo)
                        res = append(res, c[0])
                }
        }

	os.Stderr.Write([]byte{'\n'})
	return ""
}

func PasswordConfirm(prompt string, prompt2 string) string {
	var s1, s2 string
	for {
		s1 = Password(prompt)
		s2 = Password(prompt2)
		if s1 == s2 { break }
	}
	return s1
}
