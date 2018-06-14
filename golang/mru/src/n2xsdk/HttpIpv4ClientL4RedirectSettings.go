package n2xsdk

type HttpIpv4ClientL4RedirectSettings struct {
 Handler string
}

func(np *HttpIpv4ClientL4RedirectSettings) EnableL4RedirectionSettings () error {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings EnableL4RedirectionSettings, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) DisableL4RedirectionSettings () error {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings DisableL4RedirectionSettings, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) IsL4RedirectionSettingsEnabled () error {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings IsL4RedirectionSettingsEnabled, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) SetAuthenticationLevel () error {
 //parameters: SessionPoolHandle AuthenticationLevel
 //AgtHttpIpv4ClientL4RedirectSettings SetAuthenticationLevel, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetAuthenticationLevel ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetAuthenticationLevel
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) SetFirstCredentialNameValue () error {
 //parameters: SessionPoolHandle CredentialFieldName CredentialFieldValue
 //AgtHttpIpv4ClientL4RedirectSettings SetFirstCredentialNameValue, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetFirstCredentialNameValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetFirstCredentialNameValue
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetFirstCredentialFieldValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetFirstCredentialFieldValue
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) SetFirstCredentialValueModifierIncrementor () error {
 //parameters: SessionPoolHandle FieldValueModifier IncrementingRange Repeat
 //AgtHttpIpv4ClientL4RedirectSettings SetFirstCredentialValueModifierIncrementor, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetFirstCredentialValueModifierIncrementor ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetFirstCredentialValueModifierIncrementor
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) SetSecondCredentialNameValue () error {
 //parameters: SessionPoolHandle CredentialFieldName CredentialFieldValue
 //AgtHttpIpv4ClientL4RedirectSettings SetSecondCredentialNameValue, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetSecondCredentialNameValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetSecondCredentialNameValue
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetSecondCredentialFieldValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetSecondCredentialFieldValue
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) SetSecondCredentialValueModifierIncrementor () error {
 //parameters: SessionPoolHandle FieldValueModifier IncrementingRange Repeat
 //AgtHttpIpv4ClientL4RedirectSettings SetSecondCredentialValueModifierIncrementor, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetSecondCredentialValueModifierIncrementor ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetSecondCredentialValueModifierIncrementor
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) AddAuthenticationParameter () error {
 //parameters: SessionPoolHandle RowIndex
 //AgtHttpIpv4ClientL4RedirectSettings AddAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetNumAuthenticationParameters ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetNumAuthenticationParameters
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) ListAuthenticationParameters ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings ListAuthenticationParameters
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) RemoveAuthenticationParameter () error {
 //parameters: SessionPoolHandle RowIndex
 //AgtHttpIpv4ClientL4RedirectSettings RemoveAuthenticationParameter, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) RemoveAllAuthenticationParameters () error {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings RemoveAllAuthenticationParameters, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) SetAuthenticationParameterValue () error {
 //parameters: SessionPoolHandle RowIndex AuthenticationParameterValue
 //AgtHttpIpv4ClientL4RedirectSettings SetAuthenticationParameterValue, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetAuthenticationParameterValue ()(string, error) {
 //parameters: SessionPoolHandle RowIndex
 //AgtHttpIpv4ClientL4RedirectSettings GetAuthenticationParameterValue
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) SetSuccessIndicatorType () error {
 //parameters: SessionPoolHandle SuccessIndicatorType
 //AgtHttpIpv4ClientL4RedirectSettings SetSuccessIndicatorType, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetSuccessIndicatorType ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetSuccessIndicatorType
 return "", nil
}

func(np *HttpIpv4ClientL4RedirectSettings) SetSuccessIndicatorValue () error {
 //parameters: SessionPoolHandle SuccessIndicatorValue
 //AgtHttpIpv4ClientL4RedirectSettings SetSuccessIndicatorValue, m.Object, m.Name);
 return nil
}

func(np *HttpIpv4ClientL4RedirectSettings) GetSuccessIndicatorValue ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtHttpIpv4ClientL4RedirectSettings GetSuccessIndicatorValue
 return "", nil
}

