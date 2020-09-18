// Code generated by go-bindata. (@generated) DO NOT EDIT.

// Package graphql generated by go-bindata.// sources:
// schema.graphql
package graphql

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x57\xcd\x72\xdb\x36\x10\x3e\x57\x4f\x01\xa9\x17\x67\xc6\x93\x07\xd0\x4d\x56\x92\xd6\x93\xa8\x71\x23\x35\x97\x4c\x46\x03\x83\x2b\x0a\x63\x10\x60\x81\x85\x6d\x4e\xa7\xef\xde\x01\x08\x4a\xfc\x01\x48\x79\x7a\xf1\xc1\x3c\x91\xbb\xdf\x2e\xf6\x07\xfb\x43\x2e\x11\xf4\x81\x32\x20\x6b\x55\x14\xa0\x19\xec\xef\xb4\xca\x2c\x43\xf2\xcf\x8c\x10\x42\xb0\x2a\x61\x49\xb6\xa8\xb9\xcc\xe7\x9e\x52\x50\xfd\x00\x78\x27\x28\x83\xb5\xca\xfa\x4c\xc8\x38\x5d\x0e\x94\xed\x37\x8e\x3e\xbf\xf6\x98\x52\x73\x06\x11\xcc\x9d\xa3\xdf\xca\x83\x0a\x38\xe4\x28\x7a\xea\x19\x45\xc8\x95\xe6\x60\x22\xf2\xeb\x13\xb3\x06\x67\x60\x98\xe6\x25\x72\x25\xbb\x5a\xcc\x51\x69\xfc\x90\xe2\x16\x80\x71\x0f\x90\xd6\x00\xa1\x2a\x2a\xb0\x8a\x60\xbe\xd4\x9c\x1a\x46\x11\x35\xbf\xb7\x18\xb5\x75\x75\x62\xce\x67\xff\xce\x66\x33\x17\xe6\x21\x6a\xcb\x8b\x52\x40\x93\x10\xff\x51\x80\x44\xf3\x96\xac\x57\x98\xac\xb5\x92\x07\x9e\x5b\x4d\xef\xdf\x52\xf6\xaa\x52\xe6\xc0\x8f\x54\x73\xea\x6c\xd9\x82\x00\xe6\x5e\xcc\x92\xfc\x18\x88\x7d\x1f\xc0\xe6\x3f\x47\x52\xbe\x62\xc8\x1f\xc1\x0b\x49\x7c\xcb\xf9\x6b\xcc\xb9\xc4\xcd\x58\x6c\xff\xcf\xbd\xf0\x06\x9d\x6f\xc0\xb4\x96\x55\x02\x5c\xdf\xb1\xf8\x15\x1b\xa2\xc3\x25\x62\x03\x5f\x04\xbd\x07\xd1\x25\xa9\xf2\x05\x2e\xed\xbf\x7a\xf4\x88\x31\x29\x07\x82\x49\xa7\xec\x7c\x49\x59\xd2\x63\xbc\xc0\xeb\x60\x5c\x38\x29\xe2\xaa\x41\x8a\xb1\x0a\x48\xaa\xda\x3a\x81\x8b\x6e\xca\xb5\x33\x14\xa4\x2d\x5e\xaa\x3d\x58\xbb\x5a\xef\x6e\xbf\x7f\xf4\xaf\x9b\xd5\x6e\xfd\xbb\x7f\xfb\xe3\xeb\xbe\xfe\x48\x06\xe1\x5c\x7b\x41\x4f\x41\xb9\x4c\xd7\x68\xb5\x03\x6a\x40\x87\x42\x11\x62\x49\x48\x24\xed\x3d\xf0\x48\xb2\x5d\x71\x86\x83\x1f\xa0\x7a\x52\x3a\x33\xcb\x5f\xdc\xd7\x8f\x10\x96\x11\xd9\x50\xb4\x41\x3c\x34\x27\xff\xa4\xa0\xe7\x4e\xe5\x45\x80\x6a\xc9\x65\xbe\x1c\x15\xf9\x58\x83\xbc\xd0\x94\x29\x67\xfd\xc9\x26\x9c\xc1\x81\x5a\x81\x9d\x08\x73\x16\xae\x08\x37\x1f\xb8\x61\xca\x4a\x84\x6c\x49\x6e\x94\x12\x40\x65\x90\x6b\x31\x62\xa2\x0d\x7f\x07\xcf\xd8\x6b\x8b\x5c\xde\x29\x2e\xd1\xec\xd4\xb6\x04\x89\x4b\xf2\x49\x28\x8a\xcd\x4c\x78\x4e\x33\x99\x92\xe8\xd5\xc5\x3b\xfe\xba\x66\xfb\x0a\x5b\x84\x08\xb4\xa2\xe5\xba\xf5\x93\x21\x78\x04\x1f\x06\x42\x65\xe6\x3f\x4a\x7f\x9e\x0f\x3e\x64\x8b\x89\x80\xb6\xf5\xd5\x21\x5d\xec\x1a\x85\xea\xe0\xf5\x0d\x4f\xbe\x26\xf0\x3e\x7f\x4f\x36\x5c\x80\x59\xc9\x6c\xa3\x34\x2c\x12\xe9\xf0\xda\x1e\xa9\xb0\xa3\xea\x98\xd5\x1a\x24\xab\x08\xa3\x92\xdc\x43\xad\x3e\xf8\xa1\x34\x29\xdc\x41\x8b\xf1\xf4\x26\xaf\x4e\x3b\x94\x4d\xd7\xb5\x06\x55\x01\xfa\x37\xad\x6c\xd9\x9b\xa3\x47\x2a\x25\x88\xe1\x8c\x11\x8a\xd1\xf6\xd0\x1d\xa9\xb8\x8c\x37\x25\xe7\x0b\x78\x58\xbe\x1e\x72\x8b\x50\x84\xf9\x93\x03\x7a\xd2\x95\x35\x34\x3f\x9f\xf1\x2e\xb5\x2e\x78\xd1\x09\x0b\x1c\xa6\x53\x24\xfe\xe9\xde\xdb\x02\x76\x35\xab\x4d\x0e\x36\xf4\xd1\x61\xe7\xe8\x93\x35\x1c\xc0\xa5\xee\x92\xc8\x9c\x47\x7b\x7f\xd4\x7c\x86\xca\x8d\xb7\x53\x4b\xea\x6f\x09\x91\x29\xdc\x70\x03\xfa\x48\xcd\x89\x74\xf5\x00\x55\x3b\x88\xa1\xce\x9b\x50\x27\x71\xe9\x43\x06\xa2\xe6\xa6\xfa\x0c\x95\x53\xd0\xb6\xfb\xdd\x94\xa5\xd3\xa1\x49\xee\x05\x8e\x12\x99\xc8\x91\x19\x6a\x25\xc7\xd8\x8e\x24\x2c\x74\x82\x3c\x35\xb2\xc2\x74\x49\x5a\x54\x52\x3c\x76\x29\x92\x16\x03\x8c\xf6\x4d\x6f\xe2\x8c\x89\xea\x6d\x75\xa7\x44\xfd\x4f\xb4\xf0\x89\x0e\x5e\xef\x7f\x37\xd4\x40\xa7\x3d\x9f\xc9\xab\xc2\x09\x26\x98\x7f\x49\x7e\xd2\x97\x9e\x32\x61\x4b\x2f\x4a\xca\x73\xf9\xcd\x0a\x18\x5c\xf8\x0c\x64\xe5\x5a\x69\x23\x6c\xba\xb2\xbf\x5e\x3e\x31\xea\x82\xa5\xcf\x6b\x41\x8d\xb9\xa0\x30\xb7\x40\x35\x3b\x7e\x03\x63\x05\x9e\xa6\xbd\x67\xc5\x6a\x2f\x98\x7b\xa0\x0c\xba\xfc\x5a\xcd\xfe\x93\x63\xcc\x7f\x86\x85\xce\xe6\x39\x98\xc1\xfe\x1a\xa0\xdb\x13\x37\x28\x35\x9e\xbe\xe9\xfe\x68\x04\xf0\xf9\x3f\xe3\x48\x4d\xbd\xa9\x41\xe6\x0f\x6b\x8d\x72\xb7\xe3\x3d\x23\xb8\x51\xe8\x5c\xfd\xd3\x82\x6e\x16\x98\xbe\x1f\x57\xa9\xff\xb5\xeb\x89\x55\x32\xd2\x2a\xa2\x07\x04\xc3\xaf\x4c\x08\xef\xdf\x16\x0c\x0e\xfd\x0a\x8c\x58\x03\x6a\x27\xc6\xf9\xf6\x5f\x00\x00\x00\xff\xff\x86\x84\x8d\xaf\xe9\x13\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
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
	"schema.graphql": schemaGraphql,
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
// AssetDir("foo.txt") and AssetDir("nonexistent") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
