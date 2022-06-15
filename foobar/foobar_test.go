package foobar

import "testing"

func TestFooBar(t *testing.T) {
	given := 1
	wants := "1"

	get := Say(given)

	if get != wants {
		t.Errorf("Say(%d) = %s, wants %s", given, get, wants)
	}
}
