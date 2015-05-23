package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func main() {
	vmem, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory()
	fmt.Println(vmem)
	fmt.Println(swap)

	faces, _ := net.NetInterfaces()
	counters, _ := net.NetIOCounters(true)
	fmt.Println(faces)
	fmt.Println(counters)

	cpuCounts, _ := cpu.CPUCounts(true)
	cpuInfo, _ := cpu.CPUInfo()
	fmt.Println(cpuCounts)
	fmt.Println(cpuInfo)

}
