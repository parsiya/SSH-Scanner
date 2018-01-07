// SSHServer(s) structs, methods and related functions.
package scanner

import (
	"errors"
	"fmt"
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
	IsSSH      bool              // true if SSH is listening on address:port
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

type SSHServers []*SSHServer

// String converts []*SSHServer to JSON. If it cannot convert to JSON, it
// will convert each member to string using fmt.Sprintf("%+v").
func (servers *SSHServers) String() string {
	var report string
	// Try converting to JSON
	report, err := ToJSON(servers, true)
	// If cannot convert to JSON
	if err != nil {
		// Save all servers as string (this is not as good as JSON)
		for _, v := range *servers {
			report += fmt.Sprintf("%+v\n%s\n", v, strings.Repeat("-", 30))
		}
		return report
	}
	return report
}
