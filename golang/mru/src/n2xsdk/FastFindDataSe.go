package n2xsdk

type FastFindDataSe struct {
 Handler string
}

func(np *FastFindDataSe) Save () error {
 //parameters: DataSetHandle Filename
 //AgtFastFindDataSet Save, m.Object, m.Name);
 return nil
}

func(np *FastFindDataSe) GetProperties ()(string, error) {
 //parameters: DataSetHandle
 //AgtFastFindDataSet GetProperties
 return "", nil
}

func(np *FastFindDataSe) GetDefaultDirectory ()(string, error) {
 //parameters: 
 //AgtFastFindDataSet GetDefaultDirectory
 return "", nil
}

