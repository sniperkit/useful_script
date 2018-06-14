package n2xsdk

type IsisSession struct {
 Handler string
}

func(np *IsisSession) SetMaximumPacketSize () error {
 //parameters: SessionHandle MaximumPacketSize
 //AgtIsisSession SetMaximumPacketSize, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetMaximumPacketSize ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetMaximumPacketSize
 return "", nil
}

func(np *IsisSession) Restart () error {
 //parameters: SessionHandle
 //AgtIsisSession Restart, m.Object, m.Name);
 return nil
}

func(np *IsisSession) RestartAllSessions () error {
 //parameters: 
 //AgtIsisSession RestartAllSessions, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetNeighborState ()(string, error) {
 //parameters: SessionHandle NeighborHandle
 //AgtIsisSession GetNeighborState
 return "", nil
}

func(np *IsisSession) GetSutLinkState ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetSutLinkState
 return "", nil
}

func(np *IsisSession) GetSessionState ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetSessionState
 return "", nil
}

func(np *IsisSession) AddNeighbor () error {
 //parameters: SessionHandle
 //AgtIsisSession AddNeighbor, m.Object, m.Name);
 return nil
}

func(np *IsisSession) RemoveNeighbor () error {
 //parameters: SessionHandle NeighborHandle
 //AgtIsisSession RemoveNeighbor, m.Object, m.Name);
 return nil
}

func(np *IsisSession) RemoveAllNeighbors () error {
 //parameters: SessionHandle
 //AgtIsisSession RemoveAllNeighbors, m.Object, m.Name);
 return nil
}

func(np *IsisSession) ListNeighbors ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession ListNeighbors
 return "", nil
}

func(np *IsisSession) SetSutLinkIpAddress () error {
 //parameters: SessionHandle Ipv4Address Ipv4PrefixLength
 //AgtIsisSession SetSutLinkIpAddress, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetSutLinkIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetSutLinkIpAddress
 return "", nil
}

func(np *IsisSession) SetSutLinkIpv4Address () error {
 //parameters: SessionHandle Ipv4Address Ipv4PrefixLength
 //AgtIsisSession SetSutLinkIpv4Address, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetSutLinkIpv4Address ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetSutLinkIpv4Address
 return "", nil
}

func(np *IsisSession) SetSutLinkIpv6Address () error {
 //parameters: SessionHandle Ipv6Address Ipv6PrefixLength
 //AgtIsisSession SetSutLinkIpv6Address, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetSutLinkIpv6Address ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetSutLinkIpv6Address
 return "", nil
}

func(np *IsisSession) GetSutLinkLocalIpv6Address ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetSutLinkLocalIpv6Address
 return "", nil
}

func(np *IsisSession) SetAreaId () error {
 //parameters: SessionHandle AreaId
 //AgtIsisSession SetAreaId, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetAreaId ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetAreaId
 return "", nil
}

func(np *IsisSession) SetAreaList ()(string, error) {
 //parameters: SessionHandle AreaList
 //AgtIsisSession SetAreaList
 return "", nil
}

func(np *IsisSession) GetAreaList ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetAreaList
 return "", nil
}

func(np *IsisSession) SetSystemId () error {
 //parameters: SessionHandle SystemId
 //AgtIsisSession SetSystemId, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetSystemId ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetSystemId
 return "", nil
}

func(np *IsisSession) IsEnabled () error {
 //parameters: SessionHandle
 //AgtIsisSession IsEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetStateSummary ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetStateSummary
 return "", nil
}

func(np *IsisSession) GetLspDatabase ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetLspDatabase
 return "", nil
}

func(np *IsisSession) GetSessionRouter ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetSessionRouter
 return "", nil
}

func(np *IsisSession) SetTeRouterId () error {
 //parameters: SessionHandle TeRouterId
 //AgtIsisSession SetTeRouterId, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetTeRouterId ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetTeRouterId
 return "", nil
}

func(np *IsisSession) SetPseudonodeNumber () error {
 //parameters: SessionHandle PseudonodeNumber
 //AgtIsisSession SetPseudonodeNumber, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetPseudonodeNumber ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetPseudonodeNumber
 return "", nil
}

func(np *IsisSession) SetRoutingLevel () error {
 //parameters: SessionHandle RoutingLevel
 //AgtIsisSession SetRoutingLevel, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetRoutingLevel ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetRoutingLevel
 return "", nil
}

func(np *IsisSession) SetNetworkType () error {
 //parameters: SessionHandle NetworkType
 //AgtIsisSession SetNetworkType, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetNetworkType ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetNetworkType
 return "", nil
}

func(np *IsisSession) EnableWideMetrics () error {
 //parameters: SessionHandle
 //AgtIsisSession EnableWideMetrics, m.Object, m.Name);
 return nil
}

func(np *IsisSession) DisableWideMetrics () error {
 //parameters: SessionHandle
 //AgtIsisSession DisableWideMetrics, m.Object, m.Name);
 return nil
}

func(np *IsisSession) IsWideMetricsEnabled () error {
 //parameters: SessionHandle
 //AgtIsisSession IsWideMetricsEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) SetIpMode () error {
 //parameters: SessionHandle IpMode
 //AgtIsisSession SetIpMode, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetIpMode ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetIpMode
 return "", nil
}

func(np *IsisSession) EnableLspDiscardMode () error {
 //parameters: SessionHandle
 //AgtIsisSession EnableLspDiscardMode, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetLspDiscardModeFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtIsisSession GetLspDiscardModeFlag
 return "", nil
}

func(np *IsisSession) DisableLspDiscardMode () error {
 //parameters: SessionHandle
 //AgtIsisSession DisableLspDiscardMode, m.Object, m.Name);
 return nil
}

func(np *IsisSession) EnableMultiTopologies () error {
 //parameters: SessionHandle
 //AgtIsisSession EnableMultiTopologies, m.Object, m.Name);
 return nil
}

func(np *IsisSession) DisableMultiTopologies () error {
 //parameters: SessionHandle
 //AgtIsisSession DisableMultiTopologies, m.Object, m.Name);
 return nil
}

func(np *IsisSession) IsMultiTopologiesEnabled () error {
 //parameters: SessionHandle
 //AgtIsisSession IsMultiTopologiesEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) Enable3WayHandshake () error {
 //parameters: SessionHandle
 //AgtIsisSession Enable3WayHandshake, m.Object, m.Name);
 return nil
}

func(np *IsisSession) Disable3WayHandshake () error {
 //parameters: SessionHandle
 //AgtIsisSession Disable3WayHandshake, m.Object, m.Name);
 return nil
}

func(np *IsisSession) Is3WayHandshakeEnabled () error {
 //parameters: SessionHandle
 //AgtIsisSession Is3WayHandshakeEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) EnableGracefulRestart () error {
 //parameters: SessionHandle
 //AgtIsisSession EnableGracefulRestart, m.Object, m.Name);
 return nil
}

func(np *IsisSession) DisableGracefulRestart () error {
 //parameters: SessionHandle
 //AgtIsisSession DisableGracefulRestart, m.Object, m.Name);
 return nil
}

func(np *IsisSession) IsGracefulRestartEnabled () error {
 //parameters: SessionHandle
 //AgtIsisSession IsGracefulRestartEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) EnableIihPadding () error {
 //parameters: SessionHandle
 //AgtIsisSession EnableIihPadding, m.Object, m.Name);
 return nil
}

func(np *IsisSession) DisableIihPadding () error {
 //parameters: SessionHandle
 //AgtIsisSession DisableIihPadding, m.Object, m.Name);
 return nil
}

func(np *IsisSession) IsIihPaddingEnabled () error {
 //parameters: SessionHandle
 //AgtIsisSession IsIihPaddingEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) SetSutLinkParameter () error {
 //parameters: SessionHandle SutLinkParameter SutLinkParameterValue
 //AgtIsisSession SetSutLinkParameter, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetSutLinkParameter ()(string, error) {
 //parameters: SessionHandle SutLinkParameter
 //AgtIsisSession GetSutLinkParameter
 return "", nil
}

func(np *IsisSession) Enable () error {
 //parameters: SessionHandle
 //AgtIsisSession Enable, m.Object, m.Name);
 return nil
}

func(np *IsisSession) Disable () error {
 //parameters: SessionHandle
 //AgtIsisSession Disable, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetRoutersAndConnections ()(string, error) {
 //parameters: SessionHandle IsGetConnections
 //AgtIsisSession GetRoutersAndConnections
 return "", nil
}

func(np *IsisSession) SetAuthenticationMode () error {
 //parameters: SessionHandle KeyType AuthenticationMode
 //AgtIsisSession SetAuthenticationMode, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetAuthenticationMode ()(string, error) {
 //parameters: SessionHandle KeyType
 //AgtIsisSession GetAuthenticationMode
 return "", nil
}

func(np *IsisSession) EnableAuthenticationKey () error {
 //parameters: SessionHandle KeyType
 //AgtIsisSession EnableAuthenticationKey, m.Object, m.Name);
 return nil
}

func(np *IsisSession) DisableAuthenticationKey () error {
 //parameters: SessionHandle KeyType
 //AgtIsisSession DisableAuthenticationKey, m.Object, m.Name);
 return nil
}

func(np *IsisSession) IsAuthenticationKeyEnabled () error {
 //parameters: SessionHandle KeyType
 //AgtIsisSession IsAuthenticationKeyEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) EnableAuthenticationSendKey () error {
 //parameters: SessionHandle KeyType
 //AgtIsisSession EnableAuthenticationSendKey, m.Object, m.Name);
 return nil
}

func(np *IsisSession) DisableAuthenticationSendKey () error {
 //parameters: SessionHandle KeyType
 //AgtIsisSession DisableAuthenticationSendKey, m.Object, m.Name);
 return nil
}

func(np *IsisSession) IsAuthenticationSendKeyEnabled () error {
 //parameters: SessionHandle KeyType
 //AgtIsisSession IsAuthenticationSendKeyEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) EnableAuthenticationReceiveKey () error {
 //parameters: SessionHandle KeyType
 //AgtIsisSession EnableAuthenticationReceiveKey, m.Object, m.Name);
 return nil
}

func(np *IsisSession) DisableAuthenticationReceiveKey () error {
 //parameters: SessionHandle KeyType
 //AgtIsisSession DisableAuthenticationReceiveKey, m.Object, m.Name);
 return nil
}

func(np *IsisSession) IsAuthenticationReceiveKeyEnabled () error {
 //parameters: SessionHandle KeyType
 //AgtIsisSession IsAuthenticationReceiveKeyEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) SetAuthenticationSendKey () error {
 //parameters: SessionHandle KeyType Key
 //AgtIsisSession SetAuthenticationSendKey, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetAuthenticationSendKey ()(string, error) {
 //parameters: SessionHandle KeyType
 //AgtIsisSession GetAuthenticationSendKey
 return "", nil
}

func(np *IsisSession) SetAuthenticationReceiveKey () error {
 //parameters: SessionHandle KeyType Key
 //AgtIsisSession SetAuthenticationReceiveKey, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetAuthenticationReceiveKey ()(string, error) {
 //parameters: SessionHandle KeyType
 //AgtIsisSession GetAuthenticationReceiveKey
 return "", nil
}

func(np *IsisSession) EnableAllSessions () error {
 //parameters: 
 //AgtIsisSession EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *IsisSession) DisableAllSessions () error {
 //parameters: 
 //AgtIsisSession DisableAllSessions, m.Object, m.Name);
 return nil
}

func(np *IsisSession) SetMacAddressSource () error {
 //parameters: MacAddressSourceType
 //AgtIsisSession SetMacAddressSource, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetMacAddressSource ()(string, error) {
 //parameters: 
 //AgtIsisSession GetMacAddressSource
 return "", nil
}

func(np *IsisSession) EnableAuthentication () error {
 //parameters: SessionHandle AuthenticationType
 //AgtIsisSession EnableAuthentication, m.Object, m.Name);
 return nil
}

func(np *IsisSession) DisableAuthentication () error {
 //parameters: SessionHandle AuthenticationType
 //AgtIsisSession DisableAuthentication, m.Object, m.Name);
 return nil
}

func(np *IsisSession) IsAuthenticationEnabled () error {
 //parameters: SessionHandle AuthenticationType
 //AgtIsisSession IsAuthenticationEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisSession) SetSendPassword () error {
 //parameters: SessionHandle AuthenticationType Password
 //AgtIsisSession SetSendPassword, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetSendPassword ()(string, error) {
 //parameters: SessionHandle AuthenticationType
 //AgtIsisSession GetSendPassword
 return "", nil
}

func(np *IsisSession) SetReceivePassword () error {
 //parameters: SessionHandle AuthenticationType Count PasswordList
 //AgtIsisSession SetReceivePassword, m.Object, m.Name);
 return nil
}

func(np *IsisSession) GetReceivePassword ()(string, error) {
 //parameters: SessionHandle AuthenticationType
 //AgtIsisSession GetReceivePassword
 return "", nil
}

