package n2xsdk

type PimPdu struct {
 Handler string
}

func(np *PimPdu) Prohibit () error {
 //parameters: SessionHandle PduType
 //AgtPimPdu Prohibit, m.Object, m.Name);
 return nil
}

func(np *PimPdu) Permit () error {
 //parameters: SessionHandle PduType
 //AgtPimPdu Permit, m.Object, m.Name);
 return nil
}

func(np *PimPdu) BuildPdu () error {
 //parameters: SessionHandle PduType
 //AgtPimPdu BuildPdu, m.Object, m.Name);
 return nil
}

func(np *PimPdu) BuildPduUsingId () error {
 //parameters: PduType ReferencePduId
 //AgtPimPdu BuildPduUsingId, m.Object, m.Name);
 return nil
}

func(np *PimPdu) BuildPduUsingContents () error {
 //parameters: SessionHandle PduContents
 //AgtPimPdu BuildPduUsingContents, m.Object, m.Name);
 return nil
}

func(np *PimPdu) DeletePdu () error {
 //parameters: PduId
 //AgtPimPdu DeletePdu, m.Object, m.Name);
 return nil
}

func(np *PimPdu) DeleteAllPdusInSession () error {
 //parameters: SessionHandle
 //AgtPimPdu DeleteAllPdusInSession, m.Object, m.Name);
 return nil
}

func(np *PimPdu) DeleteAllPdus () error {
 //parameters: 
 //AgtPimPdu DeleteAllPdus, m.Object, m.Name);
 return nil
}

func(np *PimPdu) DecodePdu () error {
 //parameters: FormatMode PduId
 //AgtPimPdu DecodePdu, m.Object, m.Name);
 return nil
}

func(np *PimPdu) GetFieldCount ()(string, error) {
 //parameters: PduId FieldType
 //AgtPimPdu GetFieldCount
 return "", nil
}

func(np *PimPdu) GetSubfield ()(string, error) {
 //parameters: PduId FieldType Index SubfieldType
 //AgtPimPdu GetSubfield
 return "", nil
}

func(np *PimPdu) SetSubfield () error {
 //parameters: PduId FieldType Index SubfieldType SubfieldValue
 //AgtPimPdu SetSubfield, m.Object, m.Name);
 return nil
}

func(np *PimPdu) DeleteField () error {
 //parameters: PduId FieldType Index
 //AgtPimPdu DeleteField, m.Object, m.Name);
 return nil
}

func(np *PimPdu) SendPdu () error {
 //parameters: DestinationIpAddress PduId
 //AgtPimPdu SendPdu, m.Object, m.Name);
 return nil
}

func(np *PimPdu) GetNextPduInfo ()(string, error) {
 //parameters: TimeOutSec Count psaSessionHandle
 //AgtPimPdu GetNextPduInfo
 return "", nil
}

func(np *PimPdu) GetPduDetails ()(string, error) {
 //parameters: PduId PduInfoType
 //AgtPimPdu GetPduDetails
 return "", nil
}

func(np *PimPdu) IsProhibited () error {
 //parameters: SessionHandle PduType
 //AgtPimPdu IsProhibited, m.Object, m.Name);
 return nil
}

func(np *PimPdu) AddPduFilter () error {
 //parameters: SessionHandle PduType ModifierType SourceIpAddress GroupIpAddress
 //AgtPimPdu AddPduFilter, m.Object, m.Name);
 return nil
}

func(np *PimPdu) RemovePduFilter () error {
 //parameters: SessionHandle FilterId
 //AgtPimPdu RemovePduFilter, m.Object, m.Name);
 return nil
}

func(np *PimPdu) RemoveAllPduFilters () error {
 //parameters: SessionHandle
 //AgtPimPdu RemoveAllPduFilters, m.Object, m.Name);
 return nil
}

func(np *PimPdu) ListPduFilters ()(string, error) {
 //parameters: SessionHandle
 //AgtPimPdu ListPduFilters
 return "", nil
}

func(np *PimPdu) GetPduFilterDetails ()(string, error) {
 //parameters: SessionHandle FilterId
 //AgtPimPdu GetPduFilterDetails
 return "", nil
}

func(np *PimPdu) StartPduSave () error {
 //parameters: SessionHandle
 //AgtPimPdu StartPduSave, m.Object, m.Name);
 return nil
}

func(np *PimPdu) StopPduSave () error {
 //parameters: SessionHandle
 //AgtPimPdu StopPduSave, m.Object, m.Name);
 return nil
}

func(np *PimPdu) IsPduSaveEnabled () error {
 //parameters: SessionHandle
 //AgtPimPdu IsPduSaveEnabled, m.Object, m.Name);
 return nil
}

func(np *PimPdu) EnablePduNotification () error {
 //parameters: SessionHandle
 //AgtPimPdu EnablePduNotification, m.Object, m.Name);
 return nil
}

func(np *PimPdu) DisablePduNotification () error {
 //parameters: SessionHandle
 //AgtPimPdu DisablePduNotification, m.Object, m.Name);
 return nil
}

func(np *PimPdu) IsPduNotificationEnabled () error {
 //parameters: SessionHandle
 //AgtPimPdu IsPduNotificationEnabled, m.Object, m.Name);
 return nil
}

func(np *PimPdu) ListSavedPdus ()(string, error) {
 //parameters: SessionHandle
 //AgtPimPdu ListSavedPdus
 return "", nil
}

