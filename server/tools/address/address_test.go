package address

import "testing"

func TestGetAddressByIP(t *testing.T) {
	addr := New("175.11.202.69")
	t.Log(addr.GetAddress())
}
