package utils

import (
	"io/ioutil"
	"net"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/cloudfoundry/gosigar"
	"github.com/prometheus/client_golang/prometheus"
)

type CollectType string

const (
	TypeHost      CollectType = "host"
	TypeContainer CollectType = "container"
	TypeRedis     CollectType = "redis"
)

type Collector interface {
	Collect() error
	Name() string
	Type() string
}

type SysInfo struct {
	sync.Mutex

	HostName  string
	IPAddress map[string]string

	CPUInfo CPUInfo
	FSInfo  FSInfo
	Mem     sigar.Mem
	Swap    sigar.Swap
	Updated time.Time

	collectItems []CollectFunc

	InfoCounter map[string]prometheus.CounterVec
	InfoGauge   map[string]prometheus.GaugeVec
	errors      *prometheus.CounterVec
}

type CPUInfo map[string]sigar.Cpu
type FSInfo map[string]sigar.FileSystemUsage

type CollectFunc func() error

func NewSysInfo() *SysInfo {
	si := &SysInfo{}
	si.GetHostname()
	si.GetIPAddress()

	si.AddCollect(
		si.GetCPU,
		si.GetMem,
		si.GetSwap,
		si.GetFS,
		si.GetNow,
	)

	return si
}

func (s *SysInfo) Collect() error {
	for _, hf := range s.collectItems {
		if err := hf(); err != nil {
			return err
		}
	}
	return nil
}

func (s *SysInfo) Name() string {
	return s.HostName
}

func (s *SysInfo) Type() CollectType {
	return TypeHost
}

// 添加收集项
func (s *SysInfo) AddCollect(cfs ...CollectFunc) {
	if len(cfs) == 0 {
		return
	}

	if len(s.collectItems) == 0 {
		s.collectItems = make([]CollectFunc, 0, len(cfs))
	}

	s.collectItems = append(s.collectItems, cfs...)
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
	cpu := sigar.Cpu{}
	list := sigar.CpuList{}
	if err := list.Get(); err != nil {
		return err
	}
	if err := cpu.Get(); err != nil {
		return err
	}

	cpuTable := make(CPUInfo)
	for index, c := range list.List {
		cpuTable[strconv.Itoa(index+1)] = c
	}
	cpuTable["total"] = cpu

	s.CPUInfo = cpuTable
	return nil
}

// 文件系统的使用情况
func (s *SysInfo) GetFS() error {
	list := sigar.FileSystemList{}
	if err := list.Get(); err != nil {
		return err
	}

	fsTable := make(FSInfo)
	for _, fs := range list.List {
		usage := sigar.FileSystemUsage{}
		if err := usage.Get(fs.DirName); err != nil {
			return err
		}
		fsTable[fs.DevName] = usage
	}

	s.FSInfo = fsTable
	return nil
}

func (s *SysInfo) GetMem() error {
	if err := s.Mem.Get(); err != nil {
		return err
	}
	return nil
}

func (s *SysInfo) GetSwap() error {
	if err := s.Swap.Get(); err != nil {
		return err
	}
	return nil
}

func (s *SysInfo) GetNow() error {
	s.Updated = time.Now()
	return nil
}
