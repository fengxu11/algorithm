package code

import (
	"fmt"
	"testing"
)

func TestInteger(t *testing.T) {

	fmt.Println("输出: ", Integer(123))
	fmt.Println("输出: ", Integer(-123))
	fmt.Println("输出: ", Integer(120))
	fmt.Println("输出: ", Integer(100011))
	fmt.Println("输出: ", Integer(2147482311))
}
