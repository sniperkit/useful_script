package n2xsdk

type SipGatewayV4 struct {
 Handler string
}

func(np *SipGatewayV4) IsNetworkAddressTranslationEnabled () error {
 //parameters: DeviceHandle
 //AgtSipGatewayV4 IsNetworkAddressTranslationEnabled, m.Object, m.Name);
 return nil
}

func(np *SipGatewayV4) EnableNetworkAddressTranslation () error {
 //parameters: DeviceHandle
 //AgtSipGatewayV4 EnableNetworkAddressTranslation, m.Object, m.Name);
 return nil
}

func(np *SipGatewayV4) DisableNetworkAddressTranslation () error {
 //parameters: DeviceHandle
 //AgtSipGatewayV4 DisableNetworkAddressTranslation, m.Object, m.Name);
 return nil
}

func(np *SipGatewayV4) SetNatedInternaIpv4AddressRange () error {
 //parameters: DeviceHandle Ipv4Address Increment Repeat
 //AgtSipGatewayV4 SetNatedInternaIpv4AddressRange, m.Object, m.Name);
 return nil
}

func(np *SipGatewayV4) GetNatedInternaIpv4AddressRange ()(string, error) {
 //parameters: DeviceHandle
 //AgtSipGatewayV4 GetNatedInternaIpv4AddressRange
 return "", nil
}

func(np *SipGatewayV4) SetNatedExternalIpv4AddressRange () error {
 //parameters: DeviceHandle Ipv4Address Increment Repeat
 //AgtSipGatewayV4 SetNatedExternalIpv4AddressRange, m.Object, m.Name);
 return nil
}

func(np *SipGatewayV4) GetNatedExternalIpv4AddressRange ()(string, error) {
 //parameters: DeviceHandle
 //AgtSipGatewayV4 GetNatedExternalIpv4AddressRange
 return "", nil
}

func(np *SipGatewayV4) SetNatedExternalIpv4Address () error {
 //parameters: DeviceHandle Ipv4Address
 //AgtSipGatewayV4 SetNatedExternalIpv4Address, m.Object, m.Name);
 return nil
}

func(np *SipGatewayV4) GetNatedExternalIpv4Address ()(string, error) {
 //parameters: DeviceHandle
 //AgtSipGatewayV4 GetNatedExternalIpv4Address
 return "", nil
}

