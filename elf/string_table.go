package elf

import (
	"bytes"
	"errors"
	"io"
)

type StringTable struct {
	data []byte
}

func NewStringTable(data []byte) *StringTable {
	return &StringTable{data}
}

func ParseStringTable(header *ELFHeader, sectionHeader *SectionHeader, r *bytes.Reader) (*StringTable, error) {
	_, err := r.Seek(int64(sectionHeader.Offset), io.SeekStart)
	if err != nil {
		return nil, err
	}
	data := make([]byte, sectionHeader.Size)
	n, err := r.Read(data)
	if err != nil {
		return nil, err
	}
	if n != int(sectionHeader.Size) {
		return nil, errors.New("Mismatching string table size")
	}
	return NewStringTable(data), nil
}

func (s *StringTable) GetString(ix int) (string, error) {
	result := []byte{}
	if ix >= len(s.data) {
		return "", errors.New("index out of bounds")
	}
	for ix < len(s.data) && s.data[ix] != 0 {
		result = append(result, s.data[ix])
		ix++
	}
	return string(result), nil
}
