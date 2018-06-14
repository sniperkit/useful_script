package n2xsdk

type ConnectionGroupList struct {
 Handler string
}

func(np *ConnectionGroupLis) Add () error {
 //parameters: Type
 //AgtConnectionGroupList Add, m.Object, m.Name);
 return nil
}

func(np *ConnectionGroupLis) AddItems () error {
 //parameters: Type NumberOfItems
 //AgtConnectionGroupList AddItems, m.Object, m.Name);
 return nil
}

func(np *ConnectionGroupLis) Remove () error {
 //parameters: Handle
 //AgtConnectionGroupList Remove, m.Object, m.Name);
 return nil
}

func(np *ConnectionGroupLis) Copy () error {
 //parameters: Handle
 //AgtConnectionGroupList Copy, m.Object, m.Name);
 return nil
}

func(np *ConnectionGroupLis) SetName () error {
 //parameters: Handle Name
 //AgtConnectionGroupList SetName, m.Object, m.Name);
 return nil
}

func(np *ConnectionGroupLis) GetName ()(string, error) {
 //parameters: Handle
 //AgtConnectionGroupList GetName
 return "", nil
}

func(np *ConnectionGroupLis) GetHandle ()(string, error) {
 //parameters: Name
 //AgtConnectionGroupList GetHandle
 return "", nil
}

func(np *ConnectionGroupLis) GetType ()(string, error) {
 //parameters: Handle
 //AgtConnectionGroupList GetType
 return "", nil
}

func(np *ConnectionGroupLis) LockItem () error {
 //parameters: Handle
 //AgtConnectionGroupList LockItem, m.Object, m.Name);
 return nil
}

func(np *ConnectionGroupLis) UnlockItem () error {
 //parameters: Handle
 //AgtConnectionGroupList UnlockItem, m.Object, m.Name);
 return nil
}

func(np *ConnectionGroupLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtConnectionGroupList ListHandles
 return "", nil
}

func(np *ConnectionGroupLis) ListNames ()(string, error) {
 //parameters: 
 //AgtConnectionGroupList ListNames
 return "", nil
}

func(np *ConnectionGroupLis) ListHandlesByType ()(string, error) {
 //parameters: Type
 //AgtConnectionGroupList ListHandlesByType
 return "", nil
}

func(np *ConnectionGroupLis) ListTypes ()(string, error) {
 //parameters: 
 //AgtConnectionGroupList ListTypes
 return "", nil
}

func(np *ConnectionGroupLis) GetInterface ()(string, error) {
 //parameters: Type
 //AgtConnectionGroupList GetInterface
 return "", nil
}

func(np *ConnectionGroupLis) GetLockCount ()(string, error) {
 //parameters: 
 //AgtConnectionGroupList GetLockCount
 return "", nil
}

func(np *ConnectionGroupLis) ListSupportedApplicationTypes ()(string, error) {
 //parameters: 
 //AgtConnectionGroupList ListSupportedApplicationTypes
 return "", nil
}

func(np *ConnectionGroupLis) ListSupportedApplicationTypesOnPort ()(string, error) {
 //parameters: PortHandle
 //AgtConnectionGroupList ListSupportedApplicationTypesOnPort
 return "", nil
}

func(np *ConnectionGroupLis) AddConnectionGroups () error {
 //parameters: ApplicationType NumberOfClientPortHandles psaClientPortHandles NumberOfServerPortHandles psaServerPortHandles
 //AgtConnectionGroupList AddConnectionGroups, m.Object, m.Name);
 return nil
}

func(np *ConnectionGroupLis) Clear () error {
 //parameters: 
 //AgtConnectionGroupList Clear, m.Object, m.Name);
 return nil
}

func(np *ConnectionGroupLis) ListHandlesOnPort ()(string, error) {
 //parameters: PortHandle
 //AgtConnectionGroupList ListHandlesOnPort
 return "", nil
}

func(np *ConnectionGroupLis) ListHandlesByApplicationType ()(string, error) {
 //parameters: ApplicationType
 //AgtConnectionGroupList ListHandlesByApplicationType
 return "", nil
}

func(np *ConnectionGroupLis) ListHandlesOnPortByApplicationType ()(string, error) {
 //parameters: PortHandle ApplicationType
 //AgtConnectionGroupList ListHandlesOnPortByApplicationType
 return "", nil
}

func(np *ConnectionGroupLis) GetAvailableResourcesOnPort ()(string, error) {
 //parameters: PortHandle
 //AgtConnectionGroupList GetAvailableResourcesOnPort
 return "", nil
}

