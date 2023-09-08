// Package yooreceipt describes all the necessary entities for working with YooMoney Receipts.
package yooreceipt

// Unit of measurement of product quantity
type Measure string

const (
	Piece            Measure = "piece"
	Gram             Measure = "gram"
	Kilogram         Measure = "kilogram"
	Ton              Measure = "ton"
	Centimeter       Measure = "centimeter"
	Decimeter        Measure = "decimeter"
	Meter            Measure = "meter"
	SquareCentimeter Measure = "square_centimeter"
	SquareDecimeter  Measure = "square_decimeter"
	SquareMeter      Measure = "square_meter"
	Milliliter       Measure = "milliliter"
	Liter            Measure = "liter"
	CubicMeter       Measure = "cubic_meter"
	KilowattHour     Measure = "kilowatt_hour"
	Gigacalorie      Measure = "gigacalorie"
	Day              Measure = "day"
	Hour             Measure = "hour"
	Minute           Measure = "minute"
	Second           Measure = "second"
	Kilobyte         Measure = "kilobyte"
	Megabyte         Measure = "megabyte"
	Gigabyte         Measure = "gigabyte"
	Terabyte         Measure = "terabyte"
	Other            Measure = "other"
)
