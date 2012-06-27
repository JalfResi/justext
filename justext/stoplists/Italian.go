package gojustext
import ( "io"; "os"; "bytes"; "compress/gzip" )

func ItalianStoplist() ([]byte, os.Error) {
var gz *gzip.Decompressor
var err os.Error
if gz, err = gzip.NewReader(bytes.NewBuffer([]byte{
0x1f,0x8b,0x08,0x00,0x00,0x00,0x00,0x00,0x00,0xff,0x54,0x97,
0x4d,0x72,0xdc,0x46,0x12,0x85,0xf7,0x75,0x0c,0x6d,0x86,0x0b,
0x85,0xee,0xc0,0x21,0x43,0x1a,0x4e,0x70,0x18,0x94,0x28,0x1d,
0xa0,0x1a,0xa8,0x06,0x33,0x06,0x40,0x41,0x05,0x14,0x1c,0xd1,
0x27,0xf0,0x31,0xbc,0xb4,0x23,0xbc,0xf2,0xce,0x6b,0x5e,0xcc,
0xdf,0xcb,0x42,0xcb,0xe1,0x0d,0xd9,0x3f,0xa8,0xac,0xcc,0xf7,
0x5e,0xbe,0xcc,0xee,0x2d,0xa4,0x60,0x63,0x18,0x63,0xb0,0x39,
0xf4,0x69,0x0c,0x31,0x74,0xaf,0x29,0xbc,0xfd,0x1a,0xaa,0xbf,
0xe7,0x8b,0x25,0x15,0xde,0xf0,0x79,0x9e,0xc3,0xcc,0x13,0x7d,
0x0c,0xab,0x05,0x0b,0x91,0x63,0x89,0x67,0x2c,0x3c,0x8c,0xe1,
0x31,0xf2,0x9e,0x87,0x67,0x1e,0xea,0xf2,0x94,0xc2,0x62,0x6f,
0x7f,0x7a,0x80,0xa4,0x43,0x7c,0x13,0x67,0x05,0xee,0x39,0xb5,
0xe6,0x39,0x87,0xd4,0x87,0xd8,0x87,0xd7,0x18,0xce,0x35,0x0c,
0xa3,0x85,0xb5,0xc6,0xf0,0x44,0xf4,0xb5,0xf2,0x5d,0x89,0x61,
0xe2,0xc2,0x6a,0x21,0xeb,0x04,0xa7,0xfb,0x9a,0xc2,0x98,0xc3,
0xc6,0x37,0x4b,0x2c,0x9b,0xee,0xf5,0x53,0x5b,0xdc,0x72,0x58,
0x8a,0xf1,0x78,0x9c,0x67,0xaa,0x59,0xd7,0x54,0x52,0x78,0x20,
0x4e,0x78,0x98,0x43,0xb4,0xd0,0xe7,0x25,0x73,0xb4,0xe4,0xf0,
0xbd,0xa6,0x95,0xa7,0x77,0x4b,0x73,0xa2,0x22,0x5e,0xa5,0x99,
0x57,0x6b,0xa5,0x6a,0xe5,0xf9,0x48,0xd4,0xbc,0x2b,0x5f,0xf3,
0x90,0x99,0x4c,0xb9,0xb3,0xb3,0x6d,0x7b,0xfb,0xc5,0xab,0xf0,
0x87,0xbd,0x4a,0x2a,0x7c,0xfb,0x39,0x2c,0xd9,0xc2,0x94,0xe6,
0x8d,0x1b,0xa7,0x3c,0x12,0xfb,0x7b,0x8d,0x3c,0x04,0x00,0x95,
0xc0,0x67,0xe3,0x0e,0xea,0x21,0x4d,0x03,0x33,0x15,0xdf,0x32,
0xa6,0x1c,0x82,0x65,0x6e,0xdd,0x8a,0xe9,0xcc,0xdc,0x67,0xa0,
0x98,0x79,0xdc,0x73,0x8c,0xe1,0x56,0x18,0xf0,0xb6,0xaf,0xfc,
0xa3,0xda,0x97,0x38,0x73,0x92,0xe2,0x28,0xb6,0xbe,0xfd,0x01,
0x0e,0xa4,0x92,0x60,0x84,0x83,0x71,0x4f,0xbb,0xb3,0xb3,0xb5,
0x1a,0x1d,0x2e,0x7d,0x5d,0x2c,0xb5,0xf7,0x19,0x90,0x8b,0x40,
0x27,0xa3,0xc4,0xad,0x5d,0x05,0x29,0xa5,0x4a,0x3a,0x69,0xa8,
0x46,0xe2,0x03,0xf7,0x70,0xe8,0x6c,0xe3,0x14,0xee,0x85,0xd8,
0x8b,0x85,0xad,0x6e,0x1b,0x34,0xcf,0x5d,0x06,0xf5,0x2e,0xaf,
0x6f,0xbf,0x87,0x9d,0x2a,0x63,0x78,0x46,0x11,0x43,0xa9,0x0b,
0x8f,0xad,0x69,0x5a,0x94,0x0b,0xf4,0x88,0x28,0x81,0x27,0x5e,
0xbe,0x57,0x9b,0xd1,0x56,0x1e,0xb8,0xa8,0xb3,0xd2,0xc1,0xce,
0xc9,0x36,0x4a,0xe1,0xc6,0xbc,0x89,0x84,0x54,0xa8,0x65,0xe3,
0x34,0x15,0x91,0x69,0xa6,0x90,0x33,0xd7,0xac,0x69,0xbe,0xe8,
0xaf,0x4a,0x8b,0x61,0xa8,0xa9,0xf0,0xa1,0x70,0x42,0x51,0x25,
0xc3,0x0a,0xa9,0x1a,0x6c,0x08,0x34,0xc2,0x50,0xce,0x5c,0xf3,
0xce,0xd9,0xe8,0x51,0xb9,0x03,0x26,0x88,0x45,0x9e,0x1c,0x89,
0xca,0xe5,0x1c,0x27,0xe3,0x7f,0xf4,0x7a,0x1c,0xae,0x02,0x7e,
0x19,0xcc,0x4a,0x1a,0x2c,0x13,0xa2,0x8b,0x75,0x85,0xd4,0x3a,
0xa5,0xab,0x4a,0x44,0xd9,0x56,0x5d,0x5b,0xa9,0x80,0x37,0x35,
0x0a,0xfb,0x7c,0x0e,0x9f,0x1c,0xb4,0x4e,0xc2,0x80,0xa8,0x76,
0x5e,0xdc,0x37,0xc9,0xa8,0x2b,0x40,0x17,0x20,0x73,0x41,0x92,
0x3b,0xb7,0x52,0xd5,0xdc,0x25,0xd8,0x5f,0x0d,0x06,0xf9,0x50,
0x0a,0x50,0x56,0x7c,0x4b,0xd3,0xad,0x96,0xd0,0x12,0xca,0x87,
0x35,0xd4,0x59,0x94,0x27,0x84,0x10,0x9e,0x9b,0xbb,0xd7,0xb7,
0xdf,0xe8,0x47,0x4b,0xa4,0xb7,0x2e,0x9e,0x02,0x44,0x5d,0xe0,
0xd5,0xe6,0x3d,0x75,0x60,0x92,0xa5,0x68,0xfe,0x7a,0x4f,0x9d,
0x22,0xd7,0x14,0xe3,0x49,0x15,0x0a,0x44,0x9b,0x24,0x91,0xde,
0x87,0x95,0xd7,0xbd,0x48,0x94,0x56,0x51,0xf4,0x2b,0x9f,0x09,
0x0d,0x02,0x54,0xc1,0xb8,0xa5,0x32,0x49,0x1a,0xfc,0x2f,0x5c,
0x2e,0x94,0x47,0x9a,0x8f,0xfe,0x18,0x14,0x57,0x10,0xb7,0xbe,
0x33,0x0a,0x8f,0x00,0x3b,0x56,0xbe,0x01,0x37,0x12,0x5a,0x61,
0x85,0x8b,0x1c,0xdb,0x29,0xab,0x35,0xbf,0xc9,0x02,0x44,0xef,
0xd4,0xf4,0xd9,0x74,0x1d,0x31,0x8d,0x17,0x2f,0x9d,0x63,0x28,
0xd4,0xb9,0xa1,0x43,0x9c,0xa8,0x3b,0x6c,0x63,0x53,0xeb,0xd8,
0x2c,0x26,0x65,0x07,0x12,0x13,0x8a,0xca,0xe1,0x73,0x6b,0x5b,
0xe8,0xdf,0x6d,0xee,0xc0,0x47,0xda,0x71,0xdc,0x4b,0x15,0x4e,
0x6a,0x4c,0xee,0x8a,0x03,0x64,0x48,0x26,0xb9,0xcb,0xe1,0x3f,
0x31,0xdc,0x60,0x68,0x1f,0x6b,0x10,0x06,0xa2,0x9c,0x74,0xee,
0xf1,0x9e,0xde,0xae,0x4d,0xdf,0x41,0x0a,0x1f,0xc6,0xf1,0x54,
0x27,0x35,0x86,0x6e,0xd7,0xb7,0xc4,0x26,0x18,0x4e,0x72,0x3a,
0xa9,0x89,0x68,0x8a,0xdb,0x11,0x4a,0xd0,0xe9,0xfd,0xd1,0x8c,
0x0b,0xfd,0x31,0xc2,0x82,0x52,0x38,0x4b,0x5a,0x9e,0x5c,0x5f,
0xdb,0x27,0x4b,0x3d,0x9d,0x46,0xeb,0xa4,0x9d,0x07,0x4a,0x22,
0xdf,0x4f,0x06,0x7c,0x32,0xa7,0x4c,0x8a,0xe1,0x11,0x9e,0x12,
0xd5,0x62,0x86,0x2e,0x6a,0xa4,0x31,0xa3,0x5a,0x6f,0x40,0x35,
0x0e,0xc5,0x26,0xbd,0x03,0x39,0xef,0x78,0x75,0x03,0x4c,0x4d,
0x27,0x9e,0xee,0x94,0x9a,0xea,0x04,0x63,0x70,0x00,0xbb,0x84,
0xe3,0xd1,0x44,0xde,0x16,0x75,0x90,0xdf,0x2c,0xa9,0x43,0x1f,
0x2f,0x2e,0xa5,0x0f,0xe1,0xdf,0x96,0x89,0x75,0xd6,0x4b,0x2f,
0x1c,0xf6,0xd5,0x4d,0xa2,0x8d,0xc7,0xe4,0x6b,0xff,0x8b,0xd2,
0xdc,0x57,0xd4,0xd1,0xaa,0x87,0x5c,0xfe,0x0b,0x52,0xb8,0x86,
0x2b,0x75,0xe8,0x25,0x73,0x9d,0x5b,0x09,0xa9,0x11,0x26,0xaa,
0x64,0xc9,0x41,0xc6,0x97,0xe4,0x53,0xe6,0x6f,0xe5,0x85,0xb2,
0xbc,0x6b,0x9a,0x05,0xa1,0x1c,0xf8,0xa8,0x4c,0x3a,0xbb,0x35,
0x31,0x4a,0x24,0x2e,0x70,0x09,0xa5,0xc9,0xd3,0x58,0x91,0x95,
0x10,0x9f,0xb0,0x56,0x05,0x71,0xd6,0xe9,0x9e,0x58,0xd4,0x58,
0xf0,0xbe,0x88,0xa0,0x11,0x95,0xf2,0x21,0x7a,0x4f,0x21,0xd3,
0x20,0xee,0x8c,0x3c,0xd6,0x45,0x84,0xab,0xb1,0x41,0x2b,0x58,
0xe3,0x41,0xe9,0x92,0x64,0x2f,0x1b,0x18,0xc7,0x7f,0x19,0xcc,
0x15,0x32,0x9d,0x62,0xb9,0x28,0xdf,0xae,0x21,0x1a,0xf7,0x76,
0xed,0x59,0xdd,0x44,0x69,0xc4,0xe0,0x9a,0x77,0x8c,0x33,0x35,
0x95,0x9d,0x8c,0x77,0xd2,0x76,0x64,0x44,0xee,0x1a,0x22,0x0f,
0x0f,0xaa,0x61,0xf0,0x0e,0x3b,0xc8,0xe3,0x20,0x63,0x43,0xc1,
0x25,0x5d,0x3c,0x88,0x72,0x24,0xb9,0x23,0x60,0x0a,0x3b,0x10,
0x37,0xf8,0x8c,0xd1,0xb1,0xb7,0x9b,0x29,0x65,0x60,0xe6,0x19,
0xf8,0x16,0x39,0x23,0x4f,0x03,0x18,0x41,0xee,0xb9,0xab,0x66,
0xda,0xab,0x77,0x74,0x9c,0xe4,0xd9,0xe5,0xa6,0xd4,0xd4,0x12,
0xe2,0x2a,0x31,0xa8,0xde,0x31,0x78,0x85,0x42,0x0c,0x5f,0x32,
0xc6,0x72,0xa4,0x63,0xe1,0x2e,0x96,0x66,0x4f,0x9c,0xc4,0xb7,
0xb7,0x72,0x48,0xb3,0x87,0x2f,0xe5,0xfd,0xee,0x2b,0xe2,0xbb,
0x6b,0xb6,0x72,0xeb,0x63,0xb9,0x33,0x4c,0x0e,0xf3,0x44,0xae,
0x1a,0x13,0x40,0x52,0xc4,0x31,0x6d,0x0c,0xfb,0x4f,0xe9,0x27,
0x82,0x09,0x09,0x1f,0x6c,0x14,0x82,0xd3,0xcc,0xfd,0x61,0x98,
0xb4,0x0b,0x1d,0x26,0x84,0x9a,0x65,0xa2,0x25,0x02,0xca,0x0c,
0xbd,0x22,0x94,0x4f,0x1b,0xed,0x36,0xfa,0xb8,0x78,0xa2,0xcf,
0xe5,0xff,0x0e,0x3d,0x2d,0x81,0x35,0xa5,0xb2,0x8b,0x31,0x07,
0x49,0x7c,0x30,0xcc,0x04,0x02,0x66,0x71,0xcc,0xb7,0x5b,0x59,
0xe3,0xd9,0xb8,0x4c,0x9a,0xc3,0x33,0xae,0x88,0xc9,0xa4,0xb1,
0x21,0x8d,0x37,0x79,0x92,0x06,0x9a,0x6a,0x59,0x2b,0xa3,0x06,
0x2e,0x48,0x41,0x29,0xbf,0x5a,0x9c,0x04,0x64,0x17,0x67,0x64,
0x8c,0xc0,0xea,0x4a,0x77,0x52,0x50,0xed,0x5d,0xc5,0xb2,0xe6,
0xcf,0x6d,0x04,0x1f,0x2b,0x45,0x81,0x6f,0x5a,0x9d,0xe9,0xe1,
0xee,0xda,0x99,0x97,0x0c,0xce,0x98,0xac,0x7b,0x24,0x02,0xd2,
0xfd,0x0c,0x49,0x9f,0xb1,0x77,0xbe,0xef,0x28,0x27,0x70,0x01,
0x08,0x1f,0x36,0xea,0x67,0xc8,0x2e,0x3d,0x0d,0xc6,0x32,0xa4,
0x9a,0x56,0xf6,0x8a,0xc3,0x77,0x6c,0x8f,0x92,0x39,0xc9,0x2a,
0x8d,0xc6,0x3c,0xc6,0x28,0x78,0x5d,0x3e,0x34,0xef,0x20,0x43,
0xfe,0x21,0x7d,0xc6,0x4b,0x1e,0x0d,0xff,0xd5,0xd4,0xa2,0x83,
0x79,0xe6,0x1b,0x8f,0x12,0x30,0x16,0x5e,0xcb,0x7f,0x7c,0xae,
0xfe,0x37,0xbf,0xb2,0xf5,0x40,0xce,0xc7,0x36,0x76,0xf0,0x42,
0xd8,0x99,0xcc,0xaf,0xc2,0x54,0xb4,0xe7,0x0d,0xa3,0x6a,0xae,
0xcc,0x2c,0xe6,0x15,0x27,0x4f,0xcc,0xe6,0xb1,0xa5,0x13,0xf7,
0xea,0x50,0x41,0xd5,0x97,0x74,0x75,0x32,0xf7,0x71,0x93,0xf8,
0x16,0xeb,0x7c,0xf2,0x79,0x47,0x68,0xb8,0x79,0x53,0x30,0x28,
0x18,0x6d,0x11,0x1c,0x96,0x45,0x53,0x8d,0x3c,0xd4,0xe2,0x92,
0xa3,0xf5,0x1e,0x76,0xc9,0xa0,0x26,0x66,0xfa,0xf0,0xf4,0x43,
0xc9,0x1c,0x39,0x1c,0x33,0x43,0x92,0x26,0x26,0x4f,0x7e,0x49,
0x32,0x31,0x66,0xe6,0xe6,0x16,0x54,0xe8,0x53,0x6d,0x0f,0x13,
0xa1,0x84,0x12,0xe5,0x74,0x96,0xd9,0x4e,0x0f,0xed,0x5d,0x69,
0x69,0xa6,0xd8,0xe6,0xf2,0x7b,0x57,0xd6,0x87,0x63,0x23,0xf2,
0xed,0x27,0x97,0xcb,0x31,0x92,0xdb,0x8d,0x4c,0xb3,0xd1,0x4e,
0x74,0xf3,0xb9,0x2a,0x7d,0x9b,0x51,0x20,0x0d,0xd0,0x21,0x20,
0x69,0x22,0x9b,0x0f,0x5e,0xb8,0x5d,0x71,0x85,0xb5,0x63,0xcc,
0x88,0x50,0x14,0x6c,0xa1,0xd2,0x08,0x76,0xb9,0x48,0x51,0x73,
0xe6,0x0f,0xc6,0x39,0x61,0xae,0x5a,0x7e,0x72,0xf8,0x24,0xd2,
0x69,0x5b,0x85,0x27,0x67,0xfa,0xdc,0xc2,0xbd,0x2c,0x52,0x7b,
0xde,0x8d,0xc4,0x89,0xf0,0x69,0x5c,0x39,0x08,0xe7,0xda,0x32,
0x10,0x26,0x22,0x66,0xe8,0x3f,0x6b,0x7c,0xff,0x7f,0x02,0x41,
0x24,0xa7,0x65,0xcd,0x85,0xd4,0x83,0x6d,0x6b,0xdb,0x6e,0x8c,
0xf2,0x3d,0x1f,0xe2,0x02,0x90,0x5d,0x1b,0x81,0xba,0x75,0xd4,
0xcc,0xfc,0xa6,0x90,0xd6,0xc1,0x68,0x44,0x53,0x9c,0x18,0x1a,
0xef,0xd6,0xb4,0x40,0x28,0x19,0xc7,0x73,0x14,0x81,0x04,0xf2,
0xbd,0x63,0xac,0x14,0xe8,0xe5,0x6b,0xe3,0x96,0x57,0x9e,0x59,
0x14,0xb5,0x5e,0x9b,0xa4,0x33,0xba,0x62,0x81,0xf8,0x4c,0x4a,
0x7c,0x57,0xe7,0x7f,0xcc,0x38,0xb9,0x13,0x39,0xb6,0x1d,0x4c,
0xd6,0x46,0x31,0x4d,0x22,0xbe,0x85,0x1d,0x3b,0xae,0x0f,0x08,
0x2a,0x6e,0xf3,0x76,0xed,0xc8,0xad,0x8d,0x73,0xc4,0xed,0xd6,
0x2f,0x5e,0x8c,0x29,0x59,0xd7,0xb4,0x2c,0x22,0x07,0x87,0x13,
0x52,0xea,0x99,0x1b,0x6d,0x2b,0xde,0xc1,0xbe,0xe2,0x4a,0x50,
0x68,0x43,0x9e,0xa4,0x5e,0x5d,0x4c,0x6b,0x87,0x2f,0xf5,0x69,
0x64,0xa0,0x29,0xda,0x98,0x86,0xa1,0x65,0xc4,0x70,0xd0,0xf6,
0xc2,0xd6,0xa0,0x15,0x47,0x1b,0x5f,0xb3,0x68,0x7c,0xa2,0xa9,
0xfd,0x6a,0xf1,0x7e,0x2e,0x8d,0x00,0x22,0xb4,0xfa,0xc3,0xdb,
0xf3,0x89,0x4f,0xbc,0x29,0x97,0x31,0x6d,0x47,0xbf,0x6a,0x83,
0xd1,0xc8,0xba,0xe0,0x01,0xfc,0x0a,0x60,0x6d,0xe0,0xad,0x2e,
0xf0,0xd9,0x22,0xc6,0x73,0x78,0xac,0x34,0xaf,0x56,0xe0,0x90,
0x77,0xc6,0x00,0x35,0xd3,0x68,0xf0,0x25,0x19,0x0a,0x18,0x1e,
0xc6,0x46,0x00,0x8b,0xcf,0x9e,0x81,0x80,0x27,0x24,0x81,0x4d,
0x26,0x86,0x45,0x3b,0x3c,0xa0,0x4f,0x0f,0x4e,0xc0,0x7a,0xcc,
0x54,0x97,0xf6,0x07,0x69,0x44,0x4c,0x16,0xc0,0xea,0xd8,0xc7,
0x99,0x7b,0x14,0x84,0x0b,0x49,0x91,0xfe,0x53,0xed,0x68,0x86,
0xc0,0x46,0xcb,0xb1,0x7a,0x16,0x73,0xde,0xba,0x1a,0x0e,0x7f,
0xdb,0x8f,0x5b,0xf1,0x55,0x3d,0x4f,0xd0,0x8c,0xeb,0x09,0xfc,
0x5b,0x2e,0x9b,0xe5,0xfc,0x71,0x89,0x0d,0xcf,0x1b,0x7e,0x1e,
0x16,0xc6,0x0b,0xee,0xa9,0x85,0xd9,0x7f,0xba,0x1d,0x23,0x16,
0x03,0x25,0x9a,0x85,0x67,0xff,0xfd,0xf5,0xc0,0x68,0xd4,0x02,
0xc9,0x02,0xab,0x8d,0x84,0x4d,0x44,0xdd,0xc6,0x4f,0x2a,0xf6,
0x85,0xc3,0xda,0xf8,0x1d,0x46,0x8f,0x79,0xd8,0x66,0xc3,0x24,
0xf6,0x57,0x00,0x00,0x00,0xff,0xff,0x1c,0x76,0xb9,0x1c,0x83,
0x0e,0x00,0x00,
})); err != nil {
	return nil, err
}

var b bytes.Buffer
io.Copy(&b, gz)
gz.Close()
return b.Bytes(), nil}