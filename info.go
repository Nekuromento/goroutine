// Copyright 2016 Huan Du. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package goroutine

import (
	"unsafe"

	runtime1_8 "github.com/huandu/goroutine/hack/go1_8/runtime"
	runtime1_8_1 "github.com/huandu/goroutine/hack/go1_8_1/runtime"
	runtime1_8_2 "github.com/huandu/goroutine/hack/go1_8_2/runtime"
	runtime1_8_3 "github.com/huandu/goroutine/hack/go1_8_3/runtime"
	runtime1_9 "github.com/huandu/goroutine/hack/go1_9/runtime"
	runtime1_9_1 "github.com/huandu/goroutine/hack/go1_9_1/runtime"
	runtime1_9_2 "github.com/huandu/goroutine/hack/go1_9_2/runtime"
)

func getg() unsafe.Pointer

// GoroutineId return id of current goroutine.
// It's guaranteed to be unique globally during app's life time.
func GoroutineId() int64 {
	gp := getg()

	switch goVersionCode() {
	case _GO_VERSION1_8:
		return (*runtime1_8.Goroutine)(gp).Goid()
	case _GO_VERSION1_8_1:
		return (*runtime1_8_1.Goroutine)(gp).Goid()
	case _GO_VERSION1_8_2:
		return (*runtime1_8_2.Goroutine)(gp).Goid()
	case _GO_VERSION1_8_3:
		return (*runtime1_8_3.Goroutine)(gp).Goid()
	case _GO_VERSION1_9:
		return (*runtime1_9.Goroutine)(gp).Goid()
	case _GO_VERSION1_9_1:
		return (*runtime1_9_1.Goroutine)(gp).Goid()
	case _GO_VERSION1_9_2:
		return (*runtime1_9_2.Goroutine)(gp).Goid()

	default:
		panic("unsupported go version " + goVersion().String())
	}
}
