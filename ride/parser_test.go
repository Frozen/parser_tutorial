package pg

import "testing"

func TestParse1(t *testing.T) {
	rs, err := Parse([]byte(`let x=  "fsas" `))
	t.Logf("%+v", rs)
	t.Log(err)
}

func TestParse2(t *testing.T) {
	rs, err := Parse([]byte(`let answersCount = 201`))
	t.Logf("%+v", rs)
	t.Log(err)
}
