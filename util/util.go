// Copyright 2013 The Chihaya Authors. All rights reserved.
// Use of this source code is governed by the BSD 2-Clause license,
// which can be found in the LICENSE file.
// Chihaya.
package util

// MinInt returns the smaller of the two integers provided.
func MinInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the larger of the two integers provided.
func MaxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// Btoa converts a boolean value into the string form "1" or "0".
func Btoa(a bool) string {
	if a {
		return "1"
	}
	return "0"
}
