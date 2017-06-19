package pubsubby

import "testing"

func TestPubSubby(t *testing.T) {
	var ts testStruct
	p := newPubsubby()
	p.Subscribe("greeting", ts.Sub)
	p.Subscribe("greeting", ts.Sub)
	p.Subscribe("greeting", ts.Sub)
	p.Subscribe("name", ts.Sub)
	p.Publish("greeting", "Hello")
	p.Publish("name", "world")

	if ts.greetingCnt != 3 {
		t.Fatalf("invalid calls to subscription func with a key of name, expected %d and received %d", 3, ts.greetingCnt)
	}

	if ts.nameCnt != 1 {
		t.Fatalf("invalid calls to subscription func with a key of name, expected %d and received %d", 1, ts.nameCnt)
	}
}

type testStruct struct {
	greetingCnt int
	nameCnt     int
}

func (ts *testStruct) Sub(key Key, val Value) (end bool) {
	strKey := key.(string)
	//	if key ==
	switch strKey {
	case "greeting":
		ts.greetingCnt++
	case "name":
		ts.nameCnt++
	}

	return false
}
