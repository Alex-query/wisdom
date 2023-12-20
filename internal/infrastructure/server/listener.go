package server

import (
	"encoding/binary"
	"github.com/google/uuid"
	"io"
	"net"
)

type Listener struct {
	clientID   string
	connection net.Conn
}

func NewListener(connection net.Conn) Listener {
	return Listener{
		clientID:   uuid.New().String(),
		connection: connection,
	}
}

func (listener *Listener) GetClientID() string {
	return listener.clientID
}

func (listener *Listener) GetConnection() net.Conn {
	return listener.connection
}

func (listener *Listener) Read() ([]byte, error) {
	var length uint64
	err := binary.Read(listener.connection, binary.BigEndian, &length)
	if err != nil {
		return []byte{}, err
	}

	msg := make([]byte, length)
	_, err = io.ReadFull(listener.connection, msg)
	if err != nil {
		return []byte{}, err
	}

	return msg, nil
}

func (listener *Listener) Write(msg []byte) error {
	err := binary.Write(listener.connection, binary.BigEndian, uint64(len(msg)))
	if err != nil {
		return err
	}

	_, err = listener.connection.Write(msg)
	if err != nil {
		return err
	}

	return nil
}
