go-hvkvp
========


## Prepare and send Key-Value pair
```powershell
$VmMgmt = Get-WmiObject -Namespace root\virtualization\v2 -Class Msvm_VirtualSystemManagementService  
$vm = Get-WmiObject -Namespace root\virtualization\v2 -Class Msvm_ComputerSystem -Filter {ElementName = 'minishift'}

$kvpDataItem = ([WMIClass][String]::Format("\\{0}\{1}:{2}", $VmMgmt.ClassPath.Server, $VmMgmt.ClassPath.NamespacePath, "Msvm_KvpExchangeDataItem")).CreateInstance()

$kvpDataItem.Name = "IpAddress"
$kvpDataItem.Data = "10.0.75.128"
$kvpDataItem.Source = 0

$VmMgmt.AddKvpItems($Vm, $kvpDataItem.PSBase.GetText(1))
```

## Receive/Read on the host:
```
$ go build main.go
$ ./main
```
