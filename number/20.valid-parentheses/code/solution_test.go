package code

import "testing"

func Test_isValid(t *testing.T) {

	testStr1 := "()"
	testStr2 := "()[]{}"
	testStr3 := "(]"
	testStr4 := "([)]"
	testStr5 := "{[]}"

	t.Log(isValid(testStr1))
	t.Log(isValid(testStr2))
	t.Log(isValid(testStr3))
	t.Log(isValid(testStr4))
	t.Log(isValid(testStr5))

}
