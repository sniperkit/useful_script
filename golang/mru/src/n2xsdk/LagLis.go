package n2xsdk

type LagLis struct {
 Handler string
}

func(np *LagLis) Add () error {
 //parameters: Type
 //AgtLagList Add, m.Object, m.Name);
 return nil
}

func(np *LagLis) Remove () error {
 //parameters: Handle
 //AgtLagList Remove, m.Object, m.Name);
 return nil
}

func(np *LagLis) SetName () error {
 //parameters: Handle Name
 //AgtLagList SetName, m.Object, m.Name);
 return nil
}

func(np *LagLis) GetName ()(string, error) {
 //parameters: Handle
 //AgtLagList GetName
 return "", nil
}

func(np *LagLis) GetHandle ()(string, error) {
 //parameters: Name
 //AgtLagList GetHandle
 return "", nil
}

func(np *LagLis) GetType ()(string, error) {
 //parameters: Handle
 //AgtLagList GetType
 return "", nil
}

func(np *LagLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtLagList ListHandles
 return "", nil
}

func(np *LagLis) ListNames ()(string, error) {
 //parameters: 
 //AgtLagList ListNames
 return "", nil
}

func(np *LagLis) ListHandlesByType ()(string, error) {
 //parameters: Type
 //AgtLagList ListHandlesByType
 return "", nil
}

func(np *LagLis) ListTypes ()(string, error) {
 //parameters: 
 //AgtLagList ListTypes
 return "", nil
}

