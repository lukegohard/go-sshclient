package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
)

type ServerAuth struct {
	Username  string
	Password  string
	IpAddress string
	Port      string
}

var auth ServerAuth

func main() {

	//Read server auth data
	flag.Parse()
	
	//check if username and IpAddress are empty
	if auth.Username == "" {
		log.Println("[-]Please insert username.")
		os.Exit(1)
	}

	if auth.IpAddress == "" {
		log.Println("[-]Please insert IP Address.")
		os.Exit(1)
	}

	conn_config := &ssh.ClientConfig{User: auth.Username, Auth: []ssh.AuthMethod{ssh.Password(auth.Password)}, HostKeyCallback: ssh.InsecureIgnoreHostKey()}
	conn_addr := net.JoinHostPort(auth.IpAddress, auth.Port)
	
	//creating new connection
	conn, err := ssh.Dial("tcp", conn_addr, conn_config)
	CheckErr(err)
	
	//requesting a new session
	session, err := conn.NewSession()
	CheckErr(err)
	defer session.Close()

	//redirect IO of Server at the Client
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	stdin, err := session.StdinPipe()
	CheckErr(err)

	// requesting pseudoterminal
	terminal := ssh.TerminalModes{
		ssh.ECHO: 0,
	}
	err = session.RequestPty("vt220", 40, 130, terminal)
	CheckErr(err)

	//shell
	err = session.Shell()
	CheckErr(err)

	for {
		io.Copy(stdin, os.Stdin)
	}

}

func init() {

	flag.StringVar(&auth.Username, "u", "", "username")
	flag.StringVar(&auth.Password, "pw", "", "password")
	flag.StringVar(&auth.IpAddress, "a", "", "server ip address")
	flag.StringVar(&auth.Port, "p", "22", "server port")

}

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
