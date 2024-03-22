package system

import (
	"fmt"
	"os"
	"strings"

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
	totalMemory := vmStat.Total / 1024 / 1024
	usedMemory := vmStat.Used / 1024 / 1024
	si.RAM = fmt.Sprintf("%d%s / %d%s", usedMemory, "MB", totalMemory, "MB")
}

func (si *SystemInfor) getDisk() {
	totalDisk := diskStat.Total / 1024 / 1024
	usedDisk := diskStat.Used / 1024 / 1024
	si.Disk = fmt.Sprintf("%d%s / %d%s", usedDisk, "MB", totalDisk, "MB")
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

func (si *SystemInfor) GenInfoSys(disable []string) []string {
	// We want to display by order
	listSysInform := []string{
		fmt.Sprint(si.User + "@" + si.Hostname),
		"-----------------------------------",
		fmt.Sprintf("%s: %s", "Host", si.Hostname[len(si.User)+1:]),
		fmt.Sprintf("%s: %s", "Platform", si.Platform),
		fmt.Sprintf("%s: %s", "Terminal", si.TerminalFont),
		fmt.Sprintf("%s: %s", "CPU", si.CPU),
		fmt.Sprintf("%s: %s", "Memory", si.RAM),
		fmt.Sprintf("%s: %s", "Disk", si.Disk),
		fmt.Sprintf("%s: %s", "Uptime", si.Uptime),
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
