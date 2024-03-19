package system

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	Hostname  string `json:hostname`
	Platform  string `json:platform`
	CPU       string `json:cpu`
	RAM       uint64 `json:ram`
	TotalDisk uint64 `json:total_disk`
	UsedDisk  uint64 `json:used_disk`
}

func System() {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/") // If you're in Unix change this "\\" for "/"

	info := new(SystemInfo)
	info.Hostname = hostStat.Hostname
	info.Platform = fmt.Sprintf("%s %s %s", hostStat.Platform, hostStat.PlatformVersion, hostStat.KernelArch)
	info.CPU = cpuStat[0].ModelName
	info.RAM = vmStat.Total / 1024 / 1024
	info.TotalDisk = diskStat.Total / 1024 / 1024
	info.UsedDisk = diskStat.Used / 1024 / 1024

	fmt.Println("OS:", info.Platform)
	fmt.Println("Host:", info.Hostname)
	fmt.Println("CPU:", info.CPU)
	fmt.Printf("RAM: %d%s\n", info.RAM, "MiB")
	fmt.Printf("Disk: %d%s/%d%s\n", info.UsedDisk, "MB", info.TotalDisk, "MB")
}
