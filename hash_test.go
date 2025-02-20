/*
MIT License

Copyright (c) 2021 Prysmatic Labs

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package gohashtree_test

import (
	"reflect"
	"testing"

	"github.com/umbracle/gohashtree"
)

var _test_32_block = [][32]byte{
	{0x7a, 0xee, 0xd5, 0xc9, 0x66, 0x17, 0x59, 0x7f, 0x89, 0xd6, 0xd9, 0xe8, 0xa8, 0xa7, 0x01, 0x47, 0x60, 0xc6, 0x88, 0xfd, 0x2a, 0x7a, 0xf6, 0x1d, 0x10, 0x20, 0x62, 0x7e, 0x7c, 0xd0, 0x1a, 0x0b},
	{0xd4, 0x1f, 0xa7, 0x89, 0x8c, 0xf9, 0x05, 0xfc, 0x1e, 0xb0, 0x04, 0xd7, 0xaa, 0x56, 0x35, 0xec, 0x36, 0xf5, 0x0d, 0x41, 0x75, 0x64, 0x34, 0x71, 0xf0, 0x3b, 0x5b, 0xb2, 0xcc, 0xfa, 0x8c, 0xca},
	{0xf8, 0xd9, 0x9e, 0xa7, 0x9c, 0xa1, 0xe0, 0x3a, 0x19, 0x4f, 0xd3, 0x2d, 0xbd, 0x40, 0x3a, 0xa3, 0x28, 0xe8, 0xa4, 0x27, 0x58, 0x44, 0x12, 0xf7, 0x69, 0x01, 0x66, 0xfa, 0xf1, 0x97, 0x30, 0xfe},
	{0x99, 0x7c, 0x24, 0x0e, 0xed, 0x31, 0x0a, 0xda, 0x12, 0x16, 0x0e, 0x06, 0x44, 0xb8, 0x3f, 0xa2, 0x40, 0x52, 0xbc, 0x2d, 0xaf, 0x97, 0x00, 0x01, 0x5d, 0xbb, 0x0d, 0x06, 0x66, 0xb1, 0x59, 0xf2},
	{0x99, 0x43, 0x52, 0x77, 0x28, 0x39, 0x6b, 0xeb, 0x03, 0x51, 0xc4, 0x5f, 0x7d, 0xd3, 0xe1, 0x41, 0x17, 0x66, 0x7b, 0x0e, 0xc9, 0x51, 0x01, 0xa7, 0x39, 0xf3, 0xc8, 0x63, 0x95, 0xa5, 0x92, 0x6b},
	{0xce, 0x6e, 0xab, 0xd2, 0xe8, 0xad, 0x90, 0xad, 0xbe, 0xe5, 0x94, 0x96, 0xa9, 0x98, 0xe7, 0x83, 0x07, 0xa4, 0x0f, 0x8e, 0xe5, 0xb3, 0x5a, 0x05, 0xcd, 0xfd, 0xae, 0x9c, 0x07, 0xad, 0x26, 0xaa},
	{0xf5, 0xee, 0x66, 0x87, 0x00, 0xed, 0xeb, 0x8b, 0xc2, 0x7d, 0x97, 0x52, 0x2d, 0xfc, 0x0a, 0x2a, 0x32, 0x0e, 0x92, 0xd2, 0x91, 0xd1, 0x69, 0x29, 0x9d, 0xb1, 0x3a, 0x65, 0x9f, 0x8e, 0x7e, 0x2a},
	{0x88, 0x4a, 0xc8, 0x81, 0xdb, 0xa6, 0x79, 0x36, 0x54, 0xe9, 0x15, 0x5c, 0xff, 0x06, 0x35, 0x8b, 0x6e, 0x0d, 0xaa, 0x3e, 0x7a, 0x82, 0x7c, 0x4a, 0xfe, 0x8a, 0x91, 0xb4, 0x34, 0xed, 0xe3, 0x17},
	{0xe7, 0x92, 0xa4, 0x91, 0xdc, 0x1d, 0x83, 0xc8, 0x72, 0x5a, 0xd1, 0x27, 0x17, 0x78, 0x2b, 0xc7, 0x67, 0xe9, 0x56, 0xf2, 0xb4, 0x37, 0x51, 0xa1, 0x6b, 0x23, 0x8c, 0xc9, 0x03, 0x3d, 0x90, 0x1e},
	{0xc4, 0x1f, 0xcc, 0x5e, 0xcb, 0x5e, 0x7d, 0x02, 0x12, 0x3f, 0x15, 0x9f, 0x35, 0xf4, 0x49, 0x55, 0xba, 0xc6, 0x47, 0xd2, 0x85, 0x85, 0x61, 0x69, 0xa5, 0x60, 0x7a, 0x32, 0x7f, 0x8e, 0x09, 0x5f},
	{0x60, 0xb6, 0xab, 0xb5, 0x6b, 0x4d, 0xce, 0x6f, 0x1d, 0x77, 0x2e, 0x9b, 0x0d, 0x60, 0x76, 0xe3, 0xcb, 0x79, 0xbc, 0x40, 0x2d, 0x16, 0xf6, 0xa3, 0x06, 0x12, 0x36, 0x71, 0xda, 0xfd, 0x28, 0x89},
	{0x67, 0xdd, 0x7f, 0x26, 0x6d, 0x2e, 0xf3, 0xef, 0x13, 0xb6, 0x09, 0x73, 0x82, 0xbc, 0x73, 0x25, 0x83, 0xc0, 0x34, 0x90, 0xe8, 0xad, 0xf0, 0x17, 0x8d, 0xed, 0xad, 0x29, 0xf7, 0x78, 0x9c, 0x28},
	{0x00, 0xb0, 0xd5, 0xd0, 0x8e, 0x9b, 0xe5, 0xf0, 0x46, 0x8e, 0x60, 0x25, 0x95, 0xe5, 0x3a, 0x46, 0xb1, 0x07, 0x74, 0x97, 0xed, 0x0a, 0x2f, 0x9a, 0x3f, 0xf3, 0x94, 0x2f, 0xb3, 0x12, 0xa1, 0x91},
	{0x8d, 0x36, 0x16, 0xc6, 0x00, 0x88, 0xd6, 0x69, 0xb4, 0x5a, 0x71, 0x18, 0x41, 0xe5, 0x4d, 0xb2, 0xd9, 0x00, 0x7a, 0x17, 0x63, 0x6a, 0x9b, 0x2e, 0x22, 0x12, 0x5b, 0xa3, 0x74, 0x7c, 0x95, 0xc9},
	{0x4e, 0xfc, 0x5c, 0x18, 0xd1, 0x8a, 0x5b, 0x57, 0x7c, 0x86, 0x3e, 0xe2, 0x75, 0x91, 0xf2, 0xb3, 0x5f, 0xd0, 0x92, 0xbc, 0x77, 0xbe, 0x1b, 0xef, 0x1a, 0x7c, 0xe2, 0xd8, 0x8d, 0x7b, 0xef, 0xf7},
	{0xb7, 0x80, 0xc2, 0x31, 0xe6, 0x75, 0x0c, 0xad, 0x0f, 0xe8, 0xed, 0x59, 0x34, 0xdb, 0xfb, 0x41, 0xd4, 0x38, 0x73, 0x7a, 0x47, 0x01, 0xb8, 0xea, 0xea, 0x2e, 0x01, 0x8e, 0x4f, 0x09, 0x64, 0x82},
	{0x99, 0x43, 0x52, 0x77, 0x28, 0x39, 0x6b, 0xeb, 0x03, 0x51, 0xc4, 0x5f, 0x7d, 0xd3, 0xe1, 0x41, 0x17, 0x66, 0x7b, 0x0e, 0xc9, 0x51, 0x01, 0xa7, 0x39, 0xf3, 0xc8, 0x63, 0x95, 0xa5, 0x92, 0x6b},
	{0xce, 0x6e, 0xab, 0xd2, 0xe8, 0xad, 0x90, 0xad, 0xbe, 0xe5, 0x94, 0x96, 0xa9, 0x98, 0xe7, 0x83, 0x07, 0xa4, 0x0f, 0x8e, 0xe5, 0xb3, 0x5a, 0x05, 0xcd, 0xfd, 0xae, 0x9c, 0x07, 0xad, 0x26, 0xaa},
	{0xf5, 0xee, 0x66, 0x87, 0x00, 0xed, 0xeb, 0x8b, 0xc2, 0x7d, 0x97, 0x52, 0x2d, 0xfc, 0x0a, 0x2a, 0x32, 0x0e, 0x92, 0xd2, 0x91, 0xd1, 0x69, 0x29, 0x9d, 0xb1, 0x3a, 0x65, 0x9f, 0x8e, 0x7e, 0x2a},
	{0x88, 0x4a, 0xc8, 0x81, 0xdb, 0xa6, 0x79, 0x36, 0x54, 0xe9, 0x15, 0x5c, 0xff, 0x06, 0x35, 0x8b, 0x6e, 0x0d, 0xaa, 0x3e, 0x7a, 0x82, 0x7c, 0x4a, 0xfe, 0x8a, 0x91, 0xb4, 0x34, 0xed, 0xe3, 0x17},
	{0xe7, 0x92, 0xa4, 0x91, 0xdc, 0x1d, 0x83, 0xc8, 0x72, 0x5a, 0xd1, 0x27, 0x17, 0x78, 0x2b, 0xc7, 0x67, 0xe9, 0x56, 0xf2, 0xb4, 0x37, 0x51, 0xa1, 0x6b, 0x23, 0x8c, 0xc9, 0x03, 0x3d, 0x90, 0x1e},
	{0xc4, 0x1f, 0xcc, 0x5e, 0xcb, 0x5e, 0x7d, 0x02, 0x12, 0x3f, 0x15, 0x9f, 0x35, 0xf4, 0x49, 0x55, 0xba, 0xc6, 0x47, 0xd2, 0x85, 0x85, 0x61, 0x69, 0xa5, 0x60, 0x7a, 0x32, 0x7f, 0x8e, 0x09, 0x5f},
	{0x60, 0xb6, 0xab, 0xb5, 0x6b, 0x4d, 0xce, 0x6f, 0x1d, 0x77, 0x2e, 0x9b, 0x0d, 0x60, 0x76, 0xe3, 0xcb, 0x79, 0xbc, 0x40, 0x2d, 0x16, 0xf6, 0xa3, 0x06, 0x12, 0x36, 0x71, 0xda, 0xfd, 0x28, 0x89},
	{0x67, 0xdd, 0x7f, 0x26, 0x6d, 0x2e, 0xf3, 0xef, 0x13, 0xb6, 0x09, 0x73, 0x82, 0xbc, 0x73, 0x25, 0x83, 0xc0, 0x34, 0x90, 0xe8, 0xad, 0xf0, 0x17, 0x8d, 0xed, 0xad, 0x29, 0xf7, 0x78, 0x9c, 0x28},
	{0x00, 0xb0, 0xd5, 0xd0, 0x8e, 0x9b, 0xe5, 0xf0, 0x46, 0x8e, 0x60, 0x25, 0x95, 0xe5, 0x3a, 0x46, 0xb1, 0x07, 0x74, 0x97, 0xed, 0x0a, 0x2f, 0x9a, 0x3f, 0xf3, 0x94, 0x2f, 0xb3, 0x12, 0xa1, 0x91},
	{0x8d, 0x36, 0x16, 0xc6, 0x00, 0x88, 0xd6, 0x69, 0xb4, 0x5a, 0x71, 0x18, 0x41, 0xe5, 0x4d, 0xb2, 0xd9, 0x00, 0x7a, 0x17, 0x63, 0x6a, 0x9b, 0x2e, 0x22, 0x12, 0x5b, 0xa3, 0x74, 0x7c, 0x95, 0xc9},
	{0x4e, 0xfc, 0x5c, 0x18, 0xd1, 0x8a, 0x5b, 0x57, 0x7c, 0x86, 0x3e, 0xe2, 0x75, 0x91, 0xf2, 0xb3, 0x5f, 0xd0, 0x92, 0xbc, 0x77, 0xbe, 0x1b, 0xef, 0x1a, 0x7c, 0xe2, 0xd8, 0x8d, 0x7b, 0xef, 0xf7},
	{0xcd, 0x78, 0x15, 0x64, 0x2c, 0x78, 0x57, 0x74, 0x2b, 0xb7, 0xdb, 0x74, 0xe2, 0xab, 0x82, 0xbb, 0x61, 0x32, 0x3e, 0xe4, 0xb1, 0x00, 0xde, 0xb2, 0x35, 0x1e, 0x3e, 0x1c, 0x91, 0x9d, 0x87, 0xde},
	{0x17, 0xcc, 0x52, 0x5c, 0x60, 0x9e, 0xd8, 0xd4, 0xf4, 0x56, 0x28, 0x16, 0xde, 0xde, 0x73, 0xfe, 0xd9, 0x92, 0xb7, 0x99, 0x15, 0x24, 0x1b, 0x40, 0xb0, 0xda, 0x9a, 0xf8, 0x24, 0x38, 0x13, 0xbd},
	{0xd0, 0x45, 0x9b, 0xe3, 0x9a, 0xae, 0x78, 0x41, 0xcd, 0x12, 0x9a, 0x6b, 0x91, 0x58, 0x29, 0x75, 0xae, 0x21, 0xd3, 0xf2, 0x5e, 0x98, 0xab, 0x09, 0xb0, 0xaa, 0x62, 0x96, 0x35, 0x64, 0x18, 0x48},
	{0xd2, 0x5b, 0x10, 0xf1, 0x35, 0xaa, 0x04, 0x49, 0x4e, 0x51, 0x30, 0x0d, 0xb6, 0xbf, 0xa0, 0x9b, 0xa0, 0xf5, 0x66, 0x5f, 0x28, 0xc7, 0x8d, 0xa8, 0x3e, 0x0f, 0xe4, 0xa7, 0xc9, 0xd4, 0x0f, 0x7d},
	{0xb7, 0x80, 0xc2, 0x31, 0xe6, 0x75, 0x0c, 0xad, 0x0f, 0xe8, 0xed, 0x59, 0x34, 0xdb, 0xfb, 0x41, 0xd4, 0x38, 0x73, 0x7a, 0x47, 0x01, 0xb8, 0xea, 0xea, 0x2e, 0x01, 0x8e, 0x4f, 0x09, 0x64, 0x82},
	{0xe4, 0x8b, 0x12, 0xd3, 0xd0, 0x78, 0xb5, 0x5f, 0x3e, 0x9d, 0x94, 0x7f, 0x93, 0x84, 0x77, 0x77, 0xdb, 0x78, 0x41, 0xe8, 0x91, 0xfb, 0x6d, 0x0d, 0xef, 0x00, 0x30, 0x8e, 0x0a, 0xe4, 0x7b, 0xec},
	{0xe7, 0xb2, 0x76, 0xe7, 0x6c, 0xba, 0x8f, 0x8c, 0x0b, 0xf2, 0xa3, 0xad, 0xc2, 0x2d, 0x92, 0xb4, 0xd5, 0xf2, 0x83, 0x42, 0x65, 0x02, 0xd6, 0x67, 0x9a, 0x78, 0x6a, 0xc1, 0xca, 0x91, 0x87, 0x7c},
	{0x16, 0x99, 0x13, 0xf8, 0xa9, 0x20, 0x62, 0x2e, 0xc1, 0x84, 0xc0, 0x25, 0xdc, 0x35, 0x1f, 0xe6, 0x32, 0x49, 0x37, 0x79, 0x78, 0xfb, 0xf5, 0xf7, 0x34, 0xf4, 0xa5, 0x49, 0x9f, 0xc8, 0xfa, 0x8e},
	{0x28, 0x9b, 0x27, 0xae, 0x21, 0x12, 0x14, 0x57, 0x56, 0xf6, 0x9d, 0x7f, 0x0d, 0x28, 0x03, 0xbd, 0x05, 0xd0, 0x11, 0x9e, 0xf1, 0x98, 0x8e, 0x1c, 0xbe, 0xc1, 0x83, 0xdb, 0x1a, 0x65, 0x08, 0x0d},
	{0xef, 0x42, 0x3a, 0x0b, 0x2f, 0xea, 0xdf, 0xfe, 0xeb, 0xd9, 0x72, 0x9a, 0xcf, 0x5a, 0xac, 0x19, 0x09, 0x75, 0x25, 0x64, 0x61, 0x19, 0xf5, 0xcd, 0xdb, 0x9d, 0xcf, 0x4a, 0xa9, 0xf5, 0x48, 0x2c},
	{0x47, 0x69, 0xaa, 0x80, 0x3f, 0xd3, 0x02, 0x67, 0xe9, 0x8b, 0x82, 0xa8, 0x02, 0xe8, 0xcf, 0x60, 0x66, 0xaa, 0xcf, 0x05, 0x0a, 0x85, 0xeb, 0x3d, 0x87, 0x21, 0xcc, 0xe2, 0xdd, 0x6c, 0x42, 0x54},
	{0xd8, 0xb4, 0x39, 0x4f, 0x78, 0xce, 0xd8, 0xad, 0x57, 0xbe, 0xda, 0x18, 0x8f, 0x4a, 0x9b, 0x41, 0xfe, 0x58, 0x9d, 0xa1, 0xd4, 0x71, 0x6e, 0x2f, 0x04, 0xaf, 0x37, 0xa0, 0x29, 0x60, 0x6f, 0x9d},
	{0x84, 0x4a, 0x39, 0x0a, 0x5e, 0x24, 0x81, 0x2e, 0x63, 0xc9, 0xb6, 0xde, 0xc3, 0xf1, 0x82, 0x7b, 0x82, 0x14, 0x07, 0xde, 0x46, 0x03, 0x25, 0x27, 0x4d, 0x09, 0x6b, 0x7e, 0xb9, 0x82, 0x98, 0x41},
	{0x68, 0xf8, 0x98, 0x04, 0xb2, 0x61, 0x78, 0xbf, 0x8a, 0x69, 0x4d, 0xc7, 0x83, 0x4a, 0xe7, 0x77, 0xf7, 0x4b, 0x00, 0x28, 0x34, 0xe6, 0x36, 0xca, 0xa2, 0x58, 0x37, 0x61, 0x60, 0x95, 0x0d, 0xa6},
	{0x20, 0x00, 0x7e, 0x29, 0xa8, 0x6e, 0xca, 0xb8, 0x1b, 0xbc, 0x94, 0x29, 0x2b, 0x18, 0xaa, 0x56, 0x0f, 0x4c, 0x38, 0x1a, 0x7a, 0x16, 0xe8, 0xbb, 0x51, 0xb7, 0xb3, 0xe3, 0x22, 0x8e, 0x9c, 0x05},
	{0xa8, 0x0f, 0x08, 0x4d, 0xf1, 0xd1, 0xd8, 0x2c, 0xac, 0xe8, 0x73, 0x43, 0xcc, 0x73, 0x6b, 0x03, 0x40, 0x21, 0x85, 0x9b, 0x9d, 0x63, 0xa8, 0x44, 0x6a, 0x6c, 0x23, 0xe3, 0x4e, 0x76, 0xb1, 0x51},
	{0x90, 0x61, 0x31, 0xfe, 0xf7, 0x4a, 0x8f, 0x06, 0x9e, 0x75, 0x6a, 0x5a, 0x66, 0xdd, 0xa2, 0xe4, 0x9b, 0x8f, 0x98, 0xbb, 0x18, 0x9a, 0x96, 0x84, 0xfa, 0xe4, 0x3c, 0xd2, 0x2c, 0x96, 0x61, 0xd8},
	{0x96, 0xb4, 0x84, 0xa8, 0x8b, 0x6f, 0xeb, 0xc5, 0x3e, 0xa3, 0x48, 0xd5, 0x00, 0x95, 0x47, 0xda, 0xc1, 0x2d, 0x95, 0x68, 0x49, 0x29, 0x15, 0xb9, 0x36, 0x59, 0x4c, 0x0b, 0x77, 0xdc, 0x01, 0x06},
	{0x58, 0x37, 0xa7, 0x03, 0x40, 0x70, 0x91, 0xee, 0x29, 0x75, 0x10, 0xd4, 0xec, 0x01, 0x87, 0x5f, 0x2e, 0xb5, 0x56, 0xc6, 0x2d, 0xe9, 0x2b, 0xb4, 0xab, 0x95, 0x82, 0x1f, 0x11, 0xf2, 0xb8, 0xc9},
	{0x81, 0xbf, 0xb0, 0x58, 0xcc, 0xdd, 0x0e, 0xf1, 0x9c, 0x17, 0x6b, 0xa0, 0xe6, 0x42, 0x8c, 0x1a, 0x3c, 0x9c, 0x20, 0x18, 0x0b, 0x52, 0x66, 0x5a, 0xc1, 0xe5, 0xc5, 0x66, 0x35, 0xe5, 0x26, 0x4f},
	{0xca, 0x73, 0xe0, 0x95, 0x2c, 0xc7, 0xa9, 0x22, 0x58, 0x68, 0x49, 0xb3, 0x68, 0xdc, 0x34, 0xe1, 0x3b, 0x17, 0x67, 0xaa, 0x82, 0xa1, 0xb6, 0xbd, 0x69, 0x9b, 0xf6, 0x00, 0x71, 0x51, 0x08, 0xca},
	{0xce, 0x06, 0x68, 0x95, 0x13, 0x37, 0x8b, 0x32, 0xc9, 0x62, 0x38, 0xc9, 0x78, 0x90, 0x89, 0x0e, 0x3a, 0x5d, 0x85, 0x50, 0x1c, 0x4c, 0xd6, 0x80, 0xcc, 0x5f, 0x63, 0xf0, 0xc9, 0xfe, 0x7a, 0xb5},
	{0x79, 0x78, 0x8d, 0x38, 0x13, 0xdf, 0xb7, 0x37, 0x18, 0x78, 0xbd, 0x2f, 0x3e, 0xc7, 0x2c, 0x46, 0xd2, 0x74, 0x01, 0xe9, 0xa1, 0x3f, 0xfe, 0x46, 0x11, 0xb0, 0x85, 0x2f, 0x6d, 0x4b, 0x4b, 0x8e},
	{0x11, 0xce, 0x55, 0xe4, 0xba, 0xf7, 0x11, 0xcd, 0xe8, 0xa8, 0x04, 0x33, 0xbd, 0x19, 0xe8, 0xbe, 0xa1, 0x00, 0xd3, 0x28, 0xca, 0x78, 0x56, 0x6d, 0xde, 0xe5, 0x71, 0x13, 0xc2, 0xbd, 0xd8, 0xc2},
	{0x04, 0x64, 0xdb, 0xdb, 0x8b, 0x4f, 0x73, 0x0e, 0x0a, 0x9e, 0xfe, 0xd0, 0x5d, 0x92, 0x3e, 0xf8, 0xf4, 0x8b, 0xef, 0xb6, 0x6f, 0x42, 0xc9, 0xea, 0x73, 0xfb, 0xb6, 0x8e, 0x37, 0x74, 0xae, 0x39},
	{0x91, 0x1e, 0x40, 0x74, 0x23, 0xa7, 0xa8, 0x00, 0xfc, 0xa1, 0x16, 0xed, 0xcf, 0xff, 0xce, 0xea, 0x3f, 0x31, 0x54, 0xad, 0x19, 0x98, 0xcb, 0x5d, 0xfd, 0x82, 0xe2, 0x48, 0xbf, 0xc3, 0x74, 0x71},
	{0x5f, 0x45, 0x5f, 0xba, 0x82, 0x5d, 0xc4, 0x20, 0x12, 0x67, 0x65, 0x0d, 0x8b, 0x14, 0x45, 0x20, 0xd3, 0xbc, 0xb4, 0x23, 0x26, 0x98, 0xfc, 0x05, 0x8f, 0xa5, 0x99, 0xe2, 0x78, 0x74, 0x72, 0x71},
	{0xda, 0xa5, 0x2a, 0xc1, 0x13, 0xa4, 0x3b, 0xeb, 0x41, 0x51, 0x1b, 0x96, 0xa3, 0xa0, 0x5b, 0xd8, 0xed, 0x5e, 0x69, 0x67, 0xfb, 0xc5, 0x27, 0x66, 0x56, 0x8a, 0xb2, 0x1e, 0x93, 0xbf, 0xb0, 0x36},
	{0x54, 0xb8, 0x17, 0xb6, 0xd2, 0x26, 0x22, 0x93, 0xdc, 0xb5, 0xd5, 0x32, 0x1b, 0x76, 0x3c, 0xfa, 0x24, 0x04, 0xcb, 0xa0, 0x1b, 0xcb, 0xa3, 0x12, 0x20, 0x60, 0x3b, 0x59, 0xe5, 0xdf, 0xf7, 0xbf},
	{0x41, 0x42, 0x6c, 0xbf, 0xfa, 0x23, 0xcc, 0xee, 0x3e, 0xf6, 0xf3, 0xbf, 0xa1, 0x39, 0x9b, 0x6e, 0x7f, 0xfb, 0x2c, 0x7f, 0x4e, 0xf5, 0x35, 0x78, 0xb5, 0x5e, 0x77, 0x02, 0x40, 0x2a, 0xbc, 0x77},
	{0x9b, 0xc5, 0x2f, 0xb6, 0xa1, 0x3d, 0x5a, 0xc0, 0x9a, 0x23, 0xce, 0xbf, 0x9b, 0x94, 0xad, 0xd4, 0xe4, 0x6f, 0x0f, 0x0a, 0x64, 0x55, 0x22, 0x26, 0xbc, 0x8b, 0xba, 0xdf, 0xb9, 0x04, 0x3a, 0x5b},
	{0x7b, 0x66, 0x20, 0xcf, 0x63, 0xeb, 0x29, 0xb9, 0x11, 0xc5, 0x5e, 0x18, 0x98, 0x15, 0x2f, 0x69, 0x60, 0xa7, 0xf1, 0x0c, 0xc1, 0x6b, 0x6f, 0xba, 0xd3, 0x2c, 0x83, 0x7d, 0x9d, 0x8e, 0x2b, 0x74},
	{0x7b, 0x9b, 0xcd, 0x1a, 0xe3, 0xfd, 0xd9, 0xd4, 0x74, 0x2e, 0x0d, 0xbc, 0xe1, 0x3c, 0x54, 0x2c, 0xc1, 0x81, 0xb5, 0x0b, 0xa0, 0xf9, 0xd5, 0xe1, 0xca, 0x18, 0x00, 0xf9, 0xb5, 0x84, 0x85, 0xca},
	{0xe7, 0xc9, 0xe2, 0xc8, 0x33, 0x41, 0x31, 0x15, 0xb3, 0x84, 0x3f, 0x79, 0x18, 0xe9, 0x98, 0x5a, 0x51, 0x60, 0xf0, 0x5a, 0x5b, 0xf8, 0x7f, 0x5f, 0xdd, 0x70, 0x27, 0xe3, 0x8f, 0xe3, 0x39, 0xf4},
	{0x36, 0x0d, 0x5b, 0xa8, 0x0e, 0x59, 0xe2, 0x82, 0xa2, 0x39, 0xdf, 0x28, 0x34, 0x4d, 0x4f, 0x74, 0xee, 0xd8, 0x6b, 0xa0, 0xd8, 0x9d, 0xe7, 0x88, 0x05, 0x4e, 0xba, 0x6b, 0x50, 0x03, 0x89, 0xa2},
	{0x89, 0xd6, 0x81, 0x5f, 0x68, 0x39, 0x36, 0x6c, 0x25, 0xad, 0xb6, 0x43, 0xff, 0x6b, 0x5e, 0x19, 0x63, 0xd3, 0xff, 0xd0, 0xce, 0x1a, 0xa7, 0x8c, 0x7f, 0xeb, 0x5a, 0x6e, 0x99, 0xf1, 0xb4, 0xdb},
	{0x1f, 0x36, 0x6f, 0x27, 0xc8, 0x2f, 0x23, 0x81, 0xfc, 0x02, 0x80, 0x4f, 0x8b, 0x8d, 0xa8, 0x2f, 0x3d, 0x35, 0x91, 0xe3, 0x60, 0x90, 0x7c, 0x57, 0x03, 0xc3, 0xa9, 0xed, 0xb1, 0x72, 0x3e, 0x3e},
}

var _test_32_digests = [][32]byte{
	{0x22, 0xd8, 0x35, 0x89, 0xe6, 0x42, 0xe1, 0xb1, 0x40, 0xed, 0x1b, 0x48, 0x48, 0x5b, 0x44, 0xc7, 0x07, 0x9d, 0xf3, 0xb2, 0x04, 0xbe, 0x48, 0x69, 0x42, 0x1d, 0x45, 0x49, 0xf3, 0x9e, 0x2c, 0xc7},
	{0xac, 0xfe, 0x28, 0x1d, 0x11, 0x77, 0x7c, 0x1e, 0x22, 0xe0, 0xb7, 0x16, 0x0f, 0x01, 0x66, 0x92, 0xa7, 0xb3, 0xb5, 0x69, 0xed, 0x12, 0x8d, 0x93, 0xcf, 0xce, 0x27, 0x49, 0xfd, 0x1c, 0x85, 0x01},
	{0xbc, 0xb2, 0xa2, 0x0b, 0x95, 0x58, 0x91, 0x64, 0x1f, 0x3a, 0x5d, 0x80, 0xaa, 0x11, 0x49, 0xa5, 0x1b, 0xac, 0xb7, 0x1e, 0x06, 0x62, 0x45, 0x34, 0xa5, 0x66, 0xd1, 0xc7, 0x5a, 0xa9, 0x68, 0xc9},
	{0x4d, 0xe2, 0xaa, 0x4b, 0xc4, 0x6c, 0x1c, 0x3d, 0x42, 0x65, 0x34, 0x8a, 0x2c, 0x7a, 0x64, 0xa8, 0xd9, 0x8a, 0x82, 0xe4, 0x8b, 0x9c, 0xc9, 0x3c, 0x3c, 0xcd, 0x34, 0x4d, 0x71, 0x76, 0xda, 0x69},
	{0x1e, 0x00, 0xd3, 0xc6, 0x59, 0x37, 0x27, 0x6a, 0x6a, 0xae, 0xa7, 0xd8, 0x37, 0x51, 0xac, 0x74, 0x2d, 0xe0, 0xb6, 0x7e, 0xc5, 0xa8, 0xa7, 0x56, 0x5b, 0x0f, 0x10, 0xba, 0x8a, 0x40, 0xe2, 0x1c},
	{0x30, 0x96, 0xdb, 0x9d, 0xcf, 0xa9, 0x5c, 0xf4, 0xa4, 0xc4, 0xc9, 0xd5, 0xa0, 0x1e, 0xd4, 0x30, 0xe5, 0xe8, 0xad, 0x9d, 0xaa, 0x8e, 0x79, 0x1c, 0x5d, 0x6c, 0xac, 0x1a, 0xb3, 0x65, 0xb5, 0x14},
	{0x7a, 0xee, 0xd5, 0xc9, 0x66, 0x17, 0x59, 0x7f, 0x89, 0xd6, 0xd9, 0xe8, 0xa8, 0xa7, 0x01, 0x47, 0x60, 0xc6, 0x88, 0xfd, 0x2a, 0x7a, 0xf6, 0x1d, 0x10, 0x20, 0x62, 0x7e, 0x7c, 0xd0, 0x1a, 0x0b},
	{0xce, 0x0c, 0x94, 0xa7, 0x41, 0x25, 0xa5, 0xe3, 0x96, 0x77, 0xd6, 0xbd, 0x91, 0xca, 0xe6, 0x06, 0xf3, 0x90, 0xe0, 0x37, 0xcc, 0xc1, 0x2c, 0x7d, 0x97, 0x97, 0xf3, 0x56, 0xf0, 0xbd, 0x66, 0x43},
	{0xbc, 0xb2, 0xa2, 0x0b, 0x95, 0x58, 0x91, 0x64, 0x1f, 0x3a, 0x5d, 0x80, 0xaa, 0x11, 0x49, 0xa5, 0x1b, 0xac, 0xb7, 0x1e, 0x06, 0x62, 0x45, 0x34, 0xa5, 0x66, 0xd1, 0xc7, 0x5a, 0xa9, 0x68, 0xc9},
	{0x4d, 0xe2, 0xaa, 0x4b, 0xc4, 0x6c, 0x1c, 0x3d, 0x42, 0x65, 0x34, 0x8a, 0x2c, 0x7a, 0x64, 0xa8, 0xd9, 0x8a, 0x82, 0xe4, 0x8b, 0x9c, 0xc9, 0x3c, 0x3c, 0xcd, 0x34, 0x4d, 0x71, 0x76, 0xda, 0x69},
	{0x1e, 0x00, 0xd3, 0xc6, 0x59, 0x37, 0x27, 0x6a, 0x6a, 0xae, 0xa7, 0xd8, 0x37, 0x51, 0xac, 0x74, 0x2d, 0xe0, 0xb6, 0x7e, 0xc5, 0xa8, 0xa7, 0x56, 0x5b, 0x0f, 0x10, 0xba, 0x8a, 0x40, 0xe2, 0x1c},
	{0x30, 0x96, 0xdb, 0x9d, 0xcf, 0xa9, 0x5c, 0xf4, 0xa4, 0xc4, 0xc9, 0xd5, 0xa0, 0x1e, 0xd4, 0x30, 0xe5, 0xe8, 0xad, 0x9d, 0xaa, 0x8e, 0x79, 0x1c, 0x5d, 0x6c, 0xac, 0x1a, 0xb3, 0x65, 0xb5, 0x14},
	{0x7a, 0xee, 0xd5, 0xc9, 0x66, 0x17, 0x59, 0x7f, 0x89, 0xd6, 0xd9, 0xe8, 0xa8, 0xa7, 0x01, 0x47, 0x60, 0xc6, 0x88, 0xfd, 0x2a, 0x7a, 0xf6, 0x1d, 0x10, 0x20, 0x62, 0x7e, 0x7c, 0xd0, 0x1a, 0x0b},
	{0xd4, 0x1f, 0xa7, 0x89, 0x8c, 0xf9, 0x05, 0xfc, 0x1e, 0xb0, 0x04, 0xd7, 0xaa, 0x56, 0x35, 0xec, 0x36, 0xf5, 0x0d, 0x41, 0x75, 0x64, 0x34, 0x71, 0xf0, 0x3b, 0x5b, 0xb2, 0xcc, 0xfa, 0x8c, 0xca},
	{0xf8, 0xd9, 0x9e, 0xa7, 0x9c, 0xa1, 0xe0, 0x3a, 0x19, 0x4f, 0xd3, 0x2d, 0xbd, 0x40, 0x3a, 0xa3, 0x28, 0xe8, 0xa4, 0x27, 0x58, 0x44, 0x12, 0xf7, 0x69, 0x01, 0x66, 0xfa, 0xf1, 0x97, 0x30, 0xfe},
	{0x99, 0x7c, 0x24, 0x0e, 0xed, 0x31, 0x0a, 0xda, 0x12, 0x16, 0x0e, 0x06, 0x44, 0xb8, 0x3f, 0xa2, 0x40, 0x52, 0xbc, 0x2d, 0xaf, 0x97, 0x00, 0x01, 0x5d, 0xbb, 0x0d, 0x06, 0x66, 0xb1, 0x59, 0xf2},
	{0x99, 0x43, 0x52, 0x77, 0x28, 0x39, 0x6b, 0xeb, 0x03, 0x51, 0xc4, 0x5f, 0x7d, 0xd3, 0xe1, 0x41, 0x17, 0x66, 0x7b, 0x0e, 0xc9, 0x51, 0x01, 0xa7, 0x39, 0xf3, 0xc8, 0x63, 0x95, 0xa5, 0x92, 0x6b},
	{0xce, 0x6e, 0xab, 0xd2, 0xe8, 0xad, 0x90, 0xad, 0xbe, 0xe5, 0x94, 0x96, 0xa9, 0x98, 0xe7, 0x83, 0x07, 0xa4, 0x0f, 0x8e, 0xe5, 0xb3, 0x5a, 0x05, 0xcd, 0xfd, 0xae, 0x9c, 0x07, 0xad, 0x26, 0xaa},
	{0xf5, 0xee, 0x66, 0x87, 0x00, 0xed, 0xeb, 0x8b, 0xc2, 0x7d, 0x97, 0x52, 0x2d, 0xfc, 0x0a, 0x2a, 0x32, 0x0e, 0x92, 0xd2, 0x91, 0xd1, 0x69, 0x29, 0x9d, 0xb1, 0x3a, 0x65, 0x9f, 0x8e, 0x7e, 0x2a},
	{0x88, 0x4a, 0xc8, 0x81, 0xdb, 0xa6, 0x79, 0x36, 0x54, 0xe9, 0x15, 0x5c, 0xff, 0x06, 0x35, 0x8b, 0x6e, 0x0d, 0xaa, 0x3e, 0x7a, 0x82, 0x7c, 0x4a, 0xfe, 0x8a, 0x91, 0xb4, 0x34, 0xed, 0xe3, 0x17},
	{0xe7, 0x92, 0xa4, 0x91, 0xdc, 0x1d, 0x83, 0xc8, 0x72, 0x5a, 0xd1, 0x27, 0x17, 0x78, 0x2b, 0xc7, 0x67, 0xe9, 0x56, 0xf2, 0xb4, 0x37, 0x51, 0xa1, 0x6b, 0x23, 0x8c, 0xc9, 0x03, 0x3d, 0x90, 0x1e},
	{0xc4, 0x1f, 0xcc, 0x5e, 0xcb, 0x5e, 0x7d, 0x02, 0x12, 0x3f, 0x15, 0x9f, 0x35, 0xf4, 0x49, 0x55, 0xba, 0xc6, 0x47, 0xd2, 0x85, 0x85, 0x61, 0x69, 0xa5, 0x60, 0x7a, 0x32, 0x7f, 0x8e, 0x09, 0x5f},
	{0x60, 0xb6, 0xab, 0xb5, 0x6b, 0x4d, 0xce, 0x6f, 0x1d, 0x77, 0x2e, 0x9b, 0x0d, 0x60, 0x76, 0xe3, 0xcb, 0x79, 0xbc, 0x40, 0x2d, 0x16, 0xf6, 0xa3, 0x06, 0x12, 0x36, 0x71, 0xda, 0xfd, 0x28, 0x89},
	{0x67, 0xdd, 0x7f, 0x26, 0x6d, 0x2e, 0xf3, 0xef, 0x13, 0xb6, 0x09, 0x73, 0x82, 0xbc, 0x73, 0x25, 0x83, 0xc0, 0x34, 0x90, 0xe8, 0xad, 0xf0, 0x17, 0x8d, 0xed, 0xad, 0x29, 0xf7, 0x78, 0x9c, 0x28},
	{0x00, 0xb0, 0xd5, 0xd0, 0x8e, 0x9b, 0xe5, 0xf0, 0x46, 0x8e, 0x60, 0x25, 0x95, 0xe5, 0x3a, 0x46, 0xb1, 0x07, 0x74, 0x97, 0xed, 0x0a, 0x2f, 0x9a, 0x3f, 0xf3, 0x94, 0x2f, 0xb3, 0x12, 0xa1, 0x91},
	{0x8d, 0x36, 0x16, 0xc6, 0x00, 0x88, 0xd6, 0x69, 0xb4, 0x5a, 0x71, 0x18, 0x41, 0xe5, 0x4d, 0xb2, 0xd9, 0x00, 0x7a, 0x17, 0x63, 0x6a, 0x9b, 0x2e, 0x22, 0x12, 0x5b, 0xa3, 0x74, 0x7c, 0x95, 0xc9},
	{0x4e, 0xfc, 0x5c, 0x18, 0xd1, 0x8a, 0x5b, 0x57, 0x7c, 0x86, 0x3e, 0xe2, 0x75, 0x91, 0xf2, 0xb3, 0x5f, 0xd0, 0x92, 0xbc, 0x77, 0xbe, 0x1b, 0xef, 0x1a, 0x7c, 0xe2, 0xd8, 0x8d, 0x7b, 0xef, 0xf7},
	{0xcd, 0x78, 0x15, 0x64, 0x2c, 0x78, 0x57, 0x74, 0x2b, 0xb7, 0xdb, 0x74, 0xe2, 0xab, 0x82, 0xbb, 0x61, 0x32, 0x3e, 0xe4, 0xb1, 0x00, 0xde, 0xb2, 0x35, 0x1e, 0x3e, 0x1c, 0x91, 0x9d, 0x87, 0xde},
	{0x17, 0xcc, 0x52, 0x5c, 0x60, 0x9e, 0xd8, 0xd4, 0xf4, 0x56, 0x28, 0x16, 0xde, 0xde, 0x73, 0xfe, 0xd9, 0x92, 0xb7, 0x99, 0x15, 0x24, 0x1b, 0x40, 0xb0, 0xda, 0x9a, 0xf8, 0x24, 0x38, 0x13, 0xbd},
	{0xd0, 0x45, 0x9b, 0xe3, 0x9a, 0xae, 0x78, 0x41, 0xcd, 0x12, 0x9a, 0x6b, 0x91, 0x58, 0x29, 0x75, 0xae, 0x21, 0xd3, 0xf2, 0x5e, 0x98, 0xab, 0x09, 0xb0, 0xaa, 0x62, 0x96, 0x35, 0x64, 0x18, 0x48},
	{0xd2, 0x5b, 0x10, 0xf1, 0x35, 0xaa, 0x04, 0x49, 0x4e, 0x51, 0x30, 0x0d, 0xb6, 0xbf, 0xa0, 0x9b, 0xa0, 0xf5, 0x66, 0x5f, 0x28, 0xc7, 0x8d, 0xa8, 0x3e, 0x0f, 0xe4, 0xa7, 0xc9, 0xd4, 0x0f, 0x7d},
	{0xb7, 0x80, 0xc2, 0x31, 0xe6, 0x75, 0x0c, 0xad, 0x0f, 0xe8, 0xed, 0x59, 0x34, 0xdb, 0xfb, 0x41, 0xd4, 0x38, 0x73, 0x7a, 0x47, 0x01, 0xb8, 0xea, 0xea, 0x2e, 0x01, 0x8e, 0x4f, 0x09, 0x64, 0x82},
}

func TestHash(t *testing.T) {
	tests := []struct {
		name  string
		count uint32
	}{
		{
			name:  "hash 1 block",
			count: 1,
		},
		{
			name:  "hash 4 blocks",
			count: 4,
		},
		{
			name:  "hash 8 blocks",
			count: 8,
		},
		{
			name:  "hash 16 blocks",
			count: 16,
		},
		{
			name:  "hash 18 blocks",
			count: 18,
		},
		{
			name:  "hash 24 blocks",
			count: 24,
		},
		{
			name:  "hash 32 blocks",
			count: 32,
		},
		{
			name:  "hash 31 blocks",
			count: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			digests := make([]byte, tt.count*32)
			var my_test_32_block []byte
			for i := 0; i < int(2*tt.count); i += 2 {
				my_test_32_block = append(my_test_32_block, _test_32_block[i][:]...)
				my_test_32_block = append(my_test_32_block, _test_32_block[i+1][:]...)
			}
			err := gohashtree.Hash(digests, my_test_32_block)
			if err != nil {
				t.Log(err)
				t.Fail()
			}
			var my_test_32_digests []byte
			for i := 0; i < int(tt.count); i++ {
				my_test_32_digests = append(my_test_32_digests, _test_32_digests[i][:]...)
			}
			if !reflect.DeepEqual(digests, my_test_32_digests) {
				t.Logf("Digests are different\n Expected: %x\n Produced: %x\n",
					_test_32_digests[:tt.count], digests)
				t.Fail()
			}
			digests2 := make([][32]byte, tt.count)
			gohashtree.Sha256_1_generic(digests2, _test_32_block[:2*tt.count])
			if err != nil {
				t.Log(err)
				t.Fail()
			}
			if !reflect.DeepEqual(digests2, _test_32_digests[:tt.count]) {
				t.Logf("Digests are different\n Expected: %x\n Produced: %x\n",
					_test_32_digests[:tt.count], digests)
				t.Fail()
			}
		})
	}
}
