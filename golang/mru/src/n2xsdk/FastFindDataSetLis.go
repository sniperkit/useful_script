package n2xsdk

type FastFindDataSetLis struct {
 Handler string
}

func(np *FastFindDataSetLis) ListHandles ()(string, error) {
 //parameters: 
 //AgtFastFindDataSetList ListHandles
 return "", nil
}

func(np *FastFindDataSetLis) GetFastFindSessionDataSet ()(string, error) {
 //parameters: 
 //AgtFastFindDataSetList GetFastFindSessionDataSet
 return "", nil
}

func(np *FastFindDataSetLis) OpenFastFindDataSetFile () error {
 //parameters: Filename
 //AgtFastFindDataSetList OpenFastFindDataSetFile, m.Object, m.Name);
 return nil
}

func(np *FastFindDataSetLis) CloseFastFindDataSetFile () error {
 //parameters: DataSetHandle
 //AgtFastFindDataSetList CloseFastFindDataSetFile, m.Object, m.Name);
 return nil
}

