package n2xsdk

type nalysisFilterLis struct {
 Handler string
}

func(np *nalysisFilterLis) Add () error {
 //parameters: Type
 //AgtAnalysisFilterList Add, m.Object, m.Name);
 return nil
}

func(np *nalysisFilterLis) AddItems () error {
 //parameters: Type NumberOfItems
 //AgtAnalysisFilterList AddItems, m.Object, m.Name);
 return nil
}

func(np *nalysisFilterLis) Remove () error {
 //parameters: Handle
 //AgtAnalysisFilterList Remove, m.Object, m.Name);
 return nil
}

func(np *nalysisFilterLis) Copy () error {
 //parameters: Handle
 //AgtAnalysisFilterList Copy, m.Object, m.Name);
 return nil
}

func(np *nalysisFilterLis) SetName () error {
 //parameters: Handle Name
 //AgtAnalysisFilterList SetName, m.Object, m.Name);
 return nil
}

func(np *nalysisFilterLis) GetName ()(string, error) {
 //parameters: Handle
 //AgtAnalysisFilterList GetName
 return "", nil
}

func(np *nalysisFilterLis) GetHandle ()(string, error) {
 //parameters: Name
 //AgtAnalysisFilterList GetHandle
 return "", nil
}

func(np *nalysisFilterLis) GetType ()(string, error) {
 //parameters: Handle
 //AgtAnalysisFilterList GetType
 return "", nil
}

func(np *nalysisFilterLis) LockItem () error {
 //parameters: Handle
 //AgtAnalysisFilterList LockItem, m.Object, m.Name);
 return nil
}

func(np *nalysisFilterLis) UnlockItem () error {
 //parameters: Handle
 //AgtAnalysisFilterList UnlockItem, m.Object, m.Name);
 return nil
}

func(np *nalysisFilterLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtAnalysisFilterList ListHandles
 return "", nil
}

func(np *nalysisFilterLis) ListNames ()(string, error) {
 //parameters: 
 //AgtAnalysisFilterList ListNames
 return "", nil
}

func(np *nalysisFilterLis) ListHandlesByType ()(string, error) {
 //parameters: Type
 //AgtAnalysisFilterList ListHandlesByType
 return "", nil
}

func(np *nalysisFilterLis) ListTypes ()(string, error) {
 //parameters: 
 //AgtAnalysisFilterList ListTypes
 return "", nil
}

func(np *nalysisFilterLis) GetInterface ()(string, error) {
 //parameters: Type
 //AgtAnalysisFilterList GetInterface
 return "", nil
}

func(np *nalysisFilterLis) GetLockCount ()(string, error) {
 //parameters: 
 //AgtAnalysisFilterList GetLockCount
 return "", nil
}

