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
	"github.com/kindermoumoute/hashcode2022/simulator"
	"go.uber.org/atomic"
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
	inputChoice := "C"
	logger.L = logger.L.Named(inputChoice)
	input := models.ParseInput(c)

	wp := workerpool.New(runtime.NumCPU())
	// Solve each input in a different worker

	tryingAlphas := []float64{0.1}

	maxScore := atomic.NewFloat64(0)
	for _, alpha := range tryingAlphas {
		alpha := alpha
		wp.Submit(func() {
			output := Solver1(input, Solver1Parameters{})
			finalScore := simulator.Simulator(output, input)

			// if maxScore.Load() < output.FinalScore(input) { // If new score is found
			maxScore.Store(finalScore)
			logger.L.Infof("new high score is %f", finalScore)
			assertNoErr(ioutil.WriteFile(path.Join("output", inputChoice+fmt.Sprintf("_%2f_latest.txt", alpha)), []byte(output.Generate()), 0644))
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
