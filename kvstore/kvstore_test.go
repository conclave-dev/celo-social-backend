package kvstore

import (
	"testing"
)

func TestIsUpdateExists(t *testing.T) {
	exists := IsUpdateExists("test")
	if exists == true {
		t.Errorf("Received %t but key \"test\" should not exist", exists)
		return
	}
}

func TestSetUpdate(t *testing.T) {
	SetUpdate("test", "test")

	exists := IsUpdateExists("test")
	if exists != true {
		t.Errorf("Received %t but key \"test\" should be set", exists)
		return
	}
}

func TestDeleteUpdate(t *testing.T) {
	DeleteUpdate("test")

	exists := IsUpdateExists("test")
	if exists == true {
		t.Errorf("Received %t but key \"test\" should be deleted", exists)
		return
	}
}
