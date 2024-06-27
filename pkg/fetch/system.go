package fetch

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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

func removeEmptyString(s []string) []string {
	i := 0
	for i < len(s) {
		if s[i] == "" {
			s = append(s[:i], s[i+1:]...)
		}
		i++
	}
	return s
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

func getUser() string {
	return os.Getenv("USER")
}

func getShell() string {
	return os.Getenv("SHELL")
}

func getTerminal() string {
	return os.Getenv("TERM_PROGRAM")
}

func getCPU() CPUInfor {
	cpuStat, err := cpu.Info()
	if err != nil {
		log.Fatalf("error when getting cpu information: %s", err.Error())
	}
	if len(cpuStat) == 0 {
		log.Fatalf("can not get cpu information")
	}
	cpuInfor := CPUInfor{
		VendorId:  cpuStat[0].VendorID,
		Model:     cpuStat[0].Model,
		ModelName: cpuStat[0].ModelName,
		Mhz:       cpuStat[0].Mhz,
		CacheSize: cpuStat[0].CacheSize,
	}
	return cpuInfor
}

func getVM() VMInfor {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalf("error when getting vm information: %s", err.Error())
	}
	vmInfor := VMInfor{
		Total:       vmStat.Total,
		Available:   vmStat.Available,
		Used:        vmStat.Used,
		UsedPercent: vmStat.UsedPercent,
		Free:        vmStat.Free,
		Active:      vmStat.Active,
		Inactive:    vmStat.Inactive,
	}
	return vmInfor
}

func getDisk() DiskInfo {
	diskStat, err := disk.Usage("/") // If you're in Unix change this "\\" for "/"
	if err != nil {
		log.Fatalf("error when getting disk information: %s", err.Error())
	}
	diskInfor := DiskInfo{
		Total:       diskStat.Total,
		Used:        diskStat.Used,
		UsedPercent: diskStat.UsedPercent,
		Free:        diskStat.Free,
	}

	return diskInfor
}

func getGPUInfo() GPUInfo {
	gpu, err := ghw.GPU()
	if err != nil {
		log.Fatalf("error when getting gpu information: %s", err.Error())
	}
	if len(gpu.GraphicsCards) == 0 {
		log.Fatalf("cannot get gpu information")
	}
	gpuInfor := GPUInfo{
		ProductName: gpu.GraphicsCards[0].DeviceInfo.Product.Name,
		VendorName:  gpu.GraphicsCards[0].DeviceInfo.Vendor.Name,
	}

	return gpuInfor
}

func getHostName() HostNameInfor {
	hostStat, err := host.Info()
	if err != nil {
		log.Fatalf("error when getting hostname information: %s", err.Error())
	}
	hostName := HostNameInfor{
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
	return hostName
}

func (si SystemInfor) GetUptime() string {
	uptime := si.HostName.UpTime
	days, hours, mins := uptimeToDaysHoursMins(uptime)

	if days > 0 {
		return fmt.Sprintf("%d days, %d hours, %d mins", days, hours, mins)
	} else if hours > 0 {
		return fmt.Sprintf("%d hours, %d mins", hours, mins)
	} else {
		return fmt.Sprintf("%d mins", mins)
	}
}

func (si SystemInfor) GetHost() string {
	return si.HostName.HostName
}

func (si SystemInfor) GetOS() string {
	return si.HostName.OS
}

func (si SystemInfor) GetKernelVersion() string {
	return si.HostName.KernelVersion
}

func (si SystemInfor) GetCpu() string {
	return si.Cpu.ModelName
}

func (si SystemInfor) GetGpu() string {
	return si.Gpu.ProductName
}

func execLinuxCmd(command string) (string, error) {
	// Replace the command with your package manager's command
	cmd := exec.Command("sh", "-c", command) // Example for Debian-based systems

	// Execute the command
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	// Convert output to string
	outputStr := string(output)
	return outputStr, nil
}

func (si SystemInfor) GetMemmory() string {
	return fmt.Sprintf("%dMB / %dMB", si.Vm.Used/1024/1024, si.Vm.Total/1024/1024)
}

func (si SystemInfor) GetPackages() string {
	switch si.HostName.OS {
	case "linux":
		packages, err := execLinuxCmd("dpkg --list")

		if err != nil {
			fmt.Printf("err: %s", err)
			return ""
		}
		// Split output by lines
		lines := strings.Split(packages, "\n")
		packagesInstalled := 0

		// Count installed packages
		for _, line := range lines {
			fields := strings.Fields(line)
			if len(fields) > 1 && fields[0] == "ii" {
				packagesInstalled++
			}
		}

		return fmt.Sprintf("%d (dkpg)", packagesInstalled)
	case "darwin":
	case "windows":
	}
	return ""
}

func (si SystemInfor) GetResolution() string {
	switch si.HostName.OS {
	case "linux":
		cmd, err := execLinuxCmd("xrandr | grep '*' | awk '{print $1}'")
		if err != nil {
			fmt.Printf("err: %s", err)
			return ""
		}
		resolutions := removeEmptyString(strings.Split(cmd, "\n"))
		return strings.Join(resolutions, ", ")
	case "darwin":
	case "windows":
	}
	return ""
}

func (si SystemInfor) GetShell() string {
	switch si.HostName.OS {
	case "linux":
		shell := getShell()
		cmd, err := execLinuxCmd(fmt.Sprintf("%s --version | head -1 | cut -d ' ' -f 4", shell))
		if err != nil {
			fmt.Printf("err: %s", err)
			return ""
		}
		return fmt.Sprintf("%s %s", shell, strings.Trim(cmd, "\n"))
	}

	return ""
}

func (si SystemInfor) GetTheme() string {
	cmd, err := execLinuxCmd("gsettings get org.gnome.desktop.interface gtk-theme")
	if err != nil {
		fmt.Printf("err: %s", err)
		return ""
	}
	return strings.Trim(cmd, "\n")
}

func (si SystemInfor) GetIcons() string {
	cmd, err := execLinuxCmd("gsettings get org.gnome.desktop.interface icon-theme")
	if err != nil {
		fmt.Printf("err: %s", err)
		return ""
	}
	return strings.Trim(cmd, "\n")
}

func (si SystemInfor) formatInfo(label, info string) string {
	return fmt.Sprintf("%s%s: %s", label, placeHolder["${c0}"], info)
}

func (si SystemInfor) ListSysInfor(disable, seemore []string) []string {
	// We want to display by order
	listSysInform := []string{
		fmt.Sprintf(si.User + "@" + si.GetHost()),
		"-----------------------------------",
		si.formatInfo("OS", si.GetOS()),
		si.formatInfo("Host", si.GetHost()),
		si.formatInfo("Kernel", si.GetKernelVersion()),
		si.formatInfo("Uptime", si.GetUptime()),
		si.formatInfo("Packages", si.GetPackages()),
		si.formatInfo("Shell", si.GetShell()),
		si.formatInfo("Resolution", si.GetResolution()),
		si.formatInfo("Theme", si.GetTheme()),
		si.formatInfo("Icons", si.GetIcons()),
		si.formatInfo("Terminal", si.GetUptime()),
		si.formatInfo("CPU", si.GetCpu()),
		si.formatInfo("GPU", si.GetGpu()),
		si.formatInfo("Memory", si.GetMemmory()),
	}

	if len(disable) > 0 {
		for _, typeInfo := range disable {
			for index, str := range listSysInform {
				findDisable := strings.Contains(strings.ToLower(str), typeInfo)
				if findDisable {
					listSysInform = append(listSysInform[:index], listSysInform[index+1:]...)
				}
			}
		}
	}

	return listSysInform
}

func NewSysInfor() SystemInfor {
	return SystemInfor{
		User:     getUser(),
		Terminal: getTerminal(),
		HostName: getHostName(),
		Cpu:      getCPU(),
		Vm:       getVM(),
		Disk:     getDisk(),
		Gpu:      getGPUInfo(),
	}
}

//
// How to change ascii by request.
// Dynamic change color
// Dynamic change ascii
// format text color
// format ascii color
