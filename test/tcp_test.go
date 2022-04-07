package test

import (
	"fmt"
	"github.com/dcmeshio/fakemesh"
	"github.com/dcmeshio/fakemesh/fakec"
	"github.com/dcmeshio/fakemesh/fakes"
	"net"
	"testing"
)

func TestTcpServer(t *testing.T) {

	listen("0.0.0.0:5555", nil, func(conn net.Conn) {
		println(fmt.Sprintf("Conn %s <-> %s", conn.LocalAddr(), conn.RemoteAddr()))
		bc := fakemesh.NewBufferConn(conn)

		uc, qc, forerunner, err := fakes.Receive(bc)
		if err != nil {
			println(fmt.Sprintf("%s", err))
			println(fmt.Sprintf("% x", forerunner))
			println(fmt.Sprintf("%s", string(forerunner)))
			return
		}
		println(fmt.Sprintf("Link sccessful, uc: %d, qc: %d", uc, qc))
		err = fakes.Send(bc)
		if err != nil {
			println(fmt.Sprintf("%s", err))
			return
		}

		bytes := make([]byte, 1024)
		for true {
			n, e := bc.Read(bytes)
			if e != nil {
				println(fmt.Sprintf("%s", e))
				println(fmt.Sprintf("Exit."))
				return
			}
			s := string(bytes[:n])
			println(fmt.Sprintf("%s", s))
			if s == "Client handshake" {
				_, e = bc.Write([]byte("Server handshake"))
				if e != nil {
					println(fmt.Sprintf("%s", e))
					return
				}
			}
		}
	})

}

func TestTcpClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:5555")
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}
	println(fmt.Sprintf("Conn %s <-> %s", conn.LocalAddr(), conn.RemoteAddr()))
	bc := fakemesh.NewBufferConn(conn)

	err = fakec.Send(bc, "idimesh.helmsnets.com:5855", 10011, 20022)
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}

	err = fakec.Receive(bc)
	if err != nil {
		println(fmt.Sprintf("%s", err))
		return
	}

	println(fmt.Sprintf("Link successful."))

	_, err = bc.Write([]byte("Client handshake"))

	buf := make([]byte, 1024)
	for true {
		n, e := bc.Read(buf)
		if e != nil {
			println(fmt.Sprintf("%s", e))
			return
		}
		s := string(buf[:n])
		println(fmt.Sprintf("%s", s))
		if s == "Server handshake" {
			println(fmt.Sprintf("Exit."))
			return
		}
	}

}

func TestServerRead(t *testing.T) {
	listen("0.0.0.0:5555", nil, func(conn net.Conn) {
		println(fmt.Sprintf("Conn %s <-> %s", conn.LocalAddr(), conn.RemoteAddr()))
		// 测试收到内容
		buf := make([]byte, 1024)
		for true {
			n, e := conn.Read(buf)
			if e != nil {
				println(fmt.Sprintf("%s", e))
				return
			}
			println(fmt.Sprintf("%s", string(buf[:n])))
		}
	})
}

func listen(addr string, fallback func(), handle func(conn net.Conn)) {
	if fallback != nil {
		// 监听的异步回调
		go fallback()
	}
	listenAddr, _ := net.ResolveTCPAddr("tcp", addr)
	listener, err := net.ListenTCP("tcp", listenAddr)
	if err != nil {
		println(err)
		return
	}
	defer listener.Close()
	println(fmt.Sprintf("TestScoks start and listen at: %s", listenAddr.String()))
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			println(err)
			continue
		}
		// conn 被关闭时直接清除所有数据，不管没有发送的数据
		_ = conn.SetLinger(0)
		go handle(conn)
	}
}
