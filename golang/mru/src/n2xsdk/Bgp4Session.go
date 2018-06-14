package n2xsdk

type Bgp4Session struct {
 Handler string
}

func(np *Bgp4Session) SetSutIpAddress () error {
 //parameters: SessionHandle RemoteIpAddress
 //AgtBgp4Session SetSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetSutIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetSutIpAddress
 return "", nil
}

func(np *Bgp4Session) SetTesterIpAddress () error {
 //parameters: SessionHandle LocalIpAddress
 //AgtBgp4Session SetTesterIpAddress, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetTesterIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetTesterIpAddress
 return "", nil
}

func(np *Bgp4Session) SetSutAsNumber () error {
 //parameters: SessionHandle AsNumber
 //AgtBgp4Session SetSutAsNumber, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetSutAsNumber ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetSutAsNumber
 return "", nil
}

func(np *Bgp4Session) SetTesterAsNumber () error {
 //parameters: SessionHandle AsNumber
 //AgtBgp4Session SetTesterAsNumber, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetTesterAsNumber ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetTesterAsNumber
 return "", nil
}

func(np *Bgp4Session) SetTesterAsNumberType () error {
 //parameters: SessionHandle TesterAsNumberType
 //AgtBgp4Session SetTesterAsNumberType, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetTesterAsNumberType ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetTesterAsNumberType
 return "", nil
}

func(np *Bgp4Session) SetTester4ByteAsNumber () error {
 //parameters: SessionHandle AsNumber
 //AgtBgp4Session SetTester4ByteAsNumber, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) SetTester4ByteAsNumberPlain () error {
 //parameters: SessionHandle FourByteFormatType AsNumber
 //AgtBgp4Session SetTester4ByteAsNumberPlain, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetTester4ByteAsNumber ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetTester4ByteAsNumber
 return "", nil
}

func(np *Bgp4Session) GetTester4ByteAsNumberPlain ()(string, error) {
 //parameters: SessionHandle FourByteFormatType
 //AgtBgp4Session GetTester4ByteAsNumberPlain
 return "", nil
}

func(np *Bgp4Session) Enable () error {
 //parameters: SessionHandle
 //AgtBgp4Session Enable, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) Disable () error {
 //parameters: SessionHandle
 //AgtBgp4Session Disable, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetEnableFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetEnableFlag
 return "", nil
}

func(np *Bgp4Session) Open () error {
 //parameters: SessionHandle
 //AgtBgp4Session Open, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) OpenAllSessions () error {
 //parameters: 
 //AgtBgp4Session OpenAllSessions, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) Close () error {
 //parameters: SessionHandle
 //AgtBgp4Session Close, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) CloseAllSessions () error {
 //parameters: 
 //AgtBgp4Session CloseAllSessions, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) CloseSessionWithoutNotification () error {
 //parameters: SessionHandle CloseType
 //AgtBgp4Session CloseSessionWithoutNotification, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) SetNotificationParameters () error {
 //parameters: SessionHandle ErrorCode ErrorSubcode
 //AgtBgp4Session SetNotificationParameters, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetNotificationParameters ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetNotificationParameters
 return "", nil
}

func(np *Bgp4Session) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetState
 return "", nil
}

func(np *Bgp4Session) GetPortStateSummary ()(string, error) {
 //parameters: PortHandle
 //AgtBgp4Session GetPortStateSummary
 return "", nil
}

func(np *Bgp4Session) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtBgp4Session GetStateSummary
 return "", nil
}

func(np *Bgp4Session) GetLastError ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetLastError
 return "", nil
}

func(np *Bgp4Session) SetOpenParameter () error {
 //parameters: SessionHandle OpenParameter Value
 //AgtBgp4Session SetOpenParameter, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetOpenParameter ()(string, error) {
 //parameters: SessionHandle OpenParameter
 //AgtBgp4Session GetOpenParameter
 return "", nil
}

func(np *Bgp4Session) EnableKeepAliveTimerNegotiation () error {
 //parameters: SessionHandle
 //AgtBgp4Session EnableKeepAliveTimerNegotiation, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) DisableKeepAliveTimerNegotiation () error {
 //parameters: SessionHandle
 //AgtBgp4Session DisableKeepAliveTimerNegotiation, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) IsKeepAliveTimerNegotiationEnabled () error {
 //parameters: SessionHandle
 //AgtBgp4Session IsKeepAliveTimerNegotiationEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) SetRoutesPerUpdate () error {
 //parameters: SessionHandle NumRoutes
 //AgtBgp4Session SetRoutesPerUpdate, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetRoutesPerUpdate ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetRoutesPerUpdate
 return "", nil
}

func(np *Bgp4Session) SetInterUpdateDelay () error {
 //parameters: SessionHandle Delay
 //AgtBgp4Session SetInterUpdateDelay, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetInterUpdateDelay ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetInterUpdateDelay
 return "", nil
}

func(np *Bgp4Session) EnableAllEndOfRibMarkers () error {
 //parameters: SessionHandle
 //AgtBgp4Session EnableAllEndOfRibMarkers, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) DisableAllEndOfRibMarkers () error {
 //parameters: SessionHandle
 //AgtBgp4Session DisableAllEndOfRibMarkers, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) IsAllEndOfRibMarkersEnabled () error {
 //parameters: SessionHandle
 //AgtBgp4Session IsAllEndOfRibMarkersEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) EnableMd5Authentication () error {
 //parameters: SessionHandle
 //AgtBgp4Session EnableMd5Authentication, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) DisableMd5Authentication () error {
 //parameters: SessionHandle
 //AgtBgp4Session DisableMd5Authentication, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) IsMd5AuthenticationEnabled () error {
 //parameters: SessionHandle
 //AgtBgp4Session IsMd5AuthenticationEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) SetMd5AuthenticationKey () error {
 //parameters: SessionHandle Md5Key
 //AgtBgp4Session SetMd5AuthenticationKey, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetMd5AuthenticationKey ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Session GetMd5AuthenticationKey
 return "", nil
}

func(np *Bgp4Session) EnableAllSessions () error {
 //parameters: 
 //AgtBgp4Session EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) DisableAllSessions () error {
 //parameters: 
 //AgtBgp4Session DisableAllSessions, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) EnableAllPeerPools () error {
 //parameters: 
 //AgtBgp4Session EnableAllPeerPools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) DisableAllPeerPools () error {
 //parameters: 
 //AgtBgp4Session DisableAllPeerPools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) OpenAllPeerPools () error {
 //parameters: 
 //AgtBgp4Session OpenAllPeerPools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) CloseAllPeerPools () error {
 //parameters: 
 //AgtBgp4Session CloseAllPeerPools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) SetDefaultOpenParameter () error {
 //parameters: Parameter Value
 //AgtBgp4Session SetDefaultOpenParameter, m.Object, m.Name);
 return nil
}

func(np *Bgp4Session) GetDefaultOpenParameter ()(string, error) {
 //parameters: Parameter
 //AgtBgp4Session GetDefaultOpenParameter
 return "", nil
}

