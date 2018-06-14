package n2xsdk

type RipngSession struct {
 Handler string
}

func(np *RipngSession) SetInterfaceIpAddress () error {
 //parameters: SessionHandle LocalIpAddress
 //AgtRipngSession SetInterfaceIpAddress, m.Object, m.Name);
 return nil
}

func(np *RipngSession) SetSutIpAddress () error {
 //parameters: SessionHandle RemoteIpAddress
 //AgtRipngSession SetSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *RipngSession) SetRouterId () error {
 //parameters: SessionHandle RouterId
 //AgtRipngSession SetRouterId, m.Object, m.Name);
 return nil
}

func(np *RipngSession) SetSubnetMask () error {
 //parameters: SessionHandle SubnetMask
 //AgtRipngSession SetSubnetMask, m.Object, m.Name);
 return nil
}

func(np *RipngSession) SetMetric () error {
 //parameters: SessionHandle Metric
 //AgtRipngSession SetMetric, m.Object, m.Name);
 return nil
}

func(np *RipngSession) SetUpdateInterval () error {
 //parameters: SessionHandle UpdateInterval
 //AgtRipngSession SetUpdateInterval, m.Object, m.Name);
 return nil
}

func(np *RipngSession) SetTriggeredInterval () error {
 //parameters: SessionHandle TriggeredInterval
 //AgtRipngSession SetTriggeredInterval, m.Object, m.Name);
 return nil
}

func(np *RipngSession) SetUpdateControl () error {
 //parameters: SessionHandle UpdateControl
 //AgtRipngSession SetUpdateControl, m.Object, m.Name);
 return nil
}

func(np *RipngSession) SetExpirationInterval () error {
 //parameters: SessionHandle ExpirationInterval
 //AgtRipngSession SetExpirationInterval, m.Object, m.Name);
 return nil
}

func(np *RipngSession) SetGarbageInterval () error {
 //parameters: SessionHandle GarbageInterval
 //AgtRipngSession SetGarbageInterval, m.Object, m.Name);
 return nil
}

func(np *RipngSession) Enable () error {
 //parameters: SessionHandle
 //AgtRipngSession Enable, m.Object, m.Name);
 return nil
}

func(np *RipngSession) Disable () error {
 //parameters: SessionHandle
 //AgtRipngSession Disable, m.Object, m.Name);
 return nil
}

func(np *RipngSession) EnableAllSessions () error {
 //parameters: 
 //AgtRipngSession EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *RipngSession) DisableAllSessions () error {
 //parameters: 
 //AgtRipngSession DisableAllSessions, m.Object, m.Name);
 return nil
}

func(np *RipngSession) GetInterfaceIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetInterfaceIpAddress
 return "", nil
}

func(np *RipngSession) GetSutIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetSutIpAddress
 return "", nil
}

func(np *RipngSession) GetRouterId ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetRouterId
 return "", nil
}

func(np *RipngSession) GetSubnetMask ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetSubnetMask
 return "", nil
}

func(np *RipngSession) GetMetric ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetMetric
 return "", nil
}

func(np *RipngSession) GetUpdateInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetUpdateInterval
 return "", nil
}

func(np *RipngSession) GetTriggeredInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetTriggeredInterval
 return "", nil
}

func(np *RipngSession) GetUpdateControl ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetUpdateControl
 return "", nil
}

func(np *RipngSession) GetExpirationInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetExpirationInterval
 return "", nil
}

func(np *RipngSession) GetGarbageInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetGarbageInterval
 return "", nil
}

func(np *RipngSession) AdvertiseAllRoutePools () error {
 //parameters: SessionHandle
 //AgtRipngSession AdvertiseAllRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipngSession) WithdrawAllRoutePools () error {
 //parameters: SessionHandle
 //AgtRipngSession WithdrawAllRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipngSession) StartFlapAll () error {
 //parameters: 
 //AgtRipngSession StartFlapAll, m.Object, m.Name);
 return nil
}

func(np *RipngSession) StopFlapAll () error {
 //parameters: 
 //AgtRipngSession StopFlapAll, m.Object, m.Name);
 return nil
}

func(np *RipngSession) EnableAllRoutePools () error {
 //parameters: 
 //AgtRipngSession EnableAllRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipngSession) DisableAllRoutePools () error {
 //parameters: 
 //AgtRipngSession DisableAllRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipngSession) IsEnabled () error {
 //parameters: SessionHandle
 //AgtRipngSession IsEnabled, m.Object, m.Name);
 return nil
}

func(np *RipngSession) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtRipngSession GetStateSummary
 return "", nil
}

func(np *RipngSession) GetLastError ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetLastError
 return "", nil
}

func(np *RipngSession) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngSession GetState
 return "", nil
}

