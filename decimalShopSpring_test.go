package decimal

import (
	"testing"
	"fmt"
)

func TestDecimalShopspringWrapper(t *testing.T) {
	var d Decimal
	//d := NewDecimal()
	fmt.Printf("%v", d)
}

func TestDecimalFromFloat(t *testing.T) {
	d := NewFromFloat(1.0000004)
	fmt.Println(d)

	d2, _ := NewFromString("1.0000004")
	fmt.Println(d2)
}