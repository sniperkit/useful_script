package n2xsdk

type LdpStatistics struct {
 Handler string
}

func(np *LdpStatistics) ListSelectedSessions ()(string, error) {
 //parameters: 
 //AgtLdpStatistics ListSelectedSessions
 return "", nil
}

func(np *LdpStatistics) SelectSessions () error {
 //parameters: Count psaSessions
 //AgtLdpStatistics SelectSessions, m.Object, m.Name);
 return nil
}

func(np *LdpStatistics) DeselectSession () error {
 //parameters: SessionHandle
 //AgtLdpStatistics DeselectSession, m.Object, m.Name);
 return nil
}

func(np *LdpStatistics) GetSessionAccumulatedValues ()(string, error) {
 //parameters: SessionHandle
 //AgtLdpStatistics GetSessionAccumulatedValues
 return "", nil
}

