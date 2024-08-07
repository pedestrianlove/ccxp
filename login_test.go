package ccxp

import (
	"testing"
)

func TestLogin(t *testing.T) {
	client := New("jsl", "123456")
	client.Login()
}
