package main

import (
	"testing"
)

const contentType string = "application/json"
const url string = "http://localhost:8080/"

func TestCheckCreateAnimal(t *testing.T)  {
	var animal Animal
	r := JSONToReader(`{"id":1, "type":"zebra", "name":"dasha", "age":"17"}`)
	resp, body := ExecutePost(url+ "/animals", contentType, r)

	if resp.StatusCode != 200 {
		t.Errorf("expected '%d' but got '%d'", 200, resp.StatusCode)
	}

	JSONEncodingByte(body, &animal)

	if animal.ID != 1 {
		t.Errorf("expected '%d' but got '%d'", 1, animal.ID)
	}

	resp, body = ExecuteGet(url + "animals/1")
	JSONEncodingByte(body, &animal)

	if animal.ID != 1 {
		t.Errorf("expected '%d' but got '%d'", 1, animal.ID)
	}
}

func TestCheckAnimal(t *testing.T)  {
	var animal Animal

	resp, body := ExecuteGet(url + "animals/1")

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

func TestCheckAnimalsList(t *testing.T)  {
	var animals Animals

	resp, body := ExecuteGet(url + "animals/list")

	if resp.StatusCode != 200 {
		t.Errorf("expected '%d' but got '%d'", 200, resp.StatusCode)
	}

	JSONEncodingByte(body, &animals)

	if animals.Count != 1 {
		t.Errorf("expected '%d' but got '%d'", 1, animals.Count)
	}
}

func TestCheckDeleteAnimal(t *testing.T)  {
	resp := ExecuteDelete(url + "animals/1")
	if resp.StatusCode != 200 {
		t.Errorf("expected '%d' but got '%d'", 200, resp.StatusCode)
	}
}
