package models

type Output struct {
}

func (o *Output) Generate() string {
	return ""
}

func (o *Output) FinalScore() float64 {
	return 0
}
