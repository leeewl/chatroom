package user

import (
	"testing"
)

func TestEnUserGetAndDel(t *testing.T) {
	var u *user
	u = GetUser(1)
	u = GetUser(2)
	DelUser(2)

	if u == nil {
		t.Error("user getUser err")
	}
}
