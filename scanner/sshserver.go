// SSHServer(s) structs, methods and related functions.
package scanner

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"golang.org/x/crypto/ssh"
)

// Struct to hold server data
type SSHServer struct {
	Address    string            // host:port
	Host       string            // IP address
	Port       int               // port
	IsSSH      bool              // true if SSH is listening on host:port
	Banner     string            // banner text, if any
	Certs      []ssh.Certificate // server's certificates
	Hostname   string            // hostname
	PublicKeys []ssh.PublicKey   // server's public keys
}

// NewSSHServer returns a new SSHServer with address, host and port populated.
// Returns an error if address cannot be processed.
func NewSSHServer(address string) (*SSHServer, error) {
	// Process address, return error if it's not in the correct format
	host, port, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	var s SSHServer

	s.Address = address
	s.Host = host
	s.Port, err = strconv.Atoi(port)
	if err != nil {
		return nil, err
	}
	// If port is not in (0,65535]
	if 0 > s.Port || s.Port > 65535 {
		return nil, errors.New(port + " invalid port")
	}
	return &s, nil
}

// -----

// SSHServers is a slice of *SSHServer
type SSHServers []*SSHServer

// Initialize converts the list of addresses to *SSHServer and stores the
// results in the SSHServers receiver.
func (s *SSHServers) Initialize(addresses []string, logger *log.Logger) {
	for _, addr := range addresses {
		ts, err := NewSSHServer(addr)
		if err != nil {
			logger.Printf("could not processs %v\n", err)
			continue
		}
		*s = append(*s, ts)
	}
}

// Process goes through all servers and populates them.
func (s *SSHServers) Process(logger *log.Logger) {
	for _, server := range *s {
		logger.Printf("%+v", server)
	}
}

// String converts []*SSHServer to JSON. If it cannot convert to JSON, it
// will convert each member to string using fmt.Sprintf("%+v").
func (s *SSHServers) String() string {
	var report string
	// Try converting to JSON
	report, err := ToJSON(s, true)
	// If cannot convert to JSON
	if err != nil {
		// Save all servers as string (this is not as good as JSON)
		for _, v := range *s {
			report += fmt.Sprintf("%+v\n%s\n", v, strings.Repeat("-", 30))
		}
		return report
	}
	return report
}
