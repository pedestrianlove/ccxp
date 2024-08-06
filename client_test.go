package ccxp

import (
	"testing"
)

func TestClient(t *testing.T) {
	client := New("jsl", "123456")
	client.Login()

	for client.GetCaptcha() != "" {
		t.Error("captcha should be empty")
	}
}
