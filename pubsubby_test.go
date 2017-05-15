package pubsubby

import "testing"

func TestPubSubby(t *testing.T) {
	var ts testStruct
	p := newPubsubby()
	p.Subscribe("greeting", ts.Sub)
	p.Subscribe("greeting", ts.Sub)
	p.Subscribe("greeting", ts.Sub)
	p.Publish("greeting", "Hello world!")

	if ts.cnt != 3 {
		t.Fatalf("invalid calls to subscription func, expected %d and received %d", 3, ts.cnt)
	}
}

type testStruct struct {
	cnt int
}

func (ts *testStruct) Sub(val Value) (end bool) {
	ts.cnt++
	return false
}
