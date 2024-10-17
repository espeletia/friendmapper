package domain

type Pin struct {
	Lat float64
	Lon float64
}

type Viewport struct {
	NorthWest Pin
	SouthEast Pin
}

type Place struct {
	ID                string
	Type              string
	SubType           *string
	Name              string
	Description       *string
	Accessibility     int64
	AccessibilityNote *string
	Capacity          *int64
	CapacityNote      *string
	Phones            []string
	Email             *string
	Web               string
	Okres             string
	Obce              string
	Address           string
	Point             Pin
}
