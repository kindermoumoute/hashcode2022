package simulator

import (
	"fmt"

	"github.com/kindermoumoute/hashcode2022/models"
)

type Busy struct {
	Contributor *models.Contributor
	UntilDay    uint
}

func Simulator(ouput *models.Output, input *models.Input) float64 {
	var (
		finalScore uint
		numDay     uint
	)

	busy := make(map[string]Busy, len(input.Contributors))
	for _, contributor := range input.Contributors {
		busy[contributor.Name] = Busy{Contributor: contributor}
	}

	for _, executedProject := range ouput.ExecutedProjects {
		endProject := numDay

		for numRole, contributor := range executedProject.Contributors {
			role := executedProject.Project.Roles[numRole]
			if !contributor.CanDoRole(role, executedProject.Contributors) {
				panic(fmt.Errorf("contributor %q could not do role %q", contributor.Name, role.RequiredSkill))
			}

			localEndProject := busy[contributor.Name].UntilDay + executedProject.Project.Duration
			if localEndProject > endProject {
				endProject = localEndProject
			}

			contributor.Skills[role.RequiredSkill].Level++
		}

		for _, contributor := range executedProject.Contributors {
			b := busy[contributor.Name]
			b.UntilDay = endProject
			busy[contributor.Name] = b
		}
		finalScore += executedProject.Project.ScoreWithDuration(endProject - 1)
	}

	return float64(finalScore)
}
