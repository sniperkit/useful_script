package n2xsdk

type FastFindViewLis struct {
 Handler string
}

func(np *FastFindViewLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtFastFindViewList ListHandles
 return "", nil
}

func(np *FastFindViewLis) Add () error {
 //parameters: 
 //AgtFastFindViewList Add, m.Object, m.Name);
 return nil
}

func(np *FastFindViewLis) Remove () error {
 //parameters: ViewHandle
 //AgtFastFindViewList Remove, m.Object, m.Name);
 return nil
}

func(np *FastFindViewLis) Copy () error {
 //parameters: ViewHandle
 //AgtFastFindViewList Copy, m.Object, m.Name);
 return nil
}

