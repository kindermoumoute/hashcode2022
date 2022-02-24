package models

import (
	"strconv"
	"strings"

	"github.com/kindermoumoute/adventofcode/pkg"
	"github.com/kindermoumoute/hashcode2022/logger"
)

type Input struct {
	Contributors []*Contributor
	// Skills       []*Skill
	Projects []*Project

	MaxDate uint
}

func (i *Input) PrintSomeStats() {
	logger.L.Infof("project count: ", len(i.Projects))
	logger.L.Infof("contributor count: ", len(i.Contributors))
	// logger.L.Infof("contributor count: ", len(i.Contributors))
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
			Name:   contributorLine[0],
			Skills: map[string]*Skill{},
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

			newContributor.Skills[newContribSkill.Name] = newContribSkill
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
			AwardedScore: uint(intList[1]),
		}
		if newProject.BestBefore+newProject.AwardedScore > input.MaxDate {
			input.MaxDate = newProject.BestBefore + newProject.AwardedScore
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
