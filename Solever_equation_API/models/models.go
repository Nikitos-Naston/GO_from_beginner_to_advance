package models

type Coef struct {
	A int `json:"A"`
	B int `json:"B"`
	C int `json:"C"`
}

type Answer struct {
	A     int `json:"A"`
	B     int `json:"B"`
	C     int `json:"C"`
	Roots int `json:"roots"`
}
