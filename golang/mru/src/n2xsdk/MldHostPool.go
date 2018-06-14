package n2xsdk

type MldHostPool struct {
 Handler string
}

func(np *MldHostPool) SetVersion () error {
 //parameters: DeviceHandle MldVersion
 //AgtMldHostPool SetVersion, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) GetVersion ()(string, error) {
 //parameters: DeviceHandle
 //AgtMldHostPool GetVersion
 return "", nil
}

func(np *MldHostPool) SetParameter () error {
 //parameters: DeviceHandle Parameter Value
 //AgtMldHostPool SetParameter, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) GetParameter ()(string, error) {
 //parameters: DeviceHandle Parameter
 //AgtMldHostPool GetParameter
 return "", nil
}

func(np *MldHostPool) GetInfo ()(string, error) {
 //parameters: SessionList Info
 //AgtMldHostPool GetInfo
 return "", nil
}

func(np *MldHostPool) AddGroupPools () error {
 //parameters: DeviceHandle GroupPoolHandles
 //AgtMldHostPool AddGroupPools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) RemoveGroupPools () error {
 //parameters: DeviceHandle GroupPoolHandles
 //AgtMldHostPool RemoveGroupPools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) RemoveAllGroupPools () error {
 //parameters: DeviceHandle
 //AgtMldHostPool RemoveAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) ListGroupPools ()(string, error) {
 //parameters: DeviceHandle
 //AgtMldHostPool ListGroupPools
 return "", nil
}

func(np *MldHostPool) AddSourcePools () error {
 //parameters: SessionList GroupPoolHandle SourcePoolHandles
 //AgtMldHostPool AddSourcePools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) RemoveSourcePools () error {
 //parameters: SessionList GroupPoolHandle SourcePoolHandles
 //AgtMldHostPool RemoveSourcePools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) RemoveAllSourcePools () error {
 //parameters: SessionList GroupPoolHandle
 //AgtMldHostPool RemoveAllSourcePools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) ListSourcePools ()(string, error) {
 //parameters: SessionList GroupPoolHandle
 //AgtMldHostPool ListSourcePools
 return "", nil
}

func(np *MldHostPool) SetGroupPoolFilter () error {
 //parameters: SessionList GroupPoolHandle MldFilter
 //AgtMldHostPool SetGroupPoolFilter, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) GetGroupPoolFilter ()(string, error) {
 //parameters: SessionList GroupPoolHandle
 //AgtMldHostPool GetGroupPoolFilter
 return "", nil
}

func(np *MldHostPool) SetGroupPoolState () error {
 //parameters: SessionList GroupPoolHandle MldFilter SourcePoolHandles
 //AgtMldHostPool SetGroupPoolState, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) GetGroupPoolState ()(string, error) {
 //parameters: SessionList GroupPoolHandle
 //AgtMldHostPool GetGroupPoolState
 return "", nil
}

func(np *MldHostPool) GetNumberOfSourcesInGroupPool ()(string, error) {
 //parameters: SessionList GroupPoolHandle
 //AgtMldHostPool GetNumberOfSourcesInGroupPool
 return "", nil
}

func(np *MldHostPool) ListGroupPoolsJoined ()(string, error) {
 //parameters: SessionList
 //AgtMldHostPool ListGroupPoolsJoined
 return "", nil
}

func(np *MldHostPool) IsJoined () error {
 //parameters: SessionList GroupPoolHandle
 //AgtMldHostPool IsJoined, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) JoinGroupPools () error {
 //parameters: SessionList GroupPoolHandles
 //AgtMldHostPool JoinGroupPools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) JoinAllGroupPools () error {
 //parameters: SessionList
 //AgtMldHostPool JoinAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) LeaveGroupPools () error {
 //parameters: SessionList GroupPoolHandles
 //AgtMldHostPool LeaveGroupPools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) LeaveAllGroupPools () error {
 //parameters: SessionList
 //AgtMldHostPool LeaveAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) StartPduSave () error {
 //parameters: SessionList
 //AgtMldHostPool StartPduSave, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) StopPduSave () error {
 //parameters: SessionList
 //AgtMldHostPool StopPduSave, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) ListSavedPdus ()(string, error) {
 //parameters: SessionList
 //AgtMldHostPool ListSavedPdus
 return "", nil
}

func(np *MldHostPool) GetPduDetails ()(string, error) {
 //parameters: SessionList PduId PduInfoType
 //AgtMldHostPool GetPduDetails
 return "", nil
}

func(np *MldHostPool) GetAllPduDetails ()(string, error) {
 //parameters: SessionList PduId
 //AgtMldHostPool GetAllPduDetails
 return "", nil
}

func(np *MldHostPool) DeletePdus () error {
 //parameters: SessionList
 //AgtMldHostPool DeletePdus, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) GetVlanPriority ()(string, error) {
 //parameters: DeviceHandle VlanTagIndex
 //AgtMldHostPool GetVlanPriority
 return "", nil
}

func(np *MldHostPool) SetVlanPriority () error {
 //parameters: DeviceHandle VlanTagIndex VlanPriority
 //AgtMldHostPool SetVlanPriority, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) IsIpv6PriorityNoCodePointFieldSelected () error {
 //parameters: DeviceHandle
 //AgtMldHostPool IsIpv6PriorityNoCodePointFieldSelected, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) SelectIpv6PriorityNoCodePointField () error {
 //parameters: DeviceHandle
 //AgtMldHostPool SelectIpv6PriorityNoCodePointField, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) GetIpv6Priority ()(string, error) {
 //parameters: DeviceHandle
 //AgtMldHostPool GetIpv6Priority
 return "", nil
}

func(np *MldHostPool) SetIpv6Priority () error {
 //parameters: DeviceHandle Ipv6Priority
 //AgtMldHostPool SetIpv6Priority, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) IsIpv6PriorityDiffServFieldSelected () error {
 //parameters: DeviceHandle DiffServPerHopBehaviorGroup
 //AgtMldHostPool IsIpv6PriorityDiffServFieldSelected, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) SelectIpv6PriorityDiffServField () error {
 //parameters: DeviceHandle DiffServPerHopBehaviorGroup
 //AgtMldHostPool SelectIpv6PriorityDiffServField, m.Object, m.Name);
 return nil
}

func(np *MldHostPool) GetIpv6PriorityDiffServValue ()(string, error) {
 //parameters: DeviceHandle DiffServCodePointField
 //AgtMldHostPool GetIpv6PriorityDiffServValue
 return "", nil
}

func(np *MldHostPool) SetIpv6PriorityDiffServValue () error {
 //parameters: DeviceHandle DiffServCodePointConfigurableField Value
 //AgtMldHostPool SetIpv6PriorityDiffServValue, m.Object, m.Name);
 return nil
}

