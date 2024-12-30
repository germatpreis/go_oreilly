// Code generated by "stringer -type=Direction"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[North-1]
	_ = x[South-2]
	_ = x[East-3]
	_ = x[West-4]
}

const _Direction_name = "NorthSouthEastWest"

var _Direction_index = [...]uint8{0, 5, 10, 14, 18}

func (i Direction) String() string {
	i -= 1
	if i < 0 || i >= Direction(len(_Direction_index)-1) {
		return "Direction(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Direction_name[_Direction_index[i]:_Direction_index[i+1]]
}