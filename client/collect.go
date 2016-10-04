package client

import (
// "io/ioutil"
// "net"
// "os/exec"
// "strconv"
// "strings"
// "time"

// "github.com/ckeyer/monitor/types"
// sigar "github.com/cloudfoundry/gosigar"
)

// type SysInfo  types.SysInfo

// func NewSysInfo() *SysInfo {
// 	si := &SysInfo{}
// 	si.GetHostname()
// 	si.GetIPAddress()
// 	return si
// }

// func (s *SysInfo) GetHostname() error {
// 	hostname, err := exec.Command("hostname").CombinedOutput()
// 	if err != nil {
// 		hostname, err = ioutil.ReadFile("/etc/hostname")
// 	}
// 	if err != nil {
// 		hostname, err = ioutil.ReadFile("/rootfs/etc/hostname")
// 	}
// 	if err != nil {
// 		return err
// 	}

// 	s.HostName = strings.TrimSpace(string(hostname))
// 	return nil
// }

// func (s *SysInfo) GetIPAddress() error {
// 	ipTable := make(map[string]string)

// 	list, err := net.Interfaces()
// 	if err != nil {
// 		return err
// 	}

// SKIP_DEV:
// 	for _, dev := range list {
// 		for _, v := range []string{"veth", "lo", "docker"} {
// 			if strings.HasPrefix(dev.Name, v) {
// 				continue SKIP_DEV
// 			}
// 		}

// 		addrs, err := dev.Addrs()
// 		if err != nil {
// 			return err
// 		}

// 		for _, addr := range addrs {
// 			if len(strings.Split(addr.String(), ".")) == 4 {
// 				ipTable[dev.Name] = addr.String()
// 				continue SKIP_DEV
// 			}
// 		}
// 	}
// 	s.IPAddress = ipTable

// 	return nil
// }

// func (s *SysInfo) GetCPU() error {
// 	cpu := sigar.Cpu{}
// 	list := sigar.CpuList{}
// 	if err := list.Get(); err != nil {
// 		return err
// 	}
// 	if err := cpu.Get(); err != nil {
// 		return err
// 	}

// 	cpuTable := make(types.CPUInfo)
// 	for index, c := range list.List {
// 		cpuTable[strconv.Itoa(index+1)] = c
// 	}
// 	cpuTable["total"] = cpu

// 	s.CPUInfo = cpuTable
// 	return nil
// }

// // 文件系统的使用情况
// func (s *SysInfo) GetFS() error {
// 	list := sigar.FileSystemList{}
// 	if err := list.Get(); err != nil {
// 		return err
// 	}

// 	fsTable := make(types.FSInfo)
// 	for _, fs := range list.List {
// 		usage := sigar.FileSystemUsage{}
// 		if err := usage.Get(fs.DirName); err != nil {
// 			return err
// 		}
// 		fsTable[fs.DevName] = usage
// 	}

// 	s.FSInfo = fsTable
// 	return nil
// }

// func (s *SysInfo) GetMem() error {
// 	if err := s.Mem.Get(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *SysInfo) GetSwap() error {
// 	if err := s.Swap.Get(); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *SysInfo) GetNow() error {
// 	s.Updated = time.Now()
// 	return nil
// }
