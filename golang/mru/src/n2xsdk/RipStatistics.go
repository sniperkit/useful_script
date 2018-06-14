package n2xsdk

type RipStatistics struct {
 Handler string
}

func(np *RipStatistics) ListSelectedSessions ()(string, error) {
 //parameters: 
 //AgtRipStatistics ListSelectedSessions
 return "", nil
}

func(np *RipStatistics) SelectSessions () error {
 //parameters: Count psaSessions
 //AgtRipStatistics SelectSessions, m.Object, m.Name);
 return nil
}

func(np *RipStatistics) DeselectSession () error {
 //parameters: SessionHandle
 //AgtRipStatistics DeselectSession, m.Object, m.Name);
 return nil
}

func(np *RipStatistics) GetSessionAccumulatedValues ()(string, error) {
 //parameters: SessionHandle
 //AgtRipStatistics GetSessionAccumulatedValues
 return "", nil
}

