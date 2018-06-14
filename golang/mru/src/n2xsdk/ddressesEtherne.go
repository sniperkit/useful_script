package n2xsdk

type ddressesEtherne struct {
 Handler string
}

func(np *ddressesEtherne) SetEthernetFraming () error {
 //parameters: DeviceHandle HeaderInstance FramingType
 //AgtAddressesEthernet SetEthernetFraming, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetEthernetFraming ()(string, error) {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet GetEthernetFraming
 return "", nil
}

func(np *ddressesEtherne) AddVlanTag () error {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet AddVlanTag, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) RemoveVlanTag () error {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance
 //AgtAddressesEthernet RemoveVlanTag, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetNumberOfVlanTags ()(string, error) {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet GetNumberOfVlanTags
 return "", nil
}

func(np *ddressesEtherne) SetLocalMacAddressIncrementingRange () error {
 //parameters: DeviceHandle HeaderInstance FirstMacAddress MacAddressIncrement MacAddressRepeat
 //AgtAddressesEthernet SetLocalMacAddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetLocalMacAddressIncrementingRange ()(string, error) {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet GetLocalMacAddressIncrementingRange
 return "", nil
}

func(np *ddressesEtherne) SetLocalMacAddressList ()(string, error) {
 //parameters: DeviceHandle HeaderInstance MacAddressList
 //AgtAddressesEthernet SetLocalMacAddressList
 return "", nil
}

func(np *ddressesEtherne) GetLocalMacAddressList ()(string, error) {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet GetLocalMacAddressList
 return "", nil
}

func(np *ddressesEtherne) SetLocalMacAddress () error {
 //parameters: DeviceHandle HeaderInstance MacAddress
 //AgtAddressesEthernet SetLocalMacAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetLocalMacAddress ()(string, error) {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet GetLocalMacAddress
 return "", nil
}

func(np *ddressesEtherne) SetLocalMacAddressSubRangeOffsets () error {
 //parameters: DeviceHandle FirstMacAddress MsbOffsetList
 //AgtAddressesEthernet SetLocalMacAddressSubRangeOffsets, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetLocalMacAddressSubRangeOffsets ()(string, error) {
 //parameters: DeviceHandle
 //AgtAddressesEthernet GetLocalMacAddressSubRangeOffsets
 return "", nil
}

func(np *ddressesEtherne) SetLocalMacAddressSubRange () error {
 //parameters: DeviceHandle SubRangeInstance MacAddressIncrement MacAddressCount MacAddressRepeat
 //AgtAddressesEthernet SetLocalMacAddressSubRange, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetLocalMacAddressSubRange ()(string, error) {
 //parameters: DeviceHandle SubRangeInstance
 //AgtAddressesEthernet GetLocalMacAddressSubRange
 return "", nil
}

func(np *ddressesEtherne) EnableRemoteMacAddress () error {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet EnableRemoteMacAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) DisableRemoteMacAddress () error {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet DisableRemoteMacAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) IsRemoteMacAddressEnabled () error {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet IsRemoteMacAddressEnabled, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) SetRemoteMacAddressIncrementingRange () error {
 //parameters: DeviceHandle HeaderInstance FirstMacAddress MacAddressIncrement MacAddressRepeat
 //AgtAddressesEthernet SetRemoteMacAddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetRemoteMacAddressIncrementingRange ()(string, error) {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet GetRemoteMacAddressIncrementingRange
 return "", nil
}

func(np *ddressesEtherne) SetRemoteMacAddressList ()(string, error) {
 //parameters: DeviceHandle HeaderInstance MacAddressList
 //AgtAddressesEthernet SetRemoteMacAddressList
 return "", nil
}

func(np *ddressesEtherne) GetRemoteMacAddressList ()(string, error) {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet GetRemoteMacAddressList
 return "", nil
}

func(np *ddressesEtherne) SetRemoteMacAddress () error {
 //parameters: DeviceHandle HeaderInstance MacAddress
 //AgtAddressesEthernet SetRemoteMacAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetRemoteMacAddress ()(string, error) {
 //parameters: DeviceHandle HeaderInstance
 //AgtAddressesEthernet GetRemoteMacAddress
 return "", nil
}

func(np *ddressesEtherne) SetVlanTagProtocolId () error {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance ProtocolId
 //AgtAddressesEthernet SetVlanTagProtocolId, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetVlanTagProtocolId ()(string, error) {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance
 //AgtAddressesEthernet GetVlanTagProtocolId
 return "", nil
}

func(np *ddressesEtherne) SetVlanIdIncrementingRange () error {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance FirstVlanId VlanIdIncrement VlanIdRepeat
 //AgtAddressesEthernet SetVlanIdIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetVlanIdIncrementingRange ()(string, error) {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance
 //AgtAddressesEthernet GetVlanIdIncrementingRange
 return "", nil
}

func(np *ddressesEtherne) SetVlanIdIncrementingRangeAndCount () error {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance FirstVlanId VlanIdIncrement VlanIdRepeat VlanIdCount
 //AgtAddressesEthernet SetVlanIdIncrementingRangeAndCount, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetVlanIdIncrementingRangeAndCount ()(string, error) {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance
 //AgtAddressesEthernet GetVlanIdIncrementingRangeAndCount
 return "", nil
}

func(np *ddressesEtherne) SetVlanIdList ()(string, error) {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance VlanIdList
 //AgtAddressesEthernet SetVlanIdList
 return "", nil
}

func(np *ddressesEtherne) GetVlanIdList ()(string, error) {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance
 //AgtAddressesEthernet GetVlanIdList
 return "", nil
}

func(np *ddressesEtherne) SetVlanId () error {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance VlanId
 //AgtAddressesEthernet SetVlanId, m.Object, m.Name);
 return nil
}

func(np *ddressesEtherne) GetVlanId ()(string, error) {
 //parameters: DeviceHandle HeaderInstance VlanTagInstance
 //AgtAddressesEthernet GetVlanId
 return "", nil
}

