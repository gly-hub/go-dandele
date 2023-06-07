// Code generated for package asset by go-bindata DO NOT EDIT. (@generated)
// sources:
// internal/template/boot/boot.tmpl
// internal/template/cmd/apiserver.tmpl
// internal/template/cmd/cobra.tmpl
// internal/template/cmd/rpcserver.tmpl
// internal/template/config/config.tmpl
// internal/template/global/global.tmpl
// internal/template/internal/route.tmpl
// internal/template/internal/rpcapi.tmpl
// internal/template/main.tmpl
package asset

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _internalTemplateBootBootTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x52\x00\xad\xff\x70\x61\x63\x6b\x61\x67\x65\x20\x62\x6f\x6f\x74\x0a\x0a\x66\x75\x6e\x63\x20\x49\x6e\x69\x74\x28\x29\x20\x7b\x0a\x09\x2f\x2f\x20\xe5\xb0\x86\xe9\x9c\x80\xe8\xa6\x81\xe5\x88\x9d\xe5\xa7\x8b\xe5\x8c\x96\xe7\x9a\x84\xe6\x96\xb9\xe6\xb3\x95\xe5\x9c\xa8\xe8\xaf\xa5\xe5\xa4\x84\xe6\xb3\xa8\xe5\x86\x8c\x20\x54\x4f\x44\x4f\x0a\x7d\x0a\x01\x00\x00\xff\xff\xa8\x42\xb5\x5d\x52\x00\x00\x00")

func internalTemplateBootBootTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateBootBootTmpl,
		"internal/template/boot/boot.tmpl",
	)
}

func internalTemplateBootBootTmpl() (*asset, error) {
	bytes, err := internalTemplateBootBootTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/boot/boot.tmpl", size: 82, mode: os.FileMode(420), modTime: time.Unix(1686111346, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateCmdApiserverTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x56\x5d\x6f\xdc\x44\x14\x7d\xf6\xfc\x8a\x8b\xa5\x54\x76\xb5\x6b\x0b\xf1\xb6\x22\x0f\x28\xe4\x4b\x42\xe9\x2a\xdb\xf2\xd2\x22\x34\xb1\xc7\xde\x51\xec\x19\x33\xbe\xde\x6c\x14\xad\x14\x5a\x68\xa1\xa4\x02\x41\xda\x87\x50\x24\x2a\xf1\x11\xf1\xd0\x14\x15\xa9\x2a\xa2\xf4\xc7\x50\x27\xbb\xff\x02\x8d\xc7\xbb\xd9\xa6\x4d\x49\xf0\x83\x3d\xf6\x9c\x3d\x73\xef\xb9\xe7\xce\x6c\x46\x83\x75\x1a\x33\xa0\x19\x27\x84\xa7\x99\x54\x08\x0e\xb1\xec\x28\x45\x9b\x58\x4a\x16\xc8\x45\x0c\x76\xcc\xb1\x5b\xac\x79\x81\x4c\xfd\x38\xd9\x6c\x76\x8b\x35\x3f\xa2\x39\x76\x11\xb3\x66\x0d\xb2\x89\x65\x6f\x6d\x81\xd7\x36\x8c\x2b\x34\x65\x30\x18\xf8\x5c\x20\x53\x82\x26\xbe\x86\x31\x0d\x7a\x0d\x57\x2c\x9b\x21\x15\x21\x4b\xb8\x14\x3e\xcd\xb2\x84\x07\x14\xb9\x14\x67\x81\x07\x52\x44\x3c\x3e\x0b\x32\x91\x71\xcc\xd4\x29\x48\x94\x32\x59\x93\x7d\x9f\x67\xff\x01\xc8\x51\x71\x11\xf7\x4f\xa0\xf2\x2c\x7a\xfb\x1d\x3f\x90\x6b\x8a\xea\x19\x2e\x7d\xae\x65\x49\xf4\x8b\xcc\xcd\xdd\xcf\x79\x2c\x68\x62\x13\x97\x90\x1e\x55\x5a\x67\x26\x7a\x50\x5d\x86\x95\x58\x1d\xa4\x0a\xe7\xd2\x10\x66\xe1\x42\xc5\xe6\xcd\xc9\x34\xa5\x22\xdc\x22\x96\x75\x25\x67\x2d\x98\x5c\x76\xce\x54\x8f\x29\xbb\x41\x2c\xab\xd3\x95\x0a\x27\x73\x76\xc5\x02\xef\xb5\x97\x61\x0a\x33\xdf\xa7\x69\x96\x8c\x19\xaa\x5a\x75\xaa\xd9\xba\x54\x35\x16\x9a\x0c\x12\x19\xd0\xc4\x10\xf3\x84\x89\x80\x5d\xc9\x69\xcc\x5a\x80\xaa\x60\xfa\x6b\x5b\xb1\xd5\x42\xb4\x20\x2a\x44\xe0\x04\x69\x08\x17\x5f\x8a\xb5\x01\x54\xc5\x39\x5c\xfd\xc8\x64\xe5\x82\x0e\xde\xca\x19\x16\x99\xe3\x12\xcb\x1a\x68\x8e\xd5\x42\xcc\x9f\x9d\x81\x29\x25\x95\xe1\x51\x0c\x0b\x25\x40\x15\x62\x42\x36\xd0\x92\x6a\x2a\xe0\x82\xa3\x53\x2d\x38\x56\xd2\x6b\x33\x95\xf3\x1c\x99\xc0\x85\x84\xc6\xb9\xe3\x7a\x9d\x8a\xf4\x43\xaa\xda\xce\x05\x26\x7a\x0d\xb0\x99\xe8\xd9\xfa\xa1\x6f\x75\xf2\x60\xcf\x8b\x9e\xed\x92\x41\xcd\x5c\x87\xaf\xa9\x7d\x1f\x46\x9f\xdf\x39\x7a\xf6\xb0\xfc\xe2\x87\xf2\xd7\xaf\xca\x9d\x7b\xc4\x32\x36\xf4\x96\x05\xc7\xb9\x6a\xe8\x30\xd1\x73\x2b\x6c\xf9\xe7\xee\xd1\xee\xfe\x14\x76\xca\xe1\xd5\x0f\x1c\x83\x1b\x3e\x39\x38\xda\xfd\x7d\x0a\x57\xb5\x4c\x85\x58\xd5\xa3\x1a\x76\xf8\x78\xbf\xbc\xb9\x53\xfe\xf4\xc7\xe8\xc6\x7e\x20\x05\xb2\x3e\x8e\xbe\xfb\x7b\xf8\xe4\xe0\x65\xe2\x55\x16\xeb\xac\xd5\x12\xa3\x21\x53\x0b\x5a\xe7\xe3\xe1\x71\x5a\x53\xd3\x01\xf6\xe1\x62\xdd\xcd\xde\x9c\x61\x6e\x40\x48\x91\x42\x4a\xb3\xab\xa6\x14\x93\x8a\xbc\xf2\xa9\x56\x66\x78\xeb\xb7\xf2\xe1\xde\x8b\xa7\x5f\x9a\x08\x4d\x68\xff\x6c\x5f\x1f\x1e\xfc\x7c\x78\xef\xe9\xe1\xe3\xbb\xc3\x1b\xcf\xca\x47\x37\x47\xf7\xb7\x87\xbf\x7c\x7a\xb4\xf7\x59\xf9\xf5\xf5\xc3\xbb\x8f\x46\xdb\x7b\xc3\xe7\xb7\x54\x16\x0c\x9f\x7f\x3f\x7c\xb0\xf3\xe2\xaf\x1f\x47\xdb\xdf\xc2\xe5\x4b\xef\x5f\x22\x64\x5c\x71\x1d\xc9\x24\xee\xaa\xfc\xc7\xb6\xd0\x32\x7f\x73\x50\xde\xde\xd7\x5b\xd2\xe1\xfd\x3b\xe5\xed\x07\xc4\x8a\xa5\xf1\x97\x31\xe0\xb4\x38\x4b\x88\x99\x31\xbf\xb6\x43\x3d\x20\xd6\x40\xdf\x2a\x4d\x05\x36\xe0\x63\x68\xcd\x82\x69\x63\x6f\x95\xd1\x70\x81\x27\xcc\xb1\x3d\x3f\x47\x8a\x3c\x98\xec\x6c\xcd\x98\x22\xdb\xa0\x9b\x1e\xf6\xd1\x76\x89\x15\xa5\xe8\xb5\x15\x17\x98\x08\xc7\x6c\x39\xde\xa2\x62\x4c\x38\x46\x26\xa7\xe6\x77\x5d\x97\xe8\x46\x3c\x15\x6e\x9b\xb8\x74\xa6\x40\xb1\x65\x9f\xc4\x47\x8e\xdd\x04\xf8\x40\xbb\x55\xf7\xb4\x4e\xbc\xe5\xfb\x95\x7b\xbb\x32\xc7\xd6\x4c\xe8\xc3\x35\x75\x4d\xd8\x0d\x38\x35\xf5\xb6\x54\xe8\xbc\x9e\x78\x85\xe1\x86\x54\xeb\xad\x31\xf1\x4c\x3e\xcd\xc8\x33\x6f\x91\x61\xb5\xf8\x92\xcc\xd1\x71\xcf\xbb\x48\xa2\xbb\x57\x7f\xe1\x11\xd4\xbd\xb3\xc8\x70\x5e\xf4\x1c\x17\xde\x9a\x05\x3b\x53\x32\x2c\x82\xea\x10\x80\x2d\x32\xde\xd9\xde\x20\xd6\x06\xd5\xaf\x27\xd5\x3a\xaf\x62\x34\xe3\x7e\x6e\xa8\x7c\x2e\x42\xd6\xf7\xba\x98\x26\xe7\xd2\xf1\xcc\x5a\xbe\x79\xad\xff\xab\xf0\x80\x54\x8f\x4f\x0a\x8e\xda\xbf\x29\x5d\x67\x4e\xd0\xa5\x02\x64\xee\x75\xaa\x03\xc8\xe0\xcc\x61\xe4\xad\x48\xe4\xd1\xa6\xa3\xe1\x0d\x0d\x59\xd6\xae\x56\x45\x86\x06\xf5\x6e\x53\xcf\xbc\xe2\x8f\x99\x1c\x3a\xdd\x02\x43\xb9\x21\xa0\xb6\xa9\xe7\x79\xe3\xd0\xeb\x23\x52\xc7\x3f\x57\x28\xc5\x04\x5e\xe6\x29\xeb\xa0\x1a\xc7\x58\x97\x6e\x59\x44\x72\x62\x73\xd6\xe7\xd5\xdf\x08\x83\xa8\x5b\x5e\xf0\x84\x0c\xc8\xbf\x01\x00\x00\xff\xff\x7f\x67\x30\x29\xa5\x08\x00\x00")

func internalTemplateCmdApiserverTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateCmdApiserverTmpl,
		"internal/template/cmd/apiserver.tmpl",
	)
}

func internalTemplateCmdApiserverTmpl() (*asset, error) {
	bytes, err := internalTemplateCmdApiserverTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/cmd/apiserver.tmpl", size: 2213, mode: os.FileMode(420), modTime: time.Unix(1686123884, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateCmdCobraTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x91\x41\x8f\xd3\x30\x10\x85\xcf\x9e\x5f\x31\xe4\x80\x12\xd4\x26\xaa\xb8\x05\x7a\x58\xad\x7a\x43\xab\x6a\xab\x3d\x21\x0e\xae\x3d\x75\x2d\x62\x3b\x8c\xed\x65\x21\xca\x7f\x47\x4e\x4b\x05\x2b\x4e\x5c\xe7\x7b\xf3\xde\xd3\xcc\x28\xd5\x57\x69\x08\x95\xd3\x00\xd6\x8d\x81\x13\xd6\x20\x2a\x62\x0e\x1c\x2b\x10\xd5\x34\x61\xbb\xbf\xa8\x1e\xa4\x23\x9c\xe7\x4e\x39\xdd\xc9\xd1\x16\x6a\x6c\x3a\xe7\x63\xab\x82\xeb\xcc\xf0\x63\x7d\xce\xc7\xce\x84\xb5\x96\x5e\xd3\x60\x83\xef\x86\x60\x0c\xf1\x2b\x65\x1c\x4f\x9b\xf7\x9d\x0a\x47\x96\x85\x84\x58\x41\x03\xf0\x2c\x19\x39\x84\x74\xef\x34\x6e\xf1\xed\x82\xdb\xfb\xe0\x9c\xf4\x7a\x02\xf1\x14\xa9\xc7\x7f\xb4\xa9\x56\x20\x0e\xe7\xc0\xe9\x4a\x0f\xc4\xcf\xc4\x7f\x42\x3b\x90\x57\xf4\x14\xa5\xa1\x3e\x71\xa6\x15\x88\x4f\xc1\x9b\x1e\x2b\x99\xd3\x39\xb0\xfd\x49\x45\x77\xc7\x26\xf6\x78\xca\x5e\xd5\xca\x69\x7c\xf7\x57\x81\x15\x4a\x36\x11\x3f\x7f\x89\x89\xad\x37\x0d\x2e\x07\xc2\x09\x84\xb0\x27\x1c\xc8\xd7\x85\x37\xf8\x11\x37\xcb\x50\x30\xa5\xcc\xfe\x22\x8b\xed\x03\x7d\xaf\x2f\xa7\x68\x1f\x49\xd7\x15\xd3\xb7\x6c\x99\x22\xca\x84\x03\xc9\x98\x30\x78\x2a\x11\x55\xd3\x80\x10\x33\xdc\x0c\xbc\x1d\x40\xcc\x2b\x10\x7b\xe2\x68\x63\x22\x9f\xf6\x21\xa6\xc7\xec\x77\xff\x53\xf6\x95\xeb\x0c\x50\x3c\xd0\x7a\x9b\xea\x66\x02\x71\x7d\x40\x7b\xa7\xf5\xd5\xac\x96\xa3\x6d\x0f\x49\x72\x99\x37\xb7\x8d\xdd\x0b\xa9\x9c\x68\x59\xb2\xa7\x92\x80\xfd\xf6\xf7\xff\xda\x1b\xfd\xb0\x90\x37\xdb\x92\x58\xf2\x43\x6c\x77\x2f\x36\xd5\xeb\x4d\x03\x62\x86\x19\x7e\x05\x00\x00\xff\xff\xda\xd0\x84\x07\x81\x02\x00\x00")

func internalTemplateCmdCobraTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateCmdCobraTmpl,
		"internal/template/cmd/cobra.tmpl",
	)
}

func internalTemplateCmdCobraTmpl() (*asset, error) {
	bytes, err := internalTemplateCmdCobraTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/cmd/cobra.tmpl", size: 641, mode: os.FileMode(420), modTime: time.Unix(1686118345, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateCmdRpcserverTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x94\xbf\x8f\x1b\x45\x14\xc7\xeb\x99\xbf\xe2\x69\xa4\x44\xb3\xc8\x37\x2b\x44\x77\x22\x05\xb2\x2e\xd1\x35\x91\x65\x13\x1a\x82\xd0\x78\xfd\x76\x3d\xca\xec\x9b\x65\xf6\xad\xe3\xe8\xe4\x12\x09\x01\xe9\x42\x2a\x28\xa8\xa0\xa2\xa1\x02\x89\x3f\xc7\x17\xf1\x5f\xa0\xd9\x59\x1f\x06\x29\xd2\x6d\xb1\xb6\xe7\x7d\xe7\xf3\x7e\xbb\xb3\xd5\x0b\xdb\x20\xd8\xce\x49\xe9\xda\x2e\x44\x06\x2d\x85\xaa\x5b\x56\x52\xa8\x9b\x1b\x30\x8b\x2c\x79\x6a\x5b\x84\xc3\xa1\x5c\x87\xf0\x3e\x93\x23\xc6\x48\xd6\x97\x3d\xc6\x9d\xab\x30\xc9\x1a\xc7\xdb\x61\x6d\xaa\xd0\x96\x8d\x7f\x75\xb1\x1d\xd6\x65\x13\x2e\x36\x96\x36\xe8\x5d\xa0\xd2\x76\x9d\x77\x95\x65\x17\xe8\x3e\xf2\x2a\x50\xed\x9a\xfb\x28\x7d\x68\x1a\x8c\xef\x51\x72\x08\x7e\x1d\xf6\x65\xcf\xd1\x51\xb3\xff\x9f\xaa\xef\xea\x0f\x3f\x2a\xab\xb0\x8e\x36\x59\x5c\x28\x5d\x18\xd8\xf9\xf4\x23\xf4\xf9\x5d\xf6\xae\x21\xeb\x95\x2c\xa4\xdc\xd9\x98\x8a\x86\xb4\x83\xf1\xc9\x54\x29\x56\x6c\x23\xcf\xdb\x0d\x3c\x82\x87\x23\xcd\xcc\x43\xdb\x5a\xda\xdc\x48\x21\x9e\xf5\x78\x09\x77\x8f\x4a\x25\xc3\xa8\x66\x52\x88\xd5\x36\x44\xbe\xb3\xa9\x91\x02\xcb\xc5\x1c\xce\x34\x57\x7b\xdb\x76\xfe\x44\x18\x9b\xb1\x1a\xad\x53\x2f\x26\x2d\x5c\x20\xf8\x50\x59\x9f\xc1\xce\x23\x55\xf8\xac\xb7\x0d\x5e\x02\xc7\x01\xd3\xe9\x22\xe2\x72\xa0\x4b\xa8\x07\xaa\x74\xd5\x6e\xe0\x83\xff\xc4\x3a\x03\x1b\x9b\x1e\x3e\xff\x22\x67\x55\x40\x0a\x5e\xf4\xc8\x43\xa7\x0b\x29\xc4\x21\x31\x96\x03\x5d\xdd\x9f\x80\x31\x86\x98\x39\x11\x79\x88\x04\x71\xa0\x3b\xd8\x21\x95\x34\xa1\xc0\x91\x63\x3d\x3a\x3c\x55\xd2\x2c\x30\xf6\xae\x67\x24\x7e\xec\x6d\xd3\xeb\xc2\xac\x46\xe8\x67\x36\x2e\xf4\x43\xa4\xdd\x0c\x14\xd2\x4e\xa5\x8f\xf4\x9a\x92\x07\x75\x45\x3b\x55\xc8\xc3\x44\x9e\xc2\x4f\xe8\xb2\x84\xbf\xbf\x7e\xfd\xee\xaf\xdf\x8e\xdf\xfc\x74\xfc\xe5\xbb\xe3\xf7\x6f\xa5\xc8\x53\x66\xae\xc9\xf1\x7c\xfc\xaa\x91\x76\xc5\xa8\x3d\xfe\xf9\xe6\xdd\x9b\x5f\xcf\xb4\x67\x03\x3c\x5e\xd0\x93\xee\xa4\xb8\xfd\xf1\xf5\xf1\xdb\x9f\x6f\xdf\xfe\x71\xfb\xfb\x0f\x52\xa4\xed\x39\xc9\x4e\xc1\x8c\xc9\xff\x5b\x94\xf3\xcb\xb1\xab\xa0\x0d\x1b\xf4\x52\x34\x21\xd7\x37\x37\xe0\xdc\xeb\xb2\xab\x72\xef\x35\xe1\x4b\x3d\xed\x5e\x3a\xfd\xa4\x73\x45\x21\xc5\x21\x85\x54\x05\x4a\x55\x9b\xc1\x97\x70\xf9\x08\xf2\x3c\x9b\x25\xda\xcd\x63\xe7\x51\x2b\x53\xf6\x6c\xd9\x55\xe5\x83\xde\xf0\x9e\x55\x21\x45\xdd\xb2\x59\x44\x47\xec\x49\xe7\x65\x32\x4f\x22\x22\xe9\xdc\x46\x3d\x11\x8b\xe4\xe2\xab\xc1\x71\xe2\xb6\xf6\x05\xea\x6a\x6b\x09\x42\x6f\x56\xe3\x86\x14\x52\xe4\x55\x31\x4f\x03\xbb\xfa\x95\x4e\xda\x59\xb2\x5f\xa7\xbf\x8b\x38\x74\x5c\x48\xf1\xf1\x45\x3a\x3e\x73\x5a\x6b\xf5\xa0\x87\xd5\x76\xe0\x4d\x78\x49\x90\x13\x04\x63\x0c\x3c\x8f\xcf\x49\xcd\xa6\x25\xdb\x9b\x27\xc8\xf3\x21\x46\x24\xfe\xd4\xb5\xb8\xe2\xa8\x8b\x42\xa6\xa5\x98\x62\xbe\xa6\x3a\x68\x35\xdd\xc7\xbd\x63\x47\x8d\xca\x8a\x69\xf8\xc8\x79\x79\x90\xff\x04\x00\x00\xff\xff\x00\x02\x3e\xfa\x0a\x05\x00\x00")

func internalTemplateCmdRpcserverTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateCmdRpcserverTmpl,
		"internal/template/cmd/rpcserver.tmpl",
	)
}

func internalTemplateCmdRpcserverTmpl() (*asset, error) {
	bytes, err := internalTemplateCmdRpcserverTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/cmd/rpcserver.tmpl", size: 1290, mode: os.FileMode(420), modTime: time.Unix(1686123896, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateConfigConfigTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x94\xcf\x6e\x13\x3f\x10\xc7\xef\xfb\x14\x23\xff\xce\xbf\x68\x37\x29\x49\xea\x5b\x01\x21\x90\x2a\x40\x6d\x24\x0e\x15\x07\x37\x9e\x34\x2b\x9c\xb5\xb1\x9d\x84\x6a\x65\x89\x0a\x24\xa4\x4a\x08\x90\xe0\x00\x37\x24\x5a\x4e\x55\x38\x70\xe2\xc2\xcb\x40\xba\x7d\x0b\x64\xef\x26\x9b\x96\x94\xd3\xda\x1f\xcf\xf8\x3b\xff\xbc\x79\x0e\xe9\x00\x1a\x3d\x29\x85\x69\x6c\xcb\x83\x03\xd4\xe0\x5c\x24\xc2\x8a\x46\x00\xff\x6d\xe3\x04\x05\xc4\x70\xfe\xfd\xeb\xfc\xf9\xc9\xf9\xa7\x97\x90\x14\x67\xa7\xf3\xe3\x13\x68\x5e\xbc\x7a\x5d\x9c\x1e\x79\xd4\xba\x78\xff\xb1\x98\xcd\x60\xa3\x38\x3b\xfd\xfd\xee\x18\x6e\xcc\xdf\xbc\x3d\xff\xf2\x03\xda\xbf\x7e\x7e\x9e\x1f\xcd\xa0\x53\x7c\x7b\x51\xcc\x3e\x44\x00\x7d\x99\x19\x29\x70\x77\x28\xa7\x14\xac\x1e\x63\xcd\x82\x12\x05\xe8\x44\x00\x83\x54\xe0\x23\x9d\x5a\xa4\x00\x03\x26\x0c\x56\x6c\xd5\x66\x34\x16\x36\xbd\x53\x1b\x2e\xec\x96\xbc\x32\xee\x44\x79\x0e\x98\x71\x9f\x58\x74\x29\xe1\xbb\xd6\x2a\x70\x6e\x68\xad\xda\x45\x3d\x29\x33\x56\x52\x5b\x0a\xdd\xb8\x1b\xfb\x8d\xd2\x72\x10\x76\xdd\xd5\x5b\xb4\xea\xd7\x0e\x23\xc9\xbd\x4c\x12\x01\x98\x00\xef\xb3\x11\x52\x20\x79\x0e\x8d\xdd\x25\x00\xe7\x48\x04\x80\xb6\xcf\x29\xec\x91\xa4\xd9\x69\xc4\x8d\xb8\x91\xd0\x66\xab\xb3\x49\x1e\x47\x00\xfb\xcc\xe0\x43\x66\x87\x95\xeb\x96\x52\xc1\x27\xcf\xeb\x80\x77\x54\xdf\x39\x00\xc6\xb9\xa6\x40\x48\x1d\x6d\x77\x73\xb3\x8e\x36\x09\xdb\x3c\xc7\x8c\x97\x29\xd7\x17\xdc\xbe\xe9\x1c\xdf\xf7\x51\xf3\xfd\xde\xa1\xf2\x61\x8e\x0e\xcd\x53\xe1\xaf\x1a\xb1\x67\x0f\x14\x66\xb7\x64\x96\x51\x68\xc6\x25\xb9\xc7\x05\x96\x64\xa3\x06\xbd\xd4\x27\x98\xc4\x95\xcd\x76\x3a\xa8\x50\xab\x1d\x98\x28\x0b\xef\x3d\x8c\x90\xd3\xde\x50\xa3\x19\x4a\xc1\x29\x90\x24\x8e\x47\xa6\x54\x33\xb6\xac\x1f\xc0\xd8\xa0\xcf\x47\x4b\x69\x49\x00\x8a\x19\x33\x95\xda\x3b\x2c\x96\xe5\xc1\x50\x1a\xeb\x6f\x59\x54\xaf\x32\x0f\x45\x20\xad\x56\xdc\x2e\x01\x67\x96\xf9\x72\x52\x20\x8b\x25\x09\xc1\xb0\x09\x96\x92\xff\xff\x2d\xfa\x0f\xd9\xeb\x84\xd7\x48\x5f\x23\xbe\xb6\x1b\x3b\xc8\x53\xe3\x5c\xa4\xfd\xd7\xc7\x15\x16\x55\x5f\x98\x90\x59\x88\x3a\x43\x3b\x95\xfa\xc9\xaa\x3a\x6d\xfb\xa1\xf1\x19\x59\xa6\xed\x56\x18\x87\xbd\xab\xc7\x7e\xa6\x58\xdf\xa6\x93\x65\xaf\x52\x2e\x96\x6b\x36\x0e\x93\x46\xca\xf7\x97\xf9\xfe\xc9\xb1\x5d\xed\x90\x46\xc6\xd7\xe0\xa9\x7f\x6d\x6b\x78\x5a\x0d\xc6\x25\xbc\x36\xed\x9e\x66\x7d\x74\x2e\xb2\xfe\x1b\x46\x40\x2a\xcc\x02\x5d\xfe\x13\xc2\x59\xf5\x90\xea\x57\xe5\x75\xae\x76\x82\xb6\xbb\xad\xa4\x56\xfa\x13\x00\x00\xff\xff\xdd\x91\xa0\xf9\xd2\x04\x00\x00")

func internalTemplateConfigConfigTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateConfigConfigTmpl,
		"internal/template/config/config.tmpl",
	)
}

func internalTemplateConfigConfigTmpl() (*asset, error) {
	bytes, err := internalTemplateConfigConfigTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/config/config.tmpl", size: 1234, mode: os.FileMode(420), modTime: time.Unix(1686119096, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateGlobalGlobalTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x48\xcf\xc9\x4f\x4a\xcc\xe1\x02\x04\x00\x00\xff\xff\x32\x37\xa5\xbc\x0f\x00\x00\x00")

func internalTemplateGlobalGlobalTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateGlobalGlobalTmpl,
		"internal/template/global/global.tmpl",
	)
}

func internalTemplateGlobalGlobalTmpl() (*asset, error) {
	bytes, err := internalTemplateGlobalGlobalTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/global/global.tmpl", size: 15, mode: os.FileMode(420), modTime: time.Unix(1686108961, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateInternalRouteTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\x31\x6b\x1b\x31\x14\x80\x67\xbd\x5f\xa1\x6a\xba\x33\xb5\xb4\x17\x8c\x97\x96\xda\x50\x30\xd8\x85\xce\xf2\x9d\x2c\x8b\x9e\x25\x21\xeb\x5c\x4a\x31\x74\x68\x3d\xb5\x24\x10\x67\x72\xc6\x18\xbc\x04\x13\x3c\x24\x84\x84\xfc\x1a\xc5\xce\xbf\x08\xe7\x73\x82\x33\xc5\xa3\xc4\xf7\xde\xfb\x3e\xcb\x93\xef\x5c\x0a\xec\x4c\xee\x05\x80\x1a\x58\xe3\x3c\x8e\x00\x11\xa9\x7c\x3f\xef\xd2\xc4\x0c\x98\xcc\x7e\x56\xfb\x79\x97\x49\x53\x4d\xb9\x4e\x45\xa6\x8c\x66\xdc\xda\x4c\x25\xdc\x2b\xa3\xc9\x01\x78\x62\x74\x4f\xc9\x43\xc8\xa1\x70\x23\xe1\x58\xdf\x7b\x4b\x00\x15\x62\x4a\xcb\xce\x0f\x2e\xa5\x70\xf8\xed\xe9\x12\x24\x10\x03\xf4\x72\x9d\xe0\xa6\x56\xbe\x5d\xd4\x45\x31\xfe\x05\xa8\xcb\x87\x62\xfb\x74\xf8\x43\x0d\xef\x45\xd0\x86\xf7\xb6\xb3\xbd\x1d\xc5\xb4\x44\xa2\x18\x90\xea\xe1\xd2\x9d\x7e\x16\xfe\x93\x1e\x45\x31\x7e\x57\xc3\xc4\x3a\x93\xe6\xc9\xb6\xbe\x58\x8b\x18\xc3\x0f\xab\x45\x98\xfc\xdb\x09\x00\xda\x3b\x55\x8c\x46\xe4\xd9\x8d\x55\xc8\x7b\xfc\xba\x8b\x7e\x73\xdc\x36\xb8\x4e\x33\xe1\x62\x40\xa8\x88\xa7\x5f\x8c\x6c\x4a\x6d\x9c\x68\x8b\x61\x9e\xf9\x88\xd0\x4a\xfd\x65\x07\xad\xd4\x49\x8c\x19\xc3\xe1\xfe\x6e\x7d\x3a\xdf\x7d\x87\x93\xff\xe1\x66\x1a\x7e\xdf\x02\x1a\x03\x14\x52\xe1\x68\x19\xce\x16\x9b\xe5\x3c\x9c\xff\x29\x05\xd7\xb3\xeb\xf0\x77\x15\x2e\x8e\x1f\x67\x93\xcd\xd5\x72\x3d\xbd\xc4\x5f\x5b\x1f\x5b\x00\xc8\x09\x9f\x3b\x0d\x63\x78\x0a\x00\x00\xff\xff\xb1\xcf\x6a\xbd\x16\x02\x00\x00")

func internalTemplateInternalRouteTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateInternalRouteTmpl,
		"internal/template/internal/route.tmpl",
	)
}

func internalTemplateInternalRouteTmpl() (*asset, error) {
	bytes, err := internalTemplateInternalRouteTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/internal/route.tmpl", size: 534, mode: os.FileMode(420), modTime: time.Unix(1686108871, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateInternalRpcapiTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\x28\x4e\x2d\x2a\xcb\x4c\x4e\xe5\xe2\x2a\xa9\x2c\x48\x55\x08\x2a\x48\x76\x2c\xc8\x54\x28\x2e\x29\x2a\x4d\x2e\x51\xa8\xe6\xe2\xaa\xe5\x02\x04\x00\x00\xff\xff\xbd\xf9\x40\x98\x29\x00\x00\x00")

func internalTemplateInternalRpcapiTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateInternalRpcapiTmpl,
		"internal/template/internal/rpcapi.tmpl",
	)
}

func internalTemplateInternalRpcapiTmpl() (*asset, error) {
	bytes, err := internalTemplateInternalRpcapiTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/internal/rpcapi.tmpl", size: 41, mode: os.FileMode(420), modTime: time.Unix(1686108931, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _internalTemplateMainTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x48\x4c\xce\x4e\x4c\x4f\x55\xc8\x4d\xcc\xcc\xe3\xe2\xca\xcc\x2d\xc8\x2f\x2a\x51\x50\xaa\xae\x56\xd0\x0b\x80\xc8\xf8\x25\xe6\xa6\x2a\xd4\xd6\xea\x27\xe7\xa6\x28\x71\x71\xa5\x95\xe6\x25\x83\xd5\x6a\x68\x2a\x54\x73\x71\x26\xe7\xa6\xe8\xb9\x56\xa4\x26\x97\x96\xa4\x6a\x68\x72\xd5\x72\x01\x02\x00\x00\xff\xff\x12\xa7\xcc\x37\x4e\x00\x00\x00")

func internalTemplateMainTmplBytes() ([]byte, error) {
	return bindataRead(
		_internalTemplateMainTmpl,
		"internal/template/main.tmpl",
	)
}

func internalTemplateMainTmpl() (*asset, error) {
	bytes, err := internalTemplateMainTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "internal/template/main.tmpl", size: 78, mode: os.FileMode(420), modTime: time.Unix(1686118271, 0)}
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
	"internal/template/boot/boot.tmpl":       internalTemplateBootBootTmpl,
	"internal/template/cmd/apiserver.tmpl":   internalTemplateCmdApiserverTmpl,
	"internal/template/cmd/cobra.tmpl":       internalTemplateCmdCobraTmpl,
	"internal/template/cmd/rpcserver.tmpl":   internalTemplateCmdRpcserverTmpl,
	"internal/template/config/config.tmpl":   internalTemplateConfigConfigTmpl,
	"internal/template/global/global.tmpl":   internalTemplateGlobalGlobalTmpl,
	"internal/template/internal/route.tmpl":  internalTemplateInternalRouteTmpl,
	"internal/template/internal/rpcapi.tmpl": internalTemplateInternalRpcapiTmpl,
	"internal/template/main.tmpl":            internalTemplateMainTmpl,
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
	"internal": &bintree{nil, map[string]*bintree{
		"template": &bintree{nil, map[string]*bintree{
			"boot": &bintree{nil, map[string]*bintree{
				"boot.tmpl": &bintree{internalTemplateBootBootTmpl, map[string]*bintree{}},
			}},
			"cmd": &bintree{nil, map[string]*bintree{
				"apiserver.tmpl": &bintree{internalTemplateCmdApiserverTmpl, map[string]*bintree{}},
				"cobra.tmpl":     &bintree{internalTemplateCmdCobraTmpl, map[string]*bintree{}},
				"rpcserver.tmpl": &bintree{internalTemplateCmdRpcserverTmpl, map[string]*bintree{}},
			}},
			"config": &bintree{nil, map[string]*bintree{
				"config.tmpl": &bintree{internalTemplateConfigConfigTmpl, map[string]*bintree{}},
			}},
			"global": &bintree{nil, map[string]*bintree{
				"global.tmpl": &bintree{internalTemplateGlobalGlobalTmpl, map[string]*bintree{}},
			}},
			"internal": &bintree{nil, map[string]*bintree{
				"route.tmpl":  &bintree{internalTemplateInternalRouteTmpl, map[string]*bintree{}},
				"rpcapi.tmpl": &bintree{internalTemplateInternalRpcapiTmpl, map[string]*bintree{}},
			}},
			"main.tmpl": &bintree{internalTemplateMainTmpl, map[string]*bintree{}},
		}},
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
