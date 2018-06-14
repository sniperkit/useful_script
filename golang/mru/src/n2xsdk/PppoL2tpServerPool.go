package n2xsdk

type PppoL2tpServerPool struct {
 Handler string
}

func(np *PppoL2tpServerPool) IsLcpForceNegotiationEnabled () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool IsLcpForceNegotiationEnabled, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) DisableLcpForceLcpNegotiation () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool DisableLcpForceLcpNegotiation, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) EnableLcpForceLcpNegotiation () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool EnableLcpForceLcpNegotiation, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetIpv6cpEnableFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetIpv6cpEnableFlag
 return "", nil
}

func(np *PppoL2tpServerPool) SetIpv6cpEnableFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetIpv6cpEnableFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetServerAddressPoolIncrementingRange ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetServerAddressPoolIncrementingRange
 return "", nil
}

func(np *PppoL2tpServerPool) SetServerAddressPoolIncrementingRange () error {
 //parameters: DeviceHandle IncrementorValue IncrementorValueIncrement PrefixLength Repeat
 //AgtPppoL2tpServerPool SetServerAddressPoolIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetSourceServerAddressPoolIncrementingRange ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetSourceServerAddressPoolIncrementingRange
 return "", nil
}

func(np *PppoL2tpServerPool) SetSourceServerAddressPoolIncrementingRange () error {
 //parameters: DeviceHandle IncrementorValue IncrementorValueIncrement PrefixLength Repeat
 //AgtPppoL2tpServerPool SetSourceServerAddressPoolIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) SetOfferedNetmaskLengthList ()(string, error) {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetOfferedNetmaskLengthList
 return "", nil
}

func(np *PppoL2tpServerPool) GetOfferedNetmaskLengthList ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetOfferedNetmaskLengthList
 return "", nil
}

func(np *PppoL2tpServerPool) OpenIpcp () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool OpenIpcp, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) CloseIpcp () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool CloseIpcp, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) OpenIpv6cp () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool OpenIpv6cp, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) CloseIpv6cp () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool CloseIpv6cp, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) IsMlPppLcpEPDOptionEnabled () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool IsMlPppLcpEPDOptionEnabled, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) DisableMlPppLcpEPDOption () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool DisableMlPppLcpEPDOption, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) EnableMlPppLcpEPDOption () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool EnableMlPppLcpEPDOption, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetMlPppLcpEPDClass ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetMlPppLcpEPDClass
 return "", nil
}

func(np *PppoL2tpServerPool) SetMlPppLcpEPDClass () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetMlPppLcpEPDClass, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetEPDIpAddressIncrementor ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetEPDIpAddressIncrementor
 return "", nil
}

func(np *PppoL2tpServerPool) SetEPDIpAddressIncrementor () error {
 //parameters: DeviceHandle IncrementorValue IncrementorValueIncrement IncrementorValueRepeat
 //AgtPppoL2tpServerPool SetEPDIpAddressIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetEnableFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetEnableFlag
 return "", nil
}

func(np *PppoL2tpServerPool) Enable () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool Enable, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) Disable () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool Disable, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) Cancel () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool Cancel, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) Reset () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool Reset, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) Open () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool Open, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) Close () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool Close, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) CancelAttempts () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool CancelAttempts, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetNumberOfEstablishedSessions ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetNumberOfEstablishedSessions
 return "", nil
}

func(np *PppoL2tpServerPool) SetConnectionRetryFlags () error {
 //parameters: DeviceHandle ConnectionRetryFlags
 //AgtPppoL2tpServerPool SetConnectionRetryFlags, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetConnectionRetryFlags ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetConnectionRetryFlags
 return "", nil
}

func(np *PppoL2tpServerPool) SetReestablishmentFlags () error {
 //parameters: DeviceHandle ReestablishmentFlags
 //AgtPppoL2tpServerPool SetReestablishmentFlags, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetReestablishmentFlags ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetReestablishmentFlags
 return "", nil
}

func(np *PppoL2tpServerPool) SetUnlimitedConnectionAttemptsFlag () error {
 //parameters: DeviceHandle EnableFlag
 //AgtPppoL2tpServerPool SetUnlimitedConnectionAttemptsFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetUnlimitedConnectionAttemptsFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetUnlimitedConnectionAttemptsFlag
 return "", nil
}

func(np *PppoL2tpServerPool) EnableUnlimitedReestablishment () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool EnableUnlimitedReestablishment, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) DisableUnlimitedReestablishment () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool DisableUnlimitedReestablishment, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) IsUnlimitedReestablishmentEnabled () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool IsUnlimitedReestablishmentEnabled, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetStateCount ()(string, error) {
 //parameters: DeviceHandle SessionStateName
 //AgtPppoL2tpServerPool GetStateCount
 return "", nil
}

func(np *PppoL2tpServerPool) GetSessionState ()(string, error) {
 //parameters: SessionIdentifiers
 //AgtPppoL2tpServerPool GetSessionState
 return "", nil
}

func(np *PppoL2tpServerPool) GetSessionSubprotocolState ()(string, error) {
 //parameters: SessionIdentifiers SubprotocolName
 //AgtPppoL2tpServerPool GetSessionSubprotocolState
 return "", nil
}

func(np *PppoL2tpServerPool) GetNumberOfSessions ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetNumberOfSessions
 return "", nil
}

func(np *PppoL2tpServerPool) SetNumberOfSessions () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetNumberOfSessions, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetAutoStartFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetAutoStartFlag
 return "", nil
}

func(np *PppoL2tpServerPool) SetAutoStartFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetAutoStartFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetSessionLifetime ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetSessionLifetime
 return "", nil
}

func(np *PppoL2tpServerPool) SetSessionLifetime () error {
 //parameters: DeviceHandle Value Value
 //AgtPppoL2tpServerPool SetSessionLifetime, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetOpeningWindow ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetOpeningWindow
 return "", nil
}

func(np *PppoL2tpServerPool) SetOpeningWindow () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetOpeningWindow, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetMaxConnectionAttempts ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetMaxConnectionAttempts
 return "", nil
}

func(np *PppoL2tpServerPool) SetMaxConnectionAttempts () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetMaxConnectionAttempts, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetConnectionAttempts ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetConnectionAttempts
 return "", nil
}

func(np *PppoL2tpServerPool) GetMaxReestablishments ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetMaxReestablishments
 return "", nil
}

func(np *PppoL2tpServerPool) SetMaxReestablishments () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetMaxReestablishments, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) RetryConnections () error {
 //parameters: DeviceOrSession
 //AgtPppoL2tpServerPool RetryConnections, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetInitiationRateLimitFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetInitiationRateLimitFlag
 return "", nil
}

func(np *PppoL2tpServerPool) SetInitiationRateLimitFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetInitiationRateLimitFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetSessionInitiationRate ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetSessionInitiationRate
 return "", nil
}

func(np *PppoL2tpServerPool) SetSessionInitiationRate () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetSessionInitiationRate, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetLimitOpeningRateMode ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetLimitOpeningRateMode
 return "", nil
}

func(np *PppoL2tpServerPool) SetLimitOpeningRateMode () error {
 //parameters: DeviceHandle SetupInterval
 //AgtPppoL2tpServerPool SetLimitOpeningRateMode, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetSessionInitiationRateInMilliseconds ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetSessionInitiationRateInMilliseconds
 return "", nil
}

func(np *PppoL2tpServerPool) SetSessionInitiationRateInMilliseconds () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetSessionInitiationRateInMilliseconds, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetLcpOption ()(string, error) {
 //parameters: DeviceHandle PppOption
 //AgtPppoL2tpServerPool GetLcpOption
 return "", nil
}

func(np *PppoL2tpServerPool) SetLcpOption () error {
 //parameters: DeviceHandle PppOption Value
 //AgtPppoL2tpServerPool SetLcpOption, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetPppOption ()(string, error) {
 //parameters: DeviceHandle PppOption
 //AgtPppoL2tpServerPool GetPppOption
 return "", nil
}

func(np *PppoL2tpServerPool) SetPppOption () error {
 //parameters: DeviceHandle PppOption Value
 //AgtPppoL2tpServerPool SetPppOption, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) IsMlPppLcpMrruOptionEnabled () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool IsMlPppLcpMrruOptionEnabled, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) DisableMlPppLcpMrruOption () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool DisableMlPppLcpMrruOption, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) EnableMlPppLcpMrruOption () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool EnableMlPppLcpMrruOption, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetMlPppLcpMrruSize ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetMlPppLcpMrruSize
 return "", nil
}

func(np *PppoL2tpServerPool) SetMlPppLcpMrruSize () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetMlPppLcpMrruSize, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetIpcpEnableFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetIpcpEnableFlag
 return "", nil
}

func(np *PppoL2tpServerPool) SetIpcpEnableFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetIpcpEnableFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetOfferNetmaskFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetOfferNetmaskFlag
 return "", nil
}

func(np *PppoL2tpServerPool) SetOfferNetmaskFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetOfferNetmaskFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetMandatoryNetmaskNegotiationFlag ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetMandatoryNetmaskNegotiationFlag
 return "", nil
}

func(np *PppoL2tpServerPool) SetMandatoryNetmaskNegotiationFlag () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetMandatoryNetmaskNegotiationFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetOfferedNetmaskLength ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetOfferedNetmaskLength
 return "", nil
}

func(np *PppoL2tpServerPool) SetOfferedNetmaskLength () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetOfferedNetmaskLength, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetOfferedNetmaskLengthModifier ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetOfferedNetmaskLengthModifier
 return "", nil
}

func(np *PppoL2tpServerPool) SetOfferedNetmaskLengthModifier () error {
 //parameters: DeviceHandle Value Repeat Increment
 //AgtPppoL2tpServerPool SetOfferedNetmaskLengthModifier, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetOfferNameServerAddressFlag ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoL2tpServerPool GetOfferNameServerAddressFlag
 return "", nil
}

func(np *PppoL2tpServerPool) SetOfferNameServerAddressFlag () error {
 //parameters: DeviceHandle NameAddressServerType Value
 //AgtPppoL2tpServerPool SetOfferNameServerAddressFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetMandatoryNameServerAddressNegotiationFlag ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoL2tpServerPool GetMandatoryNameServerAddressNegotiationFlag
 return "", nil
}

func(np *PppoL2tpServerPool) SetMandatoryNameServerAddressNegotiationFlag () error {
 //parameters: DeviceHandle NameAddressServerType Value
 //AgtPppoL2tpServerPool SetMandatoryNameServerAddressNegotiationFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetOfferedNameServerAddress ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoL2tpServerPool GetOfferedNameServerAddress
 return "", nil
}

func(np *PppoL2tpServerPool) SetOfferedNameServerAddress () error {
 //parameters: DeviceHandle NameAddressServerType Value
 //AgtPppoL2tpServerPool SetOfferedNameServerAddress, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetOfferedNameServerAddressIncrementingValue ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoL2tpServerPool GetOfferedNameServerAddressIncrementingValue
 return "", nil
}

func(np *PppoL2tpServerPool) SetOfferedNameServerAddressIncrementingValue () error {
 //parameters: DeviceHandle NameAddressServerType Value Repeat Increment
 //AgtPppoL2tpServerPool SetOfferedNameServerAddressIncrementingValue, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetOfferedNameServerAddressList ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType
 //AgtPppoL2tpServerPool GetOfferedNameServerAddressList
 return "", nil
}

func(np *PppoL2tpServerPool) SetOfferedNameServerAddressList ()(string, error) {
 //parameters: DeviceHandle NameAddressServerType Value
 //AgtPppoL2tpServerPool SetOfferedNameServerAddressList
 return "", nil
}

func(np *PppoL2tpServerPool) GetAuthenticationProtocol ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetAuthenticationProtocol
 return "", nil
}

func(np *PppoL2tpServerPool) SetAuthenticationProtocol () error {
 //parameters: DeviceHandle AuthProtocol
 //AgtPppoL2tpServerPool SetAuthenticationProtocol, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) SetAuthenticationRestartTimer () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetAuthenticationRestartTimer, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetAuthenticationRestartTimer ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetAuthenticationRestartTimer
 return "", nil
}

func(np *PppoL2tpServerPool) SetAuthenticationMaximumAttempts () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetAuthenticationMaximumAttempts, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetAuthenticationMaximumAttempts ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetAuthenticationMaximumAttempts
 return "", nil
}

func(np *PppoL2tpServerPool) GetTesterCredentials ()(string, error) {
 //parameters: DeviceHandle CredentialParameterName
 //AgtPppoL2tpServerPool GetTesterCredentials
 return "", nil
}

func(np *PppoL2tpServerPool) SetTesterCredentials () error {
 //parameters: DeviceHandle CredentialParameterName Value
 //AgtPppoL2tpServerPool SetTesterCredentials, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetTesterCredentialsIncrementor ()(string, error) {
 //parameters: DeviceHandle CredentialParameterName
 //AgtPppoL2tpServerPool GetTesterCredentialsIncrementor
 return "", nil
}

func(np *PppoL2tpServerPool) SetTesterCredentialsIncrementor () error {
 //parameters: DeviceHandle CredentialParameterName IteratedValue IteratedValueIncrement IteratedValueRepeat
 //AgtPppoL2tpServerPool SetTesterCredentialsIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) AddTesterCredentialsNumberedIncrementor () error {
 //parameters: DeviceHandle CredentialParameterName
 //AgtPppoL2tpServerPool AddTesterCredentialsNumberedIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) RemoveTesterCredentialsNumberedIncrementor () error {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber
 //AgtPppoL2tpServerPool RemoveTesterCredentialsNumberedIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetTesterCredentialsNumberedIncrementor ()(string, error) {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber
 //AgtPppoL2tpServerPool GetTesterCredentialsNumberedIncrementor
 return "", nil
}

func(np *PppoL2tpServerPool) SetTesterCredentialsNumberedIncrementor () error {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber IteratedValue IteratedValueCount IteratedValueIncrement IteratedValueRepeat
 //AgtPppoL2tpServerPool SetTesterCredentialsNumberedIncrementor, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetTesterCredentialsIncrementorCountOverride ()(string, error) {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber
 //AgtPppoL2tpServerPool GetTesterCredentialsIncrementorCountOverride
 return "", nil
}

func(np *PppoL2tpServerPool) SetTesterCredentialsIncrementorCountOverride () error {
 //parameters: DeviceHandle CredentialParameterName IncrementorNumber IsIncrementCountOverridden
 //AgtPppoL2tpServerPool SetTesterCredentialsIncrementorCountOverride, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) EnableAcceptAnySutCredentialFlag () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool EnableAcceptAnySutCredentialFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) DisableAcceptAnySutCredentialFlag () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool DisableAcceptAnySutCredentialFlag, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) IsAcceptAnySutCredentialFlagEnabled () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool IsAcceptAnySutCredentialFlagEnabled, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) SetAcceptedSutNameOrPeerIdValue () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetAcceptedSutNameOrPeerIdValue, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetAcceptedSutNameOrPeerIdValue ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetAcceptedSutNameOrPeerIdValue
 return "", nil
}

func(np *PppoL2tpServerPool) SetAcceptedSutPasswordOrSecretValue () error {
 //parameters: DeviceHandle Value
 //AgtPppoL2tpServerPool SetAcceptedSutPasswordOrSecretValue, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetAcceptedSutPasswordOrSecretValue ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetAcceptedSutPasswordOrSecretValue
 return "", nil
}

func(np *PppoL2tpServerPool) SetAcceptedSutNameOrPeerIdModifier () error {
 //parameters: DeviceHandle Value Repeat Increment
 //AgtPppoL2tpServerPool SetAcceptedSutNameOrPeerIdModifier, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetAcceptedSutNameOrPeerIdModifier ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetAcceptedSutNameOrPeerIdModifier
 return "", nil
}

func(np *PppoL2tpServerPool) SetAcceptedSutPasswordOrSecretModifier () error {
 //parameters: DeviceHandle Value Repeat Increment
 //AgtPppoL2tpServerPool SetAcceptedSutPasswordOrSecretModifier, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetAcceptedSutPasswordOrSecretModifier ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetAcceptedSutPasswordOrSecretModifier
 return "", nil
}

func(np *PppoL2tpServerPool) ListSelectedPools ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool ListSelectedPools
 return "", nil
}

func(np *PppoL2tpServerPool) SelectPools () error {
 //parameters: SessionIdentifiers
 //AgtPppoL2tpServerPool SelectPools, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) DeselectPools () error {
 //parameters: SessionIdentifiers
 //AgtPppoL2tpServerPool DeselectPools, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetAccumulatedValues ()(string, error) {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool GetAccumulatedValues
 return "", nil
}

func(np *PppoL2tpServerPool) GetDeviceAccumulatedValues ()(string, error) {
 //parameters: DeviceHandle DeviceStatisticsList
 //AgtPppoL2tpServerPool GetDeviceAccumulatedValues
 return "", nil
}

func(np *PppoL2tpServerPool) GetSessionAccumulatedSpecifiedValues ()(string, error) {
 //parameters: SessionIdentifiers SessionStatisticsList
 //AgtPppoL2tpServerPool GetSessionAccumulatedSpecifiedValues
 return "", nil
}

func(np *PppoL2tpServerPool) ClearStatistics () error {
 //parameters: DeviceHandle
 //AgtPppoL2tpServerPool ClearStatistics, m.Object, m.Name);
 return nil
}

func(np *PppoL2tpServerPool) GetLastFailureReason ()(string, error) {
 //parameters: SessionIdentifiers
 //AgtPppoL2tpServerPool GetLastFailureReason
 return "", nil
}

