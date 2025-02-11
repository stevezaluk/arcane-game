package net

import "net"

const bufferSize = 32000

/*
BasicRead - Use a connection object to either read messages from the server or the client
*/
func BasicRead(conn net.Conn) (string, error) {
	buffer := make([]byte, bufferSize)

	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), err

}

/*
BasicWrite - Use a connection object to either write messages to the server or the client
*/
func BasicWrite(conn net.Conn, message string) error {
	buffer := []byte(message)

	_, err := conn.Write(buffer)
	if err != nil {
		return err
	}

	return nil
}
