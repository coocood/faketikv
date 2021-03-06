// Copyright 2019-present PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"hash/crc32"
	"io"
	"os"

	"github.com/pingcap/errors"
)

// GetFileSize gets the file size of the file.
func GetFileSize(path string) (uint64, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return uint64(fi.Size()), nil
}

// FileExists checks whether a file exists in the given path. It also fails if the path points to a directory or there is an error when trying to check the file.
func FileExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

// DirExists checks whether a directory exists in the given path. It also fails if the path is a file rather a directory or there is an error checking whether it exists.
func DirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// DeleteFileIfExists deletes the file if the file exists.
func DeleteFileIfExists(path string) (bool, error) {
	err := os.Remove(path)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, errors.WithStack(err)
	}
	return true, nil
}

// CalcCRC32 Calculates the given file's CRC32 checksum.
func CalcCRC32(path string) (uint32, error) {
	digest := crc32.NewIEEE()
	f, err := os.Open(path)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	_, err = io.Copy(digest, f)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return digest.Sum32(), nil
}
