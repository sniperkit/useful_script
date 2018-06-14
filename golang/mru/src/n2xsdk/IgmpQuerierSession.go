package n2xsdk

type IgmpQuerierSession struct {
 Handler string
}

func(np *IgmpQuerierSession) GetInterfaceIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession GetInterfaceIpAddress
 return "", nil
}

func(np *IgmpQuerierSession) SetSutIpAddress () error {
 //parameters: SessionHandle SutIpAddress
 //AgtIgmpQuerierSession SetSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) GetSutIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession GetSutIpAddress
 return "", nil
}

func(np *IgmpQuerierSession) Enable () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession Enable, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) EnableAllSessions () error {
 //parameters: 
 //AgtIgmpQuerierSession EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) Disable () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession Disable, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) DisableAllSessions () error {
 //parameters: 
 //AgtIgmpQuerierSession DisableAllSessions, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) GetEnableFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession GetEnableFlag
 return "", nil
}

func(np *IgmpQuerierSession) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession GetState
 return "", nil
}

func(np *IgmpQuerierSession) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtIgmpQuerierSession GetStateSummary
 return "", nil
}

func(np *IgmpQuerierSession) GetPortStateSummary ()(string, error) {
 //parameters: PortHandle
 //AgtIgmpQuerierSession GetPortStateSummary
 return "", nil
}

func(np *IgmpQuerierSession) GetVersion ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession GetVersion
 return "", nil
}

func(np *IgmpQuerierSession) GetParameter ()(string, error) {
 //parameters: SessionHandle Parameter
 //AgtIgmpQuerierSession GetParameter
 return "", nil
}

func(np *IgmpQuerierSession) GetInfo ()(string, error) {
 //parameters: SessionHandle Info
 //AgtIgmpQuerierSession GetInfo
 return "", nil
}

func(np *IgmpQuerierSession) AddGeneralQueryMessage () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession AddGeneralQueryMessage, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) RemoveGeneralQueryMessage () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession RemoveGeneralQueryMessage, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) IsGeneralQueryAdded () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession IsGeneralQueryAdded, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) SetGeneralQueryInterval () error {
 //parameters: SessionHandle QueryInterval
 //AgtIgmpQuerierSession SetGeneralQueryInterval, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) GetGeneralQueryInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession GetGeneralQueryInterval
 return "", nil
}

func(np *IgmpQuerierSession) AddGroupQueryMessage () error {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpQuerierSession AddGroupQueryMessage, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) RemoveGroupQueryMessage () error {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpQuerierSession RemoveGroupQueryMessage, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) RemoveAllQueryMessages () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession RemoveAllQueryMessages, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) ListGroupQueryMessages ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession ListGroupQueryMessages
 return "", nil
}

func(np *IgmpQuerierSession) AddSourcePools () error {
 //parameters: SessionHandle GroupPoolHandle SourcePoolHandles
 //AgtIgmpQuerierSession AddSourcePools, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) RemoveSourcePools () error {
 //parameters: SessionHandle GroupPoolHandle SourcePoolHandles
 //AgtIgmpQuerierSession RemoveSourcePools, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) RemoveAllSourcePools () error {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpQuerierSession RemoveAllSourcePools, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) ListSourcePools ()(string, error) {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpQuerierSession ListSourcePools
 return "", nil
}

func(np *IgmpQuerierSession) GetNumberOfSourcesInGroupPool ()(string, error) {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpQuerierSession GetNumberOfSourcesInGroupPool
 return "", nil
}

func(np *IgmpQuerierSession) ListGroupPoolsInvalid ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession ListGroupPoolsInvalid
 return "", nil
}

func(np *IgmpQuerierSession) IsInvalid () error {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpQuerierSession IsInvalid, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) SetGroupQueryParameters () error {
 //parameters: SessionHandle GroupPoolHandle Lmqi Lmqc
 //AgtIgmpQuerierSession SetGroupQueryParameters, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) GetGroupQueryParameters ()(string, error) {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpQuerierSession GetGroupQueryParameters
 return "", nil
}

func(np *IgmpQuerierSession) StartGeneralQueryMessage () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession StartGeneralQueryMessage, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) StopGeneralQueryMessage () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession StopGeneralQueryMessage, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) IsGeneralQueryRunning () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession IsGeneralQueryRunning, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) StartGroupQueryMessages () error {
 //parameters: SessionHandle GroupPoolHandles
 //AgtIgmpQuerierSession StartGroupQueryMessages, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) StartAllQueryMessages () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession StartAllQueryMessages, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) StopGroupQueryMessages () error {
 //parameters: SessionHandle GroupPoolHandles
 //AgtIgmpQuerierSession StopGroupQueryMessages, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) StopAllQueryMessages () error {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession StopAllQueryMessages, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) ListGroupQueryMessagesRunning ()(string, error) {
 //parameters: SessionHandle
 //AgtIgmpQuerierSession ListGroupQueryMessagesRunning
 return "", nil
}

func(np *IgmpQuerierSession) IsGroupQueryRunning () error {
 //parameters: SessionHandle GroupPoolHandle
 //AgtIgmpQuerierSession IsGroupQueryRunning, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) SetInterfaceIpAddress () error {
 //parameters: SessionHandle TesterIpAddress
 //AgtIgmpQuerierSession SetInterfaceIpAddress, m.Object, m.Name);
 return nil
}

func(np *IgmpQuerierSession) SetParameter () error {
 //parameters: 
 //AgtIgmpQuerierSession SetParameter, m.Object, m.Name);
 return nil
}

