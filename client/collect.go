package client

import (
	"io/ioutil"
	"net"
	"os/exec"
	"strings"

	sigar "github.com/cloudfoundry/gosigar"
)

type SysInfo struct {
	HostName  string
	IPAddress map[string]string
	CPUList   sigar.CpuList
	CPU       sigar.Cpu
}

func NewSysInfo() *SysInfo {
	si := &SysInfo{}
	si.GetHostname()
	si.GetIPAddress()
	return si
}

func (s *SysInfo) GetHostname() error {
	hostname, err := exec.Command("hostname").CombinedOutput()
	if err != nil {
		hostname, err = ioutil.ReadFile("/etc/hostname")
	}
	if err != nil {
		hostname, err = ioutil.ReadFile("/rootfs/etc/hostname")
	}
	if err != nil {
		return err
	}

	s.HostName = strings.TrimSpace(string(hostname))
	return nil
}

func (s *SysInfo) GetIPAddress() error {
	ipTable := make(map[string]string)

	list, err := net.Interfaces()
	if err != nil {
		return err
	}

SKIP_DEV:
	for _, dev := range list {
		for _, v := range []string{"veth", "lo", "docker"} {
			if strings.HasPrefix(dev.Name, v) {
				continue SKIP_DEV
			}
		}

		addrs, err := dev.Addrs()
		if err != nil {
			return err
		}

		for _, addr := range addrs {
			if len(strings.Split(addr.String(), ".")) == 4 {
				ipTable[dev.Name] = addr.String()
				continue SKIP_DEV
			}
		}
	}
	s.IPAddress = ipTable

	return nil
}

func (s *SysInfo) GetCPU() error {
	s.CPU = sigar.Cpu{}
	if err := s.CPUList.Get(); err != nil {
		return err
	}
	s.CPU.Get()
	return nil
}
