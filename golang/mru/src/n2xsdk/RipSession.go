package n2xsdk

type RipSession struct {
 Handler string
}

func(np *RipSession) SetInterfaceIpAddress () error {
 //parameters: SessionHandle LocalIpAddress
 //AgtRipSession SetInterfaceIpAddress, m.Object, m.Name);
 return nil
}

func(np *RipSession) GetInterfaceIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetInterfaceIpAddress
 return "", nil
}

func(np *RipSession) SetSutIpAddress () error {
 //parameters: SessionHandle RemoteIpAddress
 //AgtRipSession SetSutIpAddress, m.Object, m.Name);
 return nil
}

func(np *RipSession) GetSutIpAddress ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetSutIpAddress
 return "", nil
}

func(np *RipSession) GetRouterId ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetRouterId
 return "", nil
}

func(np *RipSession) GetSubnetMask ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetSubnetMask
 return "", nil
}

func(np *RipSession) GetMetric ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetMetric
 return "", nil
}

func(np *RipSession) GetUpdateInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetUpdateInterval
 return "", nil
}

func(np *RipSession) GetTriggeredInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetTriggeredInterval
 return "", nil
}

func(np *RipSession) GetUpdateControl ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetUpdateControl
 return "", nil
}

func(np *RipSession) GetExpirationInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetExpirationInterval
 return "", nil
}

func(np *RipSession) GetGarbageInterval ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetGarbageInterval
 return "", nil
}

func(np *RipSession) AdvertiseAllRoutePools () error {
 //parameters: SessionHandle
 //AgtRipSession AdvertiseAllRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipSession) WithdrawAllRoutePools () error {
 //parameters: SessionHandle
 //AgtRipSession WithdrawAllRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipSession) StartFlapAll () error {
 //parameters: 
 //AgtRipSession StartFlapAll, m.Object, m.Name);
 return nil
}

func(np *RipSession) StopFlapAll () error {
 //parameters: 
 //AgtRipSession StopFlapAll, m.Object, m.Name);
 return nil
}

func(np *RipSession) EnableAllRoutePools () error {
 //parameters: 
 //AgtRipSession EnableAllRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipSession) DisableAllRoutePools () error {
 //parameters: 
 //AgtRipSession DisableAllRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipSession) IsEnabled () error {
 //parameters: SessionHandle
 //AgtRipSession IsEnabled, m.Object, m.Name);
 return nil
}

func(np *RipSession) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtRipSession GetStateSummary
 return "", nil
}

func(np *RipSession) GetLastError ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetLastError
 return "", nil
}

func(np *RipSession) GetState ()(string, error) {
 //parameters: SessionHandle
 //AgtRipSession GetState
 return "", nil
}

func(np *RipSession) SetRouterId () error {
 //parameters: SessionHandle RouterId
 //AgtRipSession SetRouterId, m.Object, m.Name);
 return nil
}

func(np *RipSession) SetSubnetMask () error {
 //parameters: SessionHandle SubnetMask
 //AgtRipSession SetSubnetMask, m.Object, m.Name);
 return nil
}

func(np *RipSession) SetMetric () error {
 //parameters: SessionHandle Metric
 //AgtRipSession SetMetric, m.Object, m.Name);
 return nil
}

func(np *RipSession) SetUpdateInterval () error {
 //parameters: SessionHandle UpdateInterval
 //AgtRipSession SetUpdateInterval, m.Object, m.Name);
 return nil
}

func(np *RipSession) SetTriggeredInterval () error {
 //parameters: SessionHandle TriggeredInterval
 //AgtRipSession SetTriggeredInterval, m.Object, m.Name);
 return nil
}

func(np *RipSession) SetUpdateControl () error {
 //parameters: SessionHandle UpdateControl
 //AgtRipSession SetUpdateControl, m.Object, m.Name);
 return nil
}

func(np *RipSession) SetExpirationInterval () error {
 //parameters: SessionHandle ExpirationInterval
 //AgtRipSession SetExpirationInterval, m.Object, m.Name);
 return nil
}

func(np *RipSession) SetGarbageInterval () error {
 //parameters: SessionHandle GarbageInterval
 //AgtRipSession SetGarbageInterval, m.Object, m.Name);
 return nil
}

func(np *RipSession) Enable () error {
 //parameters: SessionHandle
 //AgtRipSession Enable, m.Object, m.Name);
 return nil
}

func(np *RipSession) Disable () error {
 //parameters: SessionHandle
 //AgtRipSession Disable, m.Object, m.Name);
 return nil
}

func(np *RipSession) EnableAllSessions () error {
 //parameters: 
 //AgtRipSession EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *RipSession) DisableAllSessions () error {
 //parameters: 
 //AgtRipSession DisableAllSessions, m.Object, m.Name);
 return nil
}

