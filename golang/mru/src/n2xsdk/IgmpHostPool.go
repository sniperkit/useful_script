package n2xsdk

type IgmpHostPool struct {
 Handler string
}

func(np *IgmpHostPool) SetVersion () error {
 //parameters: DeviceHandle IgmpVersion
 //AgtIgmpHostPool SetVersion, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) GetVersion ()(string, error) {
 //parameters: DeviceHandle
 //AgtIgmpHostPool GetVersion
 return "", nil
}

func(np *IgmpHostPool) SetParameter () error {
 //parameters: DeviceHandle Parameter Value
 //AgtIgmpHostPool SetParameter, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) GetParameter ()(string, error) {
 //parameters: DeviceHandle Parameter
 //AgtIgmpHostPool GetParameter
 return "", nil
}

func(np *IgmpHostPool) GetInfo ()(string, error) {
 //parameters: SessionList Info
 //AgtIgmpHostPool GetInfo
 return "", nil
}

func(np *IgmpHostPool) AddGroupPools () error {
 //parameters: DeviceHandle GroupPoolHandles
 //AgtIgmpHostPool AddGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) RemoveGroupPools () error {
 //parameters: DeviceHandle GroupPoolHandles
 //AgtIgmpHostPool RemoveGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) RemoveAllGroupPools () error {
 //parameters: DeviceHandle
 //AgtIgmpHostPool RemoveAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) ListGroupPools ()(string, error) {
 //parameters: DeviceHandle
 //AgtIgmpHostPool ListGroupPools
 return "", nil
}

func(np *IgmpHostPool) AddSourcePools () error {
 //parameters: SessionList GroupPoolHandle SourcePoolHandles
 //AgtIgmpHostPool AddSourcePools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) RemoveSourcePools () error {
 //parameters: SessionList GroupPoolHandle SourcePoolHandles
 //AgtIgmpHostPool RemoveSourcePools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) RemoveAllSourcePools () error {
 //parameters: SessionList GroupPoolHandle
 //AgtIgmpHostPool RemoveAllSourcePools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) ListSourcePools ()(string, error) {
 //parameters: SessionList GroupPoolHandle
 //AgtIgmpHostPool ListSourcePools
 return "", nil
}

func(np *IgmpHostPool) SetGroupPoolFilter () error {
 //parameters: SessionList GroupPoolHandle IgmpFilter
 //AgtIgmpHostPool SetGroupPoolFilter, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) GetGroupPoolFilter ()(string, error) {
 //parameters: SessionList GroupPoolHandle
 //AgtIgmpHostPool GetGroupPoolFilter
 return "", nil
}

func(np *IgmpHostPool) SetGroupPoolState () error {
 //parameters: SessionList GroupPoolHandle IgmpFilter SourcePoolHandles
 //AgtIgmpHostPool SetGroupPoolState, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) GetGroupPoolState ()(string, error) {
 //parameters: SessionList GroupPoolHandle
 //AgtIgmpHostPool GetGroupPoolState
 return "", nil
}

func(np *IgmpHostPool) GetNumberOfSourcesInGroupPool ()(string, error) {
 //parameters: SessionList GroupPoolHandle
 //AgtIgmpHostPool GetNumberOfSourcesInGroupPool
 return "", nil
}

func(np *IgmpHostPool) ListGroupPoolsJoined ()(string, error) {
 //parameters: SessionList
 //AgtIgmpHostPool ListGroupPoolsJoined
 return "", nil
}

func(np *IgmpHostPool) IsJoined () error {
 //parameters: SessionList GroupPoolHandle
 //AgtIgmpHostPool IsJoined, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) JoinGroupPools () error {
 //parameters: SessionList GroupPoolHandles
 //AgtIgmpHostPool JoinGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) JoinAllGroupPools () error {
 //parameters: SessionList
 //AgtIgmpHostPool JoinAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) LeaveGroupPools () error {
 //parameters: SessionList GroupPoolHandles
 //AgtIgmpHostPool LeaveGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) LeaveAllGroupPools () error {
 //parameters: SessionList
 //AgtIgmpHostPool LeaveAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) StartPduSave () error {
 //parameters: SessionList
 //AgtIgmpHostPool StartPduSave, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) StopPduSave () error {
 //parameters: SessionList
 //AgtIgmpHostPool StopPduSave, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) ListSavedPdus ()(string, error) {
 //parameters: SessionList
 //AgtIgmpHostPool ListSavedPdus
 return "", nil
}

func(np *IgmpHostPool) GetPduDetails ()(string, error) {
 //parameters: SessionList PduId PduInfoType
 //AgtIgmpHostPool GetPduDetails
 return "", nil
}

func(np *IgmpHostPool) GetAllPduDetails ()(string, error) {
 //parameters: SessionList PduId
 //AgtIgmpHostPool GetAllPduDetails
 return "", nil
}

func(np *IgmpHostPool) DeletePdus () error {
 //parameters: SessionList
 //AgtIgmpHostPool DeletePdus, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) IsIpv4PriorityNoCodePointFieldSelected () error {
 //parameters: DeviceHandle
 //AgtIgmpHostPool IsIpv4PriorityNoCodePointFieldSelected, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) SelectIpv4PriorityNoCodePointField () error {
 //parameters: DeviceHandle
 //AgtIgmpHostPool SelectIpv4PriorityNoCodePointField, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) GetIpv4Priority ()(string, error) {
 //parameters: DeviceHandle
 //AgtIgmpHostPool GetIpv4Priority
 return "", nil
}

func(np *IgmpHostPool) SetIpv4Priority () error {
 //parameters: DeviceHandle Ipv4Priority
 //AgtIgmpHostPool SetIpv4Priority, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) IsIpv4PriorityTypeOfServiceFieldSelected () error {
 //parameters: DeviceHandle
 //AgtIgmpHostPool IsIpv4PriorityTypeOfServiceFieldSelected, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) SelectIpv4PriorityTypeOfServiceField () error {
 //parameters: DeviceHandle
 //AgtIgmpHostPool SelectIpv4PriorityTypeOfServiceField, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) GetIpv4PriorityTypeOfServiceValue ()(string, error) {
 //parameters: DeviceHandle TosCodePointField
 //AgtIgmpHostPool GetIpv4PriorityTypeOfServiceValue
 return "", nil
}

func(np *IgmpHostPool) SetIpv4PriorityTypeOfServiceValue () error {
 //parameters: DeviceHandle TosCodePointField Value
 //AgtIgmpHostPool SetIpv4PriorityTypeOfServiceValue, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) IsIpv4PriorityDiffServFieldSelected () error {
 //parameters: DeviceHandle DiffServPerHopBehaviorGroup
 //AgtIgmpHostPool IsIpv4PriorityDiffServFieldSelected, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) SelectIpv4PriorityDiffServField () error {
 //parameters: DeviceHandle DiffServPerHopBehaviorGroup
 //AgtIgmpHostPool SelectIpv4PriorityDiffServField, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) GetIpv4PriorityDiffServValue ()(string, error) {
 //parameters: DeviceHandle DiffServCodePointField
 //AgtIgmpHostPool GetIpv4PriorityDiffServValue
 return "", nil
}

func(np *IgmpHostPool) SetIpv4PriorityDiffServValue () error {
 //parameters: DeviceHandle DiffServCodePointConfigurableField Value
 //AgtIgmpHostPool SetIpv4PriorityDiffServValue, m.Object, m.Name);
 return nil
}

func(np *IgmpHostPool) GetVlanPriority ()(string, error) {
 //parameters: DeviceHandle VlanTagIndex
 //AgtIgmpHostPool GetVlanPriority
 return "", nil
}

func(np *IgmpHostPool) SetVlanPriority () error {
 //parameters: DeviceHandle VlanTagIndex VlanPriority
 //AgtIgmpHostPool SetVlanPriority, m.Object, m.Name);
 return nil
}

