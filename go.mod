module github.com/kindermoumoute/hashcode2022

go 1.17

require (
	github.com/gammazero/workerpool v1.1.2
	github.com/kindermoumoute/adventofcode v0.0.0-20201225075527-682ab4e0f685
	github.com/stretchr/testify v1.7.0
	go.uber.org/atomic v1.9.0
	go.uber.org/zap v1.21.0
	gonum.org/v1/gonum v0.9.3
)

require (
	github.com/beefsack/go-astar v0.0.0-20171024231011-f324bbb0d6f7 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gammazero/deque v0.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/kindermoumoute/hashcode2022 v0.0.0 => .

replace github.com/kindermoumoute/hashcode2022/models v0.0.0 => ./models

replace github.com/kindermoumoute/hashcode2022/logger v0.0.0 => ./logger

replace github.com/kindermoumoute/hashcode2022/simulator v0.0.0 => ./simulator
