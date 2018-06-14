package n2xsdk

type StreamGroupLis struct {
 Handler string
}

func(np *StreamGroupLis) Add () error {
 //parameters: Type
 //AgtStreamGroupList Add, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) AddItems () error {
 //parameters: Type NumberOfItems
 //AgtStreamGroupList AddItems, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) Remove () error {
 //parameters: Handle
 //AgtStreamGroupList Remove, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) Copy () error {
 //parameters: Handle
 //AgtStreamGroupList Copy, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) SetName () error {
 //parameters: Handle Name
 //AgtStreamGroupList SetName, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) GetName ()(string, error) {
 //parameters: Handle
 //AgtStreamGroupList GetName
 return "", nil
}

func(np *StreamGroupLis) GetHandle ()(string, error) {
 //parameters: Name
 //AgtStreamGroupList GetHandle
 return "", nil
}

func(np *StreamGroupLis) GetType ()(string, error) {
 //parameters: Handle
 //AgtStreamGroupList GetType
 return "", nil
}

func(np *StreamGroupLis) LockItem () error {
 //parameters: Handle
 //AgtStreamGroupList LockItem, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) UnlockItem () error {
 //parameters: Handle
 //AgtStreamGroupList UnlockItem, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtStreamGroupList ListHandles
 return "", nil
}

func(np *StreamGroupLis) ListNames ()(string, error) {
 //parameters: 
 //AgtStreamGroupList ListNames
 return "", nil
}

func(np *StreamGroupLis) ListHandlesByType ()(string, error) {
 //parameters: Type
 //AgtStreamGroupList ListHandlesByType
 return "", nil
}

func(np *StreamGroupLis) ListTypes ()(string, error) {
 //parameters: 
 //AgtStreamGroupList ListTypes
 return "", nil
}

func(np *StreamGroupLis) GetInterface ()(string, error) {
 //parameters: Type
 //AgtStreamGroupList GetInterface
 return "", nil
}

func(np *StreamGroupLis) GetLockCount ()(string, error) {
 //parameters: 
 //AgtStreamGroupList GetLockCount
 return "", nil
}

func(np *StreamGroupLis) AddStreamGroups () error {
 //parameters: SourcePortHandle Type NumberOfStreamGroups
 //AgtStreamGroupList AddStreamGroups, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) AddStreamGroupsWithExistingProfile () error {
 //parameters: ProfileHandle Type NumberOfStreamGroups
 //AgtStreamGroupList AddStreamGroupsWithExistingProfile, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) AddStreamGroupsWithNewProfile () error {
 //parameters: SourcePortHandle Type NumberOfStreamGroups
 //AgtStreamGroupList AddStreamGroupsWithNewProfile, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) Clear () error {
 //parameters: 
 //AgtStreamGroupList Clear, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) SetGlobalRandomSeed () error {
 //parameters: GlobalRandomSeed
 //AgtStreamGroupList SetGlobalRandomSeed, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) GetGlobalRandomSeed ()(string, error) {
 //parameters: 
 //AgtStreamGroupList GetGlobalRandomSeed
 return "", nil
}

func(np *StreamGroupLis) SetPayloadProtection () error {
 //parameters: SourcePortHandle PayloadType CustomByteOffset
 //AgtStreamGroupList SetPayloadProtection, m.Object, m.Name);
 return nil
}

func(np *StreamGroupLis) GetPayloadProtection ()(string, error) {
 //parameters: SourcePortHandle
 //AgtStreamGroupList GetPayloadProtection
 return "", nil
}

func(np *StreamGroupLis) GetMaximumNumberOfStreamGroupsOnPort ()(string, error) {
 //parameters: SourcePortHandle
 //AgtStreamGroupList GetMaximumNumberOfStreamGroupsOnPort
 return "", nil
}

