// Utility functions.
package scanner

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
