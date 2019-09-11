package rbmarshal_test

import (
	"bytes"
	"encoding/hex"
	rbmarshal "github.com/damonchen/rubymarshal"
	"testing"
)

// User user
type User struct {
	Name string `ruby:"name"`
	Age  int    `ruby:"age"`
}

// Profile profile
type Profile struct {
	User User   `ruby:"user"`
	Job  string `ruby:"job"`
	Time int64  `ruby:"time"`
}

func TestProfileMarshal(t *testing.T) {

	v := Profile{
		User: User{
			Name: "damon",
			Age:  18,
		},
		Job:  "programmer",
		Time: 1568104088,
	}
	buff := bytes.NewBufferString("")
	err := rbmarshal.NewEncoder(buff).Encode(&v)
	if err != nil {
		t.Error(err)
	}

	expected := "04087b083a09757365727b073a096e616d6549220a64616d6f6e063a0645543a0861676569173a086a6f6249220f70726f6772616d6d6572063b07543a0974696d656c2b07985e775d"
	s := hex.EncodeToString(buff.Bytes())
	if expected != s {
		t.Error("expected not value")
	}
}

func TestProfileArray(t *testing.T) {
	type Ruby struct {
		profile []*Profile `ruby:"profile"`
	}

	ruby := Ruby{profile: []*Profile{
		&Profile{
			User: User{
				Name: "damon",
				Age:  18,
			},
			Job:  "programmer",
			Time: 1568104088,
		},
	}}
	buff := bytes.NewBufferString("")
	err := rbmarshal.NewEncoder(buff).Encode(&ruby)
	if err != nil {
		t.Error(err)
	}
	s := hex.EncodeToString(buff.Bytes())
	expected := "04087b063a0c70726f66696c655b067b083a09757365727b073a096e616d6549220a64616d6f6e063a0645543a0861676569173a086a6f6249220f70726f6772616d6d6572063b08543a0974696d656c2b07985e775d"
	if expected != s {
		t.Error("expected not value")
	}
}

