package code

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	fmt.Println("是否是回文数: ", IsPalindrome(121))
	fmt.Println("是否是回文数: ", IsPalindrome(-121))
	fmt.Println("是否是回文数: ", IsPalindrome(10))
	fmt.Println("是否是回文数: ", IsPalindrome(2112))

}
