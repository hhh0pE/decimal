package decimal

import (
	"fmt"
	"database/sql/driver"
	"errors"
	//"bitbucket.org/trade4crypt/logging"
	//"reflect"
	//"go.uber.org/zap"
)

const (
	MaxIntegralDigits   = 131072 // max digits before the decimal point
	MaxFractionalDigits = 16383  // max digits after the decimal point
)

//var Logger *zap.Logger
//func init() {
//	Logger = logging.ModuleLogger("decimal")
//}

// LengthError is returned from Decimal.Value when either its integral (digits
// before the decimal point) or fractional (digits after the decimal point)
// parts are too long for PostgresSQL.
type LengthError struct {
	Part string // "integral" or "fractional"
	N    int    // length of invalid part
	max  int
}

func (e LengthError) Error() string {
	return fmt.Sprintf("%s (%d digits) is too long (%d max)", e.Part, e.N, e.max)
}


// Value implements driver.Valuer.
func (d Decimal) Value() (driver.Value, error) {
	//Logger.Debug("Decimal.Value", zap.Reflect("d", d), zap.Stringer("d", d))
	//checkValue(d)
	if d.b.IsNaN(0) {
		return "NaN", nil
	}
	if d.b.IsInf(0) {
		return nil, errors.New("Decimal.Value: DECIMAL does not accept Infinities")
	}

	dl := d.Precision()  // length of d
	sl := int(d.Scale()) // length of fractional part

	if il := dl - sl; il > MaxIntegralDigits {
		return nil, &LengthError{Part: "integral", N: il, max: MaxIntegralDigits}
	}
	if sl > MaxFractionalDigits {
		return nil, &LengthError{Part: "fractional", N: sl, max: MaxFractionalDigits}
	}
	return d.String(), nil
}

// Scan implements sql.Scanner.
func (d *Decimal) Scan(val interface{}) error {
	if d == nil {
		d = new(Decimal)
	}
	//checkValue(d)
	switch t := val.(type) {
	case Decimal:
		*d = t
	case string:
		if _, ok := d.b.SetString(t); !ok {
			if err := d.b.Context.Err(); err != nil {
				return err
			}
			return fmt.Errorf("Decimal.Scan: invalid syntax: %q", t)
		}
		return nil
	case []byte:
		return d.UnmarshalText(t)
	default:
		return fmt.Errorf("Decimal.Scan: unknown value: %#v", val)
	}
	return nil
}


type NullDecimal struct {
	Decimal
	Valid bool
}

// Scan implements the Scanner interface.
func (n *NullDecimal) Scan(value interface{}) error {
	//log.Println("NullDecimal.Scan", value)
	if value == nil {
		n.Valid = false
		return nil
	}
	if err := n.Decimal.Scan(value); err != nil {
		return err
	} else {
		n.Valid = true
		return nil
	}
}

// Value implements the driver Valuer interface.
func (n NullDecimal) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Decimal.Value()
}

func (d *NullDecimal) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		d.Valid = false
		return nil
	}
	//log.Println("NullDecimal.UnmarshalJSON", data, string(data))
	return d.b.UnmarshalText(data)
}

func (d NullDecimal) MarshalJSON() ([]byte, error) {
	return d.b.MarshalText()
}

