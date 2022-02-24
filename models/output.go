package models

type Output struct {
	ExecutedProjects []*ExecutedProject
}

func (o *Output) Generate() string {
	return ""
}

func (o *Output) FinalScore(input *Input) {
}
