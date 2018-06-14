package n2xsdk

type PppoAClientPool struct {
 Handler string
}

func(np *PppoAClientPool) GetSessionId ()(string, error) {
 //parameters: SessionIdentifiers
 //AgtPppoAClientPool GetSessionId
 return "", nil
}

func(np *PppoAClientPool) SetStartingVpiVci () error {
 //parameters: DeviceHandle StartingVpi StartingVci
 //AgtPppoAClientPool SetStartingVpiVci, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetStartingVpiVci ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetStartingVpiVci
 return "", nil
}

func(np *PppoAClientPool) SetVpiVciIncrement () error {
 //parameters: DeviceHandle VpiIncrement VciIncrement
 //AgtPppoAClientPool SetVpiVciIncrement, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetVpiVciIncrement ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetVpiVciIncrement
 return "", nil
}

func(np *PppoAClientPool) SetAtmEncapsulation () error {
 //parameters: DeviceHandle Encapsulation
 //AgtPppoAClientPool SetAtmEncapsulation, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetAtmEncapsulation ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetAtmEncapsulation
 return "", nil
}

func(np *PppoAClientPool) GetIpv6cpEnableFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetIpv6cpEnableFlag
 return "", nil
}

func(np *PppoAClientPool) SetIpv6cpEnableFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetIpv6cpEnableFlag, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetIpv6cpAdvertisedIidMode ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetIpv6cpAdvertisedIidMode
 return "", nil
}

func(np *PppoAClientPool) SetIpv6cpAdvertisedIidMode () error {
 //parameters: DeviceHandle AdvertisedIidMode
 //AgtPppoAClientPool SetIpv6cpAdvertisedIidMode, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetIpv6cpIidIncrementingRange ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetIpv6cpIidIncrementingRange
 return "", nil
}

func(np *PppoAClientPool) SetIpv6cpIidIncrementingRange () error {
 //parameters: DeviceHandle IncrementorValue IncrementorValueIncrement
 //AgtPppoAClientPool SetIpv6cpIidIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetIpv6cpIidList ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetIpv6cpIidList
 return "", nil
}

func(np *PppoAClientPool) SetIpv6cpIidList ()(string, error) {
 //parameters: DeviceHandle IteratedValueList
 //AgtPppoAClientPool SetIpv6cpIidList
 return "", nil
}

func(np *PppoAClientPool) GetNegotiatedIpv6cpIid ()(string, error) {
 //parameters: SessionIdentifiers
 //AgtPppoAClientPool GetNegotiatedIpv6cpIid
 return "", nil
}

func(np *PppoAClientPool) OpenIpcp () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool OpenIpcp, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) CloseIpcp () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool CloseIpcp, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) OpenIpv6cp () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool OpenIpv6cp, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) CloseIpv6cp () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool CloseIpv6cp, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) IsMlPppLcpEPDOptionEnabled () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool IsMlPppLcpEPDOptionEnabled, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) DisableMlPppLcpEPDOption () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool DisableMlPppLcpEPDOption, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) EnableMlPppLcpEPDOption () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool EnableMlPppLcpEPDOption, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetMlPppLcpEPDClass ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetMlPppLcpEPDClass
 return "", nil
}

func(np *PppoAClientPool) SetMlPppLcpEPDClass () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetMlPppLcpEPDClass, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetEPDIpAddressIncrementor ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetEPDIpAddressIncrementor
 return "", nil
}

func(np *PppoAClientPool) SetEPDIpAddressIncrementor () error {
 //parameters: DeviceHandle IncrementorValue IncrementorValueIncrement IncrementorValueRepeat
 //AgtPppoAClientPool SetEPDIpAddressIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetEnableFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetEnableFlag
 return "", nil
}

func(np *PppoAClientPool) Enable () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool Enable, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) Disable () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool Disable, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) Cancel () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool Cancel, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) Reset () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool Reset, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) Open () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool Open, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) Close () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool Close, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) CancelAttempts () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool CancelAttempts, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetNumberOfEstablishedSessions ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetNumberOfEstablishedSessions
 return "", nil
}

func(np *PppoAClientPool) SetConnectionRetryFlags () error {
 //parameters: DeviceHandle ConnectionRetryFlags
 //AgtPppoAClientPool SetConnectionRetryFlags, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetConnectionRetryFlags ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetConnectionRetryFlags
 return "", nil
}

func(np *PppoAClientPool) SetReestablishmentFlags () error {
 //parameters: DeviceHandle ReestablishmentFlags
 //AgtPppoAClientPool SetReestablishmentFlags, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetReestablishmentFlags ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetReestablishmentFlags
 return "", nil
}

func(np *PppoAClientPool) SetUnlimitedConnectionAttemptsFlag () error {
 //parameters: DeviceHandle EnableFlag
 //AgtPppoAClientPool SetUnlimitedConnectionAttemptsFlag, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetUnlimitedConnectionAttemptsFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetUnlimitedConnectionAttemptsFlag
 return "", nil
}

func(np *PppoAClientPool) EnableUnlimitedReestablishment () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool EnableUnlimitedReestablishment, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) DisableUnlimitedReestablishment () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool DisableUnlimitedReestablishment, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) IsUnlimitedReestablishmentEnabled () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool IsUnlimitedReestablishmentEnabled, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetStateCount ()(string, error) {
 //parameters: DeviceHandle SessionStateName
 //AgtPppoAClientPool GetStateCount
 return "", nil
}

func(np *PppoAClientPool) GetSessionState ()(string, error) {
 //parameters: SessionIdentifiers
 //AgtPppoAClientPool GetSessionState
 return "", nil
}

func(np *PppoAClientPool) GetSessionSubprotocolState ()(string, error) {
 //parameters: SessionIdentifiers SubProtocolName
 //AgtPppoAClientPool GetSessionSubprotocolState
 return "", nil
}

func(np *PppoAClientPool) GetNumberOfSessions ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetNumberOfSessions
 return "", nil
}

func(np *PppoAClientPool) SetNumberOfSessions () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetNumberOfSessions, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetAutoStartFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetAutoStartFlag
 return "", nil
}

func(np *PppoAClientPool) SetAutoStartFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetAutoStartFlag, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetSessionLifetime ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetSessionLifetime
 return "", nil
}

func(np *PppoAClientPool) SetSessionLifetime () error {
 //parameters: DeviceHandle Value Value
 //AgtPppoAClientPool SetSessionLifetime, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetOpeningWindow ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetOpeningWindow
 return "", nil
}

func(np *PppoAClientPool) GetClosingWindow ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetClosingWindow
 return "", nil
}

func(np *PppoAClientPool) SetOpeningWindow () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetOpeningWindow, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) SetClosingWindow () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetClosingWindow, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) IsSustainedProlongedFlappingEnabled () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool IsSustainedProlongedFlappingEnabled, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) DisableSustainedProlongedFlapping () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool DisableSustainedProlongedFlapping, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) EnableSustainedProlongedFlapping () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool EnableSustainedProlongedFlapping, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetMaxConnectionAttempts ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetMaxConnectionAttempts
 return "", nil
}

func(np *PppoAClientPool) SetMaxConnectionAttempts () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetMaxConnectionAttempts, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetConnectionAttempts ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetConnectionAttempts
 return "", nil
}

func(np *PppoAClientPool) GetMaxReestablishments ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetMaxReestablishments
 return "", nil
}

func(np *PppoAClientPool) SetMaxReestablishments () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetMaxReestablishments, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) RetryConnections () error {
 //parameters: DeviceOrSession
 //AgtPppoAClientPool RetryConnections, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetInitiationRateLimitFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetInitiationRateLimitFlag
 return "", nil
}

func(np *PppoAClientPool) SetInitiationRateLimitFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetInitiationRateLimitFlag, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetSessionInitiationRate ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetSessionInitiationRate
 return "", nil
}

func(np *PppoAClientPool) SetSessionInitiationRate () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetSessionInitiationRate, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetLimitOpeningRateMode ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetLimitOpeningRateMode
 return "", nil
}

func(np *PppoAClientPool) SetLimitOpeningRateMode () error {
 //parameters: DeviceHandle SetupInterval
 //AgtPppoAClientPool SetLimitOpeningRateMode, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetSessionInitiationRateInMilliseconds ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetSessionInitiationRateInMilliseconds
 return "", nil
}

func(np *PppoAClientPool) SetSessionInitiationRateInMilliseconds () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetSessionInitiationRateInMilliseconds, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetClosureRateLimitFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetClosureRateLimitFlag
 return "", nil
}

func(np *PppoAClientPool) SetClosureRateLimitFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetClosureRateLimitFlag, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetSessionClosureRate ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetSessionClosureRate
 return "", nil
}

func(np *PppoAClientPool) SetSessionClosureRate () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetSessionClosureRate, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetLcpOption ()(string, error) {
 //parameters: DeviceHandle PppOption
 //AgtPppoAClientPool GetLcpOption
 return "", nil
}

func(np *PppoAClientPool) SetLcpOption () error {
 //parameters: DeviceHandle PppOption Value
 //AgtPppoAClientPool SetLcpOption, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetPppOption ()(string, error) {
 //parameters: DeviceHandle PppOption
 //AgtPppoAClientPool GetPppOption
 return "", nil
}

func(np *PppoAClientPool) SetPppOption () error {
 //parameters: DeviceHandle PppOption Value
 //AgtPppoAClientPool SetPppOption, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) IsMlPppLcpMrruOptionEnabled () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool IsMlPppLcpMrruOptionEnabled, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) DisableMlPppLcpMrruOption () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool DisableMlPppLcpMrruOption, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) EnableMlPppLcpMrruOption () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool EnableMlPppLcpMrruOption, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetMlPppLcpMrruSize ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetMlPppLcpMrruSize
 return "", nil
}

func(np *PppoAClientPool) SetMlPppLcpMrruSize () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetMlPppLcpMrruSize, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetIpcpEnableFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetIpcpEnableFlag
 return "", nil
}

func(np *PppoAClientPool) SetIpcpEnableFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetIpcpEnableFlag, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetIpcpAdvertisedAddressMode ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetIpcpAdvertisedAddressMode
 return "", nil
}

func(np *PppoAClientPool) SetIpcpAdvertisedAddressMode () error {
 //parameters: DeviceHandle AdvertisedIpValue
 //AgtPppoAClientPool SetIpcpAdvertisedAddressMode, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetIpcpAdvertisedAddressIncrementor ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetIpcpAdvertisedAddressIncrementor
 return "", nil
}

func(np *PppoAClientPool) SetIpcpAdvertisedAddressIncrementor () error {
 //parameters: DeviceHandle IncrementorValue IncrementorValueIncrement IncrementorValueRepeat
 //AgtPppoAClientPool SetIpcpAdvertisedAddressIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetNegotiateNetmaskFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetNegotiateNetmaskFlag
 return "", nil
}

func(np *PppoAClientPool) SetNegotiateNetmaskFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetNegotiateNetmaskFlag, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetNetmaskRequestType ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetNetmaskRequestType
 return "", nil
}

func(np *PppoAClientPool) SetNetmaskRequestType () error {
 //parameters: DeviceHandle NetmaskRequestType
 //AgtPppoAClientPool SetNetmaskRequestType, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetRequestedNetmaskLength ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetRequestedNetmaskLength
 return "", nil
}

func(np *PppoAClientPool) SetRequestedNetmaskLength () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetRequestedNetmaskLength, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetRequestedNetmaskLengthModifier ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetRequestedNetmaskLengthModifier
 return "", nil
}

func(np *PppoAClientPool) SetRequestedNetmaskLengthModifier () error {
 //parameters: DeviceHandle Value IncrementorValueRepeat IncrementorValueIncrement
 //AgtPppoAClientPool SetRequestedNetmaskLengthModifier, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetNegotiatedNetmask ()(string, error) {
 //parameters: SessionIdentifiers
 //AgtPppoAClientPool GetNegotiatedNetmask
 return "", nil
}

func(np *PppoAClientPool) GetRequestNameServerAddressFlag ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoAClientPool GetRequestNameServerAddressFlag
 return "", nil
}

func(np *PppoAClientPool) SetRequestNameServerAddressFlag () error {
 //parameters: DeviceHandle NameAddressServerType Value
 //AgtPppoAClientPool SetRequestNameServerAddressFlag, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetNameServerAddressRequestType ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoAClientPool GetNameServerAddressRequestType
 return "", nil
}

func(np *PppoAClientPool) SetNameServerAddressRequestType () error {
 //parameters: DeviceHandle NameAddressServerType NameServerAddressRequestType
 //AgtPppoAClientPool SetNameServerAddressRequestType, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetRequestedNameServerAddress ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoAClientPool GetRequestedNameServerAddress
 return "", nil
}

func(np *PppoAClientPool) SetRequestedNameServerAddress () error {
 //parameters: DeviceHandle NameAddressServerType IPv4Address
 //AgtPppoAClientPool SetRequestedNameServerAddress, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetRequestedNameServerAddressIncrementingValue ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoAClientPool GetRequestedNameServerAddressIncrementingValue
 return "", nil
}

func(np *PppoAClientPool) SetRequestedNameServerAddressIncrementingValue () error {
 //parameters: DeviceHandle NameAddressServerType IPv4Address IncrementorValueRepeat IncrementorValueIncrement
 //AgtPppoAClientPool SetRequestedNameServerAddressIncrementingValue, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetRequestedNameServerAddressList ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoAClientPool GetRequestedNameServerAddressList
 return "", nil
}

func(np *PppoAClientPool) SetRequestedNameServerAddressList ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType IPv4Address
 //AgtPppoAClientPool SetRequestedNameServerAddressList
 return "", nil
}

func(np *PppoAClientPool) GetNameServerAddress ()(string, error) {
 //parameters: SessionIdentifiers NameAddressServerType
 //AgtPppoAClientPool GetNameServerAddress
 return "", nil
}

func(np *PppoAClientPool) GetAuthenticationProtocol ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetAuthenticationProtocol
 return "", nil
}

func(np *PppoAClientPool) SetAuthenticationProtocol () error {
 //parameters: DeviceHandle AuthProtocol
 //AgtPppoAClientPool SetAuthenticationProtocol, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) SetAuthenticationRestartTimer () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetAuthenticationRestartTimer, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetAuthenticationRestartTimer ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetAuthenticationRestartTimer
 return "", nil
}

func(np *PppoAClientPool) SetAuthenticationMaximumAttempts () error {
 //parameters: DeviceHandle Value
 //AgtPppoAClientPool SetAuthenticationMaximumAttempts, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetAuthenticationMaximumAttempts ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetAuthenticationMaximumAttempts
 return "", nil
}

func(np *PppoAClientPool) GetTesterCredentials ()(string, error) {
 //parameters: DeviceHandle CredentialParameterName
 //AgtPppoAClientPool GetTesterCredentials
 return "", nil
}

func(np *PppoAClientPool) SetTesterCredentials () error {
 //parameters: DeviceHandle CredentialParameterName Value
 //AgtPppoAClientPool SetTesterCredentials, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetTesterCredentialsIncrementor ()(string, error) {
 //parameters: DeviceHandle CredentialParameterName
 //AgtPppoAClientPool GetTesterCredentialsIncrementor
 return "", nil
}

func(np *PppoAClientPool) SetTesterCredentialsIncrementor () error {
 //parameters: DeviceHandle CredentialParameterName IteratedValue IteratedValueIncrement IteratedValueRepeat
 //AgtPppoAClientPool SetTesterCredentialsIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) AddTesterCredentialsNumberedIncrementor () error {
 //parameters: DeviceHandle CredentialParameterName
 //AgtPppoAClientPool AddTesterCredentialsNumberedIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) RemoveTesterCredentialsNumberedIncrementor () error {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber
 //AgtPppoAClientPool RemoveTesterCredentialsNumberedIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetTesterCredentialsNumberedIncrementor ()(string, error) {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber
 //AgtPppoAClientPool GetTesterCredentialsNumberedIncrementor
 return "", nil
}

func(np *PppoAClientPool) SetTesterCredentialsNumberedIncrementor () error {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber IteratedValue IteratedValueCount IteratedValueIncrement IteratedValueRepeat
 //AgtPppoAClientPool SetTesterCredentialsNumberedIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetTesterCredentialsIncrementorCountOverride ()(string, error) {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber
 //AgtPppoAClientPool GetTesterCredentialsIncrementorCountOverride
 return "", nil
}

func(np *PppoAClientPool) SetTesterCredentialsIncrementorCountOverride () error {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber IsIncrementCountOverridden
 //AgtPppoAClientPool SetTesterCredentialsIncrementorCountOverride, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) ListSelectedPools ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool ListSelectedPools
 return "", nil
}

func(np *PppoAClientPool) SelectPools () error {
 //parameters: SessionIdentifiers
 //AgtPppoAClientPool SelectPools, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) DeselectPools () error {
 //parameters: SessionIdentifiers
 //AgtPppoAClientPool DeselectPools, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetAccumulatedValues ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoAClientPool GetAccumulatedValues
 return "", nil
}

func(np *PppoAClientPool) GetDeviceAccumulatedValues ()(string, error) {
 //parameters: DeviceHandle DeviceStatisticsList
 //AgtPppoAClientPool GetDeviceAccumulatedValues
 return "", nil
}

func(np *PppoAClientPool) GetSessionAccumulatedSpecifiedValues ()(string, error) {
 //parameters: SessionIdentifiers SessionStatisticsList
 //AgtPppoAClientPool GetSessionAccumulatedSpecifiedValues
 return "", nil
}

func(np *PppoAClientPool) ClearStatistics () error {
 //parameters: DeviceHandle
 //AgtPppoAClientPool ClearStatistics, m.Object, m.Name);
 return nil
}

func(np *PppoAClientPool) GetLastFailureReason ()(string, error) {
 //parameters: SessionIdentifiers
 //AgtPppoAClientPool GetLastFailureReason
 return "", nil
}

