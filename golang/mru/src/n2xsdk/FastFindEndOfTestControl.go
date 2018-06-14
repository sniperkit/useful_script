package n2xsdk

type FastFindEndOfTestControl struct {
 Handler string
}

func(np *FastFindEndOfTestControl) RetrieveData () error {
 //parameters: 
 //AgtFastFindEndOfTestControl RetrieveData, m.Object, m.Name);
 return nil
}

func(np *FastFindEndOfTestControl) CancelDataRetrieval () error {
 //parameters: 
 //AgtFastFindEndOfTestControl CancelDataRetrieval, m.Object, m.Name);
 return nil
}

func(np *FastFindEndOfTestControl) GetDataRetrievalStatus ()(string, error) {
 //parameters: 
 //AgtFastFindEndOfTestControl GetDataRetrievalStatus
 return "", nil
}

func(np *FastFindEndOfTestControl) EnableDataRetrievalOnAllPorts () error {
 //parameters: EnableOnAllPorts
 //AgtFastFindEndOfTestControl EnableDataRetrievalOnAllPorts, m.Object, m.Name);
 return nil
}

func(np *FastFindEndOfTestControl) IsDataRetrievalSetOnAllPortsEnabled () error {
 //parameters: 
 //AgtFastFindEndOfTestControl IsDataRetrievalSetOnAllPortsEnabled, m.Object, m.Name);
 return nil
}

func(np *FastFindEndOfTestControl) SelectPortsForDataRetrieval () error {
 //parameters: PortHandles
 //AgtFastFindEndOfTestControl SelectPortsForDataRetrieval, m.Object, m.Name);
 return nil
}

func(np *FastFindEndOfTestControl) ListSelectedPortsForDataRetrieval ()(string, error) {
 //parameters: 
 //AgtFastFindEndOfTestControl ListSelectedPortsForDataRetrieval
 return "", nil
}

func(np *FastFindEndOfTestControl) IsAutoRetrieveDataEnabled () error {
 //parameters: 
 //AgtFastFindEndOfTestControl IsAutoRetrieveDataEnabled, m.Object, m.Name);
 return nil
}

func(np *FastFindEndOfTestControl) EnableAutoRetrieveData () error {
 //parameters: 
 //AgtFastFindEndOfTestControl EnableAutoRetrieveData, m.Object, m.Name);
 return nil
}

func(np *FastFindEndOfTestControl) DisableAutoRetrieveData () error {
 //parameters: 
 //AgtFastFindEndOfTestControl DisableAutoRetrieveData, m.Object, m.Name);
 return nil
}

