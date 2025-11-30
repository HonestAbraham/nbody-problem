package sequential

import (
	"math"

	"nbody-problem/enum"
)

func ComputeForces(bodies []enum.Body, fx, fy, fz []float64) {
	const G = 1.0
	const eps2 = 1e-6

	n := len(bodies)
	for i := 0; i < n; i++ {
		bi := bodies[i]
		fxi, fyi, fzi := 0.0, 0.0, 0.0

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			bj := bodies[j]

			dx := bj.X - bi.X
			dy := bj.Y - bi.Y
			dz := bj.Z - bi.Z

			dist2 := dx*dx + dy*dy + dz*dz + eps2
			invDist := 1.0 / math.Sqrt(dist2)
			invDist3 := invDist * invDist * invDist

			F := G * bi.M * bj.M * invDist3

			fxi += F * dx
			fyi += F * dy
			fzi += F * dz
		}

		fx[i] = fxi
		fy[i] = fyi
		fz[i] = fzi
	}
}

func Integrate(bodies []enum.Body, fx, fy, fz []float64, dt float64) {
	for i := range bodies {
		b := &bodies[i]

		// acceleration = F / M
		ax := fx[i] / b.M
		ay := fy[i] / b.M
		az := fz[i] / b.M

		// update velocity
		b.Vx += ax * dt
		b.Vy += ay * dt
		b.Vz += az * dt

		// update position
		b.X += b.Vx * dt
		b.Y += b.Vy * dt
		b.Z += b.Vz * dt
	}
}
