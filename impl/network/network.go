package network

import (
	"errors"
	"fmt"
	"net"
	"strconv"

	"github.com/minelc/go-server/api/network"
)

func StartNet(port int, host string) error {
	ser, err := net.ResolveTCPAddr("tcp4", host+":"+strconv.Itoa(port))
	if err != nil {
		return errors.New("address resolution failed. Error: " + err.Error())
	}

	tcp, err := net.ListenTCP("tcp4", ser)
	if err != nil {
		return errors.New("Failed to bind " + err.Error())
	}

	go func() {
		defer tcp.Close()
		for {
			con, err := tcp.AcceptTCP()

			if err != nil {
				println(err.Error())
				break
			}

			con.SetNoDelay(true)
			go handleConnection(newConnection(con))
		}
	}()
	return nil
}

func handleConnection(conn *connection) {

	for {
		inf := make([]byte, 1024)
		sze, err := conn.Pull(inf)

		if err != nil || sze == 0 {
			conn.Stop()
			break
		}

		buf := NewBufferWith(conn.Decrypt(inf[:sze]))

		// decompression
		// decryption

		if buf.UAS()[0] == 0xFE { // LEGACY PING
			continue
		}

		packetLen := buf.PullVrI()

		bufI := NewBufferWith(buf.UAS()[buf.InI() : buf.InI()+packetLen])

		uuid := bufI.PullVrI()
		packetI := getPacketI(uuid, conn.GetState())
		if packetI == nil {
			fmt.Printf("unable to decode %v packet with uuid: %d", conn.GetState(), uuid)
			return
		}

		packetI.Pull(bufI)

		var api_conn network.Connection = conn
		packetI.Handle(&api_conn)
	}
}
