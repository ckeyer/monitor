package client

import (
	"testing"
	"time"

	"github.com/cloudfoundry/gosigar"
)

func TestNetwork(t *testing.T) {
	ns := NewSysInfo()

	ns.GetCPU()
	ns.GetFS()

	t.Logf("ip tables: %+v", ns)

	t.Error("...")
}

func TestCollect(t *testing.T) {
	return
	cpus := sigar.CpuList{}
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		cpus.Get()
		for k, c := range cpus.List {
			t.Logf("cpu(%d): %+v", k, c)
		}
		t.Log()
	}

	cpu := sigar.Cpu{}
	cpu.Get()

	t.Logf("%+v", cpu)
	t.Error("...")
}
