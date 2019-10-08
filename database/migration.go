// Package database Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// database/migrations/000001_init.down.sql
// database/migrations/000001_init.up.sql
package database

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

var __000001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func _000001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initDownSql,
		"000001_init.down.sql",
	)
}

func _000001_initDownSql() (*asset, error) {
	bytes, err := _000001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.down.sql", size: 0, mode: os.FileMode(420), modTime: time.Unix(1570529246, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __000001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x98\xcd\x8e\xdb\x36\x10\xc7\xef\x7a\x0a\xde\xf2\x81\x45\x03\x14\x6d\x2e\x39\x6d\x52\xb7\x30\x9a\x6c\x02\xc7\x05\x12\x14\x05\x41\x49\x63\x69\xb0\x14\xa9\x70\x46\x76\x9d\xa7\x2f\x28\x6b\x2d\xd1\xa2\x64\x77\x91\xcd\xc7\x51\x9c\x21\x35\x9c\xff\x6f\x86\x94\x5e\x2e\xfe\x58\xde\xbc\x48\x92\x57\xab\xc5\xf5\x7a\x21\xd6\xd7\x2f\x5f\x2f\xc4\xf2\x77\x71\xf3\x76\x2d\x16\x1f\x96\xef\xd7\xef\x85\x75\x85\x32\xf8\x59\x31\x5a\x43\x72\x0b\x8e\xd0\x1a\xc8\xc5\xe3\x44\x08\x6a\xaa\x9f\x7f\x7d\x2e\xb2\x52\x39\x95\x31\x38\xb1\x55\x6e\x8f\xa6\x78\xfc\xfc\x97\x27\xe2\xdd\x6a\xf9\xe6\x7a\xf5\x51\xfc\xb9\xf8\x78\x95\x08\xd1\xcd\x24\x81\x86\xa1\x00\x27\xae\x57\xab\xeb\x8f\x57\x49\x22\x84\xda\x2a\x56\x4e\x36\x4e\x0b\x86\x7f\xd9\x7b\xa7\xa8\x35\x9a\x42\x42\xa5\xb0\x1f\xcd\xac\xd6\x2a\xb5\x4e\xb1\x75\x24\x52\x2c\xd0\x1c\xc6\x1d\x28\x86\x5c\x2a\x16\x8c\x15\x10\xab\xaa\xe6\xcf\xde\x92\x03\x65\x0e\x6b\x1f\xfc\x71\x95\x70\xcd\x92\x2b\x3d\x7c\x33\xe6\x83\x85\xb5\xcd\x54\x30\x57\xdb\x02\xfb\x27\xa3\x2a\xe8\x1f\x6c\x0e\x12\xf3\xe3\xb3\xdd\x19\xc8\x65\xed\x70\xab\x18\xa4\x83\xda\x0e\x43\xae\x9b\x54\x63\x36\x1a\x66\xcb\x4a\x4f\x4e\xe2\x9d\x95\x1b\x95\xb1\x75\xd2\xc1\xa7\x06\x1d\x54\x60\x58\x82\x51\xa9\x86\x5c\xa4\xd6\x6a\x50\xc6\x7b\x36\x75\x1e\xc9\x48\xf2\xa4\xd7\x7a\x79\xf3\xdb\xe2\xc3\x25\x5a\x93\x78\x7b\x33\x4d\xc1\x9d\xd3\x93\x79\x8a\x1a\x02\xf7\x35\xe9\xb1\x03\x66\xaa\x5a\x99\x7d\xff\x3c\xc9\x4a\xc8\xc5\xc6\x6a\x6d\x77\x10\x70\x76\x18\x43\x53\x0c\xc6\x4a\x74\xe0\xd3\x3f\xcc\xfe\xb7\x62\xaa\x1b\x2f\x90\x38\xc2\xda\xc4\xf0\xe9\x2a\x84\x0c\x52\xe5\x15\x9a\xe1\x96\xe6\xc1\xbc\x1f\x6e\x01\x14\x2d\x66\x23\x4c\x2e\xc4\xab\x8d\x07\xd9\x3a\x84\x87\xa0\x4c\x08\xcf\x99\xd7\x5e\x56\xe0\x0a\x90\x99\xad\x2a\xe4\x61\x7e\x0e\x56\x07\xa9\x22\x38\x38\x8d\xad\xf4\xa9\x51\x54\x46\xac\x2e\x2b\x71\x1b\x16\x70\xa6\xad\x81\x00\xea\xb9\x26\xb7\x51\x8d\x66\x99\x3a\x65\xb2\xf2\x38\x21\xd6\xfb\x72\xa4\x51\xaf\xd8\x58\x77\x7b\xfa\x4c\x32\xb3\x8d\xe1\x21\xfb\x8d\xd6\x32\x80\xb3\x54\x24\x91\xa8\x01\x0a\xd8\x57\x24\x77\x78\x8b\xc1\x98\xad\xa0\x56\x05\x5c\xd8\x74\x95\x29\x9a\xa1\x77\x85\xce\xd9\xb0\xc2\xe7\xab\xa4\x06\xd3\x45\x36\xda\x86\xaf\x20\x27\x8f\xef\x6b\x21\xba\xf9\xeb\xf5\xeb\xde\xd6\x97\x63\xc4\xc8\xfb\x1a\xc6\xb6\xae\x2c\x86\x5b\xae\x1b\x2a\xa3\x62\x11\x95\xc3\x8d\x10\x2b\x57\xa8\xcf\x1e\xfb\xd3\x50\xd9\xd6\x98\x51\xeb\xf9\xf7\x3f\xc1\xfb\xe2\xe5\xe6\x2d\x3b\xc5\x59\x79\xba\xd8\xb9\x3a\x8c\x55\x4f\x5b\x8e\x53\x65\x75\x61\x55\x76\x12\x3c\x4c\xd7\x27\xc2\xc2\x00\x44\xf3\x93\xda\x7c\xd0\xf0\xb5\xa5\x68\xae\x3a\x43\xba\x9f\xc0\xa1\xb7\x4f\x20\xe1\xbb\x00\x18\xbe\xec\x26\x72\x06\xf9\x14\x74\x74\x2b\xda\x66\xb7\x61\xb9\x56\xa8\x81\xd8\x77\x87\x0e\xf9\xc0\xbf\xb7\x32\xb2\x8e\xb0\x7a\x5a\x2c\xa6\xa9\x52\x70\x83\x58\x8e\xa2\xef\xfb\x72\x0f\x56\x18\x38\xb4\x45\x31\xf6\x20\xf6\xd5\x70\xf7\x86\x3e\x90\x79\x76\x7d\xf3\x9f\x90\xa2\x35\x45\x54\x38\x47\x76\x48\x60\xcb\xf4\x18\xca\xff\x43\xb3\xbc\xd3\xfc\x81\xa8\x6e\xb8\xb4\x4e\x2a\x22\x9b\x61\x78\x4b\x08\x99\xbe\x1f\x66\x87\x2d\x04\x8a\xcf\xb2\xf1\x05\x50\xf8\xfa\x72\x8f\x24\xea\x65\x8f\xaa\x77\xa1\xfc\xb5\x3f\xff\xfc\xad\x1b\xe8\xc1\xd4\xcf\x73\x6c\x2f\xd8\x03\xc9\x66\xfb\xdc\x1c\x2e\xfe\x1e\xe2\x60\x33\x56\xa4\xb3\x9c\x51\xf6\xd4\x6b\x42\xde\xd6\x8d\x4a\x35\x61\xf1\x42\x46\x4c\x01\xca\xa5\x32\x05\xe4\x72\xe3\x1b\xd7\xb0\x93\x4e\xb6\xed\x48\xd7\x6d\xaf\x63\x97\x7e\x11\x6a\x38\x4d\x72\x09\x2a\x8f\x27\xab\xb3\x9c\x49\xd6\xa9\xd7\x44\xb2\x5a\xb7\x68\xb2\x5a\x4b\x3c\x59\xf7\x3c\x37\x2a\x85\x86\x15\xfa\x3b\x4b\xa6\x8c\xac\x6c\x8e\x9b\x7d\x70\x8c\x0c\x2e\xb2\xc7\xa0\x8e\x86\xd3\x2f\x9a\x76\x30\x1f\x8f\x44\x12\xdc\x19\x26\x8f\xd5\xde\x3e\x71\xac\xfe\x00\x07\x9c\x83\x2d\xc2\x4e\x46\x50\xfc\x6e\x8e\xbe\x68\xbf\x6a\x5b\xe1\x64\x27\xbb\x47\x27\x94\x87\x4c\x3c\x50\x43\x0c\x1b\xc5\x01\xd5\x81\xb2\xb3\xa5\x71\x8a\x41\x10\xf5\xf4\x11\xf8\xc5\x6f\x3f\xd4\xa4\x15\xf2\x37\x91\xfd\x54\x9c\x91\xfa\x31\xf5\x86\x10\x3c\x7b\x9a\xac\x4b\x38\x7c\x6c\x21\xf9\x98\x72\xff\x75\x78\xf8\xa7\xc2\x98\x6a\xe4\xbd\xd8\x21\x97\xa2\x28\x69\x6f\xb2\x2b\x91\x36\x1c\xbe\xe1\xae\x44\x92\xdc\x02\x09\x63\x59\x10\x5b\x07\x82\x4b\x10\x4b\x7f\x26\xbf\x3a\x38\x3c\x22\x61\x37\xe2\x5d\xa3\xf5\xea\x30\xf3\x11\xfd\x94\x2c\x0d\x31\xa8\x5c\x60\x37\x8b\xda\x69\x03\xa7\x55\x1b\x7f\xb7\xc4\x95\x20\x2b\x94\x48\x81\x3d\x71\x6d\xd0\x3b\xdb\xe8\x5c\xa4\x90\x44\x76\xdd\x87\xf6\xf4\xd9\xc5\xb8\x7f\x27\xd7\xbf\x51\x29\xcc\x9c\x77\xb8\xd9\xc8\xb2\x31\xb7\x97\x95\x0d\x1a\x7f\x90\xe9\xbd\x64\x3b\x53\x4c\xd6\x79\x93\xd2\x72\x1c\xc9\xd1\xd4\x96\x89\xdf\xc2\xe0\x8f\x93\xe2\xfe\xf7\x44\xcc\x7e\x59\x8d\xc6\xc4\x0c\x36\xf1\xa3\xdd\x5b\x67\x11\x1b\x57\xed\xd9\x5b\xec\xdb\x37\x6f\x96\xeb\x17\xc9\x7f\x01\x00\x00\xff\xff\x0a\xe0\x61\xd9\xde\x17\x00\x00")

func _000001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__000001_initUpSql,
		"000001_init.up.sql",
	)
}

func _000001_initUpSql() (*asset, error) {
	bytes, err := _000001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "000001_init.up.sql", size: 6110, mode: os.FileMode(420), modTime: time.Unix(1570529246, 0)}
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
	"000001_init.down.sql": _000001_initDownSql,
	"000001_init.up.sql":   _000001_initUpSql,
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
	"000001_init.down.sql": &bintree{_000001_initDownSql, map[string]*bintree{}},
	"000001_init.up.sql":   &bintree{_000001_initUpSql, map[string]*bintree{}},
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
