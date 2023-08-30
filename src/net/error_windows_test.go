// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"errors"
	"syscall"
)

var (
	errTimedout       = syscall.ETIMEDOUT
	errOpNotSupported = syscall.EOPNOTSUPP

	abortedConnRequestErrors = []error{syscall.ERROR_NETNAME_DELETED, syscall.WSAECONNRESET} // see accept in fd_windows.go
)

func isPlatformError(err error) bool {
	_, ok := err.(syscall.Errno)
	return ok
}

func isENOBUFS(err error) bool {
	// syscall.ENOBUFS is a completely made-up value on Windows: we don't expect
	// a real system call to ever actually return it. However, since it is already
	// defined in the syscall package we may as well check for it.
	return errors.Is(err, syscall.ENOBUFS)
}

func isECONNRESET(err error) bool {
	return errors.Is(err, syscall.ECONNRESET)
}

func isWSAECONNREFUSED(err error) bool {
	const WSAECONNREFUSED = syscall.Errno(10061)
	return errors.Is(err, WSAECONNREFUSED)
}
