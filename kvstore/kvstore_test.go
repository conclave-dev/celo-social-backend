package kvstore

import (
	"testing"
)

func TestDoesUpdateExist(t *testing.T) {
	exists := DoesUpdateExist("test")

	if exists == true {
		t.Errorf(`Received %t but it should be false`, exists)
		return
	}
}

func TestSetUpdate(t *testing.T) {
	SetUpdate("test", "test")

	exists := DoesUpdateExist("test")

	if exists != true {
		t.Errorf(`Received %t but it should be true`, exists)
		return
	}
}

func TestGetUpdate(t *testing.T) {
	update, err := GetUpdate("test")
	if err != nil {
		t.Errorf(`Error %s`, err)
		return
	}

	if update != "test" {
		t.Errorf(`Received %t but it should be test`, update)
		return
	}
}

func TestDeleteUpdate(t *testing.T) {
	DeleteUpdate("test")

	exists := DoesUpdateExist("test")

	if exists == true {
		t.Errorf("Received %t but key \"test\" should be deleted", exists)
		return
	}
}

func TestDoesUserExist(t *testing.T) {
	exists := DoesUserExist("test")

	if exists == true {
		t.Errorf(`Received %t but it should be false`, exists)
		return
	}
}

func TestSetUser(t *testing.T) {
	SetUser("test", "test")

	exists := DoesUserExist("test")

	if exists != true {
		t.Errorf(`Received %t but it should be true`, exists)
		return
	}
}

func TestGetUser(t *testing.T) {
	update, err := GetUser("test")
	if err != nil {
		t.Errorf(`Error %s`, err)
		return
	}

	if update != "test" {
		t.Errorf(`Received %t but it should be test`, update)
		return
	}
}

func TestDeleteUser(t *testing.T) {
	DeleteUser("test")

	exists := DoesUserExist("test")

	if exists == true {
		t.Errorf("Received %t but key \"test\" should be deleted", exists)
		return
	}
}

func TestDoesAddressExist(t *testing.T) {
	exists := DoesAddressExist("test")

	if exists == true {
		t.Errorf(`Received %t but it should be false`, exists)
		return
	}
}

func TestSetAddress(t *testing.T) {
	SetAddress("test", "test")

	exists := DoesAddressExist("test")

	if exists != true {
		t.Errorf(`Received %t but it should be true`, exists)
		return
	}
}

func TestGetAddress(t *testing.T) {
	update, err := GetAddress("test")
	if err != nil {
		t.Errorf(`Error %s`, err)
		return
	}

	if update != "test" {
		t.Errorf(`Received %t but it should be test`, update)
		return
	}
}

func TestDeleteAddress(t *testing.T) {
	DeleteAddress("test")

	exists := DoesAddressExist("test")

	if exists == true {
		t.Errorf("Received %t but key \"test\" should be deleted", exists)
		return
	}
}
