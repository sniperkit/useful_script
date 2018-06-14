package n2xsdk

type PppSessionPoolEthernetOverAtm struct {
 Handler string
}

func(np *PppSessionPoolEthernetOverAtm) SetAtmEncapsulation () error {
 //parameters: SessionPoolHandle Encapsulation
 //AgtPppSessionPoolEthernetOverAtm SetAtmEncapsulation, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetNumberOfSessionsPerSrcMacAddr () error {
 //parameters: SessionPoolHandle SessionsPerMacAddr
 //AgtPppSessionPoolEthernetOverAtm SetNumberOfSessionsPerSrcMacAddr, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetInitialAckTimeout () error {
 //parameters: SessionPoolHandle Timeout
 //AgtPppSessionPoolEthernetOverAtm SetInitialAckTimeout, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetMaxTxRetries () error {
 //parameters: SessionPoolHandle Retries
 //AgtPppSessionPoolEthernetOverAtm SetMaxTxRetries, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetServiceName () error {
 //parameters: SessionPoolHandle ServiceName
 //AgtPppSessionPoolEthernetOverAtm SetServiceName, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetAccessConcentratorName () error {
 //parameters: SessionPoolHandle AccessConcentratorName
 //AgtPppSessionPoolEthernetOverAtm SetAccessConcentratorName, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetAccessConcentratorName ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetAccessConcentratorName
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) EnableEchoVendorSpecificTag () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm EnableEchoVendorSpecificTag, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) DisableEchoVendorSpecificTag () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm DisableEchoVendorSpecificTag, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) IsEchoVendorSpecificTagEnabled () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm IsEchoVendorSpecificTagEnabled, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetStartingSourceMacAddress () error {
 //parameters: SessionPoolHandle MacAddress
 //AgtPppSessionPoolEthernetOverAtm SetStartingSourceMacAddress, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetSourceMacAddressIncrement () error {
 //parameters: SessionPoolHandle Increment
 //AgtPppSessionPoolEthernetOverAtm SetSourceMacAddressIncrement, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetStartingVpiVci () error {
 //parameters: SessionPoolHandle VpiStart VciStart
 //AgtPppSessionPoolEthernetOverAtm SetStartingVpiVci, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetVpiVciIncrement () error {
 //parameters: SessionPoolHandle VpiIncrement VciIncrement
 //AgtPppSessionPoolEthernetOverAtm SetVpiVciIncrement, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetVpiVciCountOverride ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetVpiVciCountOverride
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) SetVpiVciCountOverride () error {
 //parameters: SessionPoolHandle IsVpiCountOverride IsVciCountOverride
 //AgtPppSessionPoolEthernetOverAtm SetVpiVciCountOverride, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetVpiVciRange ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetVpiVciRange
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) SetVpiVciRange () error {
 //parameters: SessionPoolHandle VpiStart VpiCount VpiIncrement VpiRepeat VciStart VciCount VciIncrement VciRepeat
 //AgtPppSessionPoolEthernetOverAtm SetVpiVciRange, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetServiceName ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetServiceName
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetNumberOfSessionsPerSrcMacAddr ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetNumberOfSessionsPerSrcMacAddr
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetStartingSourceMacAddress ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetStartingSourceMacAddress
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetSourceMacAddressIncrement ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetSourceMacAddressIncrement
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetInitialAckTimeout ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetInitialAckTimeout
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetMaxTxRetries ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetMaxTxRetries
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetVlanFlag ()(string, error) {
 //parameters: SessionPoolHandle VlanTagNumber
 //AgtPppSessionPoolEthernetOverAtm GetVlanFlag
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetEthertype ()(string, error) {
 //parameters: SessionPoolHandle VlanTagNumber
 //AgtPppSessionPoolEthernetOverAtm GetEthertype
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetVlanId ()(string, error) {
 //parameters: SessionPoolHandle VlanTagNumber
 //AgtPppSessionPoolEthernetOverAtm GetVlanId
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetDestinationMacAddresses ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetDestinationMacAddresses
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetSessionIds ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetSessionIds
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetAtmEncapsulation ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetAtmEncapsulation
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetStartingVpiVci ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetStartingVpiVci
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetVpiVciIncrement ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetVpiVciIncrement
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) Enable () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm Enable, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) Disable () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm Disable, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetNumberOfSessions () error {
 //parameters: SessionPoolHandle NumSessions
 //AgtPppSessionPoolEthernetOverAtm SetNumberOfSessions, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetSessionLifetime () error {
 //parameters: SessionPoolHandle MinLifetime MaxLifetime
 //AgtPppSessionPoolEthernetOverAtm SetSessionLifetime, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetSessionInitiationRate () error {
 //parameters: SessionPoolHandle NumSessions Period
 //AgtPppSessionPoolEthernetOverAtm SetSessionInitiationRate, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetLcpOption () error {
 //parameters: SessionPoolHandle LcpOption Value
 //AgtPppSessionPoolEthernetOverAtm SetLcpOption, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) EnableAccm () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm EnableAccm, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) DisableAccm () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm DisableAccm, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) EnableAddressControlFieldCompression () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm EnableAddressControlFieldCompression, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) DisableAddressControlFieldCompression () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm DisableAddressControlFieldCompression, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetMode () error {
 //parameters: SessionPoolHandle Mode
 //AgtPppSessionPoolEthernetOverAtm SetMode, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetAuthenticationProtocol () error {
 //parameters: SessionPoolHandle Protocol
 //AgtPppSessionPoolEthernetOverAtm SetAuthenticationProtocol, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetAuthenticationParameter () error {
 //parameters: SessionPoolHandle Parameter Value
 //AgtPppSessionPoolEthernetOverAtm SetAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) EnableUniqueAuthenticationParameter () error {
 //parameters: SessionPoolHandle Parameter
 //AgtPppSessionPoolEthernetOverAtm EnableUniqueAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) DisableUniqueAuthenticationParameter () error {
 //parameters: SessionPoolHandle Parameter
 //AgtPppSessionPoolEthernetOverAtm DisableUniqueAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetAuthenticationParameterStartValue () error {
 //parameters: SessionPoolHandle Parameter StartValue
 //AgtPppSessionPoolEthernetOverAtm SetAuthenticationParameterStartValue, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetAuthenticationParameter ()(string, error) {
 //parameters: SessionPoolHandle Parameter
 //AgtPppSessionPoolEthernetOverAtm GetAuthenticationParameter
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) SetPeerAuthenticationParameter () error {
 //parameters: SessionPoolHandle PeerParameter Value
 //AgtPppSessionPoolEthernetOverAtm SetPeerAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetPeerAuthenticationParameter ()(string, error) {
 //parameters: SessionPoolHandle PeerParameter
 //AgtPppSessionPoolEthernetOverAtm GetPeerAuthenticationParameter
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) EnableAcceptAllAuthenticationValues () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm EnableAcceptAllAuthenticationValues, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) DisableAcceptAllAuthenticationValues () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm DisableAcceptAllAuthenticationValues, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) IsAcceptAllAuthenticationValuesEnabled () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm IsAcceptAllAuthenticationValuesEnabled, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetUserToDomainRatio () error {
 //parameters: SessionPoolHandle UserToDomainRatio
 //AgtPppSessionPoolEthernetOverAtm SetUserToDomainRatio, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetUserToDomainRatio ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetUserToDomainRatio
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) SetPasswordIncrementMethod () error {
 //parameters: SessionPoolHandle PasswordIncrementMethod
 //AgtPppSessionPoolEthernetOverAtm SetPasswordIncrementMethod, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetPasswordIncrementMethod ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetPasswordIncrementMethod
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) SetDomainCountLimit () error {
 //parameters: SessionPoolHandle DomainCountLimit
 //AgtPppSessionPoolEthernetOverAtm SetDomainCountLimit, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetDomainCountLimit ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetDomainCountLimit
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) EnableDomainNameIncrementAfterExtension () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm EnableDomainNameIncrementAfterExtension, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) DisableDomainNameIncrementAfterExtension () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm DisableDomainNameIncrementAfterExtension, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) IsDomainNameIncrementAfterExtensionEnabled () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm IsDomainNameIncrementAfterExtensionEnabled, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetAcceptedAuthenticationParameter () error {
 //parameters: SessionPoolHandle AcceptedParameter Value
 //AgtPppSessionPoolEthernetOverAtm SetAcceptedAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetAcceptedAuthenticationParameter ()(string, error) {
 //parameters: SessionPoolHandle AcceptedParameter
 //AgtPppSessionPoolEthernetOverAtm GetAcceptedAuthenticationParameter
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) SetStartingSourceIpAddress () error {
 //parameters: SessionPoolHandle IpAddress
 //AgtPppSessionPoolEthernetOverAtm SetStartingSourceIpAddress, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetSourceIpAddressModifier () error {
 //parameters: SessionPoolHandle Modifier
 //AgtPppSessionPoolEthernetOverAtm SetSourceIpAddressModifier, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) SetMaxConnectionAttempts () error {
 //parameters: SessionPoolHandle MaxConnectionAttempts
 //AgtPppSessionPoolEthernetOverAtm SetMaxConnectionAttempts, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetMaxConnectionAttempts ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetMaxConnectionAttempts
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) RetryConnections () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm RetryConnections, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetConnectionAttempts ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetConnectionAttempts
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) Cancel () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm Cancel, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetAccmFlag ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetAccmFlag
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetAddressControlFieldCompressionFlag ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetAddressControlFieldCompressionFlag
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetAuthenticationParameterStartValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetAuthenticationParameterStartValue
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetAuthenticationProtocol ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetAuthenticationProtocol
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetEnableFlag ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetEnableFlag
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetLcpOption ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetLcpOption
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetMode ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetMode
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetName ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetName
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) SetName () error {
 //parameters: SessionPoolHandle Name
 //AgtPppSessionPoolEthernetOverAtm SetName, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) GetNegotiatedIpAddresses ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetNegotiatedIpAddresses
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetNumberOfEstablishedSessions ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetNumberOfEstablishedSessions
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetNumberOfSessions ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetNumberOfSessions
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetRemoteIpAddresses ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetRemoteIpAddresses
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetSessionInitiationRate ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetSessionInitiationRate
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetSessionLifetime ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetSessionLifetime
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetSessionStateDescriptions ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetSessionStateDescriptions
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetSourceIpAddressModifier ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetSourceIpAddressModifier
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetStartingSourceIpAddress ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetStartingSourceIpAddress
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetState ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPoolEthernetOverAtm GetState
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtPppSessionPoolEthernetOverAtm GetStateSummary
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) GetUniqueAuthenticationParameterFlag ()(string, error) {
 //parameters: SessionPoolHandle Parameter
 //AgtPppSessionPoolEthernetOverAtm GetUniqueAuthenticationParameterFlag
 return "", nil
}

func(np *PppSessionPoolEthernetOverAtm) EnableAllSessions () error {
 //parameters: 
 //AgtPppSessionPoolEthernetOverAtm EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *PppSessionPoolEthernetOverAtm) DisableAllSessions () error {
 //parameters: 
 //AgtPppSessionPoolEthernetOverAtm DisableAllSessions, m.Object, m.Name);
 return nil
}

