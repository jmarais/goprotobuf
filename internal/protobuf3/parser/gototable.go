// Code generated by gocc; DO NOT EDIT.

package parser

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
)

const numNTSymbols = 52

type (
	gotoTable [numStates]gotoRow
	gotoRow   [numNTSymbols]int
)

var gotoTab = gotoTable{}

func init() {
	tab := [][]int{}
	data := []byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0xec, 0x9a, 0xb1, 0x6b, 0x14, 0x4f,
		0x14, 0xc7, 0xf7, 0xfb, 0x08, 0x3f, 0x7e, 0x84, 0x18, 0x82, 0xa4, 0x48, 0x21, 0x12, 0x82, 0x88,
		0x84, 0x10, 0x44, 0x42, 0x48, 0x21, 0x22, 0x29, 0x44, 0x14, 0x52, 0x58, 0x04, 0x8b, 0x60, 0x61,
		0x25, 0x62, 0x65, 0x21, 0x16, 0x16, 0x62, 0x21, 0xf1, 0x4c, 0x21, 0x41, 0x82, 0x65, 0x10, 0xb1,
		0x08, 0x22, 0x16, 0x12, 0x0e, 0x91, 0x20, 0x12, 0xc4, 0x22, 0xa4, 0x90, 0x20, 0x62, 0x29, 0xe2,
		0xdf, 0x20, 0x22, 0xdf, 0x91, 0xd3, 0xdb, 0xec, 0x26, 0xd9, 0xec, 0x66, 0x5f, 0x76, 0xe7, 0xf6,
		0x2e, 0xef, 0x15, 0x77, 0x73, 0xd9, 0xf7, 0xb9, 0xcf, 0x70, 0xbb, 0xb3, 0x33, 0xf3, 0x36, 0x87,
		0xdc, 0x43, 0x81, 0xb8, 0xb9, 0x00, 0xae, 0x16, 0x04, 0x3d, 0xee, 0x41, 0xe3, 0x53, 0x2d, 0x40,
		0x57, 0x10, 0xf0, 0x62, 0xdd, 0xcd, 0x05, 0xc4, 0xb5, 0x31, 0x08, 0xd0, 0x85, 0x7c, 0x31, 0x96,
		0x33, 0x3f, 0x84, 0xfe, 0xef, 0x46, 0x0f, 0x7a, 0xfb, 0x80, 0xc3, 0x5b, 0x0f, 0x1d, 0x6d, 0xbc,
		0x1c, 0x01, 0x06, 0x0b, 0x33, 0x61, 0xd8, 0x9b, 0xc9, 0xa0, 0x08, 0x1a, 0x29, 0xc7, 0x74, 0x2a,
		0x11, 0x1a, 0x2f, 0xc1, 0x54, 0x79, 0xe8, 0x8c, 0x02, 0x9a, 0x54, 0x99, 0x1a, 0x71, 0x4e, 0x03,
		0xa9, 0x4c, 0x06, 0x15, 0x0a, 0x5d, 0x2a, 0xde, 0x34, 0xad, 0x81, 0xf6, 0x6a, 0x9a, 0xd1, 0x40,
		0x2a, 0x93, 0x41, 0xad, 0x85, 0xae, 0x47, 0xcd, 0x1b, 0x15, 0xec, 0x5e, 0x71, 0xd0, 0x9d, 0x84,
		0x83, 0xee, 0x7e, 0x29, 0x2a, 0x37, 0xdb, 0x78, 0x59, 0x80, 0x7b, 0x03, 0xb8, 0x15, 0xc0, 0xad,
		0xc2, 0x7d, 0x6c, 0xfc, 0xe9, 0xf1, 0xdf, 0xa3, 0xf3, 0x6e, 0xc3, 0x7d, 0x0e, 0x53, 0x3f, 0xb9,
		0xf5, 0x7f, 0x8d, 0x65, 0xf7, 0xc5, 0xd5, 0x63, 0xdf, 0xf1, 0x75, 0x77, 0x91, 0xfb, 0xe6, 0xbe,
		0xbb, 0x1f, 0x45, 0xf6, 0xd8, 0x20, 0x83, 0xb6, 0x43, 0xee, 0xa7, 0x37, 0x95, 0x37, 0x88, 0x08,
		0xbc, 0xb9, 0x3a, 0x11, 0x22, 0xfe, 0xf3, 0xe6, 0xca, 0x0a, 0xa2, 0x87, 0xe8, 0x0e, 0xdb, 0xbd,
		0x55, 0xfd, 0x15, 0x19, 0xed, 0x79, 0xab, 0xd9, 0x41, 0x83, 0x0c, 0x32, 0x28, 0x07, 0x44, 0xf4,
		0xa7, 0xe5, 0xee, 0xba, 0xa4, 0x53, 0x0b, 0xb3, 0x82, 0x18, 0xd0, 0x60, 0x2a, 0x57, 0x75, 0x20,
		0x62, 0x88, 0x38, 0x4e, 0x1c, 0x4b, 0x38, 0x32, 0xa8, 0x30, 0x12, 0x27, 0x32, 0x9d, 0xa9, 0x91,
		0x67, 0xea, 0x1a, 0xd6, 0x60, 0x49, 0xae, 0xac, 0xbd, 0x06, 0x88, 0xd1, 0x30, 0xd5, 0xef, 0x85,
		0xc9, 0x78, 0x9d, 0xb0, 0x35, 0x8b, 0x95, 0x1c, 0x5f, 0xb0, 0xcf, 0xb3, 0x90, 0xb2, 0x5b, 0x03,
		0x31, 0xd1, 0x3e, 0xbb, 0x35, 0xe2, 0xb4, 0x37, 0x97, 0x41, 0x06, 0x19, 0xd4, 0x32, 0x88, 0xb8,
		0x40, 0x9c, 0xf7, 0x65, 0x53, 0x42, 0xc4, 0x94, 0x37, 0x57, 0x06, 0xc4, 0x78, 0xb1, 0xb4, 0x64,
		0x97, 0x41, 0x5b, 0x82, 0xb8, 0xed, 0xcd, 0x65, 0x90, 0x41, 0x29, 0x11, 0xbb, 0x07, 0x10, 0xf7,
		0x3c, 0xba, 0x66, 0x2b, 0xf8, 0x63, 0xec, 0x0e, 0x11, 0x35, 0x5f, 0x2e, 0x62, 0xbe, 0xf9, 0xbe,
		0x40, 0x3c, 0x49, 0xcc, 0xf7, 0xbf, 0x27, 0x37, 0xc8, 0xa0, 0xce, 0x82, 0x88, 0x45, 0x6f, 0xae,
		0x44, 0xfd, 0xd3, 0x1d, 0x18, 0xf1, 0x2c, 0x31, 0x73, 0xa9, 0x00, 0xdf, 0x4e, 0x28, 0xab, 0xd4,
		0x42, 0xbc, 0xcc, 0x7e, 0xae, 0xab, 0x32, 0x67, 0xc5, 0x96, 0x79, 0xe2, 0x55, 0x9e, 0x1d, 0x40,
		0x3d, 0x6c, 0x2c, 0x13, 0x6f, 0x5b, 0x55, 0xa6, 0x79, 0x97, 0xd8, 0xb5, 0xf7, 0xc4, 0x6a, 0x19,
		0x3e, 0xbf, 0x10, 0xb1, 0xe6, 0xcd, 0x95, 0x1e, 0xc4, 0x7a, 0x71, 0xae, 0x70, 0xce, 0x05, 0xb1,
		0x91, 0x98, 0x5d, 0xfc, 0x00, 0x20, 0xd2, 0x8a, 0x6c, 0xa0, 0x04, 0x04, 0x9b, 0x99, 0xbf, 0x28,
		0x42, 0xfc, 0x6e, 0xaf, 0x4b, 0xc5, 0x20, 0x25, 0x54, 0xdd, 0x1a, 0xb8, 0x0a, 0xa2, 0xf4, 0x6e,
		0xb6, 0xfa, 0xaa, 0x79, 0x16, 0xa2, 0x89, 0x43, 0xfa, 0x73, 0x4e, 0x1c, 0x94, 0x01, 0x8d, 0x11,
		0x06, 0xa5, 0x40, 0x94, 0xc1, 0xe6, 0xfb, 0x10, 0x65, 0xe7, 0x03, 0xb1, 0xc2, 0x7d, 0x06, 0xed,
		0x31, 0x28, 0xa3, 0x94, 0x91, 0xb0, 0x7d, 0xb2, 0x9a, 0x9d, 0x34, 0xc8, 0x20, 0x83, 0xda, 0x11,
		0xa2, 0x8c, 0x7b, 0x73, 0x19, 0xd4, 0x0c, 0xca, 0x84, 0x37, 0x57, 0xb5, 0x21, 0xca, 0xd9, 0xf8,
		0x87, 0xc9, 0x52, 0x2a, 0x01, 0xa0, 0x4c, 0x65, 0x80, 0x2a, 0x5b, 0xc5, 0xa1, 0x68, 0x49, 0x37,
		0xad, 0x59, 0xd2, 0x51, 0x2e, 0xe7, 0x35, 0xe6, 0x88, 0xd2, 0x47, 0xd8, 0x15, 0xca, 0x8c, 0xd2,
		0x46, 0xb9, 0xaa, 0xc1, 0x54, 0x2e, 0x83, 0x0c, 0x6a, 0x17, 0x68, 0xdb, 0xcd, 0xfa, 0x66, 0x41,
		0xf7, 0x27, 0x50, 0x6e, 0xe5, 0x02, 0x55, 0x36, 0x83, 0xca, 0x84, 0x62, 0xe5, 0x9c, 0xbb, 0x2d,
		0x7b, 0x0e, 0x60, 0x50, 0x07, 0x43, 0xd1, 0x6a, 0xa6, 0xa6, 0x5b, 0xcd, 0xcc, 0xe5, 0x35, 0xe6,
		0x08, 0xd5, 0x40, 0x79, 0x94, 0xb7, 0xee, 0x19, 0xa1, 0xf3, 0x07, 0x7d, 0x8c, 0x51, 0x16, 0xbd,
		0xb9, 0xb4, 0x50, 0xec, 0x7c, 0x3d, 0x6f, 0xdd, 0xb3, 0xd1, 0xcd, 0x61, 0xb3, 0xe4, 0xa3, 0xae,
		0x4b, 0x79, 0x91, 0x9e, 0x49, 0x79, 0x5d, 0x9c, 0x2d, 0x4b, 0xb5, 0xac, 0xc3, 0xea, 0x1a, 0x0c,
		0x07, 0x08, 0xa2, 0xac, 0xf8, 0x72, 0x51, 0x12, 0x77, 0xf4, 0xa5, 0xb8, 0x4a, 0x80, 0x28, 0x1f,
		0xbc, 0xb9, 0x0c, 0xda, 0x1f, 0x14, 0xff, 0x1f, 0x18, 0x59, 0xf3, 0xd7, 0xc1, 0x3f, 0x01, 0x00,
		0x00, 0xff, 0xff, 0x12, 0x1c, 0xae, 0xb0, 0xd6, 0x4b, 0x00, 0x00,
	}
	buf, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&tab); err != nil {
		panic(err)
	}
	for i := 0; i < numStates; i++ {
		for j := 0; j < numNTSymbols; j++ {
			gotoTab[i][j] = tab[i][j]
		}
	}
}
