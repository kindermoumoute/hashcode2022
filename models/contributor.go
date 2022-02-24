package models

import (
	"math"
)

type Contributor struct {
	Name   string
	Skills map[string]*Skill

	AssignedToProject *Project
	SkillLevelsSum    uint
}

func (c *Contributor) Score(project *Project) float64 {
	var score uint

	// Int.Max - (c.skillLevel - role.RequiredLevel + c.skills.map(sum(skillLevel)))
	for _, role := range project.Roles {
		skill, exists := c.Skills[role.RequiredSkill]
		if !exists {
			skill = &Skill{
				Name:  role.RequiredSkill,
				Level: 0,
			}
			c.Skills[role.RequiredSkill] = skill
		}

		if skill.Level-1 < role.RequiredLevel {
			continue
		}

		score += (math.MaxUint / 2) - c.SumSkills() - (skill.Level - role.RequiredLevel)
	}

	return float64(score)
}

func (c *Contributor) SumSkills() uint {
	if c.SkillLevelsSum != 0 {
		return c.SkillLevelsSum
	}

	var score uint

	for _, skill := range c.Skills {
		score += skill.Level
	}

	return score
}

func (c *Contributor) FindSkill(SkillNameToFind string) *Skill {
	if _, exists := c.Skills[SkillNameToFind]; exists {
		return c.Skills[SkillNameToFind]
	}

	c.Skills[SkillNameToFind] = &Skill{
		Name:  SkillNameToFind,
		Level: 0,
	}

	return c.Skills[SkillNameToFind]
}

func (c *Contributor) CanDoRole(role *Role, contributors []*Contributor) bool {
	skill := c.FindSkill(role.RequiredSkill)

	if skill.Level < role.RequiredLevel {
		if skill.Level+1 < role.RequiredLevel {
			return false
		}

		// mentor needed
		if !role.MentorExists(contributors) {
			return false
		}
	}

	return true
}
