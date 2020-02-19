package pg

import "testing"

func TestParse(t *testing.T) {
	rs, err := Parse([]byte(`"bla"`))
	t.Log(rs)
	t.Log(err)
}
