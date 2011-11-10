package junction

import(
 "time"
 "testing"
)

func rett()(bool) { return true }
func retf()(bool) { return false }
func retfWait()(bool) { time.Sleep(5e8); return false }
func rettWait()(bool) { time.Sleep(5e8); return true }

func TestJunction(t *testing.T) {
	if all(retf, rett) {
		t.Errorf("all(retf, rett) evaluated true, should be false.")
	}

	if all(rett, rett, rett, rett) { } else {
		t.Errorf("all(rett, rett, rett, rett) returned false, should be true.")
	}
	if all(rettWait, rettWait, retf) {
		t.Errorf("Failed with waiting.")
	}

}
