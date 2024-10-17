package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Pin struct {
	Lat float64
	Lon float64
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

func PrettyPrintPlace(place Place) {
	fmt.Printf("Place Details:\n")
	fmt.Printf("ID: %s\n", place.ID)
	fmt.Printf("Type: %s\n", place.Type)
	if place.SubType != nil {
		fmt.Printf("SubType: %s\n", *place.SubType)
	} else {
		fmt.Printf("SubType: None\n")
	}
	fmt.Printf("Name: %s\n", place.Name)
	if place.Description != nil {
		fmt.Printf("Description: %s\n", *place.Description)
	} else {
		fmt.Printf("Description: None\n")
	}
	fmt.Printf("Accessibility: %t\n", place.Accessibility)
	if place.AccessibilityNote != nil {
		fmt.Printf("Accessibility Note: %s\n", *place.AccessibilityNote)
	} else {
		fmt.Printf("Accessibility Note: None\n")
	}
	if place.Capacity != nil {
		fmt.Printf("Capacity: %d\n", *place.Capacity)
	} else {
		fmt.Printf("Capacity: None\n")
	}
	if place.CapacityNote != nil {
		fmt.Printf("Capacity Note: %s\n", *place.CapacityNote)
	} else {
		fmt.Printf("Capacity Note: None\n")
	}
	if len(place.Phones) > 0 {
		fmt.Printf("Phones: %s\n", strings.Join(place.Phones, ", "))
	} else {
		fmt.Printf("Phones: None\n")
	}
	if place.Email != nil {
		fmt.Printf("Email: %s\n", *place.Email)
	} else {
		fmt.Printf("Email: None\n")
	}
	fmt.Printf("Web: %s\n", place.Web)
	fmt.Printf("Okres: %s\n", place.Okres)
	fmt.Printf("Obce: %s\n", place.Obce)
	fmt.Printf("Address: %s\n", place.Address)
	fmt.Printf("Location: Latitude: %.6f, Longitude: %.6f\n", place.Point.Lat, place.Point.Lon)
}

func run() error {
	// items, err := parseMusicAndFestivals()
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%+v", items)
	getIndexes("../hradec/migrations/Zámky_3381043052573979413.csv")
	// got, err := parsePlaces("../hradec/migrations/Hudební_kluby_a_festival_parky_-7703215986117040335.csv",
	// "MUSIC", 0, 26, toPtr(3), 7, 22, 11, 15, 25, 24, 17, 18, nil, nil, toPtr(8), toPtr(4), toPtr(5), toPtr(6))

	// got, err := parsePlaces("../hradec/migrations/Kina_7325494378020973866.csv",
	// 	"CINEMA", 0, 25, nil, 3, 21, 10, 14, 24, 23, 16, 17, nil, nil, toPtr(4), toPtr(5), nil, nil)
	// places, err := parsePlaces("../hradec/migrations/Národní_kulturní_památky_bodové_objekty_6158076928879444772.csv",
	// 	"MONUMENTS",
	// 	0, 22, toPtr(1), 0, 18, 7, 11, 21, 20, 13, 14, nil, nil, nil, nil, nil, nil)

	// places, err := parsePlaces("../hradec/migrations/Pivovary_-340792773380756746.csv",
	// 	"BEER",
	// 	0, 29, toPtr(5), 4, 19, 8, 12, 28, 27, 14, 15, []int{20, 21, 22}, toPtr(24), nil, toPtr(2))
	// places, err := parsePlaces("../hradec/migrations/Divadla_a_filharmonie_2818140482217729379.csv",
	// 	"THEATRE",
	// 	0, 27, toPtr(2), 6, 21, 10, 14, 26, 25, 16, 17, []int{22}, toPtr(23), toPtr(7), toPtr(3), toPtr(5), nil)
	// places, err := parsePlaces("../hradec/migrations/Muzea_a_galerie_-156929117348245312.csv",
	// 	"MUSEUM_GALLERY",
	// 	0, 24, toPtr(6), 3, 20, 9, 13, 23, 22, 15, 16, nil, nil, toPtr(4), toPtr(5), nil, nil)
	// places, err := parsePlaces("../hradec/migrations/Technické_památky_-7149205952045496218.csv",
	// 	"MONUMENTS",
	// 	0, 20, toPtr(1), 0, 15, 4, 8, 19, 18, 10, 11, []int{16}, nil, nil, nil, nil, nil)
	// places, err := parsePlaces("../hradec/migrations/Letecké_sporty_8959676097614550187.csv",
	// 	"SUMMER_SPORTS",
	// 	0, 30, toPtr(5), 3, 19, 8, 12, 29, 28, 14, 15, []int{20, 21, 22, 23}, toPtr(25), nil, toPtr(2), nil, nil)
	places, err := parsePlaces("../hradec/migrations/Zámky_3381043052573979413.csv",
		"CASTLES",
		0, 19, toPtr(1), 0, 15, 4, 8, 18, 17, 10, 11, nil, nil, nil, nil, nil, nil)

	if err != nil {
		return err
	}
	for _, place := range places {
		PrettyPrintPlace(place)
		fmt.Scanln()
	}
	return nil
}

func toPtr[T any](x T) *T {
	return &x
}

func getIndexes(filepath string) error {
	openFile, err := os.Open(filepath)
	if err != nil {
		return err
	}
	reader := csv.NewReader(openFile)
	rows, err := reader.ReadAll()
	if err != nil {
		return err
	}
	length := 0
	for i, row := range rows {
		if len(row) != length && length != 0 {
			panic("unnormalized csv")
		}
		if i < 1 {
			for j, item := range row {
				fmt.Printf("%d: %s\n", j, item)
			}
		}
	}
	return nil
}

func parsePlaces(filepath, fileType string,
	nameIdx, idIdx int, // Music: 26, 0; Kino: 25, 0
	descriptionIdx *int, accessibilityIdx, // Music: 3; Kino: None, 3;
	webIdx, okresIdx, // Music: 22, 11; Kino: 21, 10
	obceIdx, latIdx, // Music: 15, 25; Kino: 14, 24
	lonIdx, streetIdx, // Music: 24, 17; Kino: 23, 16
	streetNumIdx int, // Music: 18; Kino: 17
	phoneIdxs []int, // Music: None, Kino: None
	emailIdx, accessibilityNoteIdx, subTypeIdx,
	capacityIdx, capacityNoteIdx *int, // music: None, 8, 4; Kino: None, 4, 5
) ([]Place, error) {
	openFile, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(openFile)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	length := 0
	result := []Place{}
	for i, row := range rows {
		if len(row) != length && length != 0 {
			panic("unnormalized csv")
		}
		if i < 1 {
			for j, item := range row {
				fmt.Printf("%d: %s\n", j, item)
			}
		}
		if i > 0 {
			accessibility := 0
			if row[accessibilityIdx] == "ano" || row[accessibilityIdx] == "ANO" || strings.Contains(row[accessibilityIdx], "Bez bariér") {
				accessibility = 1
			} else if row[accessibilityIdx] == "ČÁSTEČNĚ" || row[accessibilityIdx] == "částečně" {
				accessibility = 2
			} else if row[accessibilityIdx] == "ne" || row[accessibilityIdx] == "NE" {
				accessibility = 3
			}
			lat, err := strconv.ParseFloat(row[latIdx], 64)
			if err != nil {
				return nil, err
			}
			lon, err := strconv.ParseFloat(row[lonIdx], 64)
			if err != nil {
				return nil, err
			}

			var subtype *string = nil
			if subTypeIdx != nil {
				subtype = &row[*subTypeIdx]
			}

			var accessibilityNote *string = nil
			if accessibilityNoteIdx != nil {
				accessibilityNote = &row[*accessibilityNoteIdx]
			}
			var description *string = nil
			if descriptionIdx != nil {
				description = &row[*descriptionIdx]
			}
			var email *string = nil
			if emailIdx != nil {
				email = &row[*emailIdx]
			}
			var phones []string = nil
			if phoneIdxs != nil {
				phones = []string{}
				for _, idx := range phoneIdxs {
					normalizedPhone, valid := normalizeCzechPhone(row[idx])
					if valid {
						phones = append(phones, normalizedPhone)
					}
				}
			}
			var capacity *int64 = nil
			if capacityIdx != nil {
				capacityInt, err := strconv.ParseInt(row[*capacityIdx], 10, 64)
				if err != nil {
					return nil, err
				}
				capacity = &capacityInt
			}

			var capacityNote *string = nil
			if capacityNoteIdx != nil {
				capacityNote = &row[*capacityNoteIdx]
			}

			address := fmt.Sprintf("%s %s", row[streetIdx], row[streetNumIdx])

			result = append(result, Place{
				ID:                row[idIdx],
				Type:              fileType,
				SubType:           subtype,
				Name:              row[nameIdx],
				Description:       description,
				Accessibility:     accessibility,
				AccessibilityNote: accessibilityNote,
				Address:           address,
				Web:               row[webIdx],
				Phones:            phones,
				Email:             email,
				Okres:             row[okresIdx],
				Obce:              row[obceIdx],
				Capacity:          capacity,
				CapacityNote:      capacityNote,
				Point: Pin{
					Lat: lat,
					Lon: lon,
				},
			})
		}
	}
	return result, nil
}

func parseMusicAndFestivals() ([]Place, error) {
	openFile, err := os.Open("./../hradec/migrations/Hudební_kluby_a_festival_parky_-7703215986117040335.csv")
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(openFile)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	length := 0
	result := []Place{}
	for i, row := range rows {
		if len(row) != length && length != 0 {
			panic("unnormalized csv")
		}
		for j, item := range row {
			fmt.Printf("%d: %s\n", j, item)
		}
		if i > 0 {
			accessibility := false
			if row[7] == "ano" {
				accessibility = true
			}
			lat, err := strconv.ParseFloat(row[25], 64)
			if err != nil {
				return nil, err
			}
			lon, err := strconv.ParseFloat(row[24], 64)
			if err != nil {
				return nil, err
			}

			address := fmt.Sprintf("%s %s", row[17], row[18])

			result = append(result, Place{
				ID:                row[26],
				Type:              "MUSIC",
				SubType:           &row[4],
				Name:              row[0],
				Description:       &row[3],
				Accessibility:     accessibility,
				AccessibilityNote: &row[8],
				Address:           address,
				Web:               row[22],
				Phones:            nil,
				Email:             nil,
				Okres:             row[11],
				Obce:              row[15],
				Point: Pin{
					Lat: lat,
					Lon: lon,
				},
			})
		}
	}
	return result, nil
}

func parseCinemas() ([]Place, error) {
	openFile, err := os.Open("./../hradec/migrations/Kina_7325494378020973866.csv")
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(openFile)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	length := 0
	result := []Place{}
	for i, row := range rows {
		if len(row) != length && length != 0 {
			panic("unnormalized csv")
		}
		for j, item := range row {
			fmt.Printf("%d: %s\n", j, item)
		}
		if i > 0 {
			accessibility := false
			if row[3] == "ANO" || row[3] == "ČÁSTEČNĚ" {
				accessibility = true
			}
			lat, err := strconv.ParseFloat(row[24], 64)
			if err != nil {
				return nil, err
			}
			lon, err := strconv.ParseFloat(row[23], 64)
			if err != nil {
				return nil, err
			}

			address := fmt.Sprintf("%s %s", row[16], row[17])
			result = append(result, Place{
				ID:                row[25],
				Type:              "CINEMA",
				SubType:           &row[5],
				Name:              row[0],
				Description:       nil,
				Accessibility:     accessibility,
				AccessibilityNote: &row[4],
				Phones:            nil,
				Email:             nil,
				Address:           address,
				Web:               row[21],
				Okres:             row[10],
				Obce:              row[14],
				Point: Pin{
					Lat: lat,
					Lon: lon,
				},
			})
		}
	}
	return result, nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func normalizeCzechPhone(phone string) (string, bool) {
	// Remove any non-numeric characters except the leading "+"
	phone = regexp.MustCompile(`[^\d+]`).ReplaceAllString(phone, "")

	// Check if the phone number starts with a valid Czech code (+420 or just 420)
	if strings.HasPrefix(phone, "+420") {
		phone = phone
	} else if strings.HasPrefix(phone, "420") {
		phone = "+" + phone
	} else if len(phone) == 9 {
		// Assuming local 9-digit Czech numbers without country code
		phone = "+420" + phone
	} else {
		// Invalid number, return false
		return "", false
	}

	// Validate if the phone number has the correct format and length after normalization
	if matched, _ := regexp.MatchString(`^\+420\d{9}$`, phone); matched {
		return phone, true
	}

	// If it doesn't match the expected format, return false
	return "", false
}
