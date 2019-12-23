// Code generated by go-bindata. DO NOT EDIT.
// sources:
// metadata.tmpl (841B)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _metadataTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xcd\xce\x93\x40\x14\x86\xf7\x73\x15\xaf\x5f\xba\x80\x04\x99\xbd\xa6\x0b\xd3\x1a\xe3\xa2\x76\x61\xe3\xc6\xb8\x38\x85\x03\x21\x0c\x33\x66\x18\xc6\x28\x9d\x7b\x37\x50\xc0\x26\x16\x56\x1f\x3b\xce\xcf\x93\x87\xf7\x20\x25\x0e\x26\x67\x94\xac\xd9\x92\xe3\x1c\xd7\xdf\x28\xcd\xf2\x8e\x4a\x3b\xb6\x9a\x94\xcc\x9a\x5c\x36\xec\x28\x27\x47\xef\x71\x3c\xe3\xcb\xf9\x82\x8f\xc7\xcf\x97\x54\xfc\xa4\xac\xa6\x92\x31\xb7\x85\x90\x12\x1f\x94\x02\x79\xaa\x14\x5d\xd5\xbf\x56\x2a\x32\xa3\x5b\x87\x48\xf4\xfd\x5b\x58\xd2\x25\x63\x57\x27\xd8\x79\xbc\xdb\x23\x3d\x92\x23\x84\x20\x00\xa0\xef\xb1\xab\x71\x43\x46\x0d\xab\x03\xb5\x8c\x10\xb0\xc7\xcb\xbd\x1e\xc2\xcb\x88\x60\x9d\x0f\x0b\xb1\xd8\x06\x4a\x89\x4f\xec\x9e\x22\x7f\x55\x4a\xa1\x64\x87\x19\x0c\x4f\xaa\x63\x14\xd6\x34\x0f\xde\x45\xa7\x33\x44\x0d\x4e\x53\x25\x5e\x03\x46\x31\xa2\xa1\xee\x11\x42\x82\xab\x31\x2a\x06\xfa\xf1\x93\x7c\x02\x53\x0f\x5e\xcd\xf7\x67\x9b\x3f\xc6\xa1\xaa\xc0\x1b\x53\x4f\x1b\xc3\x63\xd9\x75\x56\xe3\xce\xbc\xe1\x0f\x5b\xf3\x6d\x34\x1c\xf8\x05\xa9\x96\xc7\xd1\x7b\x6a\xd3\xb0\x4f\x17\x87\x38\x81\xb3\x1d\x8b\x30\x9e\xe5\xd4\xb5\xee\xd5\x93\xd8\x80\x46\x31\x66\x91\x39\x85\x49\x71\x25\x83\x07\xf1\x49\xf9\xeb\x96\x6e\xfb\xbf\x6e\xa5\x9d\xd9\xd4\x5d\x01\x46\x7e\x51\x5d\x2e\xb6\x22\x89\x3d\xbc\x08\x8f\x3f\xe0\xdf\x00\x00\x00\xff\xff\x48\xb9\xfc\x43\x49\x03\x00\x00")

func metadataTmplBytes() ([]byte, error) {
	return bindataRead(
		_metadataTmpl,
		"metadata.tmpl",
	)
}

func metadataTmpl() (*asset, error) {
	bytes, err := metadataTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metadata.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xab, 0x26, 0xb1, 0x4b, 0x1, 0x74, 0xf, 0xb1, 0x7a, 0x27, 0x56, 0x65, 0xc9, 0xc2, 0xe1, 0xa5, 0x89, 0xb2, 0xcf, 0x73, 0xc4, 0x95, 0x1a, 0xd6, 0x47, 0xef, 0x4b, 0x9d, 0xa9, 0x69, 0x5, 0xb2}}
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

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
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

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"metadata.tmpl": metadataTmpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
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
	"metadata.tmpl": &bintree{metadataTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
