package system

import (
	"fmt"
	"os"

	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfor struct {
	User     string
	Terminal string
	HostName HostNameInfor
	Cpu      CPUInfor
	Gpu      GPUInfo
	Vm       VMInfor
	Disk     DiskInfo
}

type HostNameInfor struct {
	HostName        string
	UpTime          uint64
	BootTime        uint64
	Procs           uint64
	OS              string
	Platform        string
	PlatformFamily  string
	PlatformVersion string
	KernelVersion   string
	KernelArch      string
}

type CPUInfor struct {
	VendorId  string
	Model     string
	ModelName string
	Mhz       float64
	CacheSize int32
}

type VMInfor struct {
	Total       uint64
	Available   uint64
	Used        uint64
	UsedPercent float64
	Free        uint64
	Active      uint64
	Inactive    uint64
}

type DiskInfo struct {
	Total       uint64
	Free        uint64
	Used        uint64
	UsedPercent float64
}

type GPUInfo struct {
	ProductName string
	VendorName  string
}

func uptimeToDaysHoursMins(uptimeSeconds uint64) (days, hours, mins uint64) {
	// Calculate days
	days = uptimeSeconds / (24 * 3600)

	// Calculate remaining seconds after extracting days
	remainingSeconds := uptimeSeconds % (24 * 3600)

	// Calculate hours
	hours = remainingSeconds / 3600

	// Calculate remaining seconds after extracting hours
	remainingSeconds %= 3600

	// Calculate minutes
	mins = remainingSeconds / 60

	return days, hours, mins
}

func GetUser() string {
	return os.Getenv("USER")
}

func GetTerminal() string {
	return os.Getenv("TERM_PROGRAM")
}

func GetCPU() (*CPUInfor, error) {
	cpuStat, err := cpu.Info()
	if err != nil {
		return nil, fmt.Errorf("error when getting cpu information: %s", err.Error())
	}
	if len(cpuStat) == 0 {
		return nil, fmt.Errorf("can not get cpu information")
	}
	cpuInfor := &CPUInfor{
		VendorId:  cpuStat[0].VendorID,
		Model:     cpuStat[0].Model,
		ModelName: cpuStat[0].ModelName,
		Mhz:       cpuStat[0].Mhz,
		CacheSize: cpuStat[0].CacheSize,
	}
	return cpuInfor, nil
}

func GetVM() (*VMInfor, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("error when getting vm information: %s", err.Error())
	}
	vmInfor := &VMInfor{
		Total:       vmStat.Total,
		Available:   vmStat.Available,
		Used:        vmStat.Used,
		UsedPercent: vmStat.UsedPercent,
		Free:        vmStat.Free,
		Active:      vmStat.Active,
		Inactive:    vmStat.Inactive,
	}
	return vmInfor, nil
}

func GetDisk() (*DiskInfo, error) {
	diskStat, err := disk.Usage("/") // If you're in Unix change this "\\" for "/"
	if err != nil {
		return nil, fmt.Errorf("error when getting disk information: %s", err.Error())
	}
	diskInfor := &DiskInfo{
		Total:       diskStat.Total,
		Used:        diskStat.Used,
		UsedPercent: diskStat.UsedPercent,
		Free:        diskStat.Free,
	}

	return diskInfor, nil
}

func GetGPUInfo() (*GPUInfo, error) {
	gpu, err := ghw.GPU()
	if err != nil {
		return nil, fmt.Errorf("error when getting gpu information: %s", err.Error())
	}
	if len(gpu.GraphicsCards) == 0 {
		return nil, fmt.Errorf("cannot get gpu information")
	}
	gpuInfor := &GPUInfo{
		ProductName: gpu.GraphicsCards[0].DeviceInfo.Product.Name,
		VendorName:  gpu.GraphicsCards[0].DeviceInfo.Vendor.Name,
	}

	return gpuInfor, nil
}

func GetHostName() (*HostNameInfor, error) {
	hostStat, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("error when getting hostname information: %s", err.Error())
	}
	hostName := &HostNameInfor{
		HostName:        hostStat.Hostname,
		UpTime:          hostStat.Uptime,
		BootTime:        hostStat.BootTime,
		Procs:           hostStat.Procs,
		OS:              hostStat.OS,
		Platform:        hostStat.Platform,
		PlatformFamily:  hostStat.PlatformFamily,
		PlatformVersion: hostStat.PlatformVersion,
		KernelVersion:   hostStat.KernelVersion,
		KernelArch:      hostStat.KernelArch,
	}
	return hostName, nil
}

// func (si *SystemInfor) getUptime() {
// 	uptime := hostStat.Uptime
// 	days, hours, mins := uptimeToDaysHoursMins(uptime)

// 	if days == 0 {
// 		si.Uptime = fmt.Sprintf("%d hours, %d mins", hours, mins)
// 		return
// 	} else if days == 0 && hours == 0 {
// 		si.Uptime = fmt.Sprintf("%d mins", mins)
// 		return
// 	}
// 	si.Uptime = fmt.Sprintf("%d days, %d hours, %d mins", days, hours, mins)
// }

func (si *SystemInfor) PrintInfo(disable []string) []string {
	// We want to display by order
	// listSysInform := []string{
	// 	fmt.Sprint(si.User + "@" + si.Hostname),
	// 	"-----------------------------------",
	// 	fmt.Sprintf("%s: %s", "Host", si.Hostname[len(si.User)+1:]),
	// 	fmt.Sprintf("%s: %s", "Platform", si.Platform),
	// 	fmt.Sprintf("%s: %s", "Terminal", si.TerminalFont),
	// 	fmt.Sprintf("%s: %s", "CPU", si.CPU),
	// 	fmt.Sprintf("%s: %s", "GPU", si.GPU),
	// 	fmt.Sprintf("%s: %s", "Memory", si.RAM),
	// 	fmt.Sprintf("%s: %s", "Disk", si.Disk),
	// 	fmt.Sprintf("%s: %s", "Uptime", si.Uptime),
	// }

	// // Need to fix and refactor.
	// // This solution will error when 2 field is same name. such as: "Terminal, Terminal Front"
	// if len(disable) > 0 {
	// 	for _, typeInfo := range disable {
	// 		for index, str := range listSysInform {
	// 			findDisable := strings.Contains(strings.ToLower(str), typeInfo)
	// 			if findDisable {
	// 				listSysInform = append(listSysInform[:index], listSysInform[index+1:]...)
	// 			}
	// 		}
	// 	}
	// }
	return []string{}
}

func System() *SystemInfor {
	user := GetUser()
	terminal := GetTerminal()
	hostName, _ := GetHostName()
	cpuInfo, _ := GetCPU()
	gpuInfo, _ := GetGPUInfo()
	diskInfo, _ := GetDisk()
	vmInfo, _ := GetVM()

	info := &SystemInfor{
		User:     user,
		Terminal: terminal,
		HostName: *hostName,
		Cpu:      *cpuInfo,
		Vm:       *vmInfo,
		Disk:     *diskInfo,
		Gpu:      *gpuInfo,
	}

	return info
}

//
// How to change ascii by request.
// Dynamic change color
// Dynamic change ascii
// format text color
// format ascii color
