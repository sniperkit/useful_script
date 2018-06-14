package n2xsdk

type ddressesIp struct {
 Handler string
}

func(np *ddressesIp) SetLocalIpAddressIncrementingRange () error {
 //parameters: DeviceHandle AddressFamily FirstIpAddress IpAddressPrefixLength IpAddressIncrement IpAddressRepeat
 //AgtAddressesIp SetLocalIpAddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) GetLocalIpAddressIncrementingRange ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetLocalIpAddressIncrementingRange
 return "", nil
}

func(np *ddressesIp) SetLocalIpAddressSubRangeOffsets () error {
 //parameters: DeviceHandle AddressFamily FirstIpAddress MsbOffsetList
 //AgtAddressesIp SetLocalIpAddressSubRangeOffsets, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) GetLocalIpAddressSubRangeOffsets ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetLocalIpAddressSubRangeOffsets
 return "", nil
}

func(np *ddressesIp) SetLocalIpAddressSubRange () error {
 //parameters: DeviceHandle AddressFamily SubRangeInstance IpAddressIncrement IpAddressCount IpAddressRepeat
 //AgtAddressesIp SetLocalIpAddressSubRange, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) GetLocalIpAddressSubRange ()(string, error) {
 //parameters: DeviceHandle AddressFamily SubRangeInstance
 //AgtAddressesIp GetLocalIpAddressSubRange
 return "", nil
}

func(np *ddressesIp) SetLocalIpAddressList ()(string, error) {
 //parameters: DeviceHandle AddressFamily IpAddressList
 //AgtAddressesIp SetLocalIpAddressList
 return "", nil
}

func(np *ddressesIp) GetLocalIpAddressList ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetLocalIpAddressList
 return "", nil
}

func(np *ddressesIp) SetLocalIpAddress () error {
 //parameters: DeviceHandle AddressFamily IpAddress
 //AgtAddressesIp SetLocalIpAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) GetLocalIpAddress ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetLocalIpAddress
 return "", nil
}

func(np *ddressesIp) EnableRemoteIpAddress () error {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp EnableRemoteIpAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) DisableRemoteIpAddress () error {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp DisableRemoteIpAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) IsRemoteIpAddressEnabled () error {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp IsRemoteIpAddressEnabled, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) SetRemoteIpAddressIncrementingRange () error {
 //parameters: DeviceHandle AddressFamily FirstIpAddress IpAddressPrefixLength IpAddressIncrement IpAddressRepeat
 //AgtAddressesIp SetRemoteIpAddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) GetRemoteIpAddressIncrementingRange ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetRemoteIpAddressIncrementingRange
 return "", nil
}

func(np *ddressesIp) SetRemoteIpAddressList ()(string, error) {
 //parameters: DeviceHandle AddressFamily IpAddressList
 //AgtAddressesIp SetRemoteIpAddressList
 return "", nil
}

func(np *ddressesIp) GetRemoteIpAddressList ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetRemoteIpAddressList
 return "", nil
}

func(np *ddressesIp) SetRemoteIpAddress () error {
 //parameters: DeviceHandle AddressFamily IpAddress
 //AgtAddressesIp SetRemoteIpAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) GetRemoteIpAddress ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetRemoteIpAddress
 return "", nil
}

func(np *ddressesIp) GetLinkLocalIpv6AddressModifierType ()(string, error) {
 //parameters: DeviceHandle
 //AgtAddressesIp GetLinkLocalIpv6AddressModifierType
 return "", nil
}

func(np *ddressesIp) GetRemoteLinkLocalIpv6AddressModifierType ()(string, error) {
 //parameters: DeviceHandle
 //AgtAddressesIp GetRemoteLinkLocalIpv6AddressModifierType
 return "", nil
}

func(np *ddressesIp) GetLinkLocalIpv6AddressIncrementingRange ()(string, error) {
 //parameters: DeviceHandle
 //AgtAddressesIp GetLinkLocalIpv6AddressIncrementingRange
 return "", nil
}

func(np *ddressesIp) GetLinkLocalIpv6AddressList ()(string, error) {
 //parameters: DeviceHandle
 //AgtAddressesIp GetLinkLocalIpv6AddressList
 return "", nil
}

func(np *ddressesIp) GetLinkLocalIpv6Address ()(string, error) {
 //parameters: DeviceHandle
 //AgtAddressesIp GetLinkLocalIpv6Address
 return "", nil
}

func(np *ddressesIp) GetRemoteLinkLocalIpv6AddressIncrementingRange ()(string, error) {
 //parameters: DeviceHandle
 //AgtAddressesIp GetRemoteLinkLocalIpv6AddressIncrementingRange
 return "", nil
}

func(np *ddressesIp) GetRemoteLinkLocalIpv6AddressList ()(string, error) {
 //parameters: DeviceHandle
 //AgtAddressesIp GetRemoteLinkLocalIpv6AddressList
 return "", nil
}

func(np *ddressesIp) GetRemoteLinkLocalIpv6Address ()(string, error) {
 //parameters: DeviceHandle
 //AgtAddressesIp GetRemoteLinkLocalIpv6Address
 return "", nil
}

func(np *ddressesIp) EnableGatewayIpAddress () error {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp EnableGatewayIpAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) DisableGatewayIpAddress () error {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp DisableGatewayIpAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) IsGatewayIpAddressEnabled () error {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp IsGatewayIpAddressEnabled, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) SetGatewayIpAddressIncrementingRange () error {
 //parameters: DeviceHandle AddressFamily FirstIpAddress IpAddressPrefixLength IpAddressIncrement IpAddressRepeat
 //AgtAddressesIp SetGatewayIpAddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) GetGatewayIpAddressIncrementingRange ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetGatewayIpAddressIncrementingRange
 return "", nil
}

func(np *ddressesIp) SetGatewayIpAddressList ()(string, error) {
 //parameters: DeviceHandle AddressFamily IpAddressList
 //AgtAddressesIp SetGatewayIpAddressList
 return "", nil
}

func(np *ddressesIp) GetGatewayIpAddressList ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetGatewayIpAddressList
 return "", nil
}

func(np *ddressesIp) SetGatewayIpAddress () error {
 //parameters: DeviceHandle AddressFamily IpAddress
 //AgtAddressesIp SetGatewayIpAddress, m.Object, m.Name);
 return nil
}

func(np *ddressesIp) GetGatewayIpAddress ()(string, error) {
 //parameters: DeviceHandle AddressFamily
 //AgtAddressesIp GetGatewayIpAddress
 return "", nil
}

