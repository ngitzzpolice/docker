package windowsoci

// This file is a hack - essentially a mirror of OCI spec for Windows.

import (
	"fmt"

	"github.com/docker/go-connections/nat"
)

// WindowsSpec is the full specification for Windows containers.
type WindowsSpec struct {
	Spec

	// Windows is platform specific configuration for Windows based containers.
	Windows Windows `json:"windows"`
}

// Spec is the base configuration for the container.  It specifies platform
// independent configuration. This information must be included when the
// bundle is packaged for distribution.
type Spec struct {

	// Version is the version of the specification that is supported.
	Version string `json:"ociVersion"`
	// Platform is the host information for OS and Arch.
	Platform Platform `json:"platform"`
	// Process is the container's main process.
	Process Process `json:"process"`
	//	// Root is the root information for the container's filesystem.
	Root Root `json:"root"`
	// Hostname is the container's host name.
	Hostname string `json:"hostname,omitempty"`
	// Mounts profile configuration for adding mounts to the container's filesystem.
	Mounts []Mount `json:"mounts"`
	//	// Hooks are the commands run at various lifecycle events of the container.
	//	Hooks Hooks `json:"hooks"`
}

// Windows contains platform specific configuration for Windows based containers.
type Windows struct {
	// Resources contain information for handling resource constraints for the container
	Resources *Resources `json:"resources,omitempty"`
	// Networking contains the platform specific network settings for the container.
	Networking *Networking `json:"networking,omitempty"`
	// FirstStart is used for an optimization on first boot of Windows
	FirstStart bool `json:"first_start,omitempty"`
	// LayerFolder is the path to the current layer folder
	LayerFolder string `json:"layer_folder,omitempty"`
	// Layer paths of the parent layers
	LayerPaths []string `json:"layer_paths,omitempty"`
	// HvRuntime contains settings specific to Hyper-V containers, omitted if not using Hyper-V isolation
	HvRuntime *HvRuntime `json:"hv_runtime,omitempty"`
}

// Process contains information to start a specific application inside the container.
type Process struct {
	// Terminal creates an interactive terminal for the container.
	Terminal bool `json:"terminal"`
	// ConsoleSize contains the initial h,w of the console size
	InitialConsoleSize [2]int `json:"-"`
	// User specifies user information for the process.
	User User `json:"user"`
	// Args specifies the binary and arguments for the application to execute.
	Args []string `json:"args"`
	// ArgsEscaped specifies if the arguments are already escaped or still need to be escaped
	ArgsEscaped bool `json:"args_escaped"`
	// Env populates the process environment for the process.
	Env []string `json:"env,omitempty"`
	// Cwd is the current working directory for the process and must be
	// relative to the container's root.
	Cwd string `json:"cwd"`
}

// User contains the user information for Windows
// TODO Windows: Does this make sense?
type User struct {
	User string `json:"user,omitempty"`
}

// Root contains information about the container's root filesystem on the host.
type Root struct {
	// Path is the absolute path to the container's root filesystem.
	Path string `json:"path"`
	// Readonly makes the root filesystem for the container readonly before the process is executed.
	Readonly bool `json:"readonly"`
}

// Platform specifies OS and arch information for the host system that the container
// is created for.
type Platform struct {
	// OS is the operating system.
	OS string `json:"os"`
	// Arch is the architecture
	Arch string `json:"arch"`
}

// Mount specifies a mount for a container.
type Mount struct {
	// Destination is the path where the mount will be placed relative to the container's root.  The path and child directories MUST exist, a runtime MUST NOT create directories automatically to a mount point.
	Destination string `json:"destination"`
	// Type specifies the mount kind.
	Type string `json:"type"`
	// Source specifies the source path of the mount.  In the case of bind mounts
	// this would be the file on the host.
	Source string `json:"source"`
	// Options are fstab style mount options.
	Options []string `json:"options,omitempty"`
}

// HvRuntime contains settings specific to Hyper-V containers
type HvRuntime struct {
	// ImagePath is the path to the Utility VM image for this container
	ImagePath string `json:"image_path,omitempty"`
}

// Networking contains the platform specific network settings for the container
type Networking struct {
	// These are here temporarily to maintain support for TP4
	MacAddress string `json:"mac,omitempty"`
	Bridge     string `json:"bridge,omitempty"`
	IPAddress  string `json:"ip,omitempty"`
	// PortBindings is the port mapping between the exposed port in the
	// container and the port on the host.
	PortBindings nat.PortMap `json:"port_bindings,omitempty"`

	// Below this is what is needed for TP5 and going forward

	// List of endpoints to be attached to the container
	EndpointList []string `json:"endpoints,omitempty"`
}

// Hook specifies a command that is run at a particular event in the lifecycle of a container
type Hook struct {
	//	Path string   `json:"path"`
	//	Args []string `json:"args,omitempty"`
	//	Env  []string `json:"env,omitempty"`
}

// Hooks for container setup and teardown
// TODO Windows containerd: Is this needed?
//type Hooks struct {
//	// Prestart is a list of hooks to be run before the container process is executed.
//	// On Linux, they are run after the container namespaces are created.
//	Prestart []Hook `json:"prestart,omitempty"`
//	// Poststart is a list of hooks to be run after the container process is started.
//	Poststart []Hook `json:"poststart,omitempty"`
//	// Poststop is a list of hooks to be run after the container process exits.
//	Poststop []Hook `json:"poststop,omitempty"`
//}

// Storage contains storage resource management settings
type Storage struct {
	// Specifies maximum Iops for the system drive
	Iops *uint16 `json:"iops,omitempty"`
	// Specifies maximum bytes per second for the system drive
	Bps *uint16 `json:"bps,omitempty"`
}

// Memory contains memory settings for the container
type Memory struct {
	// Memory limit (in bytes).
	Limit *uint64 `json:"limit,omitempty"`
	// Memory reservation (in bytes).
	Reservation *uint64 `json:"reservation,omitempty"`
}

// CPU contains information for cpu resource management
type CPU struct {
	// Number of CPUs available to the container. This is an appoximation for Windows Server Containers.
	Count *uint64 `json:"count,omitempty"`
	// CPU shares (relative weight (ratio) vs. other containers with cpu shares). Range is from 1 to 10000.
	Shares *uint64 `json:"shares,omitempty"`
	// Percent of available CPUs usable by the container.
	Percent *uint64 `json:"percent,omitempty"`
}

// Network network resource management information
type Network struct {
	// Bandwidth is the maximum egress bandwidth in bytes per second
	Bandwidth *uint64 `json:"bandwidth,omitempty"`
}

// Resources has container runtime resource constraints
// TODO Windows containerd. This structure needs ratifying with the old resources
// structure used on Windows and the latest OCI spec.
type Resources struct {
	// Memory restriction configuration
	Memory *Memory `json:"memory,omitempty"`
	// CPU resource restriction configuration
	CPU *CPU `json:"cpu,omitempty"`
	// Storage restriction configuration
	Storage *Storage `json:"storage,omitempty"`
	// Network restriction configuration
	Network *Network `json:"network,omitempty"`
	// Sandbox size indicates the size to expand the system drive to if it is currently smaller
	SandboxSize *uint64 `json:"sandbox_size,omitempty"`
}

// HACK - below taken from execdriver\driver.go and driver_windows.go

//// Resources contains all resource configs for a driver.
//type Resources struct {
//	Memory            int64  `json:"memory"`
//	MemoryReservation int64  `json:"memory_reservation"`
//	CPUShares         int64  `json:"cpu_shares"`
//	BlkioWeight       uint16 `json:"blkio_weight"`
//}

//// ProcessConfig is the platform specific structure that describes a process
//// that will be run inside a container.
//type ProcessConfig struct {
//	exec.Cmd `json:"-"`
//	Tty        bool     `json:"tty"`
//	Entrypoint string   `json:"entrypoint"`
//	Arguments  []string `json:"arguments"`
//	Terminal   Terminal `json:"-"` // standard or tty terminal
//	ConsoleSize [2]int `json:"-"` // h,w of initial console size
//}

//// Network settings of the container
//type Network struct {
//	MacAddress string `json:"mac"`
//	Bridge     string `json:"bridge"`
//	IPAddress  string `json:"ip"`

//	// PortBindings is the port mapping between the exposed port in the
//	// container and the port on the host.
//	PortBindings nat.PortMap `json:"port_bindings"`

//	ContainerID string            `json:"container_id"` // id of the container to join network.
//}

//// Command wraps an os/exec.Cmd to add more metadata
//type Command struct {
//	ContainerPid  int           `json:"container_pid"` // the pid for the process inside a container
//	ID            string        `json:"id"`
//	MountLabel    string        `json:"mount_label"` // TODO Windows. More involved, but can be factored out
//	Mounts        []Mount       `json:"mounts"`
//	Network       *Network      `json:"network"`
//	ProcessConfig ProcessConfig `json:"process_config"` // Describes the init process of the container.
//	ProcessLabel  string        `json:"process_label"`  // TODO Windows. More involved, but can be factored out
//	Resources     *Resources    `json:"resources"`
//	Rootfs        string        `json:"rootfs"` // root fs of the container
//	WorkingDir    string        `json:"working_dir"`
//	TmpDir        string        `json:"tmpdir"` // Directory used to store docker tmpdirs.
//	FirstStart  bool     `json:"first_start"`  // Optimization for first boot of Windows
//	Hostname    string   `json:"hostname"`     // Windows sets the hostname in the execdriver
//	LayerFolder string   `json:"layer_folder"` // Layer folder for a command
//	LayerPaths  []string `json:"layer_paths"`  // Layer paths for a command
//	Isolation   string   `json:"isolation"`    // Isolation technology for the container
//	ArgsEscaped bool     `json:"args_escaped"` // True if args are already escaped
//	HvPartition bool     `json:"hv_partition"` // True if it's an hypervisor partition
//}

// This is a temporary hack, copied mostly from OCI for Windows support.

const (
	// VersionMajor is for an API incompatible changes
	VersionMajor = 0
	// VersionMinor is for functionality in a backwards-compatible manner
	VersionMinor = 3
	// VersionPatch is for backwards-compatible bug fixes
	VersionPatch = 0

	// VersionDev indicates development branch. Releases will be empty string.
	VersionDev = ""
)

// Version is the specification version that the package types support.
var Version = fmt.Sprintf("%d.%d.%d%s", VersionMajor, VersionMinor, VersionPatch, VersionDev)
