package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
	"github.com/spf13/cobra"
	//"github.com/itchyny/gojq"
)

type ProcessInfo struct {
	PID        int32
	Name       string
	CPUPercent float64
	MemPercent float32
	MemoryMB   uint64
}

var getCpu bool
var getMem bool
var getDisk bool
var getTop bool
var verbose bool
var sysInfoCmd = &cobra.Command{
	Use:     "sysinfo",
	Aliases: []string{"sys", "sysinfo"},
	Short:   "Get System information",
	Long:    "Get System information of the server",
	Run:     run,
}

func init() {
	sysInfoCmd.Flags().BoolVarP(&getCpu, "cpu", "c", false, "Get CPU information")
	sysInfoCmd.Flags().BoolVarP(&getMem, "mem", "m", false, "Get Memory information")
	sysInfoCmd.Flags().BoolVarP(&getDisk, "disk", "d", false, "Get Disk information")
	sysInfoCmd.Flags().BoolVarP(&getTop, "top", "t", false, "Get Top 5 processes by CPU and Memory usage")
	sysInfoCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Get verbose output")
	rootCmd.AddCommand(sysInfoCmd)
}

func run(cmd *cobra.Command, args []string) {
	if getCpu {
		physical, _ := cpu.Counts(false)
		logical, _ := cpu.Counts(true)
		fmt.Printf("Physical CPUs: %v, Logical CPUs: %v\n", physical, logical)

		fmt.Println("Waiting for 3 seconds...")
		totalPercent, _ := cpu.Percent(3*time.Second, false)
		fmt.Println("------- Total CPU Usage -------")
		for i, percent := range totalPercent {
			fmt.Printf("Total CPU %d: %f%%\n", i, percent)
		}
		if verbose {
			fmt.Println("----------------------------------------")
			fmt.Println("Waiting for 3 seconds...")
			perPercent, _ := cpu.Percent(3*time.Second, true)
			fmt.Println("------- Per logical CPUs Usage -------")
			for i, percent := range perPercent {
				fmt.Printf("CPU %d: %f%%\n", i, percent)
			}
			fmt.Println("----------------------------------------")
		}
	}
	if getMem {
		v, _ := mem.VirtualMemory()
		totalMB := v.Total / 1024 / 1024
		freeMB := v.Free / 1024 / 1024
		usedMB := v.Used / 1024 / 1024
		fmt.Printf("Total: %v MB, Free: %v MB, Used: %v MB, UsedPercent: %f%%\n", totalMB, freeMB, usedMB, v.UsedPercent)
		if verbose {
			s, _ := mem.SwapMemory()
			sin := s.Sin / 1024 / 1024 // Convert to MB
			fmt.Printf("Sin: %v MB\n", sin)

			sout := s.Sout / 1024 / 1024 // Convert to MB
			fmt.Printf("Sout: %v MB\n", sout)

			swapTotal := s.Total / 1024 / 1024 // Convert to MB
			fmt.Printf("SwapTotal: %v MB\n", swapTotal)

			swapFree := s.Free / 1024 / 1024 // Convert to MB
			fmt.Printf("SwapFree: %v MB\n", swapFree)

			swapUsed := uint64(s.UsedPercent)
			fmt.Printf("SwapUsedPercent: %v%%\n", swapUsed)

		}
	}
	if getDisk {
		diskUsage, _ := disk.Usage("/")
		totalMB := diskUsage.Total / 1024 / 1024
		freeMB := diskUsage.Free / 1024 / 1024
		usedMB := diskUsage.Used / 1024 / 1024
		fmt.Printf("Total: %v MB, Free: %v MB, Used: %v MB, UsedPercent: %f%%\n", totalMB, freeMB, usedMB, diskUsage.UsedPercent)
		if verbose {
			partitions, err := disk.Partitions(true)
			if err != nil {
				fmt.Printf("Error getting disk partitions: %v\n", err)
				return
			}
			for _, partition := range partitions {
				usage, err := disk.Usage(partition.Mountpoint)
				if err != nil {
					fmt.Printf("Error getting disk usage for partition %s: %v\n", partition.Mountpoint, err)
					continue
				}
				totalMB := usage.Total / 1024 / 1024
				freeMB := usage.Free / 1024 / 1024
				usedMB := usage.Used / 1024 / 1024
				fmt.Printf("Partition: %s, Total: %v MB, Free: %v MB, Used: %v MB, UsedPercent: %f%%\n", partition.Mountpoint, totalMB, freeMB, usedMB, usage.UsedPercent)
			}
		}
	}
	if getTop {
		getTopProcesses()
	}
}

func getTopProcesses() {
	pids, err := process.Pids()
	if err != nil {
		fmt.Printf("Error getting process list: %v\n", err)
		return
	}

	var processes []ProcessInfo

	for _, pid := range pids {
		p, err := process.NewProcess(pid)
		if err != nil {
			continue
		}

		name, err := p.Name()
		if err != nil {
			continue
		}

		cpuPercent, err := p.CPUPercent()
		if err != nil {
			cpuPercent = 0
		}

		memPercent, err := p.MemoryPercent()
		if err != nil {
			memPercent = 0
		}

		memInfo, err := p.MemoryInfo()
		var memoryMB uint64 = 0
		if err == nil {
			memoryMB = memInfo.RSS / 1024 / 1024 // Convert to MB
		}

		processes = append(processes, ProcessInfo{
			PID:        pid,
			Name:       name,
			CPUPercent: cpuPercent,
			MemPercent: memPercent,
			MemoryMB:   memoryMB,
		})
	}

	if verbose {
		fmt.Println("\nAll processes:")
		fmt.Printf("%-8s %-20s %-10s %-10s %-10s\n", "PID", "NAME", "CPU%", "MEM%", "MEM(MB)")
		fmt.Println("----------------------------------------------------------------")
		for _, p := range processes {
			fmt.Printf("%-8d %-20s %-10.2f %-10.2f %-10d\n",
				p.PID, p.Name, p.CPUPercent, p.MemPercent, p.MemoryMB)
		}
		return
	}

	// Sort by CPU usage (descending)
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].CPUPercent > processes[j].CPUPercent
	})

	fmt.Println("\nTop 5 processes by CPU usage:")
	fmt.Printf("%-8s %-20s %-10s %-10s %-10s\n", "PID", "NAME", "CPU%", "MEM%", "MEM(MB)")
	fmt.Println("----------------------------------------------------------------")
	for i := 0; i < 5 && i < len(processes); i++ {
		p := processes[i]
		fmt.Printf("%-8d %-20s %-10.2f %-10.2f %-10d\n",
			p.PID, p.Name, p.CPUPercent, p.MemPercent, p.MemoryMB)
	}

	// Sort by Memory usage (descending)
	sort.Slice(processes, func(i, j int) bool {
		return processes[i].MemPercent > processes[j].MemPercent
	})

	fmt.Println("\nTop 5 processes by Memory usage:")
	fmt.Printf("%-8s %-20s %-10s %-10s %-10s\n", "PID", "NAME", "CPU%", "MEM%", "MEM(MB)")
	fmt.Println("----------------------------------------------------------------")
	for i := 0; i < 5 && i < len(processes); i++ {
		p := processes[i]
		fmt.Printf("%-8d %-20s %-10.2f %-10.2f %-10d\n",
			p.PID, p.Name, p.CPUPercent, p.MemPercent, p.MemoryMB)
	}
}
