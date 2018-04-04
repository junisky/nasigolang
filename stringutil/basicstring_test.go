package stringutil

import (
	"testing"
)

func TestBasicString(t *testing.T) {
	t.Run("testBasic", testBasic)
}

func testBasic(t *testing.T) {
	str := ""

	if !IsEmpty(str) {
		t.Error(str, " is not empty")
	}

	str = "Abcd"
	if !IsUpperFirst(str) {
		t.Error(str, " is not upper first")
	}

	str = "abcd"
	if !IsLowerFirst(str) {
		t.Error(str, " is not lower first")
	}

	str = "abcD"
	if !IsUpperLast(str) {
		t.Error(str, " is not upper last")
	}

	str = "abcd"
	if !IsLowerLast(str) {
		t.Error(str, " is not lower last")
	}

	str = "가나다"
	if StrLen(str) != 3 {
		t.Error(StrLen(str), "failed StrLen for", str, "- actual:", StrLen(str))
	}

	str = "abc"
	if StrLen(str) != 3 {
		t.Error(StrLen(str), "failed StrLen for", str, "- actual:", StrLen(str))
	}

}
