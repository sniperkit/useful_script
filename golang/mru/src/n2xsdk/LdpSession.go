package n2xsdk

type LdpSession struct {
 Handler string
}

func(np *LdpSession) SetSutIpAddress () error {
 //parameters: SessionHandle RemoteIpAddress
 //AgtLdpSession SetSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetSutIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetSutIpAddress
 return "", nil
}

func(np *LdpSession) SetTesterIpAddress () error {
 //parameters: SessionHandle LocalIpAddress
 //AgtLdpSession SetTesterIpAddress, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetTesterIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetTesterIpAddress
 return "", nil
}

func(np *LdpSession) SetSubnetMask () error {
 //parameters: SessionHandle Mask
 //AgtLdpSession SetSubnetMask, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetSubnetMask ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetSubnetMask
 return "", nil
}

func(np *LdpSession) SetPeerType () error {
 //parameters: SessionHandle PeerType
 //AgtLdpSession SetPeerType, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetPeerType ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetPeerType
 return "", nil
}

func(np *LdpSession) SetHoldTime () error {
 //parameters: SessionHandle HoldTimeSec
 //AgtLdpSession SetHoldTime, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetHoldTime ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetHoldTime
 return "", nil
}

func(np *LdpSession) SetLabelSpaceId () error {
 //parameters: SessionHandle LabelSpaceId
 //AgtLdpSession SetLabelSpaceId, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetLabelSpaceId ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetLabelSpaceId
 return "", nil
}

func(np *LdpSession) SetKeepAliveTime () error {
 //parameters: SessionHandle KeepAliveTimeSec
 //AgtLdpSession SetKeepAliveTime, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetKeepAliveTime ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetKeepAliveTime
 return "", nil
}

func(np *LdpSession) SetRouterId () error {
 //parameters: SessionHandle RouterId
 //AgtLdpSession SetRouterId, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetRouterId ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetRouterId
 return "", nil
}

func(np *LdpSession) EnableAutoTesterRouterId () error {
 //parameters: SessionHandle
 //AgtLdpSession EnableAutoTesterRouterId, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableAutoTesterRouterId () error {
 //parameters: SessionHandle
 //AgtLdpSession DisableAutoTesterRouterId, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsAutoTesterRouterIdEnabled () error {
 //parameters: SessionHandle
 //AgtLdpSession IsAutoTesterRouterIdEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpSession) SetTesterRouterId () error {
 //parameters: SessionHandle RouterId
 //AgtLdpSession SetTesterRouterId, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetTesterRouterId ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetTesterRouterId
 return "", nil
}

func(np *LdpSession) SetSutRouterId () error {
 //parameters: SessionHandle SutRouterId
 //AgtLdpSession SetSutRouterId, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetSutRouterId ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetSutRouterId
 return "", nil
}

func(np *LdpSession) EnableTransportAddress () error {
 //parameters: SessionHandle
 //AgtLdpSession EnableTransportAddress, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableTransportAddress () error {
 //parameters: SessionHandle
 //AgtLdpSession DisableTransportAddress, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsTransportAddressEnabled () error {
 //parameters: SessionHandle
 //AgtLdpSession IsTransportAddressEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpSession) SetTransportAddress () error {
 //parameters: SessionHandle TransportAddress
 //AgtLdpSession SetTransportAddress, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetTransportAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetTransportAddress
 return "", nil
}

func(np *LdpSession) SetVcDirectionality () error {
 //parameters: SessionHandle VcDir
 //AgtLdpSession SetVcDirectionality, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetVcDirectionality ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetVcDirectionality
 return "", nil
}

func(np *LdpSession) SetMergeCapability () error {
 //parameters: SessionHandle MergeCap
 //AgtLdpSession SetMergeCapability, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetMergeCapability ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetMergeCapability
 return "", nil
}

func(np *LdpSession) SetLabelAdvertisementMode () error {
 //parameters: SessionHandle LAMode
 //AgtLdpSession SetLabelAdvertisementMode, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetLabelAdvertisementMode ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetLabelAdvertisementMode
 return "", nil
}

func(np *LdpSession) SetMaxPduLength () error {
 //parameters: SessionHandle MaxPduLengthBytes
 //AgtLdpSession SetMaxPduLength, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetMaxPduLength ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetMaxPduLength
 return "", nil
}

func(np *LdpSession) Enable () error {
 //parameters: SessionHandle
 //AgtLdpSession Enable, m.Object, m.Name);
 return nil
}

func(np *LdpSession) Disable () error {
 //parameters: SessionHandle
 //AgtLdpSession Disable, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetEnableFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetEnableFlag
 return "", nil
}

func(np *LdpSession) SetSessionAggregation () error {
 //parameters: SessionHandle
 //AgtLdpSession SetSessionAggregation, m.Object, m.Name);
 return nil
}

func(np *LdpSession) UnSetSessionAggregation () error {
 //parameters: SessionHandle
 //AgtLdpSession UnSetSessionAggregation, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsSessionAggregationSet () error {
 //parameters: SessionHandle
 //AgtLdpSession IsSessionAggregationSet, m.Object, m.Name);
 return nil
}

func(np *LdpSession) EnableLoopDetection () error {
 //parameters: SessionHandle
 //AgtLdpSession EnableLoopDetection, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableLoopDetection () error {
 //parameters: SessionHandle
 //AgtLdpSession DisableLoopDetection, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsLoopDetectionEnabled () error {
 //parameters: SessionHandle
 //AgtLdpSession IsLoopDetectionEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpSession) EnableLspDiscardMode () error {
 //parameters: SessionHandle
 //AgtLdpSession EnableLspDiscardMode, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableLspDiscardMode () error {
 //parameters: SessionHandle
 //AgtLdpSession DisableLspDiscardMode, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsLspDiscardModeEnabled () error {
 //parameters: SessionHandle
 //AgtLdpSession IsLspDiscardModeEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpSession) EnableNotificationOnUnsupportedFec () error {
 //parameters: SessionHandle
 //AgtLdpSession EnableNotificationOnUnsupportedFec, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableNotificationOnUnsupportedFec () error {
 //parameters: SessionHandle
 //AgtLdpSession DisableNotificationOnUnsupportedFec, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsNotificationOnUnsupportedFecEnabled () error {
 //parameters: SessionHandle
 //AgtLdpSession IsNotificationOnUnsupportedFecEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpSession) EnableNotificationOnUnsupportedTlv () error {
 //parameters: SessionHandle
 //AgtLdpSession EnableNotificationOnUnsupportedTlv, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableNotificationOnUnsupportedTlv () error {
 //parameters: SessionHandle
 //AgtLdpSession DisableNotificationOnUnsupportedTlv, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsNotificationOnUnsupportedTlvEnabled () error {
 //parameters: SessionHandle
 //AgtLdpSession IsNotificationOnUnsupportedTlvEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpSession) EnableFtSessionTlvFlag () error {
 //parameters: SessionHandle Flag
 //AgtLdpSession EnableFtSessionTlvFlag, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableFtSessionTlvFlag () error {
 //parameters: SessionHandle Flag
 //AgtLdpSession DisableFtSessionTlvFlag, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsFtSessionTlvFlagEnabled () error {
 //parameters: SessionHandle Flag
 //AgtLdpSession IsFtSessionTlvFlagEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpSession) EnableMd5Authentication () error {
 //parameters: SessionHandle
 //AgtLdpSession EnableMd5Authentication, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableMd5Authentication () error {
 //parameters: SessionHandle
 //AgtLdpSession DisableMd5Authentication, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsMd5AuthenticationEnabled () error {
 //parameters: SessionHandle
 //AgtLdpSession IsMd5AuthenticationEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpSession) SetMd5AuthenticationKey () error {
 //parameters: SessionHandle Md5Key
 //AgtLdpSession SetMd5AuthenticationKey, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetMd5AuthenticationKey ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetMd5AuthenticationKey
 return "", nil
}

func(np *LdpSession) EnableTLV () error {
 //parameters: SessionHandle TlvType
 //AgtLdpSession EnableTLV, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableTLV () error {
 //parameters: SessionHandle TlvType
 //AgtLdpSession DisableTLV, m.Object, m.Name);
 return nil
}

func(np *LdpSession) IsTlvEnabled () error {
 //parameters: SessionHandle TlvType
 //AgtLdpSession IsTlvEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpSession) SetFtSessionTlvTimer () error {
 //parameters: SessionHandle Timer TimerValue
 //AgtLdpSession SetFtSessionTlvTimer, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetFtSessionTlvTimer ()(string, error) {
 //parameters: SessionHandle Timer
 //AgtLdpSession GetFtSessionTlvTimer
 return "", nil
}

func(np *LdpSession) EnableConformance () error {
 //parameters: SessionHandle
 //AgtLdpSession EnableConformance, m.Object, m.Name);
 return nil
}

func(np *LdpSession) GetLastError ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetLastError
 return "", nil
}

func(np *LdpSession) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetState
 return "", nil
}

func(np *LdpSession) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtLdpSession GetStateSummary
 return "", nil
}

func(np *LdpSession) GetLspStateSummaryBySession ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpSession GetLspStateSummaryBySession
 return "", nil
}

func(np *LdpSession) GetLspStateSummary ()(string, error) {
 //parameters: 
 //AgtLdpSession GetLspStateSummary
 return "", nil
}

func(np *LdpSession) SimulateRestart () error {
 //parameters: SessionHandle CloseType DownTime
 //AgtLdpSession SimulateRestart, m.Object, m.Name);
 return nil
}

func(np *LdpSession) EnableAllSessions () error {
 //parameters: 
 //AgtLdpSession EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *LdpSession) DisableAllSessions () error {
 //parameters: 
 //AgtLdpSession DisableAllSessions, m.Object, m.Name);
 return nil
}

