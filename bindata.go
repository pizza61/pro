// Code generated by go-bindata.
// sources:
// lib/bs.glade
// lib/icon.png
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _libBsGlade = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\xdd\x6e\x9b\x30\x14\xbe\xef\x53\x78\x9e\xd4\x9b\x29\x21\x3f\xad\x34\xa9\x40\xa5\x49\x6b\x6f\x76\xd9\xa9\x97\xc8\xc0\x09\x38\x71\x6c\x66\x1f\x12\xe8\x43\xed\x66\x6f\xb0\xbd\xd8\x44\x69\xd4\x50\x9c\x1f\xc8\x2a\xad\xd3\x6e\x2a\x84\xcf\x67\xfb\xf8\xfb\xa9\x89\x7b\x5d\x2c\x05\x59\x81\x36\x5c\x49\x8f\x8e\x87\x23\x4a\x40\x46\x2a\xe6\x32\xf1\xe8\xd7\xbb\x9b\xc1\x47\x7a\xed\x9f\xb9\xef\x06\x03\x72\x0b\x12\x34\x43\x88\xc9\x9a\x63\x4a\x12\xc1\x62\x20\xd3\xe1\x64\x32\x1c\x93\xc1\xc0\x3f\x73\xb9\x44\xd0\x33\x16\x81\x7f\x46\x88\xab\xe1\x5b\xce\x35\x18\x22\x78\xe8\xd1\x04\x17\x1f\xe8\xf3\x42\xd3\xe1\x64\x44\x9d\xc7\x3a\x15\xce\x21\x42\x12\x09\x66\x8c\x47\x6f\x71\x71\xcf\x65\xac\xd6\x94\xf0\xd8\xa3\x4f\xcf\x55\x21\x21\x6e\xa6\x55\x06\x1a\x4b\x22\xd9\x12\x3c\x1a\x31\x19\xcc\x54\x94\x1b\xea\xdf\x30\x61\xc0\x75\x36\x05\xf6\xfa\x18\x66\x2c\x17\x18\xac\x79\x8c\x29\xf5\x2f\x47\xa3\x16\x22\x4a\xb9\x88\xeb\xe7\x0a\x2f\x58\x04\xa9\x12\x31\x68\xe7\xa9\xc0\xd9\xaa\x78\x51\xdd\x6a\xe4\x93\x2a\xe8\x66\xb4\xbd\x9b\x15\x37\x3c\x14\x40\xfd\x3b\x9d\xb7\xb6\xde\xa7\x5d\x1b\x46\x69\x0e\x12\x19\x72\x25\xa9\xbf\x02\x8d\x3c\x62\xc2\x0a\x6c\xf4\x62\xef\xe7\xb3\x44\x5d\xd6\xbc\xac\x95\x8e\xe9\x76\x75\x8f\xfe\x0e\xf4\xd8\x05\x86\x4a\x09\xe4\x59\x80\x50\x20\x25\xa8\x99\x34\x82\x21\x0b\x05\x78\xb4\x04\x43\xfd\xfb\x8c\x9b\x07\x82\x39\xb2\x39\x31\xeb\x5f\x3f\xe6\x04\x61\x61\xf0\xd8\xf9\x97\x4c\x27\x5c\x06\x02\x66\x48\xfd\x71\x4b\x36\x07\x60\x9a\x27\x69\x1f\x1c\xaa\xac\x07\x2a\x54\x88\x6a\xd9\x0d\x58\xd4\x9e\x08\xa2\x94\x69\x43\xfd\xc9\x85\x1d\xea\x3a\xb5\x26\x1a\xef\x32\x16\x2d\xb8\x4c\xf6\x2f\x01\x45\xc6\x64\xbc\x47\xb8\x36\xd0\x8c\x0b\xd1\x4d\x07\x99\x32\xbc\xd6\xfa\x8e\xe6\x5d\xa7\xb5\xdd\x86\xa5\x89\xdd\x0a\xed\x20\x38\x16\xd9\x0e\x85\x1c\x51\xc9\x66\x34\xd8\x5a\x39\xd5\x3e\x9d\x4e\xfa\x44\xa1\x76\x97\x5c\xc8\x0c\x08\x2e\x21\x78\x66\xac\x9e\xe4\xd8\x09\x04\x2b\x55\x8e\x81\xc1\xb2\x3a\x23\x90\xf1\x4e\x60\x8b\x13\x3b\x2f\x5f\x58\x08\x82\xbe\xac\xeb\xc9\xcb\x29\xdc\xd8\xbb\xad\x36\x67\x0d\xb6\x32\xe4\xa0\x1f\xc8\x1c\x62\xa9\xba\x4c\xd9\x30\xfc\xf8\x72\x37\xd4\x66\x7a\xb2\xd3\xf8\xb6\xa5\x36\xe6\xef\x7a\x62\x87\xfd\x6f\x43\x1d\xcc\x00\x62\xcf\x01\x62\x73\x34\xe9\xa0\xa0\xda\xd9\xf5\xff\xc7\x70\x7c\x84\x94\x76\xb3\x9a\x5e\x14\x23\xdd\xa5\xe9\xea\x2f\xf5\xc3\x71\x17\xcc\x9f\x50\x72\x57\xa8\x86\x08\xf8\x0a\x4c\xf0\x74\x17\xeb\x3e\x43\x33\xaa\xf6\xe8\x96\x10\xf7\x31\x1d\xda\xef\x2b\x4a\x2b\xd6\x36\x97\x07\xa4\x8e\x05\xec\x58\xd1\x6f\xd7\x0d\x7b\xb4\xf1\xea\x6e\x98\x9c\xe4\x86\xf7\x17\xe7\x09\x5e\x9d\x0b\xbc\x1a\xfd\xfc\xde\xc3\x15\x93\xff\xae\xd8\xe2\xf4\xad\xea\x77\x0f\x8b\xaf\xae\xdf\xe9\x69\x69\x5e\xf4\xc9\xf2\xe9\xbf\xa6\xda\xb7\x2b\xbd\x3d\x54\x1c\x2f\xbd\xbf\xff\xe3\x89\x45\x8b\x00\xcb\xec\xc0\x7d\xba\xab\x41\x0f\x7f\x72\x35\x4f\x66\x6b\xf0\x79\xc0\x75\xb6\x7e\x5c\xfa\x1d\x00\x00\xff\xff\xee\x8c\xe5\x74\xb5\x12\x00\x00")

func libBsGladeBytes() ([]byte, error) {
	return bindataRead(
		_libBsGlade,
		"lib/bs.glade",
	)
}

func libBsGlade() (*asset, error) {
	bytes, err := libBsGladeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/bs.glade", size: 4789, mode: os.FileMode(420), modTime: time.Unix(1545238809, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _libIconPng = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x96\x57\x50\x13\x6a\xd4\xae\x3f\x40\x7a\x11\x41\x45\x40\x4a\xe8\xbd\x47\x8a\x41\xa4\x48\x07\x03\x46\x6a\xa4\x37\x43\x15\x42\xe8\x06\xdc\x8a\x05\x05\x91\x20\x35\x04\xc1\x08\x31\x10\x41\xa9\x52\x6d\x48\x8d\x20\xa1\x48\xe8\x75\xd3\x11\xe9\x20\x9c\xd9\x67\xce\xed\xf9\xe7\x9f\x35\xcf\xcd\x33\xeb\x76\xad\xf7\x7d\x0c\xb7\x35\xe3\xe6\x10\xe2\x00\x00\x70\x5b\x98\x9b\xd8\x03\x00\x12\xff\x83\x8d\x05\x00\xf0\xe5\x8a\x67\x17\x00\x80\xc5\xcb\xca\xcc\x04\x9c\xfc\x37\x85\x4d\x25\x19\x00\x00\xf6\x30\x73\xe7\x08\x00\x54\x24\xff\x83\x21\x9a\x68\x80\x01\x00\xb0\xa2\x2d\x6c\xae\xb1\x4e\x73\xf1\xf3\xea\x32\x44\xbe\x12\x66\x06\x00\x08\xa2\xaf\x39\xa1\x8d\x43\x83\x83\x7d\x43\xd0\xc0\x38\xdc\xd7\x13\xed\xeb\x23\x1e\x75\x1b\x1d\x20\x6e\x66\x61\x03\x77\x4c\xe2\x11\x00\xe0\xb4\x92\x85\x89\x21\x22\x7a\x64\x15\x9f\x00\xbf\xf9\x4d\xf8\x7b\xcb\xbd\xc6\xc7\x4e\x6f\x46\x25\x3a\xb6\x78\x1e\xa9\x46\x8a\xac\xe3\x72\xc2\xfb\x3c\x7d\xb2\xfb\xda\xed\xf8\x65\xff\xb9\xaf\xd5\x6f\x89\x57\xba\x00\x71\x63\x64\x31\x95\x93\xea\xa3\xb1\xd0\x18\x29\xaa\x4e\x16\xcf\x3c\x25\xbf\x9c\x62\x64\x79\xc1\xc7\x81\x34\x56\xe5\x0d\xd5\xa6\x20\x21\x31\x19\x43\x7f\xff\xf6\xf5\xaa\x36\xc5\x25\xc6\x51\xeb\xc3\x1f\xff\xa5\xe6\x63\x31\xbb\xb0\x86\x85\x95\x85\xfc\xb1\x0f\xfe\x51\x04\x76\xc0\x96\xdd\x56\xde\x3d\xac\xec\x01\xd4\xd4\x41\xf6\x05\x60\xcd\x05\x64\x58\x00\x3f\x53\xe2\x63\x86\xd6\xaf\xc0\xc3\x0b\xa8\x3d\x73\x04\x38\x4b\x69\x36\x48\xac\x19\x1a\xad\xd5\xe0\xbd\xc2\xfd\xff\xfc\xff\x6f\x5d\xd9\xe6\xec\x17\x73\x38\x9c\xd7\x81\x82\x64\x27\x21\x9e\x2c\xf6\x97\x7c\x19\x19\x11\xe2\x97\xb1\xbe\xea\xef\x2f\xab\x77\xf9\x72\x5f\x6f\x6f\x62\xe0\x10\xc5\x90\x84\x90\x90\x95\xfd\x2e\x27\x30\xbd\xb8\xf8\x38\x72\x79\xf0\xd3\xaf\x5f\x82\x9a\xfe\xd4\xaf\x34\xd2\x0d\x4f\x6f\x6f\x26\x45\x45\x45\x5e\xa3\xde\xfe\xfe\x7f\x76\x76\x30\xff\xa2\xf5\x10\x8e\x8e\x92\xda\xda\x02\x95\x95\x95\x4f\xc1\x51\x93\xca\x64\x92\x31\x09\xc1\x97\x5b\x54\x5c\x2c\x66\x2e\x70\xee\xdc\x35\x54\x60\x60\x87\xf9\xc0\xc0\xc0\x35\x0d\x0d\x8d\xf6\x7c\xdd\x08\x36\x65\x77\xe3\xb4\x73\x8a\x0e\xcf\xb7\xb6\xd0\x45\x0e\x64\x0b\x2c\xd6\x40\x1d\x11\xb5\x36\x0a\xb1\x24\xbe\x30\x97\xe2\x73\xdb\xa7\xff\xfa\x95\x4e\xaf\x0d\xbd\x11\xac\x0d\x85\x9e\xc9\xb1\x96\xd9\x3b\x58\x1b\x6f\xb2\xaa\x76\x77\x70\x90\xe9\xd4\x9f\x65\x4e\x34\x91\x80\x64\xa9\xdc\x92\x60\x62\x62\xba\x9f\xb6\xb4\xbc\xfc\xb8\xa4\x44\x76\xdd\x28\x74\xa2\x19\x8f\x6c\x8e\xcf\x41\x47\x45\xf1\xab\xbf\xee\x0b\x68\xdf\x59\xa5\xc7\x64\xfa\xae\x8e\x54\xbf\x2e\x2f\x7f\x86\xaa\x2d\x2d\x35\x87\xcf\x21\x5b\xee\xe2\x51\xb5\x44\xe2\x83\xba\x86\x86\x34\x1e\x60\x4f\xc2\x95\x22\x14\x3e\x69\x37\x91\x48\x8f\x77\xaf\x05\x68\x09\x49\xd4\xac\x4d\xb4\x5c\xff\xfc\xf9\x73\x74\xa6\xff\x44\x73\x42\xfe\x69\x63\x7b\x7b\x3e\xd5\xbc\x92\xeb\x84\x37\x23\xcb\xcb\xcb\x92\x10\x08\xc4\xeb\x66\xf3\x14\xc0\xe1\x70\x8f\x30\xfd\x29\x29\x29\x36\x2c\x89\x9f\x64\x96\x56\x6b\x27\xe6\x6f\xbe\x29\x2b\x53\xb2\xad\xe0\xaa\xae\xf7\xff\x21\xfa\x93\x4e\xb7\xaa\xce\x4a\x4f\x87\xf8\xcd\x65\xb0\x17\x96\x71\x23\x14\x32\x87\xab\x6e\x67\xb8\xf9\xbf\xaa\xaa\x3a\xcb\x29\xa0\x12\x2f\x5e\xd6\xd4\x24\x42\x24\x12\xeb\xc2\xe7\x38\x3a\x3b\x3b\xe1\x76\x76\x0c\x9c\x9c\x9c\x4a\x4e\x15\x20\x24\x24\xc4\xe9\x23\x5a\xbe\x20\x61\xc7\x99\xca\x63\x8e\x46\x6b\x19\x19\x19\x71\x2d\x55\x05\x68\x09\xc1\xb6\x69\x37\x4e\xd6\xbf\xf2\xde\x6b\x6b\x33\xa5\xd1\x68\x93\x83\xe5\xae\x70\x12\xa2\x37\x40\xab\xf4\xed\x5b\x23\x2b\x2b\xce\xf9\x9e\x3c\x56\x7f\x7f\x7f\x25\x97\x2a\xb6\x1f\x06\x27\x30\x24\x12\xe9\xe1\x67\x7a\x02\x50\x41\x41\x5f\x47\xeb\x23\xfa\x06\x07\xfd\xe0\x39\x52\x7c\xec\x4f\x9e\x3c\x51\x0b\x1c\x54\x42\xc5\x38\x96\x39\x25\x75\x76\x76\x7a\xae\x2a\x39\xbf\x3f\xb5\xf0\x83\xf0\xbd\xea\x76\xff\xd6\xb4\x82\xa2\xe2\xd4\xf4\xb4\x46\x2c\xbe\x15\x4b\x6e\x69\x21\x1c\x6d\xf5\x5f\x9f\xfd\x26\xe0\x7e\xf1\xe5\x7b\x1c\xee\x2c\xd5\xd6\xe5\xda\x35\xe6\x75\xed\xab\x11\x11\x9a\xfb\x53\x0f\xc5\x1e\xa5\xa5\xa5\x19\x2e\x1c\xee\xed\x71\xfd\x7a\x4a\x7c\xf3\x66\x33\x36\xb7\xb1\xa9\xc9\xc2\xcc\x8c\x75\xc0\xcf\xd0\xf0\xd3\xa3\xa0\xb0\x26\xcf\xf5\xb1\x06\x33\x87\x77\x6e\xca\x4b\xfd\xa1\x61\xfc\x6a\xa5\x05\x05\xa2\x2f\x9e\x3f\xf7\x80\x9b\xca\xcb\xcb\xe7\x07\xe7\x66\x65\x65\x19\xd6\x58\x5a\x59\x75\x3c\xe6\x97\x8d\xde\x1d\x67\x4d\xe3\x14\x50\x81\x77\xd1\x46\x47\xbb\xe9\x7e\xf1\xf1\x30\x7f\xec\x64\x8a\x6a\xf3\x64\x64\xda\xc1\xe1\xa1\x05\x02\x81\x53\x57\x20\x88\x01\x59\xdb\xfc\xc9\x4c\xe7\x9b\x37\xcf\x0d\x8c\x0f\xb9\x63\x83\xcb\x52\x84\x77\xea\x30\xab\x42\xd5\xd6\x84\xf8\x4d\xf6\xd0\x3c\x83\x93\x23\x4d\x7f\x31\xef\x3f\xf3\x3d\xb4\x08\xcc\xda\xa8\x45\x6c\x6c\x6c\xe9\x82\xaf\x9f\x9f\xd0\x1f\xd1\x0b\x2c\x57\x63\x62\x74\xd7\xbb\x53\x2f\xea\x1a\x55\xe7\x10\x08\x06\x28\x50\xe6\xde\x52\x10\x3b\x13\x11\x1e\x1e\x4d\x47\xa3\xd1\xb3\x17\xc5\x48\x64\xb2\xa4\x3d\x69\x2b\x69\x2b\xbc\x3e\x5f\x0b\x0a\xb5\xeb\x27\x0f\xa0\xbc\xfc\xfd\xfd\xf5\xc0\x29\x56\xd6\x3f\xd4\x20\x6a\xbe\x2e\x6d\x3c\xd8\xc6\xa6\xa7\x94\x01\x24\xe2\x32\x33\xed\xbe\x2b\xdc\x1f\xde\xfa\xfd\x1b\x82\xc7\xe3\x69\x71\x7b\x6a\xc8\x6f\xdf\xbe\xd1\xd0\xab\xab\xab\x3e\x7f\x0f\xb6\xdb\x9a\x8f\xd3\x99\x7a\x69\x34\x29\x1d\xe3\x11\x59\x64\x60\xe0\x72\x29\x63\xd8\x7c\xb7\x20\xca\xc0\xec\xdf\x8c\xa8\x8d\x49\x43\x2a\xeb\xed\x8a\xd3\x99\xf9\xf9\xd6\xcf\xc4\x2e\x5c\x10\x27\xdc\x3d\x58\x2c\xae\x3b\x4e\x67\x82\x3b\x39\x05\xe0\x4a\xca\xc9\xe4\x40\x04\x83\xb9\xb5\xf5\x6b\xb7\x77\x04\x02\xc1\x83\x1e\x13\x13\x83\xa0\x8b\x88\x88\x48\xe9\xe9\xe5\x46\x68\x69\x5e\xbe\x7c\x98\x8e\x6c\x30\xc4\x32\x91\xc9\x64\x66\x6c\x04\x26\x32\x72\xf1\x81\xb3\x8b\x4b\xea\xd2\x0f\x5e\xbc\xad\x19\x37\x41\x3f\xae\xc6\x83\xee\x5a\x13\x64\xe8\xed\x5d\xec\xb6\xf7\xda\x22\x83\xf7\xec\x59\x93\xa6\x9d\x91\x11\x97\xa0\xf1\xc6\x5f\x5b\xc7\x2c\x35\x35\x35\x1c\xd4\x8a\x86\x06\xd7\x13\x42\x55\x8e\xa0\x8f\x86\x20\xc5\xbd\x85\x4a\x5f\x7b\x21\x6d\x35\xb9\xbc\xcc\x29\x2b\x2b\xdb\x9d\xad\xce\xb4\xb3\xb3\x53\xf1\xfe\xfd\x3d\x22\x51\x5a\x49\x59\xd9\x33\x24\x24\xa4\xc2\x66\xc4\xe3\xe9\xab\xb7\x65\x65\x26\x01\x01\x72\x3c\x62\x57\xb8\x14\x15\xc9\xac\xd3\x32\xd6\xd9\x57\x13\xae\x5e\x39\xde\xf3\xa6\x0d\x0d\x9d\x68\x33\x4c\x2a\xb5\x9b\xdc\xc5\xbe\x99\x99\x9e\xae\x40\x0d\x28\xd4\xd5\xd5\x6d\xd4\x7c\x92\x0e\x0c\x0c\xb4\x5b\xb9\x68\x70\xf7\xa2\x6a\xd3\x16\x5f\x4d\x30\xfd\x52\xd0\xc8\xec\xac\xcf\x6b\x22\x31\x67\x3e\x3e\x31\x29\x89\x21\x63\xff\xe1\xdf\x41\xd7\x50\x88\xb8\x38\x88\x58\xa0\x42\x5d\x44\xcf\x9f\x3f\x5f\xf4\x2b\x76\x01\x1f\x25\x9e\xd1\x51\xe1\xd3\xc1\x47\xd5\x8a\x89\x89\x19\xae\x0e\x14\xc8\x3b\xa3\x95\xbf\xb3\x36\x96\x62\xfc\x90\x47\xb0\x7a\xb3\xde\xad\xcf\x08\xbd\xbd\x74\x93\x5a\x77\xb8\xbb\xee\x7d\xb8\xb3\x9a\x76\xe4\x96\xaa\x0b\x30\xab\x23\x0f\x36\x36\xc2\x86\x87\x87\x9f\x92\xc2\xe3\xe2\xe2\xbe\x4c\x4c\x60\xdb\xe7\xe7\x3f\xfc\xb8\x12\x37\x3b\x97\xab\x33\x08\xf6\x36\xbe\x9c\x7e\x18\x94\x07\x15\xe2\xa6\xa1\xa0\x28\x14\x0a\xee\xea\x9a\x72\x70\x10\x4b\x1b\x1a\x2a\x7e\xf5\x8a\xc1\xa5\xfc\xe5\x48\x3d\x07\x37\xf7\x35\x1f\x9f\xd7\x7d\x9d\x4a\xe9\xe6\xbc\xc2\xc2\x29\x16\x19\x32\xdc\x39\xd4\xfe\x7e\xfe\xf4\xf4\x74\x36\x4e\x4e\x6f\xd0\xbb\x35\x97\xe5\xee\x35\xdf\x9d\x43\xab\x9b\xc0\x9e\x88\xb0\xf2\x5c\x6c\x4d\x97\x34\xe3\xaa\x39\x0c\x5b\x47\x06\x11\x89\x44\x3b\x2c\x7f\x7f\x73\x32\x66\x7d\xfc\x99\xfb\xdd\x7d\x66\x01\x15\xa4\x07\x80\x29\x39\x3a\x3a\xa2\xf6\xe6\xba\xb2\x2a\x83\x7e\x69\x8c\x63\x08\xd1\xd3\xec\xc5\x99\xba\x32\x32\x2c\x8a\x8e\x94\xd7\x6d\x91\x68\x74\xeb\xe0\xe0\x4e\x52\x85\xfc\x8d\x37\x1e\x27\xc7\x7f\xd3\xb8\x12\xd5\xfb\x59\x9f\x5b\x12\xc3\xc2\xc3\x8b\xee\xe6\x1e\x1c\x1c\x70\xbc\xd7\x4f\x38\x40\x17\xe3\x7e\x75\x75\x71\xae\xb7\x9c\x5c\xa9\xd6\x74\x46\x20\xee\xb7\xb7\x9b\x6f\x6d\x6f\x2f\x7e\xeb\xc1\x5f\x16\x98\x48\xd8\x15\x61\x9a\x9c\xb0\x97\x3b\x4b\xb3\xfb\xf0\xe2\x05\x1f\xe2\xe6\x4d\xf6\xf7\xc1\x4f\x9d\x2f\x61\x30\x98\xe1\x16\xac\xc1\xca\xca\xca\xd6\x74\x95\x8b\x72\x4d\x30\x3d\xb9\xae\x4e\xc7\xc7\xc7\x07\xfe\x03\xbb\x3d\xe0\x58\x19\xb5\xee\xc6\xd2\x5a\x58\x37\x72\x9c\x9c\xcc\xf2\x49\x44\x8b\x84\x50\x88\x8d\xb9\x5e\x00\xb3\x48\xf0\xde\xde\xd9\x91\x50\x50\x58\x22\xd8\xfd\xe0\x16\xbe\x74\x95\x68\xbf\x48\xd4\xd4\xd6\x36\xf3\xf0\x10\x3f\xcb\xc1\xcc\xe4\x51\x11\x3c\x7b\x70\xf7\x7c\x51\xc2\xe1\x8e\xe7\xe6\x6c\x47\xd1\x4b\xa8\x92\x12\x17\xd9\xa1\xbc\xa8\x0d\xaf\x1b\xc1\xf2\x90\x47\x24\x56\x3c\x68\x77\x6d\x6c\x6f\xf9\x5d\xcb\xd6\x4f\x5b\xbc\xde\x4d\xd5\xd5\x3a\xe6\xc9\xb5\xb1\x06\x5a\x18\x2b\x2b\xeb\xf8\x25\x5f\x0d\xc1\xc6\xc6\x46\x13\xe5\x92\x87\x0f\x59\x77\x56\x7e\x19\xe3\xd1\x2e\x2e\x2e\x96\xab\xef\x7c\x7c\x7d\x21\x10\x48\x92\xde\x4f\x85\x8c\x8e\x42\x84\x82\x20\x40\x55\xe7\x0a\xa9\x64\x75\xc5\xd0\x61\x98\x95\xec\x7f\x45\x5a\x99\x58\x79\x26\x2b\xbf\x8b\x16\x33\x1e\x8d\xc7\xad\xdf\x29\x10\x38\x73\x26\xa9\xbb\xdb\x1a\x35\x9a\xcc\x2d\xfc\xf4\x92\x09\xcc\x0f\x98\x5f\xbf\x5e\x72\xfc\x34\x3c\x2c\xac\x95\x74\xa3\x84\x6b\xad\xf0\xd6\x47\x34\xbb\x6a\xe3\xc6\xbd\x52\xdf\xa3\xf9\xbc\x50\xb5\xd0\x71\x58\xd3\x46\x73\xf3\x95\xaf\x6d\x6d\x45\x85\x85\x1e\xed\x7f\xb4\x83\x47\xee\xd3\x6b\x43\x97\x1f\xb8\x3a\x39\x89\x4b\x49\x99\xf0\xae\x78\x8b\x26\x7e\xfa\x24\x9a\x7b\x29\x64\xf4\x51\xc4\x02\xd5\x2e\x6e\xf4\x74\xe0\x10\xe5\x96\x9c\xbc\xfc\x99\xa6\x88\x73\x17\x2e\x3c\x82\x06\x0e\xc2\x67\x34\x34\x34\x78\x05\x04\xb4\xcb\x9e\xd5\xd4\x5c\x8a\xdf\xea\x33\x0b\x08\x0d\x1d\x4a\xf9\x50\x5b\xfb\xbd\xaf\xd8\xec\x59\xb8\x7b\x8c\x2a\xcf\x5c\xc4\xec\xc2\xc2\x8d\xe0\xae\xb9\xcd\x33\xe5\xec\x17\x51\x67\x4a\x1d\xc8\xd2\x82\x82\x82\x52\x3a\x3a\xd9\xd1\xb1\xb1\x7c\x24\x84\x65\x63\x63\x63\x3b\xab\x58\xfc\x55\x5b\xbc\xde\x35\x84\xc2\xb9\xc8\xce\x4b\x43\x42\x8d\x38\x1c\xee\xcc\xad\xfe\xf7\xbe\x5d\x8f\x6c\x4b\xf8\x83\xf9\xca\xd5\x5f\xe6\x1a\x76\x7f\x7b\x55\xcc\x0c\x80\x3d\x1c\x2e\xbe\x14\x5e\xd1\xd9\x69\xc9\x77\xfe\xbc\x0c\x5b\xe2\xc4\xdd\xc3\x67\xb1\x77\x03\x03\x15\x7f\x6f\x6c\x5c\xd8\x26\xfa\xf1\x06\x5a\x66\xcb\xb7\x4f\x4c\x88\x41\xa1\xd0\xa2\xe1\xb7\x64\xf2\x97\xbc\xd0\x66\x76\xa3\xbd\x38\xaa\x95\x8d\xcd\x80\xab\x0a\xeb\x91\x7e\x94\x04\x04\xc2\xe0\xec\xec\xdc\xb8\xd2\xfc\xbb\x4d\x78\x6f\x26\xcd\x60\xeb\x73\x79\x59\x19\xef\x99\x33\x49\xba\x11\xf3\x4f\x5b\xae\x80\x8d\xdd\xb3\x8a\x58\x2c\x16\x81\x40\xb0\xbd\xbf\x61\x6f\x7f\xcf\xd1\x2d\x86\xf9\xc1\x83\x07\x8d\xc1\x2b\xa3\xf5\xfc\xc2\xda\x21\xaf\xdb\x7c\xfd\xfc\x24\x24\x24\xee\xd5\x06\x97\x3b\xa4\x34\xc5\xed\xda\x2f\xc0\xf5\x74\x74\x26\x87\x2b\xbd\x67\xf8\x01\x68\xef\xe8\xf8\x74\x8f\x99\x33\x6c\x73\xd6\xa2\xaa\xaa\xea\x02\x3a\xc4\xd8\x98\x69\x7a\x66\x06\x92\xd5\xc5\xdc\x63\x88\xeb\x9a\x63\x63\x67\x7f\x45\x15\xfa\x98\x70\x78\x29\xf2\x4e\xf9\x2b\xbd\x84\x03\xad\x17\x2f\x5e\x94\xde\xa1\x50\x28\x3a\xd1\x1b\x85\x93\xb0\x52\xf2\x3f\xdf\x5f\x48\xf1\xb1\x1b\x19\x1b\x57\x06\xf4\x49\x74\x17\xe8\x0b\x63\x30\x98\x8a\xca\xca\xa4\x9c\x1c\xc1\xd8\xb8\xb8\xbe\xbe\xbe\xa4\xf9\x79\xbf\xd2\x01\x94\x8f\xbf\xbf\xe4\x95\x2b\xa2\xd6\x27\x10\x90\xdd\xe6\xc9\xff\x3f\x17\xa1\xff\xad\x96\xf2\x28\x6e\x0f\xa6\xd7\xfc\xfd\x60\xd3\x60\x22\x91\x54\x5e\xae\xac\xa4\xa4\xd4\x3b\x31\x91\xb6\xb6\x16\xf2\x2f\x5a\xef\xbf\x47\x51\x51\xc1\x4b\xc0\x1e\xef\x4b\xd2\x4f\x0a\x3a\xc2\x04\xab\x0a\x60\x3a\x3a\xd9\xa3\x1f\x23\xbb\x0a\xa1\x7b\xe5\xe5\xe5\xb3\xf4\xd0\x16\xe9\xd4\xd4\xd4\xe2\x25\xa1\x64\x11\xc9\x51\x71\xc6\x8b\x94\xde\xfb\xd1\x5c\xb2\xf7\x9f\x1b\x9c\x1c\x65\xac\x0c\x57\xfe\x33\xb6\x55\xc9\x92\x88\x54\x91\x45\x55\xbb\x68\xb1\xf6\x98\x6e\xdd\xeb\xe0\xa9\x64\x78\x6e\x2e\xc5\xd2\xb2\x69\x36\xfd\x56\x5f\x4d\xed\xf9\x58\x3c\xbf\x4e\xf2\x2d\xd1\x02\x3c\x9e\x3b\x2a\x2a\x8a\x1a\x9d\x69\xf9\x60\x74\x14\x39\x3b\x37\x37\x35\xe4\x8e\xfd\xd6\x7f\xdd\x80\xdb\xd1\xd1\x51\xc9\xbd\xf9\xa2\x55\x96\xca\xdc\xa8\x64\x0b\x91\xc8\xbc\x3c\x58\x6e\x6c\x6b\xcb\xb3\xfc\x20\x07\x25\xab\x26\x27\xf7\xe8\x43\x4f\x68\x60\xa0\xa4\xba\xba\xba\xf1\x4f\x22\x5c\x4e\x27\x6a\x52\xe2\xfe\x73\x45\x47\x4a\x45\x49\x2f\x2e\x0f\xa8\x9c\xd4\xe2\xf1\x78\x96\xc8\x79\x8a\x5b\x53\xe7\xa2\x88\xa6\xb8\x78\xe2\xcd\x7b\xf2\x8c\xbc\xc2\xc2\x5d\xfb\xb3\x32\x32\x32\x42\xb8\x8e\xa3\xfd\x3f\xba\x78\xc9\xac\xcc\xcc\x4a\x4b\xb5\x6b\x20\x64\xbc\xd1\xa2\xbc\x22\x4f\x3b\x74\xb3\x53\x4e\x46\x26\xf9\x8f\xf2\xe4\xf2\x72\x97\xc2\x67\x1c\xd8\xd8\xdc\x9c\x1e\x4d\xd7\x84\xc1\xf2\xa3\xfc\x22\x23\x23\x11\x79\x06\xc7\x7b\x92\x64\xb6\x1b\x8c\x93\x5b\x8b\xfd\xb4\x27\x1f\x23\x16\x52\x0b\xbb\xeb\x30\xe2\x03\x01\x5a\xa6\x0d\xad\x8f\x1e\xb1\xbd\xc0\xe3\x6d\x68\x28\xe8\xed\xe0\x60\xe9\xb1\xb1\xb1\xd3\xe8\xa0\x1f\x04\x83\x1f\x20\x7a\x4f\x4e\x55\xb5\xc4\x01\xe1\xe0\xf0\xb2\xb0\x1b\x06\x83\xe9\xb4\x28\x0c\x52\x6e\x29\xb2\x15\x3d\xac\xad\xb4\xb1\xb4\x94\x70\x29\xc7\x0c\xab\x3f\xfe\xcd\x03\xb6\x43\xc5\xb3\x86\x82\x20\x0d\x28\x46\xec\x65\xd3\xde\x00\x2d\x5a\xdd\x00\xd9\x31\x65\xd3\x86\x0c\x97\x2b\xb4\x27\x55\x44\x2e\x3b\x0c\xd3\xe9\xa7\xf3\xdc\x82\xd6\x0b\xbd\x8f\xa3\x2a\x08\x04\xb1\x01\xc5\x85\xd6\x56\x40\x2d\xd0\x57\xd4\x31\x1e\x39\xae\x05\x93\x2b\x2b\x4a\xc3\x35\x0d\x0d\x97\xa9\x05\xfa\x81\xcf\xad\x9c\x59\x5a\x8d\x5d\x1a\xa3\xd5\xa1\xaf\x6c\x4b\xf8\xab\xaa\x97\x0a\xeb\x60\x58\x1d\x3d\x3d\x8e\x06\x5e\x4d\x4d\x29\x59\xd9\x99\x0e\x59\x82\xd6\xbf\x8d\xbf\x79\x12\xc7\xe3\x13\xfd\x0b\x62\x92\x73\x2b\xea\xeb\x9d\x71\x5c\x64\x32\x99\x34\x65\x84\xd7\xa6\x38\x7f\x30\x7d\xb7\xb8\xb4\xb2\xf2\xf5\x99\xb0\x36\x1c\x81\xb8\xbf\xb9\x19\x8e\xcb\xca\xfa\xf2\xf9\x33\xa3\x46\x0e\x97\xad\xad\xad\x53\x63\x74\xa5\x47\xc9\x8d\x13\x5d\x50\x37\x62\x53\x09\xea\xee\xcc\xb0\x40\xe3\x3b\x71\xf2\x61\x68\xf4\xac\xfe\x9d\x59\x4f\x47\xe5\x98\x98\x18\x36\x4e\xce\xe2\xe3\x40\x8f\x13\xb1\xe4\x00\x46\x51\xec\xd1\x1d\x94\x51\x6e\x8b\xe8\x95\xce\x9e\x1e\xe2\xda\x8f\x50\x3f\xbf\xb6\x45\xa1\xc1\x42\x31\xe1\xfb\xe2\xca\xca\xdc\x21\x21\x21\x65\x45\xfb\xa0\x75\x7b\x37\x75\x7e\x28\x35\xad\xca\x45\x39\xf2\x74\x3c\xc8\x56\xf7\x4e\xf6\x38\x14\x50\x75\x3b\xbb\x62\x89\x41\x32\x5d\x3d\x64\x04\x57\x23\x7d\xbe\x3a\xea\xeb\xeb\xd3\x36\x6e\xf7\xbf\x61\x7f\x1f\xb7\x40\x15\x69\x3c\x5c\x6b\x28\x73\x78\x00\xd4\xee\x48\x24\x14\x04\xf8\xfb\x97\x04\xe4\xec\xec\xec\xd8\x7f\xa9\x1f\xbf\x92\x0c\xfb\x6b\xf4\x97\x2d\x2d\x0e\x74\xf7\xf4\x90\xd8\x6f\xbb\x51\xf6\x9f\xcc\x43\xf8\xc1\xc0\x61\x32\x24\xf4\x9c\x77\x83\xb3\x88\x00\xe8\x45\x16\xa4\x32\x03\x21\x21\xa1\x53\x69\x08\x0d\x90\x5c\x7a\x91\xa2\xcd\xf4\xce\x5a\x57\x7c\x06\x25\x64\x5a\x34\x32\xe2\xe2\x90\xe7\x05\x7a\x27\x26\xae\x8b\xa9\x65\xeb\xf4\xbd\x66\xe0\x5d\x32\x2d\x60\x50\xeb\x26\x8f\x0f\x0e\x3a\x54\x59\xd7\xbb\x6d\x23\x09\x0f\x25\xed\x5c\x5d\x85\xc9\x36\xea\x60\xf2\xed\xfd\x8e\x0c\xd9\x94\x97\x24\x22\x51\x1a\x8a\x6c\x0f\x98\x69\x8a\xdb\x75\x0a\x7e\xc2\x00\x86\x16\x44\x7b\xa4\xf5\xc4\x2c\xa4\xf8\x9c\x5c\x5d\x53\x8e\x85\x04\x7e\xc6\xab\x81\x63\x69\x0a\xc4\x9d\x02\xf1\x32\xae\x8e\xdf\x0f\x47\x89\x79\x78\x78\x80\xfd\xfd\xfd\x32\x53\x53\x61\x40\x19\xaa\xe8\xb5\xfb\x68\xa8\xbf\x3b\x1a\x91\x1f\xac\xe3\x6a\x18\xc4\x08\x68\x03\x03\x76\xfc\xa9\x3e\x98\x1e\x73\xa9\x4c\x4b\xd2\x94\x1e\x3e\x75\x20\x59\x1e\xc0\x91\xc8\xee\x7d\x96\xe7\x1d\x3e\xa6\x4c\x89\x09\x6d\x39\x6c\xe0\x7b\x9a\x58\xc9\x5d\xf6\xec\x9c\x9c\xa9\xea\x21\x37\xfe\xfc\xa9\x32\xe7\x0f\xa5\x53\xb5\x21\x63\x5e\xd5\x9e\x1c\x00\x5c\xb5\xb1\xe1\x86\xfe\xee\x52\x6d\x51\x92\xd3\x3c\x7d\xfa\xb4\x8e\x8e\x0e\x7c\x86\x4a\x30\xf8\x81\xfa\x82\x00\x93\xb9\xd5\x1f\xc2\xa6\x8c\x0c\x73\xf6\xf7\xf7\xed\x16\x26\xa1\x50\xe8\xd4\x68\x7d\x44\xd9\x25\x6b\xe6\xc4\x0c\xd9\xeb\xf6\x27\xec\xc3\x74\xba\xe4\xff\xbd\x88\x90\x90\x10\xda\x46\xe0\x10\xa5\x32\x57\x4d\x1a\x34\x9f\x1c\xc7\x51\x7f\x35\xbc\x6b\x39\xd6\x7b\xf7\xd4\xcb\xcb\xab\x54\x13\xcb\x0e\x31\xbf\x7e\xdd\x1f\xea\x69\x09\xb8\x8c\xae\x66\xa4\xa4\x70\xf2\xf1\xf3\x73\x77\xfa\x1c\x1f\xed\x0f\xbf\xf7\xe5\x84\xda\x12\xed\xa7\xc8\xef\x9a\xbd\xa0\x9e\x28\x70\xea\xd4\xa9\xb0\xbb\xb9\xc1\x63\x1f\xcf\x5d\xc6\xac\x5c\xb8\x23\x43\x22\x91\x9e\x9b\x4b\xd1\x4a\x0c\xa4\xa5\x1f\x0c\xdc\x93\x67\x04\xc0\x2c\x4d\x94\x48\x9d\xac\xae\xac\x34\x44\x22\x2f\x16\x15\x17\x3f\xd4\xa2\xc0\xe5\xe0\x2e\x2e\x1c\xf5\x48\xdb\x67\x2d\xb1\xdb\x0a\x5e\xde\xde\xce\xca\x4f\xdf\x27\x23\x61\xd9\x19\x19\x9a\xb3\xd3\x9d\x4f\x9e\x70\x84\xef\x6d\x78\x4e\xaf\xae\xda\xbc\xb6\x97\xb7\x93\x3b\xfb\x5f\x32\x21\x1b\xa3\xd3\xab\x6a\x6b\x23\x0a\x86\x7e\xfe\x34\x82\xc1\x60\x26\xae\x9e\x47\x1b\xd3\x70\xcc\xa3\x2a\xab\xec\xe2\xe2\xe2\x00\xd7\xaf\xdc\x6f\x1b\x1a\xf2\x4e\x31\x33\x3f\xfe\xf8\x35\x9f\xa1\x75\x7d\x4f\x4c\x54\x42\x6c\x4d\x85\xea\xd3\x03\x00\x00\x16\xd7\x6c\x4d\x28\x46\x1e\xf7\xfe\x4f\x00\x00\x00\xff\xff\x4b\x19\x55\x86\xc4\x0f\x00\x00")

func libIconPngBytes() ([]byte, error) {
	return bindataRead(
		_libIconPng,
		"lib/icon.png",
	)
}

func libIconPng() (*asset, error) {
	bytes, err := libIconPngBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "lib/icon.png", size: 4036, mode: os.FileMode(420), modTime: time.Unix(1545238812, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"lib/bs.glade": libBsGlade,
	"lib/icon.png": libIconPng,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"lib": &bintree{nil, map[string]*bintree{
		"bs.glade": &bintree{libBsGlade, map[string]*bintree{}},
		"icon.png": &bintree{libIconPng, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

