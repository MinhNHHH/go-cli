package system

import (
	"fmt"
	"os"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfor struct {
	User         string
	Hostname     string
	Platform     string
	CPU          string
	Terminal     string
	TerminalFont string
	RAM          string
	Disk         string
	Uptime       string
}

var (
	hostStat, _ = host.Info()
	cpuStat, _  = cpu.Info()
	vmStat, _   = mem.VirtualMemory()
	diskStat, _ = disk.Usage("/") // If you're in Unix change this "\\" for "/"
)

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

func (si *SystemInfor) getUser() {
	si.User = os.Getenv("USER")
}

func (si *SystemInfor) getTerminal() {
	si.Terminal = os.Getenv("TERM_PROGRAM")
}

func (si *SystemInfor) getTerminalFont() {
	si.TerminalFont = os.Getenv("TERM")
}

func (si *SystemInfor) getPlatform() {
	si.Platform = fmt.Sprintf("%s %s %s", hostStat.Platform, hostStat.PlatformVersion, hostStat.KernelArch)
}

func (si *SystemInfor) getCPU() {
	si.CPU = cpuStat[0].ModelName
}

func (si *SystemInfor) getRam() {
	si.RAM = fmt.Sprintf("%dGB", vmStat.Total/1024/1024/1024)
}

func (si *SystemInfor) getDisk() {
	totalDisk := diskStat.Total / 1024 / 1024 / 1024
	usedDisk := diskStat.Used / 1024 / 1024 / 1024
	si.Disk = fmt.Sprintf("%d%s / %d%s", usedDisk, "GB", totalDisk, "GB")
}

func (si *SystemInfor) getUptime() {
	uptime := hostStat.Uptime
	days, hours, mins := uptimeToDaysHoursMins(uptime)

	if days == 0 {
		si.Uptime = fmt.Sprintf("%d hours, %d mins", hours, mins)
		return
	} else if days == 0 && hours == 0 {
		si.Uptime = fmt.Sprintf("%d mins", mins)
		return
	}
	si.Uptime = fmt.Sprintf("%d days, %d hours, %d mins", days, hours, mins)
}

func (si *SystemInfor) getHostName() {
	si.Hostname = hostStat.Hostname
}

func (si *SystemInfor) List() []string {
	// We want to display by order
	listSysInform := []string{}
	listSysInform = append(listSysInform, si.User)
	listSysInform = append(listSysInform, "-----------------------------------")
	listSysInform = append(listSysInform, fmt.Sprintf("%s: %s", "Host", si.Hostname))
	listSysInform = append(listSysInform, fmt.Sprintf("%s: %s", "Platform", si.Platform))
	listSysInform = append(listSysInform, fmt.Sprintf("%s: %s", "Terminal", si.Terminal))
	listSysInform = append(listSysInform, fmt.Sprintf("%s: %s", "Terminal Font", si.TerminalFont))
	listSysInform = append(listSysInform, fmt.Sprintf("%s: %s", "CPU", si.CPU))
	listSysInform = append(listSysInform, fmt.Sprintf("%s: %s", "Memory", si.RAM))
	listSysInform = append(listSysInform, fmt.Sprintf("%s: %s", "Disk", si.Disk))
	listSysInform = append(listSysInform, fmt.Sprintf("%s: %s", "Uptime", si.Uptime))
	return listSysInform
}

func System() *SystemInfor {
	info := &SystemInfor{}
	info.getHostName()
	info.getCPU()
	info.getUser()
	info.getDisk()
	info.getPlatform()
	info.getRam()
	info.getTerminal()
	info.getTerminalFont()
	info.getUptime()

	return info
}

// GPU infor x
// Dynamic change color
// Dynamic change ascii
