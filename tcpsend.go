package tcpsend

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

/*
// Usage example
import "github.com/angelside/tcpsend"

ip := "127.0.0.1"
port := "9100"
data := "~DATA~"
timeout := 1

err := tcpsend.Run(ip, port, data, time.Second*timeout)
if err != nil {
	fmt.Println(err)
} else {
	fmt.Println("[OK] Data sent")
}
*/

func Run(ip, port, data string, timeout time.Duration) error {
	tcpAddr, err := net.ResolveTCPAddr("tcp", net.JoinHostPort(ip, port))
	if err != nil {
		return fmt.Errorf("error resolving TCP address: %w", err)
	}

	conn, err := net.DialTimeout("tcp4", tcpAddr.String(), timeout)
	if err != nil {
		return fmt.Errorf("failed to open connection: %w", err)
	}
	defer conn.Close()

	payload := fmt.Sprintf("%s\r\n\r\n", data)
	buf := bytes.NewBufferString(payload) // Use a buffer to reduce the number of system calls made by the conn.Write() method
	_, err = buf.WriteTo(conn)
	if err != nil {
		return err
	}

	return nil
}
