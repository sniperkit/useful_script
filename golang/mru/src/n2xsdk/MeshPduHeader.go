package n2xsdk

type MeshPduHeader struct {
 Handler string
}

func(np *MeshPduHeader) SetPduHeaders () error {
 //parameters: MeshHandle Count psaProtocolList
 //AgtMeshPduHeader SetPduHeaders, m.Object, m.Name);
 return nil
}

func(np *MeshPduHeader) SetPduHeadersByPacketType () error {
 //parameters: MeshHandle L2Protocol PacketType
 //AgtMeshPduHeader SetPduHeadersByPacketType, m.Object, m.Name);
 return nil
}

func(np *MeshPduHeader) GetFieldValueType ()(string, error) {
 //parameters: MeshHandle Protocol ProtocolInstance Field
 //AgtMeshPduHeader GetFieldValueType
 return "", nil
}

func(np *MeshPduHeader) GetFieldFixedValue ()(string, error) {
 //parameters: MeshHandle Protocol ProtocolInstance Field
 //AgtMeshPduHeader GetFieldFixedValue
 return "", nil
}

func(np *MeshPduHeader) SetFieldFixedValue () error {
 //parameters: MeshHandle Protocol ProtocolInstance Field FieldValue
 //AgtMeshPduHeader SetFieldFixedValue, m.Object, m.Name);
 return nil
}

func(np *MeshPduHeader) GetFieldIncrementingValueRange ()(string, error) {
 //parameters: MeshHandle Protocol ProtocolInstance Field
 //AgtMeshPduHeader GetFieldIncrementingValueRange
 return "", nil
}

func(np *MeshPduHeader) SetFieldIncrementingValueRange () error {
 //parameters: MeshHandle Protocol ProtocolInstance Field InFieldOffset StartValue NumValues StepSize
 //AgtMeshPduHeader SetFieldIncrementingValueRange, m.Object, m.Name);
 return nil
}

func(np *MeshPduHeader) GetFieldDecrementingValueRange ()(string, error) {
 //parameters: MeshHandle Protocol ProtocolInstance Field
 //AgtMeshPduHeader GetFieldDecrementingValueRange
 return "", nil
}

func(np *MeshPduHeader) SetFieldDecrementingValueRange () error {
 //parameters: MeshHandle Protocol ProtocolInstance Field InFieldOffset StartValue NumValues StepSize
 //AgtMeshPduHeader SetFieldDecrementingValueRange, m.Object, m.Name);
 return nil
}

func(np *MeshPduHeader) GetFieldRandomValueRange ()(string, error) {
 //parameters: MeshHandle Protocol ProtocolInstance Field
 //AgtMeshPduHeader GetFieldRandomValueRange
 return "", nil
}

func(np *MeshPduHeader) SetFieldRandomValueRange () error {
 //parameters: MeshHandle Protocol ProtocolInstance Field InFieldOffset MinValue MaxValue
 //AgtMeshPduHeader SetFieldRandomValueRange, m.Object, m.Name);
 return nil
}

func(np *MeshPduHeader) GetFieldValueList ()(string, error) {
 //parameters: MeshHandle Protocol ProtocolInstance Field
 //AgtMeshPduHeader GetFieldValueList
 return "", nil
}

func(np *MeshPduHeader) SetFieldValueList ()(string, error) {
 //parameters: MeshHandle Protocol ProtocolInstance Field Count psaValueList
 //AgtMeshPduHeader SetFieldValueList
 return "", nil
}

func(np *MeshPduHeader) EnableOptionalField () error {
 //parameters: MeshHandle Protocol ProtocolInstance OptionalField
 //AgtMeshPduHeader EnableOptionalField, m.Object, m.Name);
 return nil
}

func(np *MeshPduHeader) DisableOptionalField () error {
 //parameters: MeshHandle Protocol ProtocolInstance OptionalField
 //AgtMeshPduHeader DisableOptionalField, m.Object, m.Name);
 return nil
}

func(np *MeshPduHeader) IsOptionalFieldEnabled () error {
 //parameters: MeshHandle Protocol ProtocolInstance OptionalField
 //AgtMeshPduHeader IsOptionalFieldEnabled, m.Object, m.Name);
 return nil
}

