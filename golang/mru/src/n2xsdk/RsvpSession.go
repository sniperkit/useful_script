package n2xsdk

type RsvpSession struct {
 Handler string
}

func(np *RsvpSession) SetSutIpAddress () error {
 //parameters: SessionHandle RemoteIpAddress
 //AgtRsvpSession SetSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetSutIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetSutIpAddress
 return "", nil
}

func(np *RsvpSession) SetTesterIpAddress () error {
 //parameters: SessionHandle LocalIpAddress
 //AgtRsvpSession SetTesterIpAddress, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetTesterIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetTesterIpAddress
 return "", nil
}

func(np *RsvpSession) SetRefreshPeriod () error {
 //parameters: SessionHandle PeriodMs
 //AgtRsvpSession SetRefreshPeriod, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetRefreshPeriod ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetRefreshPeriod
 return "", nil
}

func(np *RsvpSession) SetLTimerK () error {
 //parameters: SessionHandle LTimerK
 //AgtRsvpSession SetLTimerK, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetLTimerK ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetLTimerK
 return "", nil
}

func(np *RsvpSession) EnableHelloSignalling () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableHelloSignalling, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableHelloSignalling () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableHelloSignalling, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsHelloSignallingEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsHelloSignallingEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) SetHelloTTL () error {
 //parameters: SessionHandle HelloTTL
 //AgtRsvpSession SetHelloTTL, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetHelloTTL ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetHelloTTL
 return "", nil
}

func(np *RsvpSession) SetHelloInterval () error {
 //parameters: SessionHandle HelloInterval
 //AgtRsvpSession SetHelloInterval, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetHelloInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetHelloInterval
 return "", nil
}

func(np *RsvpSession) SetHelloTimerK () error {
 //parameters: SessionHandle HelloTimerK
 //AgtRsvpSession SetHelloTimerK, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetHelloTimerK ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetHelloTimerK
 return "", nil
}

func(np *RsvpSession) SetHelloIncludeRestartCapOption () error {
 //parameters: SessionHandle RestartCapOption
 //AgtRsvpSession SetHelloIncludeRestartCapOption, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetHelloIncludeRestartCapOption ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetHelloIncludeRestartCapOption
 return "", nil
}

func(np *RsvpSession) EnableRefreshReduction () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableRefreshReduction, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableRefreshReduction () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableRefreshReduction, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsRefreshReductionEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsRefreshReductionEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) SetRefreshReductionInitTimeout () error {
 //parameters: SessionHandle RefreshReductionInitTimeout
 //AgtRsvpSession SetRefreshReductionInitTimeout, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetRefreshReductionInitTimeout ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetRefreshReductionInitTimeout
 return "", nil
}

func(np *RsvpSession) SetRefreshReductionExpBackoff () error {
 //parameters: SessionHandle RefreshReductionExpBackoff
 //AgtRsvpSession SetRefreshReductionExpBackoff, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetRefreshReductionExpBackoff ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetRefreshReductionExpBackoff
 return "", nil
}

func(np *RsvpSession) SetRefreshReductionRetryLimit () error {
 //parameters: SessionHandle RefreshReductionRetryLimit
 //AgtRsvpSession SetRefreshReductionRetryLimit, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetRefreshReductionRetryLimit ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetRefreshReductionRetryLimit
 return "", nil
}

func(np *RsvpSession) SetRefreshReductionMaxAckDelay () error {
 //parameters: SessionHandle RefreshReductionMaxAckDelay
 //AgtRsvpSession SetRefreshReductionMaxAckDelay, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetRefreshReductionMaxAckDelay ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetRefreshReductionMaxAckDelay
 return "", nil
}

func(np *RsvpSession) EnableSrefresh () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableSrefresh, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableSrefresh () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableSrefresh, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsSrefreshEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsSrefreshEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) EnableBundle () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableBundle, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableBundle () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableBundle, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsBundleEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsBundleEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) EnableEgressRroInsertion () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableEgressRroInsertion, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableEgressRroInsertion () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableEgressRroInsertion, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsEgressRroInsertionEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsEgressRroInsertionEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) EnableEgressRroLabelInsertion () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableEgressRroLabelInsertion, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableEgressRroLabelInsertion () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableEgressRroLabelInsertion, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsEgressRroLabelInsertionEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsEgressRroLabelInsertionEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) SetEgressRroLabelStart () error {
 //parameters: SessionHandle EgressRroLabelStart
 //AgtRsvpSession SetEgressRroLabelStart, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetEgressRroLabelStart ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetEgressRroLabelStart
 return "", nil
}

func(np *RsvpSession) SetEgressRroLabelIncrement () error {
 //parameters: SessionHandle EgressRroLabelIncrement
 //AgtRsvpSession SetEgressRroLabelIncrement, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetEgressRroLabelIncrement ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetEgressRroLabelIncrement
 return "", nil
}

func(np *RsvpSession) EnableEgressRroNodeIdFlag () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableEgressRroNodeIdFlag, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableEgressRroNodeIdFlag () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableEgressRroNodeIdFlag, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsEgressRroNodeIdFlagEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsEgressRroNodeIdFlagEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) SetEgressRroNodeIdPrefix () error {
 //parameters: SessionHandle EgressRroNodeIdPrefixAddress EgressRroNodeIdPrefixLength
 //AgtRsvpSession SetEgressRroNodeIdPrefix, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetEgressRroNodeIdPrefix ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetEgressRroNodeIdPrefix
 return "", nil
}

func(np *RsvpSession) EnableGreTunnel () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableGreTunnel, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableGreTunnel () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableGreTunnel, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsGreTunnelEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsGreTunnelEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) EnableGreTunnelChecksumField () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableGreTunnelChecksumField, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableGreTunnelChecksumField () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableGreTunnelChecksumField, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsGreTunnelChecksumFieldEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsGreTunnelChecksumFieldEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) SetGreTunnelLocalAddress () error {
 //parameters: SessionHandle LocalAddr
 //AgtRsvpSession SetGreTunnelLocalAddress, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetGreTunnelLocalAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetGreTunnelLocalAddress
 return "", nil
}

func(np *RsvpSession) SetGreTunnelRemoteAddress () error {
 //parameters: SessionHandle RemoteAddr
 //AgtRsvpSession SetGreTunnelRemoteAddress, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetGreTunnelRemoteAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetGreTunnelRemoteAddress
 return "", nil
}

func(np *RsvpSession) GetNullLabelOption ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetNullLabelOption
 return "", nil
}

func(np *RsvpSession) Enable () error {
 //parameters: SessionHandle
 //AgtRsvpSession Enable, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) Disable () error {
 //parameters: SessionHandle
 //AgtRsvpSession Disable, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetState
 return "", nil
}

func(np *RsvpSession) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtRsvpSession GetStateSummary
 return "", nil
}

func(np *RsvpSession) GetTunnelStateSummary ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetTunnelStateSummary
 return "", nil
}

func(np *RsvpSession) GetLastError ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetLastError
 return "", nil
}

func(np *RsvpSession) GetEnableFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetEnableFlag
 return "", nil
}

func(np *RsvpSession) GetDetailedState ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetDetailedState
 return "", nil
}

func(np *RsvpSession) AddIngressPool () error {
 //parameters: SessionHandle
 //AgtRsvpSession AddIngressPool, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) ReplicatePool () error {
 //parameters: SessionHandle PoolHandle
 //AgtRsvpSession ReplicatePool, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) RemovePool () error {
 //parameters: SessionHandle PoolHandle
 //AgtRsvpSession RemovePool, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) ListPools ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession ListPools
 return "", nil
}

func(np *RsvpSession) ListIngressPools ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession ListIngressPools
 return "", nil
}

func(np *RsvpSession) ClearPools () error {
 //parameters: SessionHandle ClearIngressPools ClearEgressPools
 //AgtRsvpSession ClearPools, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) EnableEgress () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableEgress, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableEgress () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableEgress, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsEgressEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsEgressEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetDefaultEgressPoolHandle ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetDefaultEgressPoolHandle
 return "", nil
}

func(np *RsvpSession) EnableGracefulRestart () error {
 //parameters: SessionHandle
 //AgtRsvpSession EnableGracefulRestart, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableGracefulRestart () error {
 //parameters: SessionHandle
 //AgtRsvpSession DisableGracefulRestart, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) IsGracefulRestartEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpSession IsGracefulRestartEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GracefulRestart () error {
 //parameters: SessionHandle
 //AgtRsvpSession GracefulRestart, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) ForceGracefulRestart () error {
 //parameters: SessionHandle GracefulRestartType
 //AgtRsvpSession ForceGracefulRestart, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) SetRestartCapRestartTime () error {
 //parameters: SessionHandle RestartTime
 //AgtRsvpSession SetRestartCapRestartTime, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetRestartCapRestartTime ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetRestartCapRestartTime
 return "", nil
}

func(np *RsvpSession) SetRestartCapRecoveryTime () error {
 //parameters: SessionHandle RecoveryTime
 //AgtRsvpSession SetRestartCapRecoveryTime, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetRestartCapRecoveryTime ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetRestartCapRecoveryTime
 return "", nil
}

func(np *RsvpSession) SetGracefulRestartType () error {
 //parameters: SessionHandle GracefulRestartType
 //AgtRsvpSession SetGracefulRestartType, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetGracefulRestartType ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpSession GetGracefulRestartType
 return "", nil
}

func(np *RsvpSession) SetEgressLabelOverrideRange () error {
 //parameters: MinimumLabel MaximumLabel
 //AgtRsvpSession SetEgressLabelOverrideRange, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) GetEgressLabelOverrideRange ()(string, error) {
 //parameters: 
 //AgtRsvpSession GetEgressLabelOverrideRange
 return "", nil
}

func(np *RsvpSession) EnableAllSessions () error {
 //parameters: 
 //AgtRsvpSession EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) DisableAllSessions () error {
 //parameters: 
 //AgtRsvpSession DisableAllSessions, m.Object, m.Name);
 return nil
}

func(np *RsvpSession) SetNullLabelOption () error {
 //parameters: SessionHandle
 //AgtRsvpSession SetNullLabelOption, m.Object, m.Name);
 return nil
}

