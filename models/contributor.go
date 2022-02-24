package models

type Contributor struct {
	Name   string
	Skills []*Skill

	AssignedToProject *Project
}
