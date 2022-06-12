package subsystems

var (
	SubsystemInst = []Subsystem{
		&CpusetSubsystem{},
		&MemorySubsystem{},
		&CpuSubSystem{},
	}
)

type ResourceConfig struct {
	MemoryLimit string
	CpuShare    string
	CpuSet      string
}

type Subsystem interface {
	// return the subsystem name
	Name() string

	// set cgroup attribute in this subsystem
	Set(path string, config *ResourceConfig) error

	// insert a process into cgroup
	Apply(path string, pid int) error

	// remove a cgroup
	Remove(path string) error
}
