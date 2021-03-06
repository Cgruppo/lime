// Copyright 2013 The lime Authors.
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package commands

import (
	. "github.com/limetext/text"
	"testing"
)

func TestSingleSelection(t *testing.T) {
	tests := []findTest{
		{
			[]Region{{1, 1}, {2, 2}, {3, 3}, {6, 6}},
			[]Region{{1, 1}},
		},
		{
			[]Region{{2, 2}, {3, 3}, {6, 6}},
			[]Region{{2, 2}},
		},
		{
			[]Region{{5, 5}},
			[]Region{{5, 5}},
		},
	}

	runFindTest(&tests, t, "single_selection")
}

func TestSelectAll(t *testing.T) {
	tests := []findTest{
		{
			[]Region{{1, 1}, {2, 2}, {3, 3}, {6, 6}},
			[]Region{{0, 36}},
		},
		{
			[]Region{{2, 2}, {3, 3}, {6, 6}},
			[]Region{{0, 36}},
		},
		{
			[]Region{{5, 5}},
			[]Region{{0, 36}},
		},
	}

	runFindTest(&tests, t, "select_all")
}
