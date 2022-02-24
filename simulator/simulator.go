package simulator

import "github.com/kindermoumoute/hashcode2022/models"

type Busy struct {
	Contributor *models.Contributor
	UntilDay    int
}

func Simulator(ouput *models.Output, input *models.Input) float64 {
	var (
		finalScore int
		numDay     int
	)

	busy := make(map[string]Busy, len(input.Contributors))
	for _, contributor := range input.Contributors {
		busy[contributor.Name] = Busy{Contributor: contributor}
	}

	for _, project := range ouput.ExecutedProjects {
		daySoonerFree := numDay

		for numRole, contributor := range project.Contributors {
			if _, exists := busy[contributor.Name]; !exists {
				busy[contributor.Name] = Busy{
					Contributor: contributor,
					UntilDay:    numDay,
				}
			}

		}
	}

	return float64(finalScore)
}
