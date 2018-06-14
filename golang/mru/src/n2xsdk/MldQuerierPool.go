package n2xsdk

type MldQuerierPool struct {
 Handler string
}

func(np *MldQuerierPool) SetVersion () error {
 //parameters: DeviceHandle MldVersion
 //AgtMldQuerierPool SetVersion, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) GetVersion ()(string, error) {
 //parameters: DeviceHandle
 //AgtMldQuerierPool GetVersion
 return "", nil
}

func(np *MldQuerierPool) SetParameter () error {
 //parameters: DeviceHandle Parameter Value
 //AgtMldQuerierPool SetParameter, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) GetParameter ()(string, error) {
 //parameters: DeviceHandle Parameter
 //AgtMldQuerierPool GetParameter
 return "", nil
}

func(np *MldQuerierPool) GetInfo ()(string, error) {
 //parameters: DeviceHandle Info
 //AgtMldQuerierPool GetInfo
 return "", nil
}

func(np *MldQuerierPool) SetGeneralQueryInterval () error {
 //parameters: DeviceHandle QueryInterval
 //AgtMldQuerierPool SetGeneralQueryInterval, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) GetGeneralQueryInterval ()(string, error) {
 //parameters: DeviceHandle
 //AgtMldQuerierPool GetGeneralQueryInterval
 return "", nil
}

func(np *MldQuerierPool) AddGroupQueryMessage () error {
 //parameters: DeviceHandle GroupPoolHandles
 //AgtMldQuerierPool AddGroupQueryMessage, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) RemoveGroupQueryMessage () error {
 //parameters: DeviceHandle GroupPoolHandle
 //AgtMldQuerierPool RemoveGroupQueryMessage, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) RemoveAllQueryMessages () error {
 //parameters: DeviceHandle
 //AgtMldQuerierPool RemoveAllQueryMessages, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) ListGroupQueryMessages ()(string, error) {
 //parameters: DeviceHandle
 //AgtMldQuerierPool ListGroupQueryMessages
 return "", nil
}

func(np *MldQuerierPool) AddSourcePools () error {
 //parameters: DeviceHandle GroupPoolHandle SourcePoolHandles
 //AgtMldQuerierPool AddSourcePools, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) RemoveSourcePools () error {
 //parameters: DeviceHandle GroupPoolHandle SourcePoolHandles
 //AgtMldQuerierPool RemoveSourcePools, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) RemoveAllSourcePools () error {
 //parameters: DeviceHandle GroupPoolHandle
 //AgtMldQuerierPool RemoveAllSourcePools, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) ListSourcePools ()(string, error) {
 //parameters: DeviceHandle GroupPoolHandle
 //AgtMldQuerierPool ListSourcePools
 return "", nil
}

func(np *MldQuerierPool) GetNumberOfSourcesInGroupPool ()(string, error) {
 //parameters: DeviceHandle GroupPoolHandle
 //AgtMldQuerierPool GetNumberOfSourcesInGroupPool
 return "", nil
}

func(np *MldQuerierPool) SetGroupQueryParameters () error {
 //parameters: DeviceHandle GroupPoolHandle Lmqi Lmqc
 //AgtMldQuerierPool SetGroupQueryParameters, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) GetGroupQueryParameters ()(string, error) {
 //parameters: DeviceHandle GroupPoolHandle
 //AgtMldQuerierPool GetGroupQueryParameters
 return "", nil
}

func(np *MldQuerierPool) StartGeneralQueryMessage () error {
 //parameters: DeviceHandle
 //AgtMldQuerierPool StartGeneralQueryMessage, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) StopGeneralQueryMessage () error {
 //parameters: DeviceHandle
 //AgtMldQuerierPool StopGeneralQueryMessage, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) IsGeneralQueryRunning () error {
 //parameters: DeviceHandle
 //AgtMldQuerierPool IsGeneralQueryRunning, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) StartGroupQueryMessages () error {
 //parameters: DeviceHandle GroupPoolHandle
 //AgtMldQuerierPool StartGroupQueryMessages, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) StartAllQueryMessages () error {
 //parameters: DeviceHandle
 //AgtMldQuerierPool StartAllQueryMessages, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) StopGroupQueryMessages () error {
 //parameters: DeviceHandle GroupPoolHandle
 //AgtMldQuerierPool StopGroupQueryMessages, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) StopAllQueryMessages () error {
 //parameters: DeviceHandle
 //AgtMldQuerierPool StopAllQueryMessages, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) ListGroupQueryMessagesRunning ()(string, error) {
 //parameters: DeviceHandle
 //AgtMldQuerierPool ListGroupQueryMessagesRunning
 return "", nil
}

func(np *MldQuerierPool) IsGroupQueryRunning () error {
 //parameters: DeviceHandle GroupPoolHandle
 //AgtMldQuerierPool IsGroupQueryRunning, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) GetVlanPriority ()(string, error) {
 //parameters: DeviceHandle VlanTagIndex
 //AgtMldQuerierPool GetVlanPriority
 return "", nil
}

func(np *MldQuerierPool) SetVlanPriority () error {
 //parameters: DeviceHandle VlanTagIndex VlanPriority
 //AgtMldQuerierPool SetVlanPriority, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) IsIpv6PriorityNoCodePointFieldSelected () error {
 //parameters: DeviceHandle
 //AgtMldQuerierPool IsIpv6PriorityNoCodePointFieldSelected, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) SelectIpv6PriorityNoCodePointField () error {
 //parameters: DeviceHandle
 //AgtMldQuerierPool SelectIpv6PriorityNoCodePointField, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) GetIpv6Priority ()(string, error) {
 //parameters: DeviceHandle
 //AgtMldQuerierPool GetIpv6Priority
 return "", nil
}

func(np *MldQuerierPool) SetIpv6Priority () error {
 //parameters: DeviceHandle Ipv6Priority
 //AgtMldQuerierPool SetIpv6Priority, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) IsIpv6PriorityDiffServFieldSelected () error {
 //parameters: DeviceHandle DiffServPerHopBehaviorGroup
 //AgtMldQuerierPool IsIpv6PriorityDiffServFieldSelected, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) SelectIpv6PriorityDiffServField () error {
 //parameters: DeviceHandle DiffServPerHopBehaviorGroup
 //AgtMldQuerierPool SelectIpv6PriorityDiffServField, m.Object, m.Name);
 return nil
}

func(np *MldQuerierPool) GetIpv6PriorityDiffServValue ()(string, error) {
 //parameters: DeviceHandle DiffServCodePointField
 //AgtMldQuerierPool GetIpv6PriorityDiffServValue
 return "", nil
}

func(np *MldQuerierPool) SetIpv6PriorityDiffServValue () error {
 //parameters: DeviceHandle DiffServCodePointConfigurableField Value
 //AgtMldQuerierPool SetIpv6PriorityDiffServValue, m.Object, m.Name);
 return nil
}

