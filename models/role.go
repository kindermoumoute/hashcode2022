package models

type Role struct {
	ID            int
	Project       *Project
	RequiredSkill string
	RequiredLevel uint

	AssignedContributor *Contributor
}

func (r *Role) MentorExists(contributors []*Contributor) bool {
	for _, contributor := range contributors {
		skill := contributor.FindSkill(r.RequiredSkill)
		if skill.Level >= r.RequiredLevel {
			return true
		}
	}

	return false
}

func (r *Role) MentorExistsMap(contributors map[string]*Contributor, contribUsed map[string]bool) bool {
	for _, contributor := range contributors {
		skill := contributor.FindSkill(r.RequiredSkill)
		if skill.Level >= r.RequiredLevel && contribUsed[contributor.Name] {
			return true
		}
	}

	return false
}
