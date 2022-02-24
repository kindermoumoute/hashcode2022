package main

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/gammazero/workerpool"
	"github.com/kindermoumoute/hashcode2022/logger"
	"github.com/kindermoumoute/hashcode2022/models"
	"go.uber.org/atomic"
)

var (

	//go:embed input/a.txt
	a string

	//go:embed input/b.txt
	b string

	//go:embed input/c.txt
	c string

	//go:embed input/d.txt
	d string

	//go:embed input/e.txt
	e string

	//go:embed input/f.txt
	f string

	allInputs = map[string]string{
		"A": a,
		"B": b,
		"C": c,
		"D": d,
		"E": e,
		"F": f,
	}
)

func main() {
	inputChoice := "A"
	if os.Args[1] != "" {
		inputChoice = strings.ToLower(os.Args[1])
	}
	rawInput, ok := allInputs[inputChoice]
	if !ok {
		assertNoErr(fmt.Errorf("input %s does not exist", inputChoice))
	}
	logger.L = logger.L.Named(inputChoice)
	input := models.ParseInput(rawInput)

	wp := workerpool.New(runtime.NumCPU())
	// Solve each input in a different worker

	tryingAlphas := []float64{0.1}

	maxScore := atomic.NewFloat64(0)
	for _, alpha := range tryingAlphas {
		alpha := alpha
		wp.Submit(func() {
			output := Solver1(input, Solver1Parameters{})
			finalScore := output.FinalScore()

			if maxScore.Load() < output.FinalScore() { // If new score is found
				maxScore.Store(finalScore)
				logger.L.Infof("new high score is %f", output.FinalScore())
				assertNoErr(ioutil.WriteFile(path.Join("output", inputChoice+fmt.Sprintf("_%2f_latest.txt", alpha)), []byte(output.Generate()), 0644))
			}
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
