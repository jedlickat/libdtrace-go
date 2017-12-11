package dtrace

import (
	"fmt"
)

// #include <sys/types.h>
// #include <dtrace.h>
import "C"

// DTraceHandle is the main handle for interaction with DTrace.
type DTraceHandle struct {
	dhLibHandle *C.dtrace_hdl_t
}

// NewDTraceHandle allocates new handle.
func NewDTraceHandle() *DTraceHandle {
	handle := new(DTraceHandle)
	var err C.int

	handle.dhLibHandle = C.dtrace_open(C.DTRACE_VERSION, 0, &err)
	if handle.dhLibHandle == nil {
		fmt.Printf("Failed to initialize DTrace\n")
		fmt.Printf("  %s\n", C.GoString(C.dtrace_errmsg(nil, err)))

		return nil
	}

	return handle
}

// FreeDTraceHandle releases handle.
func FreeDTraceHandle(handle *DTraceHandle) {

	if handle == nil {
		return
	}

	C.dtrace_close(handle.dhLibHandle)
}
