// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hull

import (
	"sort"

	"github.com/deslaughter/algo"
)

type pointsByXThenY []algo.Point

func (p pointsByXThenY) Len() int      { return len(p) }
func (p pointsByXThenY) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p pointsByXThenY) Less(i, j int) bool {
	if p[i].X == p[j].X {
		return p[i].Y < p[j].Y
	}
	return p[i].X < p[j].X
}

// cross returns the cross product between three points
func cross(p1, p2, p3 algo.Point) float64 {
	return (p2.X-p1.X)*(p3.Y-p1.Y) - (p2.Y-p1.Y)*(p3.X-p1.X)
}

// MonotoneChain implements the Monotone Chain algorithm for calculating the
// convex hull of a slice of points.
//
// References:
//   https://en.wikibooks.org/wiki/Algorithm_Implementation/Geometry/Convex_hull/Monotone_chain
func MonotoneChain(points algo.Points) algo.Points {

	// Sort points by X coordinate, then by the Y coordinate if the
	// X coordinates are equal
	sort.Sort(pointsByXThenY(points))

	// Create points slice to hold hull points
	hull := algo.Points{}

	// Calculate the number of points
	n := len(points)

	k := 0

	// Loop through points and find the lower hull points
	for i := 0; i < n; i++ {
		for k >= 2 && cross(hull[k-2], hull[k-1], points[i]) <= 0 {
			k--
		}
		hull = append(hull[:k], points[i])
		k++
	}

	// Loop through points and find the upper hull points
	for i, t := n-2, k+1; i >= 0; i-- {
		for k >= t && cross(hull[k-2], hull[k-1], points[i]) <= 0 {
			k--
		}
		hull = append(hull[:k], points[i])
		k++
	}

	return hull
}
