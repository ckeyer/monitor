package utils

import (
	"testing"
)

func TestCollect(t *testing.T) {
	var ns Collector = NewSysInfo()

	if err := ns.Collect(); err != nil {
		t.Error(err)
	}

	t.Logf("%+v", ns)

	t.Error("...")
}
