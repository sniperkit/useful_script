package n2xsdk

type SipUserAgentPool struct {
 Handler string
}

func(np *SipUserAgentPool) SetOutboundProxyIpv4Address () error {
 //parameters: DeviceHandle Ipv4Address Increment Repeat
 //AgtSipUserAgentPool SetOutboundProxyIpv4Address, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetOutboundProxyIpv4Address ()(string, error) {
 //parameters: DeviceHandle
 //AgtSipUserAgentPool GetOutboundProxyIpv4Address
 return "", nil
}

func(np *SipUserAgentPool) IsNetworkAddressTranslationEnabled () error {
 //parameters: DeviceHandle
 //AgtSipUserAgentPool IsNetworkAddressTranslationEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableNetworkAddressTranslation () error {
 //parameters: DeviceHandle
 //AgtSipUserAgentPool EnableNetworkAddressTranslation, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableNetworkAddressTranslation () error {
 //parameters: DeviceHandle
 //AgtSipUserAgentPool DisableNetworkAddressTranslation, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetNatedInternaIpv4AddressRange () error {
 //parameters: DeviceHandle Ipv4Address Increment Repeat
 //AgtSipUserAgentPool SetNatedInternaIpv4AddressRange, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetNatedInternaIpv4AddressRange ()(string, error) {
 //parameters: DeviceHandle
 //AgtSipUserAgentPool GetNatedInternaIpv4AddressRange
 return "", nil
}

func(np *SipUserAgentPool) SetNatedExternalIpv4AddressRange () error {
 //parameters: DeviceHandle Ipv4Address Increment Repeat
 //AgtSipUserAgentPool SetNatedExternalIpv4AddressRange, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetNatedExternalIpv4AddressRange ()(string, error) {
 //parameters: DeviceHandle
 //AgtSipUserAgentPool GetNatedExternalIpv4AddressRange
 return "", nil
}

func(np *SipUserAgentPool) EstablishSession () error {
 //parameters: DeviceOrSession
 //AgtSipUserAgentPool EstablishSession, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) TerminateSession () error {
 //parameters: DeviceOrSession
 //AgtSipUserAgentPool TerminateSession, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) RegisterSession () error {
 //parameters: DeviceOrSession
 //AgtSipUserAgentPool RegisterSession, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DeregisterSession () error {
 //parameters: DeviceOrSession
 //AgtSipUserAgentPool DeregisterSession, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetLocalUsernamePrefix () error {
 //parameters: EmulationHandle LocalUserNamePrefix
 //AgtSipUserAgentPool SetLocalUsernamePrefix, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetLocalUsernameSuffixIncrementingRange () error {
 //parameters: EmulationHandle LocalUserNameSuffix Increment Repeat
 //AgtSipUserAgentPool SetLocalUsernameSuffixIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetLocalDomain () error {
 //parameters: EmulationHandle LocalDomain
 //AgtSipUserAgentPool SetLocalDomain, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetLocalDomain ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetLocalDomain
 return "", nil
}

func(np *SipUserAgentPool) GetLocalUsernamePrefix ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetLocalUsernamePrefix
 return "", nil
}

func(np *SipUserAgentPool) GetLocalUsernameSuffixIncrementingRange ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetLocalUsernameSuffixIncrementingRange
 return "", nil
}

func(np *SipUserAgentPool) IsUseRemoteAddressOfRecordEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsUseRemoteAddressOfRecordEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableUseRemoteAddressOfRecord () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableUseRemoteAddressOfRecord, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableUseRemoteAddressOfRecord () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableUseRemoteAddressOfRecord, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetRemoteUsernamePrefix () error {
 //parameters: EmulationHandle RemoteUserNamePrefix
 //AgtSipUserAgentPool SetRemoteUsernamePrefix, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetRemoteUsernameSuffixIncrementingRange () error {
 //parameters: EmulationHandle RemoteUserNameSuffix Increment Repeat
 //AgtSipUserAgentPool SetRemoteUsernameSuffixIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRemoteUsernameSuffixIncrementingRange ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRemoteUsernameSuffixIncrementingRange
 return "", nil
}

func(np *SipUserAgentPool) SetRemoteDomain () error {
 //parameters: EmulationHandle RemoteDomain
 //AgtSipUserAgentPool SetRemoteDomain, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRemoteUsernamePrefix ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRemoteUsernamePrefix
 return "", nil
}

func(np *SipUserAgentPool) GetRemoteDomain ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRemoteDomain
 return "", nil
}

func(np *SipUserAgentPool) SetRemoteHostAddressV4IncrementingRange () error {
 //parameters: EmulationHandle Ipv4Address Increment Repeat
 //AgtSipUserAgentPool SetRemoteHostAddressV4IncrementingRange, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetRemoteHostL3AddressFamily () error {
 //parameters: EmulationHandle RegistrarAddressType
 //AgtSipUserAgentPool SetRemoteHostL3AddressFamily, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRemoteHostL3AddressFamily ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRemoteHostL3AddressFamily
 return "", nil
}

func(np *SipUserAgentPool) GetRemoteHostAddressV4IncrementingRange ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRemoteHostAddressV4IncrementingRange
 return "", nil
}

func(np *SipUserAgentPool) SetRemoteHostAddressV6IncrementingRange () error {
 //parameters: EmulationHandle Ipv6Address Increment Repeat
 //AgtSipUserAgentPool SetRemoteHostAddressV6IncrementingRange, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRemoteHostAddressV6IncrementingRange ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRemoteHostAddressV6IncrementingRange
 return "", nil
}

func(np *SipUserAgentPool) SetMediaPayloadType () error {
 //parameters: EmulationHandle MediaPayloadType
 //AgtSipUserAgentPool SetMediaPayloadType, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetMediaPayloadType ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetMediaPayloadType
 return "", nil
}

func(np *SipUserAgentPool) SetMediaPort () error {
 //parameters: EmulationHandle Port
 //AgtSipUserAgentPool SetMediaPort, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetMediaPort ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetMediaPort
 return "", nil
}

func(np *SipUserAgentPool) SetSipPort () error {
 //parameters: EmulationHandle Port
 //AgtSipUserAgentPool SetSipPort, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetSipPort ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetSipPort
 return "", nil
}

func(np *SipUserAgentPool) IsRegistrarEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsRegistrarEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableRegistrar () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableRegistrar, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableRegistrar () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableRegistrar, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetRegistrarIpv4Address () error {
 //parameters: EmulationHandle Ipv4Address
 //AgtSipUserAgentPool SetRegistrarIpv4Address, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRegistrarIpv4Address ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRegistrarIpv4Address
 return "", nil
}

func(np *SipUserAgentPool) SetRegistrarIpv6Address () error {
 //parameters: EmulationHandle Ipv6Address
 //AgtSipUserAgentPool SetRegistrarIpv6Address, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRegistrarIpv6Address ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRegistrarIpv6Address
 return "", nil
}

func(np *SipUserAgentPool) SetRegistrarL3AddressFamily () error {
 //parameters: EmulationHandle L3AddressFamily
 //AgtSipUserAgentPool SetRegistrarL3AddressFamily, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRegistrarL3AddressFamily ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRegistrarL3AddressFamily
 return "", nil
}

func(np *SipUserAgentPool) SetRegistrarHostname () error {
 //parameters: EmulationHandle RegistrarHostname
 //AgtSipUserAgentPool SetRegistrarHostname, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRegistrarHostname ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRegistrarHostname
 return "", nil
}

func(np *SipUserAgentPool) IsAutomaticRegistrationEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsAutomaticRegistrationEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableAutomaticRegistration () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableAutomaticRegistration, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableAutomaticRegistration () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableAutomaticRegistration, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetDesiredRegistrationExpiryTime () error {
 //parameters: EmulationHandle DesiredRegistrationExpiryTime
 //AgtSipUserAgentPool SetDesiredRegistrationExpiryTime, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetDesiredRegistrationExpiryTime ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetDesiredRegistrationExpiryTime
 return "", nil
}

func(np *SipUserAgentPool) SetMaxRegistrationAttempts () error {
 //parameters: EmulationHandle MaxRegistrationAttempts
 //AgtSipUserAgentPool SetMaxRegistrationAttempts, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetMaxRegistrationAttempts ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetMaxRegistrationAttempts
 return "", nil
}

func(np *SipUserAgentPool) SetRegistrationAttemptDelayTime () error {
 //parameters: EmulationHandle RegistrationAttemptDelayTime
 //AgtSipUserAgentPool SetRegistrationAttemptDelayTime, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRegistrationAttemptDelayTime ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRegistrationAttemptDelayTime
 return "", nil
}

func(np *SipUserAgentPool) SetTransactionTimerT1 () error {
 //parameters: EmulationHandle Time
 //AgtSipUserAgentPool SetTransactionTimerT1, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetTransactionTimerT2 () error {
 //parameters: EmulationHandle Time
 //AgtSipUserAgentPool SetTransactionTimerT2, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetTransactionTimerT4 () error {
 //parameters: EmulationHandle Time
 //AgtSipUserAgentPool SetTransactionTimerT4, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetTransactionTimerT1 ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetTransactionTimerT1
 return "", nil
}

func(np *SipUserAgentPool) GetTransactionTimerT2 ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetTransactionTimerT2
 return "", nil
}

func(np *SipUserAgentPool) GetTransactionTimerT4 ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetTransactionTimerT4
 return "", nil
}

func(np *SipUserAgentPool) IsResourcePriorityEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsResourcePriorityEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableResourcePriority () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableResourcePriority, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableResourcePriority () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableResourcePriority, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetResourcePriority () error {
 //parameters: EmulationHandle ResourcePriority
 //AgtSipUserAgentPool SetResourcePriority, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetResourcePriority ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetResourcePriority
 return "", nil
}

func(np *SipUserAgentPool) IsSessionRefreshTimerEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsSessionRefreshTimerEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableSessionRefreshTimer () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableSessionRefreshTimer, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableSessionRefreshTimer () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableSessionRefreshTimer, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetDesiredSessionInterval () error {
 //parameters: EmulationHandle DesiredSessionInterval
 //AgtSipUserAgentPool SetDesiredSessionInterval, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetDesiredSessionInterval ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetDesiredSessionInterval
 return "", nil
}

func(np *SipUserAgentPool) GetMinimumSessionInterval ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetMinimumSessionInterval
 return "", nil
}

func(np *SipUserAgentPool) SetMinimumSessionInterval () error {
 //parameters: EmulationHandle MinimumSessionInterval
 //AgtSipUserAgentPool SetMinimumSessionInterval, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetRoundTripDelay () error {
 //parameters: EmulationHandle RoundTripInterval
 //AgtSipUserAgentPool SetRoundTripDelay, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetRoundTripDelay ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetRoundTripDelay
 return "", nil
}

func(np *SipUserAgentPool) SetSessionRefresher () error {
 //parameters: EmulationHandle UserAgentType
 //AgtSipUserAgentPool SetSessionRefresher, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetSessionRefresher ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetSessionRefresher
 return "", nil
}

func(np *SipUserAgentPool) IsCallAcceptDelayEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsCallAcceptDelayEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableCallAcceptDelay () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableCallAcceptDelay, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableCallAcceptDelay () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableCallAcceptDelay, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetResponseDelayInterval () error {
 //parameters: EmulationHandle Time
 //AgtSipUserAgentPool SetResponseDelayInterval, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetResponseDelayInterval ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetResponseDelayInterval
 return "", nil
}

func(np *SipUserAgentPool) IsPrivacyHeaderEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsPrivacyHeaderEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnablePrivacyHeader () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnablePrivacyHeader, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisablePrivacyHeader () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisablePrivacyHeader, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetPrivacyType () error {
 //parameters: EmulationHandle PrivacyType
 //AgtSipUserAgentPool SetPrivacyType, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetPrivacyType ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetPrivacyType
 return "", nil
}

func(np *SipUserAgentPool) IsAuthenticationEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsAuthenticationEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableAuthentication () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableAuthentication, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableAuthentication () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableAuthentication, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) IsAuthenticationCredentialEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsAuthenticationCredentialEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableAuthenticationCredential () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableAuthenticationCredential, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableAuthenticationCredential () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableAuthenticationCredential, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) IsCredentialInDialogEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsCredentialInDialogEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableCredentialInDialog () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableCredentialInDialog, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableCredentialInDialog () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableCredentialInDialog, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) SetAuthCredentialUsernamePrefix () error {
 //parameters: EmulationHandle AuthCredentialUsernamePrefix
 //AgtSipUserAgentPool SetAuthCredentialUsernamePrefix, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetAuthCredentialUsernamePrefix ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetAuthCredentialUsernamePrefix
 return "", nil
}

func(np *SipUserAgentPool) SetAuthCredentialUsernameSuffixIncrementingRange () error {
 //parameters: EmulationHandle AuthCredentialUsernameSuffix Increment Repeat
 //AgtSipUserAgentPool SetAuthCredentialUsernameSuffixIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetAuthCredentialUsernameSuffixIncrementingRange ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetAuthCredentialUsernameSuffixIncrementingRange
 return "", nil
}

func(np *SipUserAgentPool) SetAuthCredentialPassword () error {
 //parameters: EmulationHandle AuthCredentialPassword
 //AgtSipUserAgentPool SetAuthCredentialPassword, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetAuthCredentialPassword ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetAuthCredentialPassword
 return "", nil
}

func(np *SipUserAgentPool) SetAuthCredentialPasswordModifier () error {
 //parameters: EmulationHandle AuthCredentialPasswordModifier Increment Repeat
 //AgtSipUserAgentPool SetAuthCredentialPasswordModifier, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetAuthCredentialPasswordModifier ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetAuthCredentialPasswordModifier
 return "", nil
}

func(np *SipUserAgentPool) IsAuthenticationChallengeEnabled () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool IsAuthenticationChallengeEnabled, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) EnableAuthenticationChallenge () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool EnableAuthenticationChallenge, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) DisableAuthenticationChallenge () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool DisableAuthenticationChallenge, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetChallengeAuthenticationScheme ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetChallengeAuthenticationScheme
 return "", nil
}

func(np *SipUserAgentPool) SetChallengeAuthenticationRealm () error {
 //parameters: EmulationHandle ChallengeAuthenticationRealm
 //AgtSipUserAgentPool SetChallengeAuthenticationRealm, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetChallengeAuthenticationRealm ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetChallengeAuthenticationRealm
 return "", nil
}

func(np *SipUserAgentPool) SetAuthChallengeUsernamePrefix () error {
 //parameters: EmulationHandle AuthChallengeUsernamePrefix
 //AgtSipUserAgentPool SetAuthChallengeUsernamePrefix, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetAuthChallengeUsernamePrefix ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetAuthChallengeUsernamePrefix
 return "", nil
}

func(np *SipUserAgentPool) SetAuthChallengeUsernameSuffixIncrementingRange () error {
 //parameters: EmulationHandle AuthChallengeUsernameSuffix Increment Repeat
 //AgtSipUserAgentPool SetAuthChallengeUsernameSuffixIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetAuthChallengeUsernameSuffixIncrementingRange ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetAuthChallengeUsernameSuffixIncrementingRange
 return "", nil
}

func(np *SipUserAgentPool) SetAuthChallengePassword () error {
 //parameters: EmulationHandle AuthChallengePassword
 //AgtSipUserAgentPool SetAuthChallengePassword, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetAuthChallengePassword ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetAuthChallengePassword
 return "", nil
}

func(np *SipUserAgentPool) SetAuthChallengePasswordModifier () error {
 //parameters: EmulationHandle AuthChallengePasswordModifier Increment Repeat
 //AgtSipUserAgentPool SetAuthChallengePasswordModifier, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetAuthChallengePasswordModifier ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetAuthChallengePasswordModifier
 return "", nil
}

func(np *SipUserAgentPool) SetChallengeOpaque () error {
 //parameters: EmulationHandle ChallengeOpaque
 //AgtSipUserAgentPool SetChallengeOpaque, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetChallengeOpaque ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetChallengeOpaque
 return "", nil
}

func(np *SipUserAgentPool) SetChallengeAlgorithm () error {
 //parameters: EmulationHandle ChallengeAlgorithm
 //AgtSipUserAgentPool SetChallengeAlgorithm, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetChallengeAlgorithm ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetChallengeAlgorithm
 return "", nil
}

func(np *SipUserAgentPool) SetChallengeQop () error {
 //parameters: EmulationHandle ChallengeQop
 //AgtSipUserAgentPool SetChallengeQop, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetChallengeQop ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetChallengeQop
 return "", nil
}

func(np *SipUserAgentPool) SetNonceExpireTime () error {
 //parameters: EmulationHandle NonceExpireTime
 //AgtSipUserAgentPool SetNonceExpireTime, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetNonceExpireTime ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetNonceExpireTime
 return "", nil
}

func(np *SipUserAgentPool) SetMaximumChallengeAttempts () error {
 //parameters: EmulationHandle MaximumChallengeAttempts
 //AgtSipUserAgentPool SetMaximumChallengeAttempts, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) GetMaximumChallengeAttempts ()(string, error) {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool GetMaximumChallengeAttempts
 return "", nil
}

func(np *SipUserAgentPool) SendTrafficParams () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool SendTrafficParams, m.Object, m.Name);
 return nil
}

func(np *SipUserAgentPool) ResetTrafficParams () error {
 //parameters: EmulationHandle
 //AgtSipUserAgentPool ResetTrafficParams, m.Object, m.Name);
 return nil
}

