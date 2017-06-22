// Copyright 2017 Derek Slaughter. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hull

import (
	"math/rand"
	"testing"
	"time"

	"sort"

	"github.com/deslaughter/algo"
)

var randomPoints algo.Points
var smallPoints = algo.Points{
	algo.Point{X: -1, Y: 0},
	algo.Point{X: 0, Y: -1},
	algo.Point{X: 1, Y: 0},
	algo.Point{X: 0, Y: 1},
	algo.Point{X: 0, Y: 0},
}

func init() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomPoints = make(algo.Points, 1000)

	for i := range randomPoints {
		randomPoints[i] = algo.Point{X: r.Float64(), Y: r.Float64()}
	}

	sort.Sort(pointsByXThenY(smallPoints))
	sort.Sort(pointsByXThenY(randomPoints))
}

func TestMonotoneChain(t *testing.T) {

	hull := MonotoneChain(smallPoints)

	if len(hull) != 5 {
		t.Fatalf("len(hull) = %d, expected 5", len(hull))
	}
	if hull[0].X != -1 || hull[0].Y != 0 {
		t.Fatalf("hull[0] = %v, expected {-1, 0}", hull[0])
	}
	if hull[1].X != 0 || hull[1].Y != -1 {
		t.Fatalf("hull[1] = %v, expected {0, -1}", hull[1])
	}
	if hull[2].X != 1 || hull[2].Y != 0 {
		t.Fatalf("hull[2] = %v, expected {1, 0}", hull[2])
	}
	if hull[3].X != 0 || hull[3].Y != 1 {
		t.Fatalf("hull[3] = %v, expected {0, 1}", hull[3])
	}
	if hull[4].X != -1 || hull[4].Y != 0 {
		t.Fatalf("hull[4] = %v, expected {-1, 0}", hull[4])
	}
}

func BenchmarkMonotoneChainSmall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MonotoneChain(smallPoints)
	}
}

func BenchmarkMonotoneChainRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MonotoneChain(randomPoints)
	}
}
