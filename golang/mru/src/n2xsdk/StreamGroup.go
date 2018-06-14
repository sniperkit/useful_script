package n2xsdk

type StreamGroup struct {
 Handler string
}

func(np *StreamGroup) GetType ()(string, error) {
 //parameters: Handle
 //AgtStreamGroup GetType
 return "", nil
}

func(np *StreamGroup) GetName ()(string, error) {
 //parameters: Handle
 //AgtStreamGroup GetName
 return "", nil
}

func(np *StreamGroup) GetLockCount ()(string, error) {
 //parameters: Handle
 //AgtStreamGroup GetLockCount
 return "", nil
}

func(np *StreamGroup) GetSourcePort ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetSourcePort
 return "", nil
}

func(np *StreamGroup) SetSourceEndpointType () error {
 //parameters: StreamGroupHandle EndpointType
 //AgtStreamGroup SetSourceEndpointType, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetSourceEndpointType ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetSourceEndpointType
 return "", nil
}

func(np *StreamGroup) SetSourceEndpoint () error {
 //parameters: StreamGroupHandle Count psaEndpointIdentifier
 //AgtStreamGroup SetSourceEndpoint, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetSourceEndpoint ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetSourceEndpoint
 return "", nil
}

func(np *StreamGroup) Refresh () error {
 //parameters: StreamGroupHandle
 //AgtStreamGroup Refresh, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) SetExpectedDestinationPorts () error {
 //parameters: StreamGroupHandle Count psaDestinationPorts
 //AgtStreamGroup SetExpectedDestinationPorts, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetExpectedDestinationPorts ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetExpectedDestinationPorts
 return "", nil
}

func(np *StreamGroup) ListAllL2Protocols ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup ListAllL2Protocols
 return "", nil
}

func(np *StreamGroup) ListAllPacketTypes ()(string, error) {
 //parameters: 
 //AgtStreamGroup ListAllPacketTypes
 return "", nil
}

func(np *StreamGroup) SetPduHeaders () error {
 //parameters: StreamGroupHandle Count psaProtocolList
 //AgtStreamGroup SetPduHeaders, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) SetPduHeadersByPacketType () error {
 //parameters: StreamGroupHandle L2Protocol PacketType
 //AgtStreamGroup SetPduHeadersByPacketType, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) AppendHeader () error {
 //parameters: StreamGroupHandle Protocol
 //AgtStreamGroup AppendHeader, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetDefaultL2Protocol ()(string, error) {
 //parameters: PortHandle
 //AgtStreamGroup GetDefaultL2Protocol
 return "", nil
}

func(np *StreamGroup) SetLengthMode () error {
 //parameters: StreamGroupHandle LengthMode
 //AgtStreamGroup SetLengthMode, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetLengthMode ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetLengthMode
 return "", nil
}

func(np *StreamGroup) SetLength () error {
 //parameters: StreamGroupHandle LengthType Count psaLengthParameterList
 //AgtStreamGroup SetLength, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetLength ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetLength
 return "", nil
}

func(np *StreamGroup) GetStreamTag ()(string, error) {
 //parameters: StreamGroupHandle StreamIndex
 //AgtStreamGroup GetStreamTag
 return "", nil
}

func(np *StreamGroup) GetStreamId ()(string, error) {
 //parameters: StreamGroupHandle StreamIndex
 //AgtStreamGroup GetStreamId
 return "", nil
}

func(np *StreamGroup) SetRepeatCount () error {
 //parameters: StreamGroupHandle RepeatCount
 //AgtStreamGroup SetRepeatCount, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetRepeatCount ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetRepeatCount
 return "", nil
}

func(np *StreamGroup) EnableTestPayload () error {
 //parameters: StreamGroupHandle
 //AgtStreamGroup EnableTestPayload, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) DisableTestPayload () error {
 //parameters: StreamGroupHandle
 //AgtStreamGroup DisableTestPayload, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) IsTestPayloadEnabled () error {
 //parameters: StreamGroupHandle
 //AgtStreamGroup IsTestPayloadEnabled, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) SetProfile () error {
 //parameters: StreamGroupHandle ProfileHandle
 //AgtStreamGroup SetProfile, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetProfile ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetProfile
 return "", nil
}

func(np *StreamGroup) GetPdu ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetPdu
 return "", nil
}

func(np *StreamGroup) SetStreamGenerationParameter () error {
 //parameters: StreamGroupHandle Protocol ProtocolInstance Field
 //AgtStreamGroup SetStreamGenerationParameter, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetStreamGenerationParameter ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetStreamGenerationParameter
 return "", nil
}

func(np *StreamGroup) GetNumberOfStreams ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetNumberOfStreams
 return "", nil
}

func(np *StreamGroup) GetStreamGenerationFieldValue ()(string, error) {
 //parameters: StreamGroupHandle StreamIndex
 //AgtStreamGroup GetStreamGenerationFieldValue
 return "", nil
}

func(np *StreamGroup) SetFieldModifiersRelation () error {
 //parameters: StreamGroupHandle FieldModifiersRelation
 //AgtStreamGroup SetFieldModifiersRelation, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetFieldModifiersRelation ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetFieldModifiersRelation
 return "", nil
}

func(np *StreamGroup) SetLinkedFieldModifiers () error {
 //parameters: StreamGroupHandle NumProtocols psaProtocols NumProtocolInstances psaProtocolInstances NumFields psaFields
 //AgtStreamGroup SetLinkedFieldModifiers, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetLinkedFieldModifiers ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetLinkedFieldModifiers
 return "", nil
}

func(np *StreamGroup) Enable () error {
 //parameters: StreamGroupHandle
 //AgtStreamGroup Enable, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) Disable () error {
 //parameters: StreamGroupHandle
 //AgtStreamGroup Disable, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) IsEnabled () error {
 //parameters: StreamGroupHandle
 //AgtStreamGroup IsEnabled, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) SetL2Error () error {
 //parameters: StreamGroupHandle L2Error
 //AgtStreamGroup SetL2Error, m.Object, m.Name);
 return nil
}

func(np *StreamGroup) GetL2Error ()(string, error) {
 //parameters: StreamGroupHandle
 //AgtStreamGroup GetL2Error
 return "", nil
}

