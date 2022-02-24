package models

type Project struct {
	Name         string
	Duration     uint    // in days
	BestBefore   uint    // if the project last day of work is strictly before the indicated day, it earns the full score. If it’s late, it gets one point less for each day it is late, but no less than zero points.
	AwardedScore float64 // score awarded for completing the project
	Roles        []*Role
}
