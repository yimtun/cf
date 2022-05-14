package host

import "golang.org/x/crypto/ssh"

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

// get ip user password from  api-server 
func Connect(cfuser,ip,port,username,password string) {

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", ip+":"+port, config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	fd := int(os.Stdin.Fd())
	state, err := terminal.MakeRaw(fd)
	if err != nil {
		fmt.Println(err)
	}
	defer terminal.Restore(fd, state)

	w, h, err := terminal.GetSize(fd)
	if err != nil {
		fmt.Println(err)
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm-256color", h, w, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal("Unable to setup stdin for session", err)
	}
	//	go io.Copy(stdin, os.Stdin)

	go func() {
		for {
			buf := make([]byte, 2048)

			n, err := os.Stdin.Read(buf)
			if err == nil {
				//fmt.Println("print:", string(buf[:n])) //if print must set ssh.ECHO 0
				//fmt.Printf("%s", string(buf[:n]))
				//fmt.Printf("%c", string(buf[:n]))
				//	fmt.Sprintf("%s", string(buf[:n]))
			}

			s := string(buf[:n])
			//if s == "w" {
			//	fmt.Printf("%s", "you  cant't use shutdown")
			//} else {
			fmt.Fprint(stdin, s)
			//}
			//		d := s + "\n"
			//	tracefile(s)

		}
	}()

	stdout, err := session.StdoutPipe()
	if err != nil {
		log.Fatal("Unable to setup stdout for session", err)
	}
	//go io.Copy(os.Stdout, stdout)
	go func() {
		for {
			buf := make([]byte, 2048)

			n, err := stdout.Read(buf)
			if err == nil {
				//          fmt.Println("print:", string(buf[:n])) //if print must set ssh.ECHO 0
			}

			s := string(buf[:n])
			fmt.Fprint(os.Stdout, s)
			//      d := s + "\n"
			tracefile(s,cfuser+"@"+ip+"txt")

		}
	}()

	stderr, err := session.StderrPipe()
	if err != nil {
		log.Fatal("Unable to setup stderr for session", err)
	}
	go io.Copy(os.Stderr, stderr)
	/*
		if err := session.Shell(); err != nil {
						log.Fatalf("failed to start shell: %s", err)
								}
	*/
	//session.Run("mysql -u xx -p xx -h ")
	session.Run("bash")

	//session.Wait()
}

func tracefile(str_content string,recordFile string ) {
	fd, _ := os.OpenFile(recordFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	//fd_time := time.Now().Format("2006-01-02 15:04:05")
	//fd_content := strings.Join([]string{"======", fd_time, "=====", str_content, "\n"}, "")
	//fd_content := strings.Join([]string{str_content}, "")
	//buf := []byte(fd_content)
	fd.WriteString(str_content)
	//fd.Write(buf)
	fd.Close()
}

func checkCurrentInput(a string) (string, error) {
	for {
		var b string = ""
		c := a + b
		if c == "ww" {
			return b, nil
		}
	}

}
