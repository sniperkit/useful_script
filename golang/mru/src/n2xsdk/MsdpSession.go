package n2xsdk

type MsdpSession struct {
 Handler string
}

func(np *MsdpSession) SetSutIpAddress () error {
 //parameters: SessionHandle SutIpAddress
 //AgtMsdpSession SetSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) GetSutIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpSession GetSutIpAddress
 return "", nil
}

func(np *MsdpSession) SetInterfaceIpAddress () error {
 //parameters: SessionHandle TesterIpAddress TesterIpAddressPrefix
 //AgtMsdpSession SetInterfaceIpAddress, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) SetMsdpTimer () error {
 //parameters: SessionHandle TimerType Period
 //AgtMsdpSession SetMsdpTimer, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) GetMsdpTimer ()(string, error) {
 //parameters: SessionHandle TimerType
 //AgtMsdpSession GetMsdpTimer
 return "", nil
}

func(np *MsdpSession) AddMembershipPools () error {
 //parameters: SessionHandle GroupPoolHandles
 //AgtMsdpSession AddMembershipPools, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) SetPoolRpAddress () error {
 //parameters: SessionHandle MembershipPoolHandle RpIpAddress
 //AgtMsdpSession SetPoolRpAddress, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) GetPoolRpAddress ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtMsdpSession GetPoolRpAddress
 return "", nil
}

func(np *MsdpSession) GetLastError ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpSession GetLastError
 return "", nil
}

func(np *MsdpSession) SetRouterId () error {
 //parameters: SessionHandle RouterId
 //AgtMsdpSession SetRouterId, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) Enable () error {
 //parameters: SessionHandle
 //AgtMsdpSession Enable, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) Disable () error {
 //parameters: SessionHandle
 //AgtMsdpSession Disable, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) EnableAll () error {
 //parameters: 
 //AgtMsdpSession EnableAll, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) DisableAll () error {
 //parameters: 
 //AgtMsdpSession DisableAll, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) GetInterfaceIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpSession GetInterfaceIpAddress
 return "", nil
}

func(np *MsdpSession) GetRouterId ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpSession GetRouterId
 return "", nil
}

func(np *MsdpSession) GetLinkLocalInterfaceAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpSession GetLinkLocalInterfaceAddress
 return "", nil
}

func(np *MsdpSession) GetNumberOfGroupAddresses ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtMsdpSession GetNumberOfGroupAddresses
 return "", nil
}

func(np *MsdpSession) GetNumberOfSourceAddresses ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtMsdpSession GetNumberOfSourceAddresses
 return "", nil
}

func(np *MsdpSession) RemoveMembershipPools () error {
 //parameters: SessionHandle Count psaMembershipPoolHandles
 //AgtMsdpSession RemoveMembershipPools, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) RemoveAllMembershipPools () error {
 //parameters: SessionHandle
 //AgtMsdpSession RemoveAllMembershipPools, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) ListMembershipPools ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpSession ListMembershipPools
 return "", nil
}

func(np *MsdpSession) SetGroupPoolHandle () error {
 //parameters: SessionHandle MembershipPoolHandle GroupPoolHandle
 //AgtMsdpSession SetGroupPoolHandle, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) GetGroupPoolHandle ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtMsdpSession GetGroupPoolHandle
 return "", nil
}

func(np *MsdpSession) AddSourcePools () error {
 //parameters: SessionHandle MembershipPoolHandle NumSources psaSourcePoolHandles
 //AgtMsdpSession AddSourcePools, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) RemoveSourcePools () error {
 //parameters: SessionHandle MembershipPoolHandle Count psaSourcePoolHandles
 //AgtMsdpSession RemoveSourcePools, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) RemoveAllSourcePools () error {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtMsdpSession RemoveAllSourcePools, m.Object, m.Name);
 return nil
}

func(np *MsdpSession) ListSourcePools ()(string, error) {
 //parameters: SessionHandle MembershipPoolHandle
 //AgtMsdpSession ListSourcePools
 return "", nil
}

func(np *MsdpSession) GetEnableFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpSession GetEnableFlag
 return "", nil
}

func(np *MsdpSession) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpSession GetState
 return "", nil
}

func(np *MsdpSession) GetPortStateSummary ()(string, error) {
 //parameters: PortHandle
 //AgtMsdpSession GetPortStateSummary
 return "", nil
}

func(np *MsdpSession) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtMsdpSession GetStateSummary
 return "", nil
}

func(np *MsdpSession) GetMulticastRoutingProtocol ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpSession GetMulticastRoutingProtocol
 return "", nil
}

