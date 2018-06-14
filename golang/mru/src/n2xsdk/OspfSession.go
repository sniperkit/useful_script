package n2xsdk

type OspfSession struct {
 Handler string
}

func(np *OspfSession) SetAreaId () error {
 //parameters: SessionHandle AreaId
 //AgtOspfSession SetAreaId, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetAreaId ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetAreaId
 return "", nil
}

func(np *OspfSession) SetNetworkType () error {
 //parameters: SessionHandle NetType
 //AgtOspfSession SetNetworkType, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetNetworkType ()(string, error) {
 //parameters: 
 //AgtOspfSession GetNetworkType
 return "", nil
}

func(np *OspfSession) SetMaxLsasPerPacket () error {
 //parameters: SessionHandle MaxLsas
 //AgtOspfSession SetMaxLsasPerPacket, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetMaxLsasPerPacket ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetMaxLsasPerPacket
 return "", nil
}

func(np *OspfSession) EnableNeighborDiscovery () error {
 //parameters: SessionHandle
 //AgtOspfSession EnableNeighborDiscovery, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableNeighborDiscovery () error {
 //parameters: SessionHandle
 //AgtOspfSession DisableNeighborDiscovery, m.Object, m.Name);
 return nil
}

func(np *OspfSession) IsNeighborDiscoveryEnabled () error {
 //parameters: SessionHandle
 //AgtOspfSession IsNeighborDiscoveryEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfSession) EnableLsaDiscardMode () error {
 //parameters: SessionHandle
 //AgtOspfSession EnableLsaDiscardMode, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableLsaDiscardMode () error {
 //parameters: SessionHandle
 //AgtOspfSession DisableLsaDiscardMode, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetLsaDiscardModeFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetLsaDiscardModeFlag
 return "", nil
}

func(np *OspfSession) EnableAutoLsaRefreshMode () error {
 //parameters: SessionHandle
 //AgtOspfSession EnableAutoLsaRefreshMode, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableAutoLsaRefreshMode () error {
 //parameters: SessionHandle
 //AgtOspfSession DisableAutoLsaRefreshMode, m.Object, m.Name);
 return nil
}

func(np *OspfSession) IsLsaRefreshModeEnabled () error {
 //parameters: SessionHandle
 //AgtOspfSession IsLsaRefreshModeEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfSession) EnableGracefulRestart () error {
 //parameters: SessionHandle
 //AgtOspfSession EnableGracefulRestart, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableGracefulRestart () error {
 //parameters: SessionHandle
 //AgtOspfSession DisableGracefulRestart, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetNeighborState ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetNeighborState
 return "", nil
}

func(np *OspfSession) GetSpecifiedNeighborState ()(string, error) {
 //parameters: SessionHandle NeighborHandle
 //AgtOspfSession GetSpecifiedNeighborState
 return "", nil
}

func(np *OspfSession) GetInterfaceState ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetInterfaceState
 return "", nil
}

func(np *OspfSession) GetPortStateSummary ()(string, error) {
 //parameters: PortHandle
 //AgtOspfSession GetPortStateSummary
 return "", nil
}

func(np *OspfSession) GetPortOspfv3StateSummary ()(string, error) {
 //parameters: PortHandle
 //AgtOspfSession GetPortOspfv3StateSummary
 return "", nil
}

func(np *OspfSession) GetPortOspfv2StateSummary ()(string, error) {
 //parameters: PortHandle
 //AgtOspfSession GetPortOspfv2StateSummary
 return "", nil
}

func(np *OspfSession) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtOspfSession GetStateSummary
 return "", nil
}

func(np *OspfSession) GetOspfv3StateSummary ()(string, error) {
 //parameters: 
 //AgtOspfSession GetOspfv3StateSummary
 return "", nil
}

func(np *OspfSession) GetOspfv2StateSummary ()(string, error) {
 //parameters: 
 //AgtOspfSession GetOspfv2StateSummary
 return "", nil
}

func(np *OspfSession) GetLastError ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetLastError
 return "", nil
}

func(np *OspfSession) AddNeighbor () error {
 //parameters: SessionHandle
 //AgtOspfSession AddNeighbor, m.Object, m.Name);
 return nil
}

func(np *OspfSession) RemoveNeighbor () error {
 //parameters: SessionHandle NeighborHandle
 //AgtOspfSession RemoveNeighbor, m.Object, m.Name);
 return nil
}

func(np *OspfSession) RemoveAllNeighbors () error {
 //parameters: SessionHandle
 //AgtOspfSession RemoveAllNeighbors, m.Object, m.Name);
 return nil
}

func(np *OspfSession) ListNeighbors ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession ListNeighbors
 return "", nil
}

func(np *OspfSession) GetLsaDatabase ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetLsaDatabase
 return "", nil
}

func(np *OspfSession) GetSessionRouter ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetSessionRouter
 return "", nil
}

func(np *OspfSession) GetVersion ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetVersion
 return "", nil
}

func(np *OspfSession) GetLinkLocalInterfaceAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetLinkLocalInterfaceAddress
 return "", nil
}

func(np *OspfSession) GetDrRouterId ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetDrRouterId
 return "", nil
}

func(np *OspfSession) GetBdrRouterId ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetBdrRouterId
 return "", nil
}

func(np *OspfSession) GetInternalMessageExchangeFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetInternalMessageExchangeFlag
 return "", nil
}

func(np *OspfSession) GetLsaReAssertMode ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetLsaReAssertMode
 return "", nil
}

func(np *OspfSession) SetLsaReAssertMode () error {
 //parameters: SessionHandle LsaReAssertMode
 //AgtOspfSession SetLsaReAssertMode, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetHostRouteFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetHostRouteFlag
 return "", nil
}

func(np *OspfSession) GetAuthenticationKey ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetAuthenticationKey
 return "", nil
}

func(np *OspfSession) GetKeyIdentifier ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetKeyIdentifier
 return "", nil
}

func(np *OspfSession) IsGreTunnelChecksumFieldEnabled () error {
 //parameters: SessionHandle
 //AgtOspfSession IsGreTunnelChecksumFieldEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfSession) Restart () error {
 //parameters: 
 //AgtOspfSession Restart, m.Object, m.Name);
 return nil
}

func(np *OspfSession) RestartAllSessions () error {
 //parameters: 
 //AgtOspfSession RestartAllSessions, m.Object, m.Name);
 return nil
}

func(np *OspfSession) EnableBehaviorOverride () error {
 //parameters: SessionHandle BehaviorOverride
 //AgtOspfSession EnableBehaviorOverride, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableBehaviorOverride () error {
 //parameters: SessionHandle BehaviorOverride
 //AgtOspfSession DisableBehaviorOverride, m.Object, m.Name);
 return nil
}

func(np *OspfSession) EnableAllSessions () error {
 //parameters: 
 //AgtOspfSession EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *OspfSession) EnableAllOspfv2Sessions () error {
 //parameters: 
 //AgtOspfSession EnableAllOspfv2Sessions, m.Object, m.Name);
 return nil
}

func(np *OspfSession) EnableAllOspfv3Sessions () error {
 //parameters: 
 //AgtOspfSession EnableAllOspfv3Sessions, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableAllSessions () error {
 //parameters: 
 //AgtOspfSession DisableAllSessions, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableAllOspfv2Sessions () error {
 //parameters: 
 //AgtOspfSession DisableAllOspfv2Sessions, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableAllOspfv3Sessions () error {
 //parameters: 
 //AgtOspfSession DisableAllOspfv3Sessions, m.Object, m.Name);
 return nil
}

func(np *OspfSession) SetInterfaceAddress () error {
 //parameters: 
 //AgtOspfSession SetInterfaceAddress, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetInterfaceAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetInterfaceAddress
 return "", nil
}

func(np *OspfSession) Enable () error {
 //parameters: 
 //AgtOspfSession Enable, m.Object, m.Name);
 return nil
}

func(np *OspfSession) Disable () error {
 //parameters: 
 //AgtOspfSession Disable, m.Object, m.Name);
 return nil
}

func(np *OspfSession) IsEnabled () error {
 //parameters: SessionHandle
 //AgtOspfSession IsEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfSession) SetInterfaceParameter () error {
 //parameters: 
 //AgtOspfSession SetInterfaceParameter, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetInterfaceParameter ()(string, error) {
 //parameters: SessionHandle Parameter
 //AgtOspfSession GetInterfaceParameter
 return "", nil
}

func(np *OspfSession) SetInternalMessageExchangeFlag () error {
 //parameters: 
 //AgtOspfSession SetInternalMessageExchangeFlag, m.Object, m.Name);
 return nil
}

func(np *OspfSession) EnableHostRoute () error {
 //parameters: 
 //AgtOspfSession EnableHostRoute, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableHostRoute () error {
 //parameters: 
 //AgtOspfSession DisableHostRoute, m.Object, m.Name);
 return nil
}

func(np *OspfSession) IsGreTunnelEnabled () error {
 //parameters: SessionHandle
 //AgtOspfSession IsGreTunnelEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfSession) EnableGreTunnel () error {
 //parameters: 
 //AgtOspfSession EnableGreTunnel, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableGreTunnel () error {
 //parameters: 
 //AgtOspfSession DisableGreTunnel, m.Object, m.Name);
 return nil
}

func(np *OspfSession) SetGreTunnelLocalAddress () error {
 //parameters: 
 //AgtOspfSession SetGreTunnelLocalAddress, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetGreTunnelLocalAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetGreTunnelLocalAddress
 return "", nil
}

func(np *OspfSession) SetGreTunnelRemoteAddress () error {
 //parameters: 
 //AgtOspfSession SetGreTunnelRemoteAddress, m.Object, m.Name);
 return nil
}

func(np *OspfSession) GetGreTunnelRemoteAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtOspfSession GetGreTunnelRemoteAddress
 return "", nil
}

func(np *OspfSession) EnableGreTunnelChecksumField () error {
 //parameters: 
 //AgtOspfSession EnableGreTunnelChecksumField, m.Object, m.Name);
 return nil
}

func(np *OspfSession) DisableGreTunnelChecksumField () error {
 //parameters: 
 //AgtOspfSession DisableGreTunnelChecksumField, m.Object, m.Name);
 return nil
}

func(np *OspfSession) IsGracefulRestartEnabled () error {
 //parameters: SessionHandle
 //AgtOspfSession IsGracefulRestartEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfSession) IsBehaviorOverrideEnabled () error {
 //parameters: SessionHandle BehaviorOverride
 //AgtOspfSession IsBehaviorOverrideEnabled, m.Object, m.Name);
 return nil
}

