// Utility functions.
package scanner

import (
	"bufio"
	"encoding/json"
	"os"
)

// RemoveDuplicates reads a []string and removes duplicates from the slice.
// There are more efficient ways of doing this but for our purpose this way is
// easier to read. We create a second map[string]bool, go through the addresses
// and check if the address is in the map, if so we continue. If not, we add it
// to the map. In the end, the new map contains only unique addresses and we
// convert it back to []string.
func RemoveDuplicates(s []string) []string {
	m := make(map[string]bool, len(s))
	for _, v := range s {
		// If item is already in the map move to the next item
		if _, present := m[v]; present {
			continue
		}
		m[v] = true
	}
	// Now everything in m with value of true (or any value for that matter)
	// is a unique address (host:port).
	var uniques []string
	for k, _ := range m {
		uniques = append(uniques, k)
	}
	return uniques
}

// ToJSON converts input to JSON. If prettyPrint is set to True it will call
// MarshallIndent with 4 spaces.
// If your struct does not work here, make sure struct fields start with a
// capital letter. Otherwise they are not visible to the json package methods.
func ToJSON(s interface{}, prettyPrint bool) (string, error) {
	var js []byte
	var err error

	// Pretty print if specified
	if prettyPrint {
		js, err = json.MarshalIndent(s, "", "    ") // 4 spaces
	} else {
		js, err = json.Marshal(s)
	}

	// Check for marshalling errors
	if err != nil {
		return "", nil
	}

	return string(js), nil
}

// ReadTargetFile opens a file and attempts to read targets from it. Returns a
// string slice of target addresses. Each target should on its own line and in
// the correct "host:port" format.
func ReadTargetFile(file string) ([]string, error) {

	var adds []string

	// Open the file and read it
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	// Close file
	defer f.Close()

	// Read line by line and add addresses
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		adds = append(adds, scanner.Text())
	}

	// Catch scanner errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return adds, nil
}

// WriteReport stores results to file. Preferably uses ToJSON. If it cannot,
// prints them as string with .
func WriteReport(file string, servers SSHServers) error {

	outfile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer outfile.Close()

	// Try to serialize servers
	report := servers.String()
	// Write serialized date to file
	_, err = outfile.WriteString(report)
	if err != nil {
		return err
	}
	return nil
}
