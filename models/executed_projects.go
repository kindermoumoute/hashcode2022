package models

type ExecutedProject struct {
	Project      *Project
	Contributors []*Contributor
}
