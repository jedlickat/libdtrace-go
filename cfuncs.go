package dtrace

/*
#cgo CFLAGS: -D_LARGEFILE64_SOURCE
#cgo LDFLAGS: -ldtrace

#include <sys/types.h>
#include <dtrace.h>

// Gateway for probe iterator
int addProbe_cgo(dtrace_hdl_t *h, const dtrace_probedesc_t *pd,
	void *arg) {
		int addProbe(dtrace_hdl_t *, const dtrace_probedesc_t *, void *);
		return addProbe(h, pd, arg);
}
*/
import "C"
