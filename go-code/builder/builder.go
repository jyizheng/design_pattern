package main

// Car struct defines the product
type Car struct {
	Make  string
	Model string
	Year  int
	Color string
}

// CarBuilder is the interface that builders will implement
type CarBuilder interface {
	SetMake(make string) CarBuilder
	SetModel(model string) CarBuilder
	SetYear(year int) CarBuilder
	SetColor(color string) CarBuilder
	Build() Car
}


// carBuilder is a concrete builder for building a Car
type carBuilder struct {
	car Car
}

// NewCarBuilder returns a new instance of a carBuilder
func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

func (b *carBuilder) SetMake(make string) CarBuilder {
	b.car.Make = make
	return b
}

func (b *carBuilder) SetModel(model string) CarBuilder {
	b.car.Model = model
	return b
}

func (b *carBuilder) SetYear(year int) CarBuilder {
	b.car.Year = year
	return b
}

func (b *carBuilder) SetColor(color string) CarBuilder {
	b.car.Color = color
	return b
}

func (b *carBuilder) Build() Car {
	return b.car
}

func main() {
	builder := NewCarBuilder()

	car := builder.
		SetMake("Toyota").
		SetModel("Corolla").
		SetYear(2020).
		SetColor("Red").
		Build()

	println("Car built:", car.Make, car.Model, car.Year, car.Color)
	// Output: Car built: Toyota Corolla 2020 Red
}

