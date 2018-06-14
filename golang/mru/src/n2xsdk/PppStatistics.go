package n2xsdk

type PppStatistics struct {
 Handler string
}

func(np *PppStatistics) ListSelectedPools ()(string, error) {
 //parameters: 
 //AgtPppStatistics ListSelectedPools
 return "", nil
}

func(np *PppStatistics) SelectPools () error {
 //parameters: Count psaPoolHandles
 //AgtPppStatistics SelectPools, m.Object, m.Name);
 return nil
}

func(np *PppStatistics) DeselectPool () error {
 //parameters: SessionPoolHandle
 //AgtPppStatistics DeselectPool, m.Object, m.Name);
 return nil
}

func(np *PppStatistics) GetAccumulatedValues ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtPppStatistics GetAccumulatedValues
 return "", nil
}

