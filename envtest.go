package envtest

import (
	"fmt"
	"strings"
	"os/exec"
)

func Vbox_linux() bool {
    var vbox_vm bool

    out, _ := exec.Command("/bin/bash", "-c", `timeout 5 grep -ri 'virtualbox' /proc/ 2>/dev/null`).CombinedOutput()

    if fmt.Sprintf("%s", out) != "" && strings.Contains(fmt.Sprintf("%s", out), `VirtualBoxVM`) == false {
 		vbox_vm = true
    } else {
        vbox_vm = false
    }

    return vbox_vm
}

func Vbox_windows() bool {
	var vbox_vm bool

    out, _ := exec.Command("powershell", "-NoProfile", "get-PnpDevice", "-PresentOnly").CombinedOutput()

    if strings.Contains(fmt.Sprintf("%s", out), `VBOX`) {
    	vbox_vm = true
    } else {
    	vbox_vm = false
    }

    return vbox_vm
}

func Vmware_windows() bool {
	var vmware_vm bool

	out, _ := exec.Command("powershell", "-NoProfile", "get-PnpDevice", "-PresentOnly").CombinedOutput()
	
	if strings.Contains(fmt.Sprintf("%s", out), `VMware`) {
		vmware_vm = true
	} else {
		vmware_vm = false
	}
	return vmware_vm
}

func Vmware_linux() bool {
	var vmware_vm bool

	out, _ := exec.Command("/bin/bash", "-c", `timeout 5 grep -ri "vmware" /proc 2>/dev/null`).CombinedOutput()
	if strings.Contains(fmt.Sprintf("%s", out), `VMware`) {
		vmware_vm = true
	} else {
		vmware_vm = false
	}

	return vmware_vm
}

func Docker_linux() bool {
	var container bool

    out, _ := exec.Command("/bin/bash", "-c", `timeout 5 grep -ri "docker/containers" /proc 2>/dev/null`).CombinedOutput()

    if strings.Contains(fmt.Sprintf("%s", out), `docker/container`) {
 		container = true
    } else {
        container = false
    }

    return container

}
/*
func docker_windows() bool {
	var container bool
	return container
}
*/
/*
func lxd_linux() bool {
	var container bool
	return container
}
*/
/*
func xen_linux() bool {
	var container bool
	return container
}
*/
