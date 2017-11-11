package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type AbsoluteTemperature float64
type Lb float64
type Kg float64
type Ft float64
type Meter float64

const (
  AbsoluteZeroC Celsius = -273.15
  FreezingC Celsius = 0
  BoilingC Celsius = 100
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (a AbsoluteTemperature) String() string { return fmt.Sprintf("%gK", a) }
func (l Lb) String() string { return fmt.Sprintf("%glb", l) }
func (k Kg) String() string { return fmt.Sprintf("%gkg", k) }
func (f Ft) String() string { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

