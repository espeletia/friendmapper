package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ifthisWorks(t *testing.T) {
	t.Run("plswork", func(t *testing.T) {
		items, err := parseMusicAndFestivals()
		if err != nil {
			t.Error(err)
			return
		}
		got, err := parsePlaces("../hradec/migrations/Hudební_kluby_a_festival_parky_-7703215986117040335.csv",
			"MUSIC", 0, 26, toPtr(3), 7, 22, 11, 15, 25, 24, 17, 18, nil, nil, toPtr(8), toPtr(4), toPtr(5), toPtr(6))
		assert.Equal(t, items, got)
	})
	t.Run("plswork2", func(t *testing.T) {
		items, err := parseCinemas()
		if err != nil {
			t.Error(err)
			return
		}
		got, err := parsePlaces("../hradec/migrations/Kina_7325494378020973866.csv",
			"CINEMA", 0, 25, nil, 3, 21, 10, 14, 24, 23, 16, 17, nil, nil, toPtr(4), toPtr(5), nil, nil)
		if err != nil {
			t.Error(err)
			return
		}

		assert.Equal(t, items, got)
	})

	places, err := parsePlaces("../hradec/migrations/Ostatní_letní_sporty_-6176368377486830231.csv",
		"SUMMER_SPORTS",
		0, 32, toPtr(5), 4, 19, 8, 12, 31, 30, 14, 15, []int{20, 21, 22, 23, 24}, toPtr(26), nil, toPtr(2), nil, nil)
	places, err = parsePlaces("../hradec/migrations/Letecké_sporty_8959676097614550187.csv",
		"SUMMER_SPORTS",
		0, 30, toPtr(5), 3, 19, 8, 12, 29, 28, 14, 15, []int{20, 21, 22, 23}, toPtr(25), nil, toPtr(2), nil, nil)

	places, err = parsePlaces("../hradec/migrations/Pivovary_-340792773380756746.csv",
		"BEER",
		0, 29, toPtr(5), 4, 19, 8, 12, 28, 27, 14, 15, []int{20, 21, 22}, toPtr(24), nil, toPtr(2), nil, nil)

	places, err = parsePlaces("../hradec/migrations/Národní_kulturní_památky_bodové_objekty_6158076928879444772.csv",
		"CULTURAL_RESERVATIONS",
		0, 22, toPtr(1), 0, 18, 7, 11, 21, 20, 13, 14, nil, nil, nil, nil, nil, nil)

	places, err = parsePlaces("../hradec/migrations/Divadla_a_filharmonie_2818140482217729379.csv",
		"THEATRE",
		0, 27, toPtr(2), 6, 21, 10, 14, 26, 25, 16, 17, []int{22}, toPtr(23), toPtr(7), toPtr(3), toPtr(5), nil)

	places, err = parsePlaces("../hradec/migrations/Muzea_a_galerie_-156929117348245312.csv",
		"MUSEUM_GALLERY",
		0, 24, toPtr(6), 3, 20, 9, 13, 23, 22, 15, 16, nil, nil, toPtr(4), toPtr(5), nil, nil)

	places, err = parsePlaces("../hradec/migrations/Technické_památky_-7149205952045496218.csv",
		"TECHNICAL_RESERVATIONS",
		0, 20, toPtr(1), 0, 15, 4, 8, 19, 18, 10, 11, []int{16}, nil, nil, nil, nil, nil)

	places, err = parsePlaces("../hradec/migrations/Zámky_3381043052573979413.csv",
		"CASTLES",
		0, 19, toPtr(1), 0, 15, 4, 8, 18, 17, 10, 11, nil, nil, nil, nil, nil, nil)

}
