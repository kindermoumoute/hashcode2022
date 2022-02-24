package models

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
)

type Input struct {
	Contributors []*Contributor
	// Skills       []*Skill
	Projects []*Project
}

func ParseInput(s string) *Input {
	input := &Input{}
	lines := strings.Split(s, "\n")
	line0 := pkg.ParseIntList(lines[0], " ")
	i := 1
	for contributorCount := line0[0]; contributorCount > 0; contributorCount-- {
		contributorLine := strings.Split(lines[i], " ")
		i++
		newContributor := &Contributor{
			Name: contributorLine[0],
		}
		skillsCount, _ := strconv.Atoi(contributorLine[1])
		for ; skillsCount > 0; skillsCount-- {
			skillLine := strings.Split(lines[i], " ")
			i++
			newContribSkill := &Skill{
				Name: skillLine[0],
			}
			tmp, _ := strconv.Atoi(skillLine[1])
			newContribSkill.Level = uint(tmp)

			newContributor.Skills = append(newContributor.Skills, newContribSkill)
		}
		input.Contributors = append(input.Contributors, newContributor)
	}

	for projectCount := line0[1]; projectCount > 0; projectCount-- {
		projectLine := strings.Split(lines[i], " ")
		i++
		intList := pkg.ParseIntList(strings.Join(projectLine[1:], " "), " ")

		newProject := &Project{
			Name:         projectLine[0],
			Duration:     uint(intList[0]),
			BestBefore:   uint(intList[2]),
			AwardedScore: float64(intList[1]),
		}

		roleID := 0
		for roleCount := intList[3]; roleCount > 0; roleCount-- {
			roleLine := strings.Split(lines[i], " ")
			i++
			newRole := &Role{
				ID:            roleID,
				Project:       newProject,
				RequiredSkill: roleLine[0],
				RequiredLevel: uint(pkg.MustAtoi(roleLine[1])),
			}

			newProject.Roles = append(newProject.Roles, newRole)
			roleID++
		}
		input.Projects = append(input.Projects, newProject)
	}
	return input
}
