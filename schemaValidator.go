package main

import (
	"fmt"
	"github.com/xeipuuv/gojsonschema"
	"path/filepath"
)

func SchemaUpload(schemaPath string) gojsonschema.JSONLoader {
	schema, _ := filepath.Abs(schemaPath)
	schemaLoader := gojsonschema.NewReferenceLoader("file:///" + schema)
	return schemaLoader
}

func DocumentUpload(doc []byte) gojsonschema.JSONLoader {
	documentLoader := gojsonschema.NewStringLoader(string(doc))
	return documentLoader
}

func CheckSchema(expectedSchema gojsonschema.JSONLoader, actualSchema gojsonschema.JSONLoader) bool {
	result, err := gojsonschema.Validate(expectedSchema, actualSchema)
	if err != nil {
		panic(err.Error())
	}

	if result.Valid() {
		return true
	} else {
		fmt.Printf("The document is not valid. see errors :\n")
		for _, desc := range result.Errors() {
			fmt.Printf("- %s\n", desc)
		}
		return false
	}
}
