package SAPOscol

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AdminCode string

const (
	AdminCodeSAPOscolSetCollectionInterval AdminCode = "SAPOscol-SetCollectionInterval"

	AdminCodeSAPOscolSetTraceLevel AdminCode = "SAPOscol-SetTraceLevel"

	AdminCodeSAPOscolRefreshData AdminCode = "SAPOscol-RefreshData"
)

type RequestCode string

const (
	RequestCodeSAPOscolGetDirectoryFreespace RequestCode = "SAPOscol-GetDirectoryFreespace"

	RequestCodeSAPOscolGetSyslogExtract RequestCode = "SAPOscol-GetSyslogExtract"

	RequestCodeSAPOscolCallOsscript RequestCode = "SAPOscol-CallOsscript"

	RequestCodeSAPOscolGetHardwareConfig RequestCode = "SAPOscol-GetHardwareConfig"
)

type GetCpuConsumption struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetCpuConsumption"`
}

type GetCpuConsumptionResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetCpuConsumptionResponse"`

	Result *CPUConsumption `xml:"Result,omitempty"`
}

type GetSingleCpu struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetSingleCpu"`
}

type GetSingleCpuResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetSingleCpuResponse"`

	Result *ArrayOfCPU `xml:"Result,omitempty"`
}

type GetCpuHistory struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetCpuHistory"`
}

type GetCpuHistoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetCpuHistoryResponse"`

	Result *ArrayOfCPUHistory `xml:"Result,omitempty"`
}

type GetMemory struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetMemory"`
}

type GetMemoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetMemoryResponse"`

	Result *Memory `xml:"Result,omitempty"`
}

type GetMemoryHistory struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetMemoryHistory"`
}

type GetMemoryHistoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetMemoryHistoryResponse"`

	Result *ArrayOfMemoryHistory `xml:"Result,omitempty"`
}

type GetDisk struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetDisk"`
}

type GetDiskResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetDiskResponse"`

	Result *ArrayOfDisk `xml:"Result,omitempty"`
}

type GetDiskHistory struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetDiskHistory"`
}

type GetDiskHistoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetDiskHistoryResponse"`

	Result *ArrayOfDiskHistory `xml:"Result,omitempty"`
}

type GetFilesystem struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetFilesystem"`
}

type GetFilesystemResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetFilesystemResponse"`

	Result *ArrayOfFilesystem `xml:"Result,omitempty"`
}

type GetFilesystemHistory struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetFilesystemHistory"`
}

type GetFilesystemHistoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetFilesystemHistoryResponse"`

	Result *ArrayOfFilesystemHistory `xml:"Result,omitempty"`
}

type GetLAN struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetLAN"`
}

type GetLANResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetLANResponse"`

	Result *ArrayOfLAN `xml:"Result,omitempty"`
}

type GetLANHistory struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetLANHistory"`
}

type GetLANHistoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetLANHistoryResponse"`

	Result *ArrayOfLANHistory `xml:"Result,omitempty"`
}

type GetTopCPUProcesses struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetTopCPUProcesses"`
}

type GetTopCPUProcessesResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetTopCPUProcessesResponse"`

	Result *ArrayOfProcess `xml:"Result,omitempty"`
}

type GetProcPattern struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetProcPattern"`
}

type GetProcPatternResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetProcPatternResponse"`

	Result *ArrayOfProcessPattern `xml:"Result,omitempty"`
}

type GetConfiguration struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetConfiguration"`
}

type GetConfigurationResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetConfigurationResponse"`

	Result *ArrayOfOSParameter `xml:"Result,omitempty"`
}

type GetOsData struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOsData"`
}

type GetOsDataResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOsDataResponse"`

	Result *ArrayOfOsData `xml:"Result,omitempty"`
}

type GetOS390AddressSpaceResourceData struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390AddressSpaceResourceData"`

	SystemId int32 `xml:"SystemId,omitempty"`
}

type GetOS390AddressSpaceResourceDataResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390AddressSpaceResourceDataResponse"`

	Result *ArrayOfOS390AddressSpaceResourceData `xml:"Result,omitempty"`
}

type GetOS390CPUPagingActivity struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390CPUPagingActivity"`

	SystemId int32 `xml:"SystemId,omitempty"`
}

type GetOS390CPUPagingActivityResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390CPUPagingActivityResponse"`

	Result *ArrayOfOS390CPUPagingActivity `xml:"Result,omitempty"`
}

type GetOS390CPUPagingActivityHistory struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390CPUPagingActivityHistory"`

	SystemId int32 `xml:"SystemId,omitempty"`
}

type GetOS390CPUPagingActivityHistoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390CPUPagingActivityHistoryResponse"`

	Result *ArrayOfOS390CPUPagingActivityHistory `xml:"Result,omitempty"`
}

type GetOS390StorageResource struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390StorageResource"`

	SystemId int32 `xml:"SystemId,omitempty"`
}

type GetOS390StorageResourceResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390StorageResourceResponse"`

	Result *ArrayOfOS390StorageResource `xml:"Result,omitempty"`
}

type GetOS390StorageResourceHistory struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390StorageResourceHistory"`

	SystemId int32 `xml:"SystemId,omitempty"`
}

type GetOS390StorageResourceHistoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390StorageResourceHistoryResponse"`

	Result *ArrayOfOS390StorageResourceHistory `xml:"Result,omitempty"`
}

type GetOS390CollectorDetails struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390CollectorDetails"`

	SystemId int32 `xml:"SystemId,omitempty"`
}

type GetOS390CollectorDetailsResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOS390CollectorDetailsResponse"`

	Result *ArrayOfOS390CollectorDetails `xml:"Result,omitempty"`
}

type GetOs390WLMData struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOs390WLMData"`

	SystemId int32 `xml:"SystemId,omitempty"`
}

type GetOs390WLMDataResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetOs390WLMDataResponse"`

	POL *ArrayOfOS390WLMPolicy `xml:"POL,omitempty"`

	PolAllRead bool `xml:"PolAllRead,omitempty"`

	MAP *ArrayOfOS390WLMMap `xml:"MAP,omitempty"`

	MapAllRead bool `xml:"MapAllRead,omitempty"`

	ENC *ArrayOfOS390WLMEnclaves `xml:"ENC,omitempty"`

	EncAllRead bool `xml:"EncAllRead,omitempty"`

	SCL *ArrayOfOS390WLMServiceClass `xml:"SCL,omitempty"`

	SclAllRead bool `xml:"SclAllRead,omitempty"`

	SCP *ArrayOfOS390WLMServiceClassPeriodData `xml:"SCP,omitempty"`

	ScpAllRead bool `xml:"ScpAllRead,omitempty"`

	RTD *ArrayOfOS390WLMResponseTimeDistribution `xml:"RTD,omitempty"`

	RtdAllRead bool `xml:"RtdAllRead,omitempty"`
}

type GetAS400PoolAll struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetAS400PoolAll"`
}

type GetAS400PoolAllResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetAS400PoolAllResponse"`

	Result *AS400PoolAll `xml:"Result,omitempty"`
}

type GetAS400PoolData struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetAS400PoolData"`
}

type GetAS400PoolDataResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetAS400PoolDataResponse"`

	Result *ArrayOfAS400PoolData `xml:"Result,omitempty"`
}

type GetAS400PoolHistory struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetAS400PoolHistory"`
}

type GetAS400PoolHistoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetAS400PoolHistoryResponse"`

	Result *ArrayOfAS400PoolHistory `xml:"Result,omitempty"`
}

type SetAdminParameter struct {
	XMLName xml.Name `xml:"urn:SAPOscol SetAdminParameter"`

	Parameter *AdminCode `xml:"Parameter,omitempty"`

	Value int32 `xml:"Value,omitempty"`
}

type SetAdminParameterResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol SetAdminParameterResponse"`

	Rtc int32 `xml:"rtc,omitempty"`
}

type SendRequest struct {
	XMLName xml.Name `xml:"urn:SAPOscol SendRequest"`

	Request *RequestCode `xml:"Request,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type SendRequestResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol SendRequestResponse"`

	Result string `xml:"Result,omitempty"`

	Rtc int32 `xml:"rtc,omitempty"`
}

type SendRequestAsync struct {
	XMLName xml.Name `xml:"urn:SAPOscol SendRequestAsync"`

	Request *RequestCode `xml:"Request,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type SendRequestAsyncResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol SendRequestAsyncResponse"`

	Result string `xml:"Result,omitempty"`

	Rtc int32 `xml:"rtc,omitempty"`

	RequestIdentifier int32 `xml:"RequestIdentifier,omitempty"`
}

type GetResult struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetResult"`

	Request *RequestCode `xml:"Request,omitempty"`

	RequestIdentifier int32 `xml:"RequestIdentifier,omitempty"`
}

type GetResultResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetResultResponse"`

	Result string `xml:"Result,omitempty"`

	Rtc int32 `xml:"rtc,omitempty"`
}

type GetVersion struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetVersion"`
}

type GetVersionResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetVersionResponse"`

	Version string `xml:"Version,omitempty"`
}

type GetHwConfText struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetHwConfText"`
}

type GetHwConfTextResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetHwConfTextResponse"`

	HWFile *File `xml:"hWFile,omitempty"`
}

type GetHwConfXML struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetHwConfXML"`
}

type GetHwConfXMLResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol GetHwConfXMLResponse"`

	XmlFile *File `xml:"xmlFile,omitempty"`
}

type RequestLogonFile struct {
	XMLName xml.Name `xml:"urn:SAPOscol RequestLogonFile"`

	User string `xml:"user,omitempty"`
}

type RequestLogonFileResponse struct {
	XMLName xml.Name `xml:"urn:SAPOscol RequestLogonFileResponse"`

	Filename string `xml:"filename,omitempty"`
}

type CPUConsumption struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	NumberOfCpus int32 `xml:"NumberOfCpus,omitempty"`

	LoadAvgPerMin int32 `xml:"LoadAvgPerMin,omitempty"`

	LoadAvgPer5Min int32 `xml:"LoadAvgPer5Min,omitempty"`

	LoadAvgPer15Min int32 `xml:"LoadAvgPer15Min,omitempty"`

	Interrupts int32 `xml:"Interrupts,omitempty"`

	SystemCalls int32 `xml:"SystemCalls,omitempty"`

	ContextSwitches int32 `xml:"ContextSwitches,omitempty"`

	CpuUser int32 `xml:"CpuUser,omitempty"`

	CpuSystem int32 `xml:"CpuSystem,omitempty"`

	CpuIdle int32 `xml:"CpuIdle,omitempty"`

	IdleWithoutWait int32 `xml:"IdleWithoutWait,omitempty"`

	CpuWait int32 `xml:"CpuWait,omitempty"`
}

type CPU struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	CpuUser int32 `xml:"CpuUser,omitempty"`

	CpuSystem int32 `xml:"CpuSystem,omitempty"`

	CpuIdle int32 `xml:"CpuIdle,omitempty"`
}

type ArrayOfCPU struct {
	Item []*CPU `xml:"item,omitempty"`
}

type CPUHistory struct {
	Identifier int32 `xml:"Identifier,omitempty"`

	Hour int32 `xml:"Hour,omitempty"`

	CpuUser int32 `xml:"CpuUser,omitempty"`

	CpuSystem int32 `xml:"CpuSystem,omitempty"`

	CpuIdle int32 `xml:"CpuIdle,omitempty"`

	Interrupts int32 `xml:"Interrupts,omitempty"`

	SystemCalls int32 `xml:"SystemCalls,omitempty"`

	ContextSwitches int32 `xml:"ContextSwitches,omitempty"`

	MaxLoadAverage15 int32 `xml:"MaxLoadAverage15,omitempty"`

	MinLoadAverage15 int32 `xml:"MinLoadAverage15,omitempty"`

	IdleTrue int32 `xml:"IdleTrue,omitempty"`

	WaitTrue int32 `xml:"WaitTrue,omitempty"`
}

type ArrayOfCPUHistory struct {
	Item []*CPUHistory `xml:"item,omitempty"`
}

type Memory struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	NumberOfPagesIn int32 `xml:"NumberOfPagesIn,omitempty"`

	NumberOfPagesOut int32 `xml:"NumberOfPagesOut,omitempty"`

	DataPagedIn int32 `xml:"DataPagedIn,omitempty"`

	DataPagedOut int32 `xml:"DataPagedOut,omitempty"`

	FreeMemory int32 `xml:"FreeMemory,omitempty"`

	PhysMemory int32 `xml:"PhysMemory,omitempty"`

	ConfiguredSwap int32 `xml:"ConfiguredSwap,omitempty"`

	FreeSwap int32 `xml:"FreeSwap,omitempty"`

	SwapSize int32 `xml:"SwapSize,omitempty"`

	MaxSwapSize int32 `xml:"MaxSwapSize,omitempty"`

	FreeMemInclFsCache int32 `xml:"FreeMemInclFsCache,omitempty"`
}

type MemoryHistory struct {
	Identifier int32 `xml:"Identifier,omitempty"`

	Hour int32 `xml:"Hour,omitempty"`

	NumberOfPagesIn int32 `xml:"NumberOfPagesIn,omitempty"`

	NumberOfPagesOut int32 `xml:"NumberOfPagesOut,omitempty"`

	DataPagedIn int32 `xml:"DataPagedIn,omitempty"`

	DataPagedOut int32 `xml:"DataPagedOut,omitempty"`

	MinFreeMemory int32 `xml:"MinFreeMemory,omitempty"`

	MaxFreeMemory int32 `xml:"MaxFreeMemory,omitempty"`

	AvgFreeMemory int32 `xml:"AvgFreeMemory,omitempty"`

	MinSwapSize int32 `xml:"MinSwapSize,omitempty"`

	MaxSwapSize int32 `xml:"MaxSwapSize,omitempty"`

	AvgSwapSize int32 `xml:"AvgSwapSize,omitempty"`

	MinSwapFree int32 `xml:"MinSwapFree,omitempty"`

	MaxSwapFree int32 `xml:"MaxSwapFree,omitempty"`

	AvgSwapFree int32 `xml:"AvgSwapFree,omitempty"`
}

type ArrayOfMemoryHistory struct {
	Item []*MemoryHistory `xml:"item,omitempty"`
}

type Disk struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	Diskname string `xml:"Diskname,omitempty"`

	Utilization int32 `xml:"Utilization,omitempty"`

	AverageQueueLength int32 `xml:"AverageQueueLength,omitempty"`

	AverageWaitTime int32 `xml:"AverageWaitTime,omitempty"`

	AverageServiceTime int32 `xml:"AverageServiceTime,omitempty"`

	KBperSec int32 `xml:"KBperSec,omitempty"`

	OperationsPerSec int32 `xml:"OperationsPerSec,omitempty"`
}

type ArrayOfDisk struct {
	Item []*Disk `xml:"item,omitempty"`
}

type DiskHistory struct {
	Identifier int32 `xml:"Identifier,omitempty"`

	Hour int32 `xml:"Hour,omitempty"`

	Diskname string `xml:"Diskname,omitempty"`

	AverageUtilization int32 `xml:"AverageUtilization,omitempty"`

	AverageQueueLength int32 `xml:"AverageQueueLength,omitempty"`

	AverageWaitTime int32 `xml:"AverageWaitTime,omitempty"`

	AverageServiceTime int32 `xml:"AverageServiceTime,omitempty"`

	MBperHour int32 `xml:"MBperHour,omitempty"`

	OperationsPerHour int32 `xml:"OperationsPerHour,omitempty"`
}

type ArrayOfDiskHistory struct {
	Item []*DiskHistory `xml:"item,omitempty"`
}

type Filesystem struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	FsysName string `xml:"FsysName,omitempty"`

	Capacity int32 `xml:"Capacity,omitempty"`

	FreeSpace int32 `xml:"FreeSpace,omitempty"`

	Status int32 `xml:"Status,omitempty"`
}

type ArrayOfFilesystem struct {
	Item []*Filesystem `xml:"item,omitempty"`
}

type FilesystemHistory struct {
	Identifier int32 `xml:"Identifier,omitempty"`

	Hour int32 `xml:"Hour,omitempty"`

	FsysName string `xml:"FsysName,omitempty"`

	Capacity int32 `xml:"Capacity,omitempty"`

	FreeSpace int32 `xml:"FreeSpace,omitempty"`

	Status int32 `xml:"Status,omitempty"`
}

type ArrayOfFilesystemHistory struct {
	Item []*FilesystemHistory `xml:"item,omitempty"`
}

type OsData struct {
	Objectname string `xml:"objectname,omitempty"`

	Attributename string `xml:"attributename,omitempty"`

	Mteclass string `xml:"mteclass,omitempty"`

	Attributegroup string `xml:"attributegroup,omitempty"`

	F1class string `xml:"f1-class,omitempty"`

	F1no int32 `xml:"f1-no,omitempty"`

	Visibility int32 `xml:"visibility,omitempty"`

	Mtetype int32 `xml:"mtetype,omitempty"`

	Subtype int32 `xml:"subtype,omitempty"`

	Perfunit string `xml:"perf-unit,omitempty"`

	Perfdecimals int32 `xml:"perf-decimals,omitempty"`

	Perfalclass string `xml:"perf-al-class,omitempty"`

	Perfalno int32 `xml:"perf-al-no,omitempty"`

	Perfthreshg2y int32 `xml:"perf-thresh-g2y,omitempty"`

	Perfthreshy2r int32 `xml:"perf-thresh-y2r,omitempty"`

	Perfthreshr2y int32 `xml:"perf-thresh-r2y,omitempty"`

	Perfthreshy2g int32 `xml:"perf-thresh-y2g,omitempty"`

	Perfsnaptotal int32 `xml:"perf-snap-total,omitempty"`

	Perfsnapnumber int32 `xml:"perf-snap-number,omitempty"`

	Perfsnaptimehigh int32 `xml:"perf-snap-time-high,omitempty"`

	Perfsnaptimelow int32 `xml:"perf-snap-time-low,omitempty"`

	Perflmintotal int32 `xml:"perf-lmin-total,omitempty"`

	Perflminnumber int32 `xml:"perf-lmin-number,omitempty"`

	Perfl5mintotal int32 `xml:"perf-l5min-total,omitempty"`

	Perfl5minnumber int32 `xml:"perf-l5min-number,omitempty"`

	Perfl15mintotal int32 `xml:"perf-l15min-total,omitempty"`

	Perfl15minnumber int32 `xml:"perf-l15min-number,omitempty"`

	Perflhourtotal int32 `xml:"perf-lhour-total,omitempty"`

	Perflhournumber int32 `xml:"perf-lhour-number,omitempty"`

	Textclass string `xml:"text-class,omitempty"`

	Textno int32 `xml:"text-no,omitempty"`

	Textvalue string `xml:"text-value,omitempty"`

	Color int32 `xml:"color,omitempty"`

	Mscmode int32 `xml:"msc-mode,omitempty"`

	Msctimespanminutes int32 `xml:"msc-time-span-minutes,omitempty"`

	Staalertmode int32 `xml:"sta-alert-mode,omitempty"`

	Staalertshift int32 `xml:"sta-alert-shift,omitempty"`

	Parameter1 string `xml:"parameter1,omitempty"`

	Parameter2 string `xml:"parameter2,omitempty"`

	Parameter3 string `xml:"parameter3,omitempty"`

	Parameter4 string `xml:"parameter4,omitempty"`
}

type ArrayOfOsData struct {
	Item []*OsData `xml:"item,omitempty"`
}

type LAN struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	LanName string `xml:"LanName,omitempty"`

	PacketsIn int32 `xml:"PacketsIn,omitempty"`

	PacketsOut int32 `xml:"PacketsOut,omitempty"`

	ErrorsIn int32 `xml:"ErrorsIn,omitempty"`

	ErrorsOut int32 `xml:"ErrorsOut,omitempty"`

	Collisions int32 `xml:"Collisions,omitempty"`
}

type ArrayOfLAN struct {
	Item []*LAN `xml:"item,omitempty"`
}

type LANHistory struct {
	Identifier int32 `xml:"Identifier,omitempty"`

	Hour int32 `xml:"Hour,omitempty"`

	LanName string `xml:"LanName,omitempty"`

	PacketsIn int32 `xml:"PacketsIn,omitempty"`

	PacketsOut int32 `xml:"PacketsOut,omitempty"`

	ErrorsIn int32 `xml:"ErrorsIn,omitempty"`

	ErrorsOut int32 `xml:"ErrorsOut,omitempty"`

	Collisions int32 `xml:"Collisions,omitempty"`
}

type ArrayOfLANHistory struct {
	Item []*LANHistory `xml:"item,omitempty"`
}

type Process struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	ProcessID int32 `xml:"ProcessID,omitempty"`

	Username string `xml:"Username,omitempty"`

	Command string `xml:"Command,omitempty"`

	CpuUtilization int32 `xml:"CpuUtilization,omitempty"`

	CpuTime string `xml:"CpuTime,omitempty"`

	ResidentSize int32 `xml:"ResidentSize,omitempty"`

	Priority int32 `xml:"Priority,omitempty"`

	PrivPages int32 `xml:"PrivPages,omitempty"`
}

type ArrayOfProcess struct {
	Item []*Process `xml:"item,omitempty"`
}

type ProcessPattern struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	ProcessNamePattern string `xml:"ProcessNamePattern,omitempty"`

	UserNamePattern string `xml:"UserNamePattern,omitempty"`

	NumberOfProc int32 `xml:"NumberOfProc,omitempty"`

	CpuUtil int32 `xml:"CpuUtil,omitempty"`

	ResSize int64 `xml:"ResSize,omitempty"`

	PrivPages int64 `xml:"PrivPages,omitempty"`
}

type ArrayOfProcessPattern struct {
	Item []*ProcessPattern `xml:"item,omitempty"`
}

type OSParameter struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	ParameterName string `xml:"ParameterName,omitempty"`

	CurrentValue string `xml:"CurrentValue,omitempty"`

	DefinedValue string `xml:"DefinedValue,omitempty"`
}

type ArrayOfOSParameter struct {
	Item []*OSParameter `xml:"item,omitempty"`
}

type OS390AddressSpaceResourceData struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	Jobname string `xml:"jobname,omitempty"`

	TransactionResidenyTime uint32 `xml:"TransactionResidenyTime,omitempty"`

	TransactionElapsedTime uint32 `xml:"TransactionElapsedTime,omitempty"`

	StepTCBTime uint32 `xml:"StepTCBTime,omitempty"`

	DeviceConnectTime uint32 `xml:"DeviceConnectTime,omitempty"`

	ProcessorTime uint32 `xml:"ProcessorTime,omitempty"`

	PageInCount uint32 `xml:"PageInCount,omitempty"`

	PagesSwapped uint32 `xml:"PagesSwapped,omitempty"`

	RealFrames uint32 `xml:"RealFrames,omitempty"`

	FixedFramesBelow uint32 `xml:"FixedFramesBelow,omitempty"`

	PrivateFramesBelow uint32 `xml:"PrivateFramesBelow,omitempty"`

	VIOPages uint32 `xml:"VIOPages,omitempty"`

	NonVIOPages uint32 `xml:"NonVIOPages,omitempty"`

	LSQAPages uint32 `xml:"LSQAPages,omitempty"`

	LSQAPagesExpanded uint32 `xml:"LSQAPagesExpanded,omitempty"`

	LSQAFixedFrames uint32 `xml:"LSQAFixedFrames,omitempty"`

	NonLSQAFixedFrames uint32 `xml:"NonLSQAFixedFrames,omitempty"`

	EXCPCount uint32 `xml:"EXCPCount,omitempty"`

	Reserved1 uint32 `xml:"Reserved1,omitempty"`

	Reserved2 uint32 `xml:"Reserved2,omitempty"`

	Reserved3 uint32 `xml:"Reserved3,omitempty"`
}

type ArrayOfOS390AddressSpaceResourceData struct {
	Item []*OS390AddressSpaceResourceData `xml:"item,omitempty"`
}

type OS390CPUPagingActivity struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	MeasurementIntervalSeconds int32 `xml:"MeasurementIntervalSeconds,omitempty"`

	AvgProcessorUtilization int32 `xml:"AvgProcessorUtilization,omitempty"`

	AvgSystemPagingRate int32 `xml:"AvgSystemPagingRate,omitempty"`

	AvgMVSCpuUtilization int32 `xml:"AvgMVSCpuUtilization,omitempty"`

	PagesInPerSecond int32 `xml:"PagesInPerSecond,omitempty"`

	PagesOutPerSecond int32 `xml:"PagesOutPerSecond,omitempty"`

	PrivatePagesInPerSecond int32 `xml:"PrivatePagesInPerSecond,omitempty"`

	PrivatePagesOutPerSecond int32 `xml:"PrivatePagesOutPerSecond,omitempty"`

	PagesToExpandedStorage int32 `xml:"PagesToExpandedStorage,omitempty"`

	PagesMigratedFromExpanded int32 `xml:"PagesMigratedFromExpanded,omitempty"`

	BlockedPagesPagedIn int32 `xml:"BlockedPagesPagedIn,omitempty"`

	BlocksPagedIn int32 `xml:"BlocksPagedIn,omitempty"`

	SmfId string `xml:"SmfId,omitempty"`

	SystemName string `xml:"SystemName,omitempty"`

	Reserved int32 `xml:"Reserved,omitempty"`
}

type ArrayOfOS390CPUPagingActivity struct {
	Item []*OS390CPUPagingActivity `xml:"item,omitempty"`
}

type OS390CPUPagingActivityHistory struct {
	Identifier int32 `xml:"Identifier,omitempty"`

	Hour int32 `xml:"Hour,omitempty"`

	MeasurementIntervalSeconds int32 `xml:"MeasurementIntervalSeconds,omitempty"`

	NumberOfSnapShots int32 `xml:"NumberOfSnapShots,omitempty"`

	MinProcessorUtilization int32 `xml:"MinProcessorUtilization,omitempty"`

	MaxProcessorUtilization int32 `xml:"MaxProcessorUtilization,omitempty"`

	AvgProcessorUtilization int32 `xml:"AvgProcessorUtilization,omitempty"`

	MinSystemPagingRate int32 `xml:"MinSystemPagingRate,omitempty"`

	MaxSystemPagingRate int32 `xml:"MaxSystemPagingRate,omitempty"`

	AvgSystemPagingRate int32 `xml:"AvgSystemPagingRate,omitempty"`

	MinMVSCpuUtilization int32 `xml:"MinMVSCpuUtilization,omitempty"`

	MaxMVSCpuUtilization int32 `xml:"MaxMVSCpuUtilization,omitempty"`

	AvgMVSCpuUtilization int32 `xml:"AvgMVSCpuUtilization,omitempty"`

	PagesInPerSecond int32 `xml:"PagesInPerSecond,omitempty"`

	PagesOutPerSecond int32 `xml:"PagesOutPerSecond,omitempty"`

	PrivatePagesInPerSecond int32 `xml:"PrivatePagesInPerSecond,omitempty"`

	PrivatePagesOutPerSecond int32 `xml:"PrivatePagesOutPerSecond,omitempty"`

	PagesToExpandedStorage int32 `xml:"PagesToExpandedStorage,omitempty"`

	PagesMigratedFromExpanded int32 `xml:"PagesMigratedFromExpanded,omitempty"`

	BlockedPagesPagedIn int32 `xml:"BlockedPagesPagedIn,omitempty"`

	BlocksPagedIn int32 `xml:"BlocksPagedIn,omitempty"`

	SmfId string `xml:"SmfId,omitempty"`

	SystemName string `xml:"SystemName,omitempty"`

	Reserved int32 `xml:"Reserved,omitempty"`
}

type ArrayOfOS390CPUPagingActivityHistory struct {
	Item []*OS390CPUPagingActivityHistory `xml:"item,omitempty"`
}

type OS390StorageResource struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	MeasurementIntervalSeconds int32 `xml:"MeasurementIntervalSeconds,omitempty"`

	HighUICCount int32 `xml:"HighUICCount,omitempty"`

	ExpandedStorageMigrationAge int32 `xml:"ExpandedStorageMigrationAge,omitempty"`

	AvailableCentralStorageFrames int32 `xml:"AvailableCentralStorageFrames,omitempty"`

	AvailableExpandedStorageFrames int32 `xml:"AvailableExpandedStorageFrames,omitempty"`

	SmfId string `xml:"SmfId,omitempty"`

	SystemName string `xml:"SystemName,omitempty"`

	Reserved1 int32 `xml:"Reserved1,omitempty"`

	Reserved2 int32 `xml:"Reserved2,omitempty"`
}

type ArrayOfOS390StorageResource struct {
	Item []*OS390StorageResource `xml:"item,omitempty"`
}

type OS390StorageResourceHistory struct {
	Identifier int32 `xml:"Identifier,omitempty"`

	Hour int32 `xml:"Hour,omitempty"`

	MeasurementIntervalSeconds int32 `xml:"MeasurementIntervalSeconds,omitempty"`

	NumberOfSnapShots int32 `xml:"NumberOfSnapShots,omitempty"`

	MinUICCount int32 `xml:"MinUICCount,omitempty"`

	MaxUICCount int32 `xml:"maxUICCount,omitempty"`

	AvgUICCount int32 `xml:"AvgUICCount,omitempty"`

	MinExpandedStorageMigrationAge int32 `xml:"MinExpandedStorageMigrationAge,omitempty"`

	MaxExpandedStorageMigrationAge int32 `xml:"MaxExpandedStorageMigrationAge,omitempty"`

	AvgExpandedStorageMigrationAge int32 `xml:"AvgExpandedStorageMigrationAge,omitempty"`

	MinAvailableCentralStorageFrames int32 `xml:"MinAvailableCentralStorageFrames,omitempty"`

	MaxAvailableCentralStorageFrames int32 `xml:"MaxAvailableCentralStorageFrames,omitempty"`

	AvgAvailableCentralStorageFrames int32 `xml:"AvgAvailableCentralStorageFrames,omitempty"`

	MinAvailableExpandedStorageFrames int32 `xml:"MinAvailableExpandedStorageFrames,omitempty"`

	MaxAvailableExpandedStorageFrames int32 `xml:"MaxAvailableExpandedStorageFrames,omitempty"`

	AvgAvailableExpandedStorageFrames int32 `xml:"AvgAvailableExpandedStorageFrames,omitempty"`

	SmfId string `xml:"SmfId,omitempty"`

	SystemName string `xml:"SystemName,omitempty"`

	Reserved1 int32 `xml:"Reserved1,omitempty"`

	Reserved2 int32 `xml:"Reserved2,omitempty"`

	Reserved3 int32 `xml:"Reserved3,omitempty"`

	Reserved4 int32 `xml:"Reserved4,omitempty"`
}

type ArrayOfOS390StorageResourceHistory struct {
	Item []*OS390StorageResourceHistory `xml:"item,omitempty"`
}

type OS390CollectorDetails struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	Text string `xml:"Text,omitempty"`
}

type ArrayOfOS390CollectorDetails struct {
	Item []*OS390CollectorDetails `xml:"item,omitempty"`
}

type OS390WLMPolicy struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	SystemName string `xml:"SystemName,omitempty"`

	PolicyName string `xml:"PolicyName,omitempty"`

	PolicyDescription string `xml:"PolicyDescription,omitempty"`

	AdminName string `xml:"AdminName,omitempty"`

	SystemMode uint16 `xml:"SystemMode,omitempty"`

	StartTime string `xml:"StartTime,omitempty"`

	EndTime string `xml:"EndTime,omitempty"`

	Duration uint64 `xml:"Duration,omitempty"`

	MinTime uint64 `xml:"MinTime,omitempty"`

	Status int32 `xml:"Status,omitempty"`
}

type ArrayOfOS390WLMPolicy struct {
	Item []*OS390WLMPolicy `xml:"item,omitempty"`
}

type OS390WLMMap struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	BucketMap00 int32 `xml:"BucketMap00,omitempty"`

	BucketMap01 int32 `xml:"BucketMap01,omitempty"`

	BucketMap02 int32 `xml:"BucketMap02,omitempty"`

	BucketMap03 int32 `xml:"BucketMap03,omitempty"`

	BucketMap04 int32 `xml:"BucketMap04,omitempty"`

	BucketMap05 int32 `xml:"BucketMap05,omitempty"`

	BucketMap06 int32 `xml:"BucketMap06,omitempty"`

	BucketMap07 int32 `xml:"BucketMap07,omitempty"`

	BucketMap08 int32 `xml:"BucketMap08,omitempty"`

	BucketMap09 int32 `xml:"BucketMap09,omitempty"`

	BucketMap10 int32 `xml:"BucketMap10,omitempty"`

	BucketMap11 int32 `xml:"BucketMap11,omitempty"`

	BucketMap12 int32 `xml:"BucketMap12,omitempty"`

	BucketMap13 int32 `xml:"BucketMap13,omitempty"`
}

type ArrayOfOS390WLMMap struct {
	Item []*OS390WLMMap `xml:"item,omitempty"`
}

type OS390WLMEnclaves struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	Name string `xml:"Name,omitempty"`

	SCName string `xml:"SCName,omitempty"`

	TPName string `xml:"TPName,omitempty"`

	UserID string `xml:"UserID,omitempty"`

	SSType string `xml:"SSType,omitempty"`

	FuncName string `xml:"FuncName,omitempty"`

	SCIndex int32 `xml:"SCIndex,omitempty"`

	PeriodInd int32 `xml:"PeriodInd,omitempty"`

	Value1 uint32 `xml:"value1,omitempty"`

	Value2 uint32 `xml:"value2,omitempty"`
}

type ArrayOfOS390WLMEnclaves struct {
	Item []*OS390WLMEnclaves `xml:"item,omitempty"`
}

type OS390WLMServiceClass struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	SCName string `xml:"SCName,omitempty"`

	SCDescription string `xml:"SCDescription,omitempty"`

	PeriodIndex int32 `xml:"PeriodIndex,omitempty"`

	PeriodNumber int32 `xml:"PeriodNumber,omitempty"`

	Workload string `xml:"Workload,omitempty"`

	ResourceGroup string `xml:"ResourceGroup,omitempty"`

	MinCapacity uint64 `xml:"MinCapacity,omitempty"`

	MaxCapacity uint64 `xml:"MaxCapacity,omitempty"`

	Reserved1 uint64 `xml:"Reserved1,omitempty"`

	Reserved2 uint64 `xml:"Reserved2,omitempty"`
}

type ArrayOfOS390WLMServiceClass struct {
	Item []*OS390WLMServiceClass `xml:"item,omitempty"`
}

type OS390WLMServiceClassPeriodData struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	GoalType int32 `xml:"GoalType,omitempty"`

	GoalPercValue int32 `xml:"GoalPercValue,omitempty"`

	Importance int32 `xml:"Importance,omitempty"`

	GoalValue uint32 `xml:"GoalValue,omitempty"`

	Duration uint32 `xml:"Duration,omitempty"`

	CPUUnits1 uint32 `xml:"CPUUnits1,omitempty"`

	CPUUnits2 uint32 `xml:"CPUUnits2,omitempty"`

	SRBUnits1 uint32 `xml:"SRBUnits1,omitempty"`

	SRBUnits2 uint32 `xml:"SRBUnits2,omitempty"`

	CompletedTAs uint32 `xml:"CompletedTAs,omitempty"`

	AbortedTAs uint32 `xml:"AbortedTAs,omitempty"`

	UsingSamples uint32 `xml:"UsingSamples,omitempty"`

	DelaySamples uint32 `xml:"DelaySamples,omitempty"`

	ElapsedTime1 uint32 `xml:"ElapsedTime1,omitempty"`

	ElapsedTime2 uint32 `xml:"ElapsedTime2,omitempty"`

	ExecutionTime1 uint32 `xml:"ExecutionTime1,omitempty"`

	ExecutionTime2 uint32 `xml:"ExecutionTime2,omitempty"`

	RTDIndex int32 `xml:"RTDIndex,omitempty"`
}

type ArrayOfOS390WLMServiceClassPeriodData struct {
	Item []*OS390WLMServiceClassPeriodData `xml:"item,omitempty"`
}

type OS390WLMResponseTimeDistribution struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	Bucket00 int32 `xml:"Bucket00,omitempty"`

	Bucket01 int32 `xml:"Bucket01,omitempty"`

	Bucket02 int32 `xml:"Bucket02,omitempty"`

	Bucket03 int32 `xml:"Bucket03,omitempty"`

	Bucket04 int32 `xml:"Bucket04,omitempty"`

	Bucket05 int32 `xml:"Bucket05,omitempty"`

	Bucket06 int32 `xml:"Bucket06,omitempty"`

	Bucket07 int32 `xml:"Bucket07,omitempty"`

	Bucket08 int32 `xml:"Bucket08,omitempty"`

	Bucket09 int32 `xml:"Bucket09,omitempty"`

	Bucket10 int32 `xml:"Bucket10,omitempty"`

	Bucket11 int32 `xml:"Bucket11,omitempty"`

	Bucket12 int32 `xml:"Bucket12,omitempty"`

	Bucket13 int32 `xml:"Bucket13,omitempty"`
}

type ArrayOfOS390WLMResponseTimeDistribution struct {
	Item []*OS390WLMResponseTimeDistribution `xml:"item,omitempty"`
}

type AS400PoolAll struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	ActToWait int32 `xml:"ActToWait,omitempty"`

	ActToInel int32 `xml:"ActToInel,omitempty"`

	WaitToInel int32 `xml:"WaitToInel,omitempty"`
}

type AS400PoolData struct {
	CollectionTime time.Time `xml:"CollectionTime,omitempty"`

	Identifier int32 `xml:"Identifier,omitempty"`

	PoolNbr int32 `xml:"PoolNbr,omitempty"`

	MaxActive int32 `xml:"MaxActive,omitempty"`

	PoolSize int32 `xml:"PoolSize,omitempty"`

	MachineRes int32 `xml:"MachineRes,omitempty"`

	DBFaults int32 `xml:"DBFaults,omitempty"`

	NonDBFaults int32 `xml:"NonDBFaults,omitempty"`

	DBPages int32 `xml:"DBPages,omitempty"`

	NonDBPages int32 `xml:"NonDBPages,omitempty"`

	ActToWait int32 `xml:"ActToWait,omitempty"`

	ActToInel int32 `xml:"ActToInel,omitempty"`

	WaitToInel int32 `xml:"WaitToInel,omitempty"`

	R3PoolFlag int32 `xml:"R3PoolFlag,omitempty"`
}

type ArrayOfAS400PoolData struct {
	Item []*AS400PoolData `xml:"item,omitempty"`
}

type AS400PoolHistory struct {
	Identifier int32 `xml:"Identifier,omitempty"`

	Hour int32 `xml:"Hour,omitempty"`

	PoolNbr int32 `xml:"PoolNbr,omitempty"`

	MaxActive int32 `xml:"MaxActive,omitempty"`

	PoolSize int32 `xml:"PoolSize,omitempty"`

	MachineRes int32 `xml:"MachineRes,omitempty"`

	DBFaults int32 `xml:"DBFaults,omitempty"`

	NonDBFaults int32 `xml:"NonDBFaults,omitempty"`

	DBPages int32 `xml:"DBPages,omitempty"`

	NonDBPages int32 `xml:"NonDBPages,omitempty"`

	ActToWait int32 `xml:"ActToWait,omitempty"`

	ActToInel int32 `xml:"ActToInel,omitempty"`

	WaitToInel int32 `xml:"WaitToInel,omitempty"`

	R3PoolFlag int32 `xml:"R3PoolFlag,omitempty"`
}

type ArrayOfAS400PoolHistory struct {
	Item []*AS400PoolHistory `xml:"item,omitempty"`
}

type File struct {
	Name string `xml:"name,omitempty"`

	Content string `xml:"content,omitempty"`
}

type SAPOscol interface {

	/* Service definition of function SAPOscol__GetCpuConsumption */
	GetCpuConsumption(request *GetCpuConsumption) (*GetCpuConsumptionResponse, error)

	/* Service definition of function SAPOscol__GetSingleCpu */
	GetSingleCpu(request *GetSingleCpu) (*GetSingleCpuResponse, error)

	/* Service definition of function SAPOscol__GetCpuHistory */
	GetCpuHistory(request *GetCpuHistory) (*GetCpuHistoryResponse, error)

	/* Service definition of function SAPOscol__GetMemory */
	GetMemory(request *GetMemory) (*GetMemoryResponse, error)

	/* Service definition of function SAPOscol__GetMemoryHistory */
	GetMemoryHistory(request *GetMemoryHistory) (*GetMemoryHistoryResponse, error)

	/* Service definition of function SAPOscol__GetDisk */
	GetDisk(request *GetDisk) (*GetDiskResponse, error)

	/* Service definition of function SAPOscol__GetDiskHistory */
	GetDiskHistory(request *GetDiskHistory) (*GetDiskHistoryResponse, error)

	/* Service definition of function SAPOscol__GetFilesystem */
	GetFilesystem(request *GetFilesystem) (*GetFilesystemResponse, error)

	/* Service definition of function SAPOscol__GetFilesystemHistory */
	GetFilesystemHistory(request *GetFilesystemHistory) (*GetFilesystemHistoryResponse, error)

	/* Service definition of function SAPOscol__GetLAN */
	GetLAN(request *GetLAN) (*GetLANResponse, error)

	/* Service definition of function SAPOscol__GetLANHistory */
	GetLANHistory(request *GetLANHistory) (*GetLANHistoryResponse, error)

	/* Service definition of function SAPOscol__GetTopCPUProcesses */
	GetTopCPUProcesses(request *GetTopCPUProcesses) (*GetTopCPUProcessesResponse, error)

	/* Service definition of function SAPOscol__GetProcPattern */
	GetProcPattern(request *GetProcPattern) (*GetProcPatternResponse, error)

	/* Service definition of function SAPOscol__GetConfiguration */
	GetConfiguration(request *GetConfiguration) (*GetConfigurationResponse, error)

	/* Service definition of function SAPOscol__GetOsData */
	GetOsData(request *GetOsData) (*GetOsDataResponse, error)

	/* Service definition of function SAPOscol__GetOS390AddressSpaceResourceData */
	GetOS390AddressSpaceResourceData(request *GetOS390AddressSpaceResourceData) (*GetOS390AddressSpaceResourceDataResponse, error)

	/* Service definition of function SAPOscol__GetOS390CPUPagingActivity */
	GetOS390CPUPagingActivity(request *GetOS390CPUPagingActivity) (*GetOS390CPUPagingActivityResponse, error)

	/* Service definition of function SAPOscol__GetOS390CPUPagingActivityHistory */
	GetOS390CPUPagingActivityHistory(request *GetOS390CPUPagingActivityHistory) (*GetOS390CPUPagingActivityHistoryResponse, error)

	/* Service definition of function SAPOscol__GetOS390StorageResource */
	GetOS390StorageResource(request *GetOS390StorageResource) (*GetOS390StorageResourceResponse, error)

	/* Service definition of function SAPOscol__GetOS390StorageResourceHistory */
	GetOS390StorageResourceHistory(request *GetOS390StorageResourceHistory) (*GetOS390StorageResourceHistoryResponse, error)

	/* Service definition of function SAPOscol__GetOS390CollectorDetails */
	GetOS390CollectorDetails(request *GetOS390CollectorDetails) (*GetOS390CollectorDetailsResponse, error)

	/* Service definition of function SAPOscol__GetOs390WLMData */
	GetOs390WLMData(request *GetOs390WLMData) (*GetOs390WLMDataResponse, error)

	/* Service definition of function SAPOscol__GetAS400PoolAll */
	GetAS400PoolAll(request *GetAS400PoolAll) (*GetAS400PoolAllResponse, error)

	/* Service definition of function SAPOscol__GetAS400PoolData */
	GetAS400PoolData(request *GetAS400PoolData) (*GetAS400PoolDataResponse, error)

	/* Service definition of function SAPOscol__GetAS400PoolHistory */
	GetAS400PoolHistory(request *GetAS400PoolHistory) (*GetAS400PoolHistoryResponse, error)

	/* Service definition of function SAPOscol__SetAdminParameter */
	SetAdminParameter(request *SetAdminParameter) (*SetAdminParameterResponse, error)

	/* Service definition of function SAPOscol__SendRequest */
	SendRequest(request *SendRequest) (*SendRequestResponse, error)

	/* Service definition of function SAPOscol__SendRequestAsync */
	SendRequestAsync(request *SendRequestAsync) (*SendRequestAsyncResponse, error)

	/* Service definition of function SAPOscol__GetResult */
	GetResult(request *GetResult) (*GetResultResponse, error)

	/* Service definition of function SAPOscol__GetVersion */
	GetVersion(request *GetVersion) (*GetVersionResponse, error)

	/* Service definition of function SAPOscol__GetHwConfText */
	GetHwConfText(request *GetHwConfText) (*GetHwConfTextResponse, error)

	/* Service definition of function SAPOscol__GetHwConfXML */
	GetHwConfXML(request *GetHwConfXML) (*GetHwConfXMLResponse, error)

	/* Returns a logon file for local authentication. */
	RequestLogonFile(request *RequestLogonFile) (*RequestLogonFileResponse, error)
}

type sAPOscol struct {
	client *soap.Client
}

func NewSAPOscol(client *soap.Client) SAPOscol {
	return &sAPOscol{
		client: client,
	}
}

func (service *sAPOscol) GetCpuConsumption(request *GetCpuConsumption) (*GetCpuConsumptionResponse, error) {
	response := new(GetCpuConsumptionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetSingleCpu(request *GetSingleCpu) (*GetSingleCpuResponse, error) {
	response := new(GetSingleCpuResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetCpuHistory(request *GetCpuHistory) (*GetCpuHistoryResponse, error) {
	response := new(GetCpuHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetMemory(request *GetMemory) (*GetMemoryResponse, error) {
	response := new(GetMemoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetMemoryHistory(request *GetMemoryHistory) (*GetMemoryHistoryResponse, error) {
	response := new(GetMemoryHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetDisk(request *GetDisk) (*GetDiskResponse, error) {
	response := new(GetDiskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetDiskHistory(request *GetDiskHistory) (*GetDiskHistoryResponse, error) {
	response := new(GetDiskHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetFilesystem(request *GetFilesystem) (*GetFilesystemResponse, error) {
	response := new(GetFilesystemResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetFilesystemHistory(request *GetFilesystemHistory) (*GetFilesystemHistoryResponse, error) {
	response := new(GetFilesystemHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetLAN(request *GetLAN) (*GetLANResponse, error) {
	response := new(GetLANResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetLANHistory(request *GetLANHistory) (*GetLANHistoryResponse, error) {
	response := new(GetLANHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetTopCPUProcesses(request *GetTopCPUProcesses) (*GetTopCPUProcessesResponse, error) {
	response := new(GetTopCPUProcessesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetProcPattern(request *GetProcPattern) (*GetProcPatternResponse, error) {
	response := new(GetProcPatternResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetConfiguration(request *GetConfiguration) (*GetConfigurationResponse, error) {
	response := new(GetConfigurationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetOsData(request *GetOsData) (*GetOsDataResponse, error) {
	response := new(GetOsDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetOS390AddressSpaceResourceData(request *GetOS390AddressSpaceResourceData) (*GetOS390AddressSpaceResourceDataResponse, error) {
	response := new(GetOS390AddressSpaceResourceDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetOS390CPUPagingActivity(request *GetOS390CPUPagingActivity) (*GetOS390CPUPagingActivityResponse, error) {
	response := new(GetOS390CPUPagingActivityResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetOS390CPUPagingActivityHistory(request *GetOS390CPUPagingActivityHistory) (*GetOS390CPUPagingActivityHistoryResponse, error) {
	response := new(GetOS390CPUPagingActivityHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetOS390StorageResource(request *GetOS390StorageResource) (*GetOS390StorageResourceResponse, error) {
	response := new(GetOS390StorageResourceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetOS390StorageResourceHistory(request *GetOS390StorageResourceHistory) (*GetOS390StorageResourceHistoryResponse, error) {
	response := new(GetOS390StorageResourceHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetOS390CollectorDetails(request *GetOS390CollectorDetails) (*GetOS390CollectorDetailsResponse, error) {
	response := new(GetOS390CollectorDetailsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetOs390WLMData(request *GetOs390WLMData) (*GetOs390WLMDataResponse, error) {
	response := new(GetOs390WLMDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetAS400PoolAll(request *GetAS400PoolAll) (*GetAS400PoolAllResponse, error) {
	response := new(GetAS400PoolAllResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetAS400PoolData(request *GetAS400PoolData) (*GetAS400PoolDataResponse, error) {
	response := new(GetAS400PoolDataResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetAS400PoolHistory(request *GetAS400PoolHistory) (*GetAS400PoolHistoryResponse, error) {
	response := new(GetAS400PoolHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) SetAdminParameter(request *SetAdminParameter) (*SetAdminParameterResponse, error) {
	response := new(SetAdminParameterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) SendRequest(request *SendRequest) (*SendRequestResponse, error) {
	response := new(SendRequestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) SendRequestAsync(request *SendRequestAsync) (*SendRequestAsyncResponse, error) {
	response := new(SendRequestAsyncResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetResult(request *GetResult) (*GetResultResponse, error) {
	response := new(GetResultResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetVersion(request *GetVersion) (*GetVersionResponse, error) {
	response := new(GetVersionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetHwConfText(request *GetHwConfText) (*GetHwConfTextResponse, error) {
	response := new(GetHwConfTextResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) GetHwConfXML(request *GetHwConfXML) (*GetHwConfXMLResponse, error) {
	response := new(GetHwConfXMLResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPOscol) RequestLogonFile(request *RequestLogonFile) (*RequestLogonFileResponse, error) {
	response := new(RequestLogonFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
