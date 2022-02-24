package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/gammazero/workerpool"
	"github.com/kindermoumoute/hashcode2022/logger"
	"github.com/kindermoumoute/hashcode2022/models"
)

var (
	//go:embed input/a_an_example.in.txt
	a string

	//go:embed input/b_better_start_small.in.txt
	b string

	//go:embed input/c_collaboration.in.txt
	c string

	//go:embed input/d_dense_schedule.in.txt
	d string

	//go:embed input/e_exceptional_skills.in.txt
	e string

	//go:embed input/f_find_great_mentors.in.txt
	f string
)

func main() {

	wp := workerpool.New(runtime.NumCPU())
	// Solve each input in a different worker

	tryingAlphas := map[string]string{
		"A": a,
		"B": b,
		"C": c,
		"D": d,
		"E": e,
		"F": f,
	}

	// maxScore := atomic.NewFloat64(0)
	for inputChoice, rawInput := range tryingAlphas {
		inputChoice := inputChoice
		input := models.ParseInput(rawInput)
		wp.Submit(func() {
			output := Solver1(input, Solver1Parameters{})
			// finalScore := simulator.Simulator(output, input)

			// if maxScore.Load() < output.FinalScore(input) { // If new score is found
			// maxScore.Store(finalScore)
			// logger.L.Infof("new high score is %f", finalScore)
			logger.L.Infof("new computed %s", inputChoice)
			assertNoErr(ioutil.WriteFile(path.Join("output", inputChoice+fmt.Sprintf("_%s_latest.txt", inputChoice)), []byte(output.Generate()), 0644))
			// }
		})
	}

	wp.StopWait()
}

func assertNoErr(err error) {
	if err != nil {
		logger.L.Error(err)
		os.Exit(1)
	}
}
