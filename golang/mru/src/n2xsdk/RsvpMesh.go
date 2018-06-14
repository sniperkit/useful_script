package n2xsdk

type RsvpMesh struct {
 Handler string
}

func(np *RsvpMesh) GetType ()(string, error) {
 //parameters: Handle
 //AgtRsvpMesh GetType
 return "", nil
}

func(np *RsvpMesh) GetName ()(string, error) {
 //parameters: Handle
 //AgtRsvpMesh GetName
 return "", nil
}

func(np *RsvpMesh) GetLockCount ()(string, error) {
 //parameters: Handle
 //AgtRsvpMesh GetLockCount
 return "", nil
}

func(np *RsvpMesh) Enable () error {
 //parameters: MeshHandle
 //AgtRsvpMesh Enable, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) Disable () error {
 //parameters: MeshHandle
 //AgtRsvpMesh Disable, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) IsEnabled () error {
 //parameters: MeshHandle
 //AgtRsvpMesh IsEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) SetTrafficOrientation () error {
 //parameters: MeshHandle Orientation
 //AgtRsvpMesh SetTrafficOrientation, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) GetTrafficOrientation ()(string, error) {
 //parameters: MeshHandle
 //AgtRsvpMesh GetTrafficOrientation
 return "", nil
}

func(np *RsvpMesh) SetLengthMode () error {
 //parameters: MeshHandle LengthMode
 //AgtRsvpMesh SetLengthMode, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) GetLengthMode ()(string, error) {
 //parameters: MeshHandle
 //AgtRsvpMesh GetLengthMode
 return "", nil
}

func(np *RsvpMesh) SetLength () error {
 //parameters: MeshHandle LengthType Count psaLengthParameterList
 //AgtRsvpMesh SetLength, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) GetLength ()(string, error) {
 //parameters: MeshHandle
 //AgtRsvpMesh GetLength
 return "", nil
}

func(np *RsvpMesh) ListSources ()(string, error) {
 //parameters: MeshHandle
 //AgtRsvpMesh ListSources
 return "", nil
}

func(np *RsvpMesh) ListDestinations ()(string, error) {
 //parameters: MeshHandle
 //AgtRsvpMesh ListDestinations
 return "", nil
}

func(np *RsvpMesh) ListStreamGroups ()(string, error) {
 //parameters: MeshHandle
 //AgtRsvpMesh ListStreamGroups
 return "", nil
}

func(np *RsvpMesh) ListStreamGroupsBySource ()(string, error) {
 //parameters: MeshHandle Source
 //AgtRsvpMesh ListStreamGroupsBySource
 return "", nil
}

func(np *RsvpMesh) ListStreamGroupProfiles ()(string, error) {
 //parameters: MeshHandle StreamGroupCount psaStreamGroups
 //AgtRsvpMesh ListStreamGroupProfiles
 return "", nil
}

func(np *RsvpMesh) EnableStreamGeneration () error {
 //parameters: MeshHandle
 //AgtRsvpMesh EnableStreamGeneration, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) DisableStreamGeneration () error {
 //parameters: MeshHandle
 //AgtRsvpMesh DisableStreamGeneration, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) IsStreamGenerationEnabled () error {
 //parameters: MeshHandle
 //AgtRsvpMesh IsStreamGenerationEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) SetAverageLoad () error {
 //parameters: MeshHandle AverageLoad Units
 //AgtRsvpMesh SetAverageLoad, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) GetAverageLoad ()(string, error) {
 //parameters: MeshHandle
 //AgtRsvpMesh GetAverageLoad
 return "", nil
}

func(np *RsvpMesh) IsMeshAddedOnNewProfiles () error {
 //parameters: MeshHandle
 //AgtRsvpMesh IsMeshAddedOnNewProfiles, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) AddIngressLsps () error {
 //parameters: MeshHandle PoolHandleCount PoolHandles PoolIndexCount PoolIndices PortCount DestinationPorts
 //AgtRsvpMesh AddIngressLsps, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) AddIngressLspPools () error {
 //parameters: MeshHandle PoolHandleCount PoolHandles PortCount DestinationPorts
 //AgtRsvpMesh AddIngressLspPools, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) RemoveIngressLsps () error {
 //parameters: MeshHandle PoolHandleCount PoolHandles PoolIndexCount PoolIndices
 //AgtRsvpMesh RemoveIngressLsps, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) RemoveIngressLspPools () error {
 //parameters: MeshHandle Count PoolHandles
 //AgtRsvpMesh RemoveIngressLspPools, m.Object, m.Name);
 return nil
}

func(np *RsvpMesh) ListIngressLsps ()(string, error) {
 //parameters: MeshHandle
 //AgtRsvpMesh ListIngressLsps
 return "", nil
}

func(np *RsvpMesh) ListIngressLspPools ()(string, error) {
 //parameters: MeshHandle
 //AgtRsvpMesh ListIngressLspPools
 return "", nil
}

func(np *RsvpMesh) GetIngressLspStreamGroup ()(string, error) {
 //parameters: MeshHandle PoolHandle PoolIndex
 //AgtRsvpMesh GetIngressLspStreamGroup
 return "", nil
}

func(np *RsvpMesh) GetIngressLspPoolStreamGroup ()(string, error) {
 //parameters: MeshHandle PoolHandle
 //AgtRsvpMesh GetIngressLspPoolStreamGroup
 return "", nil
}

func(np *RsvpMesh) GetIngressLspForStreamGroup ()(string, error) {
 //parameters: MeshHandle StreamGroup
 //AgtRsvpMesh GetIngressLspForStreamGroup
 return "", nil
}

func(np *RsvpMesh) GetIngressLspPoolForStreamGroup ()(string, error) {
 //parameters: MeshHandle StreamGroup
 //AgtRsvpMesh GetIngressLspPoolForStreamGroup
 return "", nil
}

