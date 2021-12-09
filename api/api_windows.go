package api

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/sys/windows"
)

var (
	user32                                  = windows.MustLoadDLL("user32.dll")
	dxva2                                   = windows.MustLoadDLL("Dxva2.dll")
	enumDisplayMonitors                     = user32.MustFindProc("EnumDisplayMonitors")
	setMonitorBrightness                    = dxva2.MustFindProc("SetMonitorBrightness")
	getPhysicalMonitorsFromHMONITOR         = dxva2.MustFindProc("GetPhysicalMonitorsFromHMONITOR")
	getNumberOfPhysicalMonitorsFromHMONITOR = dxva2.MustFindProc("GetNumberOfPhysicalMonitorsFromHMONITOR")
)

type RECT struct {
	Left, Top, Right, Bottom int32
}

type Handle uintptr
type PhysicalMonitor struct {
	PhysicalMonitor            Handle
	PhysicalMonitorDescription [128]uint16
}

type HMONITOR Handle
type HDC uintptr
type LPRECT uintptr
type LPARAM uintptr

type WindowsDisplay struct {
}

// func GetDisplays(ctx context.Context) {
// 	display := make(chan WindowsDisplay)
// 	errrs := make(chan error)
// 	cb := syscall.NewCallback(func(h HMONITOR, u HDC, v LPRECT, g LPARAM) uintptr {
// 		var nMonitors uint32
// 		if _, _, err := getNumberOfPhysicalMonitorsFromHMONITOR.Call(uintptr(h), uintptr(unsafe.Pointer(&nMonitors))); err != nil {
//
// 		}
//
// 		physicalMonitor := make([]PhysicalMonitor, nMonitors)
// 		if _, _, err := getPhysicalMonitorsFromHMONITOR.Call(uintptr(h), uintptr(nMonitors), uintptr(unsafe.Pointer(&physicalMonitor[0]))); err != nil {
//
// 		}
// 		// spew.Dump(physicalMonitor)
//
// 		spew.Dump(windows.UTF16ToString(physicalMonitor[0].PhysicalMonitorDescription[:]))
// 		spew.Dump(a1, a2, a3)
//
// 		a1, a2, a3 = setMonitorBrightness.Call(uintptr(physicalMonitor[0].PhysicalMonitor), uintptr(brightness))
// 		spew.Dump(a1, a2, a3)
// 		fmt.Println("DONE")
// 		return 1
// 	})
//
// 	n1l := uintptr(unsafe.Pointer(nil))
// 	_, _, _ = enumDisplayMonitors.Call(n1l, n1l, cb, 0)
// }

func GetMonitors(brightness int) {
	cb := syscall.NewCallback(func(h HMONITOR, u HDC, v LPRECT, g LPARAM) uintptr {
		var nMonitors uint32
		a1, a2, a3 := getNumberOfPhysicalMonitorsFromHMONITOR.Call(uintptr(h), uintptr(unsafe.Pointer(&nMonitors)))
		spew.Dump(nMonitors)
		spew.Dump(a1, a2, a3)

		physicalMonitor := make([]PhysicalMonitor, nMonitors)
		a1, a2, a3 = getPhysicalMonitorsFromHMONITOR.Call(uintptr(h), uintptr(nMonitors), uintptr(unsafe.Pointer(&physicalMonitor[0])))
		// spew.Dump(physicalMonitor)

		spew.Dump(windows.UTF16ToString(physicalMonitor[0].PhysicalMonitorDescription[:]))
		spew.Dump(a1, a2, a3)

		a1, a2, a3 = setMonitorBrightness.Call(uintptr(physicalMonitor[0].PhysicalMonitor), uintptr(brightness))
		spew.Dump(a1, a2, a3)
		fmt.Println("DONE")
		return 1
	})
	getMonitors(cb)
}

func getMonitors(cb uintptr) {
	n1l := uintptr(unsafe.Pointer(nil))
	_, _, _ = enumDisplayMonitors.Call(n1l, n1l, cb, 0)
	fmt.Println("Done?")
}
