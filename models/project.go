package models

type Project struct {
	Name         string
	Duration     uint // in days
	BestBefore   uint // if the project last day of work is strictly before the indicated day, it earns the full score. If it’s late, it gets one point less for each day it is late, but no less than zero points.
	AwardedScore uint // score awarded for completing the project
	Roles        []*Role

	Assigned      bool
	AssignedScore float64
	CurrentScore  float64
}

/*
	Computes the project score at current time
*/
func (p *Project) Score(currentTime uint) float64 {
	if p.Assigned {
		return 0
	}

	// Total points obtained on completion
	points := p.AwardedScore + (p.BestBefore - (currentTime + p.Duration))
	p.CurrentScore = float64(points) / float64(p.Duration*uint(len(p.Roles)))
	return p.CurrentScore
}

/*
	1- Chooses contributors
	2- Assigns them to project
	3- Sets assignation flag to true
*/
func (p *Project) Assign(availableContributors map[string]*Contributor) bool {
	// sortedContributors := []*Contributor(nil)
	// for _, contributor := range availableContributors {
	// 	sortedContributors = append(sortedContributors, contributor)
	// }
	// sort.Slice(sortedContributors, func(i, j int) bool {
	// 	return sortedContributors[i].Score(p) > sortedContributors[j].Score(p)
	// })
	team := make(map[int]*Contributor) // index is role ID
	contribUsed := make(map[string]bool)
roleLoop:
	for _, role := range p.Roles {
		hasContributor := false
		for _, contributor := range availableContributors { // random ordering
			if contributor.AssignedToProject != nil || contribUsed[contributor.Name] {
				continue
			}

			// Add to the team if has enough skills
			contributorSkill, hasSkill := contributor.Skills[role.RequiredSkill]
			if !hasSkill {
				contributor.Skills[role.RequiredSkill] = &Skill{
					Name:  role.RequiredSkill,
					Level: 0,
				}
				contributorSkill = contributor.Skills[role.RequiredSkill]
			}

			if (role.RequiredLevel <= contributorSkill.Level) ||
				role.RequiredLevel-1 == contributorSkill.Level && role.MentorExistsMap(availableContributors, contribUsed) {
				team[role.ID] = contributor
				hasContributor = true
				contribUsed[contributor.Name] = true
				continue roleLoop
			}
		}
		if !hasContributor {
			return false
		}
	}

	for roleID, contributor := range team {
		contributor.AssignedToProject = p
		p.Roles[roleID].AssignedContributor = contributor
	}

	for _, contributor := range team {
		contributor.AssignedToProject = p
	}

	p.Assigned = true
	p.AssignedScore = p.CurrentScore
	return true
}

func (p *Project) ScoreWithDuration(duration uint) uint {
	if duration > p.Duration {
		return 0
	}

	if duration < p.BestBefore {
		return p.AwardedScore
	}

	return p.AwardedScore - (duration - p.BestBefore)
}
