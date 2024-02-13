package objects

import "testing"

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestAddInt1(t *testing.T) {

	t1 := NewSmogInteger(4)
	t2 := NewSmogInteger(6)

	objects := make([]SmogObjectInterface, 1)
	objects[0] = t2

	result, err := t1.Send("+", objects)
	if result.(*SmogInteger).GetValue() != 10 || err != nil {
		t.Fatalf(`Wrong after Op %q, %v, want "", error`, result, nil)
	}

	result, err = t1.Send("-", objects)
	if result.(*SmogInteger).GetValue() != -2 || err != nil {
		t.Fatalf(`Wrong after Op %q, %v, want "", error`, result, nil)
	}

	result, err = t1.Send("*", objects)
	if result.(*SmogInteger).GetValue() != 24 || err != nil {
		t.Fatalf(`Wrong after Op %q, %v, want "", error`, result, nil)
	}

	result, err = t1.Send("/", objects)
	if result.(*SmogInteger).GetValue() != 0 || err != nil {
		t.Fatalf(`Wrong after Op %q, %v, want "", error`, result, nil)
	}

	result, err = t1.Send("%", objects)
	if result == nil {
		t.Fatalf("result is nil")
	}
	if result.(*SmogInteger).GetValue() != 4 || err != nil {
		t.Fatalf(`Wrong after Op %q, %v, want "", error`, result, nil)
	}

}
