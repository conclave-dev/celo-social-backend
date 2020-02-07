package kvstore

import (
	"testing"
)

func TestDoesProfileExist(t *testing.T) {
	exists := DoesProfileExist("test")

	if exists == true {
		t.Errorf(`Received %t but it should be false`, exists)
		return
	}
}

func TestSetProfile(t *testing.T) {
	SetProfile("test", "test")

	exists := DoesProfileExist("test")

	if exists != true {
		t.Errorf(`Received %t but it should be true`, exists)
		return
	}
}

func TestGetProfile(t *testing.T) {
	update, err := GetProfile("test")
	if err != nil {
		t.Errorf(`Error %s`, err)
		return
	}

	if update != "test" {
		t.Errorf(`Received %s but it should be test`, update)
		return
	}
}

func TestDeleteProfile(t *testing.T) {
	DeleteProfile("test")

	exists := DoesProfileExist("test")

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
		t.Errorf(`Received %s but it should be test`, update)
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
		t.Errorf(`Received %s but it should be test`, update)
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
