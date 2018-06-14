package n2xsdk

type PduHeader struct {
 Handler string
}

func(np *PduHeader) IsPduResolved () error {
 //parameters: PduHandle
 //AgtPduHeader IsPduResolved, m.Object, m.Name);
 return nil
}

func(np *PduHeader) ListProtocolsInHeader ()(string, error) {
 //parameters: PduHandle
 //AgtPduHeader ListProtocolsInHeader
 return "", nil
}

func(np *PduHeader) GetProtocolBitOffset ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance
 //AgtPduHeader GetProtocolBitOffset
 return "", nil
}

func(np *PduHeader) ListAvailableProtocolFields ()(string, error) {
 //parameters: Protocol
 //AgtPduHeader ListAvailableProtocolFields
 return "", nil
}

func(np *PduHeader) ListProtocolFieldsInHeader ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance
 //AgtPduHeader ListProtocolFieldsInHeader
 return "", nil
}

func(np *PduHeader) GetFieldBitOffset ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldBitOffset
 return "", nil
}

func(np *PduHeader) GetFieldInstanceBitOffset ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance
 //AgtPduHeader GetFieldInstanceBitOffset
 return "", nil
}

func(np *PduHeader) ListOptionalFields ()(string, error) {
 //parameters: PduHandle Protocol
 //AgtPduHeader ListOptionalFields
 return "", nil
}

func(np *PduHeader) EnableOptionalField () error {
 //parameters: PduHandle Protocol ProtocolInstance OptionalField
 //AgtPduHeader EnableOptionalField, m.Object, m.Name);
 return nil
}

func(np *PduHeader) DisableOptionalField () error {
 //parameters: PduHandle Protocol ProtocolInstance OptionalField
 //AgtPduHeader DisableOptionalField, m.Object, m.Name);
 return nil
}

func(np *PduHeader) IsOptionalFieldEnabled () error {
 //parameters: PduHandle Protocol ProtocolInstance OptionalField
 //AgtPduHeader IsOptionalFieldEnabled, m.Object, m.Name);
 return nil
}

func(np *PduHeader) EnableOptionalFieldInstance () error {
 //parameters: PduHandle Protocol ProtocolInstance OptionalField FieldInstance
 //AgtPduHeader EnableOptionalFieldInstance, m.Object, m.Name);
 return nil
}

func(np *PduHeader) DisableOptionalFieldInstance () error {
 //parameters: PduHandle Protocol ProtocolInstance OptionalField FieldInstance
 //AgtPduHeader DisableOptionalFieldInstance, m.Object, m.Name);
 return nil
}

func(np *PduHeader) IsOptionalFieldInstanceEnabled () error {
 //parameters: PduHandle Protocol ProtocolInstance OptionalField FieldInstance
 //AgtPduHeader IsOptionalFieldInstanceEnabled, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldLength ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldLength
 return "", nil
}

func(np *PduHeader) SetFieldLength () error {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldLength
 //AgtPduHeader SetFieldLength, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldInstanceLength ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance
 //AgtPduHeader GetFieldInstanceLength
 return "", nil
}

func(np *PduHeader) SetFieldInstanceLength () error {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance FieldLength
 //AgtPduHeader SetFieldInstanceLength, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldValueType ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldValueType
 return "", nil
}

func(np *PduHeader) GetFieldFixedValue ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldFixedValue
 return "", nil
}

func(np *PduHeader) SetFieldFixedValue () error {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldValue
 //AgtPduHeader SetFieldFixedValue, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldIncrementingValueRange ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldIncrementingValueRange
 return "", nil
}

func(np *PduHeader) SetFieldIncrementingValueRange () error {
 //parameters: PduHandle Protocol ProtocolInstance Field InFieldOffset StartValue NumValues StepSize
 //AgtPduHeader SetFieldIncrementingValueRange, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldDecrementingValueRange ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldDecrementingValueRange
 return "", nil
}

func(np *PduHeader) SetFieldDecrementingValueRange () error {
 //parameters: PduHandle Protocol ProtocolInstance Field InFieldOffset StartValue NumValues StepSize
 //AgtPduHeader SetFieldDecrementingValueRange, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldRandomValueRange ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldRandomValueRange
 return "", nil
}

func(np *PduHeader) SetFieldRandomValueRange () error {
 //parameters: PduHandle Protocol ProtocolInstance Field InFieldOffset MinValue MaxValue
 //AgtPduHeader SetFieldRandomValueRange, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldValueList ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldValueList
 return "", nil
}

func(np *PduHeader) SetFieldValueList ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field Count psaValueList
 //AgtPduHeader SetFieldValueList
 return "", nil
}

func(np *PduHeader) GetFieldValueRangeRepeatCount ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldValueRangeRepeatCount
 return "", nil
}

func(np *PduHeader) SetFieldValueRangeRepeatCount () error {
 //parameters: PduHandle Protocol ProtocolInstance Field ValueRangeRepeatCount
 //AgtPduHeader SetFieldValueRangeRepeatCount, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldInstanceFixedValue ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance
 //AgtPduHeader GetFieldInstanceFixedValue
 return "", nil
}

func(np *PduHeader) SetFieldInstanceFixedValue () error {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance FieldValue
 //AgtPduHeader SetFieldInstanceFixedValue, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldInstanceIncrementingValueRange ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance
 //AgtPduHeader GetFieldInstanceIncrementingValueRange
 return "", nil
}

func(np *PduHeader) SetFieldInstanceIncrementingValueRange () error {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance InFieldOffset StartValue NumValues StepSize
 //AgtPduHeader SetFieldInstanceIncrementingValueRange, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldInstanceDecrementingValueRange ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance
 //AgtPduHeader GetFieldInstanceDecrementingValueRange
 return "", nil
}

func(np *PduHeader) SetFieldInstanceDecrementingValueRange () error {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance InFieldOffset StartValue NumValues StepSize
 //AgtPduHeader SetFieldInstanceDecrementingValueRange, m.Object, m.Name);
 return nil
}

func(np *PduHeader) IsAutoGenerationEnabled () error {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader IsAutoGenerationEnabled, m.Object, m.Name);
 return nil
}

func(np *PduHeader) EnableAutoGeneration () error {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader EnableAutoGeneration, m.Object, m.Name);
 return nil
}

func(np *PduHeader) DisableAutoGeneration () error {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader DisableAutoGeneration, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetRepeatInstanceCount ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetRepeatInstanceCount
 return "", nil
}

func(np *PduHeader) SetRepeatInstanceCount () error {
 //parameters: PduHandle Protocol ProtocolInstance Field RepeatCount
 //AgtPduHeader SetRepeatInstanceCount, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetRepeatInstanceCountForFieldInstance ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance
 //AgtPduHeader GetRepeatInstanceCountForFieldInstance
 return "", nil
}

func(np *PduHeader) SetRepeatInstanceCountForFieldInstance () error {
 //parameters: PduHandle Protocol ProtocolInstance Field FieldInstance RepeatCount
 //AgtPduHeader SetRepeatInstanceCountForFieldInstance, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetRepeatInstanceValues ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field StartInstance EndInstance
 //AgtPduHeader GetRepeatInstanceValues
 return "", nil
}

func(np *PduHeader) SetRepeatInstanceValueRange () error {
 //parameters: PduHandle Protocol ProtocolInstance Field StartInstance EndInstance InFieldOffset StartValue StepSize
 //AgtPduHeader SetRepeatInstanceValueRange, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetSegmentValueType ()(string, error) {
 //parameters: PduHandle Offset
 //AgtPduHeader GetSegmentValueType
 return "", nil
}

func(np *PduHeader) GetSegmentFixedValue ()(string, error) {
 //parameters: PduHandle Offset
 //AgtPduHeader GetSegmentFixedValue
 return "", nil
}

func(np *PduHeader) SetSegmentFixedValue () error {
 //parameters: PduHandle Offset SegmentLength SegmentValue
 //AgtPduHeader SetSegmentFixedValue, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetSegmentIncrementingValueRange ()(string, error) {
 //parameters: PduHandle Offset
 //AgtPduHeader GetSegmentIncrementingValueRange
 return "", nil
}

func(np *PduHeader) SetSegmentIncrementingValueRange () error {
 //parameters: PduHandle Offset SegmentLength SegmentStartValue NumValues StepSize
 //AgtPduHeader SetSegmentIncrementingValueRange, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetSegmentDecrementingValueRange ()(string, error) {
 //parameters: PduHandle Offset
 //AgtPduHeader GetSegmentDecrementingValueRange
 return "", nil
}

func(np *PduHeader) SetSegmentDecrementingValueRange () error {
 //parameters: PduHandle Offset SegmentLength SegmentStartValue NumValues StepSize
 //AgtPduHeader SetSegmentDecrementingValueRange, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetSegmentRandomValueRange ()(string, error) {
 //parameters: PduHandle Offset
 //AgtPduHeader GetSegmentRandomValueRange
 return "", nil
}

func(np *PduHeader) SetSegmentRandomValueRange () error {
 //parameters: PduHandle Offset SegmentLength SegmentMinValue SegmentMaxValue
 //AgtPduHeader SetSegmentRandomValueRange, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetSegmentValueList ()(string, error) {
 //parameters: PduHandle Offset
 //AgtPduHeader GetSegmentValueList
 return "", nil
}

func(np *PduHeader) SetSegmentValueList ()(string, error) {
 //parameters: PduHandle Offset SegmentLength Count psaSegmentValueList
 //AgtPduHeader SetSegmentValueList
 return "", nil
}

func(np *PduHeader) GetSegmentValueRangeRepeatCount ()(string, error) {
 //parameters: PduHandle Offset
 //AgtPduHeader GetSegmentValueRangeRepeatCount
 return "", nil
}

func(np *PduHeader) SetSegmentValueRangeRepeatCount () error {
 //parameters: PduHandle Offset ValueRangeRepeatCount
 //AgtPduHeader SetSegmentValueRangeRepeatCount, m.Object, m.Name);
 return nil
}

func(np *PduHeader) GetFieldFullName ()(string, error) {
 //parameters: PduHandle Protocol ProtocolInstance Field
 //AgtPduHeader GetFieldFullName
 return "", nil
}

