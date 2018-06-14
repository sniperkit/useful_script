package n2xsdk

type ContinuityCheckDataBase struct {
 Handler string
}

func(np *ContinuityCheckDataBase) Enable () error {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase Enable, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) Disable () error {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase Disable, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) IsEnabled () error {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase IsEnabled, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) EnableAutoDiscovery () error {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase EnableAutoDiscovery, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) DisableAutoDiscovery () error {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase DisableAutoDiscovery, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) IsAutoDiscoveryEnabled () error {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase IsAutoDiscoveryEnabled, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetCcdbRowId ()(string, error) {
 //parameters: DeviceHandle VlanId1 VlanId2 Level MacAddress MepId
 //AgtContinuityCheckDataBase GetCcdbRowId
 return "", nil
}

func(np *ContinuityCheckDataBase) GetCcdbRowIdList ()(string, error) {
 //parameters: DeviceHandle VlanId1 VlanId2 Level MacAddress MepId
 //AgtContinuityCheckDataBase GetCcdbRowIdList
 return "", nil
}

func(np *ContinuityCheckDataBase) AddRemoteMep () error {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase AddRemoteMep, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) ListRemoteMeps ()(string, error) {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase ListRemoteMeps
 return "", nil
}

func(np *ContinuityCheckDataBase) RemoveRemoteMep () error {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase RemoveRemoteMep, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetCcdbNumRows ()(string, error) {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase GetCcdbNumRows
 return "", nil
}

func(np *ContinuityCheckDataBase) RemoveAllRemoteMeps () error {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase RemoveAllRemoteMeps, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetRemoteMepId () error {
 //parameters: DeviceHandle RemoteMepKey MepId
 //AgtContinuityCheckDataBase SetRemoteMepId, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetRemoteMepId ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetRemoteMepId
 return "", nil
}

func(np *ContinuityCheckDataBase) SetRemoteMepMacAddress () error {
 //parameters: DeviceHandle RemoteMepKey MacAddress
 //AgtContinuityCheckDataBase SetRemoteMepMacAddress, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetRemoteMepMacAddress ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetRemoteMepMacAddress
 return "", nil
}

func(np *ContinuityCheckDataBase) SetRemoteMepMdLevel () error {
 //parameters: DeviceHandle RemoteMepKey Level
 //AgtContinuityCheckDataBase SetRemoteMepMdLevel, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetRemoteMepMdLevel ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetRemoteMepMdLevel
 return "", nil
}

func(np *ContinuityCheckDataBase) SetRemoteMepMegLevel () error {
 //parameters: DeviceHandle RemoteMepKey Level
 //AgtContinuityCheckDataBase SetRemoteMepMegLevel, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetRemoteMepMegLevel ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetRemoteMepMegLevel
 return "", nil
}

func(np *ContinuityCheckDataBase) GetRemoteMepState ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetRemoteMepState
 return "", nil
}

func(np *ContinuityCheckDataBase) GetRemoteMepLastStateChange ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetRemoteMepLastStateChange
 return "", nil
}

func(np *ContinuityCheckDataBase) GetRemoteMepRdiBit ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetRemoteMepRdiBit
 return "", nil
}

func(np *ContinuityCheckDataBase) EnableExpectedConnectivity () error {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase EnableExpectedConnectivity, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) DisableExpectedConnectivity () error {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase DisableExpectedConnectivity, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) IsExpectedConnectivityEnabled () error {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase IsExpectedConnectivityEnabled, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetExpectedConnectivity () error {
 //parameters: DeviceHandle
 //AgtContinuityCheckDataBase SetExpectedConnectivity, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetContinuityCheckInterval () error {
 //parameters: DeviceHandle RemoteMepKey CcInterval
 //AgtContinuityCheckDataBase SetContinuityCheckInterval, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetContinuityCheckInterval ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetContinuityCheckInterval
 return "", nil
}

func(np *ContinuityCheckDataBase) SetStandard () error {
 //parameters: DeviceHandle RemoteMepKey Standard
 //AgtContinuityCheckDataBase SetStandard, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetStandard ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetStandard
 return "", nil
}

func(np *ContinuityCheckDataBase) AddVlan () error {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase AddVlan, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) ListVlans ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase ListVlans
 return "", nil
}

func(np *ContinuityCheckDataBase) RemoveVlan () error {
 //parameters: DeviceHandle RemoteMepKey VlanStackIndex
 //AgtContinuityCheckDataBase RemoveVlan, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetVlanId () error {
 //parameters: DeviceHandle RemoteMepKey VlanStackIndex VlanId
 //AgtContinuityCheckDataBase SetVlanId, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetVlanId ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey VlanStackIndex
 //AgtContinuityCheckDataBase GetVlanId
 return "", nil
}

func(np *ContinuityCheckDataBase) SetMdNameFormat () error {
 //parameters: DeviceHandle RemoteMepKey MdNameFormat
 //AgtContinuityCheckDataBase SetMdNameFormat, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetMdNameString () error {
 //parameters: DeviceHandle RemoteMepKey MdName Length
 //AgtContinuityCheckDataBase SetMdNameString, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetMdNameMacAddrTwoOct () error {
 //parameters: DeviceHandle RemoteMepKey MacAddress Integer
 //AgtContinuityCheckDataBase SetMdNameMacAddrTwoOct, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetMdNameCharString () error {
 //parameters: DeviceHandle RemoteMepKey MdName Length
 //AgtContinuityCheckDataBase SetMdNameCharString, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetMdNameFormat ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetMdNameFormat
 return "", nil
}

func(np *ContinuityCheckDataBase) GetMdNameMacAddrTwoOct ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetMdNameMacAddrTwoOct
 return "", nil
}

func(np *ContinuityCheckDataBase) GetMdNameString ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetMdNameString
 return "", nil
}

func(np *ContinuityCheckDataBase) GetMdNameCharString ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetMdNameCharString
 return "", nil
}

func(np *ContinuityCheckDataBase) SetSmaNameFormat () error {
 //parameters: DeviceHandle RemoteMepKey SmaNameFormat
 //AgtContinuityCheckDataBase SetSmaNameFormat, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetSmaNamePriVid () error {
 //parameters: DeviceHandle RemoteMepKey VlanId
 //AgtContinuityCheckDataBase SetSmaNamePriVid, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetSmaNameCharString () error {
 //parameters: DeviceHandle RemoteMepKey SmaName Length
 //AgtContinuityCheckDataBase SetSmaNameCharString, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetSmaNameTwoOct () error {
 //parameters: DeviceHandle RemoteMepKey Integer
 //AgtContinuityCheckDataBase SetSmaNameTwoOct, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) SetSmaNameVpnId () error {
 //parameters: DeviceHandle RemoteMepKey VpnId
 //AgtContinuityCheckDataBase SetSmaNameVpnId, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetSmaNameFormat ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetSmaNameFormat
 return "", nil
}

func(np *ContinuityCheckDataBase) GetSmaNamePriVid ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetSmaNamePriVid
 return "", nil
}

func(np *ContinuityCheckDataBase) GetSmaNameCharString ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetSmaNameCharString
 return "", nil
}

func(np *ContinuityCheckDataBase) GetSmaNameTwoOct ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetSmaNameTwoOct
 return "", nil
}

func(np *ContinuityCheckDataBase) GetSmaNameVpnId ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetSmaNameVpnId
 return "", nil
}

func(np *ContinuityCheckDataBase) SetMegIdFormat () error {
 //parameters: DeviceHandle RemoteMepKey MegIdFormat
 //AgtContinuityCheckDataBase SetMegIdFormat, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetMegIdFormat ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetMegIdFormat
 return "", nil
}

func(np *ContinuityCheckDataBase) SetMegIdValue () error {
 //parameters: DeviceHandle RemoteMepKey MegId
 //AgtContinuityCheckDataBase SetMegIdValue, m.Object, m.Name);
 return nil
}

func(np *ContinuityCheckDataBase) GetMegIdValue ()(string, error) {
 //parameters: DeviceHandle RemoteMepKey
 //AgtContinuityCheckDataBase GetMegIdValue
 return "", nil
}

