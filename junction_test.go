package junction

import(
 "testing"
)

var junctionTests = []bool{
 false,
}

func TestJunction(t *testing.T) {
 for _, jt := range junctionTests {
  result := Junction()
  if result != false {
   t.Errorf("Prompt is wrong: '%s'\n", jt)
  }
 }
}
