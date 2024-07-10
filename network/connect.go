package network

import (
	"crypto/cipher"
	"fmt"
	"net"

	"crypto/rand"

	api "github.com/minelc/go-server-api"
	"github.com/minelc/go-server-api/network"
	"github.com/minelc/go-server/network/crypto"
)

type connection struct {
	new bool
	tcp *net.TCPConn

	state network.PacketState

	certify Certify
}

func newConnection(conn *net.TCPConn) *connection {
	return &connection{
		new:     true,
		tcp:     conn,
		state:   network.SHAKE,
		certify: Certify{},
	}
}

func (c *connection) Address() net.Addr {
	return c.tcp.RemoteAddr()
}

func (c *connection) GetState() network.PacketState {
	return c.state
}

func (c *connection) SetState(state network.PacketState) {
	c.state = state
}

type Certify struct {
	name string

	used bool
	data []byte

	encrypt cipher.Stream
	decrypt cipher.Stream
}

func (c *connection) Encrypt(data []byte) (output []byte) {
	if !c.certify.used {
		return data
	}

	output = make([]byte, len(data))
	c.certify.encrypt.XORKeyStream(output, data)

	return
}

func (c *connection) Decrypt(data []byte) (output []byte) {
	if !c.certify.used {
		return data
	}

	output = make([]byte, len(data))
	c.certify.decrypt.XORKeyStream(output, data)

	return
}

func (c *connection) CertifyName() string {
	return c.certify.name
}

func (c *connection) CertifyData() []byte {
	return c.certify.data
}

func (c *connection) CertifyUpdate(secret []byte) {
	encrypt, decrypt, err := crypto.NewEncryptAndDecrypt(secret)

	c.certify.encrypt = encrypt
	c.certify.decrypt = decrypt

	if err != nil {
		panic(fmt.Errorf("failed to enable encryption for user: %s\n%v", c.CertifyName(), err))
	}

	c.certify.used = true
	c.certify.data = secret
}

func (c *connection) CertifyValues(name string) {
	c.certify.name = name
	c.certify.data = randomByteArray(4)
}

func randomByteArray(len int) []byte {
	array := make([]byte, len)
	_, _ = rand.Read(array)

	return array
}

func (c *connection) Pull(data []byte) (len int, err error) {
	len, err = c.tcp.Read(data)
	return
}

func (c *connection) Push(data []byte) (len int, err error) {
	len, err = c.tcp.Write(data)
	return
}

func (c *connection) Stop() (err error) {
	var api_conn network.Connection = c
	api.GetServer().Disconnect(api_conn)
	err = c.tcp.Close()
	return
}

func (c *connection) SendPacket(packet network.PacketO) {
	bufO := NewBuffer()
	temp := NewBuffer()

	// write buffer
	bufO.PushVrI(packet.UUID())
	packet.Push(bufO)

	temp.PushVrI(bufO.Len())
	temp.PushUAS(bufO.UAS(), false)

	c.tcp.Write(c.Encrypt(temp.UAS()))
}
