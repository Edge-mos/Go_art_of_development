package conv

import "fmt"

type Celsius float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

//-------------------------------------------

type Fahrenheit float64

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)
