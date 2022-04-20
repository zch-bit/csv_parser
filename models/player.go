package models

type Player struct {
	PlayerID     string `json:"player_id" gorm:"primary_key"`
	BirthYear    string `json:"birth_year"`
	BirthMonth   string `json:"birth_month"`
	BirthDay     string `json:"birth_day"`
	BirthCountry string `json:"birth_country"`
	BirthState   string `json:"birth_state"`
	BirthCity    string `json:"birth_city"`
	DeathYear    string `json:"death_year"`
	DeathMonth   string `json:"death_month"`
	DeathDay     string `json:"death_day"`
	DeathCountry string `json:"death_country"`
	DeathState   string `json:"death_state"`
	DeathCity    string `json:"death_city"`
	NameFirst    string `json:"name_first"`
	NameLast     string `json:"name_last"`
	NameGiven    string `json:"name_given"`
	Weight       string `json:"weight"`
	Height       string `json:"height"`
	Bats         string `json:"bats"`
	Throws       string `json:"throws"`
	Debut        string `json:"debut"`
	FinalGame    string `json:"final_game"`
	RetroID      string `json:"retro_id"`
	BbrefID      string `json:"bbref_id"`
}

func ParsePlayer(record []string) *Player {
	return &Player{
		PlayerID:     record[0],
		BirthYear:    record[1],
		BirthMonth:   record[2],
		BirthDay:     record[3],
		BirthCountry: record[4],
		BirthState:   record[5],
		BirthCity:    record[6],
		DeathYear:    record[7],
		DeathMonth:   record[8],
		DeathDay:     record[9],
		DeathCountry: record[10],
		DeathState:   record[11],
		DeathCity:    record[12],
		NameFirst:    record[13],
		NameLast:     record[14],
		NameGiven:    record[15],
		Weight:       record[16],
		Height:       record[17],
		Bats:         record[18],
		Throws:       record[19],
		Debut:        record[20],
		FinalGame:    record[21],
		RetroID:      record[22],
		BbrefID:      record[23],
	}
}
