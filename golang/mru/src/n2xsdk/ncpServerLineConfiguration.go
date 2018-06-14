package n2xsdk

type ncpServerLineConfiguration struct {
 Handler string
}

func(np *ncpServerLineConfiguration) DisableAutoLineConfigPortManagement () error {
 //parameters: SessionPoolHandle
 //AgtAncpServerLineConfiguration DisableAutoLineConfigPortManagement, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) EnableAutoLineConfigPortManagement () error {
 //parameters: SessionPoolHandle
 //AgtAncpServerLineConfiguration EnableAutoLineConfigPortManagement, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) IsAutoLineConfigPortManagementEnabled () error {
 //parameters: SessionPoolHandle
 //AgtAncpServerLineConfiguration IsAutoLineConfigPortManagementEnabled, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) AddProfileEntry () error {
 //parameters: SessionPoolHandle CircuitId ProfileName
 //AgtAncpServerLineConfiguration AddProfileEntry, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) ListProfileEntries ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtAncpServerLineConfiguration ListProfileEntries
 return "", nil
}

func(np *ncpServerLineConfiguration) RemoveProfileEntry () error {
 //parameters: SessionPoolHandle CircuitId
 //AgtAncpServerLineConfiguration RemoveProfileEntry, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) RemoveAllProfileEntries () error {
 //parameters: SessionPoolHandle
 //AgtAncpServerLineConfiguration RemoveAllProfileEntries, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) GetProfileName ()(string, error) {
 //parameters: SessionPoolHandle CircuitId
 //AgtAncpServerLineConfiguration GetProfileName
 return "", nil
}

func(np *ncpServerLineConfiguration) SetDefaultProfileName () error {
 //parameters: SessionPoolHandle ProfileName
 //AgtAncpServerLineConfiguration SetDefaultProfileName, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) GetDefaultProfileName ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtAncpServerLineConfiguration GetDefaultProfileName
 return "", nil
}

func(np *ncpServerLineConfiguration) SetResultField () error {
 //parameters: SessionPoolHandle AncpResult
 //AgtAncpServerLineConfiguration SetResultField, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) GetResultField ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtAncpServerLineConfiguration GetResultField
 return "", nil
}

func(np *ncpServerLineConfiguration) AddTlv () error {
 //parameters: SessionPoolHandle TlvType TlvLength TlvValue
 //AgtAncpServerLineConfiguration AddTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) ListTlvs ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtAncpServerLineConfiguration ListTlvs
 return "", nil
}

func(np *ncpServerLineConfiguration) RemoveTlv () error {
 //parameters: SessionPoolHandle TlvType
 //AgtAncpServerLineConfiguration RemoveTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) RemoveAllTlvs () error {
 //parameters: SessionPoolHandle
 //AgtAncpServerLineConfiguration RemoveAllTlvs, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) SetTlv () error {
 //parameters: SessionPoolHandle TlvType TlvLength TlvValue
 //AgtAncpServerLineConfiguration SetTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) GetTlv ()(string, error) {
 //parameters: SessionPoolHandle TlvType
 //AgtAncpServerLineConfiguration GetTlv
 return "", nil
}

func(np *ncpServerLineConfiguration) AddSubTlv () error {
 //parameters: SessionPoolHandle TlvType SubTlvType SubTlvLength SubTlvValue
 //AgtAncpServerLineConfiguration AddSubTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) ListSubTlvs ()(string, error) {
 //parameters: SessionPoolHandle TlvType
 //AgtAncpServerLineConfiguration ListSubTlvs
 return "", nil
}

func(np *ncpServerLineConfiguration) RemoveSubTlv () error {
 //parameters: SessionPoolHandle TlvType SubTlvType
 //AgtAncpServerLineConfiguration RemoveSubTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) RemoveAllSubTlvs () error {
 //parameters: SessionPoolHandle TlvType
 //AgtAncpServerLineConfiguration RemoveAllSubTlvs, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) SetSubTlv () error {
 //parameters: SessionPoolHandle TlvType SubTlvType SubTlvLength SubTlvValue
 //AgtAncpServerLineConfiguration SetSubTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerLineConfiguration) GetSubTlv ()(string, error) {
 //parameters: SessionPoolHandle TlvType SubTlvType
 //AgtAncpServerLineConfiguration GetSubTlv
 return "", nil
}

