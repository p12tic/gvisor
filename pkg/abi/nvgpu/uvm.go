// Copyright 2023 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nvgpu

// UVM ioctl commands.
const (
	// From kernel-open/nvidia-uvm/uvm_linux_ioctl.h:
	UVM_INITIALIZE   = 0x30000001
	UVM_DEINITIALIZE = 0x30000002

	// From kernel-open/nvidia-uvm/uvm_ioctl.h:
	UVM_CREATE_RANGE_GROUP             = 23
	UVM_DESTROY_RANGE_GROUP            = 24
	UVM_REGISTER_GPU_VASPACE           = 25
	UVM_UNREGISTER_GPU_VASPACE         = 26
	UVM_REGISTER_CHANNEL               = 27
	UVM_UNREGISTER_CHANNEL             = 28
	UVM_ENABLE_PEER_ACCESS             = 29
	UVM_DISABLE_PEER_ACCESS            = 30
	UVM_SET_RANGE_GROUP                = 31
	UVM_MAP_EXTERNAL_ALLOCATION        = 33
	UVM_FREE                           = 34
	UVM_REGISTER_GPU                   = 37
	UVM_UNREGISTER_GPU                 = 38
	UVM_PAGEABLE_MEM_ACCESS            = 39
	UVM_SET_PREFERRED_LOCATION         = 42
	UVM_DISABLE_READ_DUPLICATION       = 45
	UVM_MIGRATE_RANGE_GROUP            = 53
	UVM_TOOLS_READ_PROCESS_MEMORY      = 62
	UVM_TOOLS_WRITE_PROCESS_MEMORY     = 63
	UVM_MAP_DYNAMIC_PARALLELISM_REGION = 65
	UVM_UNMAP_EXTERNAL                 = 66
	UVM_ALLOC_SEMAPHORE_POOL           = 68
	UVM_VALIDATE_VA_RANGE              = 72
	UVM_CREATE_EXTERNAL_RANGE          = 73
	UVM_MM_INITIALIZE                  = 75
)

// +marshal
type UVM_INITIALIZE_PARAMS struct {
	Flags    uint64 `nvproxy:"same"`
	RMStatus uint32
	Pad0     [4]byte
}

// UVM_INITIALIZE_PARAMS flags, from kernel-open/nvidia-uvm/uvm_types.h.
const (
	UVM_INIT_FLAGS_MULTI_PROCESS_SHARING_MODE = 0x2
)

// +marshal
type UVM_CREATE_RANGE_GROUP_PARAMS struct {
	RangeGroupID uint64 `nvproxy:"same"`
	RMStatus     uint32
	Pad0         [4]byte
}

// +marshal
type UVM_DESTROY_RANGE_GROUP_PARAMS struct {
	RangeGroupID uint64 `nvproxy:"same"`
	RMStatus     uint32
	Pad0         [4]byte
}

// +marshal
type UVM_REGISTER_GPU_VASPACE_PARAMS struct {
	GPUUUID  NvUUID `nvproxy:"same"`
	RMCtrlFD int32
	HClient  Handle
	HVASpace Handle
	RMStatus uint32
}

// GetFrontendFD implements HasFrontendFD.GetFrontendFD.
func (p *UVM_REGISTER_GPU_VASPACE_PARAMS) GetFrontendFD() int32 {
	return p.RMCtrlFD
}

// SetFrontendFD implements HasFrontendFD.SetFrontendFD.
func (p *UVM_REGISTER_GPU_VASPACE_PARAMS) SetFrontendFD(fd int32) {
	p.RMCtrlFD = fd
}

// +marshal
type UVM_UNREGISTER_GPU_VASPACE_PARAMS struct {
	GPUUUID  NvUUID `nvproxy:"same"`
	RMStatus uint32
}

// +marshal
type UVM_REGISTER_CHANNEL_PARAMS struct {
	GPUUUID  NvUUID `nvproxy:"same"`
	RMCtrlFD int32
	HClient  Handle
	HChannel Handle
	Pad      [4]byte
	Base     uint64
	Length   uint64
	RMStatus uint32
	Pad0     [4]byte
}

// GetFrontendFD implements HasFrontendFD.GetFrontendFD.
func (p *UVM_REGISTER_CHANNEL_PARAMS) GetFrontendFD() int32 {
	return p.RMCtrlFD
}

// SetFrontendFD implements HasFrontendFD.SetFrontendFD.
func (p *UVM_REGISTER_CHANNEL_PARAMS) SetFrontendFD(fd int32) {
	p.RMCtrlFD = fd
}

// +marshal
type UVM_UNREGISTER_CHANNEL_PARAMS struct {
	GPUUUID  NvUUID `nvproxy:"same"`
	HClient  Handle
	HChannel Handle
	RMStatus uint32
}

// +marshal
type UVM_ENABLE_PEER_ACCESS_PARAMS struct {
	GPUUUIDA NvUUID `nvproxy:"same"`
	GPUUUIDB NvUUID
	RMStatus uint32
}

// +marshal
type UVM_DISABLE_PEER_ACCESS_PARAMS struct {
	GPUUUIDA NvUUID `nvproxy:"same"`
	GPUUUIDB NvUUID
	RMStatus uint32
}

// +marshal
type UVM_SET_RANGE_GROUP_PARAMS struct {
	RangeGroupID  uint64 `nvproxy:"same"`
	RequestedBase uint64
	Length        uint64
	RMStatus      uint32
	Pad0          [4]byte
}

// +marshal
type UVM_MAP_EXTERNAL_ALLOCATION_PARAMS struct {
	Base               uint64 `nvproxy:"same"`
	Length             uint64
	Offset             uint64
	PerGPUAttributes   [UVM_MAX_GPUS]UvmGpuMappingAttributes
	GPUAttributesCount uint64
	RMCtrlFD           int32
	HClient            Handle
	HMemory            Handle
	RMStatus           uint32
}

// GetFrontendFD implements HasFrontendFD.GetFrontendFD.
func (p *UVM_MAP_EXTERNAL_ALLOCATION_PARAMS) GetFrontendFD() int32 {
	return p.RMCtrlFD
}

// SetFrontendFD implements HasFrontendFD.SetFrontendFD.
func (p *UVM_MAP_EXTERNAL_ALLOCATION_PARAMS) SetFrontendFD(fd int32) {
	p.RMCtrlFD = fd
}

// +marshal
type UVM_MAP_EXTERNAL_ALLOCATION_PARAMS_V550 struct {
	Base               uint64 `nvproxy:"UVM_MAP_EXTERNAL_ALLOCATION_PARAMS"`
	Length             uint64
	Offset             uint64
	PerGPUAttributes   [UVM_MAX_GPUS_V2]UvmGpuMappingAttributes
	GPUAttributesCount uint64
	RMCtrlFD           int32
	HClient            Handle
	HMemory            Handle
	RMStatus           uint32
}

// GetFrontendFD implements HasFrontendFD.GetFrontendFD.
func (p *UVM_MAP_EXTERNAL_ALLOCATION_PARAMS_V550) GetFrontendFD() int32 {
	return p.RMCtrlFD
}

// SetFrontendFD implements HasFrontendFD.SetFrontendFD.
func (p *UVM_MAP_EXTERNAL_ALLOCATION_PARAMS_V550) SetFrontendFD(fd int32) {
	p.RMCtrlFD = fd
}

// +marshal
type UVM_FREE_PARAMS struct {
	Base     uint64 `nvproxy:"same"`
	Length   uint64
	RMStatus uint32
	Pad0     [4]byte
}

// +marshal
type UVM_REGISTER_GPU_PARAMS struct {
	GPUUUID     NvUUID `nvproxy:"same"`
	NumaEnabled uint8
	Pad         [3]byte
	NumaNodeID  int32
	RMCtrlFD    int32
	HClient     Handle
	HSMCPartRef Handle
	RMStatus    uint32
}

// GetFrontendFD implements HasFrontendFD.GetFrontendFD.
func (p *UVM_REGISTER_GPU_PARAMS) GetFrontendFD() int32 {
	return p.RMCtrlFD
}

// SetFrontendFD implements HasFrontendFD.SetFrontendFD.
func (p *UVM_REGISTER_GPU_PARAMS) SetFrontendFD(fd int32) {
	p.RMCtrlFD = fd
}

// +marshal
type UVM_UNREGISTER_GPU_PARAMS struct {
	GPUUUID  NvUUID `nvproxy:"same"`
	RMStatus uint32
}

// +marshal
type UVM_PAGEABLE_MEM_ACCESS_PARAMS struct {
	PageableMemAccess uint8 `nvproxy:"same"`
	Pad               [3]byte
	RMStatus          uint32
}

// +marshal
type UVM_SET_PREFERRED_LOCATION_PARAMS struct {
	RequestedBase     uint64 `nvproxy:"same"`
	Length            uint64
	PreferredLocation NvUUID
	RMStatus          uint32
	Pad0              [4]byte
}

// +marshal
type UVM_SET_PREFERRED_LOCATION_PARAMS_V550 struct {
	RequestedBase        uint64 `nvproxy:"UVM_SET_PREFERRED_LOCATION_PARAMS"`
	Length               uint64
	PreferredLocation    NvUUID
	PreferredCPUNumaNode int32
	RMStatus             uint32
}

// +marshal
type UVM_DISABLE_READ_DUPLICATION_PARAMS struct {
	RequestedBase uint64 `nvproxy:"same"`
	Length        uint64
	RMStatus      uint32
	Pad0          [4]byte
}

// +marshal
type UVM_MIGRATE_RANGE_GROUP_PARAMS struct {
	RangeGroupID    uint64 `nvproxy:"same"`
	DestinationUUID NvUUID
	RMStatus        uint32
	Pad0            [4]byte
}

// +marshal
type UVM_TOOLS_READ_PROCESS_MEMORY_PARAMS struct {
	Buffer    uint64 `nvproxy:"same"`
	Size      uint64
	TargetVA  uint64
	BytesRead uint64
	RMStatus  uint32
	Pad0      [4]byte
}

// +marshal
type UVM_TOOLS_WRITE_PROCESS_MEMORY_PARAMS struct {
	Buffer       uint64 `nvproxy:"same"`
	Size         uint64
	TargetVA     uint64
	BytesWritten uint64
	RMStatus     uint32
	Pad0         [4]byte
}

// +marshal
type UVM_MAP_DYNAMIC_PARALLELISM_REGION_PARAMS struct {
	Base     uint64 `nvproxy:"same"`
	Length   uint64
	GPUUUID  NvUUID
	RMStatus uint32
	Pad0     [4]byte
}

// +marshal
type UVM_UNMAP_EXTERNAL_PARAMS struct {
	Base     uint64 `nvproxy:"same"`
	Length   uint64
	GPUUUID  NvUUID
	RMStatus uint32
	Pad0     [4]byte
}

// +marshal
type UVM_ALLOC_SEMAPHORE_POOL_PARAMS struct {
	Base               uint64 `nvproxy:"same"`
	Length             uint64
	PerGPUAttributes   [UVM_MAX_GPUS]UvmGpuMappingAttributes
	GPUAttributesCount uint64
	RMStatus           uint32
	Pad0               [4]byte
}

// +marshal
type UVM_ALLOC_SEMAPHORE_POOL_PARAMS_V550 struct {
	Base               uint64 `nvproxy:"UVM_ALLOC_SEMAPHORE_POOL_PARAMS"`
	Length             uint64
	PerGPUAttributes   [UVM_MAX_GPUS_V2]UvmGpuMappingAttributes
	GPUAttributesCount uint64
	RMStatus           uint32
	Pad0               [4]byte
}

// +marshal
type UVM_VALIDATE_VA_RANGE_PARAMS struct {
	Base     uint64 `nvproxy:"same"`
	Length   uint64
	RMStatus uint32
	Pad0     [4]byte
}

// +marshal
type UVM_CREATE_EXTERNAL_RANGE_PARAMS struct {
	Base     uint64 `nvproxy:"same"`
	Length   uint64
	RMStatus uint32
	Pad0     [4]byte
}

// +marshal
type UVM_MM_INITIALIZE_PARAMS struct {
	UvmFD  int32 `nvproxy:"same"`
	Status uint32
}

// From kernel-open/nvidia-uvm/uvm_types.h:

const (
	UVM_MAX_GPUS    = NV_MAX_DEVICES
	UVM_MAX_GPUS_V2 = NV_MAX_DEVICES * NV_MAX_SUBDEVICES
)

// +marshal
type UvmGpuMappingAttributes struct {
	GPUUUID            NvUUID
	GPUMappingType     uint32
	GPUCachingType     uint32
	GPUFormatType      uint32
	GPUElementBits     uint32
	GPUCompressionType uint32
}
