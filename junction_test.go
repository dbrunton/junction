package junction

import(
 "testing"
)

func rett()(bool) { return true }
func retf()(bool) { return false }

func TestJunction(t *testing.T) {
	if all(rett, rett, retf) {
		t.Errorf("all(rett, rett, retf) evaluated true, should be false.")
	}
}
