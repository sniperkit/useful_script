package n2xsdk

type DeviceMesh struct {
 Handler string
}

func(np *DeviceMesh) GetType ()(string, error) {
 //parameters: Handle
 //AgtDeviceMesh GetType
 return "", nil
}

func(np *DeviceMesh) GetName ()(string, error) {
 //parameters: Handle
 //AgtDeviceMesh GetName
 return "", nil
}

func(np *DeviceMesh) GetLockCount ()(string, error) {
 //parameters: Handle
 //AgtDeviceMesh GetLockCount
 return "", nil
}

func(np *DeviceMesh) Enable () error {
 //parameters: MeshHandle
 //AgtDeviceMesh Enable, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) Disable () error {
 //parameters: MeshHandle
 //AgtDeviceMesh Disable, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) IsEnabled () error {
 //parameters: MeshHandle
 //AgtDeviceMesh IsEnabled, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) SetTrafficOrientation () error {
 //parameters: MeshHandle Orientation
 //AgtDeviceMesh SetTrafficOrientation, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetTrafficOrientation ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh GetTrafficOrientation
 return "", nil
}

func(np *DeviceMesh) SetLengthMode () error {
 //parameters: MeshHandle LengthMode
 //AgtDeviceMesh SetLengthMode, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetLengthMode ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh GetLengthMode
 return "", nil
}

func(np *DeviceMesh) SetLength () error {
 //parameters: MeshHandle LengthType Count psaLengthParameterList
 //AgtDeviceMesh SetLength, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetLength ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh GetLength
 return "", nil
}

func(np *DeviceMesh) ListSources ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh ListSources
 return "", nil
}

func(np *DeviceMesh) ListDestinations ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh ListDestinations
 return "", nil
}

func(np *DeviceMesh) ListStreamGroups ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh ListStreamGroups
 return "", nil
}

func(np *DeviceMesh) ListStreamGroupsBySource ()(string, error) {
 //parameters: MeshHandle Source
 //AgtDeviceMesh ListStreamGroupsBySource
 return "", nil
}

func(np *DeviceMesh) ListStreamGroupProfiles ()(string, error) {
 //parameters: MeshHandle StreamGroupCount psaStreamGroups
 //AgtDeviceMesh ListStreamGroupProfiles
 return "", nil
}

func(np *DeviceMesh) EnableStreamGeneration () error {
 //parameters: MeshHandle
 //AgtDeviceMesh EnableStreamGeneration, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) DisableStreamGeneration () error {
 //parameters: MeshHandle
 //AgtDeviceMesh DisableStreamGeneration, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) IsStreamGenerationEnabled () error {
 //parameters: MeshHandle
 //AgtDeviceMesh IsStreamGenerationEnabled, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) SetAverageLoad () error {
 //parameters: MeshHandle AverageLoad Units
 //AgtDeviceMesh SetAverageLoad, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetAverageLoad ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh GetAverageLoad
 return "", nil
}

func(np *DeviceMesh) IsMeshAddedOnNewProfiles () error {
 //parameters: MeshHandle
 //AgtDeviceMesh IsMeshAddedOnNewProfiles, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) SetEndpointTypes () error {
 //parameters: MeshHandle SourceType DestinationType
 //AgtDeviceMesh SetEndpointTypes, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetEndpointTypes ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh GetEndpointTypes
 return "", nil
}

func(np *DeviceMesh) EnableManualSelectionOfAddressTypes () error {
 //parameters: MeshHandle
 //AgtDeviceMesh EnableManualSelectionOfAddressTypes, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) DisableManualSelectionOfAddressTypes () error {
 //parameters: MeshHandle
 //AgtDeviceMesh DisableManualSelectionOfAddressTypes, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) IsManualSelectionOfAddressTypesEnabled () error {
 //parameters: MeshHandle
 //AgtDeviceMesh IsManualSelectionOfAddressTypesEnabled, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) SetEndpointAddressTypes () error {
 //parameters: MeshHandle SourceAddressType DestinationAddressType
 //AgtDeviceMesh SetEndpointAddressTypes, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetEndpointAddressTypes ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh GetEndpointAddressTypes
 return "", nil
}

func(np *DeviceMesh) ListSupportedEndpointsByAddressType ()(string, error) {
 //parameters: EndpointType EndpointAddressType
 //AgtDeviceMesh ListSupportedEndpointsByAddressType
 return "", nil
}

func(np *DeviceMesh) ListSupportedEndpointAddressTypes ()(string, error) {
 //parameters: EndpointType Endpoint
 //AgtDeviceMesh ListSupportedEndpointAddressTypes
 return "", nil
}

func(np *DeviceMesh) SetTrafficDistribution () error {
 //parameters: MeshHandle Distribution
 //AgtDeviceMesh SetTrafficDistribution, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetTrafficDistribution ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh GetTrafficDistribution
 return "", nil
}

func(np *DeviceMesh) SetDeviceUnderTestType () error {
 //parameters: MeshHandle DutType
 //AgtDeviceMesh SetDeviceUnderTestType, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetDeviceUnderTestType ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh GetDeviceUnderTestType
 return "", nil
}

func(np *DeviceMesh) AddSources () error {
 //parameters: MeshHandle Count psaSources
 //AgtDeviceMesh AddSources, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) RemoveSources () error {
 //parameters: MeshHandle Count psaSources
 //AgtDeviceMesh RemoveSources, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) AddDestinations () error {
 //parameters: MeshHandle Count psaDestinations
 //AgtDeviceMesh AddDestinations, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) RemoveDestinations () error {
 //parameters: MeshHandle Count psaDestinations
 //AgtDeviceMesh RemoveDestinations, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) RemoveDestinationsForSources () error {
 //parameters: MeshHandle NumSources psaSources NumDestinations psaDestinations
 //AgtDeviceMesh RemoveDestinationsForSources, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) SetRemoteHostIpv6AddressSuffix () error {
 //parameters: MeshHandle Endpoint HostIpv6AddressSuffix
 //AgtDeviceMesh SetRemoteHostIpv6AddressSuffix, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetRemoteHostIpv6AddressSuffix ()(string, error) {
 //parameters: MeshHandle Endpoint
 //AgtDeviceMesh GetRemoteHostIpv6AddressSuffix
 return "", nil
}

func(np *DeviceMesh) SetRemoteHostCount () error {
 //parameters: MeshHandle Endpoint HostCountPerClient
 //AgtDeviceMesh SetRemoteHostCount, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetRemoteHostCount ()(string, error) {
 //parameters: MeshHandle Endpoint
 //AgtDeviceMesh GetRemoteHostCount
 return "", nil
}

func(np *DeviceMesh) Clear () error {
 //parameters: MeshHandle
 //AgtDeviceMesh Clear, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetForwardStreamGroup ()(string, error) {
 //parameters: MeshHandle Source Destination
 //AgtDeviceMesh GetForwardStreamGroup
 return "", nil
}

func(np *DeviceMesh) GetReverseStreamGroup ()(string, error) {
 //parameters: MeshHandle Source Destination
 //AgtDeviceMesh GetReverseStreamGroup
 return "", nil
}

func(np *DeviceMesh) GetEndpointsForStreamGroup ()(string, error) {
 //parameters: MeshHandle StreamGroup
 //AgtDeviceMesh GetEndpointsForStreamGroup
 return "", nil
}

func(np *DeviceMesh) ListSourcesForDestination ()(string, error) {
 //parameters: MeshHandle Destination
 //AgtDeviceMesh ListSourcesForDestination
 return "", nil
}

func(np *DeviceMesh) ListDestinationsForSource ()(string, error) {
 //parameters: MeshHandle Source
 //AgtDeviceMesh ListDestinationsForSource
 return "", nil
}

func(np *DeviceMesh) UpdateStreamGroups () error {
 //parameters: MeshHandle
 //AgtDeviceMesh UpdateStreamGroups, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) SetStreamGenerationParameter () error {
 //parameters: MeshHandle Protocol ProtocolInstance Field
 //AgtDeviceMesh SetStreamGenerationParameter, m.Object, m.Name);
 return nil
}

func(np *DeviceMesh) GetStreamGenerationParameter ()(string, error) {
 //parameters: MeshHandle
 //AgtDeviceMesh GetStreamGenerationParameter
 return "", nil
}

func(np *DeviceMesh) ResetPduHeaders () error {
 //parameters: MeshHandle
 //AgtDeviceMesh ResetPduHeaders, m.Object, m.Name);
 return nil
}

