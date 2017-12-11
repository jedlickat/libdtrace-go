package dtrace

// #include <sys/types.h>
// #include <dtrace.h>
// int addProbe_cgo(dtrace_hdl_t *h, const dtrace_probedesc_t *pd, void *arg);
import "C"
import (
	"fmt"
	"unsafe"
)

// Probe description
type Probe struct {
	dpProbeDesc *C.dtrace_probedesc_t
}

// GetID returns probe ID
func (probe *Probe) GetID() int {
	return int(probe.dpProbeDesc.dtpd_id)
}

// GetProviderName returns provider name
func (probe *Probe) GetProviderName() string {
	return C.GoString(&probe.dpProbeDesc.dtpd_provider[0])
}

// GetModuleName returns module name
func (probe *Probe) GetModuleName() string {
	return C.GoString(&probe.dpProbeDesc.dtpd_mod[0])
}

// GetFuncName returns function name
func (probe *Probe) GetFuncName() string {
	return C.GoString(&probe.dpProbeDesc.dtpd_func[0])
}

// GetProbeName returns probe name
func (probe *Probe) GetProbeName() string {
	return C.GoString(&probe.dpProbeDesc.dtpd_name[0])
}

//export addProbe
func addProbe(handle *C.dtrace_hdl_t, pdesc *C.dtrace_probedesc_t,
	arg *C.char) int {
	probe := Probe{dpProbeDesc: pdesc}
	fmt.Printf("%d: %s: %s: %s: %s\n",
		probe.GetID(), probe.GetProviderName(), probe.GetModuleName(),
		probe.GetFuncName(), probe.GetProbeName())
	return 0
}

// GetProbes return
func GetProbes(handle *DTraceHandle) []Probe {
	// Iterate over all probes and fill in the array
	C.dtrace_probe_iter(handle.dhLibHandle, nil,
		(*C.dtrace_probe_f)(unsafe.Pointer(C.addProbe_cgo)), nil)

	return nil
}
