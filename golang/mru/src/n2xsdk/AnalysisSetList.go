package n2xsdk

type AnalysisSetList struct {
 Handler string
}

func(np *nalysisSetLis) Add () error {
 //parameters: Type
 //AgtAnalysisSetList Add, m.Object, m.Name);
 return nil
}

func(np *nalysisSetLis) AddItems () error {
 //parameters: Type NumberOfItems
 //AgtAnalysisSetList AddItems, m.Object, m.Name);
 return nil
}

func(np *nalysisSetLis) Remove () error {
 //parameters: Handle
 //AgtAnalysisSetList Remove, m.Object, m.Name);
 return nil
}

func(np *nalysisSetLis) Copy () error {
 //parameters: Handle
 //AgtAnalysisSetList Copy, m.Object, m.Name);
 return nil
}

func(np *nalysisSetLis) SetName () error {
 //parameters: Handle Name
 //AgtAnalysisSetList SetName, m.Object, m.Name);
 return nil
}

func(np *nalysisSetLis) GetName ()(string, error) {
 //parameters: Handle
 //AgtAnalysisSetList GetName
 return "", nil
}

func(np *nalysisSetLis) GetHandle ()(string, error) {
 //parameters: Name
 //AgtAnalysisSetList GetHandle
 return "", nil
}

func(np *nalysisSetLis) GetType ()(string, error) {
 //parameters: Handle
 //AgtAnalysisSetList GetType
 return "", nil
}

func(np *nalysisSetLis) LockItem () error {
 //parameters: Handle
 //AgtAnalysisSetList LockItem, m.Object, m.Name);
 return nil
}

func(np *nalysisSetLis) UnlockItem () error {
 //parameters: Handle
 //AgtAnalysisSetList UnlockItem, m.Object, m.Name);
 return nil
}

func(np *nalysisSetLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtAnalysisSetList ListHandles
 return "", nil
}

func(np *nalysisSetLis) ListNames ()(string, error) {
 //parameters: 
 //AgtAnalysisSetList ListNames
 return "", nil
}

func(np *nalysisSetLis) ListHandlesByType ()(string, error) {
 //parameters: Type
 //AgtAnalysisSetList ListHandlesByType
 return "", nil
}

func(np *nalysisSetLis) ListTypes ()(string, error) {
 //parameters: 
 //AgtAnalysisSetList ListTypes
 return "", nil
}

func(np *nalysisSetLis) GetInterface ()(string, error) {
 //parameters: Type
 //AgtAnalysisSetList GetInterface
 return "", nil
}

func(np *nalysisSetLis) GetLockCount ()(string, error) {
 //parameters: 
 //AgtAnalysisSetList GetLockCount
 return "", nil
}

func(np *nalysisSetLis) GetHandleByRequestID ()(string, error) {
 //parameters: AnalysisRequestID
 //AgtAnalysisSetList GetHandleByRequestID
 return "", nil
}

func(np *nalysisSetLis) WaitForAnalysisSets () error {
 //parameters: NumRequests psaAnalysisRequestIDs
 //AgtAnalysisSetList WaitForAnalysisSets, m.Object, m.Name);
 return nil
}

