package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	in := ParseInput(`3 3
Anna 1
C++ 2
Bob 2
HTML 5
CSS 5
Maria 1
Python 3
Logging 5 10 5 1
C++ 3
WebServer 7 10 7 2
HTML 3
C++ 2
WebChat 10 20 20 2
Python 3
HTML 3
`)

	require.Len(t, in.Contributors, 3)
	require.Len(t, in.Projects, 3)
	assert.Equal(t, "Logging", in.Contributors[0].Name)
	assert.Equal(t, Skill{
		Name:  "C++",
		Level: 2,
	}, *in.Contributors[0].Skills[0])

	assert.Equal(t, "Anna", in.Projects[0].Name)
	assert.Equal(t, "Anna", in.Projects[0].AwardedScore)
	assert.Equal(t, "Anna", in.Projects[0].BestBefore)
	assert.Equal(t, "Anna", in.Projects[0].Duration)
	assert.Equal(t, "Anna", in.Projects[0].Roles[0].ID)
	assert.Equal(t, "Anna", in.Projects[0].Roles[0].RequiredLevel)
	assert.Equal(t, "Anna", in.Projects[0].Roles[0].RequiredSkill)
}
