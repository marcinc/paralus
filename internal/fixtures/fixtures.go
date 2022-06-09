// Code generated by vfsgen; DO NOT EDIT.

package fixtures

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Fixtures statically implements the virtual filesystem provided to vfsgen.
var Fixtures = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2022, 6, 9, 10, 25, 17, 658768698, time.UTC),
		},
		"/agent_templates.yaml": &vfsgen۰CompressedFileInfo{
			name:             "agent_templates.yaml",
			modTime:          time.Date(2022, 6, 9, 10, 25, 17, 658768698, time.UTC),
			uncompressedSize: 4367,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xe4\x57\xc1\x8e\x9b\x30\x10\xbd\xf7\x2b\x46\xdc\x43\xe9\xa9\x12\xb7\x88\x54\x6a\x0f\xdb\xae\xb2\xdb\x0f\x70\xcd\x40\x2d\x19\x1b\xd9\x43\xb4\x68\xb5\xff\x5e\x19\x4c\x0a\x0d\x44\xa0\x94\x4d\x56\x9b\x1b\xd8\x33\xcc\xcc\x7b\xcf\x79\x16\x84\x85\x8d\x3f\x00\x6c\xa0\x40\x62\x29\x23\xe6\x9e\xdc\x4f\xb1\x02\x63\x28\x99\x61\xb2\xb2\x1b\x8b\x8a\x4c\xbd\x29\x11\x8d\x50\xf9\xc6\xa2\x39\xa0\xf1\x3b\x53\x61\x4b\xc9\xea\xef\x4d\x40\x70\xdf\x46\xc0\x43\x13\x01\xf7\x6d\x04\x3c\x34\x11\x41\x13\x62\x4b\xe4\xdd\x67\x84\xca\x0c\xdb\x63\x36\xf5\x29\xbf\x8d\x55\xa4\xf7\x98\x0b\x4b\x68\x62\x20\x53\x61\x6f\x61\x5b\x96\x46\x1f\x70\xf0\x5e\xe4\x4a\x1b\xbc\xab\x24\x89\x52\xe2\x68\x28\x61\x51\x4a\x46\xf8\x58\x97\xae\xf0\x7e\x85\x00\xbf\xb5\x25\x1b\x83\x7f\x72\x03\x72\x6f\x62\x08\x9e\x9f\x21\x6c\x4b\xf4\xad\x7d\xd5\x96\xe0\xe5\x25\x38\x6e\x05\xa0\x36\xa3\x5b\x71\xc9\xbf\x3c\x11\x1a\xc5\x64\x70\x9a\x6d\xd8\x74\xfc\x39\x8a\x3e\x9d\x49\xf4\x4d\xcd\x4c\x14\x1e\x1f\x6b\x4b\x58\x34\x79\x61\x85\xc4\xa1\x3d\xf0\x90\xcb\xca\xcd\x36\x94\x9a\x33\xb9\xb4\x05\xa1\x92\x36\xfc\xd1\xa3\x11\x43\xd0\xad\xe9\x8a\x7e\x64\xe3\xcb\x0b\xf9\xca\xa5\x40\x45\x37\x4c\xbe\xa4\x29\x70\xe5\xa1\x70\x6d\x70\x63\x50\xb2\x7a\x8e\x80\x13\x6d\x10\xf6\x6e\x37\x0c\xa5\x21\xd9\x2f\x94\x36\x3e\x62\xec\xd3\x87\x29\x1e\x3e\x72\xad\x14\x72\xd2\xfd\xd2\x5c\xd7\x73\x84\xff\xb7\xbc\xeb\x8a\x7e\x5c\xf3\xae\xba\x66\x1a\x6d\xc8\x72\xd5\xaf\x0f\x2a\xcb\x3b\x92\x9f\x07\x29\xc5\x8c\x55\x92\x9a\x6e\xfe\x1f\x3e\x19\x93\x76\x6d\x61\x9c\x3d\x95\xdf\x00\x42\x95\x9d\x23\xba\x9f\x16\xcd\x02\xcd\x79\x38\x5d\xd4\x00\xcd\x7f\xf7\xb9\x8f\x8f\xca\x72\x52\xc4\x5b\x47\xa8\x5e\xc8\x14\xe7\x96\xb1\xa7\x3f\x85\x57\x90\xf8\x9d\x78\xc2\x74\x19\x81\xdc\x28\x6f\x8b\x3e\x3c\x5d\xc4\xa0\x64\xe7\x4f\x6e\xd7\xc9\x02\x0e\x25\xbb\x13\x16\x35\xa8\xc2\x79\x58\xc7\xaa\xbb\x0e\xb2\xd3\xc0\x26\xbb\x0b\xa0\x1d\xf1\x43\x3c\x6d\xb9\xef\x0c\x4f\x74\x99\xb5\xea\x52\x8d\xb8\xb6\x08\x56\x49\x3d\xee\xdb\xa2\xeb\xfb\xb6\x21\x9d\xe6\xba\x94\x23\xdd\x4f\x6e\x19\x33\x89\x7b\xab\x86\xc3\x93\x36\xe9\x0e\xe4\x1b\x3d\x94\xfa\xc6\x63\x0e\x4c\xdb\xbc\xf7\x9f\x3e\xeb\x58\x9a\xf0\x2a\x97\xc0\xfb\x7a\x7e\xe5\x0d\xe2\xeb\xef\x41\x3c\x7d\x57\x57\xa8\x3f\x01\x00\x00\xff\xff\x2f\xb6\x4c\xa4\x0f\x11\x00\x00"),
		},
		"/relay_agent_template.yaml": &vfsgen۰CompressedFileInfo{
			name:             "relay_agent_template.yaml",
			modTime:          time.Date(2022, 6, 9, 10, 25, 17, 658768698, time.UTC),
			uncompressedSize: 3904,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xd4\x57\x5f\x6f\xe2\x46\x10\x7f\xf7\xa7\x58\xf1\x72\x2f\x35\x24\x97\xde\xe9\x64\xe9\x1e\x7c\x90\x5c\x50\xf9\x63\x01\x69\xef\x54\x55\x68\xb2\x3b\xc0\x2a\xeb\x5d\x6b\x77\x4c\x62\x45\xf9\xee\xd5\xda\x18\x4c\x42\xda\x34\x47\x5a\xd5\x4f\x78\xfe\xfe\xe6\xb7\x83\x77\x06\x32\xf9\x2b\x5a\x27\x8d\x8e\x18\x64\x99\xeb\xac\x4f\x83\x1b\xa9\x45\xc4\x7a\x98\x29\x53\xa4\xa8\x29\x48\x91\x40\x00\x41\x14\x30\xa6\xe0\x1a\x95\xf3\xbf\x18\xb3\x98\x85\x20\x84\xf7\x5d\xbf\x0f\x2d\x2a\x28\x42\x58\x7a\x8f\x5a\xcb\x55\xee\x08\x6d\xc4\x5a\xf7\xf7\xac\xdd\x33\xb7\x5a\x19\x10\x3d\x20\x68\x77\x2b\x55\xbf\xc7\x1e\x1e\x5a\x01\x63\x1a\x52\x8c\x58\xeb\x26\xbf\xc6\x30\xb3\xe6\x6e\x13\x2a\x7c\x91\xa3\xcb\x80\x63\xc4\x32\xb0\xa0\x72\x17\xba\xc2\x11\xa6\x81\xcb\x90\x7b\xa4\x99\x35\x4b\x8b\xce\xf5\x10\x84\x92\x1a\xa7\xc8\x8d\x16\x2e\x62\xa7\x9f\x4e\x4e\x82\x12\xa9\x92\x1c\xbc\xa0\x7c\x5b\x4b\x4f\xc8\xa5\x74\x64\x6c\x31\x90\xa9\xa4\x88\x9d\x7a\x43\x87\x0a\x39\x19\x5b\x95\x9f\x02\xf1\xd5\xa0\xc1\x07\xf3\x14\xbe\xaa\x06\xc2\x34\x53\x40\xb8\x09\xdc\xe0\xdb\x3f\x6a\x2f\xc7\xeb\xb3\x30\x56\x33\xe2\x1f\x6e\x34\x81\xd4\x68\xb7\x91\x43\x06\x76\xd9\xc8\x13\xb2\x30\x4c\x8d\xc0\xcf\x5c\xc9\xfa\x54\x6b\xb9\x32\xcb\x50\xe1\x1a\xd5\xe7\xb3\xad\x1c\xf5\xba\xe9\x5c\x9d\x68\x32\xee\xcd\x47\xf1\xf0\x7c\xab\x60\x6c\x0d\x2a\xc7\x0b\x6b\xd2\xa8\x21\x64\x6c\x21\x51\x89\x09\x2e\xf6\xa5\xbe\xdc\x5d\x8f\xae\x4f\x1f\x29\x4b\xa7\x04\x68\x15\x6d\x69\x6b\xfb\xc4\xcf\xe2\x98\x26\x71\xf7\xdf\x06\x53\xb6\xe7\x13\x44\xc3\xf8\xdb\xbc\xd7\x8f\x07\xd3\xbf\x47\xc3\x8d\x5e\xc8\xe5\x10\xb2\x5f\xb0\x38\x00\xea\x06\x8b\x88\xa5\x70\xd7\x93\xa0\xdc\x23\xdd\x73\xff\xab\x2a\xe4\x4b\x9a\x66\xf7\x98\x8c\xa4\xd1\xa0\x22\x46\x36\x7f\x5a\xd0\xe4\x7c\x10\x7f\x9f\x77\xc7\xa3\x8b\xfe\xd7\x61\x9c\x1c\x3c\xf6\x1f\x87\x52\x67\xf3\xd4\x8d\xaf\x66\xf3\x64\x32\xfe\xf6\xfd\x38\x14\xae\x88\x32\x97\x78\x64\x07\x49\xac\x30\x57\x81\xfe\x19\x33\x7b\x58\xe7\xf1\xd5\xec\xf2\x7c\x34\xeb\x77\xe3\x59\x7f\x3c\x3a\x0e\xf4\x12\x5b\x9c\xd3\xea\xc8\xc8\x2f\x67\xb3\xe4\xd8\x14\xbf\x05\xc3\x1e\xe7\xf4\x7f\xd1\x0b\xa3\xf1\x31\x61\x6a\xf3\x16\x18\xe3\xc1\x60\xfc\xdb\xbc\x3f\x9a\x9e\x77\xaf\x26\xe7\xf3\x2f\xe3\xf1\x6c\x3a\x9b\xc4\xc9\x71\x30\x83\x52\xe6\xb6\xaf\x1d\xf2\xdc\xe2\x17\x63\xc8\x91\x85\xec\x48\x25\xc8\x14\x96\x78\x68\xe0\x98\xf8\xf9\xa4\xef\xb5\x7b\x9f\x93\xd2\x3e\xc9\x95\x4a\x8c\x92\xbc\x88\x58\x7f\x31\x32\x94\x58\x74\xcd\x2b\xef\xd5\xd3\x49\xf5\x10\xda\x54\x6a\xf0\x80\x87\xe8\x9c\xcf\x58\xde\x12\x1d\x81\xeb\x4e\x43\xe9\xaf\xd5\xbf\x72\xda\x40\xbc\x90\x6a\x57\xf0\xda\xa8\x3c\xc5\xa1\xc9\x35\xed\xdd\xdd\xa9\x97\x6c\xd2\x20\xf1\xce\x13\x12\x8f\x71\x33\x48\x2d\xa9\x7b\x60\x90\xe0\x26\x4d\x41\x8b\x26\x1e\xb7\x6a\x0e\x10\xbc\xf1\x92\x2b\x3f\x63\xb1\x50\xb3\x8f\x1f\x3e\x9c\x7d\x7c\x7c\x94\x16\x97\xd2\x91\x2d\xda\x16\x16\x50\x84\x28\x96\xd8\xd6\x48\x9d\xf2\xb5\x73\x9d\xbb\xe2\xda\xdc\x45\xa7\xed\xb3\xb3\x57\x1d\xaa\x43\x0a\x4b\x00\xbb\x9b\xd3\xa2\x33\xb9\xe5\xe8\x22\x76\xff\xb0\x95\x96\xfd\x2a\xa9\xf0\xf5\xe2\x1d\x35\x5b\x3b\xb3\x72\x2d\x15\x2e\x51\x3c\xea\xc6\xb7\x3b\xf8\xcc\x4a\x53\xa2\x51\xe0\xdc\xa8\xfa\xb3\x6c\xc6\xdf\xcd\xd4\x1d\x72\x2b\x49\x72\x50\xc1\xb6\x2a\x02\x4b\x75\xac\x58\xdd\x42\x51\xd7\xec\xf8\x0a\x45\xae\xd0\x56\x91\x04\x2e\x20\x57\x14\x6e\xc5\xc1\x61\x0a\x76\xf4\x34\x30\x7f\xb5\xc0\x31\x41\x2b\x8d\xd8\x4d\xda\x27\x41\xb3\x5b\xf7\x3a\x65\xf3\xc1\x68\xf2\xb9\xc9\x3f\x34\x02\x23\xf6\xf3\xfb\x93\xe0\xd8\xf3\xcc\x8f\x45\x09\xc2\x30\x0c\x1e\x4d\x82\xd5\xd6\xd4\xad\xab\x09\xea\x01\x9e\xd7\x9e\x2f\x59\x82\xea\x19\x2e\x62\xad\x4f\xad\x72\x17\x51\x50\xb8\x88\xbd\xfb\xfd\xbe\x45\xe6\x06\x75\x2b\x7a\x1a\x64\xe6\x15\x3e\xc0\x4f\x2d\x10\xc2\x1e\x32\x99\xa2\x26\x5b\xc4\x42\xd8\xca\x0e\xb5\xc8\x8c\xd4\x74\xc8\xf6\xd2\x38\xf2\x5d\x50\x59\x7a\xa2\x0e\x26\xdd\x6c\x2c\x3b\xcb\x7a\x87\x99\x3d\x8b\xb3\x69\xe0\x7d\x1e\xfe\x78\xf7\xdf\xee\x96\x2f\xed\x99\xe7\x57\xcc\x3f\x03\x00\x00\xff\xff\x98\x7c\x0c\x06\x40\x0f\x00\x00"),
		},
		"/relay_template.yaml": &vfsgen۰CompressedFileInfo{
			name:             "relay_template.yaml",
			modTime:          time.Date(2022, 6, 9, 10, 25, 17, 658768698, time.UTC),
			uncompressedSize: 6038,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x02\xff\xec\x97\xc1\x6e\xdb\x38\x13\xc7\xef\x7a\x0a\x22\x77\x39\x09\x5a\x7c\x28\x04\xf4\xa0\xcf\x76\x12\x23\x89\x2c\xc8\x4a\x8a\x9e\x0c\x56\x9a\x38\x44\x28\x92\x4b\x8e\x9c\x35\x8a\xbe\xfb\x82\xb2\x24\x4b\xb6\xe4\x06\xae\x50\x6c\xb1\xc9\x29\x26\x67\x86\xbf\x19\x8f\xc9\xf9\xbb\xae\xeb\x50\xc5\x1e\x41\x1b\x26\x85\x47\xd6\x97\xce\x0b\x13\xa9\x47\x02\x9a\x81\x51\x34\x01\x27\x03\xa4\x29\x45\xea\x39\x84\x70\xfa\x0d\xb8\xb1\xff\x11\x92\x48\x81\x5a\x72\x57\x71\x2a\xc0\xab\x3e\x72\xd0\x6e\x46\x05\x5d\x81\x76\x08\x11\x34\x03\x8f\x28\xaa\x29\xcf\x8d\x6b\x36\x06\x21\x73\x1c\x7b\x68\xf7\xa9\xa1\x5d\x31\x08\x02\x1f\x25\xcf\x33\x18\x73\xca\xb2\x16\x41\x3b\xe2\x4b\xfe\x0d\x5c\xa5\xe5\xdf\x1b\x97\xe6\x29\xc3\xd2\xa0\x00\x3f\x38\xd7\x28\x48\x6c\x08\x9a\x24\x60\xcc\xbd\x4c\xa1\xc8\xc4\x25\x11\xd0\xf4\x8b\x66\x08\x73\x91\x80\x43\x88\x06\x23\x73\x9d\x40\x99\xa8\x86\xbf\x72\x30\x58\x7e\x22\xc4\xa0\xd4\x74\x05\x1e\xb9\xbc\x66\x0e\x21\xeb\x82\xd4\x86\xf3\xc8\x15\xe3\x50\x9e\x46\x9c\xfd\xda\x52\xa5\xcc\x79\x9d\xea\x04\x14\x97\x9b\x0c\x04\xf6\x56\x98\x2a\xd5\x95\xea\x91\x2a\xbc\x29\x7f\x0d\x8a\xb3\x84\x1a\x8f\x5c\x3a\x84\x28\x2d\x57\x1a\x8c\x99\x00\x4d\x39\x13\xb0\x80\x44\x8a\xd4\x6e\x7e\xba\xb8\x70\x08\x31\xc0\x21\x41\xa9\xb7\x48\x19\xc5\xe4\xf9\xae\xc1\x78\x8c\x12\x21\x53\x9c\x22\x94\xae\x8d\x2c\xed\x1f\x6f\x45\x39\x16\x87\x90\x0a\xbd\xea\x3b\xca\x04\xe8\xda\xd7\x25\x54\xaf\x1a\x91\x5c\xe2\xba\x99\x4c\xe1\xb3\x01\xbd\x2e\xfa\x70\xb7\xce\xe5\xca\xe5\xb0\x06\xfe\xf9\x43\xbd\x0e\x62\xdd\x74\xde\x16\x37\x9c\x4f\x96\x81\x7f\x3f\xad\x37\x08\x59\x53\x9e\xc3\x95\x96\x99\xd7\x58\x24\xe4\x89\x01\x4f\x23\x78\x6a\xaf\xda\x84\x5a\x1d\xde\xde\x2c\x9c\x42\x8a\xcf\x5e\x5d\x98\x91\x3d\xb8\x97\x63\x11\xfa\xe3\xdf\x0d\xb3\xbd\x00\x0e\x88\xfc\xc8\xbf\x7b\x58\x2c\xa3\xe9\x9d\xff\x75\x19\x4e\xa7\xd1\x62\x1a\x3d\xce\xde\x42\x97\x48\xf1\xc4\x56\xf7\x54\xdd\xc2\xa6\x03\xf2\x05\x36\x1e\x31\x20\x50\x6f\x46\x0a\x40\x33\xb1\x1a\x3d\x4b\x83\x7b\x66\xbd\x77\xc0\x36\xfc\x01\xf0\x62\x1a\xc4\xd1\xd7\xa5\x3f\x99\x44\x83\x22\xd2\x34\xd5\xbf\x88\xf6\xff\xf9\x3c\x5e\xc4\x91\x1f\xfe\x2b\xe9\x76\xdf\xf0\x2c\xb8\x5e\xc6\xf3\xdb\x69\x30\x0c\xa2\xfd\x72\x47\x28\x5f\x40\x0c\x42\xf8\xb0\x98\x46\x43\xe2\xe5\x66\x78\xbc\x9b\xf9\x22\x5e\x86\xf3\x28\x1e\x10\xd1\xfe\x36\x46\x4a\x6a\x1c\x04\x73\x3c\x0f\x82\xe9\x38\x9e\x0f\x5a\xca\x44\x0a\x51\xbc\x1f\x03\xd6\x73\x07\x3a\x70\x51\x77\xb0\xc3\x56\x36\x98\xc6\x5f\xe6\xd1\xed\x72\x36\x19\x86\x53\x03\xa7\x1b\x01\xf8\x2a\xf5\xcb\x88\xa5\x27\x32\x6e\xff\x0e\x90\xfd\x87\xc9\x2c\x5e\x86\x7e\x7c\x33\x0c\x6c\x31\x98\x8d\x14\xc5\xe7\x13\x39\x59\x56\x4c\x5b\x67\xdf\xbf\x93\xd1\x44\xbe\x0a\x2e\x69\x3a\xb1\x4f\x54\x64\xcb\x30\xb3\xbb\xe4\xc7\x8f\xb3\xb6\x7d\x98\x73\x1e\x4a\xce\x92\x8d\x47\x66\x4f\x81\xc4\x50\x83\xbd\x18\x9d\x9f\x9d\x5e\x1b\x20\xe8\x8c\x09\x8a\x4c\x8a\x7b\x30\xc6\xc6\x2c\x5e\xc8\xf3\x14\xd6\xe7\x8d\x4d\x3b\x52\x1c\x73\x2a\x21\xec\x68\x58\x9b\x55\x33\x63\x2e\xb0\x35\xb7\x64\x76\xa5\x3c\x46\x2a\x3c\x2f\xf1\xce\xab\xe1\xb6\x0d\xdf\x5c\x4d\x85\xa9\x4e\x1a\xf3\xdc\x20\xe8\x2b\xa6\xeb\x97\x53\x83\x41\xaa\xb1\xb2\xf0\xf9\x2b\xdd\x18\xe7\x00\xf9\x5a\xd3\x04\x42\xd0\x4c\xa6\xf5\x18\xf8\xe1\xc2\x69\x32\x37\x86\xae\x43\x08\x42\x54\xd7\x00\xdf\x6c\x8b\xc4\x2e\x04\xc7\x87\x78\xa7\x57\x95\x8c\xab\x7e\x7b\x9b\x26\xa8\x1b\xa9\x7f\x28\xae\x62\xec\xba\xd4\x23\x67\x87\xb5\x3f\x2b\x26\xe1\xfa\x5d\xed\xea\xc6\x45\xb1\xed\xa7\xa9\x2e\xbb\xb1\x63\x90\xe9\xf2\x0b\x01\xf4\x8d\x34\x58\x7a\xed\x9e\xc6\x3e\xe3\xd8\x6e\x96\xd6\xbb\x97\xaa\xcb\xfa\xc1\xb4\xad\xf7\x2e\xe3\x2e\x97\x71\x65\xd2\xed\x57\xdf\x8b\xbd\xbf\x47\x9b\x4a\x28\x35\x36\x01\x8f\x7a\x59\xc8\x3d\xa7\xbd\xfb\xad\xf7\xac\x60\x6b\x32\x9b\x14\x7e\xfd\x6d\xb3\x00\xbd\x66\x7b\x52\xf6\x97\x24\x94\xcd\xa5\x14\x8f\xdb\x38\x98\x28\xb7\xa0\x2e\x5a\x7d\x9b\xea\xc7\x8f\x5b\x89\xa1\xb4\x44\x99\x48\xee\x91\x78\x1c\x16\x2b\x48\xf5\x0a\x8a\x84\x2b\xab\x2a\xce\x33\x62\x19\xc8\xcd\xa4\x60\x28\x6d\xeb\x34\x62\x7e\xba\xb8\x78\x4b\xd0\xd2\xac\xad\xdd\xfa\x04\xd6\x69\x42\x95\x0a\x21\xb1\xb8\x35\xcc\x1b\x75\xab\x8b\x94\xf1\x63\x3f\xd7\xc6\xfe\x4f\xca\xdf\xa7\x59\xff\x57\x48\xd6\xb6\xc0\x3d\x55\xc0\x56\x38\x83\xa8\xd8\x2a\xd8\xc9\x52\xb6\xf6\x7f\x17\xac\xef\x82\xf5\x5d\xb0\xfe\xb1\x82\xf5\x3f\x3e\x5d\xbb\xad\x7b\xec\xcf\x1a\xb1\x7f\xdf\x00\xfc\x4f\x00\x00\x00\xff\xff\xda\x9c\xd6\x9a\x96\x17\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/agent_templates.yaml"].(os.FileInfo),
		fs["/relay_agent_template.yaml"].(os.FileInfo),
		fs["/relay_template.yaml"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
