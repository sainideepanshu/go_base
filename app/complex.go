package app

type Complex struct {
	Real      float64
	Imaginary float64
}

func NewComplex(real float64, imag float64) *Complex {
	c := Complex{
		Real:      real,
		Imaginary: imag,
	}
	return &c
}

func (c *Complex) Add(other *Complex) *Complex {
	return &Complex{
		Real:      c.Real + other.Real,
		Imaginary: c.Imaginary + other.Imaginary,
	}
}

func (c *Complex) Subtract(other *Complex) *Complex {
	return &Complex{
		Real:      c.Real - other.Real,
		Imaginary: c.Imaginary - other.Imaginary,
	}
}

func (c *Complex) Multiply(other *Complex) *Complex {
	return &Complex{
		Real:      c.Real*other.Real - c.Imaginary*other.Imaginary,
		Imaginary: c.Real*other.Imaginary + c.Imaginary*other.Real,
	}
}

func (c *Complex) Display() {
	println("Complex Number:", c.Real, "+", c.Imaginary, "i")
}