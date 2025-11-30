package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"nbody-problem/enum"
	"nbody-problem/plotter"
	"nbody-problem/sequential"
)

func initBodies(n int, seed int64) []enum.Body {
	rng := rand.New(rand.NewSource(seed))
	bodies := make([]enum.Body, n)
	for i := range bodies {
		bodies[i] = enum.Body{
			X:  2*rng.Float64() - 1,
			Y:  2*rng.Float64() - 1,
			Z:  2*rng.Float64() - 1,
			Vx: 0,
			Vy: 0,
			Vz: 0,
			M:  1.0,
		}
	}
	return bodies
}

func printBodies(bodies []enum.Body) {
	for i, b := range bodies {
		fmt.Printf(
			"Body[%d]: X=%.6f Y=%.6f Z=%.6f  Vx=%.6f Vy=%.6f Vz=%.6f  M=%.2f\n",
			i, b.X, b.Y, b.Z, b.Vx, b.Vy, b.Vz, b.M,
		)
	}
}

func visualizeBodies(bodies []enum.Body) {
	FILENAME := "snap"
	plotter.PlotBodies(bodies, FILENAME)
}

func runSequential(bodies []enum.Body, fx []float64, fy []float64, fz []float64, dt *float64, nSteps *int) {
	for step := 0; step < *nSteps; step++ {
		sequential.ComputeForces(bodies, fx, fy, fz)
		sequential.Integrate(bodies, fx, fy, fz, *dt)

		if step%5 == 0 {
			printBodies(bodies)
			plotter.PlotBodies(bodies, fmt.Sprintf("frame_%04d.png", step))
		}
	}
}

func runParallel(bodies []enum.Body, fx []float64, fy []float64, fz []float64, dt *float64, nSteps *int) {
	//TODO
}

func main() {
	runMode := flag.String("mode", "seq", "mode (sequential/parallel)")
	nBodies := flag.Int("n", 1024, "number of bodies")
	nSteps := flag.Int("steps", 100, "number of time steps")
	dt := flag.Float64("dt", 0.01, "time step")
	seed := flag.Int64("seed", time.Now().UnixNano(), "random seed")

	flag.Parse()

	// Allocate bodies + force buffers
	bodies := initBodies(*nBodies, *seed)
	fx := make([]float64, *nBodies)
	fy := make([]float64, *nBodies)
	fz := make([]float64, *nBodies)

	start := time.Now()

	if *runMode == "seq" {
		runSequential(bodies, fx, fy, fz, dt, nSteps)
	} else {
		runParallel(bodies, fx, fy, fz, dt, nSteps)
	}

	elapsed := time.Since(start)
	fmt.Printf("[%s] Simulated %d bodies for %d steps in %v\n", *runMode, *nBodies, *nSteps, elapsed)
	printBodies(bodies)

	// visualizeBodies(bodies)
}
