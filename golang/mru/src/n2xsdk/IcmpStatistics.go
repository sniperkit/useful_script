package n2xsdk

type IcmpStatistics struct {
 Handler string
}

func(np *IcmpStatistics) SelectPorts () error {
 //parameters: Count saPorts
 //AgtIcmpStatistics SelectPorts, m.Object, m.Name);
 return nil
}

func(np *IcmpStatistics) ListSelectedPorts ()(string, error) {
 //parameters: 
 //AgtIcmpStatistics ListSelectedPorts
 return "", nil
}

func(np *IcmpStatistics) GetAccumulatedValues ()(string, error) {
 //parameters: PortHandle
 //AgtIcmpStatistics GetAccumulatedValues
 return "", nil
}

