package n2xsdk

type EndpointMesh struct {
 Handler string
}

func(np *EndpointMesh) GetType ()(string, error) {
 //parameters: Handle
 //AgtEndpointMesh GetType
 return "", nil
}

func(np *EndpointMesh) GetName ()(string, error) {
 //parameters: Handle
 //AgtEndpointMesh GetName
 return "", nil
}

func(np *EndpointMesh) GetLockCount ()(string, error) {
 //parameters: Handle
 //AgtEndpointMesh GetLockCount
 return "", nil
}

func(np *EndpointMesh) Enable () error {
 //parameters: MeshHandle
 //AgtEndpointMesh Enable, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) Disable () error {
 //parameters: MeshHandle
 //AgtEndpointMesh Disable, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) IsEnabled () error {
 //parameters: MeshHandle
 //AgtEndpointMesh IsEnabled, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) SetTrafficOrientation () error {
 //parameters: MeshHandle Orientation
 //AgtEndpointMesh SetTrafficOrientation, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) GetTrafficOrientation ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh GetTrafficOrientation
 return "", nil
}

func(np *EndpointMesh) SetLengthMode () error {
 //parameters: MeshHandle LengthMode
 //AgtEndpointMesh SetLengthMode, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) GetLengthMode ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh GetLengthMode
 return "", nil
}

func(np *EndpointMesh) SetLength () error {
 //parameters: MeshHandle LengthType Count psaLengthParameterList
 //AgtEndpointMesh SetLength, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) GetLength ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh GetLength
 return "", nil
}

func(np *EndpointMesh) ListSources ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh ListSources
 return "", nil
}

func(np *EndpointMesh) ListDestinations ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh ListDestinations
 return "", nil
}

func(np *EndpointMesh) ListStreamGroups ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh ListStreamGroups
 return "", nil
}

func(np *EndpointMesh) ListStreamGroupsBySource ()(string, error) {
 //parameters: MeshHandle Source
 //AgtEndpointMesh ListStreamGroupsBySource
 return "", nil
}

func(np *EndpointMesh) ListStreamGroupProfiles ()(string, error) {
 //parameters: MeshHandle StreamGroupCount psaStreamGroups
 //AgtEndpointMesh ListStreamGroupProfiles
 return "", nil
}

func(np *EndpointMesh) EnableStreamGeneration () error {
 //parameters: MeshHandle
 //AgtEndpointMesh EnableStreamGeneration, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) DisableStreamGeneration () error {
 //parameters: MeshHandle
 //AgtEndpointMesh DisableStreamGeneration, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) IsStreamGenerationEnabled () error {
 //parameters: MeshHandle
 //AgtEndpointMesh IsStreamGenerationEnabled, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) SetAverageLoad () error {
 //parameters: MeshHandle AverageLoad Units
 //AgtEndpointMesh SetAverageLoad, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) GetAverageLoad ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh GetAverageLoad
 return "", nil
}

func(np *EndpointMesh) IsMeshAddedOnNewProfiles () error {
 //parameters: MeshHandle
 //AgtEndpointMesh IsMeshAddedOnNewProfiles, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) AddSources () error {
 //parameters: MeshHandle NumSources psaSources
 //AgtEndpointMesh AddSources, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) RemoveSources () error {
 //parameters: MeshHandle NumSources psaSources
 //AgtEndpointMesh RemoveSources, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) AddDestinations () error {
 //parameters: MeshHandle NumDestinations psaDestinations
 //AgtEndpointMesh AddDestinations, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) RemoveDestinations () error {
 //parameters: MeshHandle NumDestinations psaDestinations
 //AgtEndpointMesh RemoveDestinations, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) SetEndpointTypes () error {
 //parameters: MeshHandle SourceType DestinationType
 //AgtEndpointMesh SetEndpointTypes, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) GetEndpointTypes ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh GetEndpointTypes
 return "", nil
}

func(np *EndpointMesh) SetTrafficDistribution () error {
 //parameters: MeshHandle Distribution
 //AgtEndpointMesh SetTrafficDistribution, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) GetTrafficDistribution ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh GetTrafficDistribution
 return "", nil
}

func(np *EndpointMesh) SetDeviceUnderTestType () error {
 //parameters: MeshHandle DutType
 //AgtEndpointMesh SetDeviceUnderTestType, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) GetDeviceUnderTestType ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh GetDeviceUnderTestType
 return "", nil
}

func(np *EndpointMesh) EnableFullMeshTransmitWithinPool () error {
 //parameters: MeshHandle
 //AgtEndpointMesh EnableFullMeshTransmitWithinPool, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) DisableFullMeshTransmitWithinPool () error {
 //parameters: MeshHandle
 //AgtEndpointMesh DisableFullMeshTransmitWithinPool, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) IsFullMeshTransmitWithinPoolEnabled () error {
 //parameters: MeshHandle
 //AgtEndpointMesh IsFullMeshTransmitWithinPoolEnabled, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) ListStreamGroupsBySourceDestinationPair ()(string, error) {
 //parameters: MeshHandle Source Destination
 //AgtEndpointMesh ListStreamGroupsBySourceDestinationPair
 return "", nil
}

func(np *EndpointMesh) GetEndpointsForStreamGroup ()(string, error) {
 //parameters: MeshHandle StreamGroup
 //AgtEndpointMesh GetEndpointsForStreamGroup
 return "", nil
}

func(np *EndpointMesh) Refresh () error {
 //parameters: MeshHandle
 //AgtEndpointMesh Refresh, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) SetIpMode () error {
 //parameters: MeshHandle IpMode
 //AgtEndpointMesh SetIpMode, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) GetIpMode ()(string, error) {
 //parameters: MeshHandle
 //AgtEndpointMesh GetIpMode
 return "", nil
}

func(np *EndpointMesh) EnableEmulationEncapsulation () error {
 //parameters: MeshHandle
 //AgtEndpointMesh EnableEmulationEncapsulation, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) DisableEmulationEncapsulation () error {
 //parameters: MeshHandle
 //AgtEndpointMesh DisableEmulationEncapsulation, m.Object, m.Name);
 return nil
}

func(np *EndpointMesh) IsEmulationEncapsulationEnabled () error {
 //parameters: MeshHandle
 //AgtEndpointMesh IsEmulationEncapsulationEnabled, m.Object, m.Name);
 return nil
}

