package n2xsdk

type LacPool struct {
 Handler string
}

func(np *LacPool) SetLacIpv4AddressIncrementingRange () error {
 //parameters: SessionPoolHandle Ipv4Address PrefixLength Increment Repeat
 //AgtLacPool SetLacIpv4AddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetLacIpv4AddressIncrementingRange ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetLacIpv4AddressIncrementingRange
 return "", nil
}

func(np *LacPool) SetLacIpv4AddressList ()(string, error) {
 //parameters: SessionPoolHandle Ipv4Address
 //AgtLacPool SetLacIpv4AddressList
 return "", nil
}

func(np *LacPool) GetLacIpv4AddressList ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetLacIpv4AddressList
 return "", nil
}

func(np *LacPool) SetLacIpv4Address () error {
 //parameters: SessionPoolHandle Ipv4Address
 //AgtLacPool SetLacIpv4Address, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetLacIpv4Address ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetLacIpv4Address
 return "", nil
}

func(np *LacPool) SetLnsIpv4AddressIncrementingRange () error {
 //parameters: SessionPoolHandle Ipv4Address PrefixLength Increment Repeat
 //AgtLacPool SetLnsIpv4AddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetLnsIpv4AddressIncrementingRange ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetLnsIpv4AddressIncrementingRange
 return "", nil
}

func(np *LacPool) SetLnsIpv4AddressList ()(string, error) {
 //parameters: SessionPoolHandle Ipv4Address
 //AgtLacPool SetLnsIpv4AddressList
 return "", nil
}

func(np *LacPool) GetLnsIpv4AddressList ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetLnsIpv4AddressList
 return "", nil
}

func(np *LacPool) SetLnsIpv4Address () error {
 //parameters: SessionPoolHandle Ipv4Address
 //AgtLacPool SetLnsIpv4Address, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetLnsIpv4Address ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetLnsIpv4Address
 return "", nil
}

func(np *LacPool) EnableCallingNumberAvp () error {
 //parameters: SessionPoolHandle
 //AgtLacPool EnableCallingNumberAvp, m.Object, m.Name);
 return nil
}

func(np *LacPool) DisableCallingNumberAvp () error {
 //parameters: SessionPoolHandle
 //AgtLacPool DisableCallingNumberAvp, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsCallingNumberAvpEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLacPool IsCallingNumberAvpEnabled, m.Object, m.Name);
 return nil
}

func(np *LacPool) SetCallingNumberAvpValue () error {
 //parameters: SessionPoolHandle CallingNumberAvpValue
 //AgtLacPool SetCallingNumberAvpValue, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetCallingNumberAvpValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetCallingNumberAvpValue
 return "", nil
}

func(np *LacPool) EnableL2tpHeaderLengthBit () error {
 //parameters: SessionPoolHandle
 //AgtLacPool EnableL2tpHeaderLengthBit, m.Object, m.Name);
 return nil
}

func(np *LacPool) DisableL2tpHeaderLengthBit () error {
 //parameters: SessionPoolHandle
 //AgtLacPool DisableL2tpHeaderLengthBit, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsL2tpHeaderLengthBitEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLacPool IsL2tpHeaderLengthBitEnabled, m.Object, m.Name);
 return nil
}

func(np *LacPool) EnableProxyAuthentication () error {
 //parameters: SessionPoolHandle
 //AgtLacPool EnableProxyAuthentication, m.Object, m.Name);
 return nil
}

func(np *LacPool) DisableProxyAuthentication () error {
 //parameters: SessionPoolHandle
 //AgtLacPool DisableProxyAuthentication, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsProxyAuthenticationEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLacPool IsProxyAuthenticationEnabled, m.Object, m.Name);
 return nil
}

func(np *LacPool) SetProxyAuthenType () error {
 //parameters: SessionPoolHandle ProxyAuthenticationType
 //AgtLacPool SetProxyAuthenType, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetProxyAuthenType ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenType
 return "", nil
}

func(np *LacPool) SetProxyAuthenName () error {
 //parameters: SessionPoolHandle ProxyAuthenName
 //AgtLacPool SetProxyAuthenName, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetProxyAuthenName ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenName
 return "", nil
}

func(np *LacPool) GetProxyAuthenNameValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenNameValue
 return "", nil
}

func(np *LacPool) SetProxyAuthenticationNameModifierIncrementor () error {
 //parameters: SessionPoolHandle ProxyAuthenticationNameModifier Count Increment Repeat
 //AgtLacPool SetProxyAuthenticationNameModifierIncrementor, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetProxyAuthenticationNameModifierIncrementor ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenticationNameModifierIncrementor
 return "", nil
}

func(np *LacPool) SetProxyAuthenticationNameModifierList ()(string, error) {
 //parameters: SessionPoolHandle ProxyAuthenticationNameModifier
 //AgtLacPool SetProxyAuthenticationNameModifierList
 return "", nil
}

func(np *LacPool) GetProxyAuthenticationNameModifierList ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenticationNameModifierList
 return "", nil
}

func(np *LacPool) SetProxyAuthenticationNameModifier () error {
 //parameters: SessionPoolHandle ProxyAuthenticationNameModifier
 //AgtLacPool SetProxyAuthenticationNameModifier, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetProxyAuthenticationNameModifier ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenticationNameModifier
 return "", nil
}

func(np *LacPool) SetProxyAuthenPassword () error {
 //parameters: SessionPoolHandle ProxyAuthenPassword
 //AgtLacPool SetProxyAuthenPassword, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetProxyAuthenPassword ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenPassword
 return "", nil
}

func(np *LacPool) GetProxyAuthenPasswordValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenPasswordValue
 return "", nil
}

func(np *LacPool) SetProxyAuthenticationPasswordModifierIncrementor () error {
 //parameters: SessionPoolHandle ProxyAuthenticationPasswordModifier Count Increment Repeat
 //AgtLacPool SetProxyAuthenticationPasswordModifierIncrementor, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetProxyAuthenticationPasswordModifierIncrementor ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenticationPasswordModifierIncrementor
 return "", nil
}

func(np *LacPool) SetProxyAuthenticationPasswordModifierList ()(string, error) {
 //parameters: SessionPoolHandle ProxyAuthenticationPasswordModifier
 //AgtLacPool SetProxyAuthenticationPasswordModifierList
 return "", nil
}

func(np *LacPool) GetProxyAuthenticationPasswordModifierList ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenticationPasswordModifierList
 return "", nil
}

func(np *LacPool) SetProxyAuthenticationPasswordModifier () error {
 //parameters: SessionPoolHandle ProxyAuthenticationPasswordModifier
 //AgtLacPool SetProxyAuthenticationPasswordModifier, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetProxyAuthenticationPasswordModifier ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetProxyAuthenticationPasswordModifier
 return "", nil
}

func(np *LacPool) EnableAutoStart () error {
 //parameters: SessionPoolHandle
 //AgtLacPool EnableAutoStart, m.Object, m.Name);
 return nil
}

func(np *LacPool) DisableAutoStart () error {
 //parameters: SessionPoolHandle
 //AgtLacPool DisableAutoStart, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsAutoStartEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLacPool IsAutoStartEnabled, m.Object, m.Name);
 return nil
}

func(np *LacPool) SetMode () error {
 //parameters: SessionPoolHandle Mode
 //AgtLacPool SetMode, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetMode ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetMode
 return "", nil
}

func(np *LacPool) SetNumTunnels () error {
 //parameters: SessionPoolHandle NumTunnels
 //AgtLacPool SetNumTunnels, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetNumTunnels ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetNumTunnels
 return "", nil
}

func(np *LacPool) EnableSutGatewayIpv4Address () error {
 //parameters: SessionPoolHandle
 //AgtLacPool EnableSutGatewayIpv4Address, m.Object, m.Name);
 return nil
}

func(np *LacPool) DisableSutGatewayIpv4Address () error {
 //parameters: SessionPoolHandle
 //AgtLacPool DisableSutGatewayIpv4Address, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsSutGatewayIpv4AddressEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLacPool IsSutGatewayIpv4AddressEnabled, m.Object, m.Name);
 return nil
}

func(np *LacPool) SetSutGatewayIpv4AddressIncrementingRange () error {
 //parameters: SessionPoolHandle Ipv4Address PrefixLength Increment Repeat
 //AgtLacPool SetSutGatewayIpv4AddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetSutGatewayIpv4AddressIncrementingRange ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetSutGatewayIpv4AddressIncrementingRange
 return "", nil
}

func(np *LacPool) SetSutGatewayIpv4AddressList ()(string, error) {
 //parameters: SessionPoolHandle Ipv4Address
 //AgtLacPool SetSutGatewayIpv4AddressList
 return "", nil
}

func(np *LacPool) GetSutGatewayIpv4AddressList ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetSutGatewayIpv4AddressList
 return "", nil
}

func(np *LacPool) SetSutGatewayIpv4Address () error {
 //parameters: SessionPoolHandle Ipv4Address
 //AgtLacPool SetSutGatewayIpv4Address, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetSutGatewayIpv4Address ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetSutGatewayIpv4Address
 return "", nil
}

func(np *LacPool) GetPeersPerTunnel ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetPeersPerTunnel
 return "", nil
}

func(np *LacPool) SetTunnelLifeTime () error {
 //parameters: SessionPoolHandle TunnelLifeTime
 //AgtLacPool SetTunnelLifeTime, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetTunnelLifeTime ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetTunnelLifeTime
 return "", nil
}

func(np *LacPool) SetInitialAckTimeOut () error {
 //parameters: SessionPoolHandle InitialAckTimeOut
 //AgtLacPool SetInitialAckTimeOut, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetInitialAckTimeOut ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetInitialAckTimeOut
 return "", nil
}

func(np *LacPool) SetMaxTxRetries () error {
 //parameters: SessionPoolHandle MaxTxRetries
 //AgtLacPool SetMaxTxRetries, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetMaxTxRetries ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetMaxTxRetries
 return "", nil
}

func(np *LacPool) SetHostName () error {
 //parameters: SessionPoolHandle HostName
 //AgtLacPool SetHostName, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetHostName ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetHostName
 return "", nil
}

func(np *LacPool) EnableAuthenticateSut () error {
 //parameters: SessionPoolHandle
 //AgtLacPool EnableAuthenticateSut, m.Object, m.Name);
 return nil
}

func(np *LacPool) DisableAuthenticateSut () error {
 //parameters: SessionPoolHandle
 //AgtLacPool DisableAuthenticateSut, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsAuthenticateSutEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLacPool IsAuthenticateSutEnabled, m.Object, m.Name);
 return nil
}

func(np *LacPool) SetSessionOpeningRate () error {
 //parameters: SessionPoolHandle SessionOpeningRate
 //AgtLacPool SetSessionOpeningRate, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetSessionOpeningRate ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetSessionOpeningRate
 return "", nil
}

func(np *LacPool) SetSharedSecret () error {
 //parameters: SessionPoolHandle SharedSecret
 //AgtLacPool SetSharedSecret, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetSharedSecret ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetSharedSecret
 return "", nil
}

func(np *LacPool) SetHelloInterval () error {
 //parameters: SessionPoolHandle HelloInterval
 //AgtLacPool SetHelloInterval, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetHelloInterval ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetHelloInterval
 return "", nil
}

func(np *LacPool) EnableExcludeHdlcHeader () error {
 //parameters: SessionPoolHandle
 //AgtLacPool EnableExcludeHdlcHeader, m.Object, m.Name);
 return nil
}

func(np *LacPool) DisableExcludeHdlcHeader () error {
 //parameters: SessionPoolHandle
 //AgtLacPool DisableExcludeHdlcHeader, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsExcludeHdlcHeaderEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLacPool IsExcludeHdlcHeaderEnabled, m.Object, m.Name);
 return nil
}

func(np *LacPool) EnableL2tpHeaderOffset () error {
 //parameters: SessionPoolHandle
 //AgtLacPool EnableL2tpHeaderOffset, m.Object, m.Name);
 return nil
}

func(np *LacPool) DisableL2tpHeaderOffset () error {
 //parameters: SessionPoolHandle
 //AgtLacPool DisableL2tpHeaderOffset, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsL2tpHeaderOffsetEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLacPool IsL2tpHeaderOffsetEnabled, m.Object, m.Name);
 return nil
}

func(np *LacPool) EnableReceiveWindowSize () error {
 //parameters: SessionPoolHandle
 //AgtLacPool EnableReceiveWindowSize, m.Object, m.Name);
 return nil
}

func(np *LacPool) DisableReceiveWindowSize () error {
 //parameters: SessionPoolHandle
 //AgtLacPool DisableReceiveWindowSize, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsReceiveWindowSizeEnabled () error {
 //parameters: SessionPoolHandle
 //AgtLacPool IsReceiveWindowSizeEnabled, m.Object, m.Name);
 return nil
}

func(np *LacPool) SetReceiveWindowSizeValue () error {
 //parameters: SessionPoolHandle ReceiveWindowSizeValue
 //AgtLacPool SetReceiveWindowSizeValue, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetReceiveWindowSizeValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtLacPool GetReceiveWindowSizeValue
 return "", nil
}

func(np *LacPool) open () error {
 //parameters: SessionIdentifiers
 //AgtLacPool open, m.Object, m.Name);
 return nil
}

func(np *LacPool) close () error {
 //parameters: SessionIdentifiers
 //AgtLacPool close, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsIpv4PriorityNoCodePointFieldSelected () error {
 //parameters: DeviceHandle
 //AgtLacPool IsIpv4PriorityNoCodePointFieldSelected, m.Object, m.Name);
 return nil
}

func(np *LacPool) SelectIpv4PriorityNoCodePointField () error {
 //parameters: DeviceHandle
 //AgtLacPool SelectIpv4PriorityNoCodePointField, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetIpv4Priority ()(string, error) {
 //parameters: DeviceHandle
 //AgtLacPool GetIpv4Priority
 return "", nil
}

func(np *LacPool) SetIpv4Priority () error {
 //parameters: DeviceHandle Ipv4Priority
 //AgtLacPool SetIpv4Priority, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsIpv4PriorityTypeOfServiceFieldSelected () error {
 //parameters: DeviceHandle
 //AgtLacPool IsIpv4PriorityTypeOfServiceFieldSelected, m.Object, m.Name);
 return nil
}

func(np *LacPool) SelectIpv4PriorityTypeOfServiceField () error {
 //parameters: DeviceHandle
 //AgtLacPool SelectIpv4PriorityTypeOfServiceField, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetIpv4PriorityTypeOfServiceValue ()(string, error) {
 //parameters: DeviceHandle TosCodePointField
 //AgtLacPool GetIpv4PriorityTypeOfServiceValue
 return "", nil
}

func(np *LacPool) SetIpv4PriorityTypeOfServiceValue () error {
 //parameters: DeviceHandle TosCodePointField Value
 //AgtLacPool SetIpv4PriorityTypeOfServiceValue, m.Object, m.Name);
 return nil
}

func(np *LacPool) IsIpv4PriorityDiffServFieldSelected () error {
 //parameters: DeviceHandle DiffServPerHopBehaviorGroup
 //AgtLacPool IsIpv4PriorityDiffServFieldSelected, m.Object, m.Name);
 return nil
}

func(np *LacPool) SelectIpv4PriorityDiffServField () error {
 //parameters: DeviceHandle DiffServPerHopBehaviorGroup
 //AgtLacPool SelectIpv4PriorityDiffServField, m.Object, m.Name);
 return nil
}

func(np *LacPool) GetIpv4PriorityDiffServValue ()(string, error) {
 //parameters: DeviceHandle DiffServCodePointField
 //AgtLacPool GetIpv4PriorityDiffServValue
 return "", nil
}

func(np *LacPool) SetIpv4PriorityDiffServValue () error {
 //parameters: DeviceHandle DiffServCodePointConfigurableField Value
 //AgtLacPool SetIpv4PriorityDiffServValue, m.Object, m.Name);
 return nil
}

