package main

import "testing"

func TestBuilder(t *testing.T) {
	builder := NewCarBuilder()

	car := builder.
		SetMake("Toyota").
		SetModel("Corolla").
		SetYear(2020).
		SetColor("Red").
		Build()

	expectedMake := "Toyota"
	expectedModel := "Corolla"
	expectedYear := 2020
	expectedColor := "Red"

	if car.Make != expectedMake {
		t.Errorf("Expected make to be %s, but got %s", expectedMake, car.Make)
	}

	if car.Model != expectedModel {
		t.Errorf("Expected model to be %s, but got %s", expectedModel, car.Model)
	}

	if car.Year != expectedYear {
		t.Errorf("Expected year to be %d, but got %d", expectedYear, car.Year)
	}

	if car.Color != expectedColor {
		t.Errorf("Expected color to be %s, but got %s", expectedColor, car.Color)
	}
}

func TestBuilder_SetMake(t *testing.T) {
	builder := NewCarBuilder()
	make := "Toyota"

	result := builder.SetMake(make)

	if result != builder {
		t.Errorf("Expected SetMake to return the builder, but got %v", result)
	}

	if builder.car.Make != make {
		t.Errorf("Expected make to be %s, but got %s", make, builder.car.Make)
	}
}

func TestBuilder_SetModel(t *testing.T) {
	builder := NewCarBuilder()
	model := "Corolla"

	result := builder.SetModel(model)

	if result != builder {
		t.Errorf("Expected SetModel to return the builder, but got %v", result)
	}

	if builder.car.Model != model {
		t.Errorf("Expected model to be %s, but got %s", model, builder.car.Model)
	}
}

func TestBuilder_SetYear(t *testing.T) {
	builder := NewCarBuilder()
	year := 2020

	result := builder.SetYear(year)

	if result != builder {
		t.Errorf("Expected SetYear to return the builder, but got %v", result)
	}

	if builder.car.Year != year {
		t.Errorf("Expected year to be %d, but got %d", year, builder.car.Year)
	}
}

func TestBuilder_SetColor(t *testing.T) {
	builder := NewCarBuilder()
	color := "Red"

	result := builder.SetColor(color)

	if result != builder {
		t.Errorf("Expected SetColor to return the builder, but got %v", result)
	}

	if builder.car.Color != color {
		t.Errorf("Expected color to be %s, but got %s", color, builder.car.Color)
	}
}

func TestBuilder_Build(t *testing.T) {
	builder := NewCarBuilder()
	make := "Toyota"
	model := "Corolla"
	year := 2020
	color := "Red"

	car := builder.
		SetMake(make).
		SetModel(model).
		SetYear(year).
		SetColor(color).
		Build()

	if car.Make != make {
		t.Errorf("Expected make to be %s, but got %s", make, car.Make)
	}

	if car.Model != model {
		t.Errorf("Expected model to be %s, but got %s", model, car.Model)
	}

	if car.Year != year {
		t.Errorf("Expected year to be %d, but got %d", year, car.Year)
	}

	if car.Color != color {
		t.Errorf("Expected color to be %s, but got %s", color, car.Color)
	}
}