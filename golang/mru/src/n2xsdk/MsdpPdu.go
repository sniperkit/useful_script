package n2xsdk

type MsdpPdu struct {
 Handler string
}

func(np *MsdpPdu) BuildPdu () error {
 //parameters: SessionHandle PduType
 //AgtMsdpPdu BuildPdu, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) BuildPduUsingId () error {
 //parameters: PduType ReferencePduId
 //AgtMsdpPdu BuildPduUsingId, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) BuildPduUsingContents () error {
 //parameters: SessionHandle PduContents
 //AgtMsdpPdu BuildPduUsingContents, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) DeletePdu () error {
 //parameters: PduId
 //AgtMsdpPdu DeletePdu, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) DeleteAllPdusInSession () error {
 //parameters: SessionHandle
 //AgtMsdpPdu DeleteAllPdusInSession, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) DeleteAllPdus () error {
 //parameters: 
 //AgtMsdpPdu DeleteAllPdus, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) DecodePdu () error {
 //parameters: FormatMode PduId
 //AgtMsdpPdu DecodePdu, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) GetFieldCount ()(string, error) {
 //parameters: PduId FieldType
 //AgtMsdpPdu GetFieldCount
 return "", nil
}

func(np *MsdpPdu) GetSubfield ()(string, error) {
 //parameters: PduId FieldType Index SubfieldType
 //AgtMsdpPdu GetSubfield
 return "", nil
}

func(np *MsdpPdu) SetSubfield () error {
 //parameters: PduId FieldType Index SubfieldType SubfieldValue
 //AgtMsdpPdu SetSubfield, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) DeleteField () error {
 //parameters: PduId FieldType Index
 //AgtMsdpPdu DeleteField, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) SendPdu () error {
 //parameters: DestinationIpAddress PduId
 //AgtMsdpPdu SendPdu, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) GetNextPduInfo ()(string, error) {
 //parameters: TimeOutSec Count psaSessionHandle
 //AgtMsdpPdu GetNextPduInfo
 return "", nil
}

func(np *MsdpPdu) GetPduDetails ()(string, error) {
 //parameters: PduId PduInfoType
 //AgtMsdpPdu GetPduDetails
 return "", nil
}

func(np *MsdpPdu) Prohibit () error {
 //parameters: SessionHandle PduType
 //AgtMsdpPdu Prohibit, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) Permit () error {
 //parameters: SessionHandle PduType
 //AgtMsdpPdu Permit, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) IsProhibited () error {
 //parameters: SessionHandle PduType
 //AgtMsdpPdu IsProhibited, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) AddPduFilter () error {
 //parameters: SessionHandle PduType ModifierType SourceIpAddress GroupIpAddress
 //AgtMsdpPdu AddPduFilter, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) RemovePduFilter () error {
 //parameters: SessionHandle FilterId
 //AgtMsdpPdu RemovePduFilter, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) RemoveAllPduFilters () error {
 //parameters: SessionHandle
 //AgtMsdpPdu RemoveAllPduFilters, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) ListPduFilters ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpPdu ListPduFilters
 return "", nil
}

func(np *MsdpPdu) GetPduFilterDetails ()(string, error) {
 //parameters: SessionHandle FilterId
 //AgtMsdpPdu GetPduFilterDetails
 return "", nil
}

func(np *MsdpPdu) StartPduSave () error {
 //parameters: SessionHandle
 //AgtMsdpPdu StartPduSave, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) StopPduSave () error {
 //parameters: SessionHandle
 //AgtMsdpPdu StopPduSave, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) IsPduSaveEnabled () error {
 //parameters: SessionHandle
 //AgtMsdpPdu IsPduSaveEnabled, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) EnablePduNotification () error {
 //parameters: SessionHandle
 //AgtMsdpPdu EnablePduNotification, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) DisablePduNotification () error {
 //parameters: SessionHandle
 //AgtMsdpPdu DisablePduNotification, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) IsPduNotificationEnabled () error {
 //parameters: SessionHandle
 //AgtMsdpPdu IsPduNotificationEnabled, m.Object, m.Name);
 return nil
}

func(np *MsdpPdu) ListSavedPdus ()(string, error) {
 //parameters: SessionHandle
 //AgtMsdpPdu ListSavedPdus
 return "", nil
}

