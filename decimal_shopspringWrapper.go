package decimal

import (
	"errors"
)

func init() {
	Context64.OperatingMode = Go
	Context32.OperatingMode = Go
	Context128.OperatingMode = Go
	ContextUnlimited.OperatingMode = Go
}

var zeroDecimal Decimal
func Zero() Decimal {
	return zeroDecimal
}

type Decimal struct {
	b Big
}

func (d *Decimal) UnmarshalJSON(data []byte) error {
	if data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1:len(data)-1]
	}

	if len(data) == 0 {
		d.b.SetFloat64(0)
		return nil
	}

	return d.b.UnmarshalText(data)
}

func (d Decimal) MarshalJSON() ([]byte, error) {
	bytes, err := d.b.MarshalText()
	bytes = append([]byte(`"`), append(bytes, '"')...)
	return bytes, err
}

func(d Decimal) String() string {
	//if d.b == nil {
	//	return "0"
	//}
	return d.b.String()
}

func(d Decimal) Equal(d2 Decimal) bool {
	//checkValue(&d)
	//checkValue(&d2)
	return d.b.Cmp(&d2.b) == 0
}

func(d Decimal) GreaterThan(d2 Decimal) bool {
	//checkValue(&d)
	//checkValue(&d2)
	return d.b.Cmp(&d2.b) == 1
}

func(d Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	//checkValue(&d)
	//checkValue(&d2)
	r := d.b.Cmp(&d2.b)
	return r == 1 || r == 0
}

func(d Decimal) LessThan(d2 Decimal) bool {
	//checkValue(&d)
	//checkValue(&d2)
	return d.b.Cmp(&d2.b) == -1
}

func(d Decimal) LessThanOrEqual(d2 Decimal) bool {
	//checkValue(&d)
	//checkValue(&d2)
	r := d.b.Cmp(&d2.b)
	return r == -1 || r == 0
}

func(d Decimal) Mul(d2 Decimal) Decimal {
	//checkValue(&d)
	//checkValue(&d2)
	return Decimal{*d.b.mul(&d.b, &d2.b)}
}

func(d Decimal) Div(d2 Decimal) Decimal {
	//checkValue(&d)
	//checkValue(&d2)
	return Decimal{*d.b.Quo(&d.b, &d2.b)}
}

func(d Decimal) Add(d2 Decimal) Decimal {
	//checkValue(&d)
	//checkValue(&d2)
	return Decimal{*d.b.add(&d.b, &d2.b)}
}

func(d Decimal) Sub(d2 Decimal) Decimal {
	//checkValue(&d)
	//checkValue(&d2)
	return Decimal{*d.b.sub(&d.b, &d2.b)}
}

func(d Decimal) Modulo(d2 Decimal) Decimal {
	return Decimal{*d.b.Rem(&d.b, &d2.b)}
}

func(d Decimal) IsZero() bool {
	//if d.b == nil {
	//	return true
	//}
	return d.Equal(Zero())
}

func(d Decimal) StringFixed(places int) string {
	//checkValue(&d)
	rounded := d.b.Round(places)
	return rounded.String()
}

func(d Decimal) Round(places int) Decimal {
	//checkValue(&d)
	d.b.Round(places)
	return d
}

func(d Decimal) Truncate(places int32) Decimal {
	//checkValue(&d)
	ctx := d.b.Context
	ctx.RoundingMode = ToZero
	ctx.Precision = int(places)
	d.b = *ctx.Round(&d.b)
	return d
}

func (d *Decimal) MarshalText() ([]byte, error) {
	//checkValue(d)
	return d.b.MarshalText()
}

func(d *Decimal) UnmarshalText(data []byte) error {
	//checkValue(d)
	return d.b.UnmarshalText(data)
}


func(d *Decimal) Precision() int {
	//checkValue(d)
	return d.b.Precision()
}

func(d *Decimal) Scale() int {
	//checkValue(d)
	return d.b.Scale()
}

func(d *Decimal) Sign() int {
	//checkValue(d)
	return d.b.Sign()
}

func(d *Decimal) P() *Decimal {
	//checkValue(d)
	return d
}

func(d Decimal) V() Decimal {
	//checkValue(&d)
	return d
}

func NewDecimal() Decimal {
	if pool.c == nil {
		return Decimal{}
	}
	return pool.Get()
}

func NewFromString(str string) (Decimal, error) {
	var isOk bool
	d := NewDecimal()
	_, isOk = d.b.SetString(str)
	if !isOk {
		return d, errors.New("Decimal.NewFromString(\""+str+"\"): Cannot parse string")
	}
	return d, nil
}

func NewFromFloat(f float64) Decimal {
	//bigFloat := big.NewFloat(f)
	d := NewDecimal()
	d.b.SetFloat64(f)
	//Context64.Round(&d.b)
	d.b.Reduce()
	//Context64.Reduce(&d.b)
	return d
}

//func checkValue(decimal *Decimal) {
//	return
//	if decimal.b == nil {
//		decimal.b = WithContext(Context64)
//	}
//}