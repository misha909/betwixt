// Code generated by go-bindata.
// sources:
// objdefs/ipso.json
// objdefs/oma.json
// DO NOT EDIT!

package betwixt

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

var _objdefsIpsoJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\xed\x6e\xdb\xc6\xd2\xfe\xdf\xab\x58\x18\x28\x90\x02\x8d\x22\xdb\x4d\xe2\xe6\x9f\x6b\xa7\xef\x2b\x20\x8e\x0c\xdb\x6d\xda\x1e\x9f\x1f\x2b\x72\x25\x6d\x4b\x72\x75\xf8\xe1\x0f\x14\x05\x7a\x21\x3d\x37\x97\x2b\x39\x33\xb3\xcb\x2f\x89\x94\x56\x8a\x68\x49\x16\x81\x20\xb0\x24\x6a\x39\xfb\x68\x76\xe6\x99\x99\xdd\xe1\x9f\x5f\x31\x76\xd0\xbf\x0f\x44\x78\xf0\x8e\x1d\xf4\x2e\xaf\xfb\x07\xdf\xd2\x5b\x83\xdf\x85\x13\x47\xf0\xe6\xbf\xe0\x25\x63\x7f\xd2\xff\xf0\x41\xcf\x85\xf7\x8e\x8f\xba\xdd\x6f\xd3\x77\x3e\x72\x5f\xa4\x5f\x66\xe7\x72\x24\x63\xee\xb1\x5e\x30\x49\xe2\x83\xec\x9a\x8b\xc4\x8b\xe5\xc4\xc3\xeb\xe2\x30\x11\xf9\xfb\x3c\x70\x79\xac\xc2\x47\xf8\x60\xc8\xbd\x28\xff\xe4\x5c\x44\x4e\x28\x27\xb1\x54\x01\x0e\xfe\x7f\x02\x44\x94\x0e\x73\xcd\xf8\x12\xc7\x67\x43\x15\xb2\x40\x05\x2f\xa3\x89\x70\xe4\x10\x3e\x8e\x44\x10\xa9\x30\xca\x6f\x7c\x25\x22\x95\x84\x8e\xc8\x67\x52\x9c\x4d\x3e\xa3\xd7\xaf\xf3\x19\x95\x66\x55\x9a\x10\xbb\x8e\x79\x2c\x0e\x4a\x17\xf6\x27\x22\xe4\x28\x26\xde\xe2\xe0\xaa\xfc\x61\x61\xde\xa5\xe9\xcd\x4c\xbe\x88\x4a\x49\xf0\x9b\xc7\x09\xc9\x31\x50\xca\x2b\x8f\x3d\x85\xd0\xcd\x58\x30\x27\x09\x43\x11\xc4\x2c\x42\x31\x99\x1a\x32\x5e\x06\xec\x20\xfb\xfe\x5f\xdf\xce\x07\xe3\xd0\x02\x8c\x33\x95\x04\x31\x28\x4e\x03\x70\xcc\x7e\x3a\x8d\x87\x0c\x62\x0b\x38\xfc\xc4\x03\x61\xee\x04\xbb\xe3\x5e\xa2\x11\x71\xe8\xb5\x46\xc8\x15\x31\x68\xb9\x70\xed\x71\x39\xb2\xc0\xe5\x52\x79\x3c\x94\xf1\xe3\x5c\x60\x3e\x35\x86\x8c\x95\xa6\x4c\x8c\x8c\x08\x49\x0c\xaf\xcb\xeb\x8a\x47\xa0\x39\x3f\xc0\x38\x82\x07\xec\x45\x97\xdd\xde\x26\xdd\xee\xb1\xcb\x3e\xaa\xd0\x87\x8b\x0e\xd3\xd7\x57\xe2\x4e\x84\x91\x70\xbf\xb1\x07\xf0\xd8\x02\xc0\x73\x31\x00\xd5\x72\xe6\x2f\xb4\xe6\x00\x9c\x51\xad\x9f\x02\x49\xa6\xf0\xc0\x8f\x16\x02\xeb\x1a\xd9\x19\x48\x2b\x95\x0b\x80\x32\x3f\xea\xd8\xe3\xf3\xda\x7e\xe1\x01\xfc\x91\x88\xe7\x81\xf4\xbe\x31\x8c\xd4\x84\xff\x27\x11\x73\xd1\x20\xe9\x48\xb9\x52\x79\x69\x11\xda\x42\xf1\xf6\x75\xb5\x0d\xba\x26\x13\xcf\x48\x90\x8d\x58\x9e\x28\x0e\x65\x30\x5a\xa8\x08\x31\x5c\x9e\xae\x2e\xed\x96\xd8\x0b\xf4\x57\x32\x00\xdb\x83\x0a\x72\xd9\xbb\xa2\x8b\x8a\x8b\xc7\xfc\xf5\xef\xaf\x0a\xf8\xcc\xb8\xde\xc3\xf9\xae\xb7\x9f\xc4\xcd\xfa\x5e\x45\x37\x98\x75\xbe\x60\x5a\x13\x1c\x70\x25\xf7\xfb\x7a\xbe\xfb\xd5\x93\xb2\xf0\xbf\xab\x9a\x85\xc6\x1d\xb0\x46\xcd\xda\x10\xd4\x68\xff\x14\x1e\xbb\xe6\x6a\x8a\x70\x34\xea\x68\xde\xd6\xe8\xd3\xe9\x64\xe2\x49\x87\xc0\x59\x6c\x42\x9a\x03\xce\xd2\x86\xf0\x82\xb4\x45\x7b\x62\x56\x20\xa1\xa7\x87\x62\x25\xd3\xf2\xf9\xef\x7f\x3e\xbc\x3f\xff\xfc\xf7\x7f\x97\x35\x2d\x47\xd5\xa6\xe5\x34\xe0\x9e\x1a\x35\x46\xea\xb9\x1e\x7e\xed\x9c\xfe\x4d\x0d\xa7\x2f\x4e\x87\x9d\x99\x25\xfb\x33\x39\xa7\x4d\x50\xfb\xa1\xa7\xf8\x94\x1b\xbf\xe2\xc1\x48\xf4\xc3\xf7\x41\xe2\xd3\xbd\xbb\x2f\x0f\xad\x6d\x4f\x46\x75\x51\x53\x8a\xd0\x5a\xb3\x90\x37\x35\xf4\xff\x02\xc8\xcc\x85\xe0\x51\x12\x0a\xb7\x39\xbc\x16\xaf\x9e\x0a\xc0\x2a\xf0\xf0\x65\x20\xfd\xc4\x37\x78\xf8\xa9\xdc\x83\xc7\xa2\x4b\x8e\x24\x71\x35\x75\x0f\xf4\xa4\xff\x91\xc1\x3b\x21\x91\x2a\x6b\xa4\xaa\x03\x82\x0b\xfe\xb0\x4b\x48\xf1\x87\x27\x40\xaa\x9a\xf9\xa3\x4e\x91\xb6\xef\x00\x4c\x25\x85\x8a\xc7\x3c\x66\x0e\xb8\xad\x41\x1d\x64\xf6\xd0\x7c\x57\xab\x44\xbb\x02\x4d\x49\x83\xd6\x09\x4d\x75\x3c\xa4\x63\x0b\xd4\x1d\x98\x0c\x9b\x5d\x6d\xd1\x0e\x44\x45\x73\xa4\x67\xb1\x9a\xf2\x4b\x2d\xed\xc9\x2d\x11\xfc\x4b\x43\x8d\xb9\x14\xe8\x54\x86\xec\x12\x8c\x14\x22\x5b\xe6\x42\xfb\x16\x78\x56\xc6\x9d\xcb\x72\xc3\xe3\xb9\xdc\x70\xed\x51\xe7\xcd\x58\x46\x8c\xee\xa0\x28\x17\xcd\x24\xfe\xd4\x23\x43\x19\xcd\x7b\x45\x4b\x93\x40\x7c\xc0\xee\x65\x3c\x86\x55\xf5\xc8\xfe\x90\xb0\xb4\x30\xe2\xd0\xe2\x19\xc6\x2c\x31\x0f\x31\xe4\x8e\xe8\xac\x44\x27\x5b\x17\x56\x0b\x4d\xeb\xc2\xea\xa0\xa9\x33\xc9\xc5\x75\xb3\x44\x14\xd2\x58\x82\x63\xcd\x61\x48\x96\x02\xc9\xc3\x90\x24\x8d\x5c\xad\x23\x91\xb5\xf8\xb3\x4d\x1a\xe4\xde\x90\x4d\x90\x27\x07\xda\xe7\xd7\x79\x36\x2b\x7f\x06\xe8\x86\xbe\x42\x5c\x97\x8e\xec\x8f\xeb\xea\x75\x69\x08\xae\x9d\x5b\x93\xe6\x9b\x7b\x9e\xba\xd7\xb5\x86\xfc\x4a\x9d\x13\x4a\x8d\xba\x5e\x58\x1d\xd6\x23\x63\x3f\xe0\x68\xd0\x11\xa8\xaa\xef\xe8\x65\x8a\xfc\x89\xb3\x24\x90\x30\xbe\xe3\xa8\xd0\x45\xdc\x80\x3c\xe1\x57\x7e\x3a\xfb\xe9\x82\xa5\x59\x04\x82\xbc\xc3\x6e\xc6\x49\x44\xee\x21\x05\x5f\x0f\xe3\x8a\xa1\x0c\x8c\xf7\x90\x78\x43\xb8\x7d\xe9\x9b\xa9\x29\x08\xc5\x44\x85\x31\xde\x24\x89\xe8\x56\x78\xa5\x9e\x60\x87\xdd\xde\x06\xd7\x69\xd2\xc2\x4c\x1a\x7f\x42\x98\x9f\xbc\x13\x01\x0b\xc9\x0e\xc2\x3d\x4d\x42\x03\xe7\xa8\xa7\x35\x80\x5b\x7b\x1c\x33\xe4\xd2\xcc\x56\x39\x89\x8f\x3a\x23\x02\x3e\xf0\xcc\x9c\xa4\x0b\xef\xc8\xa1\xb6\x41\xa9\xfc\xe9\x58\xae\x0c\xe1\x7e\xde\x23\x1b\x86\xca\x67\x32\x8e\x98\xae\xe0\xb2\xde\x39\xce\x3a\x93\x12\xcc\xdc\x63\xe6\x2c\x79\x85\x47\x95\x94\xa4\x13\x2e\x4e\x9c\xe0\x47\x69\x80\x91\x06\x0a\xa4\x79\x90\x91\x9e\xe7\x4a\xde\x73\xe7\x92\x0a\x59\xdd\xe7\xdc\xe8\x07\x38\x00\x58\x87\xf4\x36\x2c\x41\x0c\x80\x69\x88\x4e\x9b\x8b\xd8\x38\xc0\x6d\x0a\xe3\xc9\xd4\x77\x7f\x69\x63\x53\xfa\xda\x26\x4c\xea\xd1\x6a\x38\x61\x52\x53\x22\x30\xc1\xfe\x56\x15\x05\xd6\xa3\x7d\x1f\x78\x14\xa3\xe9\x4b\xe1\x2a\x23\xa9\xf9\x03\xa2\x7d\xbd\x94\xde\xbd\xad\xf1\xee\x06\x47\x2d\xfa\xb6\x2c\xe3\x29\x44\x0c\x02\x44\xb8\x48\x50\x46\x00\x4b\x22\x7e\xa2\x33\xea\x20\xca\x67\xc2\x43\x8c\x91\xd0\xdd\x08\x9f\x64\x87\xef\x20\x61\x83\x4f\x22\x99\xd8\xef\xf1\xd8\xfd\x1c\x9d\x65\x50\x63\xbc\xee\xbc\x90\xe6\xac\x7f\xb4\xdf\x99\xb9\xb5\x6c\x09\x39\xae\xdb\x12\xd2\xf3\xbc\x04\x7c\x37\x2f\xee\xa5\xfa\xf2\xd0\xae\x30\x6a\x3a\x0d\xf1\xc0\x7d\x18\x92\x22\xb1\x28\xab\xe2\x7b\x0f\x6d\x9c\xd0\xc6\x09\xdb\x05\x70\x1b\x27\xb4\x71\xc2\x2e\x21\xda\xc6\x09\x6d\x9c\xb0\x31\xed\xab\xdc\x52\x44\xce\x5f\x45\xb8\xa3\xcf\x24\x6c\x77\x3f\x48\x58\x9e\xf4\x96\x13\xb0\x59\x06\x19\xc8\x6e\x96\x7b\xc6\xf9\x2c\x1d\x4a\x38\x53\xa1\x84\x15\xfd\xab\xd9\xb6\x77\x49\xf2\xae\x95\xfb\xa5\x43\xa6\xf3\xa6\x7a\x6b\xba\xa7\x13\xdf\x43\x96\xaf\xe8\x72\x78\xc3\x15\x1e\x7f\x64\x13\x1e\x82\x54\xb1\x68\xcf\xe5\x94\x6a\x62\x93\x32\x94\xed\xd9\x9c\xea\xb3\x39\xed\xd9\x89\xf6\xec\x84\x55\xa0\x3c\x17\x91\xef\x6b\x38\xf9\x0f\x49\xf4\x48\xa4\xc0\x13\x3c\xd4\xf6\x6a\x43\x79\x97\x95\x4f\x1d\x9d\x93\x91\xcd\xf2\x76\x7a\xdd\xa0\xbf\xd1\xeb\xc8\x54\x43\x1d\x9a\xa0\x7e\x8b\xce\x24\xd9\x23\x57\xcd\xbd\x35\x62\x30\x3a\x41\xf8\x1c\x90\x2b\x42\x64\x50\x1b\xe0\xdc\xaa\x41\xb3\x72\xcc\x35\x7b\xa6\x0a\x6e\x7f\x8d\xbe\xb9\xf0\xf2\x1d\x9b\xa9\xbf\x47\x63\x95\x78\xee\xd4\x46\x29\x16\x17\x08\x88\x59\x69\x30\x75\x5d\xe6\x9e\xfa\xd8\xcf\xd3\xa4\x1d\x86\x35\x7a\x90\x45\x81\x17\x53\x77\xd2\x05\x66\x9d\xb2\xca\x88\x88\x8d\x89\x04\x5f\xa5\xf1\x4b\x16\xad\xdc\x69\x1e\x8e\xfc\x3c\xce\x23\xc6\xec\x3a\x5d\x24\x9f\x17\xe7\xcc\x4a\xdc\x61\xa7\x41\x96\x81\x2a\x48\xa9\xf7\x05\x50\x95\x7d\x14\x0a\xb8\xa9\xc9\xd4\xb2\x17\x09\x78\x9f\x77\xf0\xea\x9b\xd5\xb6\x81\xb5\x09\xaa\x36\x41\xd5\x26\xa8\x76\x12\xdd\xed\x4d\x50\x35\xc7\x37\xdb\x04\x55\xf3\x18\x5b\x10\xdb\x36\x41\x95\x5e\xd2\x16\xb2\x67\x71\x5c\x8b\x8a\xad\xbf\x92\x6d\xc5\x72\xbf\xab\x66\xb9\xff\x9f\xf8\xd2\x2d\x1e\x4e\xde\x0c\xc5\x1d\x1b\x31\xaa\xf8\x6d\xf6\xd9\x96\x90\xdb\x29\x59\x17\x32\xdb\x50\x98\x1c\x4a\xf6\x45\xaa\xb6\xc3\x8f\xea\xc0\x65\x1c\xee\xa6\xb9\xee\xd7\x2d\xd3\x6d\x99\xee\xb6\x01\xbc\xef\x4c\xb7\xca\xa0\xef\x05\xd5\x7d\x16\x0a\xbb\xb7\x54\xb7\xad\xc5\xd6\x21\xd3\x52\xdd\x67\x4f\x75\x5f\x57\x53\xdd\x4b\x72\x50\x05\x99\xd6\xc8\x79\x17\xd1\x5c\x75\x27\xf0\x90\x91\xf6\x91\x45\x92\x58\xc1\x77\xe1\x13\x15\x8b\xd9\x6b\x9b\xe6\xbc\x38\xc2\x40\x21\x1d\xd7\x25\x3e\xbc\x3a\x14\xe6\x05\x49\xd3\x61\x3d\x6f\xae\x00\x85\x82\x21\x9e\x58\x1a\x3d\x82\xf1\xf5\xe4\x40\xeb\x63\x76\x7b\x3d\xb3\x21\x8c\x8c\xfb\x12\x56\xa1\xbc\xad\x2d\x5e\xc6\x16\x9f\xd4\xd8\xe2\x1e\x55\xea\xe0\x9f\x50\x78\xe2\xae\xf0\x43\x6f\x9d\x65\xfe\x64\x5d\xbe\x2f\x4d\xc3\x1a\x20\x8b\x18\xaa\x69\x7c\xbe\x8c\x38\x2d\x06\x28\xe5\x9a\xc5\x79\xcc\x67\xf4\x3a\x7c\xed\x7f\xb4\x87\xd1\x22\x52\xda\x7d\x1c\x8d\xd1\x6c\x12\xc7\xfa\x80\xc8\xd0\xf7\xdd\x47\xb1\x4a\x1b\xd7\x47\xd7\x4f\x16\x06\x40\x3b\x0f\x60\x95\x1a\xae\x13\xc0\x6a\x1f\x7b\x96\x7b\xf8\x2d\x47\x70\xbc\xcc\xf6\xa6\x12\x8a\x7a\xd5\xd2\x36\x4a\xe4\xf4\xb3\xa4\x86\xd2\x1b\x8c\x76\x14\xde\x49\x87\x36\x43\x85\xd6\xe9\x8e\x93\xee\x9b\x4a\x64\x4f\xb5\x08\x9a\xa5\x9e\xe5\xb4\x69\x1e\xb4\xcd\xed\xd7\x58\x5d\x39\xaf\x04\xb0\x1d\x80\x0d\xb4\xb0\x84\x6a\x91\x09\x82\x4a\xde\x87\x32\xd6\xa7\xf3\x8b\x9b\xcb\xb2\xab\xf0\xb4\xbd\xe2\xf6\x1b\xcb\x4e\x0e\x6d\x38\x4e\x99\xce\x6e\x9f\xd2\xfe\x7c\x7a\x65\x4d\x74\xa6\xe6\x62\x8d\x93\x05\xd5\x79\x0e\x30\xa5\x0e\xa6\x3c\x97\x35\x3b\xea\x43\x0b\xc2\xf3\x2c\xc0\x4c\x03\xc5\x46\xc1\x5c\x94\x06\x7e\x16\x48\x36\xcc\x7b\x0e\x17\xf1\x9e\x67\x01\x62\xb5\x3a\xae\x13\xc6\x85\xec\x67\x07\x70\x5c\x8a\x01\x4d\x61\xd9\x30\x07\x3a\xac\xe6\x40\x57\xa9\x10\x3b\xc0\x82\x16\x29\x6a\xc6\x83\xa6\xa1\x6d\x98\x08\x1d\x55\x13\xa1\xcb\x42\xd2\x6d\x6b\xb4\x75\xf6\xf8\x8e\x39\xa7\x3e\xf0\xc4\x4c\xaa\x30\xdd\x6c\x9e\x12\x20\x07\x24\x4e\x7c\xfa\xf2\x12\xe8\x54\xd3\x9f\x34\xaf\x66\xa9\x71\x0d\x6e\x94\x5e\x8c\x12\x2c\x12\x17\x17\xde\x27\x50\x1c\x51\x46\xa4\xa0\x59\x8e\x12\xc3\xa1\x74\x24\xe6\xb9\xad\xc1\xa9\xa6\x33\x3a\x3b\x79\x36\x6d\x02\xb6\x3c\x9b\x4a\xa9\xec\x99\x10\xef\x55\x2d\x6d\x36\x7f\xcd\x2f\x2d\x64\x66\x6b\xaa\xbf\x26\xb5\x86\x2b\x2a\x4d\x33\xbd\x35\xf3\x9e\x5f\x54\x30\xa0\x3a\x01\x4f\xef\xcd\xa2\xc4\x19\xe3\xde\x92\xfe\xc7\x57\xfd\x1f\x7f\xa4\x09\x63\x27\x36\x30\xd1\xd2\x07\xde\x31\xea\xb0\x53\xe6\xa3\x58\x2f\xcd\x19\x27\xdd\xd6\x10\x3d\x26\x25\xf5\x07\xa2\xd0\xeb\xac\xd0\x11\xc2\x34\x27\x83\x7f\x49\x24\x86\x89\x87\x37\x07\x7f\xea\xb2\x89\xf4\x54\xcc\xee\x65\x08\x63\x85\xae\x08\xa3\x52\xef\x88\xce\x6c\xb9\x82\x67\xf5\x02\x3d\x81\xa1\xa7\x1b\x84\xc2\x4b\xe9\x9b\x72\xad\x3e\x1f\x41\x3e\x65\x0c\x42\x0c\x84\x80\x99\xdd\xcb\xd8\x19\x53\xab\xb3\x55\x0a\x06\xbb\xdf\xcd\x03\x9d\xf5\x69\x4d\x0b\x0f\x03\xd6\xb0\xd0\xe2\xe1\xf3\xdf\xff\x5c\x28\xba\xf0\xcc\x53\xa6\xb5\xae\xbd\x8d\xac\x41\xab\x1f\xbc\xea\x0f\x87\x8d\x60\xf4\xe5\x27\xf8\xb4\x70\xb7\xb7\x01\x36\xbe\xb3\x9e\x68\xdd\x29\x3d\xdf\x5f\x40\xe9\x1a\x3d\x28\xc3\x8a\x2d\x3d\xdf\xb1\xee\xcb\xc3\x6e\xb7\x86\x84\x7c\xbd\x40\x6b\x64\x5e\xa1\xc3\x02\xa3\x3e\x1b\x8b\xeb\xd0\x93\xa3\x71\x4c\x86\x01\xc9\x9e\x88\x89\x82\xdc\x8f\x25\x58\x10\x5c\x75\x60\x12\x7a\x41\x2c\x46\xe9\xa1\x36\x58\x86\xf1\x3d\xae\xc4\x2e\x55\xa7\x40\xa0\xe9\x4d\x6c\x4b\xa8\x57\xb5\x97\xe9\xc3\x1a\x04\x23\xb0\x39\xd8\x6b\x20\x8e\x84\xb3\x70\x69\x92\xf5\x82\x38\x15\xae\x55\x81\x1b\xd5\x1b\x32\xec\x6f\xf9\xc9\x10\x3e\x9e\xd3\xbd\xae\xa6\xd3\xfa\x5c\xb2\xa3\x4f\x13\x2e\x01\x68\x4d\xe0\x9c\x64\x96\x7e\xaa\xf1\xf3\xd3\x42\x6b\x61\xde\x4e\xd3\xde\x43\xc6\xff\x68\x78\xb4\xec\x54\xd8\x36\xd2\x30\x4f\xdc\x89\xec\xe1\x46\xa9\xbf\xbb\x24\x37\xf4\x09\xdc\xd0\xb2\xee\xfc\xa4\xda\x9d\x5f\x0b\x7c\x6a\x8f\x5c\xeb\x0e\x81\xa5\x76\xc5\x92\x83\x8d\xa9\xc9\x67\x24\xb3\x02\x3e\xbe\xcd\x91\xf6\xc6\xa1\xf2\x3c\x5c\xb7\x06\x00\x8e\x7a\x63\xda\xce\x96\x3b\x89\x52\x6f\x52\x72\xbc\x30\xde\x04\xa7\x84\x83\xc0\x6d\xc4\x03\x1a\x03\xbc\x97\x6e\x07\x94\x9e\x8d\x37\xcd\x4d\x67\xfb\xb2\xe2\xf7\x7c\x1e\xe3\x0d\x81\x2e\x44\x91\x72\x24\x11\x91\xbc\xdf\x7b\x71\x5b\x03\xc9\x8b\x7c\x83\xc6\xe0\x5e\x6e\x85\xb0\x43\xb8\xeb\xe6\x73\xd4\x1a\xef\xc1\x87\xba\x13\x78\xda\xa9\x75\x15\x37\x5f\x53\xc6\xdd\xee\x4d\x32\x1b\xea\x59\x30\xd7\xa4\xbc\xad\xa9\x51\x9c\xd1\x0f\xb5\x03\x76\x24\xf3\x75\x45\x43\x8b\x06\x5a\xf9\x5a\xdd\x60\xf1\x4c\xb8\xb3\x67\xcf\x6c\x78\x2a\x06\xf9\x7d\xed\xe6\x3f\x63\x57\xb7\xac\x85\xba\xf5\x06\xc0\xdb\xdb\xc5\x4d\x13\x32\x43\xab\x6d\xe0\x92\x2e\xe9\xb0\xa6\x07\xf8\x07\x05\x41\xff\x99\x36\xfd\x8d\x7a\xa5\x7e\x16\x6c\x92\x27\x1a\x92\x0d\xf2\x61\xbc\x97\x80\xc2\x04\x7e\x1d\x41\xf9\xa8\xd4\x0d\x11\x1d\x54\xe8\x7e\xca\x6f\xe3\x36\xa6\x24\x56\xbe\xd9\xf9\x55\x50\xbc\x17\xd8\xa5\xda\x93\xbe\x34\x61\x2c\xc5\xe1\x2b\x1d\x7d\x38\x39\xaa\x26\x3e\xef\xef\xd0\x98\xf6\x74\x2b\x6e\xd9\x10\x93\x5f\xa8\x67\x96\x6b\x52\x90\xb0\x32\x13\xb6\x14\x78\xdb\x67\x6e\xaa\xd3\xfe\xd7\x98\x94\x6d\x8e\x56\xaf\x05\x03\xa4\xce\xf7\x63\xa1\x79\x47\x49\x87\x34\x34\xf7\xd2\xf3\x74\x76\x59\xff\xbf\x44\xb7\x91\x93\xa3\x9a\x6e\x23\x89\x49\x92\xf5\x02\xdc\x70\xb7\x3d\xd8\x64\x86\x68\x46\xaa\xaa\xa7\xd8\xa6\xb3\x48\x7b\x3e\xcd\x60\xb7\x04\x50\x35\xfe\x1e\x43\x15\xcc\x2a\xc6\x8f\xec\x03\x52\xef\xed\xf6\x6f\x4e\x41\x5c\x03\x8a\xc6\x81\xb1\x9b\xdc\xcd\x85\xc2\x11\xf2\x2e\x4d\xb9\x17\x94\x8c\x52\x83\x64\xb9\xc8\x62\x85\x6a\x12\x4a\x8a\x40\x78\x34\x46\x9c\xd1\x16\xe6\xee\x72\x09\x70\xdf\x56\x73\x87\xbb\x91\x36\xeb\xa7\xee\xef\xec\xd2\xd9\x6c\x78\xb6\xb6\x9c\x83\x76\xa1\x9a\xad\xa6\x75\x3a\x53\x32\x4a\x22\x3e\xca\x38\x87\xfe\x15\x28\xe4\xa3\x5f\x65\xfa\x4c\x5c\x89\x9a\x50\xb3\xd2\xa3\xb7\xf8\x84\x03\x7a\xe6\x68\xd5\xc0\x4b\xfc\x1e\x27\x35\x56\x01\xd4\xe6\xec\xd1\xf1\x36\xcb\xe2\x1a\xf9\x25\x5c\x9c\x9b\x83\x73\xcb\xb4\xb8\xc2\xd4\xca\x0e\x9a\x62\x1e\x4f\xff\x0e\x3a\x37\x5b\x5c\x39\xe6\x47\xc3\x60\x0e\x1f\x19\xa2\x3d\x39\x84\x95\xe5\x52\x8c\x15\xe1\xa9\x69\x8b\xfb\x81\x12\x54\x1b\x60\x3c\x30\x91\x8c\xdb\x98\x34\x99\x09\x5d\xf3\x70\xfb\xc3\xfb\x73\x8c\x78\x0d\xef\xc1\x6b\xd2\xcd\xf9\x00\x46\x9e\x5e\xd3\x98\x40\x10\x16\xe8\x87\xa4\x50\x0d\x6b\x48\x94\x89\x62\xee\x72\xf6\x4d\x5f\x9d\xdd\x1b\x87\xf9\x7a\x36\xef\x46\xc7\x4d\xb3\xe6\x6d\x26\x72\x4e\x87\x28\x84\xfb\x5a\x97\xf4\x98\xe9\xbc\x64\xa0\xcb\x07\xe9\xd3\x5a\x28\xb5\xa0\x87\x58\x67\xbc\xbd\x13\x67\x53\xda\xb0\x7b\x07\xc2\xee\x7d\xd8\x85\x1a\x2b\x6c\x7c\xa7\x0b\xdd\x00\xd5\xa7\x71\x9e\x3d\xd6\x56\x04\x93\xc7\xb8\x82\xf7\x62\x0b\xc0\x4d\x4d\xd1\x5f\x1b\xd9\x67\x5d\xc2\x5a\x58\xad\xd1\xe5\x19\xb3\xc1\x29\x4b\x03\xbb\x66\x7f\x93\xf6\x00\x00\x97\xbe\x4c\x16\x9f\xda\xae\x17\xe2\x8b\xc3\xee\x37\x18\x6a\x81\x71\x3a\xc4\xcf\xcd\x91\xac\x2e\xfd\x3d\x1c\x2e\x81\x6e\x5b\x37\x6b\xeb\x66\x8d\xd6\xcd\x72\xcb\xf7\x25\x65\x33\x2b\xfe\x59\xd7\x97\x57\xef\x40\xdb\x30\xff\x34\x9b\xf1\xa6\xf9\xe7\xb5\x8f\xe9\x90\x4b\x2f\x19\x4d\xf1\x4e\x7d\x7d\x48\x3d\x13\x1b\x61\x9f\xab\xa4\xe9\xf6\xd7\x89\x67\x0a\x4c\xbf\xac\xd9\x8a\xba\x37\xdb\xf9\x6a\x7d\xf9\x52\x00\x3c\x33\x57\x5e\x5c\xa1\xbb\xef\xd0\x9f\xd0\xb1\x54\x3c\xe7\x76\xfd\xee\x5d\xff\x3a\xad\x7b\x6f\xd8\xbd\x17\x17\xc1\x93\x38\xf9\xba\xe7\xa2\x3b\x8e\xf0\x44\xa8\xa8\xf3\xfd\x93\x6d\xf6\x28\x3e\x0b\x5d\x37\x7e\x30\x89\x08\xce\x0e\x5f\x1e\x33\xfe\x80\x4b\xbd\x28\xd9\x3e\x3c\x08\xbd\x6d\xaa\xb4\xc5\x88\x36\xd4\x54\x69\xa7\x13\x88\x4f\xfd\x44\xc2\x9a\xe6\x0a\xbf\x3c\xc7\x5e\x4a\xa4\x74\xa5\xe6\x35\x10\xe2\x28\x53\xb1\xfb\x85\x2c\xe4\x12\xc0\x55\x1b\xc2\x5f\x9f\xef\x72\xad\x43\xee\xd7\x65\x91\xab\x36\x78\xbf\xed\x1f\x72\xbf\xcd\x20\x67\x45\x3b\x6a\x9a\xae\x5e\xf0\x51\x20\xe2\x2d\x64\x1d\x7e\x41\x30\xdd\x9a\xb5\x50\x69\xf2\x27\x3c\x4a\x1f\x62\x2f\x57\x3c\x88\xd1\x1a\xfc\xd6\xe0\x57\x61\xdb\x1a\xfc\xd6\xe0\x6f\x15\x72\xb3\x06\x7f\x01\x72\x35\x39\x4e\x63\x35\xcf\x53\xab\xb9\x35\x18\xce\xa6\x53\x8e\xdf\xd4\xa5\x53\xce\x85\xc5\xde\xab\x69\xff\xb0\xac\xaf\xac\xe9\xda\xf8\x03\x7f\xea\xf0\xbc\xb2\x43\x79\xc0\xb8\x0c\x69\xc7\x40\x54\xf3\x18\x9e\x41\x2a\xe8\xb6\xf4\x29\xcf\x05\xb2\x6c\x54\xfe\xc7\x25\x37\xad\xc8\xe1\xaf\xb6\x19\x79\xdb\x8c\x7c\xdb\x00\xde\xf7\x66\xe4\x6d\xda\x6c\xb7\x10\x6d\x7b\x91\xe3\xc8\x1b\xef\x7f\xdb\xf6\x22\x6f\x7b\x91\x6f\xc7\xa3\xa1\xbf\xc2\xbf\xfe\xfa\x5f\x00\x00\x00\xff\xff\x6f\xfe\x86\xa3\xd0\xad\x00\x00")

func objdefsIpsoJsonBytes() ([]byte, error) {
	return bindataRead(
		_objdefsIpsoJson,
		"objdefs/ipso.json",
	)
}

func objdefsIpsoJson() (*asset, error) {
	bytes, err := objdefsIpsoJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "objdefs/ipso.json", size: 44496, mode: os.FileMode(420), modTime: time.Unix(1456661427, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _objdefsOmaJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xdc\x5b\x5b\x93\xda\x36\x14\x7e\xcf\xaf\xd0\x30\xd3\x99\x76\x26\x64\xb8\xef\xb6\x6f\x5c\x92\x74\x27\xc0\x52\x58\x36\x0f\x9d\x3e\x08\xfb\xc0\xaa\x31\x12\x95\x64\x36\x9b\x4c\xfe\x7b\x25\x2f\xb0\x38\xc8\xf8\x86\xb1\xd9\x17\x06\x6c\x6c\x9f\xef\xe8\xe8\x9c\xef\x3b\x92\xbf\xbf\x41\xa8\x74\xfb\x48\x81\x97\xfe\x50\x5f\x06\xed\xd2\x5b\xef\xc8\xec\x5f\xb0\xa4\x50\xc7\xfe\x56\x3f\x11\xfa\xee\x7d\xaa\x13\x37\xb6\x3a\x56\x79\xbb\xfd\x39\xc4\x4b\xd0\x17\xf6\x3f\x0f\x6a\x03\x34\x01\xcb\xe5\x44\x3e\x95\x76\xe7\x07\xae\x23\xc9\xca\xd1\xff\x91\xdc\x85\x97\xe3\x98\xda\x58\x32\xfe\xf4\xf3\x89\x31\x08\xe6\x72\x0b\x5e\x1e\xbd\xff\x78\x83\x09\x46\x33\xf8\x1a\x38\x9a\x8e\x6f\x4a\xbe\x7f\xdd\xae\x80\x63\x49\x18\xd5\x37\x2f\x0d\xfd\x27\xf7\x2c\x9d\x63\x47\x80\xff\x64\x80\xb9\x3e\x93\xef\x9e\x56\x9e\x11\x42\x72\x42\x17\xfe\xbb\x8f\x31\x5d\xc0\x2d\x7f\x4f\xdd\xa5\xf7\xf0\x4a\xb9\xd6\x6c\xa2\xd9\x93\x54\x40\x7d\x7f\xbc\xc7\x8e\x0b\xea\x83\x78\xcf\xf3\xec\x84\x47\xef\xea\xdd\xc1\x5f\x2b\x48\x5d\xfc\x5b\x69\x77\xd9\x8f\xb7\x47\x1c\x55\x35\x3a\xaa\xc3\x98\x54\x76\xe2\xd5\xc6\x59\xb9\x38\x6a\xc6\x98\x13\x11\x45\xcd\x88\x62\x1b\x6f\x68\xc0\x6c\xc8\x05\x02\xa1\x32\x6c\xa0\xeb\x49\x06\xb8\x1e\x75\x78\xeb\x46\xc7\x8c\xdc\x99\x43\x2c\xf4\x09\x9e\x10\xe3\xe8\xc6\x06\x2a\xf7\xa7\xe5\x39\x3d\xc4\x56\xf8\x3f\x17\x22\xa2\x69\x04\x0c\xb3\x37\x9f\x2f\x15\x54\x33\x28\x76\x39\x48\x8d\xa6\xf8\x08\x5a\x66\x04\x83\x09\xba\x88\x19\xa8\xb2\x65\xc6\x49\xf6\x2a\xd0\x41\x1d\x42\x6d\x55\x0d\xbc\xa8\x1d\x61\xae\xce\x48\xe0\x22\xcf\x21\x3f\xea\xac\x56\xe4\x9a\xd4\x07\xba\x90\x0f\x2f\xfe\x6a\x45\xf5\xd5\x75\xa8\xaf\x5e\xa6\x46\x71\x1d\x55\xaf\x95\x1b\xd7\x49\x0b\x78\xbd\x86\x1a\xd7\x51\xfd\xf5\x7b\x38\xd3\xd1\xce\x1b\xba\xcb\x59\x4e\x75\x5c\x4f\xc1\x88\x64\xc4\x4c\xdb\x26\x0f\x8c\xcb\x2d\x98\x9b\x5e\x16\x20\x0e\xcf\xc6\x4e\x24\xd5\x72\xab\xd9\xac\x27\x48\x25\x55\xe4\x5d\x18\x99\xb1\x99\x29\x5b\xd7\x21\xaa\xe2\xa1\x3f\x99\x63\xa3\xdb\xf9\x1c\xdd\x91\x65\x41\x32\xee\x94\x12\x4f\x2a\x94\xc4\x1e\xc2\xcd\xb7\x7f\xde\xec\xe1\xf5\xab\x88\x6a\xa0\x8a\xf0\x31\xd2\x73\x6b\x88\x18\xc1\x38\xce\xad\xa8\x9d\x2d\x16\xcd\xc9\x87\xcc\x41\x86\x84\xdf\xf8\x73\xbe\xf1\x17\x5f\x4d\xf4\x60\x8e\x95\x91\x68\x40\x28\x59\xba\x4b\x34\x02\x4e\x98\x9d\x09\xc6\x04\xd9\x28\x36\x48\xb3\x32\xd8\x81\xc4\x5f\x5f\x03\x48\xb3\x60\xe8\x11\x81\x67\xce\xd1\xf0\x7c\x9f\x19\xa8\x8d\xfc\x4f\xa3\x0d\x36\xf6\x7b\x29\x9e\xb9\xf2\x62\x47\xc7\xac\x1b\x86\x4c\x92\x39\xb1\x3c\x0c\x68\xa2\xcc\xd1\x9c\xef\xf3\x03\x50\xb4\xc1\x6d\x6b\x7d\xa7\x4a\x9c\x43\x68\x4e\x29\x26\x46\x67\xc2\x4c\xfd\x37\x54\x36\x1f\xeb\xa3\x34\xa0\xee\x1e\x00\xad\x98\x10\x44\xc7\xd9\x5a\x17\x0e\x81\xd8\x1c\x6d\x6f\x85\x30\x07\xe4\x10\x21\xd5\x60\x10\x8a\x9a\xef\x6a\xef\xaa\xef\xaa\xa9\x18\xfe\x18\x16\x44\xb7\x9c\xbc\x61\x9f\xae\x14\x1c\x15\xe0\x9c\x2c\x16\xc7\x59\x6b\xd2\x89\x1a\xd5\x4b\x71\x18\x4b\x2d\x80\xb1\xb4\x2d\x45\x36\x04\xea\x32\x2a\xb9\x0a\x9c\x58\xcc\xc5\x87\x22\x25\x75\x79\xee\xda\x16\x9c\xb4\x34\x92\x92\x96\x46\x3a\xd2\xb2\x75\x0e\x15\x12\x53\x15\xe1\x45\xf5\x52\x25\x29\xb5\xab\xc4\xa3\x76\x66\x12\xd4\xee\xf6\x13\x26\xad\x03\xe4\xa7\x96\x5f\xad\xf2\x8c\xc8\xf8\x6e\x29\xd7\x6b\x57\xad\x6b\xa4\x3f\xaf\xd2\x75\x55\xfd\xd3\x1c\x3d\xaf\x9a\x14\x83\x02\x67\x19\x44\xe1\x79\xb1\x7e\x90\x17\x7b\xb0\x26\x16\x18\x33\xa1\x3f\xe1\x65\x24\xe2\xd4\x6d\xdd\x39\xb6\xa4\xcb\x43\x86\xa8\x18\x34\xd0\x9c\xb1\x74\x4f\xd5\x89\xd0\xd8\x29\x08\x88\xa0\x35\x1a\x4e\xf0\x05\xa1\x30\x4f\xfd\x0f\x84\x2f\x1f\x35\x27\xba\x07\x2e\x94\xc5\x17\x00\xc4\x2c\x8d\xc6\xa0\xb8\xed\x51\x45\x71\x3e\xc2\x15\x5f\x18\x7d\x50\x53\x5a\x3d\x4d\xd3\x54\xc8\x04\xc4\x89\xc7\xc0\x2c\x80\xda\x6b\x4c\x1c\x4f\xe0\x8d\xd8\xa3\x6e\xdf\x6e\x12\x5d\xa2\x98\xca\xba\xec\x56\xca\x57\x49\xaa\x48\xe4\x52\x6b\xd6\x4f\xfb\x8e\x41\xf7\xcc\x91\x78\x71\x5c\x0a\x9e\xcd\x3f\x3b\x11\xbc\xbc\x4f\x25\x87\x7c\x08\xbb\x2e\xe7\xf0\xf3\x93\xf2\x47\xd8\x4e\xb5\x44\xd1\xc1\x52\x82\x9a\xac\x7d\x58\x83\x93\x4f\xc2\x8c\x10\xdd\xd5\x4a\x25\x00\xfe\x2f\x49\xe2\x5e\xdd\x2e\xb2\x4c\x09\xa0\x2e\xb0\xd4\x29\xee\x03\x87\x84\x11\x9f\x5d\xdf\xe7\x53\x27\xdd\x12\xc6\x7b\xce\x19\x57\xfc\xf9\xf8\x52\x71\xc2\x48\x3f\xe5\x32\x95\x99\xca\x78\x55\x07\x45\x03\x51\x90\x02\x54\x35\xd3\x99\x4d\xbe\x09\x5d\x43\xca\xb0\xbb\xe8\xad\x1f\x44\x04\x61\xa6\x32\xd3\xbb\xae\xee\x12\x86\x30\x81\x0c\x21\xc4\x1b\x09\x33\xa3\xd1\x23\xf0\x8d\x65\xd4\xe7\x0c\x82\x20\x92\x62\x08\xd8\x08\xe2\xae\x56\x8c\xeb\x56\xe1\x76\x05\x5f\xd9\xe0\xed\x08\x49\xc8\x69\xce\xdf\xd9\x6b\x1c\x28\x58\xa5\xf2\x29\x58\x92\xac\x9f\x77\xb7\xa8\x14\xc8\x7c\x5d\xd5\x68\x92\xf6\x94\xdd\xbd\x21\xc8\x47\xc6\xbf\xa0\x0e\xe0\x8c\x54\xed\x29\x53\x68\x08\xf3\x3d\x05\x98\x73\x95\x83\x80\x6a\x80\x6d\xc2\xd0\x84\x2c\xa8\x92\xb7\x13\xc9\xbd\xfd\x30\xc5\x68\x29\xee\x0a\xb6\xdd\x59\xa6\x92\xbd\x7d\x42\xbf\xa0\xbf\x5c\xc5\x6f\x8e\x6f\xb4\xcb\x96\x8c\x44\x43\x60\x2e\x12\x37\x23\xd4\xb6\x6d\x0e\x42\x64\x22\xb0\xce\x21\x7a\xc7\xcc\x95\x7a\x2b\x42\xb6\x48\xce\xa2\x7c\xbd\x78\x9a\x4a\xe2\x90\x6f\x9e\xc1\x05\x21\xb8\x45\x52\x05\x66\x3d\xdc\x1e\x0d\x8b\x3f\xe6\x66\xa1\xdb\x05\xc7\xc9\x68\xb1\xe5\x94\xe9\xc3\xac\x61\x27\x83\x61\xb7\x20\x41\xba\x17\x81\x11\x4b\x70\xc0\xfe\xa6\x41\x37\x47\x48\x71\x48\x59\xf3\x80\x94\xed\xda\xb0\xcf\xeb\xc7\x39\x92\xb1\x11\xb6\xbe\x84\xf4\xa4\x32\x5b\xfd\x89\xb5\x77\xdb\x4c\xc4\x36\xe6\x87\xbd\x23\x93\xeb\x16\x85\xf3\xbd\x23\x63\x26\x78\x3f\xc5\x98\xc1\x3d\xc5\x68\x95\x9b\xa9\xdb\x44\x86\x18\x9f\xe7\x3e\x80\x04\xef\xc4\x54\xa3\xbf\x13\x13\xd0\x2e\x78\xde\x72\xf2\x22\x56\xb7\xef\xb8\xe5\xd3\x3c\x88\xb1\xc5\xc8\x4c\x0b\x37\x78\xd4\x7d\x95\x51\xc5\x1c\xe8\x4a\xb9\x95\x68\x05\x3a\xde\xea\x73\xeb\xa0\x4c\xf4\x99\xe5\xa7\x97\xe7\xaf\x0f\x7d\x65\x80\x74\x93\x36\x3a\x33\xc8\xae\x3b\xfa\xd0\x83\x74\xab\xd2\x7d\x46\x17\x17\x0f\x2d\x60\xf3\x8b\x93\xdd\xa0\x45\xe6\xda\x66\x68\xe9\x74\xfc\x94\x5a\xc0\x25\x56\x13\x36\x27\x19\x7f\x12\x6c\xe6\xbc\x7e\x0f\x0e\xb3\x72\xeb\x4f\x98\xde\x15\xda\x01\x1b\xc3\x1c\xb8\x40\x92\xa1\xfa\xc7\xd1\x08\x7d\x6c\xf7\x90\x58\x81\x15\x75\xfb\x6c\x70\xc7\x58\x48\xbc\x5c\xe5\x32\xff\x0e\xb7\xfd\xe7\x93\xf4\xaf\x0e\x92\xbe\xaf\x61\xab\xd9\x0f\x11\x92\x58\x22\xc7\x1a\xa0\x5f\xcb\xba\xfb\x8a\xba\xcc\xa5\x32\xaf\xbd\x2f\x69\x3b\xb6\x1a\xc3\xf8\x62\x30\x98\xd3\xba\x1a\x83\x1e\x96\xb8\x68\x3a\xfe\x13\x71\x58\xb9\xe3\xc9\x9b\x34\xb9\x7d\xfc\x4a\xd0\x99\xb3\xfb\x00\x7f\x45\x03\x10\x42\x8b\xd5\x09\xf9\x56\xb8\x25\x71\x8d\x30\x55\x3a\x6f\xaf\x15\x04\x05\xee\x95\x80\x0c\x58\x21\x94\x98\xcb\x5b\x9e\xd9\xb6\xad\xb4\x4b\x82\x6f\xf4\xb7\x1f\xff\x07\x00\x00\xff\xff\xf0\x56\x4e\x35\x0c\x45\x00\x00")

func objdefsOmaJsonBytes() ([]byte, error) {
	return bindataRead(
		_objdefsOmaJson,
		"objdefs/oma.json",
	)
}

func objdefsOmaJson() (*asset, error) {
	bytes, err := objdefsOmaJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "objdefs/oma.json", size: 17676, mode: os.FileMode(420), modTime: time.Unix(1456665022, 0)}
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
	"objdefs/ipso.json": objdefsIpsoJson,
	"objdefs/oma.json":  objdefsOmaJson,
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
	"objdefs": &bintree{nil, map[string]*bintree{
		"ipso.json": &bintree{objdefsIpsoJson, map[string]*bintree{}},
		"oma.json":  &bintree{objdefsOmaJson, map[string]*bintree{}},
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
