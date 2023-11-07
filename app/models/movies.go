package models

type Movie struct {
	Title       string
	Description string
	duration    uint8
	// Kinds map[]string
}

// type Movies []Movie

type Movies []Movie

func (u Movie) TableName() string {
	return "Movies"
}
