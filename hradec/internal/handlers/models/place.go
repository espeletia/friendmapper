package models

type Pin struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type Viewport struct {
	NorthWest Pin `json:"north_west"`
	SouthEast Pin `json:"south_east"`
}

/*
{
  "north_west": {
    "lat": 50.21305216043979,
    "lon": 15.838433504104616
  },
  "south_east": {
    "lat": 50.207415043358225,
    "lon": 15.828391313552858
  }
}*/

type Place struct {
	ID                string   `json:"id"`
	Type              string   `json:"type"`
	SubType           *string  `json:"sub_type,omitempty"`
	Name              string   `json:"name"`
	Description       *string  `json:"description,omitempty"`
	Accessibility     int64    `json:"accessibility"`
	AccessibilityNote *string  `json:"accessibility_note,omitempty"`
	Capacity          *int64   `json:"capacity,omitempty"`
	CapacityNote      *string  `json:"capacity_note,omitempty"`
	Phones            []string `json:"phones"`
	Email             *string  `json:"email,omitempty"`
	Web               string   `json:"web"`
	Okres             string   `json:"okres"`
	Obce              string   `json:"obce"`
	Address           string   `json:"address"`
	Point             Pin      `json:"point"`
}

