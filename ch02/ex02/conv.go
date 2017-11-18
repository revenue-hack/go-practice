package tempconv

func CToF(c Celsius) Fahrenheit          { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius          { return Celsius((f - 32) * 5 / 9) }
func CToA(c Celsius) AbsoluteTemperature { return AbsoluteTemperature(c - AbsoluteZeroC) }

func LbToKg(l Lb) Kg { return Kg(l * 0.454) }
func KgToLb(k Kg) Lb { return Lb(k * 2.205) }

func FtToMeter(f Ft) Meter { return Meter(f * 3.28) }
func MeterToFt(m Meter) Ft { return Ft(m / 3.28) }
