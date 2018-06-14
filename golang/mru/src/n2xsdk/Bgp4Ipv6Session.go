package n2xsdk

type Bgp4Ipv6Session struct {
 Handler string
}

func(np *Bgp4Ipv6Session) SetSutIpAddress () error {
 //parameters: SessionHandle RemoteIpAddress
 //AgtBgp4Ipv6Session SetSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) SetTesterIpAddress () error {
 //parameters: SessionHandle LocalIpAddress
 //AgtBgp4Ipv6Session SetTesterIpAddress, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) SetBgp4Identifier () error {
 //parameters: SessionHandle Bgp4Identifier
 //AgtBgp4Ipv6Session SetBgp4Identifier, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetBgp4Identifier ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetBgp4Identifier
 return "", nil
}

func(np *Bgp4Ipv6Session) GetSutIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetSutIpAddress
 return "", nil
}

func(np *Bgp4Ipv6Session) GetTesterIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetTesterIpAddress
 return "", nil
}

func(np *Bgp4Ipv6Session) SetSutAsNumber () error {
 //parameters: SessionHandle AsNumber
 //AgtBgp4Ipv6Session SetSutAsNumber, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetSutAsNumber ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetSutAsNumber
 return "", nil
}

func(np *Bgp4Ipv6Session) SetTesterAsNumber () error {
 //parameters: SessionHandle AsNumber
 //AgtBgp4Ipv6Session SetTesterAsNumber, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetTesterAsNumber ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetTesterAsNumber
 return "", nil
}

func(np *Bgp4Ipv6Session) SetTesterAsNumberType () error {
 //parameters: SessionHandle TesterAsNumberType
 //AgtBgp4Ipv6Session SetTesterAsNumberType, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetTesterAsNumberType ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetTesterAsNumberType
 return "", nil
}

func(np *Bgp4Ipv6Session) SetTester4ByteAsNumber () error {
 //parameters: SessionHandle AsNumber
 //AgtBgp4Ipv6Session SetTester4ByteAsNumber, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) SetTester4ByteAsNumberPlain () error {
 //parameters: SessionHandle FourByteFormatType AsNumber
 //AgtBgp4Ipv6Session SetTester4ByteAsNumberPlain, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetTester4ByteAsNumber ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetTester4ByteAsNumber
 return "", nil
}

func(np *Bgp4Ipv6Session) GetTester4ByteAsNumberPlain ()(string, error) {
 //parameters: SessionHandle FourByteFormatType
 //AgtBgp4Ipv6Session GetTester4ByteAsNumberPlain
 return "", nil
}

func(np *Bgp4Ipv6Session) Enable () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session Enable, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) Disable () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session Disable, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetEnableFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetEnableFlag
 return "", nil
}

func(np *Bgp4Ipv6Session) Open () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session Open, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) OpenAllSessions () error {
 //parameters: 
 //AgtBgp4Ipv6Session OpenAllSessions, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) Close () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session Close, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) CloseAllSessions () error {
 //parameters: 
 //AgtBgp4Ipv6Session CloseAllSessions, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) CloseSessionWithoutNotification () error {
 //parameters: SessionHandle CloseType
 //AgtBgp4Ipv6Session CloseSessionWithoutNotification, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) SetNotificationParameters () error {
 //parameters: SessionHandle ErrorCode ErrorSubcode
 //AgtBgp4Ipv6Session SetNotificationParameters, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetNotificationParameters ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetNotificationParameters
 return "", nil
}

func(np *Bgp4Ipv6Session) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetState
 return "", nil
}

func(np *Bgp4Ipv6Session) GetPortStateSummary ()(string, error) {
 //parameters: PortHandle
 //AgtBgp4Ipv6Session GetPortStateSummary
 return "", nil
}

func(np *Bgp4Ipv6Session) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtBgp4Ipv6Session GetStateSummary
 return "", nil
}

func(np *Bgp4Ipv6Session) GetLastError ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetLastError
 return "", nil
}

func(np *Bgp4Ipv6Session) SetOpenParameter () error {
 //parameters: SessionHandle OpenParameter Value
 //AgtBgp4Ipv6Session SetOpenParameter, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetOpenParameter ()(string, error) {
 //parameters: SessionHandle OpenParameter
 //AgtBgp4Ipv6Session GetOpenParameter
 return "", nil
}

func(np *Bgp4Ipv6Session) EnableKeepAliveTimerNegotiation () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session EnableKeepAliveTimerNegotiation, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) DisableKeepAliveTimerNegotiation () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session DisableKeepAliveTimerNegotiation, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) IsKeepAliveTimerNegotiationEnabled () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session IsKeepAliveTimerNegotiationEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) SetRoutesPerUpdate () error {
 //parameters: SessionHandle NumRoutes
 //AgtBgp4Ipv6Session SetRoutesPerUpdate, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetRoutesPerUpdate ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetRoutesPerUpdate
 return "", nil
}

func(np *Bgp4Ipv6Session) SetInterUpdateDelay () error {
 //parameters: SessionHandle Delay
 //AgtBgp4Ipv6Session SetInterUpdateDelay, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetInterUpdateDelay ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetInterUpdateDelay
 return "", nil
}

func(np *Bgp4Ipv6Session) EnableAllEndOfRibMarkers () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session EnableAllEndOfRibMarkers, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) DisableAllEndOfRibMarkers () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session DisableAllEndOfRibMarkers, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) IsAllEndOfRibMarkersEnabled () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session IsAllEndOfRibMarkersEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) EnableMd5Authentication () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session EnableMd5Authentication, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) DisableMd5Authentication () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session DisableMd5Authentication, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) IsMd5AuthenticationEnabled () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session IsMd5AuthenticationEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) SetMd5AuthenticationKey () error {
 //parameters: SessionHandle Md5Key
 //AgtBgp4Ipv6Session SetMd5AuthenticationKey, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetMd5AuthenticationKey ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6Session GetMd5AuthenticationKey
 return "", nil
}

func(np *Bgp4Ipv6Session) EnableAllSessions () error {
 //parameters: 
 //AgtBgp4Ipv6Session EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) DisableAllSessions () error {
 //parameters: 
 //AgtBgp4Ipv6Session DisableAllSessions, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) EnableAllPeerPools () error {
 //parameters: 
 //AgtBgp4Ipv6Session EnableAllPeerPools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) DisableAllPeerPools () error {
 //parameters: 
 //AgtBgp4Ipv6Session DisableAllPeerPools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) OpenAllPeerPools () error {
 //parameters: 
 //AgtBgp4Ipv6Session OpenAllPeerPools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) CloseAllPeerPools () error {
 //parameters: 
 //AgtBgp4Ipv6Session CloseAllPeerPools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) SetDefaultOpenParameter () error {
 //parameters: Parameter Value
 //AgtBgp4Ipv6Session SetDefaultOpenParameter, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6Session) GetDefaultOpenParameter ()(string, error) {
 //parameters: Parameter
 //AgtBgp4Ipv6Session GetDefaultOpenParameter
 return "", nil
}

