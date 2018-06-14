package n2xsdk

type PatternLis struct {
 Handler string
}

func(np *PatternLis) Add () error {
 //parameters: Type
 //AgtPatternList Add, m.Object, m.Name);
 return nil
}

func(np *PatternLis) AddItems () error {
 //parameters: Type NumberOfItems
 //AgtPatternList AddItems, m.Object, m.Name);
 return nil
}

func(np *PatternLis) Remove () error {
 //parameters: Handle
 //AgtPatternList Remove, m.Object, m.Name);
 return nil
}

func(np *PatternLis) Copy () error {
 //parameters: Handle
 //AgtPatternList Copy, m.Object, m.Name);
 return nil
}

func(np *PatternLis) SetName () error {
 //parameters: Handle Name
 //AgtPatternList SetName, m.Object, m.Name);
 return nil
}

func(np *PatternLis) GetName ()(string, error) {
 //parameters: Handle
 //AgtPatternList GetName
 return "", nil
}

func(np *PatternLis) GetHandle ()(string, error) {
 //parameters: Name
 //AgtPatternList GetHandle
 return "", nil
}

func(np *PatternLis) GetType ()(string, error) {
 //parameters: Handle
 //AgtPatternList GetType
 return "", nil
}

func(np *PatternLis) LockItem () error {
 //parameters: Handle
 //AgtPatternList LockItem, m.Object, m.Name);
 return nil
}

func(np *PatternLis) UnlockItem () error {
 //parameters: Handle
 //AgtPatternList UnlockItem, m.Object, m.Name);
 return nil
}

func(np *PatternLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtPatternList ListHandles
 return "", nil
}

func(np *PatternLis) ListNames ()(string, error) {
 //parameters: 
 //AgtPatternList ListNames
 return "", nil
}

func(np *PatternLis) ListHandlesByType ()(string, error) {
 //parameters: Type
 //AgtPatternList ListHandlesByType
 return "", nil
}

func(np *PatternLis) ListTypes ()(string, error) {
 //parameters: 
 //AgtPatternList ListTypes
 return "", nil
}

func(np *PatternLis) GetInterface ()(string, error) {
 //parameters: Type
 //AgtPatternList GetInterface
 return "", nil
}

func(np *PatternLis) GetLockCount ()(string, error) {
 //parameters: 
 //AgtPatternList GetLockCount
 return "", nil
}

