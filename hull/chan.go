// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hull

import (
	"sort"

	"runtime"

	"github.com/deslaughter/algo"
)

// Chan implements Chan's algorithm for calculating the convex hull of a
// slice of points.
//
// References:
//   https://en.wikipedia.org/wiki/Chan%27s_algorithm
func Chan(points algo.Points) algo.Points {

	// Sort points by X coordinate, then by the Y coordinate if the
	// X coordinates are equal
	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	// Get the number of groups from the number of cpus
	groups := runtime.NumCPU()

	pointsChan := make(chan algo.Points)

	chunkSize := (len(points) + groups - 1) / groups

	for i := 0; i < len(points); i += chunkSize {

		end := i + chunkSize

		if end > len(points) {
			end = len(points)
		}

		go func(i, j int) {
			pointsChan <- MonotoneChain(points[i:j])
		}(i, end)
	}

	groupPoints := make(algo.Points, 0)

	for i := 0; i < groups; i++ {
		groupPoints = append(groupPoints, <-pointsChan...)
	}

	hull := MonotoneChain(groupPoints)

	return hull
}
