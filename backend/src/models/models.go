package models

const SECRET = "halou"

var TermId int
var OpenSelectCourse int

type Teacher struct {
	Tid       string `json:"tid"`
	Did       string `json:"did"`
	Tname     string `json:"tname"`
	Gender    string `json:"gender"`
	birthday  string `json:"birthday"`
	education string `json:"education"`
	wage      string `json:"wage"`
}
