package main

import (
	"fmt"
	"log"
	"testing"
)

const contentType string = "application/json"
const url string = "http://localhost:8080/"

func TestCheckCreateAnimal(t *testing.T) {

	var animal Animal
	r := JSONToReader(`{"id":1, "type":"zebra", "name":"dasha", "age":"17"}`)
	resp, body, err := executePost(fmt.Sprintf("%s%s", url, "/animals"), contentType, r)

	if CheckSchema(SchemaUpload("../zooTest/schemas/animal.json"), DocumentUpload(body)) != true {
		t.Errorf("expected true but got false")
	}

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected '%d' but got '%d'", 200, resp.StatusCode)
	}

	JSONEncodingByte(body, &animal)

	if animal.ID != 1 {
		t.Errorf("expected '%d' but got '%d'", 1, animal.ID)
	}

	resp, body, err = executeGet(fmt.Sprintf("%s%s", url, "animals/1"))
	JSONEncodingByte(body, &animal)

	if CheckSchema(SchemaUpload("../zooTest/schemas/animal.json"), DocumentUpload(body)) != true {
		t.Errorf("expected true but got false")
	}

	if animal.ID != 1 {
		t.Errorf("expected '%d' but got '%d'", 1, animal.ID)
	}
}

func TestCheckAnimal(t *testing.T) {
	var animal Animal

	resp, body, err := executeGet(fmt.Sprintf("%s%s", url, "animals/1"))

	if CheckSchema(SchemaUpload("../zooTest/schemas/animal.json"), DocumentUpload(body)) != true {
		t.Errorf("expected true but got false")
	}

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected '%d' but got '%d'", 200, resp.StatusCode)
	}

	JSONEncodingByte(body, &animal)

	if animal.Name != "dasha" {
		t.Errorf("expected '%s' but got '%s'", "dasha", animal.Name)
	}

	if animal.Age != "17" {
		t.Errorf("expected '%s' but got '%s'", "17", animal.Age)
	}
}

func TestCheckAnimalsList(t *testing.T) {
	var animals Animals

	resp, body, err := executeGet(fmt.Sprintf("%s%s", url, "animals/list"))

	if CheckSchema(SchemaUpload("../zooTest/schemas/animals.json"), DocumentUpload(body)) != true {
		t.Errorf("expected true but got false")
	}

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("expected '%d' but got '%d'", 200, resp.StatusCode)
	}

	JSONEncodingByte(body, &animals)

	if animals.Count != 1 {
		t.Errorf("expected '%d' but got '%d'", 1, animals.Count)
	}
}

func TestCheckDeleteAnimal(t *testing.T) {
	resp, err := executeDelete(fmt.Sprintf("%s%s", url, "animals/1"))
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("expected '%d' but got '%d'", 200, resp.StatusCode)
	}
}
