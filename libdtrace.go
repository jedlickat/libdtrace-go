package dtrace

import (
	"fmt"
)

// #cgo LDFLAGS: -ldtrace
// #define _LARGEFILE64_SOURCE 
// #include <sys/types.h>
// #include <dtrace.h>
import "C"

type dtraceHandle struct {
	dh_lib_handle	*C.dtrace_hdl_t
}

func NewDTraceHandle() *dtraceHandle {
	handle := new(dtraceHandle)
	var err C.int

	handle.dh_lib_handle = C.dtrace_open(C.DTRACE_VERSION, 0, &err)
	if handle.dh_lib_handle == nil {
		fmt.Printf("Failed to initialize DTrace\n");
		fmt.Printf("  %s\n", C.GoString(C.dtrace_errmsg(nil, err)))
	}

	return handle
}

func FreeDTraceHandle(handle *dtraceHandle) {
	C.dtrace_close(handle.dh_lib_handle)
	fmt.Println("Closing handle")
}
