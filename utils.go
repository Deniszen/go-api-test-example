package main

import (
	"bytes"
	"encoding/json"
	"io"
)

// JSONEncodingReaderCloser encoding io.ReadCloser to model
func JSONEncodingReaderCloser(closer io.ReadCloser, itf interface{}) {
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(closer)
	jsonString := buf.String()
	_ = json.Unmarshal([]byte(jsonString), &itf)
}

// JSONEncodingByte encoding bytes to model
func JSONEncodingByte(closer []byte, itf interface{}) {
	_ = json.Unmarshal(closer, itf)
}

// JSONToReader JSON to bytes.Reader
func JSONToReader(json string) *bytes.Reader {
	data := []byte(json)
	return bytes.NewReader(data)
}
