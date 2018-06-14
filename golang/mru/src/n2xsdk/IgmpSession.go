package n2xsdk

type IgmpSession struct {
 Handler string
}

func(np *IgmpSession) SetInterfaceIpAddress () error {
 //parameters: SessionHandle TesterIpAddress
 //AgtIgmpSession SetInterfaceIpAddress, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) GetInterfaceIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession GetInterfaceIpAddress
 return "", nil
}

func(np *IgmpSession) SetSutIpAddress () error {
 //parameters: SessionHandle SutIpAddress
 //AgtIgmpSession SetSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) GetSutIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession GetSutIpAddress
 return "", nil
}

func(np *IgmpSession) Enable () error {
 //parameters: SessionHandle
 //AgtIgmpSession Enable, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) Disable () error {
 //parameters: SessionHandle
 //AgtIgmpSession Disable, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) GetParameter ()(string, error) {
 //parameters: SessionHandle Parameter
 //AgtIgmpSession GetParameter
 return "", nil
}

func(np *IgmpSession) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession GetState
 return "", nil
}

func(np *IgmpSession) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtIgmpSession GetStateSummary
 return "", nil
}

func(np *IgmpSession) GetPortStateSummary ()(string, error) {
 //parameters: PortHandle
 //AgtIgmpSession GetPortStateSummary
 return "", nil
}

func(np *IgmpSession) GetVersion ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession GetVersion
 return "", nil
}

func(np *IgmpSession) GetInfo ()(string, error) {
 //parameters: SessionHandle Info
 //AgtIgmpSession GetInfo
 return "", nil
}

func(np *IgmpSession) ListIgmpReservedGroups ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession ListIgmpReservedGroups
 return "", nil
}

func(np *IgmpSession) AddGroupPools () error {
 //parameters: SessionHandle GroupPoolHandles
 //AgtIgmpSession AddGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) RemoveGroupPools () error {
 //parameters: SessionHandle GroupPoolHandles
 //AgtIgmpSession RemoveGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) RemoveAllGroupPools () error {
 //parameters: SessionHandle
 //AgtIgmpSession RemoveAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) ListGroupPools ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession ListGroupPools
 return "", nil
}

func(np *IgmpSession) AddSourcePools () error {
 //parameters: SessionHandle GroupPoolHandle SourcePoolHandles
 //AgtIgmpSession AddSourcePools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) RemoveSourcePools () error {
 //parameters: SessionHandle GroupPoolHandle SourcePoolHandles
 //AgtIgmpSession RemoveSourcePools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) RemoveAllSourcePools () error {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpSession RemoveAllSourcePools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) ListSourcePools ()(string, error) {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpSession ListSourcePools
 return "", nil
}

func(np *IgmpSession) SetGroupPoolFilter () error {
 //parameters: SessionHandle GroupPoolHandle IgmpFilter
 //AgtIgmpSession SetGroupPoolFilter, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) GetGroupPoolFilter ()(string, error) {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpSession GetGroupPoolFilter
 return "", nil
}

func(np *IgmpSession) SetGroupPoolState () error {
 //parameters: SessionHandle GroupPoolHandle IgmpFilter SourcePoolHandles
 //AgtIgmpSession SetGroupPoolState, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) GetGroupPoolState ()(string, error) {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpSession GetGroupPoolState
 return "", nil
}

func(np *IgmpSession) GetNumberOfSourcesInGroupPool ()(string, error) {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpSession GetNumberOfSourcesInGroupPool
 return "", nil
}

func(np *IgmpSession) ListGroupPoolsInvalid ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession ListGroupPoolsInvalid
 return "", nil
}

func(np *IgmpSession) IsInvalid () error {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpSession IsInvalid, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) ListGroupPoolsJoined ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession ListGroupPoolsJoined
 return "", nil
}

func(np *IgmpSession) IsJoined () error {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpSession IsJoined, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) JoinGroupPools () error {
 //parameters: SessionHandle GroupPoolHandles
 //AgtIgmpSession JoinGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) JoinAllGroupPools () error {
 //parameters: SessionHandle
 //AgtIgmpSession JoinAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) LeaveGroupPools () error {
 //parameters: SessionHandle GroupPoolHandles
 //AgtIgmpSession LeaveGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) LeaveAllGroupPools () error {
 //parameters: SessionHandle
 //AgtIgmpSession LeaveAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) GetEnableFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession GetEnableFlag
 return "", nil
}

func(np *IgmpSession) SetPduBuffer () error {
 //parameters: SessionHandle MaxBufSize Cyclic
 //AgtIgmpSession SetPduBuffer, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) StartPduSave () error {
 //parameters: SessionHandle
 //AgtIgmpSession StartPduSave, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) StopPduSave () error {
 //parameters: SessionHandle
 //AgtIgmpSession StopPduSave, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) ListSavedPdus ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpSession ListSavedPdus
 return "", nil
}

func(np *IgmpSession) GetPduDetails ()(string, error) {
 //parameters: SessionHandle PduId PduInfoType
 //AgtIgmpSession GetPduDetails
 return "", nil
}

func(np *IgmpSession) DeleteAllPdusInSession () error {
 //parameters: SessionHandle
 //AgtIgmpSession DeleteAllPdusInSession, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) DeleteAllPdus () error {
 //parameters: 
 //AgtIgmpSession DeleteAllPdus, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) CreateSessionPool () error {
 //parameters: SessionHandles
 //AgtIgmpSession CreateSessionPool, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) DeleteSessionPool () error {
 //parameters: SessionPoolHandle
 //AgtIgmpSession DeleteSessionPool, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) ListSessionsInPool ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtIgmpSession ListSessionsInPool
 return "", nil
}

func(np *IgmpSession) ListSessionPools ()(string, error) {
 //parameters: 
 //AgtIgmpSession ListSessionPools
 return "", nil
}

func(np *IgmpSession) DisableSessionPoolUpdateInfo () error {
 //parameters: 
 //AgtIgmpSession DisableSessionPoolUpdateInfo, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) EnableSessionPoolUpdateInfo () error {
 //parameters: 
 //AgtIgmpSession EnableSessionPoolUpdateInfo, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) IsSessionPoolUpdateInfoEnabled () error {
 //parameters: 
 //AgtIgmpSession IsSessionPoolUpdateInfoEnabled, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) SessionPoolJoinGroupPools () error {
 //parameters: SessionPoolHandle GroupPoolHandles Burst Delay
 //AgtIgmpSession SessionPoolJoinGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) SessionPoolLeaveGroupPools () error {
 //parameters: SessionPoolHandle GroupPoolHandles Burst Delay
 //AgtIgmpSession SessionPoolLeaveGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) SessionPoolJoinAllGroupPools () error {
 //parameters: SessionPoolHandle Burst Delay
 //AgtIgmpSession SessionPoolJoinAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) SessionPoolLeaveAllGroupPools () error {
 //parameters: SessionPoolHandle Burst Delay
 //AgtIgmpSession SessionPoolLeaveAllGroupPools, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) EnableAllSessions () error {
 //parameters: 
 //AgtIgmpSession EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) DisableAllSessions () error {
 //parameters: 
 //AgtIgmpSession DisableAllSessions, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) SetParameter () error {
 //parameters: 
 //AgtIgmpSession SetParameter, m.Object, m.Name);
 return nil
}

func(np *IgmpSession) SessionPoolSetPduBuffer () error {
 //parameters: SessionPoolHandle MaxBufSize
 //AgtIgmpSession SessionPoolSetPduBuffer, m.Object, m.Name);
 return nil
}

