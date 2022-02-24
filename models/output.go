package models

import (
	"strconv"
	"strings"
)

type Output struct {
	ExecutedProjects []*ExecutedProject
}

func (o *Output) Generate() string {
	s := strconv.Itoa(len(o.ExecutedProjects)) + "\n"
	for _, project := range o.ExecutedProjects {
		s += project.Project.Name + "\n"
		contributors := []string(nil)
		for _, contributor := range project.Contributors {
			contributors = append(contributors, contributor.Name)
		}
		s += strings.Join(contributors, " ") + "\n"
	}
	return s
}

func (o *Output) FinalScore(input *Input) float64 {
	return 0
}
