package SAPControl

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type STATECOLOR string

const (
	STATECOLORSAPControlGRAY STATECOLOR = "SAPControl-GRAY"

	STATECOLORSAPControlGREEN STATECOLOR = "SAPControl-GREEN"

	STATECOLORSAPControlYELLOW STATECOLOR = "SAPControl-YELLOW"

	STATECOLORSAPControlRED STATECOLOR = "SAPControl-RED"
)

type VISIBLELEVEL string

const (
	VISIBLELEVELSAPControlUNKNOWN VISIBLELEVEL = "SAPControl-UNKNOWN"

	VISIBLELEVELSAPControlOPERATOR VISIBLELEVEL = "SAPControl-OPERATOR"

	VISIBLELEVELSAPControlEXPERT VISIBLELEVEL = "SAPControl-EXPERT"

	VISIBLELEVELSAPControlDEVELOPER VISIBLELEVEL = "SAPControl-DEVELOPER"
)

type J2EEPSTATE string

const (
	J2EEPSTATESAPControlJ2EESTOPPED J2EEPSTATE = "SAPControl-J2EE-STOPPED"

	J2EEPSTATESAPControlJ2EESTARTING J2EEPSTATE = "SAPControl-J2EE-STARTING"

	J2EEPSTATESAPControlJ2EECORERUNNING J2EEPSTATE = "SAPControl-J2EE-CORE-RUNNING"

	J2EEPSTATESAPControlJ2EERUNNING J2EEPSTATE = "SAPControl-J2EE-RUNNING"

	J2EEPSTATESAPControlJ2EESTOPPING J2EEPSTATE = "SAPControl-J2EE-STOPPING"

	J2EEPSTATESAPControlJ2EEMAINTENANCE J2EEPSTATE = "SAPControl-J2EE-MAINTENANCE"

	J2EEPSTATESAPControlJ2EEUNKNOWN J2EEPSTATE = "SAPControl-J2EE-UNKNOWN"
)

type StartStopOption string

const (
	StartStopOptionSAPControlALLINSTANCES StartStopOption = "SAPControl-ALL-INSTANCES"

	StartStopOptionSAPControlSCSINSTANCES StartStopOption = "SAPControl-SCS-INSTANCES"

	StartStopOptionSAPControlDIALOGINSTANCES StartStopOption = "SAPControl-DIALOG-INSTANCES"

	StartStopOptionSAPControlABAPINSTANCES StartStopOption = "SAPControl-ABAP-INSTANCES"

	StartStopOptionSAPControlJ2EEINSTANCES StartStopOption = "SAPControl-J2EE-INSTANCES"

	StartStopOptionSAPControlPRIORITYLEVEL StartStopOption = "SAPControl-PRIORITY-LEVEL"

	StartStopOptionSAPControlTREXINSTANCES StartStopOption = "SAPControl-TREX-INSTANCES"

	StartStopOptionSAPControlENQREPINSTANCES StartStopOption = "SAPControl-ENQREP-INSTANCES"

	StartStopOptionSAPControlHDBINSTANCES StartStopOption = "SAPControl-HDB-INSTANCES"

	StartStopOptionSAPControlALLNOHDBINSTANCES StartStopOption = "SAPControl-ALLNOHDB-INSTANCES"
)

type LogFileConfigOperation string

const (
	LogFileConfigOperationSAPControlSETLOGFILES LogFileConfigOperation = "SAPControl-SET-LOGFILES"

	LogFileConfigOperationSAPControlADDLOGFILES LogFileConfigOperation = "SAPControl-ADD-LOGFILES"

	LogFileConfigOperationSAPControlREMOVELOGFILES LogFileConfigOperation = "SAPControl-REMOVE-LOGFILES"
)

type RESTRICTIONTYPE string

const (
	RESTRICTIONTYPESAPControlRESTRICTNONE RESTRICTIONTYPE = "SAPControl-RESTRICT-NONE"

	RESTRICTIONTYPESAPControlRESTRICTINT RESTRICTIONTYPE = "SAPControl-RESTRICT-INT"

	RESTRICTIONTYPESAPControlRESTRICTFLOAT RESTRICTIONTYPE = "SAPControl-RESTRICT-FLOAT"

	RESTRICTIONTYPESAPControlRESTRICTINTRANGE RESTRICTIONTYPE = "SAPControl-RESTRICT-INTRANGE"

	RESTRICTIONTYPESAPControlRESTRICTFLOATRANGE RESTRICTIONTYPE = "SAPControl-RESTRICT-FLOATRANGE"

	RESTRICTIONTYPESAPControlRESTRICTENUM RESTRICTIONTYPE = "SAPControl-RESTRICT-ENUM"

	RESTRICTIONTYPESAPControlRESTRICTBOOL RESTRICTIONTYPE = "SAPControl-RESTRICT-BOOL"
)

type PERSISTENCETYPE string

const (
	PERSISTENCETYPESAPControlDYNAMIC PERSISTENCETYPE = "SAPControl-DYNAMIC"

	PERSISTENCETYPESAPControlPERSIST PERSISTENCETYPE = "SAPControl-PERSIST"

	PERSISTENCETYPESAPControlDYNAMICPERSIST PERSISTENCETYPE = "SAPControl-DYNAMIC-PERSIST"
)

type HAVerificationState string

const (
	HAVerificationStateSAPControlHASUCCESS HAVerificationState = "SAPControl-HA-SUCCESS"

	HAVerificationStateSAPControlHAWARNING HAVerificationState = "SAPControl-HA-WARNING"

	HAVerificationStateSAPControlHAERROR HAVerificationState = "SAPControl-HA-ERROR"
)

type HACheckCategory string

const (
	HACheckCategorySAPControlSAPCONFIGURATION HACheckCategory = "SAPControl-SAP-CONFIGURATION"

	HACheckCategorySAPControlSAPSTATE HACheckCategory = "SAPControl-SAP-STATE"

	HACheckCategorySAPControlHACONFIGURATION HACheckCategory = "SAPControl-HA-CONFIGURATION"

	HACheckCategorySAPControlHASTATE HACheckCategory = "SAPControl-HA-STATE"
)

type SnapshotZip []byte

type Start struct {
	XMLName xml.Name `xml:"urn:SAPControl Start"`

	Runlevel string `xml:"runlevel,omitempty"`
}

type StartResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl StartResponse"`
}

type InstanceStart struct {
	XMLName xml.Name `xml:"urn:SAPControl InstanceStart"`

	Host string `xml:"host,omitempty"`

	Nr int32 `xml:"nr,omitempty"`

	Runlevel string `xml:"runlevel,omitempty"`
}

type InstanceStartResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl InstanceStartResponse"`
}

type Bootstrap struct {
	XMLName xml.Name `xml:"urn:SAPControl Bootstrap"`

	Host string `xml:"host,omitempty"`

	Nr int32 `xml:"nr,omitempty"`
}

type BootstrapResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl BootstrapResponse"`
}

type Stop struct {
	XMLName xml.Name `xml:"urn:SAPControl Stop"`

	Softtimeout int32 `xml:"softtimeout,omitempty"`

	IsSystemStop int32 `xml:"IsSystemStop,omitempty"`
}

type StopResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl StopResponse"`
}

type InstanceStop struct {
	XMLName xml.Name `xml:"urn:SAPControl InstanceStop"`

	Host string `xml:"host,omitempty"`

	Nr int32 `xml:"nr,omitempty"`

	Softtimeout int32 `xml:"softtimeout,omitempty"`
}

type InstanceStopResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl InstanceStopResponse"`
}

type Shutdown struct {
	XMLName xml.Name `xml:"urn:SAPControl Shutdown"`

	IsSystemStop int32 `xml:"IsSystemStop,omitempty"`
}

type ShutdownResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ShutdownResponse"`
}

type ParameterValue struct {
	XMLName xml.Name `xml:"urn:SAPControl ParameterValue"`

	Parameter string `xml:"parameter,omitempty"`
}

type ParameterValueResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ParameterValueResponse"`

	Value string `xml:"value,omitempty"`
}

type GetProcessList struct {
	XMLName xml.Name `xml:"urn:SAPControl GetProcessList"`
}

type GetProcessListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetProcessListResponse"`

	Process *ArrayOfOSProcess `xml:"process,omitempty"`
}

type GetStartProfile struct {
	XMLName xml.Name `xml:"urn:SAPControl GetStartProfile"`
}

type GetStartProfileResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetStartProfileResponse"`

	Name string `xml:"name,omitempty"`

	Lines *ArrayOfString `xml:"lines,omitempty"`
}

type GetTraceFile struct {
	XMLName xml.Name `xml:"urn:SAPControl GetTraceFile"`
}

type GetTraceFileResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetTraceFileResponse"`

	Name string `xml:"name,omitempty"`

	Lines *ArrayOfString `xml:"lines,omitempty"`
}

type GetAlertTree struct {
	XMLName xml.Name `xml:"urn:SAPControl GetAlertTree"`
}

type GetAlertTreeResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetAlertTreeResponse"`

	Tree *ArrayOfAlertNode `xml:"tree,omitempty"`
}

type GetAlerts struct {
	XMLName xml.Name `xml:"urn:SAPControl GetAlerts"`

	RootTid string `xml:"RootTid,omitempty"`
}

type GetAlertsResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetAlertsResponse"`

	RootTidName string `xml:"RootTidName,omitempty"`

	Alert *ArrayOfAlert `xml:"alert,omitempty"`
}

type RestartService struct {
	XMLName xml.Name `xml:"urn:SAPControl RestartService"`
}

type RestartServiceResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl RestartServiceResponse"`
}

type StopService struct {
	XMLName xml.Name `xml:"urn:SAPControl StopService"`
}

type StopServiceResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl StopServiceResponse"`
}

type GetEnvironment struct {
	XMLName xml.Name `xml:"urn:SAPControl GetEnvironment"`
}

type GetEnvironmentResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetEnvironmentResponse"`

	Env *ArrayOfString `xml:"env,omitempty"`
}

type ListDeveloperTraces struct {
	XMLName xml.Name `xml:"urn:SAPControl ListDeveloperTraces"`
}

type ListDeveloperTracesResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ListDeveloperTracesResponse"`

	File *ArrayOfDirEntry `xml:"file,omitempty"`
}

type ReadDeveloperTrace struct {
	XMLName xml.Name `xml:"urn:SAPControl ReadDeveloperTrace"`

	Filename string `xml:"filename,omitempty"`

	Size int32 `xml:"size,omitempty"`
}

type ReadDeveloperTraceResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ReadDeveloperTraceResponse"`

	Name string `xml:"name,omitempty"`

	Lines *ArrayOfString `xml:"lines,omitempty"`
}

type RestartInstance struct {
	XMLName xml.Name `xml:"urn:SAPControl RestartInstance"`

	Softtimeout int32 `xml:"softtimeout,omitempty"`

	Runlevel string `xml:"runlevel,omitempty"`
}

type RestartInstanceResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl RestartInstanceResponse"`
}

type SendSignal struct {
	XMLName xml.Name `xml:"urn:SAPControl SendSignal"`

	Pid int32 `xml:"pid,omitempty"`

	Signal string `xml:"signal,omitempty"`
}

type SendSignalResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl SendSignalResponse"`
}

type GetVersionInfo struct {
	XMLName xml.Name `xml:"urn:SAPControl GetVersionInfo"`
}

type GetVersionInfoResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetVersionInfoResponse"`

	Version *ArrayOfInstanceVersionInfo `xml:"version,omitempty"`
}

type GetQueueStatistic struct {
	XMLName xml.Name `xml:"urn:SAPControl GetQueueStatistic"`
}

type GetQueueStatisticResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetQueueStatisticResponse"`

	Queue *ArrayOfTaskHandlerQueue `xml:"queue,omitempty"`
}

type GetInstanceProperties struct {
	XMLName xml.Name `xml:"urn:SAPControl GetInstanceProperties"`
}

type GetInstancePropertiesResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetInstancePropertiesResponse"`

	Properties *ArrayOfInstanceProperties `xml:"properties,omitempty"`
}

type OSExecute struct {
	XMLName xml.Name `xml:"urn:SAPControl OSExecute"`

	Command string `xml:"command,omitempty"`

	Async int32 `xml:"async,omitempty"`

	Timeout int32 `xml:"timeout,omitempty"`

	Protocolfile string `xml:"protocolfile,omitempty"`
}

type OSExecuteResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl OSExecuteResponse"`

	Exitcode int32 `xml:"exitcode,omitempty"`

	Pid int32 `xml:"pid,omitempty"`

	Lines *ArrayOfString `xml:"lines,omitempty"`
}

type ReadLogFile struct {
	XMLName xml.Name `xml:"urn:SAPControl ReadLogFile"`

	Filename string `xml:"filename,omitempty"`

	Filter string `xml:"filter,omitempty"`

	Language string `xml:"language,omitempty"`

	Maxentries int32 `xml:"maxentries,omitempty"`

	Statecookie string `xml:"statecookie,omitempty"`
}

type ReadLogFileResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ReadLogFileResponse"`

	Format string `xml:"format,omitempty"`

	Startcookie string `xml:"startcookie,omitempty"`

	Endcookie string `xml:"endcookie,omitempty"`

	Fields *ArrayOfString `xml:"fields,omitempty"`
}

type ListLogFiles struct {
	XMLName xml.Name `xml:"urn:SAPControl ListLogFiles"`
}

type ListLogFilesResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ListLogFilesResponse"`

	File *ArrayOfLogFile `xml:"file,omitempty"`
}

type AnalyseLogFiles struct {
	XMLName xml.Name `xml:"urn:SAPControl AnalyseLogFiles"`

	Starttime string `xml:"starttime,omitempty"`

	Endtime string `xml:"endtime,omitempty"`

	Severitylevel int32 `xml:"severity-level,omitempty"`

	Maxentries int32 `xml:"maxentries,omitempty"`
}

type AnalyseLogFilesResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl AnalyseLogFilesResponse"`

	Format string `xml:"format,omitempty"`

	Fields *ArrayOfString `xml:"fields,omitempty"`
}

type GetAccessPointList struct {
	XMLName xml.Name `xml:"urn:SAPControl GetAccessPointList"`
}

type GetAccessPointListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetAccessPointListResponse"`

	Accesspoint *ArrayOfAccessPoint `xml:"accesspoint,omitempty"`
}

type GetSystemInstanceList struct {
	XMLName xml.Name `xml:"urn:SAPControl GetSystemInstanceList"`

	Timeout int32 `xml:"timeout,omitempty"`
}

type GetSystemInstanceListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetSystemInstanceListResponse"`

	Instance *ArrayOfSAPInstance `xml:"instance,omitempty"`
}

type GetSystemUpdateList struct {
	XMLName xml.Name `xml:"urn:SAPControl GetSystemUpdateList"`
}

type GetSystemUpdateListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetSystemUpdateListResponse"`

	Instance *ArrayOfUpdateInstance `xml:"instance,omitempty"`
}

type StartSystem struct {
	XMLName xml.Name `xml:"urn:SAPControl StartSystem"`

	Options *StartStopOption `xml:"options,omitempty"`

	Prioritylevel string `xml:"prioritylevel,omitempty"`

	Waittimeout int32 `xml:"waittimeout,omitempty"`

	Runlevel string `xml:"runlevel,omitempty"`
}

type StartSystemResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl StartSystemResponse"`
}

type StopSystem struct {
	XMLName xml.Name `xml:"urn:SAPControl StopSystem"`

	Options *StartStopOption `xml:"options,omitempty"`

	Prioritylevel string `xml:"prioritylevel,omitempty"`

	Softtimeout int32 `xml:"softtimeout,omitempty"`

	Waittimeout int32 `xml:"waittimeout,omitempty"`
}

type StopSystemResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl StopSystemResponse"`
}

type RestartSystem struct {
	XMLName xml.Name `xml:"urn:SAPControl RestartSystem"`

	Options *StartStopOption `xml:"options,omitempty"`

	Prioritylevel string `xml:"prioritylevel,omitempty"`

	Softtimeout int32 `xml:"softtimeout,omitempty"`

	Waittimeout int32 `xml:"waittimeout,omitempty"`

	Runlevel string `xml:"runlevel,omitempty"`
}

type RestartSystemResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl RestartSystemResponse"`
}

type UpdateSystem struct {
	XMLName xml.Name `xml:"urn:SAPControl UpdateSystem"`

	Softtimeout int32 `xml:"softtimeout,omitempty"`

	Waittimeout int32 `xml:"waittimeout,omitempty"`

	Force bool `xml:"force,omitempty"`
}

type UpdateSystemResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl UpdateSystemResponse"`
}

type UpdateSCSInstance struct {
	XMLName xml.Name `xml:"urn:SAPControl UpdateSCSInstance"`
}

type UpdateSCSInstanceResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl UpdateSCSInstanceResponse"`
}

type CheckUpdateSystem struct {
	XMLName xml.Name `xml:"urn:SAPControl CheckUpdateSystem"`
}

type CheckUpdateSystemResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl CheckUpdateSystemResponse"`
}

type AccessCheck struct {
	XMLName xml.Name `xml:"urn:SAPControl AccessCheck"`

	Function string `xml:"function,omitempty"`
}

type AccessCheckResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl AccessCheckResponse"`
}

type GetProcessParameter struct {
	XMLName xml.Name `xml:"urn:SAPControl GetProcessParameter"`

	Processtype string `xml:"processtype,omitempty"`

	Pid int32 `xml:"pid,omitempty"`
}

type GetProcessParameterResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetProcessParameterResponse"`

	Parameter *ArrayOfProfileParameter `xml:"parameter,omitempty"`
}

type SetProcessParameter struct {
	XMLName xml.Name `xml:"urn:SAPControl SetProcessParameter"`

	Processtype string `xml:"processtype,omitempty"`

	Pid int32 `xml:"pid,omitempty"`

	Parameter *ArrayOfSetProfileParameter `xml:"parameter,omitempty"`
}

type SetProcessParameterResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl SetProcessParameterResponse"`
}

type SetProcessParameter2 struct {
	XMLName xml.Name `xml:"urn:SAPControl SetProcessParameter2"`

	Processtype string `xml:"processtype,omitempty"`

	Pid int32 `xml:"pid,omitempty"`

	Persistence *PERSISTENCETYPE `xml:"persistence,omitempty"`

	Parameter *ArrayOfSetProfileParameter `xml:"parameter,omitempty"`
}

type SetProcessParameter2Response struct {
	XMLName xml.Name `xml:"urn:SAPControl SetProcessParameter2Response"`
}

type ShmDetach struct {
	XMLName xml.Name `xml:"urn:SAPControl ShmDetach"`
}

type ShmDetachResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ShmDetachResponse"`
}

type GetSecNetworkId struct {
	XMLName xml.Name `xml:"urn:SAPControl GetSecNetworkId"`

	Serviceip string `xml:"service-ip,omitempty"`

	Serviceport int32 `xml:"service-port,omitempty"`

	Version int32 `xml:"version,omitempty"`

	Challenge string `xml:"challenge,omitempty"`
}

type GetSecNetworkIdResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetSecNetworkIdResponse"`

	Key string `xml:"key,omitempty"`

	Proof string `xml:"proof,omitempty"`
}

type GetNetworkId struct {
	XMLName xml.Name `xml:"urn:SAPControl GetNetworkId"`

	Serviceip string `xml:"service-ip,omitempty"`

	Serviceport int32 `xml:"service-port,omitempty"`

	Version int32 `xml:"version,omitempty"`
}

type GetNetworkIdResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetNetworkIdResponse"`

	Key string `xml:"key,omitempty"`
}

type CreateSnapshot struct {
	XMLName xml.Name `xml:"urn:SAPControl CreateSnapshot"`

	Description string `xml:"description,omitempty"`

	Datcolparam string `xml:"datcol-param,omitempty"`

	Analyseseveritylevel int32 `xml:"analyse-severity-level,omitempty"`

	Analysestarttime string `xml:"analyse-starttime,omitempty"`

	Analyseendtime string `xml:"analyse-endtime,omitempty"`

	Analyzemaxentries int32 `xml:"analyze-maxentries,omitempty"`

	Maxentries int32 `xml:"maxentries,omitempty"`

	Logfiles *ArrayOfString `xml:"logfiles,omitempty"`
}

type CreateSnapshotResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl CreateSnapshotResponse"`

	Filename string `xml:"filename,omitempty"`
}

type ReadSnapshot struct {
	XMLName xml.Name `xml:"urn:SAPControl ReadSnapshot"`

	Filename string `xml:"filename,omitempty"`
}

type ReadSnapshotResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ReadSnapshotResponse"`

	Snapshot *SnapshotZip `xml:"snapshot,omitempty"`
}

type ListSnapshots struct {
	XMLName xml.Name `xml:"urn:SAPControl ListSnapshots"`
}

type ListSnapshotsResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ListSnapshotsResponse"`

	Snapshots *ArrayOfSnapshotInfo `xml:"snapshots,omitempty"`
}

type DeleteSnapshots struct {
	XMLName xml.Name `xml:"urn:SAPControl DeleteSnapshots"`

	Snapshots *ArrayOfString `xml:"snapshots,omitempty"`
}

type DeleteSnapshotsResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl DeleteSnapshotsResponse"`
}

type ABAPReadSyslog struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPReadSyslog"`
}

type ABAPReadSyslogResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPReadSyslogResponse"`

	Log *ArrayOfSyslogEntry `xml:"log,omitempty"`
}

type ABAPReadRawSyslog struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPReadRawSyslog"`
}

type ABAPReadRawSyslogResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPReadRawSyslogResponse"`

	Log *ArrayOfRawSyslogEntry `xml:"log,omitempty"`
}

type ABAPGetWPTable struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPGetWPTable"`
}

type ABAPGetWPTableResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPGetWPTableResponse"`

	Workprocess *ArrayOfWorkProcess `xml:"workprocess,omitempty"`
}

type ABAPAcknowledgeAlerts struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPAcknowledgeAlerts"`

	R3Client string `xml:"R3Client,omitempty"`

	R3User string `xml:"R3User,omitempty"`

	R3Password string `xml:"R3Password,omitempty"`

	Aid *ArrayOfString `xml:"Aid,omitempty"`
}

type ABAPAcknowledgeAlertsResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPAcknowledgeAlertsResponse"`

	Alert *ArrayOfInt `xml:"alert,omitempty"`
}

type ABAPGetComponentList struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPGetComponentList"`
}

type ABAPGetComponentListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPGetComponentListResponse"`

	Component *ArrayOfABAPComponent `xml:"component,omitempty"`
}

type ABAPCheckRFCDestinations struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPCheckRFCDestinations"`
}

type ABAPCheckRFCDestinationsResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPCheckRFCDestinationsResponse"`

	Destination *ArrayOfString `xml:"destination,omitempty"`
}

type ABAPGetSystemWPTable struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPGetSystemWPTable"`

	Activeonly bool `xml:"activeonly,omitempty"`
}

type ABAPGetSystemWPTableResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ABAPGetSystemWPTableResponse"`

	Workprocess *ArrayOfSystemWorkProcess `xml:"workprocess,omitempty"`
}

type RequestTicket struct {
	XMLName xml.Name `xml:"urn:SAPControl RequestTicket"`

	Host string `xml:"host,omitempty"`

	Nr int32 `xml:"nr,omitempty"`

	Challenge string `xml:"challenge,omitempty"`
}

type RequestTicketResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl RequestTicketResponse"`
}

type SendTicket struct {
	XMLName xml.Name `xml:"urn:SAPControl SendTicket"`

	Proof string `xml:"proof,omitempty"`

	Challenge string `xml:"challenge,omitempty"`
}

type SendTicketResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl SendTicketResponse"`
}

type ConfigureLogFileList struct {
	XMLName xml.Name `xml:"urn:SAPControl ConfigureLogFileList"`

	Operation *LogFileConfigOperation `xml:"operation,omitempty"`

	Logfiles *ArrayOfString `xml:"logfiles,omitempty"`
}

type ConfigureLogFileListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ConfigureLogFileListResponse"`
}

type GetLogFileList struct {
	XMLName xml.Name `xml:"urn:SAPControl GetLogFileList"`
}

type GetLogFileListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetLogFileListResponse"`

	Logfiles *ArrayOfString `xml:"logfiles,omitempty"`
}

type RequestLogonFile struct {
	XMLName xml.Name `xml:"urn:SAPControl RequestLogonFile"`

	User string `xml:"user,omitempty"`
}

type RequestLogonFileResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl RequestLogonFileResponse"`

	Filename string `xml:"filename,omitempty"`
}

type UpdateSystemPKI struct {
	XMLName xml.Name `xml:"urn:SAPControl UpdateSystemPKI"`

	Force bool `xml:"force,omitempty"`
}

type UpdateSystemPKIResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl UpdateSystemPKIResponse"`
}

type UpdateInstancePSE struct {
	XMLName xml.Name `xml:"urn:SAPControl UpdateInstancePSE"`

	Force bool `xml:"force,omitempty"`
}

type UpdateInstancePSEResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl UpdateInstancePSEResponse"`
}

type HACheckConfig struct {
	XMLName xml.Name `xml:"urn:SAPControl HACheckConfig"`
}

type HACheckConfigResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl HACheckConfigResponse"`

	Check *ArrayOfHACheck `xml:"check,omitempty"`
}

type HACheckFailoverConfig struct {
	XMLName xml.Name `xml:"urn:SAPControl HACheckFailoverConfig"`
}

type HACheckFailoverConfigResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl HACheckFailoverConfigResponse"`

	Check *ArrayOfHACheck `xml:"check,omitempty"`
}

type HAGetFailoverConfig struct {
	XMLName xml.Name `xml:"urn:SAPControl HAGetFailoverConfig"`
}

type HAGetFailoverConfigResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl HAGetFailoverConfigResponse"`

	HAActive bool `xml:"HAActive,omitempty"`

	HAProductVersion string `xml:"HAProductVersion,omitempty"`

	HASAPInterfaceVersion string `xml:"HASAPInterfaceVersion,omitempty"`

	HADocumentation string `xml:"HADocumentation,omitempty"`

	HAActiveNode string `xml:"HAActiveNode,omitempty"`

	HANodes *ArrayOfString `xml:"HANodes,omitempty"`
}

type HAFailoverToNode struct {
	XMLName xml.Name `xml:"urn:SAPControl HAFailoverToNode"`

	Node string `xml:"node,omitempty"`
}

type HAFailoverToNodeResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl HAFailoverToNodeResponse"`
}

type StartBypassHA struct {
	XMLName xml.Name `xml:"urn:SAPControl StartBypassHA"`

	Runlevel string `xml:"runlevel,omitempty"`
}

type StartBypassHAResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl StartBypassHAResponse"`
}

type StopBypassHA struct {
	XMLName xml.Name `xml:"urn:SAPControl StopBypassHA"`

	Softtimeout int32 `xml:"softtimeout,omitempty"`

	IsSystemStop int32 `xml:"IsSystemStop,omitempty"`
}

type StopBypassHAResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl StopBypassHAResponse"`
}

type GetCallstack struct {
	XMLName xml.Name `xml:"urn:SAPControl GetCallstack"`

	Pid int32 `xml:"pid,omitempty"`
}

type GetCallstackResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl GetCallstackResponse"`

	Lines *ArrayOfString `xml:"lines,omitempty"`
}

type J2EEGetProcessList struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetProcessList"`
}

type J2EEGetProcessListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetProcessListResponse"`

	Process *ArrayOfJ2EEProcess `xml:"process,omitempty"`
}

type J2EEGetProcessList2 struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetProcessList2"`
}

type J2EEGetProcessList2Response struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetProcessList2Response"`

	Process *ArrayOfJ2EEProcess2 `xml:"process,omitempty"`
}

type J2EEControlProcess struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEControlProcess"`

	Processname string `xml:"processname,omitempty"`

	Function string `xml:"function,omitempty"`
}

type J2EEControlProcessResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEControlProcessResponse"`
}

type J2EEControlCluster struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEControlCluster"`

	Processname string `xml:"processname,omitempty"`

	Function string `xml:"function,omitempty"`

	Host string `xml:"host,omitempty"`

	Nr int32 `xml:"nr,omitempty"`
}

type J2EEControlClusterResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEControlClusterResponse"`
}

type J2EEGetThreadList struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetThreadList"`
}

type J2EEGetThreadListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetThreadListResponse"`

	Thread *ArrayOfJ2EEThread `xml:"thread,omitempty"`
}

type J2EEGetThreadList2 struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetThreadList2"`
}

type J2EEGetThreadList2Response struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetThreadList2Response"`

	Thread *ArrayOfJ2EEThread2 `xml:"thread,omitempty"`
}

type J2EEGetSessionList struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetSessionList"`
}

type J2EEGetSessionListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetSessionListResponse"`

	Session *ArrayOfJ2EESession `xml:"session,omitempty"`
}

type J2EEGetWebSessionList struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetWebSessionList"`
}

type J2EEGetWebSessionListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetWebSessionListResponse"`

	Session *ArrayOfJ2EEWebSession `xml:"session,omitempty"`
}

type J2EEGetWebSessionList2 struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetWebSessionList2"`
}

type J2EEGetWebSessionList2Response struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetWebSessionList2Response"`

	Session *ArrayOfJ2EEWebSession2 `xml:"session,omitempty"`
}

type J2EEGetCacheStatistic struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetCacheStatistic"`
}

type J2EEGetCacheStatisticResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetCacheStatisticResponse"`

	Cache *ArrayOfJ2EECache `xml:"cache,omitempty"`
}

type J2EEGetCacheStatistic2 struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetCacheStatistic2"`
}

type J2EEGetCacheStatistic2Response struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetCacheStatistic2Response"`

	Cache *ArrayOfJ2EECache2 `xml:"cache,omitempty"`
}

type J2EEGetApplicationAliasList struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetApplicationAliasList"`
}

type J2EEGetApplicationAliasListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetApplicationAliasListResponse"`

	Alias *ArrayOfJ2EEApplicationAlias `xml:"alias,omitempty"`
}

type J2EEGetVMGCHistory struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetVMGCHistory"`
}

type J2EEGetVMGCHistoryResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetVMGCHistoryResponse"`

	Gc *ArrayOfGCInfo `xml:"gc,omitempty"`
}

type J2EEGetVMGCHistory2 struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetVMGCHistory2"`
}

type J2EEGetVMGCHistory2Response struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetVMGCHistory2Response"`

	Gc *ArrayOfGCInfo2 `xml:"gc,omitempty"`
}

type J2EEGetVMHeapInfo struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetVMHeapInfo"`
}

type J2EEGetVMHeapInfoResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetVMHeapInfoResponse"`

	Heap *ArrayOfHeapInfo `xml:"heap,omitempty"`
}

type J2EEGetEJBSessionList struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetEJBSessionList"`
}

type J2EEGetEJBSessionListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetEJBSessionListResponse"`

	Ejbsession *ArrayOfJ2EEEJBSession `xml:"ejbsession,omitempty"`
}

type J2EEGetRemoteObjectList struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetRemoteObjectList"`
}

type J2EEGetRemoteObjectListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetRemoteObjectListResponse"`

	Remoteobject *ArrayOfJ2EERemoteObject `xml:"remoteobject,omitempty"`
}

type J2EEGetClusterMsgList struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetClusterMsgList"`
}

type J2EEGetClusterMsgListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetClusterMsgListResponse"`

	Msg *ArrayOfJ2EEClusterMsg `xml:"msg,omitempty"`
}

type J2EEGetSharedTableInfo struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetSharedTableInfo"`
}

type J2EEGetSharedTableInfoResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetSharedTableInfoResponse"`

	Jsf *ArrayOfJ2EESharedTableInfo `xml:"jsf,omitempty"`
}

type J2EEEnableDbgSession struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEEnableDbgSession"`

	Processname string `xml:"processname,omitempty"`

	Flags string `xml:"flags,omitempty"`

	Client string `xml:"client,omitempty"`
}

type J2EEEnableDbgSessionResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEEnableDbgSessionResponse"`

	Key string `xml:"key,omitempty"`

	Port int32 `xml:"port,omitempty"`
}

type J2EEDisableDbgSession struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEDisableDbgSession"`

	Key string `xml:"key,omitempty"`
}

type J2EEDisableDbgSessionResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEDisableDbgSessionResponse"`
}

type J2EEGetThreadCallStack struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetThreadCallStack"`

	Index int32 `xml:"index,omitempty"`
}

type J2EEGetThreadCallStackResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetThreadCallStackResponse"`

	Name string `xml:"name,omitempty"`

	Lines *ArrayOfString `xml:"lines,omitempty"`
}

type J2EEGetThreadTaskStack struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetThreadTaskStack"`

	Index int32 `xml:"index,omitempty"`
}

type J2EEGetThreadTaskStackResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetThreadTaskStackResponse"`

	Name string `xml:"name,omitempty"`

	Lines *ArrayOfString `xml:"lines,omitempty"`
}

type J2EEGetComponentList struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetComponentList"`
}

type J2EEGetComponentListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEGetComponentListResponse"`

	Component *ArrayOfJ2EEComponentInfo `xml:"component,omitempty"`
}

type J2EEControlComponents struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEControlComponents"`

	ProcessName string `xml:"processName,omitempty"`

	Operation string `xml:"operation,omitempty"`

	ComponentType string `xml:"componentType,omitempty"`

	ComponentNames string `xml:"componentNames,omitempty"`
}

type J2EEControlComponentsResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl J2EEControlComponentsResponse"`
}

type ICMGetThreadList struct {
	XMLName xml.Name `xml:"urn:SAPControl ICMGetThreadList"`
}

type ICMGetThreadListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ICMGetThreadListResponse"`

	Thread *ArrayOfICMThread `xml:"thread,omitempty"`
}

type ICMGetConnectionList struct {
	XMLName xml.Name `xml:"urn:SAPControl ICMGetConnectionList"`
}

type ICMGetConnectionListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ICMGetConnectionListResponse"`

	Connection *ArrayOfICMConnection `xml:"connection,omitempty"`
}

type ICMGetCacheEntries struct {
	XMLName xml.Name `xml:"urn:SAPControl ICMGetCacheEntries"`
}

type ICMGetCacheEntriesResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ICMGetCacheEntriesResponse"`

	Entry *ArrayOfICMCacheEntry `xml:"entry,omitempty"`
}

type ICMGetProxyConnectionList struct {
	XMLName xml.Name `xml:"urn:SAPControl ICMGetProxyConnectionList"`
}

type ICMGetProxyConnectionListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl ICMGetProxyConnectionListResponse"`

	Connection *ArrayOfICMProxyConnection `xml:"connection,omitempty"`
}

type WebDispGetServerList struct {
	XMLName xml.Name `xml:"urn:SAPControl WebDispGetServerList"`
}

type WebDispGetServerListResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl WebDispGetServerListResponse"`

	Server *ArrayOfWebDispServer `xml:"server,omitempty"`
}

type EnqGetLockTable struct {
	XMLName xml.Name `xml:"urn:SAPControl EnqGetLockTable"`
}

type EnqGetLockTableResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl EnqGetLockTableResponse"`

	Lock *ArrayOfEnqLock `xml:"lock,omitempty"`
}

type EnqRemoveLocks struct {
	XMLName xml.Name `xml:"urn:SAPControl EnqRemoveLocks"`

	Lock *ArrayOfEnqLock `xml:"lock,omitempty"`
}

type EnqRemoveLocksResponse struct {
	XMLName xml.Name `xml:"urn:SAPControl EnqRemoveLocksResponse"`
}

type EnqGetStatistic struct {
	XMLName xml.Name `xml:"urn:SAPControl EnqGetStatistic"`
}

type EnqStatistic struct {
	XMLName xml.Name `xml:"urn:SAPControl EnqStatistic"`

	Ownernow int32 `xml:"owner-now,omitempty"`

	Ownerhigh int32 `xml:"owner-high,omitempty"`

	Ownermax int32 `xml:"owner-max,omitempty"`

	Ownerstate *STATECOLOR `xml:"owner-state,omitempty"`

	Argumentsnow int32 `xml:"arguments-now,omitempty"`

	Argumentshigh int32 `xml:"arguments-high,omitempty"`

	Argumentsmax int32 `xml:"arguments-max,omitempty"`

	Argumentsstate *STATECOLOR `xml:"arguments-state,omitempty"`

	Locksnow int32 `xml:"locks-now,omitempty"`

	Lockshigh int32 `xml:"locks-high,omitempty"`

	Locksmax int32 `xml:"locks-max,omitempty"`

	Locksstate *STATECOLOR `xml:"locks-state,omitempty"`

	Enqueuerequests int64 `xml:"enqueue-requests,omitempty"`

	Enqueuerejects int64 `xml:"enqueue-rejects,omitempty"`

	Enqueueerrors int64 `xml:"enqueue-errors,omitempty"`

	Dequeuerequests int64 `xml:"dequeue-requests,omitempty"`

	Dequeueerrors int64 `xml:"dequeue-errors,omitempty"`

	Dequeueallrequests int64 `xml:"dequeue-all-requests,omitempty"`

	Cleanuprequests int64 `xml:"cleanup-requests,omitempty"`

	Backuprequests int64 `xml:"backup-requests,omitempty"`

	Reportingrequests int64 `xml:"reporting-requests,omitempty"`

	Compressrequests int64 `xml:"compress-requests,omitempty"`

	Verifyrequests int64 `xml:"verify-requests,omitempty"`

	Locktime float64 `xml:"lock-time,omitempty"`

	Lockwaittime float64 `xml:"lock-wait-time,omitempty"`

	Servertime float64 `xml:"server-time,omitempty"`

	Replicationstate *STATECOLOR `xml:"replication-state,omitempty"`
}

type ArrayOfString struct {
	Item []string `xml:"item,omitempty"`
}

type ArrayOfInt struct {
	Item []int32 `xml:"item,omitempty"`
}

type OSProcess struct {
	Name string `xml:"name,omitempty"`

	Description string `xml:"description,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`

	Textstatus string `xml:"textstatus,omitempty"`

	Starttime string `xml:"starttime,omitempty"`

	Elapsedtime string `xml:"elapsedtime,omitempty"`

	Pid int32 `xml:"pid,omitempty"`
}

type ArrayOfOSProcess struct {
	Item []*OSProcess `xml:"item,omitempty"`
}

type AlertNode struct {
	Name string `xml:"name,omitempty"`

	Parent int32 `xml:"parent,omitempty"`

	ActualValue *STATECOLOR `xml:"ActualValue,omitempty"`

	Description string `xml:"description,omitempty"`

	Time string `xml:"Time,omitempty"`

	AnalyseTool string `xml:"AnalyseTool,omitempty"`

	VisibleLevel *VISIBLELEVEL `xml:"VisibleLevel,omitempty"`

	HighAlertValue *STATECOLOR `xml:"HighAlertValue,omitempty"`

	AlDescription string `xml:"AlDescription,omitempty"`

	AlTime string `xml:"AlTime,omitempty"`

	Tid string `xml:"Tid,omitempty"`
}

type ArrayOfAlertNode struct {
	Item []*AlertNode `xml:"item,omitempty"`
}

type DirEntry struct {
	Filename string `xml:"filename,omitempty"`

	Size int64 `xml:"size,omitempty"`

	Modtime string `xml:"modtime,omitempty"`
}

type ArrayOfDirEntry struct {
	Item []*DirEntry `xml:"item,omitempty"`
}

type SyslogEntry struct {
	Time string `xml:"Time,omitempty"`

	Typ string `xml:"Typ,omitempty"`

	Client string `xml:"Client,omitempty"`

	User string `xml:"User,omitempty"`

	Tcode string `xml:"Tcode,omitempty"`

	MNo string `xml:"MNo,omitempty"`

	Text string `xml:"Text,omitempty"`

	Severity *STATECOLOR `xml:"Severity,omitempty"`
}

type ArrayOfSyslogEntry struct {
	Item []*SyslogEntry `xml:"item,omitempty"`
}

type ArrayOfRawSyslogEntry struct {
	Item []string `xml:"item,omitempty"`
}

type TaskHandlerQueue struct {
	Typ string `xml:"Typ,omitempty"`

	Now int32 `xml:"Now,omitempty"`

	High int32 `xml:"High,omitempty"`

	Max int32 `xml:"Max,omitempty"`

	Writes int32 `xml:"Writes,omitempty"`

	Reads int32 `xml:"Reads,omitempty"`
}

type ArrayOfTaskHandlerQueue struct {
	Item []*TaskHandlerQueue `xml:"item,omitempty"`
}

type WorkProcess struct {
	No int32 `xml:"No,omitempty"`

	Typ string `xml:"Typ,omitempty"`

	Pid int32 `xml:"Pid,omitempty"`

	Status string `xml:"Status,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	Start string `xml:"Start,omitempty"`

	Err string `xml:"Err,omitempty"`

	Sem string `xml:"Sem,omitempty"`

	Cpu string `xml:"Cpu,omitempty"`

	Time string `xml:"Time,omitempty"`

	Program string `xml:"Program,omitempty"`

	Client string `xml:"Client,omitempty"`

	User string `xml:"User,omitempty"`

	Action string `xml:"Action,omitempty"`

	Table string `xml:"Table,omitempty"`
}

type ArrayOfWorkProcess struct {
	Item []*WorkProcess `xml:"item,omitempty"`
}

type SystemWorkProcess struct {
	Instance string `xml:"Instance,omitempty"`

	No int32 `xml:"No,omitempty"`

	Typ string `xml:"Typ,omitempty"`

	Pid int32 `xml:"Pid,omitempty"`

	Status string `xml:"Status,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	Start string `xml:"Start,omitempty"`

	Err string `xml:"Err,omitempty"`

	Sem string `xml:"Sem,omitempty"`

	Cpu string `xml:"Cpu,omitempty"`

	Time string `xml:"Time,omitempty"`

	Program string `xml:"Program,omitempty"`

	Client string `xml:"Client,omitempty"`

	User string `xml:"User,omitempty"`

	Action string `xml:"Action,omitempty"`

	Table string `xml:"Table,omitempty"`
}

type ArrayOfSystemWorkProcess struct {
	Item []*SystemWorkProcess `xml:"item,omitempty"`
}

type Alert struct {
	Object string `xml:"Object,omitempty"`

	Attribute string `xml:"Attribute,omitempty"`

	Value *STATECOLOR `xml:"Value,omitempty"`

	Description string `xml:"Description,omitempty"`

	Time string `xml:"Time,omitempty"`

	Tid string `xml:"Tid,omitempty"`

	Aid string `xml:"Aid,omitempty"`
}

type ArrayOfAlert struct {
	Item []*Alert `xml:"item,omitempty"`
}

type J2EEProcess struct {
	TelnetPort int32 `xml:"telnetPort,omitempty"`

	Name string `xml:"name,omitempty"`

	Pid int32 `xml:"pid,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Restart string `xml:"restart,omitempty"`

	ExitCode string `xml:"exitCode,omitempty"`

	State *J2EEPSTATE `xml:"state,omitempty"`

	Statetext string `xml:"statetext,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	ElapsedTime string `xml:"elapsedTime,omitempty"`

	RestartCount int32 `xml:"restartCount,omitempty"`

	ErrorCount int32 `xml:"errorCount,omitempty"`

	Cpu string `xml:"cpu,omitempty"`

	Debug string `xml:"debug,omitempty"`
}

type ArrayOfJ2EEProcess struct {
	Item []*J2EEProcess `xml:"item,omitempty"`
}

type J2EEProcess2 struct {
	TelnetPort int32 `xml:"telnetPort,omitempty"`

	Name string `xml:"name,omitempty"`

	Pid int32 `xml:"pid,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Restart string `xml:"restart,omitempty"`

	ExitCode string `xml:"exitCode,omitempty"`

	State *J2EEPSTATE `xml:"state,omitempty"`

	Statetext string `xml:"statetext,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	ElapsedTime string `xml:"elapsedTime,omitempty"`

	RestartCount int32 `xml:"restartCount,omitempty"`

	ErrorCount int32 `xml:"errorCount,omitempty"`

	Cpu string `xml:"cpu,omitempty"`

	Debug string `xml:"debug,omitempty"`

	ClusterId int32 `xml:"clusterId,omitempty"`
}

type ArrayOfJ2EEProcess2 struct {
	Item []*J2EEProcess2 `xml:"item,omitempty"`
}

type J2EEThread struct {
	Processname string `xml:"processname,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	UpdateTime string `xml:"updateTime,omitempty"`

	TaskupdateTime string `xml:"taskupdateTime,omitempty"`

	SubtaskupdateTime string `xml:"subtaskupdateTime,omitempty"`

	Task string `xml:"task,omitempty"`

	Subtask string `xml:"subtask,omitempty"`

	Name string `xml:"name,omitempty"`

	Classname string `xml:"classname,omitempty"`

	User string `xml:"user,omitempty"`

	Pool string `xml:"pool,omitempty"`

	State string `xml:"state,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`
}

type ArrayOfJ2EEThread struct {
	Item []*J2EEThread `xml:"item,omitempty"`
}

type J2EEThread2 struct {
	Processname string `xml:"processname,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	UpdateTime string `xml:"updateTime,omitempty"`

	TaskupdateTime string `xml:"taskupdateTime,omitempty"`

	SubtaskupdateTime string `xml:"subtaskupdateTime,omitempty"`

	Task string `xml:"task,omitempty"`

	Subtask string `xml:"subtask,omitempty"`

	Name string `xml:"name,omitempty"`

	Classname string `xml:"classname,omitempty"`

	User string `xml:"user,omitempty"`

	Pool string `xml:"pool,omitempty"`

	State string `xml:"state,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`

	Index int32 `xml:"index,omitempty"`
}

type ArrayOfJ2EEThread2 struct {
	Item []*J2EEThread2 `xml:"item,omitempty"`
}

type J2EESession struct {
	Processname string `xml:"processname,omitempty"`

	IdHash int32 `xml:"IdHash,omitempty"`

	Size int32 `xml:"size,omitempty"`

	Timeout int32 `xml:"timeout,omitempty"`

	ActiveRequests int32 `xml:"activeRequests,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	UpdateTime string `xml:"updateTime,omitempty"`

	Sticky string `xml:"sticky,omitempty"`

	Corrupt string `xml:"corrupt,omitempty"`

	BackingStore string `xml:"backingStore,omitempty"`
}

type ArrayOfJ2EESession struct {
	Item []*J2EESession `xml:"item,omitempty"`
}

type J2EEWebSession struct {
	Processname string `xml:"processname,omitempty"`

	IdHash int32 `xml:"IdHash,omitempty"`

	Size int32 `xml:"size,omitempty"`

	Timeout int32 `xml:"timeout,omitempty"`

	ActiveRequests int32 `xml:"activeRequests,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	UpdateTime string `xml:"updateTime,omitempty"`

	State string `xml:"state,omitempty"`

	BackingStore string `xml:"backingStore,omitempty"`

	User string `xml:"user,omitempty"`
}

type ArrayOfJ2EEWebSession struct {
	Item []*J2EEWebSession `xml:"item,omitempty"`
}

type J2EEWebSession2 struct {
	Processname string `xml:"processname,omitempty"`

	IdHash int32 `xml:"IdHash,omitempty"`

	Size int32 `xml:"size,omitempty"`

	Timeout int32 `xml:"timeout,omitempty"`

	ActiveRequests int32 `xml:"activeRequests,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	UpdateTime string `xml:"updateTime,omitempty"`

	State string `xml:"state,omitempty"`

	BackingStore string `xml:"backingStore,omitempty"`

	User string `xml:"user,omitempty"`

	AppName string `xml:"AppName,omitempty"`
}

type ArrayOfJ2EEWebSession2 struct {
	Item []*J2EEWebSession2 `xml:"item,omitempty"`
}

type J2EECache struct {
	Cachename string `xml:"cachename,omitempty"`

	Processname string `xml:"processname,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Size int64 `xml:"size,omitempty"`

	AttrSize int64 `xml:"attrSize,omitempty"`

	KeysSize int64 `xml:"keysSize,omitempty"`

	CachedObjects int32 `xml:"cachedObjects,omitempty"`

	UsedObjects int32 `xml:"usedObjects,omitempty"`

	Puts int32 `xml:"puts,omitempty"`

	Gets int32 `xml:"gets,omitempty"`

	Hits int32 `xml:"hits,omitempty"`

	Changes int32 `xml:"changes,omitempty"`

	Removes int32 `xml:"removes,omitempty"`

	Evictions int32 `xml:"evictions,omitempty"`

	InstanceInvalidations int32 `xml:"instanceInvalidations,omitempty"`

	ClusterInvalidations int32 `xml:"clusterInvalidations,omitempty"`

	UpdateTime string `xml:"updateTime,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`
}

type ArrayOfJ2EECache struct {
	Item []*J2EECache `xml:"item,omitempty"`
}

type J2EECache2 struct {
	Description string `xml:"description,omitempty"`

	Owner string `xml:"owner,omitempty"`

	Processname string `xml:"processname,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Size int64 `xml:"size,omitempty"`

	AttrSize int64 `xml:"attrSize,omitempty"`

	KeysSize int64 `xml:"keysSize,omitempty"`

	CachedObjects int32 `xml:"cachedObjects,omitempty"`

	UsedObjects int32 `xml:"usedObjects,omitempty"`

	Puts int32 `xml:"puts,omitempty"`

	Gets int32 `xml:"gets,omitempty"`

	Hits int32 `xml:"hits,omitempty"`

	Changes int32 `xml:"changes,omitempty"`

	Removes int32 `xml:"removes,omitempty"`

	Evictions int32 `xml:"evictions,omitempty"`

	InstanceInvalidations int32 `xml:"instanceInvalidations,omitempty"`

	ClusterInvalidations int32 `xml:"clusterInvalidations,omitempty"`

	UpdateTime string `xml:"updateTime,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`
}

type ArrayOfJ2EECache2 struct {
	Item []*J2EECache2 `xml:"item,omitempty"`
}

type J2EEApplicationAlias struct {
	AppName string `xml:"AppName,omitempty"`

	Alias string `xml:"Alias,omitempty"`

	TotalRequests int32 `xml:"TotalRequests,omitempty"`

	AppActive string `xml:"AppActive,omitempty"`

	IgnoreCookie string `xml:"IgnoreCookie,omitempty"`
}

type ArrayOfJ2EEApplicationAlias struct {
	Item []*J2EEApplicationAlias `xml:"item,omitempty"`
}

type InstanceVersionInfo struct {
	Filename string `xml:"Filename,omitempty"`

	VersionInfo string `xml:"VersionInfo,omitempty"`

	Time string `xml:"Time,omitempty"`
}

type ArrayOfInstanceVersionInfo struct {
	Item []*InstanceVersionInfo `xml:"item,omitempty"`
}

type GCInfo struct {
	Processname string `xml:"processname,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Reason string `xml:"reason,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	Duration int32 `xml:"duration,omitempty"`

	CpuTime int32 `xml:"cpuTime,omitempty"`

	ObjBytesBefore int64 `xml:"objBytesBefore,omitempty"`

	ObjBytesAfter int64 `xml:"objBytesAfter,omitempty"`

	ObjBytesFreed int64 `xml:"objBytesFreed,omitempty"`

	ClsBytesBefore int64 `xml:"clsBytesBefore,omitempty"`

	ClsBytesAfter int64 `xml:"clsBytesAfter,omitempty"`

	ClsBytesFreed int64 `xml:"clsBytesFreed,omitempty"`

	HeapSize int64 `xml:"heapSize,omitempty"`

	UnloadedClasses int32 `xml:"unloadedClasses,omitempty"`
}

type ArrayOfGCInfo struct {
	Item []*GCInfo `xml:"item,omitempty"`
}

type GCInfo2 struct {
	Processname string `xml:"processname,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Reason string `xml:"reason,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	Duration int32 `xml:"duration,omitempty"`

	CpuTime int32 `xml:"cpuTime,omitempty"`

	ObjBytesBefore int64 `xml:"objBytesBefore,omitempty"`

	ObjBytesAfter int64 `xml:"objBytesAfter,omitempty"`

	ObjBytesFreed int64 `xml:"objBytesFreed,omitempty"`

	ClsBytesBefore int64 `xml:"clsBytesBefore,omitempty"`

	ClsBytesAfter int64 `xml:"clsBytesAfter,omitempty"`

	ClsBytesFreed int64 `xml:"clsBytesFreed,omitempty"`

	HeapSize int64 `xml:"heapSize,omitempty"`

	UnloadedClasses int32 `xml:"unloadedClasses,omitempty"`

	PageFaults int64 `xml:"pageFaults,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`
}

type ArrayOfGCInfo2 struct {
	Item []*GCInfo2 `xml:"item,omitempty"`
}

type HeapInfo struct {
	Processname string `xml:"processname,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Size int64 `xml:"size,omitempty"`

	CommitSize int64 `xml:"commitSize,omitempty"`

	MaxUsedSize int64 `xml:"maxUsedSize,omitempty"`

	InitialSize int64 `xml:"initialSize,omitempty"`

	MaxSize int64 `xml:"maxSize,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`
}

type ArrayOfHeapInfo struct {
	Item []*HeapInfo `xml:"item,omitempty"`
}

type InstanceProperty struct {
	Property string `xml:"property,omitempty"`

	Propertytype string `xml:"propertytype,omitempty"`

	Value string `xml:"value,omitempty"`
}

type ArrayOfInstanceProperties struct {
	Item []*InstanceProperty `xml:"item,omitempty"`
}

type LogFile struct {
	Filename string `xml:"filename,omitempty"`

	Size int64 `xml:"size,omitempty"`

	Modtime string `xml:"modtime,omitempty"`

	Format string `xml:"format,omitempty"`
}

type ArrayOfLogFile struct {
	Item []*LogFile `xml:"item,omitempty"`
}

type AccessPoint struct {
	Address string `xml:"address,omitempty"`

	Port int32 `xml:"port,omitempty"`

	Protocol string `xml:"protocol,omitempty"`

	Processname string `xml:"processname,omitempty"`

	Active string `xml:"active,omitempty"`
}

type ArrayOfAccessPoint struct {
	Item []*AccessPoint `xml:"item,omitempty"`
}

type SAPInstance struct {
	Hostname string `xml:"hostname,omitempty"`

	InstanceNr int32 `xml:"instanceNr,omitempty"`

	HttpPort int32 `xml:"httpPort,omitempty"`

	HttpsPort int32 `xml:"httpsPort,omitempty"`

	StartPriority string `xml:"startPriority,omitempty"`

	Features string `xml:"features,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`
}

type ArrayOfSAPInstance struct {
	Item []*SAPInstance `xml:"item,omitempty"`
}

type UpdateInstance struct {
	Hostname string `xml:"hostname,omitempty"`

	InstanceNr int32 `xml:"instanceNr,omitempty"`

	Status string `xml:"status,omitempty"`

	Starttime string `xml:"starttime,omitempty"`

	Endtime string `xml:"endtime,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`
}

type ArrayOfUpdateInstance struct {
	Item []*UpdateInstance `xml:"item,omitempty"`
}

type J2EEEJBSession struct {
	IdHash int32 `xml:"IdHash,omitempty"`

	State string `xml:"state,omitempty"`

	Size int32 `xml:"size,omitempty"`

	ActiveRequests int32 `xml:"activeRequests,omitempty"`

	TotalRequests int32 `xml:"totalRequests,omitempty"`

	BackingStore string `xml:"backingStore,omitempty"`

	Processname string `xml:"processname,omitempty"`

	StartTime string `xml:"startTime,omitempty"`

	UpdateTime string `xml:"updateTime,omitempty"`

	ResponseTime int32 `xml:"responseTime,omitempty"`

	User string `xml:"user,omitempty"`

	Transaction string `xml:"transaction,omitempty"`

	Ejb string `xml:"ejb,omitempty"`

	Application string `xml:"application,omitempty"`

	Reference string `xml:"reference,omitempty"`
}

type ArrayOfJ2EEEJBSession struct {
	Item []*J2EEEJBSession `xml:"item,omitempty"`
}

type J2EERemoteObject struct {
	IdHash int32 `xml:"IdHash,omitempty"`

	Address string `xml:"address,omitempty"`

	Port int32 `xml:"port,omitempty"`

	Protocol string `xml:"protocol,omitempty"`

	Direction string `xml:"direction,omitempty"`

	Stubs int32 `xml:"stubs,omitempty"`

	Implementations int32 `xml:"implementations,omitempty"`

	CreationTime string `xml:"creationTime,omitempty"`

	UpdateTime string `xml:"updateTime,omitempty"`

	Processname string `xml:"processname,omitempty"`
}

type ArrayOfJ2EERemoteObject struct {
	Item []*J2EERemoteObject `xml:"item,omitempty"`
}

type J2EEClusterMsg struct {
	Service string `xml:"service,omitempty"`

	Id string `xml:"id,omitempty"`

	Count int64 `xml:"count,omitempty"`

	Length int64 `xml:"length,omitempty"`

	Avglength int64 `xml:"avg-length,omitempty"`

	Maxlength int64 `xml:"max-length,omitempty"`

	Countp2pmsg int64 `xml:"count-p2p-msg,omitempty"`

	Countp2prequest int64 `xml:"count-p2p-request,omitempty"`

	Countp2preply int64 `xml:"count-p2p-reply,omitempty"`

	Countbroadcastmsg int64 `xml:"count-broadcast-msg,omitempty"`

	Countbroadcastrequest int64 `xml:"count-broadcast-request,omitempty"`
}

type ArrayOfJ2EEClusterMsg struct {
	Item []*J2EEClusterMsg `xml:"item,omitempty"`
}

type J2EESharedTableInfo struct {
	Table string `xml:"table,omitempty"`

	Used int32 `xml:"used,omitempty"`

	Peak int32 `xml:"peak,omitempty"`

	Limit int32 `xml:"limit,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`
}

type ArrayOfJ2EESharedTableInfo struct {
	Item []*J2EESharedTableInfo `xml:"item,omitempty"`
}

type J2EEComponentInfo struct {
	Type_ string `xml:"type,omitempty"`

	Name string `xml:"name,omitempty"`

	Startupmode string `xml:"startupmode,omitempty"`

	Status string `xml:"status,omitempty"`

	Expectedstatus string `xml:"expectedstatus,omitempty"`

	Details string `xml:"details,omitempty"`

	Dispstatus *STATECOLOR `xml:"dispstatus,omitempty"`
}

type ArrayOfJ2EEComponentInfo struct {
	Item []*J2EEComponentInfo `xml:"item,omitempty"`
}

type ICMThread struct {
	Name string `xml:"name,omitempty"`

	Id string `xml:"id,omitempty"`

	Requests int64 `xml:"requests,omitempty"`

	Status string `xml:"status,omitempty"`

	Requesttype string `xml:"requesttype,omitempty"`
}

type ArrayOfICMThread struct {
	Item []*ICMThread `xml:"item,omitempty"`
}

type ICMConnection struct {
	Conid string `xml:"conid,omitempty"`

	Protocol string `xml:"protocol,omitempty"`

	Role string `xml:"role,omitempty"`

	Requesttype string `xml:"requesttype,omitempty"`

	Peeraddress string `xml:"peer-address,omitempty"`

	Peerport int32 `xml:"peer-port,omitempty"`

	Localaddress string `xml:"local-address,omitempty"`

	Localport int32 `xml:"local-port,omitempty"`

	Proctimeout int32 `xml:"proc-timeout,omitempty"`

	Keepalivetimeout int32 `xml:"keepalive-timeout,omitempty"`

	Connectiontime string `xml:"connection-time,omitempty"`

	Nihdl int32 `xml:"nihdl,omitempty"`
}

type ArrayOfICMConnection struct {
	Item []*ICMConnection `xml:"item,omitempty"`
}

type ICMCacheEntry struct {
	Name string `xml:"name,omitempty"`

	Version int32 `xml:"version,omitempty"`

	Size int64 `xml:"size,omitempty"`

	Valid bool `xml:"valid,omitempty"`

	Cache string `xml:"cache,omitempty"`

	Creationtime string `xml:"creation-time,omitempty"`

	Lastaccesstime string `xml:"last-access-time,omitempty"`

	Expirationtime string `xml:"expiration-time,omitempty"`

	Cacheurl string `xml:"cacheurl,omitempty"`
}

type ArrayOfICMCacheEntry struct {
	Item []*ICMCacheEntry `xml:"item,omitempty"`
}

type ICMProxyConnection struct {
	Conid string `xml:"conid,omitempty"`

	Role string `xml:"role,omitempty"`

	Peeraddress string `xml:"peer-address,omitempty"`

	Peerport int32 `xml:"peer-port,omitempty"`

	Localaddress string `xml:"local-address,omitempty"`

	Localport int32 `xml:"local-port,omitempty"`

	Status string `xml:"status,omitempty"`

	Nihdl int32 `xml:"nihdl,omitempty"`

	Socket int32 `xml:"socket,omitempty"`

	Partner string `xml:"partner,omitempty"`

	Internalconvid string `xml:"internal-convid,omitempty"`

	Externalconvid string `xml:"external-convid,omitempty"`

	Snccontextlength int32 `xml:"snc-context-length,omitempty"`

	Failoverstatus string `xml:"failover-status,omitempty"`

	Disconnecttime string `xml:"disconnect-time,omitempty"`

	Objectid string `xml:"objectid,omitempty"`

	Tiduidmode string `xml:"tid-uid-mode,omitempty"`
}

type ArrayOfICMProxyConnection struct {
	Item []*ICMProxyConnection `xml:"item,omitempty"`
}

type WebDispServer struct {
	Sid string `xml:"sid,omitempty"`

	Instance string `xml:"instance,omitempty"`

	Hostname string `xml:"hostname,omitempty"`

	Protocol string `xml:"protocol,omitempty"`

	Type_ string `xml:"type,omitempty"`

	Status string `xml:"status,omitempty"`

	Capacity int32 `xml:"capacity,omitempty"`

	Load int32 `xml:"load,omitempty"`

	Port int32 `xml:"port,omitempty"`

	Curconn int32 `xml:"cur-conn,omitempty"`

	Peakconn int32 `xml:"peak-conn,omitempty"`

	Maxconn int32 `xml:"max-conn,omitempty"`

	Secport int32 `xml:"sec-port,omitempty"`

	Seccurconn int32 `xml:"sec-cur-conn,omitempty"`

	Secpeakconn int32 `xml:"sec-peak-conn,omitempty"`

	Secmaxconn int32 `xml:"sec-max-conn,omitempty"`

	Reqcntstateless int64 `xml:"req-cnt-stateless,omitempty"`

	Reqcntstateful int64 `xml:"req-cnt-stateful,omitempty"`

	Reqcntgroup int64 `xml:"req-cnt-group,omitempty"`

	Resptimemin int64 `xml:"resp-time-min,omitempty"`

	Resptimeavg int64 `xml:"resp-time-avg,omitempty"`

	Resptimelast int64 `xml:"resp-time-last,omitempty"`

	Pingtimelast int64 `xml:"ping-time-last,omitempty"`
}

type ArrayOfWebDispServer struct {
	Item []*WebDispServer `xml:"item,omitempty"`
}

type ParameterRestriction struct {
	Type_ *RESTRICTIONTYPE `xml:"type,omitempty"`

	Intmin int64 `xml:"int-min,omitempty"`

	Intmax int64 `xml:"int-max,omitempty"`

	Floatmin float64 `xml:"float-min,omitempty"`

	Floatmax float64 `xml:"float-max,omitempty"`

	Valuemap *ArrayOfString `xml:"valuemap,omitempty"`
}

type SubProfileParameter struct {
	Name string `xml:"name,omitempty"`

	Description string `xml:"description,omitempty"`

	Unit string `xml:"unit,omitempty"`

	Mandatory bool `xml:"mandatory,omitempty"`

	Restriction *ParameterRestriction `xml:"restriction,omitempty"`
}

type ArrayOfSubProfileParameter struct {
	Item []*SubProfileParameter `xml:"item,omitempty"`
}

type ProfileParameter struct {
	Name string `xml:"name,omitempty"`

	Group string `xml:"group,omitempty"`

	Description string `xml:"description,omitempty"`

	Unit string `xml:"unit,omitempty"`

	Restriction *ParameterRestriction `xml:"restriction,omitempty"`

	Subpara *ArrayOfSubProfileParameter `xml:"sub-para,omitempty"`

	Value string `xml:"value,omitempty"`

	Dynamicvalue string `xml:"dynamic-value,omitempty"`

	Values *ArrayOfString `xml:"values,omitempty"`

	Dynamicvalues *ArrayOfString `xml:"dynamic-values,omitempty"`
}

type ArrayOfProfileParameter struct {
	Item []*ProfileParameter `xml:"item,omitempty"`
}

type SetProfileParameter struct {
	Name string `xml:"name,omitempty"`

	Value string `xml:"value,omitempty"`

	Values *ArrayOfString `xml:"values,omitempty"`
}

type ArrayOfSetProfileParameter struct {
	Item []*SetProfileParameter `xml:"item,omitempty"`
}

type EnqLock struct {
	Lockname string `xml:"lock-name,omitempty"`

	Lockarg string `xml:"lock-arg,omitempty"`

	Lockmode string `xml:"lock-mode,omitempty"`

	Owner string `xml:"owner,omitempty"`

	Ownervb string `xml:"owner-vb,omitempty"`

	Usecountowner int32 `xml:"use-count-owner,omitempty"`

	Usecountownervb int32 `xml:"use-count-owner-vb,omitempty"`

	Client string `xml:"client,omitempty"`

	User string `xml:"user,omitempty"`

	Transaction string `xml:"transaction,omitempty"`

	Object string `xml:"object,omitempty"`

	Backup bool `xml:"backup,omitempty"`
}

type ArrayOfEnqLock struct {
	Item []*EnqLock `xml:"item,omitempty"`
}

type SnapshotInfo struct {
	Filename string `xml:"filename,omitempty"`

	Size int64 `xml:"size,omitempty"`

	Modtime string `xml:"modtime,omitempty"`

	Description string `xml:"description,omitempty"`
}

type ArrayOfSnapshotInfo struct {
	Item []*SnapshotInfo `xml:"item,omitempty"`
}

type HACheck struct {
	State *HAVerificationState `xml:"state,omitempty"`

	Category *HACheckCategory `xml:"category,omitempty"`

	Description string `xml:"description,omitempty"`

	Comment string `xml:"comment,omitempty"`
}

type ArrayOfHACheck struct {
	Item []*HACheck `xml:"item,omitempty"`
}

type ABAPComponent struct {
	Component string `xml:"component,omitempty"`

	Release string `xml:"release,omitempty"`

	Patchlevel string `xml:"patchlevel,omitempty"`

	Componenttype string `xml:"componenttype,omitempty"`

	Description string `xml:"description,omitempty"`
}

type ArrayOfABAPComponent struct {
	Item []*ABAPComponent `xml:"item,omitempty"`
}

type SAPControlPortType interface {

	/* Triggers an instance start and returns immediately. */
	Start(request *Start) (*StartResponse, error)

	/* Triggers a start of another instance of the system and returns immediately. */
	InstanceStart(request *InstanceStart) (*InstanceStartResponse, error)

	/* Executes bootstrap of actual or other instance of the system, webservice restarts after return. */
	Bootstrap(request *Bootstrap) (*BootstrapResponse, error)

	/* Triggers an instance stop and returns immediately. */
	Stop(request *Stop) (*StopResponse, error)

	/* Triggers a stop of another instance of the system and returns immediately. */
	InstanceStop(request *InstanceStop) (*InstanceStopResponse, error)

	/* Triggers a graceful instance stop and returns immediately. */
	Shutdown(request *Shutdown) (*ShutdownResponse, error)

	/* Returns a SAP profile parameter value for a given profile parameter. If the given profile parameter is empty it returns a string with all known parameter value pairs separated by newline. */
	ParameterValue(request *ParameterValue) (*ParameterValueResponse, error)

	/* Returns a list of all processes directly started by the webservice according to the SAP start profile. */
	GetProcessList(request *GetProcessList) (*GetProcessListResponse, error)

	/* Returns start profile name and content as array of lines. */
	GetStartProfile(request *GetStartProfile) (*GetStartProfileResponse, error)

	/* Returns webservice trace file name and content as array of lines. */
	GetTraceFile(request *GetTraceFile) (*GetTraceFileResponse, error)

	/* Returns CCMS Alert tree as an array, parent-child node relationship is encoded via the parent index of each node (similar to rz20 transaction). */
	GetAlertTree(request *GetAlertTree) (*GetAlertTreeResponse, error)

	/* Returns a list of all CCMS alerts for a given node and it's child nodes. */
	GetAlerts(request *GetAlerts) (*GetAlertsResponse, error)

	/* Triggers webservice restart without any effects on the eventually running instance and returns immediately. */
	RestartService(request *RestartService) (*RestartServiceResponse, error)

	/* Triggers webservice stop without any effects on the eventually running instance and returns immediately. After that the webservice needs to started in order to use the interface again. */
	StopService(request *StopService) (*StopServiceResponse, error)

	/* Returns the process environment as an array of strings. */
	GetEnvironment(request *GetEnvironment) (*GetEnvironmentResponse, error)

	/* Returns a list of all instance trace files (superseded by ListLogFiles). */
	ListDeveloperTraces(request *ListDeveloperTraces) (*ListDeveloperTracesResponse, error)

	/* Returns the content of a given trace file. Use size=0 to read the entire file, size bigger 0 to read the first size bytes, size smaller 0 to read the last size bytes (superseded by ReadLogFile). */
	ReadDeveloperTrace(request *ReadDeveloperTrace) (*ReadDeveloperTraceResponse, error)

	/* Triggers an instance restart and returns immediately. */
	RestartInstance(request *RestartInstance) (*RestartInstanceResponse, error)

	/* Sends a given OS signal to a given process. The signal can be given by name or number. */
	SendSignal(request *SendSignal) (*SendSignalResponse, error)

	/* Returns a list version information for the most important files of the instance */
	GetVersionInfo(request *GetVersionInfo) (*GetVersionInfoResponse, error)

	/* Returns a list of queue information of work processes and icm (similar to dpmon). */
	GetQueueStatistic(request *GetQueueStatistic) (*GetQueueStatisticResponse, error)

	/* Returns a list of available instance features and information how to get it. */
	GetInstanceProperties(request *GetInstanceProperties) (*GetInstancePropertiesResponse, error)

	/* Executes an external OS command. */
	OSExecute(request *OSExecute) (*OSExecuteResponse, error)

	/* Returns the content of a given log file. Use statecookie to specify starting position (""=beginning), use maxentries to specify upper limit of returned entries (0=all) and reading direction (positive: forward, negative: backward). */
	ReadLogFile(request *ReadLogFile) (*ReadLogFileResponse, error)

	/* Returns a list of all instance log files. */
	ListLogFiles(request *ListLogFiles) (*ListLogFilesResponse, error)

	/* Service definition of function SAPControl__AnalyseLogFiles */
	AnalyseLogFiles(request *AnalyseLogFiles) (*AnalyseLogFilesResponse, error)

	/* Returns a list of instance network access points. */
	GetAccessPointList(request *GetAccessPointList) (*GetAccessPointListResponse, error)

	/* Returns a list of SAP instances of the SAP system. */
	GetSystemInstanceList(request *GetSystemInstanceList) (*GetSystemInstanceListResponse, error)

	/* Returns a list of instances restarting by ongoing UpdateSystem execution. */
	GetSystemUpdateList(request *GetSystemUpdateList) (*GetSystemUpdateListResponse, error)

	/* Triggers start of entire system or parts of it. */
	StartSystem(request *StartSystem) (*StartSystemResponse, error)

	/* Triggers stop or soft shutdown of entire system or parts of it. */
	StopSystem(request *StopSystem) (*StopSystemResponse, error)

	/* Triggers restart of entire system or parts of it. */
	RestartSystem(request *RestartSystem) (*RestartSystemResponse, error)

	/* Triggers sequential restart of all instance to activate changes. */
	UpdateSystem(request *UpdateSystem) (*UpdateSystemResponse, error)

	/* Service definition of function SAPControl__UpdateSCSInstance */
	UpdateSCSInstance(request *UpdateSCSInstance) (*UpdateSCSInstanceResponse, error)

	/* Checks prerequisites for executing UpdateSystem. */
	CheckUpdateSystem(request *CheckUpdateSystem) (*CheckUpdateSystemResponse, error)

	/* Checks execute permission of webservice function. */
	AccessCheck(request *AccessCheck) (*AccessCheckResponse, error)

	/* Returns actual Profile Parameters for a given process. */
	GetProcessParameter(request *GetProcessParameter) (*GetProcessParameterResponse, error)

	/* Sets dynamic Profile Parameters for a given process. */
	SetProcessParameter(request *SetProcessParameter) (*SetProcessParameterResponse, error)

	/* Sets dynamic and persistent Profile Parameters for a given process. */
	SetProcessParameter2(request *SetProcessParameter2) (*SetProcessParameter2Response, error)

	/* Triggers sapstarstrv to detach from shared memory (for internal usage only). */
	ShmDetach(request *ShmDetach) (*ShmDetachResponse, error)

	/* Returns a unique network ID and verification proof for a given service. */
	GetSecNetworkId(request *GetSecNetworkId) (*GetSecNetworkIdResponse, error)

	/* Returns a unique network ID for a given service. */
	GetNetworkId(request *GetNetworkId) (*GetNetworkIdResponse, error)

	/* Creates a zipped instance snapshot on the server side. */
	CreateSnapshot(request *CreateSnapshot) (*CreateSnapshotResponse, error)

	/* Downloads a zipped snapshot from the server. */
	ReadSnapshot(request *ReadSnapshot) (*ReadSnapshotResponse, error)

	/* Returns a list of all available snapshots on the server side. */
	ListSnapshots(request *ListSnapshots) (*ListSnapshotsResponse, error)

	/* Deletes snapshot(s) on the server side. */
	DeleteSnapshots(request *DeleteSnapshots) (*DeleteSnapshotsResponse, error)

	/* Reads the SAP ABAP Syslog and returns it as an array of entries (similar to sm21 transaction). */
	ABAPReadSyslog(request *ABAPReadSyslog) (*ABAPReadSyslogResponse, error)

	/* Reads the SAP ABAP Syslog and returns it as an array of raw entries (similar to sm21 transaction). */
	ABAPReadRawSyslog(request *ABAPReadRawSyslog) (*ABAPReadRawSyslogResponse, error)

	/* Returns a list of the ABAP workprocesses (similar to sm50 transaction). */
	ABAPGetWPTable(request *ABAPGetWPTable) (*ABAPGetWPTableResponse, error)

	/* Acknowledge CCMS Alerts in the SAP ABAP system. Requires SAP user credentials and a list of alert ids to acknowledge. Returns a list of success code for each alerts (1=success, 0=failure). */
	ABAPAcknowledgeAlerts(request *ABAPAcknowledgeAlerts) (*ABAPAcknowledgeAlertsResponse, error)

	/* Returns a list of installed ABAP components in the system as defined in CVERS database table. */
	ABAPGetComponentList(request *ABAPGetComponentList) (*ABAPGetComponentListResponse, error)

	/* Returns a list of single point of failure RFC destination definitions. */
	ABAPCheckRFCDestinations(request *ABAPCheckRFCDestinations) (*ABAPCheckRFCDestinationsResponse, error)

	/* Returns a list of all or all active ABAP workprocesses in the system (similar to sm66 transaction). */
	ABAPGetSystemWPTable(request *ABAPGetSystemWPTable) (*ABAPGetSystemWPTableResponse, error)

	/* Internal usage only. */
	RequestTicket(request *RequestTicket) (*RequestTicketResponse, error)

	/* Internal usage only. */
	SendTicket(request *SendTicket) (*SendTicketResponse, error)

	/* Configures a list of logfiles accessable via ReadLogFile (Hostagent mode only). */
	ConfigureLogFileList(request *ConfigureLogFileList) (*ConfigureLogFileListResponse, error)

	/* Returns a list of configured logfiles accessable via ReadLogFile (Hostagent mode only). */
	GetLogFileList(request *GetLogFileList) (*GetLogFileListResponse, error)

	/* Returns a logon file for local authentication. */
	RequestLogonFile(request *RequestLogonFile) (*RequestLogonFileResponse, error)

	/* Updates the system internal private key infrastructure (system root PSE and instance PSE(s)) of the entire system if necessary. Use force flag to enforce update.  */
	UpdateSystemPKI(request *UpdateSystemPKI) (*UpdateSystemPKIResponse, error)

	/* Updates instance PSE of system internal private key infrastructure if necessary. Use force flag to enforce update. */
	UpdateInstancePSE(request *UpdateInstancePSE) (*UpdateInstancePSEResponse, error)

	/* Checks high availability configurration and status of the system. */
	HACheckConfig(request *HACheckConfig) (*HACheckConfigResponse, error)

	/* Checks HA failover third party configuration and status of an instnace. */
	HACheckFailoverConfig(request *HACheckFailoverConfig) (*HACheckFailoverConfigResponse, error)

	/* Returns HA failover third party information. */
	HAGetFailoverConfig(request *HAGetFailoverConfig) (*HAGetFailoverConfigResponse, error)

	/* Moves instance to given HA cluster node. */
	HAFailoverToNode(request *HAFailoverToNode) (*HAFailoverToNodeResponse, error)

	/* Triggers an instance start without checking HA product integration and returns immediately (for HA product internal use only). */
	StartBypassHA(request *StartBypassHA) (*StartBypassHAResponse, error)

	/* Triggers an instance stop without checking HA product integration and returns immediately (for HA product internal use only). */
	StopBypassHA(request *StopBypassHA) (*StopBypassHAResponse, error)

	/* Returns callstack of all threads of specified process. */
	GetCallstack(request *GetCallstack) (*GetCallstackResponse, error)

	/* Returns a list of J2EE processes controlled by jcontrol (superseded by J2EEGetProcessList2). */
	J2EEGetProcessList(request *J2EEGetProcessList) (*J2EEGetProcessListResponse, error)

	/* Returns a list of J2EE processes controlled by jcontrol (supersedes J2EEGetProcessList). */
	J2EEGetProcessList2(request *J2EEGetProcessList2) (*J2EEGetProcessList2Response, error)

	/* Performs a given function (StartInstance, StopInstance, RestartInstance, BootInstance, RebootInstance) on the J2EE instance or (EnableProcess, StartProcess, DisableProcess, StopProcess, SoftStopProcess, ActivateProcess, DeactivateProcess, RestartProcess, SoftRestartProcess, DumpStackTrace, EnableDebugging, DisableDebugging, IncrementTrace, DecrementTrace) on a given J2EE process. */
	J2EEControlProcess(request *J2EEControlProcess) (*J2EEControlProcessResponse, error)

	/* Service definition of function SAPControl__J2EEControlCluster */
	J2EEControlCluster(request *J2EEControlCluster) (*J2EEControlClusterResponse, error)

	/* Returns a list of threads in the J2EE instance (superseded by J2EEGetThreadList2). */
	J2EEGetThreadList(request *J2EEGetThreadList) (*J2EEGetThreadListResponse, error)

	/* Returns a list of threads in the J2EE instance (supersedes J2EEGetThreadList). */
	J2EEGetThreadList2(request *J2EEGetThreadList2) (*J2EEGetThreadList2Response, error)

	/* Returns a list of (HTTP) sessions in the J2EE instance (superseded by J2EEGetWebSessionList, J2EEGetWebSessionList2). */
	J2EEGetSessionList(request *J2EEGetSessionList) (*J2EEGetSessionListResponse, error)

	/* Returns a list of (HTTP) sessions in the J2EE instance (supersedes J2EEGetSessionList, superseded by J2EEGetWebSessionList2). */
	J2EEGetWebSessionList(request *J2EEGetWebSessionList) (*J2EEGetWebSessionListResponse, error)

	/* Returns a list of (HTTP) sessions in the J2EE instance (supersedes J2EEGetSessionList, J2EEGetWebSessionList). */
	J2EEGetWebSessionList2(request *J2EEGetWebSessionList2) (*J2EEGetWebSessionList2Response, error)

	/* Returns a list of caches in the J2EE instance (superseded by J2EEGetCacheStatistic2). */
	J2EEGetCacheStatistic(request *J2EEGetCacheStatistic) (*J2EEGetCacheStatisticResponse, error)

	/* Returns a list of caches in the J2EE instance (supersedes J2EEGetCacheStatistic). */
	J2EEGetCacheStatistic2(request *J2EEGetCacheStatistic2) (*J2EEGetCacheStatistic2Response, error)

	/* Returns a list of application aliases in the J2EE instance. */
	J2EEGetApplicationAliasList(request *J2EEGetApplicationAliasList) (*J2EEGetApplicationAliasListResponse, error)

	/* Returns a list of JAVA VM garbage collections in the J2EE instance (superseded by J2EEGetVMGCHistory2). */
	J2EEGetVMGCHistory(request *J2EEGetVMGCHistory) (*J2EEGetVMGCHistoryResponse, error)

	/* Returns a list of JAVA VM garbage collections in the J2EE instance (supersedes J2EEGetVMGCHistory). */
	J2EEGetVMGCHistory2(request *J2EEGetVMGCHistory2) (*J2EEGetVMGCHistory2Response, error)

	/* Returns a list of JAVA VM heap information. */
	J2EEGetVMHeapInfo(request *J2EEGetVMHeapInfo) (*J2EEGetVMHeapInfoResponse, error)

	/* Returns a list of EJB sessions in the J2EE instance. */
	J2EEGetEJBSessionList(request *J2EEGetEJBSessionList) (*J2EEGetEJBSessionListResponse, error)

	/* Returns a list of remote object connections in the J2EE instance. */
	J2EEGetRemoteObjectList(request *J2EEGetRemoteObjectList) (*J2EEGetRemoteObjectListResponse, error)

	/* Returns a list of J2EE cluster communication statistic from the message server. */
	J2EEGetClusterMsgList(request *J2EEGetClusterMsgList) (*J2EEGetClusterMsgListResponse, error)

	/* Returns a list of SAP startup framework shared memory table information. */
	J2EEGetSharedTableInfo(request *J2EEGetSharedTableInfo) (*J2EEGetSharedTableInfoResponse, error)

	/* Creates a J2EE debug session on  given node or automatically selected node (processname = "") with the given debug flags (blank separated keywords "SuspendAll", "CodeIsolate", "LoadIsolate", "MigrateSessions") for the given client name. Returns debug key an network port. */
	J2EEEnableDbgSession(request *J2EEEnableDbgSession) (*J2EEEnableDbgSessionResponse, error)

	/* Removes a J2EE debug session created by J2EEEnableDbgSession. */
	J2EEDisableDbgSession(request *J2EEDisableDbgSession) (*J2EEDisableDbgSessionResponse, error)

	/* Returns java callstack of specified J2EE thread or all J2EE threads (index = -1). */
	J2EEGetThreadCallStack(request *J2EEGetThreadCallStack) (*J2EEGetThreadCallStackResponse, error)

	/* Returns java taskstack of specified J2EE thread or all J2EE threads (index = -1). */
	J2EEGetThreadTaskStack(request *J2EEGetThreadTaskStack) (*J2EEGetThreadTaskStackResponse, error)

	/* Returns a list of J2EE services and applications. */
	J2EEGetComponentList(request *J2EEGetComponentList) (*J2EEGetComponentListResponse, error)

	/* Performs a given function (start/stop/restart) on the given J2EE components on the given J2EE process. */
	J2EEControlComponents(request *J2EEControlComponents) (*J2EEControlComponentsResponse, error)

	/* Returns a list of threads used by ICM. */
	ICMGetThreadList(request *ICMGetThreadList) (*ICMGetThreadListResponse, error)

	/* Returns a list of incoming network connections handled by ICM. */
	ICMGetConnectionList(request *ICMGetConnectionList) (*ICMGetConnectionListResponse, error)

	/* Returns a list of objects cached by ICM. */
	ICMGetCacheEntries(request *ICMGetCacheEntries) (*ICMGetCacheEntriesResponse, error)

	/* Returns a list of outgoing network proxy connections handled by ICM. */
	ICMGetProxyConnectionList(request *ICMGetProxyConnectionList) (*ICMGetProxyConnectionListResponse, error)

	/* Returns a list backend servers connected to the webdispatcher. */
	WebDispGetServerList(request *WebDispGetServerList) (*WebDispGetServerListResponse, error)

	/* Returns enque locking table. */
	EnqGetLockTable(request *EnqGetLockTable) (*EnqGetLockTableResponse, error)

	/* Removes locks from enque locking table. */
	EnqRemoveLocks(request *EnqRemoveLocks) (*EnqRemoveLocksResponse, error)

	/* Returns enque statistic. */
	EnqGetStatistic(request *EnqGetStatistic) (*EnqStatistic, error)
}

type sAPControlPortType struct {
	client *soap.Client
}

func NewSAPControlPortType(client *soap.Client) SAPControlPortType {
	return &sAPControlPortType{
		client: client,
	}
}

func (service *sAPControlPortType) Start(request *Start) (*StartResponse, error) {
	response := new(StartResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) InstanceStart(request *InstanceStart) (*InstanceStartResponse, error) {
	response := new(InstanceStartResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) Bootstrap(request *Bootstrap) (*BootstrapResponse, error) {
	response := new(BootstrapResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) Stop(request *Stop) (*StopResponse, error) {
	response := new(StopResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) InstanceStop(request *InstanceStop) (*InstanceStopResponse, error) {
	response := new(InstanceStopResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) Shutdown(request *Shutdown) (*ShutdownResponse, error) {
	response := new(ShutdownResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ParameterValue(request *ParameterValue) (*ParameterValueResponse, error) {
	response := new(ParameterValueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetProcessList(request *GetProcessList) (*GetProcessListResponse, error) {
	response := new(GetProcessListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetStartProfile(request *GetStartProfile) (*GetStartProfileResponse, error) {
	response := new(GetStartProfileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetTraceFile(request *GetTraceFile) (*GetTraceFileResponse, error) {
	response := new(GetTraceFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetAlertTree(request *GetAlertTree) (*GetAlertTreeResponse, error) {
	response := new(GetAlertTreeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetAlerts(request *GetAlerts) (*GetAlertsResponse, error) {
	response := new(GetAlertsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) RestartService(request *RestartService) (*RestartServiceResponse, error) {
	response := new(RestartServiceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) StopService(request *StopService) (*StopServiceResponse, error) {
	response := new(StopServiceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetEnvironment(request *GetEnvironment) (*GetEnvironmentResponse, error) {
	response := new(GetEnvironmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ListDeveloperTraces(request *ListDeveloperTraces) (*ListDeveloperTracesResponse, error) {
	response := new(ListDeveloperTracesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ReadDeveloperTrace(request *ReadDeveloperTrace) (*ReadDeveloperTraceResponse, error) {
	response := new(ReadDeveloperTraceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) RestartInstance(request *RestartInstance) (*RestartInstanceResponse, error) {
	response := new(RestartInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) SendSignal(request *SendSignal) (*SendSignalResponse, error) {
	response := new(SendSignalResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetVersionInfo(request *GetVersionInfo) (*GetVersionInfoResponse, error) {
	response := new(GetVersionInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetQueueStatistic(request *GetQueueStatistic) (*GetQueueStatisticResponse, error) {
	response := new(GetQueueStatisticResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetInstanceProperties(request *GetInstanceProperties) (*GetInstancePropertiesResponse, error) {
	response := new(GetInstancePropertiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) OSExecute(request *OSExecute) (*OSExecuteResponse, error) {
	response := new(OSExecuteResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ReadLogFile(request *ReadLogFile) (*ReadLogFileResponse, error) {
	response := new(ReadLogFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ListLogFiles(request *ListLogFiles) (*ListLogFilesResponse, error) {
	response := new(ListLogFilesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) AnalyseLogFiles(request *AnalyseLogFiles) (*AnalyseLogFilesResponse, error) {
	response := new(AnalyseLogFilesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetAccessPointList(request *GetAccessPointList) (*GetAccessPointListResponse, error) {
	response := new(GetAccessPointListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetSystemInstanceList(request *GetSystemInstanceList) (*GetSystemInstanceListResponse, error) {
	response := new(GetSystemInstanceListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetSystemUpdateList(request *GetSystemUpdateList) (*GetSystemUpdateListResponse, error) {
	response := new(GetSystemUpdateListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) StartSystem(request *StartSystem) (*StartSystemResponse, error) {
	response := new(StartSystemResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) StopSystem(request *StopSystem) (*StopSystemResponse, error) {
	response := new(StopSystemResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) RestartSystem(request *RestartSystem) (*RestartSystemResponse, error) {
	response := new(RestartSystemResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) UpdateSystem(request *UpdateSystem) (*UpdateSystemResponse, error) {
	response := new(UpdateSystemResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) UpdateSCSInstance(request *UpdateSCSInstance) (*UpdateSCSInstanceResponse, error) {
	response := new(UpdateSCSInstanceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) CheckUpdateSystem(request *CheckUpdateSystem) (*CheckUpdateSystemResponse, error) {
	response := new(CheckUpdateSystemResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) AccessCheck(request *AccessCheck) (*AccessCheckResponse, error) {
	response := new(AccessCheckResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetProcessParameter(request *GetProcessParameter) (*GetProcessParameterResponse, error) {
	response := new(GetProcessParameterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) SetProcessParameter(request *SetProcessParameter) (*SetProcessParameterResponse, error) {
	response := new(SetProcessParameterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) SetProcessParameter2(request *SetProcessParameter2) (*SetProcessParameter2Response, error) {
	response := new(SetProcessParameter2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ShmDetach(request *ShmDetach) (*ShmDetachResponse, error) {
	response := new(ShmDetachResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetSecNetworkId(request *GetSecNetworkId) (*GetSecNetworkIdResponse, error) {
	response := new(GetSecNetworkIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetNetworkId(request *GetNetworkId) (*GetNetworkIdResponse, error) {
	response := new(GetNetworkIdResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) CreateSnapshot(request *CreateSnapshot) (*CreateSnapshotResponse, error) {
	response := new(CreateSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ReadSnapshot(request *ReadSnapshot) (*ReadSnapshotResponse, error) {
	response := new(ReadSnapshotResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ListSnapshots(request *ListSnapshots) (*ListSnapshotsResponse, error) {
	response := new(ListSnapshotsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) DeleteSnapshots(request *DeleteSnapshots) (*DeleteSnapshotsResponse, error) {
	response := new(DeleteSnapshotsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ABAPReadSyslog(request *ABAPReadSyslog) (*ABAPReadSyslogResponse, error) {
	response := new(ABAPReadSyslogResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ABAPReadRawSyslog(request *ABAPReadRawSyslog) (*ABAPReadRawSyslogResponse, error) {
	response := new(ABAPReadRawSyslogResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ABAPGetWPTable(request *ABAPGetWPTable) (*ABAPGetWPTableResponse, error) {
	response := new(ABAPGetWPTableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ABAPAcknowledgeAlerts(request *ABAPAcknowledgeAlerts) (*ABAPAcknowledgeAlertsResponse, error) {
	response := new(ABAPAcknowledgeAlertsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ABAPGetComponentList(request *ABAPGetComponentList) (*ABAPGetComponentListResponse, error) {
	response := new(ABAPGetComponentListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ABAPCheckRFCDestinations(request *ABAPCheckRFCDestinations) (*ABAPCheckRFCDestinationsResponse, error) {
	response := new(ABAPCheckRFCDestinationsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ABAPGetSystemWPTable(request *ABAPGetSystemWPTable) (*ABAPGetSystemWPTableResponse, error) {
	response := new(ABAPGetSystemWPTableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) RequestTicket(request *RequestTicket) (*RequestTicketResponse, error) {
	response := new(RequestTicketResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) SendTicket(request *SendTicket) (*SendTicketResponse, error) {
	response := new(SendTicketResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ConfigureLogFileList(request *ConfigureLogFileList) (*ConfigureLogFileListResponse, error) {
	response := new(ConfigureLogFileListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetLogFileList(request *GetLogFileList) (*GetLogFileListResponse, error) {
	response := new(GetLogFileListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) RequestLogonFile(request *RequestLogonFile) (*RequestLogonFileResponse, error) {
	response := new(RequestLogonFileResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) UpdateSystemPKI(request *UpdateSystemPKI) (*UpdateSystemPKIResponse, error) {
	response := new(UpdateSystemPKIResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) UpdateInstancePSE(request *UpdateInstancePSE) (*UpdateInstancePSEResponse, error) {
	response := new(UpdateInstancePSEResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) HACheckConfig(request *HACheckConfig) (*HACheckConfigResponse, error) {
	response := new(HACheckConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) HACheckFailoverConfig(request *HACheckFailoverConfig) (*HACheckFailoverConfigResponse, error) {
	response := new(HACheckFailoverConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) HAGetFailoverConfig(request *HAGetFailoverConfig) (*HAGetFailoverConfigResponse, error) {
	response := new(HAGetFailoverConfigResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) HAFailoverToNode(request *HAFailoverToNode) (*HAFailoverToNodeResponse, error) {
	response := new(HAFailoverToNodeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) StartBypassHA(request *StartBypassHA) (*StartBypassHAResponse, error) {
	response := new(StartBypassHAResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) StopBypassHA(request *StopBypassHA) (*StopBypassHAResponse, error) {
	response := new(StopBypassHAResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) GetCallstack(request *GetCallstack) (*GetCallstackResponse, error) {
	response := new(GetCallstackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetProcessList(request *J2EEGetProcessList) (*J2EEGetProcessListResponse, error) {
	response := new(J2EEGetProcessListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetProcessList2(request *J2EEGetProcessList2) (*J2EEGetProcessList2Response, error) {
	response := new(J2EEGetProcessList2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEControlProcess(request *J2EEControlProcess) (*J2EEControlProcessResponse, error) {
	response := new(J2EEControlProcessResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEControlCluster(request *J2EEControlCluster) (*J2EEControlClusterResponse, error) {
	response := new(J2EEControlClusterResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetThreadList(request *J2EEGetThreadList) (*J2EEGetThreadListResponse, error) {
	response := new(J2EEGetThreadListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetThreadList2(request *J2EEGetThreadList2) (*J2EEGetThreadList2Response, error) {
	response := new(J2EEGetThreadList2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetSessionList(request *J2EEGetSessionList) (*J2EEGetSessionListResponse, error) {
	response := new(J2EEGetSessionListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetWebSessionList(request *J2EEGetWebSessionList) (*J2EEGetWebSessionListResponse, error) {
	response := new(J2EEGetWebSessionListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetWebSessionList2(request *J2EEGetWebSessionList2) (*J2EEGetWebSessionList2Response, error) {
	response := new(J2EEGetWebSessionList2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetCacheStatistic(request *J2EEGetCacheStatistic) (*J2EEGetCacheStatisticResponse, error) {
	response := new(J2EEGetCacheStatisticResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetCacheStatistic2(request *J2EEGetCacheStatistic2) (*J2EEGetCacheStatistic2Response, error) {
	response := new(J2EEGetCacheStatistic2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetApplicationAliasList(request *J2EEGetApplicationAliasList) (*J2EEGetApplicationAliasListResponse, error) {
	response := new(J2EEGetApplicationAliasListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetVMGCHistory(request *J2EEGetVMGCHistory) (*J2EEGetVMGCHistoryResponse, error) {
	response := new(J2EEGetVMGCHistoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetVMGCHistory2(request *J2EEGetVMGCHistory2) (*J2EEGetVMGCHistory2Response, error) {
	response := new(J2EEGetVMGCHistory2Response)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetVMHeapInfo(request *J2EEGetVMHeapInfo) (*J2EEGetVMHeapInfoResponse, error) {
	response := new(J2EEGetVMHeapInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetEJBSessionList(request *J2EEGetEJBSessionList) (*J2EEGetEJBSessionListResponse, error) {
	response := new(J2EEGetEJBSessionListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetRemoteObjectList(request *J2EEGetRemoteObjectList) (*J2EEGetRemoteObjectListResponse, error) {
	response := new(J2EEGetRemoteObjectListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetClusterMsgList(request *J2EEGetClusterMsgList) (*J2EEGetClusterMsgListResponse, error) {
	response := new(J2EEGetClusterMsgListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetSharedTableInfo(request *J2EEGetSharedTableInfo) (*J2EEGetSharedTableInfoResponse, error) {
	response := new(J2EEGetSharedTableInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEEnableDbgSession(request *J2EEEnableDbgSession) (*J2EEEnableDbgSessionResponse, error) {
	response := new(J2EEEnableDbgSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEDisableDbgSession(request *J2EEDisableDbgSession) (*J2EEDisableDbgSessionResponse, error) {
	response := new(J2EEDisableDbgSessionResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetThreadCallStack(request *J2EEGetThreadCallStack) (*J2EEGetThreadCallStackResponse, error) {
	response := new(J2EEGetThreadCallStackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetThreadTaskStack(request *J2EEGetThreadTaskStack) (*J2EEGetThreadTaskStackResponse, error) {
	response := new(J2EEGetThreadTaskStackResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEGetComponentList(request *J2EEGetComponentList) (*J2EEGetComponentListResponse, error) {
	response := new(J2EEGetComponentListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) J2EEControlComponents(request *J2EEControlComponents) (*J2EEControlComponentsResponse, error) {
	response := new(J2EEControlComponentsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ICMGetThreadList(request *ICMGetThreadList) (*ICMGetThreadListResponse, error) {
	response := new(ICMGetThreadListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ICMGetConnectionList(request *ICMGetConnectionList) (*ICMGetConnectionListResponse, error) {
	response := new(ICMGetConnectionListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ICMGetCacheEntries(request *ICMGetCacheEntries) (*ICMGetCacheEntriesResponse, error) {
	response := new(ICMGetCacheEntriesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) ICMGetProxyConnectionList(request *ICMGetProxyConnectionList) (*ICMGetProxyConnectionListResponse, error) {
	response := new(ICMGetProxyConnectionListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) WebDispGetServerList(request *WebDispGetServerList) (*WebDispGetServerListResponse, error) {
	response := new(WebDispGetServerListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) EnqGetLockTable(request *EnqGetLockTable) (*EnqGetLockTableResponse, error) {
	response := new(EnqGetLockTableResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) EnqRemoveLocks(request *EnqRemoveLocks) (*EnqRemoveLocksResponse, error) {
	response := new(EnqRemoveLocksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *sAPControlPortType) EnqGetStatistic(request *EnqGetStatistic) (*EnqStatistic, error) {
	response := new(EnqStatistic)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
