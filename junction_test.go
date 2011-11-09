package junction

import(
// "time"
 "testing"
)

func rett()(bool) { return true }
func retf()(bool) { return false }

func TestJunction(t *testing.T) {
	if all(retf, rett) {
		t.Errorf("all(retf, rett) evaluated true, should be false.")
	}

}
