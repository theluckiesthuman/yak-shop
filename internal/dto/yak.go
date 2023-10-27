package dto

type Herd struct {
	Yaks []Yak `xml:"labyak" json:"herd"`
}

type Yak struct {
	Name          string  `xml:"name,attr" json:"name"`
	Age           float64 `xml:"age,attr" json:"age"`
	Sex           string  `xml:"sex,attr" json:"sex,omitempty"`
	AgeLastShaved float64 `json:"age-last-shaved,omitempty"`
}
