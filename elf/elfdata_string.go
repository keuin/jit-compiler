// Code generated by "stringer -type=ELFData"; DO NOT EDIT.

package elf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ELFDATANONE-0]
	_ = x[ELFDATA2LSB-1]
	_ = x[ELFDATA2MSB-2]
}

const _ELFData_name = "ELFDATANONEELFDATA2LSBELFDATA2MSB"

var _ELFData_index = [...]uint8{0, 11, 22, 33}

func (i ELFData) String() string {
	if i >= ELFData(len(_ELFData_index)-1) {
		return "ELFData(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ELFData_name[_ELFData_index[i]:_ELFData_index[i+1]]
}