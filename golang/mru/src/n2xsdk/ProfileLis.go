package n2xsdk

type ProfileLis struct {
 Handler string
}

func(np *ProfileLis) Add () error {
 //parameters: Type
 //AgtProfileList Add, m.Object, m.Name);
 return nil
}

func(np *ProfileLis) AddItems () error {
 //parameters: Type NumberOfItems
 //AgtProfileList AddItems, m.Object, m.Name);
 return nil
}

func(np *ProfileLis) Remove () error {
 //parameters: Handle
 //AgtProfileList Remove, m.Object, m.Name);
 return nil
}

func(np *ProfileLis) Copy () error {
 //parameters: Handle
 //AgtProfileList Copy, m.Object, m.Name);
 return nil
}

func(np *ProfileLis) SetName () error {
 //parameters: Handle Name
 //AgtProfileList SetName, m.Object, m.Name);
 return nil
}

func(np *ProfileLis) GetName ()(string, error) {
 //parameters: Handle
 //AgtProfileList GetName
 return "", nil
}

func(np *ProfileLis) GetHandle ()(string, error) {
 //parameters: Name
 //AgtProfileList GetHandle
 return "", nil
}

func(np *ProfileLis) GetType ()(string, error) {
 //parameters: Handle
 //AgtProfileList GetType
 return "", nil
}

func(np *ProfileLis) LockItem () error {
 //parameters: Handle
 //AgtProfileList LockItem, m.Object, m.Name);
 return nil
}

func(np *ProfileLis) UnlockItem () error {
 //parameters: Handle
 //AgtProfileList UnlockItem, m.Object, m.Name);
 return nil
}

func(np *ProfileLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtProfileList ListHandles
 return "", nil
}

func(np *ProfileLis) ListNames ()(string, error) {
 //parameters: 
 //AgtProfileList ListNames
 return "", nil
}

func(np *ProfileLis) ListHandlesByType ()(string, error) {
 //parameters: Type
 //AgtProfileList ListHandlesByType
 return "", nil
}

func(np *ProfileLis) ListTypes ()(string, error) {
 //parameters: 
 //AgtProfileList ListTypes
 return "", nil
}

func(np *ProfileLis) GetInterface ()(string, error) {
 //parameters: Type
 //AgtProfileList GetInterface
 return "", nil
}

func(np *ProfileLis) GetLockCount ()(string, error) {
 //parameters: 
 //AgtProfileList GetLockCount
 return "", nil
}

func(np *ProfileLis) AddProfile () error {
 //parameters: SourcePortHandle Type
 //AgtProfileList AddProfile, m.Object, m.Name);
 return nil
}

func(np *ProfileLis) ListProfilesOnPort ()(string, error) {
 //parameters: SourcePortHandle
 //AgtProfileList ListProfilesOnPort
 return "", nil
}

func(np *ProfileLis) ListProfilesOnPortByType ()(string, error) {
 //parameters: SourcePortHandle ProfileType
 //AgtProfileList ListProfilesOnPortByType
 return "", nil
}

func(np *ProfileLis) GetMaximumNumberOfProfilesOnPort ()(string, error) {
 //parameters: SourcePortHandle
 //AgtProfileList GetMaximumNumberOfProfilesOnPort
 return "", nil
}

