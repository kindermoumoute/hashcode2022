package main

import (
	"sort"

	"github.com/kindermoumoute/hashcode2022/models"
)

type Solver1Parameters struct {
}

func Solver1(in *models.Input, params Solver1Parameters) *models.Output { // could return an intermediary Solution1 struct if needed
	remainingProjects := in.Projects
	availableContributors := make(map[string]*models.Contributor)
	executedProjects := []*models.ExecutedProject(nil)
	for _, contributor := range in.Contributors {
		availableContributors[contributor.Name] = contributor
	}
	planningAvailability := make(map[uint][]*models.Project)
	for day := uint(0); day < in.MaxDate; day++ {
		dayPlanning, exist := planningAvailability[day]
		if day != 0 && !exist {
			continue
		}
		// Resets CurrentScore for all projects
		for _, project := range remainingProjects {
			project.Score(day)
		}
		for _, project := range dayPlanning {
			executedProject := &models.ExecutedProject{
				Project: project,
			}
			for _, role := range project.Roles {
				executedProject.Contributors = append(executedProject.Contributors, role.AssignedContributor)
				role.AssignedContributor.AssignedToProject = nil
				availableContributors[role.AssignedContributor.Name] = role.AssignedContributor

				oldScore := role.AssignedContributor.Skills[role.RequiredSkill].Level
				requiredScore := role.RequiredLevel
				if oldScore <= requiredScore {
					role.AssignedContributor.Skills[role.RequiredSkill].Level++
				}
			}
			executedProjects = append(executedProjects, executedProject)
		}

		// Sort remaining projects by score

		sort.Slice(remainingProjects, func(i, j int) bool {
			return remainingProjects[i].Score(day) > remainingProjects[j].Score(day)
		})
		newRemainingProjects := []*models.Project(nil)
		for _, project := range remainingProjects {
			if project.Assign(availableContributors) {
				for _, role := range project.Roles {
					delete(availableContributors, role.AssignedContributor.Name)
				}
				endDay := day + project.Duration + 1
				planningAvailability[endDay] = append(planningAvailability[endDay], project)
			} else {
				newRemainingProjects = append(newRemainingProjects, project)
			}
		}

		remainingProjects = newRemainingProjects
	}

	return &models.Output{
		ExecutedProjects: executedProjects,
	}
}
