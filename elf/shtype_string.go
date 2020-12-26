// Code generated by "stringer -type=SHType"; DO NOT EDIT.

package elf

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SHT_NULL-0]
	_ = x[SHT_PROGBITS-1]
	_ = x[SHT_SYMTAB-2]
	_ = x[SHT_STRTAB-3]
	_ = x[SHT_RELA-4]
	_ = x[SHT_HASH-5]
	_ = x[SHT_DYNAMIC-6]
	_ = x[SHT_NOTE-7]
	_ = x[SHT_NOBITS-8]
	_ = x[SHT_REL-9]
	_ = x[SHT_SHLIB-10]
	_ = x[SHT_DYNSYM-11]
	_ = x[SHT_NUM-12]
	_ = x[SHT_LOPROC-1879048192]
	_ = x[SHT_HIPROC-2147483647]
	_ = x[SHT_LOUSER-2147483648]
	_ = x[SHT_HIUSER-4294967295]
}

const (
	_SHType_name_0 = "SHT_NULLSHT_PROGBITSSHT_SYMTABSHT_STRTABSHT_RELASHT_HASHSHT_DYNAMICSHT_NOTESHT_NOBITSSHT_RELSHT_SHLIBSHT_DYNSYMSHT_NUM"
	_SHType_name_1 = "SHT_LOPROC"
	_SHType_name_2 = "SHT_HIPROCSHT_LOUSER"
	_SHType_name_3 = "SHT_HIUSER"
)

var (
	_SHType_index_0 = [...]uint8{0, 8, 20, 30, 40, 48, 56, 67, 75, 85, 92, 101, 111, 118}
	_SHType_index_2 = [...]uint8{0, 10, 20}
)

func (i SHType) String() string {
	switch {
	case i <= 12:
		return _SHType_name_0[_SHType_index_0[i]:_SHType_index_0[i+1]]
	case i == 1879048192:
		return _SHType_name_1
	case 2147483647 <= i && i <= 2147483648:
		i -= 2147483647
		return _SHType_name_2[_SHType_index_2[i]:_SHType_index_2[i+1]]
	case i == 4294967295:
		return _SHType_name_3
	default:
		return "SHType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
