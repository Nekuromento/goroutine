// Copyright 2016 Huan Du. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package goroutine

import (
	"runtime"

	"github.com/huandu/goroutine/version"
)

type _GoVersion int

const (
	_GO_VERSION_INVALID _GoVersion = iota

	_GO_VERSION1_8
	_GO_VERSION1_8_1
	_GO_VERSION1_8_2
	_GO_VERSION1_8_3
	_GO_VERSION1_9
	_GO_VERSION1_9_1
	_GO_VERSION1_9_2
)

var (
	_goVersion     version.Version
	_goVersionCode = _GO_VERSION_INVALID
)

func init() {
	var err error
	v := runtime.Version()
	_goVersion, err = version.Parse(v)

	if err != nil {
		return
	} else if _goVersion.Equal("1", "8") {
		_goVersionCode = _GO_VERSION1_8
	} else if _goVersion.Equal("1", "8", "1") {
		_goVersionCode = _GO_VERSION1_8_1
	} else if _goVersion.Equal("1", "8", "2") {
		_goVersionCode = _GO_VERSION1_8_2
	} else if _goVersion.Equal("1", "8", "3") {
		_goVersionCode = _GO_VERSION1_8_3
	} else if _goVersion.Equal("1", "9") {
		_goVersionCode = _GO_VERSION1_9
	} else if _goVersion.Equal("1", "9", "1") {
		_goVersionCode = _GO_VERSION1_9_1
	} else if _goVersion.Equal("1", "9", "2") {
		_goVersionCode = _GO_VERSION1_9_2
	}
}

func goVersionCode() _GoVersion {
	return _goVersionCode
}

func goVersion() version.Version {
	return _goVersion
}
