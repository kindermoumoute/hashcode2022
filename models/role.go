package models

type Role struct {
	ID            int
	Project       *Project
	RequiredSkill string
	RequiredLevel uint

	// TODO: find possible contributor for this role?
}
