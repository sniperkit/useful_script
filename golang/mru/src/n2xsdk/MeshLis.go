package n2xsdk

type MeshLis struct {
 Handler string
}

func(np *MeshLis) Add () error {
 //parameters: Type
 //AgtMeshList Add, m.Object, m.Name);
 return nil
}

func(np *MeshLis) AddItems () error {
 //parameters: Type NumberOfItems
 //AgtMeshList AddItems, m.Object, m.Name);
 return nil
}

func(np *MeshLis) Remove () error {
 //parameters: Handle
 //AgtMeshList Remove, m.Object, m.Name);
 return nil
}

func(np *MeshLis) Copy () error {
 //parameters: Handle
 //AgtMeshList Copy, m.Object, m.Name);
 return nil
}

func(np *MeshLis) SetName () error {
 //parameters: Handle Name
 //AgtMeshList SetName, m.Object, m.Name);
 return nil
}

func(np *MeshLis) GetName ()(string, error) {
 //parameters: Handle
 //AgtMeshList GetName
 return "", nil
}

func(np *MeshLis) GetHandle ()(string, error) {
 //parameters: Name
 //AgtMeshList GetHandle
 return "", nil
}

func(np *MeshLis) GetType ()(string, error) {
 //parameters: Handle
 //AgtMeshList GetType
 return "", nil
}

func(np *MeshLis) LockItem () error {
 //parameters: Handle
 //AgtMeshList LockItem, m.Object, m.Name);
 return nil
}

func(np *MeshLis) UnlockItem () error {
 //parameters: Handle
 //AgtMeshList UnlockItem, m.Object, m.Name);
 return nil
}

func(np *MeshLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtMeshList ListHandles
 return "", nil
}

func(np *MeshLis) ListNames ()(string, error) {
 //parameters: 
 //AgtMeshList ListNames
 return "", nil
}

func(np *MeshLis) ListHandlesByType ()(string, error) {
 //parameters: Type
 //AgtMeshList ListHandlesByType
 return "", nil
}

func(np *MeshLis) ListTypes ()(string, error) {
 //parameters: 
 //AgtMeshList ListTypes
 return "", nil
}

func(np *MeshLis) GetInterface ()(string, error) {
 //parameters: Type
 //AgtMeshList GetInterface
 return "", nil
}

func(np *MeshLis) GetLockCount ()(string, error) {
 //parameters: 
 //AgtMeshList GetLockCount
 return "", nil
}

func(np *MeshLis) AddMeshWithNewProfiles () error {
 //parameters: Type
 //AgtMeshList AddMeshWithNewProfiles, m.Object, m.Name);
 return nil
}

