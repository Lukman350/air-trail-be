package utils

import (
	"fmt"
	"log"
	"net"
)

// UdpMulticast wraps the UDP multicast reader
type UdpMulticast struct {
	Group string // multicast group address, e.g. "239.0.0.1"
	Port  int    // port, e.g. 12345
	Conn  *net.UDPConn
}

// Join sets up the UDP multicast connection
func (u *UdpMulticast) Join() error {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", u.Group, u.Port))
	if err != nil {
		return fmt.Errorf("resolve UDP addr failed: %w", err)
	}

	// Listen for multicast packets
	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		return fmt.Errorf("listen multicast failed: %w", err)
	}

	// Optionally tune the buffer size
	if err := conn.SetReadBuffer(1024 * 1024); err != nil {
		log.Printf("failed to set read buffer: %v", err)
	}

	u.Conn = conn
	return nil
}

// ReadLoop continuously reads from the stream and sends data into channel
func (u *UdpMulticast) ReadLoop(out chan<- []byte) {
	buf := make([]byte, 65536) // UDP max safe size

	for {
		n, _, err := u.Conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("read error: %v", err)
			close(out)
			return
		}
		// Copy slice before sending (avoid overwriting buffer)
		data := make([]byte, n)
		copy(data, buf[:n])
		out <- data
	}
}

// Close closes the connection
func (u *UdpMulticast) Close() {
	if u.Conn != nil {
		_ = u.Conn.Close()
	}
}
