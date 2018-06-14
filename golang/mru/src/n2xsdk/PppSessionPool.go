package n2xsdk

type PppSessionPool struct {
 Handler string
}

func(np *PppSessionPool) Enable () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool Enable, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) Disable () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool Disable, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetNumberOfSessions () error {
 //parameters: SessionPoolHandle NumSessions
 //AgtPppSessionPool SetNumberOfSessions, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetSessionLifetime () error {
 //parameters: SessionPoolHandle MinLifetime MaxLifetime
 //AgtPppSessionPool SetSessionLifetime, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetSessionInitiationRate () error {
 //parameters: SessionPoolHandle NumSessions Period
 //AgtPppSessionPool SetSessionInitiationRate, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetLcpOption () error {
 //parameters: SessionPoolHandle LcpOption Value
 //AgtPppSessionPool SetLcpOption, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) EnableAccm () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool EnableAccm, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) DisableAccm () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool DisableAccm, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) EnableAddressControlFieldCompression () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool EnableAddressControlFieldCompression, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) DisableAddressControlFieldCompression () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool DisableAddressControlFieldCompression, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetMode () error {
 //parameters: SessionPoolHandle Mode
 //AgtPppSessionPool SetMode, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetAuthenticationProtocol () error {
 //parameters: SessionPoolHandle Protocol
 //AgtPppSessionPool SetAuthenticationProtocol, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetAuthenticationParameter () error {
 //parameters: SessionPoolHandle Parameter Value
 //AgtPppSessionPool SetAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) EnableUniqueAuthenticationParameter () error {
 //parameters: SessionPoolHandle Parameter
 //AgtPppSessionPool EnableUniqueAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) DisableUniqueAuthenticationParameter () error {
 //parameters: SessionPoolHandle Parameter
 //AgtPppSessionPool DisableUniqueAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetAuthenticationParameterStartValue () error {
 //parameters: SessionPoolHandle Parameter StartValue
 //AgtPppSessionPool SetAuthenticationParameterStartValue, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetAuthenticationParameter ()(string, error) {
 //parameters: SessionPoolHandle Parameter
 //AgtPppSessionPool GetAuthenticationParameter
 return "", nil
}

func(np *PppSessionPool) SetPeerAuthenticationParameter () error {
 //parameters: SessionPoolHandle PeerParameter Value
 //AgtPppSessionPool SetPeerAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetPeerAuthenticationParameter ()(string, error) {
 //parameters: SessionPoolHandle PeerParameter
 //AgtPppSessionPool GetPeerAuthenticationParameter
 return "", nil
}

func(np *PppSessionPool) EnableAcceptAllAuthenticationValues () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool EnableAcceptAllAuthenticationValues, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) DisableAcceptAllAuthenticationValues () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool DisableAcceptAllAuthenticationValues, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) IsAcceptAllAuthenticationValuesEnabled () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool IsAcceptAllAuthenticationValuesEnabled, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetUserToDomainRatio () error {
 //parameters: SessionPoolHandle UserToDomainRatio
 //AgtPppSessionPool SetUserToDomainRatio, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetUserToDomainRatio ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetUserToDomainRatio
 return "", nil
}

func(np *PppSessionPool) SetPasswordIncrementMethod () error {
 //parameters: SessionPoolHandle PasswordIncrementMethod
 //AgtPppSessionPool SetPasswordIncrementMethod, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetPasswordIncrementMethod ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetPasswordIncrementMethod
 return "", nil
}

func(np *PppSessionPool) SetDomainCountLimit () error {
 //parameters: SessionPoolHandle DomainCountLimit
 //AgtPppSessionPool SetDomainCountLimit, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetDomainCountLimit ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetDomainCountLimit
 return "", nil
}

func(np *PppSessionPool) EnableDomainNameIncrementAfterExtension () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool EnableDomainNameIncrementAfterExtension, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) DisableDomainNameIncrementAfterExtension () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool DisableDomainNameIncrementAfterExtension, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) IsDomainNameIncrementAfterExtensionEnabled () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool IsDomainNameIncrementAfterExtensionEnabled, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetAcceptedAuthenticationParameter () error {
 //parameters: SessionPoolHandle AcceptedParameter Value
 //AgtPppSessionPool SetAcceptedAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetAcceptedAuthenticationParameter ()(string, error) {
 //parameters: SessionPoolHandle AcceptedParameter
 //AgtPppSessionPool GetAcceptedAuthenticationParameter
 return "", nil
}

func(np *PppSessionPool) SetStartingSourceIpAddress () error {
 //parameters: SessionPoolHandle IpAddress
 //AgtPppSessionPool SetStartingSourceIpAddress, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetSourceIpAddressModifier () error {
 //parameters: SessionPoolHandle Modifier
 //AgtPppSessionPool SetSourceIpAddressModifier, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) SetMaxConnectionAttempts () error {
 //parameters: SessionPoolHandle MaxConnectionAttempts
 //AgtPppSessionPool SetMaxConnectionAttempts, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetMaxConnectionAttempts ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetMaxConnectionAttempts
 return "", nil
}

func(np *PppSessionPool) RetryConnections () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool RetryConnections, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetConnectionAttempts ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetConnectionAttempts
 return "", nil
}

func(np *PppSessionPool) Cancel () error {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool Cancel, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetAccmFlag ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetAccmFlag
 return "", nil
}

func(np *PppSessionPool) GetAddressControlFieldCompressionFlag ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetAddressControlFieldCompressionFlag
 return "", nil
}

func(np *PppSessionPool) GetAuthenticationParameterStartValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetAuthenticationParameterStartValue
 return "", nil
}

func(np *PppSessionPool) GetAuthenticationProtocol ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetAuthenticationProtocol
 return "", nil
}

func(np *PppSessionPool) GetEnableFlag ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetEnableFlag
 return "", nil
}

func(np *PppSessionPool) GetLcpOption ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetLcpOption
 return "", nil
}

func(np *PppSessionPool) GetMode ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetMode
 return "", nil
}

func(np *PppSessionPool) GetName ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetName
 return "", nil
}

func(np *PppSessionPool) SetName () error {
 //parameters: SessionPoolHandle Name
 //AgtPppSessionPool SetName, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) GetNegotiatedIpAddresses ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetNegotiatedIpAddresses
 return "", nil
}

func(np *PppSessionPool) GetNumberOfEstablishedSessions ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetNumberOfEstablishedSessions
 return "", nil
}

func(np *PppSessionPool) GetNumberOfSessions ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetNumberOfSessions
 return "", nil
}

func(np *PppSessionPool) GetRemoteIpAddresses ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetRemoteIpAddresses
 return "", nil
}

func(np *PppSessionPool) GetSessionInitiationRate ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetSessionInitiationRate
 return "", nil
}

func(np *PppSessionPool) GetSessionLifetime ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetSessionLifetime
 return "", nil
}

func(np *PppSessionPool) GetSessionStateDescriptions ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetSessionStateDescriptions
 return "", nil
}

func(np *PppSessionPool) GetSourceIpAddressModifier ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetSourceIpAddressModifier
 return "", nil
}

func(np *PppSessionPool) GetStartingSourceIpAddress ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetStartingSourceIpAddress
 return "", nil
}

func(np *PppSessionPool) GetState ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppSessionPool GetState
 return "", nil
}

func(np *PppSessionPool) GetStateSummary ()(string, error) {
 //parameters: 
 //AgtPppSessionPool GetStateSummary
 return "", nil
}

func(np *PppSessionPool) GetUniqueAuthenticationParameterFlag ()(string, error) {
 //parameters: SessionPoolHandle Parameter
 //AgtPppSessionPool GetUniqueAuthenticationParameterFlag
 return "", nil
}

func(np *PppSessionPool) EnableAllSessions () error {
 //parameters: 
 //AgtPppSessionPool EnableAllSessions, m.Object, m.Name);
 return nil
}

func(np *PppSessionPool) DisableAllSessions () error {
 //parameters: 
 //AgtPppSessionPool DisableAllSessions, m.Object, m.Name);
 return nil
}

