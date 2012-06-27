package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func SundaneseStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x4c,0x98,
0x51,0x96,0xdb,0x36,0xd2,0x85,0xdf,0x6b,0x19,0x79,0xca,0x83,
0xa2,0x3d,0xb4,0xe3,0xf3,0x27,0xb6,0x63,0x27,0xe7,0xef,0xce,
0x78,0xf2,0x08,0x89,0x68,0x12,0x4d,0x12,0xe0,0x90,0x80,0x35,
0x5a,0x52,0xaf,0xa3,0x37,0x36,0xdf,0x2d,0xaa,0xdb,0x3e,0xc9,
0xb1,0x65,0x09,0x04,0x0a,0x55,0xb7,0xee,0xbd,0xc5,0x3a,0x44,
0xeb,0x92,0xe5,0xc6,0x9f,0x39,0x58,0xbe,0x86,0x97,0xe7,0x1a,
0xac,0x3c,0xda,0x53,0x6c,0xb9,0xb7,0x90,0x3b,0xfb,0x90,0xbb,
0x92,0x5f,0x9e,0xb7,0x14,0x8e,0xf6,0xd7,0x5a,0x96,0x94,0xb7,
0x64,0x1f,0xc3,0x25,0xd8,0x16,0xa6,0xb0,0x85,0x21,0x3d,0x25,
0xfb,0x14,0x4e,0x6d,0x09,0xf5,0xe5,0x39,0xf3,0x4c,0xb3,0x31,
0x9c,0xc3,0x1c,0x6a,0xc8,0x56,0x8b,0x05,0x7b,0x17,0xd6,0x50,
0x0f,0xac,0x5f,0xd9,0xbc,0x63,0xaf,0x60,0x29,0xdb,0xd8,0x2c,
0x6d,0x2c,0xb5,0x9a,0x8c,0xb5,0x6c,0x98,0x62,0xb3,0x1a,0x5a,
0xb6,0x53,0x62,0xc9,0x92,0x46,0x82,0xb0,0xcf,0x61,0x7b,0x79,
0x1e,0x92,0xdd,0xce,0x99,0x38,0xa3,0x8b,0xab,0x85,0xcd,0xc2,
0x55,0xcf,0x12,0xf7,0x03,0xd7,0xa8,0x43,0xa8,0x16,0xd6,0x68,
0x4f,0x81,0x1b,0x8d,0x81,0xaf,0x1f,0xcb,0x6a,0xa7,0xab,0xb1,
0xcb,0x6a,0x0f,0xbe,0x6d,0xb4,0xd2,0xbf,0x3c,0xdb,0x6f,0xeb,
0xcb,0x73,0x5f,0xd6,0x14,0xf2,0xd1,0xde,0x6b,0x83,0x29,0xb6,
0x4b,0x1a,0x74,0xff,0x48,0x84,0x95,0x30,0x78,0xf4,0x92,0xea,
0x60,0x25,0xdb,0x1c,0x72,0xbf,0x72,0x3b,0x8f,0xa6,0x8f,0x6d,
0xb3,0x1c,0x97,0x64,0x5c,0x76,0xb0,0xc7,0xb5,0xcc,0xdc,0xd8,
0x2e,0x43,0x3a,0x0f,0xf6,0xa9,0xf0,0xf4,0xfd,0x67,0xbb,0x10,
0xdc,0x89,0xd4,0x70,0x19,0x16,0xde,0x7f,0x3e,0x92,0x44,0xb6,
0x19,0xc8,0xf2,0x16,0x4f,0xad,0xda,0x7d,0x98,0x5f,0x9e,0xf9,
0x7f,0xb0,0x33,0x0f,0x77,0x8d,0x2b,0xc4,0xdc,0x87,0xe1,0x60,
0xa9,0x92,0xa6,0x13,0xff,0xad,0x61,0x08,0x76,0x2a,0x7d,0xe0,
0xce,0xaf,0xc9,0x1c,0xc2,0xb7,0x48,0xf0,0xf1,0x35,0xf6,0x5c,
0xb4,0xfa,0xcc,0x5a,0x7e,0xda,0x6c,0x88,0x59,0xa1,0x87,0x5c,
0xf5,0x4d,0x97,0x08,0x99,0xcb,0xd6,0x81,0x2c,0x5f,0x55,0x99,
0xf7,0xc9,0x66,0x6e,0x9b,0x7b,0x72,0xce,0x69,0x6f,0x77,0x7a,
0xa7,0x50,0xf5,0x05,0x47,0x90,0xc4,0xb6,0xc5,0xce,0xc2,0xb4,
0x15,0xbb,0xc4,0x0b,0x61,0x8d,0xe4,0xa0,0xf4,0xd1,0x1a,0xbf,
0xf7,0xf6,0x31,0x2e,0xfa,0x4b,0xb5,0xa7,0x80,0x1f,0xc3,0xe8,
0x45,0x7d,0xcd,0x91,0x55,0x72,0xed,0x27,0x47,0xb6,0xf0,0xa0,
0x2f,0x61,0xac,0x8d,0x33,0x72,0x07,0xa2,0x8e,0xaf,0x1f,0x0e,
0xd6,0xf1,0x2b,0x75,0x9a,0xd8,0x76,0x60,0x59,0x01,0x8b,0xab,
0x39,0xfc,0x52,0xdd,0xec,0xbe,0xe5,0x0e,0x84,0x35,0xb2,0x7a,
0x89,0x54,0xf5,0x41,0xb7,0xb8,0x81,0x4d,0x67,0x0f,0xa6,0x3c,
0x96,0x1c,0x39,0x9a,0xb3,0xc6,0x10,0xb6,0xb6,0x90,0x4d,0x5d,
0x9c,0xdd,0xa6,0x72,0x22,0xa8,0xc2,0x83,0x27,0x22,0x9c,0x8c,
0x2b,0xe9,0x26,0xba,0xed,0x9c,0xe6,0x04,0xe2,0x46,0xd5,0x8a,
0xb0,0x29,0x02,0x00,0xd8,0x88,0xdf,0x4b,0x0b,0x46,0xa9,0x06,
0xc9,0x50,0xea,0x55,0x9b,0xd8,0x58,0x1b,0xcf,0x61,0xe1,0x4b,
0x1e,0xf7,0xef,0xda,0x4c,0xc6,0x56,0xbb,0xf3,0x80,0x78,0x10,
0xac,0xea,0xe0,0x2d,0x2c,0x71,0x15,0x98,0xa7,0x09,0xa8,0xc5,
0xbc,0x57,0x52,0x77,0x04,0xdc,0xd9,0xae,0x9e,0xb7,0x32,0x47,
0x7b,0x57,0xa8,0xe0,0x91,0x9a,0xb1,0x9a,0xc3,0xfd,0x9f,0x07,
0x5b,0x0a,0xd5,0x9a,0xc2,0x4c,0x8c,0x9c,0x4f,0xeb,0xc4,0x99,
0x96,0x12,0xb0,0xe9,0x15,0xf6,0xe0,0x79,0xb6,0xad,0x17,0xed,
0x3c,0xb2,0xe3,0xe6,0x59,0xa7,0x74,0xbf,0xa6,0x35,0x9e,0x0a,
0x78,0x20,0x85,0x69,0xe5,0xb7,0x1d,0x64,0x73,0x3a,0x91,0xeb,
0xb7,0x9f,0x0f,0xea,0x17,0x4e,0x3b,0x85,0x1e,0xf0,0xd0,0x80,
0xb4,0xa6,0x30,0x17,0xb6,0x34,0xed,0x89,0x09,0x57,0xda,0x64,
0x4b,0xe3,0xec,0x87,0x1d,0xec,0x37,0x72,0x51,0x8f,0xf4,0x01,
0xf7,0xf1,0xcf,0xb4,0x70,0xa2,0x47,0x26,0x1d,0x74,0x25,0xbd,
0x5b,0xd5,0x0e,0x3d,0x59,0xf1,0x0d,0x14,0xb5,0x6a,0xfb,0xe3,
0x2e,0x47,0xc1,0xc9,0xea,0xa5,0x28,0xb3,0x94,0xc0,0x72,0x10,
0xd0,0x08,0x2a,0xcc,0x69,0x03,0x05,0x54,0x43,0xf5,0x20,0x22,
0x1e,0x9f,0xd2,0xc5,0x31,0xf3,0xa9,0x65,0xf2,0x1a,0x88,0xf9,
0x94,0x48,0x35,0x51,0xdf,0xb7,0x11,0x8a,0x99,0x93,0x9e,0x20,
0x99,0x9c,0x69,0x24,0x8a,0xd4,0x0b,0xb4,0xa7,0xd8,0x84,0x08,
0x56,0x4e,0xaa,0x65,0x19,0xed,0x31,0xad,0x04,0xe7,0x77,0x7b,
0x7b,0xf4,0x08,0x3c,0x92,0x5d,0xc9,0x5e,0x62,0x35,0xf5,0xfb,
0x00,0x61,0x9c,0x42,0x07,0xd7,0xad,0x05,0x36,0xea,0x9a,0xa2,
0xca,0x4f,0x8d,0x5a,0xd4,0x34,0xb5,0xb7,0x28,0x8e,0xb7,0x68,
0x69,0x38,0x22,0xa4,0x79,0x28,0x3f,0x6c,0x31,0x90,0x08,0x2a,
0x70,0xb5,0x56,0xe9,0x4f,0x1d,0x9c,0xb6,0x9a,0x66,0x40,0x39,
0x5d,0xed,0xb7,0x96,0x95,0x09,0xff,0x89,0x90,0x9f,0xa0,0x36,
0x51,0xd9,0x53,0x98,0x84,0x43,0x92,0x0b,0xa9,0x79,0x1e,0xc9,
0x0e,0x0c,0xbb,0xb2,0xec,0xda,0x0e,0xaa,0xf3,0x8d,0xdf,0xa6,
0x36,0xd0,0x4a,0xaa,0x30,0x07,0xc2,0x4f,0x35,0x8d,0xea,0x9d,
0xab,0xb0,0x27,0x2a,0xdd,0xbb,0xf1,0x16,0xb0,0xae,0x36,0x4d,
0xb4,0xed,0x07,0xfa,0xff,0x91,0x83,0xa1,0x68,0xfa,0x02,0x3c,
0x6d,0xe4,0x93,0x93,0x3f,0x35,0xf8,0x93,0x06,0x55,0x92,0xbd,
0xa7,0x1b,0xa4,0x3f,0x9f,0x9c,0x47,0xc3,0x6c,0xce,0x7a,0x09,
0x64,0x7e,0x8f,0xd0,0x3b,0xd5,0x4f,0x3b,0xda,0xff,0xc1,0x86,
0x77,0x08,0x80,0xda,0xec,0xaa,0x0e,0xc8,0x57,0xbf,0xcf,0xad,
0xfb,0x09,0x9b,0x9a,0xac,0x09,0xf0,0x42,0xeb,0x9f,0x28,0x89,
0x2a,0x74,0xb0,0x36,0x72,0x03,0xf2,0xf2,0xfa,0xcd,0x11,0x4a,
0xa7,0x98,0xba,0x54,0x93,0x04,0x44,0x2e,0xdc,0xc7,0xa9,0xcc,
0xe0,0xb4,0xff,0x9e,0x85,0x23,0x14,0x3e,0x39,0xa7,0x52,0xaa,
0x4b,0x69,0x13,0x6c,0x44,0x96,0x4f,0xb0,0x5b,0x1b,0x4d,0x2c,
0xf3,0x8b,0x53,0x4d,0x7a,0x84,0x5f,0x05,0x81,0xc9,0x9e,0xda,
0x2c,0x00,0x38,0x20,0x26,0x3d,0x48,0xe3,0x5f,0x12,0x44,0xd0,
0xdb,0x1d,0x22,0x51,0x1e,0x2b,0x5b,0xad,0x08,0x43,0xb1,0x65,
0x85,0x1a,0x4e,0x69,0x4a,0xf5,0x6a,0xef,0x3f,0x7d,0xd0,0x2d,
0x9f,0xc2,0x9a,0xa3,0x90,0x44,0xdd,0xa3,0x48,0xda,0x77,0x5d,
0x00,0x47,0xa1,0x8d,0x76,0xc2,0x3b,0xd2,0x37,0x2b,0x14,0x74,
0xdf,0xe6,0xa8,0x7d,0x55,0x82,0x09,0x1c,0x59,0x0c,0x50,0xd4,
0x7d,0xd3,0x15,0xf4,0x5d,0x8d,0x6b,0xdb,0x9c,0x50,0xc8,0x26,
0x6c,0x2b,0xda,0xf9,0x41,0x42,0x21,0x22,0x51,0x3b,0xcd,0xad,
0xcb,0x75,0xa9,0x4f,0x95,0x80,0xc7,0x5c,0x2e,0x22,0x26,0xa7,
0xfb,0x37,0x7e,0xb7,0x36,0xb7,0x99,0xf8,0x74,0xc9,0x83,0x95,
0x6f,0x14,0x8b,0x48,0x1d,0x91,0x6f,0x61,0x1c,0x85,0xa3,0x27,
0x91,0xfb,0xc9,0x59,0x74,0x21,0x49,0xac,0x40,0x38,0xd0,0xa6,
0x96,0xcf,0x35,0x41,0xde,0xc2,0x65,0xd4,0x4e,0xdf,0x6f,0x7a,
0xe4,0xec,0xc7,0x47,0xd2,0x94,0xab,0x7d,0x91,0x5e,0xeb,0x8e,
0x3d,0xa0,0x16,0x0b,0xe3,0x09,0x5e,0x9e,0xc7,0xd0,0x45,0x20,
0xde,0x01,0x44,0xa7,0x19,0xa1,0x79,0x29,0x70,0x31,0xfc,0xba,
0x51,0x4e,0x7a,0xf8,0x5d,0x84,0x6a,0xe8,0xc8,0x5c,0x10,0xb4,
0x57,0xfe,0x54,0xa3,0x67,0x50,0x84,0x94,0xcc,0x0b,0x6c,0x44,
0xbe,0x09,0x71,0xbb,0x2d,0xc6,0x42,0xec,0xaa,0x7e,0x1f,0x5d,
0x1d,0x0e,0xb7,0xd4,0x1d,0x29,0x71,0x18,0xd1,0x8a,0x55,0xb9,
0x5b,0x94,0x2b,0x3d,0x04,0x10,0x80,0x10,0x71,0xf3,0x88,0x74,
0x5e,0x72,0xfa,0x3b,0x09,0xcb,0x64,0x3e,0x41,0xba,0x5c,0x48,
0x4d,0x06,0xce,0x38,0x6f,0x2f,0x14,0x60,0xdc,0x5b,0xe3,0x49,
0x4d,0x09,0x57,0x43,0x58,0xa2,0x0a,0x25,0xfb,0x95,0xa8,0xf6,
0x9c,0x02,0x47,0xf5,0x59,0x78,0xa2,0x17,0xe8,0x8a,0x9d,0x54,
0xb6,0xab,0xe7,0x6a,0x2b,0xdc,0x0f,0x35,0xa1,0x3c,0xf6,0x8d,
0xbf,0xc2,0x29,0x4e,0x76,0x6e,0xab,0x27,0xec,0x04,0xdb,0x90,
0x70,0x14,0x8b,0x74,0xad,0xa5,0xf5,0x83,0xfd,0xe1,0xb4,0xcd,
0xb9,0x1d,0x7e,0xa0,0x4b,0xb5,0x4d,0x60,0xe5,0x11,0x9b,0x31,
0x36,0x8a,0x43,0x37,0xde,0xef,0x32,0xcd,0x89,0x8a,0xed,0xaf,
0x15,0x46,0xba,0x09,0x35,0xf7,0x1a,0xc9,0x97,0x50,0x47,0xbe,
0x21,0xec,0x85,0x1b,0x78,0x07,0x4b,0x85,0x48,0x00,0x3d,0xe6,
0x51,0xab,0x8d,0x9e,0xd4,0x7d,0x9f,0x5f,0x55,0x0b,0x3b,0x51,
0xd7,0x04,0xd7,0xab,0xca,0xae,0x26,0xf6,0x65,0x67,0x44,0xc8,
0x65,0x96,0xa9,0xc3,0x43,0xe9,0x66,0xa0,0xe1,0x76,0xbf,0xa5,
0x4d,0x38,0x32,0x34,0x19,0x59,0x15,0x30,0x38,0x25,0x5e,0xac,
0x48,0x68,0xa8,0xa1,0xd7,0x6d,0xe4,0x7a,0x14,0x02,0x28,0x86,
0xab,0x14,0xcf,0xfe,0x6a,0x2b,0x0f,0xec,0x9d,0xbe,0xdb,0x04,
0x8f,0xf9,0x16,0xcd,0xc6,0x4f,0x27,0x59,0x19,0x00,0x7a,0x65,
0x13,0x30,0x16,0x06,0xea,0x59,0xc9,0xc4,0x87,0x0d,0x3d,0x33,
0xd2,0xc0,0x7a,0x92,0x42,0xbb,0x0b,0x4c,0x55,0x58,0xfa,0x5b,
0x77,0x3f,0x90,0xfe,0x15,0x6b,0x11,0x86,0x09,0x6f,0x47,0x68,
0x00,0x02,0x5a,0x40,0xc9,0x42,0xa6,0xc3,0x67,0x1d,0x2f,0xcd,
0x40,0x64,0x02,0x2c,0xb1,0xe7,0xe1,0x77,0x50,0xc8,0x59,0xac,
0x84,0x6e,0xd5,0xf6,0x35,0x4e,0x71,0x11,0xce,0x79,0xaa,0x83,
0xd7,0x05,0x4f,0xb7,0x40,0x23,0xed,0xa8,0x05,0xf1,0xbf,0x60,
0x71,0x8a,0x07,0x49,0x19,0x04,0x36,0xb0,0xc2,0xa3,0x84,0x13,
0x37,0xb7,0x1c,0x88,0x2f,0xd7,0x7b,0xed,0x25,0x6e,0x47,0x1e,
0x44,0x19,0x27,0x57,0x23,0xb7,0x16,0x7c,0x88,0x8d,0x34,0x0e,
0xa4,0x0f,0xe8,0xa2,0x9f,0xd8,0x4b,0x09,0x01,0xe9,0x21,0xa5,
0x63,0x5a,0xc3,0x2f,0xfa,0x83,0xa0,0x5d,0xaf,0x65,0xea,0x1a,
0x72,0x89,0xe4,0x83,0x4f,0x28,0xa0,0xef,0xc5,0x1f,0xe0,0x82,
0x92,0x8d,0xf6,0xb3,0xfb,0x43,0xf2,0x86,0x40,0xfb,0xc5,0x91,
0xac,0x08,0xd3,0x9d,0xa2,0x1e,0x59,0x0a,0xd5,0x84,0xf6,0x9b,
0x6e,0x03,0x9c,0xc1,0x9c,0x2c,0xa1,0x3c,0xa2,0x2c,0x89,0xff,
0xf5,0x90,0x66,0x57,0xaa,0x58,0x23,0x88,0x09,0x67,0xe8,0x4e,
0x1c,0xce,0xf7,0x3f,0x14,0x4c,0x7d,0x75,0xf5,0x94,0xa0,0xf6,
0x2e,0x52,0x7d,0xfa,0x26,0x2e,0x14,0x81,0x4e,0x40,0x8e,0x6f,
0x54,0x9e,0x91,0x7b,0x78,0xf4,0x5b,0xac,0xf2,0xcd,0xb3,0xdd,
0x41,0x4b,0xa1,0xc2,0x7f,0x43,0xdb,0xe8,0xeb,0x5d,0x32,0x59,
0xed,0x5c,0x16,0x5d,0x0c,0x10,0x6e,0xe5,0x18,0x57,0xa1,0x98,
0xe9,0x89,0x9d,0x25,0x66,0x99,0xb6,0x21,0xd1,0x15,0xb3,0x38,
0xa4,0x4b,0xdf,0x2d,0xc6,0xaf,0xb2,0x3b,0x7d,0xac,0xf2,0xda,
0x42,0x82,0x12,0x3d,0x21,0x90,0x53,0x91,0x61,0xd3,0x04,0xd2,
0x9e,0x1a,0xe6,0xa9,0x51,0x9b,0xb1,0xd0,0x52,0xc0,0x97,0xc8,
0x16,0x3c,0x84,0x8b,0x21,0xd0,0x03,0x69,0x2c,0x64,0xa9,0x6a,
0xd2,0xa6,0x76,0x95,0x45,0x63,0x73,0xfc,0x4e,0x99,0x22,0xd8,
0x15,0x7e,0x49,0xe7,0xcd,0x67,0xca,0x99,0x4f,0x30,0x71,0xc3,
0x0e,0x38,0x52,0x21,0xf8,0xb9,0x20,0xb0,0xd9,0xce,0x2e,0x2d,
0x0b,0x34,0x2e,0xcb,0xe2,0xbc,0x87,0x37,0x52,0xf5,0x96,0x42,
0x9b,0xe8,0x60,0x81,0xd6,0x7a,0xcd,0x19,0xc2,0x21,0xf2,0xa3,
0x4b,0x88,0xea,0xdc,0xe8,0xcb,0xa2,0x1e,0xa1,0x87,0xa9,0x81,
0x60,0x24,0x06,0xa6,0x12,0xdc,0x71,0x23,0x63,0xd2,0x86,0x48,
0xbb,0x8e,0x9a,0x9c,0x31,0x19,0x45,0x70,0x91,0x32,0x43,0xf7,
0xb2,0x80,0x97,0x88,0xa9,0x22,0x97,0x64,0x90,0xb6,0x54,0xa0,
0x77,0xcc,0x03,0x20,0x03,0x38,0xa4,0xae,0x4d,0x07,0xfb,0x7f,
0x30,0x4a,0x76,0x5d,0xb6,0xbd,0xa3,0x73,0x5c,0x7b,0x99,0xdf,
0x65,0xa5,0x7b,0xef,0xbf,0x0f,0x09,0x78,0x06,0x49,0x2f,0xed,
0x4d,0xf4,0xf6,0xd3,0xbf,0x7f,0x32,0xda,0x43,0xf6,0x91,0x2a,
0x6e,0x31,0x33,0xa9,0x20,0x24,0x17,0xb0,0xeb,0xda,0x44,0x16,
0xfb,0x4c,0x05,0xce,0xf6,0x55,0x2a,0x0b,0x68,0xd4,0xc7,0x40,
0xb5,0xe1,0x20,0xae,0xac,0xad,0x5a,0x2b,0x85,0xc2,0x95,0x8d,
0x64,0x80,0x60,0xdc,0x1e,0x33,0x52,0xc1,0xf4,0x40,0x5e,0xae,
0x87,0xfb,0x90,0x4c,0x66,0x05,0xd7,0xe4,0x4d,0xfd,0xc8,0xf7,
0x3f,0xef,0x63,0xdc,0xb9,0xcc,0x33,0x15,0xf8,0x23,0x50,0x94,
0x3a,0x14,0xae,0x2c,0xb2,0x18,0x95,0x20,0xec,0x9e,0xbb,0xb9,
0x4a,0x76,0x59,0x72,0x21,0x77,0xbb,0x9e,0x5d,0x22,0x64,0xa6,
0x71,0x68,0x50,0xb6,0xdd,0x11,0x93,0x36,0x3c,0x45,0xcf,0xce,
0x23,0xb8,0xc0,0xe6,0x8c,0x37,0x0c,0xb9,0x9d,0xf5,0xf9,0x41,
0x4a,0xb7,0x7b,0xe0,0x4e,0x5c,0x35,0xd8,0x87,0x47,0x53,0x1a,
0x31,0x3e,0x1d,0xee,0x97,0x74,0xa5,0x7c,0x9e,0x1a,0x28,0x6c,
0x99,0x3e,0x92,0xbe,0xc8,0xb9,0xa5,0x79,0x29,0x34,0x07,0x54,
0x7e,0x87,0x62,0xd5,0x50,0xc5,0x24,0xe3,0x9e,0x78,0xc5,0xaa,
0xed,0x31,0x7a,0x8a,0xdb,0x07,0x0e,0x97,0x6c,0xcd,0x3c,0xb8,
0x0b,0x67,0x3e,0xf9,0x25,0x92,0xe7,0xf4,0xcd,0x75,0x21,0x78,
0x10,0x85,0xfd,0xea,0xe4,0xb5,0x50,0xbc,0x1e,0x27,0x63,0x32,
0x0f,0x37,0x4a,0xea,0xe8,0x62,0x77,0x5b,0x5f,0x9c,0x61,0xa1,
0x85,0xb3,0x1a,0x43,0xac,0x7f,0xb3,0xde,0xa4,0xed,0xa3,0x8b,
0xdd,0xdf,0x39,0x55,0x76,0x9a,0xbf,0xab,0x15,0xc3,0x2d,0x3f,
0x4f,0xb4,0x25,0x80,0xb3,0xcf,0x12,0xa9,0xba,0x36,0x86,0xaa,
0x15,0x68,0x32,0x22,0x10,0xd9,0x17,0xb9,0x73,0xec,0xbc,0xcb,
0xce,0x19,0xbe,0x59,0x88,0x99,0x1b,0xb8,0xa7,0x76,0x1d,0xec,
0x23,0x9a,0xae,0xd0,0xf6,0xf9,0xfd,0x4f,0x06,0xa7,0x8b,0x26,
0xe6,0xcd,0x05,0x09,0xae,0xdf,0x18,0xda,0xf8,0x27,0xe5,0xa5,
0xdc,0x20,0x54,0x76,0x70,0xb3,0x4f,0x3e,0xf9,0x90,0x52,0x0d,
0x55,0x4b,0x10,0x5a,0x68,0x54,0x9a,0x13,0x03,0xa7,0xd9,0x86,
0x62,0x4e,0x05,0x3d,0xd3,0x2d,0x35,0xc0,0xef,0xa4,0xe8,0x7a,
0xd4,0x05,0xb6,0x86,0x6a,0x53,0xcf,0x1a,0x77,0x6a,0x1a,0xc4,
0xa0,0xea,0x33,0x04,0x78,0xb6,0x87,0x32,0x5e,0x0b,0x65,0x91,
0xd7,0xfd,0x0c,0xe0,0x66,0xc8,0x42,0xe1,0xfe,0x90,0x93,0x1c,
0x25,0x72,0x3d,0xdc,0xb0,0xcf,0x43,0x43,0x2c,0x28,0x4c,0x97,
0xce,0x85,0xe1,0x4d,0x89,0x02,0x75,0x1b,0x0a,0xbd,0xa9,0xfa,
0x40,0xa0,0x10,0xf0,0xdd,0x3e,0x74,0x72,0x94,0x53,0x8a,0x74,
0x9a,0xa8,0x61,0x27,0x80,0x67,0xef,0x23,0x48,0x91,0xb4,0xef,
0xfb,0xd1,0xf7,0xc2,0xb0,0x54,0x44,0x10,0xc7,0x60,0x42,0x0b,
0x90,0xe8,0xb0,0x46,0x8d,0x95,0x52,0x0e,0x28,0x83,0x9c,0xf2,
0x61,0x8d,0xdb,0x9c,0xfc,0x95,0x00,0x16,0x77,0x07,0x95,0x68,
0xe3,0xab,0xd4,0x03,0xd1,0x92,0x55,0x93,0x52,0xee,0xf2,0x1e,
0x30,0x05,0x94,0xaa,0x42,0x66,0x9a,0x45,0xc3,0x93,0x4c,0xa0,
0x4f,0x32,0x7a,0x2d,0x40,0x75,0x7b,0x62,0xa6,0x22,0x11,0xf7,
0xb8,0xd2,0x74,0x5d,0x93,0xad,0x13,0x71,0xf8,0x8c,0xca,0x54,
0x75,0xd5,0x04,0xe7,0x69,0x00,0x38,0xb2,0x8c,0x38,0x87,0x82,
0x84,0xd9,0x38,0xbb,0xc0,0x6c,0x6c,0x2e,0x42,0xe2,0x29,0xc5,
0xe1,0x36,0xe2,0xd5,0xbc,0x90,0x84,0x57,0xc3,0x24,0xfe,0xf2,
0xd7,0x1b,0x1b,0x0f,0xe8,0xa5,0xc5,0x8a,0xb0,0xf8,0xe5,0x45,
0x7a,0xb2,0x19,0x32,0x1d,0xe5,0xc2,0xd1,0x94,0x36,0x33,0x70,
0x16,0xf8,0xfe,0xb4,0x7f,0x4f,0x89,0x61,0x97,0x9b,0x76,0xca,
0xb9,0xc4,0x1c,0xf6,0x59,0x79,0x9f,0x46,0x98,0xac,0xe5,0x1c,
0xdf,0x5e,0xfc,0x50,0x9f,0x99,0xab,0x4b,0x90,0xdb,0x2a,0xe1,
0xc6,0xbc,0xaf,0xf2,0x74,0x32,0x1e,0x9e,0x07,0x6e,0x76,0x92,
0x18,0xfa,0x45,0x75,0xcf,0x87,0x7f,0xc9,0x28,0x2f,0xa2,0x48,
0x61,0x81,0x42,0x40,0xb4,0x52,0xf2,0xc1,0xb9,0xf9,0x77,0x35,
0x1f,0xbb,0x8b,0xe0,0xdc,0x52,0xf9,0x54,0xac,0xa1,0x55,0x09,
0x5f,0x65,0x79,0x64,0xd0,0xa0,0xc4,0x9b,0x13,0x44,0xa9,0xfc,
0xa8,0x2e,0xf9,0x34,0x02,0x77,0xda,0xcf,0x19,0xf1,0xbd,0x19,
0x85,0xc7,0x35,0xfe,0xa7,0xc5,0x7c,0x96,0x4d,0x61,0xf5,0x99,
0x5a,0xc0,0xf5,0x08,0x1a,0x7c,0x2a,0x32,0x77,0x1b,0x20,0x7c,
0xef,0x43,0x8a,0x0c,0x31,0x9f,0x98,0x1b,0xd4,0xf4,0x8c,0x4d,
0xf4,0x3b,0xec,0xdd,0xf0,0x7d,0xb2,0x0b,0xc0,0x74,0xa7,0x42,
0x6d,0xfd,0x05,0x17,0xd5,0x39,0xfd,0xdf,0xf5,0x68,0x0d,0x3c,
0x8b,0x72,0xf1,0x27,0x89,0x3c,0xeb,0xac,0x2b,0x95,0xde,0x4c,
0x08,0xeb,0x19,0x18,0xe9,0xf1,0xdd,0x99,0x1e,0xa9,0x15,0x81,
0x5e,0x94,0x86,0xdb,0x1b,0xa1,0xcf,0xa2,0x17,0x9e,0xf6,0x11,
0x5f,0x00,0xc6,0xcc,0x00,0xec,0x33,0x82,0xb4,0xd9,0x83,0x97,
0x68,0x37,0x43,0xf6,0x91,0x94,0xfc,0x30,0x5b,0xec,0xe6,0xcc,
0xdb,0xf3,0xac,0x59,0x69,0xbd,0x1e,0xec,0xeb,0xdb,0x1b,0x1c,
0xe4,0x74,0x69,0xa2,0xe0,0x7f,0x94,0x21,0xbd,0xc1,0x93,0xaa,
0xbc,0xca,0x86,0x8c,0x89,0xcc,0xc0,0xf1,0x86,0x68,0x7e,0xce,
0x98,0x8f,0xa4,0x89,0x03,0xfa,0xc5,0xbc,0x8e,0x22,0x0f,0xea,
0x0c,0xfc,0x31,0x13,0x12,0x5e,0xa0,0xf0,0xad,0x4c,0x55,0xa2,
0x43,0x41,0xd7,0xc8,0xb0,0x22,0x1a,0x97,0x6d,0x62,0xda,0xb6,
0xf7,0x2f,0xcf,0x17,0xa8,0xb3,0x79,0x6f,0xc2,0x91,0xf0,0x13,
0x0c,0xe7,0x34,0x32,0x06,0xee,0x2f,0x5b,0x5d,0xc8,0xd0,0xdd,
0xb2,0x52,0x76,0xe4,0x9a,0xca,0x1e,0x70,0x82,0x9a,0x19,0x5c,
0x7f,0x9d,0x45,0xfa,0x46,0x3a,0x89,0x42,0xf2,0x85,0x9b,0xce,
0x74,0x35,0xac,0xfe,0x65,0x9f,0xc3,0x37,0x66,0x64,0xc9,0x3d,
0xbb,0x2d,0x65,0xdb,0xd2,0x09,0x9c,0x76,0xf1,0x31,0x65,0x5a,
0xcc,0xdf,0x97,0xa8,0xfd,0x09,0x00,0x85,0xc2,0xab,0x60,0x35,
0x59,0x96,0x60,0x7d,0x8c,0x82,0x28,0x8c,0xde,0x97,0xda,0xc8,
0xb2,0xf8,0x44,0xca,0x9c,0xaf,0x30,0x33,0x99,0xa4,0xac,0xfe,
0x92,0xd2,0x87,0x0c,0x9a,0x55,0x2c,0x0b,0xc7,0xca,0x3b,0x40,
0x62,0x5c,0x4d,0x44,0x72,0x13,0xb1,0xaa,0xc1,0x52,0x9f,0x34,
0x88,0xc2,0x6f,0x10,0x0d,0x9d,0x84,0x86,0xec,0xef,0x9b,0x3a,
0xdc,0x01,0xcf,0x42,0xcc,0x15,0xd7,0x20,0x7d,0xd1,0xc9,0x44,
0xf6,0x93,0x5e,0x5b,0x32,0x21,0x02,0x17,0x12,0xbd,0x45,0x72,
0x84,0x4e,0xe8,0xfd,0xd3,0xd1,0x03,0x22,0xea,0xc0,0x4f,0x0c,
0x0d,0x6b,0x39,0xab,0x83,0xc1,0x0b,0x7e,0x61,0xdf,0x54,0x4c,
0x86,0x10,0x91,0xda,0x7f,0x60,0x22,0xf0,0xcd,0x3a,0x9a,0xae,
0x36,0xbd,0x25,0x7a,0x40,0x6c,0x7c,0xa8,0xb9,0xcd,0xcc,0x3c,
0x45,0x1d,0x7d,0x3a,0x55,0xae,0x6e,0x2f,0x88,0x8e,0x50,0x10,
0x78,0xf4,0xb7,0x63,0x30,0x14,0x7e,0x40,0x2f,0xea,0x84,0x36,
0x7f,0x2b,0x40,0x4f,0x46,0xa4,0x82,0x31,0xf9,0xf6,0xd6,0x52,
0x6f,0xaa,0xa6,0xf6,0xcb,0x1a,0xae,0xf6,0x0d,0xbd,0x2b,0xfe,
0xaa,0xf6,0x68,0x7f,0x8e,0xb5,0x48,0xda,0xb2,0x26,0xde,0xb2,
0x8a,0xcf,0xef,0x18,0x6f,0x70,0xb5,0xbb,0x44,0x5c,0x82,0x5e,
0x43,0x20,0xb9,0x3b,0xf1,0xcc,0x85,0xe1,0xdb,0x33,0x08,0xde,
0xf5,0x1c,0x8c,0x48,0xf6,0x14,0x9d,0x9b,0x52,0x6c,0x06,0x4d,
0xcf,0x80,0xa4,0xcb,0x33,0x37,0xb9,0x54,0x52,0xb4,0xe4,0xa0,
0x56,0xb4,0x8e,0x25,0x79,0x45,0x54,0x8a,0x91,0xb8,0xc1,0x00,
0x0b,0xc6,0x8a,0xae,0x95,0x91,0xe0,0xa6,0xa7,0xb0,0xf7,0x8c,
0xeb,0x00,0x0c,0xce,0x95,0xef,0xbf,0xbf,0xff,0xab,0xd7,0x25,
0xda,0x57,0x8d,0xdf,0x23,0x55,0x52,0xf5,0x07,0x59,0x51,0xc1,
0x8b,0xb1,0x7b,0xb7,0xbb,0xce,0x72,0x63,0x41,0xec,0xf1,0xde,
0x7a,0x2f,0xf9,0xfe,0xcb,0x9d,0xf4,0x22,0xed,0x0f,0x50,0xb8,
0xab,0xbf,0xdc,0xd3,0xcb,0xe4,0x89,0x81,0x5a,0x2f,0x9e,0xc5,
0xed,0x3c,0xff,0xbf,0x00,0x00,0x00,0xff,0xff,0x2c,0x7b,0x15,
0x75,0x53,0x17,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}