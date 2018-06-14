package n2xsdk

type LmpSession struct {
 Handler string
}

func(np *LmpSession) NegotiateHelloConfig () error {
 //parameters: SessionHandle
 //AgtLmpSession NegotiateHelloConfig, m.Object, m.Name);
 return nil
}

func(np *LmpSession) SetLocalNodeId () error {
 //parameters: SessionHandle NodeId
 //AgtLmpSession SetLocalNodeId, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetLocalNodeId ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetLocalNodeId
 return "", nil
}

func(np *LmpSession) SetRemoteNodeId () error {
 //parameters: SessionHandle NodeId
 //AgtLmpSession SetRemoteNodeId, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetRemoteNodeId ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetRemoteNodeId
 return "", nil
}

func(np *LmpSession) SetTesterIpccAddress () error {
 //parameters: SessionHandle LocalIpccAddress
 //AgtLmpSession SetTesterIpccAddress, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetTesterIpccAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetTesterIpccAddress
 return "", nil
}

func(np *LmpSession) SetSutIpccAddress () error {
 //parameters: SessionHandle RemoteIpccAddress
 //AgtLmpSession SetSutIpccAddress, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetSutIpccAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetSutIpccAddress
 return "", nil
}

func(np *LmpSession) SetCcId () error {
 //parameters: SessionHandle CcId
 //AgtLmpSession SetCcId, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetCcId ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetCcId
 return "", nil
}

func(np *LmpSession) Enable () error {
 //parameters: SessionHandle
 //AgtLmpSession Enable, m.Object, m.Name);
 return nil
}

func(np *LmpSession) Disable () error {
 //parameters: SessionHandle
 //AgtLmpSession Disable, m.Object, m.Name);
 return nil
}

func(np *LmpSession) EnableGreTunnel () error {
 //parameters: SessionHandle
 //AgtLmpSession EnableGreTunnel, m.Object, m.Name);
 return nil
}

func(np *LmpSession) DisableGreTunnel () error {
 //parameters: SessionHandle
 //AgtLmpSession DisableGreTunnel, m.Object, m.Name);
 return nil
}

func(np *LmpSession) IsGreTunnelEnabled () error {
 //parameters: SessionHandle
 //AgtLmpSession IsGreTunnelEnabled, m.Object, m.Name);
 return nil
}

func(np *LmpSession) SetGreTunnelLocalAddress () error {
 //parameters: SessionHandle LocalAddr
 //AgtLmpSession SetGreTunnelLocalAddress, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetGreTunnelLocalAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetGreTunnelLocalAddress
 return "", nil
}

func(np *LmpSession) SetGreTunnelRemoteAddress () error {
 //parameters: SessionHandle RemoteAddr
 //AgtLmpSession SetGreTunnelRemoteAddress, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetGreTunnelRemoteAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetGreTunnelRemoteAddress
 return "", nil
}

func(np *LmpSession) EnableGreTunnelChecksumField () error {
 //parameters: SessionHandle
 //AgtLmpSession EnableGreTunnelChecksumField, m.Object, m.Name);
 return nil
}

func(np *LmpSession) DisableGreTunnelChecksumField () error {
 //parameters: SessionHandle
 //AgtLmpSession DisableGreTunnelChecksumField, m.Object, m.Name);
 return nil
}

func(np *LmpSession) IsGreTunnelChecksumFieldEnabled () error {
 //parameters: SessionHandle
 //AgtLmpSession IsGreTunnelChecksumFieldEnabled, m.Object, m.Name);
 return nil
}

func(np *LmpSession) SetConfigRetransmitInterval () error {
 //parameters: SessionHandle Interval
 //AgtLmpSession SetConfigRetransmitInterval, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetConfigRetransmitInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetConfigRetransmitInterval
 return "", nil
}

func(np *LmpSession) SetHelloInterval () error {
 //parameters: SessionHandle HelloInterval
 //AgtLmpSession SetHelloInterval, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetHelloInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetHelloInterval
 return "", nil
}

func(np *LmpSession) SetHelloDeadInterval () error {
 //parameters: SessionHandle HelloDeadInterval
 //AgtLmpSession SetHelloDeadInterval, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetHelloDeadInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetHelloDeadInterval
 return "", nil
}

func(np *LmpSession) EnableHelloConfigNegotiable () error {
 //parameters: SessionHandle
 //AgtLmpSession EnableHelloConfigNegotiable, m.Object, m.Name);
 return nil
}

func(np *LmpSession) DisableHelloConfigNegotiable () error {
 //parameters: SessionHandle
 //AgtLmpSession DisableHelloConfigNegotiable, m.Object, m.Name);
 return nil
}

func(np *LmpSession) IsHelloConfigNegotiable () error {
 //parameters: SessionHandle
 //AgtLmpSession IsHelloConfigNegotiable, m.Object, m.Name);
 return nil
}

func(np *LmpSession) SetActiveConfig () error {
 //parameters: SessionHandle
 //AgtLmpSession SetActiveConfig, m.Object, m.Name);
 return nil
}

func(np *LmpSession) SetPassiveConfig () error {
 //parameters: SessionHandle
 //AgtLmpSession SetPassiveConfig, m.Object, m.Name);
 return nil
}

func(np *LmpSession) IsActiveConfig () error {
 //parameters: SessionHandle
 //AgtLmpSession IsActiveConfig, m.Object, m.Name);
 return nil
}

func(np *LmpSession) Open () error {
 //parameters: SessionHandle
 //AgtLmpSession Open, m.Object, m.Name);
 return nil
}

func(np *LmpSession) OpenAllSessions () error {
 //parameters: 
 //AgtLmpSession OpenAllSessions, m.Object, m.Name);
 return nil
}

func(np *LmpSession) Close () error {
 //parameters: SessionHandle
 //AgtLmpSession Close, m.Object, m.Name);
 return nil
}

func(np *LmpSession) CloseAllSessions () error {
 //parameters: 
 //AgtLmpSession CloseAllSessions, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetState
 return "", nil
}

func(np *LmpSession) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtLmpSession GetStateSummary
 return "", nil
}

func(np *LmpSession) SetTeLinkCorrelationMode () error {
 //parameters: SessionHandle TeLinkCorrelationMode
 //AgtLmpSession SetTeLinkCorrelationMode, m.Object, m.Name);
 return nil
}

func(np *LmpSession) GetTeLinkCorrelationMode ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession GetTeLinkCorrelationMode
 return "", nil
}

func(np *LmpSession) AddTeLink () error {
 //parameters: SessionHandle EmulationPduHandle
 //AgtLmpSession AddTeLink, m.Object, m.Name);
 return nil
}

func(np *LmpSession) RemoveTeLink () error {
 //parameters: SessionHandle TeLinkHandle
 //AgtLmpSession RemoveTeLink, m.Object, m.Name);
 return nil
}

func(np *LmpSession) ListTeLinks ()(string, error) {
 //parameters: SessionHandle
 //AgtLmpSession ListTeLinks
 return "", nil
}

func(np *LmpSession) EnableAllSessions () error {
 //parameters: 
 //AgtLmpSession EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *LmpSession) DisableAllSessions () error {
 //parameters: 
 //AgtLmpSession DisableAllSessions, m.Object, m.Name);
 return nil
}

