package main

import (
	"bytes"
	"encoding/json"
	"io"
)

func JSONEncodingReaderCloser(closer io.ReadCloser, itf interface{}) {
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(closer)
	jsonString := buf.String()
	_ = json.Unmarshal([]byte(jsonString), &itf)
}

func JSONEncodingByte(closer []byte, itf interface{}) {
	_ = json.Unmarshal(closer, &itf)
}

func JSONToReader(json string) *bytes.Reader {
	data := []byte(json)
	return bytes.NewReader(data)
}
