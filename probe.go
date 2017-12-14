package dtrace

// #include <sys/types.h>
// #include <dtrace.h>
// int addProbe_cgo(dtrace_hdl_t *h, const dtrace_probedesc_t *pd, void *arg);
import "C"
import (
	"unsafe"
)

// Probe description
type Probe struct {
	ProbeID		int
	ProviderName	string
	ModuleName	string
	FunctionName	string
	ProbeName	string
}

//
// For the moment we use global array to build a list of probes
// until we build something better to support parallel iterations.
var probes []Probe

//export addProbe
func addProbe(handle *C.dtrace_hdl_t, pdesc *C.dtrace_probedesc_t,
	arg *C.char) int {
	probe := Probe{
		ProbeID:	int(pdesc.dtpd_id),
		ProviderName:	C.GoString(&pdesc.dtpd_provider[0]),
		ModuleName:	C.GoString(&pdesc.dtpd_mod[0]),
		FunctionName:	C.GoString(&pdesc.dtpd_func[0]),
		ProbeName:	C.GoString(&pdesc.dtpd_name[0]),
	}
	probes = append(probes, probe)
	return 0
}

// GetProbes return
func GetProbes(handle *DTraceHandle) []Probe {
	probes = make([]Probe, 1)

	// Iterate over all probes and fill in the array
	C.dtrace_probe_iter(handle.dhLibHandle, nil,
		(*C.dtrace_probe_f)(unsafe.Pointer(C.addProbe_cgo)), nil)

	return probes
}
