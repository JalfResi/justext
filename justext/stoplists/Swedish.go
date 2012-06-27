package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func SwedishStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x4c,0x98,
0xcf,0x76,0xdc,0xb6,0x19,0xc5,0xf7,0x78,0x8c,0xae,0xb2,0x48,
0xe7,0xa4,0xbb,0x6e,0x95,0xd8,0xb1,0x9d,0x54,0xae,0x4f,0x65,
0xa7,0xa7,0x4b,0x8c,0x88,0xe1,0x40,0x04,0x41,0x1e,0x10,0x64,
0x3d,0xfb,0xf6,0x4d,0xa4,0x17,0xe8,0x42,0x0f,0xd0,0x79,0xb1,
0xfe,0xee,0x07,0x8e,0xd3,0x73,0x12,0x4b,0x33,0x04,0x81,0xef,
0xcf,0xbd,0xf7,0xbb,0xd0,0xf4,0x78,0x76,0xd1,0xf9,0xcd,0x85,
0xec,0x96,0x69,0x74,0xd7,0xe7,0xe2,0x7c,0xad,0xae,0xc6,0x94,
0xdc,0x7c,0x7d,0x71,0x63,0xe8,0x5c,0xc7,0xc3,0xd3,0xf5,0xb5,
0xb8,0xcd,0x17,0x17,0x78,0xda,0x05,0x77,0xe6,0xd7,0x2e,0x54,
0x77,0x2a,0xd7,0x97,0xcc,0xa7,0xec,0x78,0x7b,0x8b,0x9d,0x5b,
0x73,0x17,0x8a,0x5b,0x62,0xef,0xde,0xf3,0xe5,0xc8,0xab,0x1f,
0x5c,0xcc,0x35,0xb0,0xf5,0xc6,0x87,0x63,0x0a,0x1b,0x4f,0xb3,
0x0b,0x29,0xb1,0xee,0x0d,0x5f,0x0d,0xac,0x3b,0x7b,0xf6,0x1c,
0xf9,0xa5,0x53,0x04,0xd7,0x17,0xce,0x39,0x55,0x9e,0xfb,0xdc,
0x15,0xcf,0x2a,0x0e,0x22,0x80,0xa5,0x7a,0x16,0xbc,0xb8,0x53,
0xcc,0x79,0xd1,0x37,0x5d,0xe7,0x16,0x3e,0x4f,0x8f,0x83,0x7e,
0x8c,0x53,0x75,0xd7,0xd7,0x8d,0xd7,0xea,0x66,0xa1,0xa7,0xc4,
0x8e,0x7d,0xc8,0x84,0x76,0x8a,0x8f,0x83,0x5b,0x42,0xc7,0x17,
0x6f,0x02,0x51,0xb0,0xb1,0x5b,0xfc,0x58,0x15,0xfa,0xe2,0xce,
0x53,0xe6,0xa8,0xec,0xab,0x7b,0x9b,0xdd,0x5a,0xdd,0x3a,0xcf,
0x1c,0xf4,0xca,0x92,0x61,0x25,0x4e,0x12,0xe0,0x3f,0x76,0x19,
0xf8,0x7f,0x4a,0x71,0xf0,0xa4,0x9a,0x06,0xa2,0xea,0x38,0x9a,
0x74,0xa8,0x09,0x39,0x79,0xf7,0xc5,0x92,0x3f,0xf3,0xd8,0x13,
0xf6,0x89,0x0c,0xbd,0x1b,0x2f,0x8f,0xb6,0x92,0xc2,0x79,0x97,
0x95,0x5e,0x8a,0x7d,0xcf,0xb2,0xeb,0x33,0xf5,0xa1,0x7a,0xbd,
0x27,0xb0,0xec,0x4b,0x50,0x40,0x23,0x19,0x86,0xe4,0xde,0x5a,
0xf6,0x0b,0x15,0x5b,0x06,0xd5,0x9f,0xbd,0x3d,0xd9,0xb8,0xca,
0xaa,0x51,0x4f,0xea,0xc4,0xd6,0xed,0xb9,0x77,0x9f,0x54,0x83,
0x93,0x7b,0x4f,0x16,0x6b,0x25,0xc3,0x07,0x6a,0x10,0xfb,0xc0,
0x77,0x1c,0x39,0x86,0x45,0xfd,0xec,0x62,0xaf,0x23,0x4e,0xc5,
0x8f,0xae,0x7f,0x9a,0x0a,0xf5,0x3e,0x52,0xd2,0x27,0x15,0xbe,
0x0b,0xcb,0xc2,0xde,0xf4,0xdc,0x22,0xa0,0x20,0xda,0xdf,0xbe,
0x56,0x26,0xb1,0x2a,0x71,0x1d,0x5b,0xa7,0xde,0xe5,0x8b,0x35,
0x84,0x9d,0x37,0x4a,0xc6,0xca,0xbd,0x31,0xe7,0x40,0x7c,0x83,
0xda,0xbf,0x6f,0x4c,0x8f,0xbf,0x86,0x71,0x26,0x1b,0x9f,0x37,
0x92,0xed,0x16,0xf7,0x1b,0x00,0xe9,0xd5,0x89,0x13,0x11,0x1f,
0x55,0xa3,0x8f,0x2a,0x48,0xf6,0x63,0xa6,0x34,0x2f,0x3d,0x5f,
0x5c,0xff,0xdd,0x3e,0x53,0x32,0xdb,0x9a,0x98,0x09,0x5e,0x60,
0x52,0x9b,0x16,0xc7,0xbe,0x2b,0x21,0x71,0xc6,0x8b,0x20,0x13,
0x01,0x48,0x70,0x9f,0xcf,0x14,0x8f,0x42,0xd0,0x47,0xf5,0x5a,
0xf8,0x39,0xa6,0x48,0x99,0x38,0x5d,0x09,0x6a,0x6f,0xc0,0x71,
0x22,0xe0,0x05,0x98,0x3d,0xb9,0x7e,0x5f,0x02,0xf8,0x28,0x2b,
0x58,0xba,0x70,0x74,0x65,0x97,0xbe,0x90,0x01,0x99,0x17,0xbe,
0x1c,0x14,0x34,0xcb,0xcb,0xe2,0x1e,0x63,0x61,0xfb,0x69,0xa8,
0xd3,0x91,0x32,0x74,0xe1,0x31,0x8c,0xfa,0xe5,0x69,0xcd,0x91,
0xe6,0xcd,0xb5,0x7d,0x7c,0x0b,0x0e,0x42,0xee,0x3c,0x05,0xf7,
0x73,0x89,0x09,0xd0,0x64,0xca,0xf0,0xb4,0x72,0xd0,0xf5,0x85,
0x2f,0xd7,0x7e,0x5d,0x6a,0x6c,0xd8,0x16,0xdd,0xf2,0xb4,0xb5,
0x37,0x29,0xa0,0xb1,0x2c,0x1c,0x8b,0x65,0x07,0x06,0xf8,0x08,
0xf9,0x78,0x9d,0x08,0x79,0x95,0xf2,0x83,0x3b,0xc1,0xc1,0x92,
0xc9,0xee,0x67,0x31,0x72,0x49,0x6b,0xa5,0x52,0x64,0xb3,0xb8,
0xbc,0xba,0x8f,0xe1,0x9f,0xbc,0x3a,0xba,0x23,0x3d,0x57,0x86,
0x86,0xd2,0x46,0x5e,0x62,0x2a,0xb6,0xdd,0x2a,0x76,0xaa,0x61,
0x3b,0x4a,0x68,0x71,0x54,0xfb,0x7d,0xae,0x3e,0x11,0x3e,0xe8,
0x84,0x15,0x81,0x33,0xc7,0x20,0x10,0xce,0xc9,0x57,0x0a,0x54,
0xae,0xcf,0x23,0x7b,0x9c,0xc1,0x22,0xaf,0xdc,0xb3,0x11,0x5b,
0xc4,0xce,0xf7,0x42,0xc8,0x93,0x28,0x15,0x37,0x80,0x42,0xe0,
0x28,0x83,0x0e,0xa4,0x96,0x7e,0x04,0x12,0xf9,0xc2,0x2f,0x9b,
0x3b,0x73,0x6e,0xaf,0x90,0xa7,0x71,0x28,0x31,0xf7,0x06,0xa9,
0x35,0x3b,0xea,0xcd,0x5b,0xa0,0xc0,0x14,0xe6,0xcb,0xc3,0x9d,
0x61,0x8f,0x5e,0x25,0x72,0x10,0x28,0x55,0x97,0x87,0x0a,0xcf,
0xce,0x53,0x1a,0x25,0x18,0x70,0x81,0xdc,0x0e,0xee,0xc1,0x14,
0x27,0xd1,0x97,0xeb,0xbf,0x24,0x2d,0x7f,0x10,0x0a,0x44,0x30,
0x7b,0xce,0x17,0x1b,0x55,0x12,0x4c,0x20,0x0a,0xbf,0x81,0x0a,
0xda,0xc3,0x0b,0xc9,0x8e,0x6b,0xfa,0x92,0xa7,0x22,0x1e,0x41,
0x75,0x7e,0x7c,0xd6,0xb3,0x9f,0x4d,0xd0,0x96,0xa1,0xa0,0x53,
0x54,0xc8,0xf8,0xa5,0x9e,0xd2,0x4b,0x44,0xc5,0x68,0x2a,0xa5,
0x60,0xe7,0x4b,0xdf,0xc3,0x90,0x86,0x20,0x29,0xa2,0x90,0xbc,
0xf9,0x4e,0x1c,0x20,0xc4,0x5f,0xa6,0x33,0x04,0x40,0x7c,0x54,
0x12,0xf1,0x66,0x2f,0xf7,0x41,0x9c,0x5f,0xd0,0xc3,0x37,0x8d,
0x74,0x17,0x1d,0xf0,0x7b,0x77,0xa1,0x90,0x50,0xa7,0x98,0x9e,
0x09,0x57,0x15,0x43,0x48,0x68,0xa9,0xd8,0x79,0xf4,0x3c,0x17,
0x54,0x5c,0xa2,0x95,0x92,0x3b,0x70,0x9c,0x75,0xc6,0x5f,0x47,
0x63,0xbe,0xd5,0x50,0xda,0x22,0x5c,0x28,0x28,0xf4,0xa2,0xa3,
0x39,0x0f,0x0d,0x48,0xa0,0xd9,0x3a,0xb3,0xb3,0xd9,0xd8,0x51,
0xdc,0xbd,0xb4,0xff,0xe9,0xfa,0x9c,0x36,0x97,0x80,0x96,0x3b,
0x51,0xfe,0x13,0xdc,0x14,0x93,0x89,0x82,0xda,0xe9,0x84,0xcc,
0x0b,0xdf,0xda,0x80,0x40,0x28,0x6d,0x2d,0x1d,0x39,0xeb,0xfa,
0x3c,0x43,0x81,0x45,0x51,0xf1,0x4e,0x6c,0xea,0x7a,0xcf,0xe4,
0x38,0xc6,0xd4,0x11,0x06,0xcb,0xfd,0x78,0xd4,0x97,0x8f,0x56,
0x18,0x9b,0x16,0xd4,0x59,0x03,0x64,0x16,0x13,0xb4,0x79,0xab,
0xa3,0x03,0x9a,0x84,0x92,0x78,0xf2,0xab,0x2f,0xc9,0xfd,0xa4,
0x7f,0x7c,0xea,0x28,0x1d,0xbd,0x7c,0x16,0x24,0xf2,0x2a,0x75,
0x55,0x74,0x2c,0x14,0xfb,0x77,0x79,0x51,0x2f,0x0c,0x59,0xb0,
0x2e,0x9c,0xaf,0x2f,0x36,0x60,0x4e,0xe9,0x82,0x54,0x49,0xe3,
0x68,0x8f,0x86,0x19,0x19,0x95,0x65,0x88,0xc0,0x5b,0xea,0x08,
0xa2,0xdf,0x36,0xc4,0x93,0x44,0xee,0xd9,0xec,0x47,0x0b,0x9e,
0x13,0x96,0xa6,0x05,0x16,0x3e,0x75,0x11,0x62,0x29,0x32,0xb5,
0x38,0x99,0x64,0x59,0x2c,0x14,0x8e,0x66,0x6a,0x30,0x5c,0xd0,
0x07,0x13,0x43,0x58,0x37,0xc6,0xf4,0x64,0x29,0x31,0x0f,0x4d,
0xd6,0x81,0xe5,0x4a,0xc1,0x69,0xa4,0xd7,0x33,0xaa,0x69,0x7d,
0x73,0x6f,0x50,0xc1,0xb9,0x4c,0x27,0x90,0x80,0xdc,0x9e,0xd5,
0x83,0xd9,0xbd,0xa7,0x02,0x53,0x89,0xfe,0x20,0xe0,0x70,0x46,
0x55,0x7b,0xc2,0xd7,0x9d,0x2f,0x07,0x68,0xc7,0x04,0x61,0xa6,
0x58,0x60,0x40,0xc6,0x08,0x0c,0x46,0x9b,0xd6,0xc2,0x65,0x84,
0x95,0x62,0xa0,0x44,0x3d,0x30,0xb4,0x1e,0x1b,0xcc,0xc1,0x85,
0xa3,0x16,0xa1,0x90,0xee,0x3e,0x09,0x0c,0x6f,0x5f,0xe6,0x79,
0xf1,0xac,0xe1,0xd4,0x0e,0x10,0xbf,0xf6,0xaa,0x31,0x65,0xd4,
0xb1,0x75,0x95,0x12,0x52,0xbd,0x93,0x75,0x3b,0x26,0xcd,0x73,
0x8d,0x4d,0x8d,0x13,0x1b,0xb6,0xf4,0x65,0x3f,0x44,0xe1,0x68,
0xb0,0x58,0x29,0xc7,0xac,0xb7,0x7c,0x39,0x06,0x2b,0x3e,0xb5,
0x45,0x9d,0xbe,0xc1,0xe7,0x80,0x02,0x09,0x90,0x0d,0xbd,0x34,
0x89,0xa1,0x64,0xd0,0x0d,0x7c,0x4c,0x46,0x02,0x0a,0xcd,0x1c,
0x3d,0xd5,0xdf,0x5f,0xfa,0x5e,0x72,0x98,0x60,0xc2,0xad,0xdd,
0x40,0x46,0x86,0x44,0x5c,0xaf,0x1a,0x26,0x99,0xc1,0xaa,0x62,
0x99,0xda,0x21,0x3d,0x82,0x02,0xe5,0x90,0x1c,0x11,0xd5,0x82,
0xad,0xd1,0xbc,0xe0,0x5f,0xb8,0xcb,0x4c,0xab,0x51,0x07,0xb1,
0xc8,0x84,0x12,0x44,0x6d,0x37,0x50,0x9a,0x1e,0x68,0x7b,0x25,
0x33,0x33,0xef,0xdb,0x52,0x31,0x90,0xfa,0x68,0x7e,0xf3,0xe9,
0xa1,0x29,0x36,0x6f,0xa3,0xb6,0xf5,0xa2,0xf5,0x9f,0x4c,0x2f,
0xfb,0x48,0xd4,0x77,0xe6,0x5b,0x50,0xb3,0x83,0x50,0x34,0xcf,
0xac,0xac,0x65,0x42,0x47,0xa5,0x3f,0x41,0xa0,0x17,0x10,0x29,
0x5d,0x6c,0x9c,0xdd,0x73,0x62,0xc4,0x94,0x4e,0x5f,0x18,0xbc,
0x3b,0x54,0x83,0x79,0x30,0x90,0x41,0xa6,0x73,0xd0,0xe0,0x08,
0x8c,0x50,0x4b,0x44,0x57,0x18,0x87,0x0f,0x9a,0x45,0x09,0x18,
0x30,0x22,0x18,0x92,0x9c,0xf7,0xbd,0x63,0xa3,0xa5,0x72,0xf8,
0x5a,0x96,0x19,0x00,0xf7,0xb4,0x5b,0x26,0x4b,0x00,0x75,0x6f,
0x51,0x99,0x26,0xe1,0xa8,0x10,0x3d,0x21,0x2f,0x80,0xf7,0xdb,
0xbd,0x4b,0x6c,0x69,0xdb,0xbb,0xcf,0x08,0x52,0x73,0x47,0xe8,
0x54,0xf3,0x58,0x22,0x16,0xe0,0xf3,0xee,0xbe,0xb9,0x95,0x63,
0xa0,0xc3,0xda,0x74,0xad,0xa6,0xe8,0xb0,0xae,0xd7,0x10,0x7a,
0xb5,0x73,0x81,0x4f,0xb1,0xa6,0x7b,0xf4,0x40,0x66,0x4c,0x98,
0x02,0x6f,0x3f,0x37,0xf4,0xc8,0x54,0xac,0x55,0x56,0xf3,0xf5,
0x26,0x3b,0xf0,0x82,0x67,0x2a,0xab,0xa8,0x74,0x2b,0x85,0x70,
0x2c,0xbf,0xf5,0x76,0x2d,0xd3,0xcc,0x5e,0xc8,0x37,0xaf,0x6b,
0x33,0xa2,0x35,0x80,0x2b,0x6b,0xf9,0x3e,0x55,0x5c,0x10,0xd5,
0x44,0x3f,0xc0,0x9b,0x23,0x51,0xec,0xaa,0xbc,0xb8,0x77,0x66,
0xff,0x20,0xfc,0xd4,0xc6,0xd9,0xc7,0xe6,0x2d,0xa4,0xbb,0x0d,
0xb7,0x3b,0xf5,0x69,0x4d,0x28,0x83,0xc4,0x39,0xc9,0x13,0x76,
0x52,0x46,0x0d,0x46,0x3a,0x2c,0x39,0xa2,0xd0,0x93,0xc8,0xc3,
0x7c,0xf0,0x99,0x84,0x07,0xda,0x13,0x51,0x52,0xd2,0x46,0x84,
0x44,0x19,0xa2,0x2e,0xfc,0xb6,0x17,0x83,0xa6,0x95,0xa8,0xfe,
0x9c,0xed,0x54,0x13,0x10,0x1c,0x0d,0xb2,0xec,0x3e,0x7c,0x30,
0x88,0x26,0x1c,0x9c,0xd4,0xf2,0x36,0x56,0x60,0xf6,0xc4,0x04,
0x6e,0xe1,0x1c,0x34,0xe5,0x25,0xa6,0x5b,0xe4,0xc4,0x80,0xfc,
0xe1,0x99,0x88,0x42,0x23,0xfa,0x28,0x58,0x36,0x8b,0x25,0x2c,
0x6e,0xae,0x4c,0x90,0x93,0x42,0x10,0x41,0xda,0xc3,0x60,0x6e,
0x31,0xee,0xc9,0x6f,0xa8,0x71,0x63,0x17,0x2f,0x91,0xa3,0x7c,
0x96,0xaa,0xb9,0x1c,0x1c,0xee,0x63,0xc8,0x34,0x0c,0xea,0xeb,
0x5b,0x01,0x78,0x6e,0x39,0x2b,0x2c,0x00,0x8c,0x30,0x3b,0x42,
0xdc,0x69,0x87,0x0d,0x82,0xcf,0xca,0x25,0x9e,0xa7,0xd9,0x61,
0x07,0x40,0xe4,0x3d,0x5e,0xe5,0x96,0xb9,0xc8,0xe0,0x37,0x0a,
0x8e,0xc5,0xf1,0x27,0x2d,0x18,0x32,0x4d,0x5c,0x9e,0x56,0x8b,
0x1c,0x0d,0xe3,0x85,0xcf,0x46,0x02,0xab,0x86,0x77,0x7f,0x93,
0x99,0xaa,0xee,0x8d,0xd7,0xbd,0x01,0x27,0x27,0x8b,0x62,0xe3,
0x30,0x64,0xe8,0x5c,0xdd,0x47,0x9c,0x66,0x46,0xb2,0xfb,0x26,
0x54,0xef,0x55,0xce,0x85,0x9d,0xe4,0x6e,0xd0,0xaa,0x85,0xaa,
0x01,0x82,0x8d,0xa9,0x83,0x92,0xf2,0x6c,0xf7,0x28,0x58,0xde,
0x75,0xa1,0x5e,0x0c,0xdf,0xda,0xa2,0xd9,0xec,0x06,0x52,0xf0,
0xf7,0x11,0x5d,0x96,0x8f,0xe6,0xdd,0xe3,0x51,0xa5,0x7e,0x91,
0x9c,0xde,0x98,0x2d,0xa5,0x02,0x65,0xba,0x96,0xec,0x20,0x85,
0x6a,0x3c,0xbf,0xfb,0x11,0xe1,0xa7,0x8e,0x05,0xa7,0xcf,0x9e,
0x92,0x19,0xf7,0x41,0x2c,0x64,0xae,0x2f,0x0c,0x90,0xbe,0xcd,
0x8c,0xeb,0xf3,0x57,0xe1,0x90,0x8d,0x4e,0x88,0x8d,0x29,0x7f,
0x82,0xc2,0xef,0xae,0xaf,0x35,0xd0,0xb4,0x1e,0xf9,0x85,0x3c,
0x6d,0x24,0x52,0x67,0x5a,0x83,0x40,0x01,0x3a,0xc4,0xca,0xfd,
0x9d,0x10,0x22,0x3f,0xbf,0x13,0x05,0x34,0x9f,0x12,0x0a,0x6a,
0x0e,0xb8,0x73,0xff,0x98,0x00,0xe6,0x49,0xb3,0x6a,0xd7,0xdd,
0x7a,0x99,0x8d,0xbf,0x7d,0xe7,0xee,0x36,0x77,0x82,0x44,0x0a,
0xfd,0x17,0x5c,0xc4,0xe2,0xee,0x94,0x5e,0x33,0xf8,0x0b,0xeb,
0x6b,0x48,0x58,0xc5,0x98,0x2d,0x42,0xd2,0x57,0x82,0x8a,0xa2,
0xc0,0x7b,0x39,0x98,0xe6,0x37,0xc8,0xc9,0x28,0x25,0xe3,0x5f,
0xfc,0x57,0x23,0x80,0x4a,0x1d,0x69,0xea,0x27,0xbf,0xe2,0x98,
0x56,0xe3,0x35,0x70,0x37,0x4d,0xe7,0x6a,0x31,0x5a,0x8e,0x92,
0xa9,0x9b,0x1e,0xdb,0x2d,0x46,0xce,0x66,0x18,0xff,0xfb,0x1f,
0x09,0xeb,0x7e,0xa5,0x70,0xe1,0x49,0x17,0x2a,0xd5,0x45,0x31,
0x7d,0x3e,0x4f,0xa3,0x70,0x31,0x18,0x2d,0x9e,0xcd,0x32,0x6a,
0x6a,0x9d,0xd7,0x6d,0xed,0x16,0x3f,0x34,0xa0,0x9e,0xf7,0x29,
0x68,0xaa,0xc2,0xca,0x22,0xa6,0xc6,0x0d,0x0a,0xcd,0x68,0x34,
0xd4,0x1d,0xcc,0x6b,0x98,0xd8,0x09,0x12,0x93,0x7c,0x2e,0x82,
0x58,0x75,0x43,0x6a,0x9a,0xad,0x03,0x6e,0x42,0xaf,0xae,0x24,
0xae,0x0e,0x66,0x04,0x5e,0x6e,0xfe,0xe1,0xff,0x24,0x72,0x1f,
0x62,0xbc,0x2c,0x16,0x59,0x95,0xcd,0x2a,0x55,0x2f,0x14,0x3d,
0xb4,0xfb,0xd9,0x7d,0x7c,0x3c,0xfb,0xdb,0xed,0x4c,0x8d,0xe2,
0x56,0xd2,0x69,0xce,0x1c,0x24,0x5e,0x98,0x42,0x03,0x50,0x31,
0x0d,0xee,0x5d,0xd0,0x05,0x14,0x80,0x03,0x6e,0x86,0x58,0x51,
0x28,0x38,0x47,0x95,0x74,0x9f,0x33,0x79,0x37,0x3f,0x37,0x6b,
0xf8,0xbd,0x7b,0x1f,0xb2,0x98,0xfe,0x0d,0x31,0x88,0x92,0x38,
0x6d,0xd6,0xa9,0x71,0x18,0x56,0xb4,0x7b,0xd2,0x2d,0xba,0x86,
0xd8,0xbd,0x03,0xef,0x94,0x17,0x96,0x5a,0x77,0xa0,0xd7,0x62,
0x08,0x36,0x5b,0x4f,0x0a,0x5f,0x6e,0x00,0xbe,0xb4,0x01,0x79,
0x7d,0x1d,0x68,0x55,0x81,0x24,0xf2,0x4c,0x92,0x06,0x93,0x4b,
0x18,0x39,0x61,0x77,0xcd,0xb1,0xc8,0x23,0x12,0x85,0x89,0x3f,
0xd6,0xfe,0x83,0x84,0xf4,0x23,0x51,0x51,0xe0,0x32,0x71,0x91,
0x47,0x56,0xd3,0x7a,0x3c,0xca,0x69,0x37,0x14,0xa3,0x75,0xfb,
0x4d,0x66,0xb1,0x94,0x61,0xb2,0x4c,0xaa,0x35,0x94,0x43,0xec,
0x0a,0x6e,0x17,0xcc,0x7d,0x64,0x26,0x9c,0xd1,0x86,0xb9,0xd3,
0xc4,0x6d,0xc6,0xd3,0xfd,0x45,0xc6,0x8b,0x32,0xe8,0x92,0xc1,
0xc4,0x6c,0x60,0xf8,0x69,0xe2,0x92,0x70,0x91,0xe7,0x64,0x10,
0xb4,0xcb,0x8c,0xfc,0xf6,0xab,0xba,0x86,0xde,0xea,0xde,0xf1,
0x5d,0x74,0x0f,0xfa,0x7b,0xc3,0x84,0x6c,0xc9,0xf6,0x44,0xf5,
0xc4,0x7a,0x67,0x2b,0x84,0x28,0xf4,0x47,0xc7,0x99,0x01,0x68,
0xec,0x33,0xb4,0xce,0xe8,0xa4,0x2e,0x19,0x34,0xfc,0xf0,0x2b,
0xd7,0x8a,0x7b,0x9f,0xc6,0xeb,0xab,0x52,0xdc,0x90,0x24,0x7b,
0x71,0x74,0xbf,0xde,0x50,0x22,0x47,0x5f,0x5b,0x54,0x38,0x18,
0xc9,0xd2,0x37,0xbf,0xfa,0xa7,0x3f,0xff,0xf0,0xc3,0x1f,0x5b,
0xb1,0xef,0xf6,0x49,0x1d,0xdb,0x9f,0x47,0xa4,0x4e,0xed,0xbe,
0xd1,0x4c,0x97,0xab,0x5c,0x86,0xd2,0xc4,0x8c,0xf2,0xee,0xf1,
0x1c,0x4e,0xbb,0xb7,0x03,0x37,0xf6,0x76,0x9d,0xf6,0x66,0x73,
0xd1,0xf6,0x19,0xe1,0xd5,0x5f,0x2e,0x16,0xa9,0x73,0x8d,0x3c,
0xfe,0x84,0xe8,0x42,0x17,0x19,0x5c,0xe5,0x2f,0x24,0xed,0xd8,
0x3d,0xdc,0x68,0xc9,0x2d,0x0b,0xae,0x55,0x2e,0xcb,0xb8,0x56,
0x79,0x56,0xcd,0x6a,0x2c,0xeb,0xd8,0x5c,0xb2,0x62,0xe0,0xc6,
0x26,0x37,0x28,0xe6,0xa8,0x94,0x36,0xa6,0x19,0x23,0x1b,0x97,
0x17,0xca,0x79,0xa7,0x2b,0xcc,0xbb,0x60,0xcd,0x3e,0x4e,0x62,
0xda,0xd2,0xaa,0xc9,0xdd,0x01,0x18,0xec,0xd6,0x09,0x0b,0xa2,
0x11,0x14,0x34,0x3b,0xac,0x8f,0x0a,0x4b,0x7c,0x15,0x15,0x65,
0x94,0xb8,0x76,0x88,0x36,0x0a,0x9c,0xe9,0xb0,0xdb,0x69,0x6e,
0x63,0x8f,0x62,0xab,0xc1,0x11,0xf4,0xcb,0x4f,0x5a,0xb2,0xc4,
0x2c,0xc1,0x15,0x7a,0x78,0xa4,0x46,0x81,0x01,0xdf,0x13,0x0e,
0xf2,0x81,0x9f,0xea,0xdd,0x1d,0x23,0xc4,0xee,0xf8,0x87,0x46,
0xe8,0x1b,0x2b,0x0c,0x62,0xa0,0xf3,0x09,0xb5,0x17,0xa9,0x22,
0xe3,0xe1,0xbe,0xfd,0xf9,0xa7,0xad,0x93,0x4a,0x11,0x83,0x0a,
0xc1,0x38,0x13,0x5c,0x7f,0xf7,0xa7,0xc5,0xfd,0x2f,0x00,0x00,
0xff,0xff,0x84,0xd3,0x66,0x7c,0x06,0x13,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}