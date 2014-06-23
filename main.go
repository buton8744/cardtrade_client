package main

import (
        "bufio"
        "fmt"
        "net"
        "os"
        "strings"
	"ex_cardtrade/proto/packet"
)

func main() {
        if len(os.Args) != 5 {
                fmt.Println("Usage: ", os.Args[0], "host, port, userid, password")
                os.Exit(1)
        }
        host := os.Args[1]
	port := os.Args[2]
	userid := os.Args[3]
	password := os.Args[4]
        conn, err := net.Dial("tcp", host+":"+port)
        checkError(err)
	client := &Client{0, userid, conn)

	if !client.Login(password) {
		fmt.Println("Login failed")
	}
	fmt.Println("Login Success")
        reader := bufio.NewReader(os.Stdin)
	go netReader(conn)
        for {
                line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
                line = strings.TrimRight(line, " \t\r\n")
		strings.Split()
		switch line {
		case "/":
		case "/":
		case "/":
		default:
			c.Chating(line)
		}
		if !netWriter(conn, line) {
			break
		}
        }
}

func checkError(err error) {
        if err != nil {
                fmt.Println("Fatal error ", err.Error())
        }
}

func (c *Client) netReader(buffer []byte) bool {
	for i := 0; i < 2048; i++ {
		buffer[i] = 0x00
	}
	for {
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			conn.Close()
			Log(err)
			return false
		}
		Log("Read", bytesRead, "bytes")
	}
	return false
}

func (c *Client)netWriter(buffer []byte) bool {
	bytesWrite, err := conn.Write(buffer)
	if err != nil {
		conn.Close()
		Log(err)
		return false
	}
	Log("Write", bytesWrite, "bytes")
	return true
}

func Log(v ...interface{}) {
	fmt.Println(v...)
}

type Client struct {
	useruid int64
	userid string
	conn *net.Conn
}

func (c *Client) Login(password string) bool {
	send_packet := &packet.Packet{}
	send_packet.SendSignInReq(c.userid, password)
	buffer := make([]byte, 2048)
	send_packet.Byte(buffer)
	c.netWriter(buffer)
	c.netReader(buffer)
	recv_packet := &packet.Packet{}
	recv_packet.Read(buffer)
	login_packet, err := recv_packet.RecvSignInAck()
	if login_packet.Type != packet.PacketType_SIGNINACK {
		fmt.Println("first packet must be the type of SignInAck")
		return false
	}
	return login_packet.Result
}
