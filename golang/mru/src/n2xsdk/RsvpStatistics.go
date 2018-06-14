package n2xsdk

type RsvpStatistics struct {
 Handler string
}

func(np *RsvpStatistics) ListSelectedSessions ()(string, error) {
 //parameters: 
 //AgtRsvpStatistics ListSelectedSessions
 return "", nil
}

func(np *RsvpStatistics) SelectSessions () error {
 //parameters: Count psaSessions
 //AgtRsvpStatistics SelectSessions, m.Object, m.Name);
 return nil
}

func(np *RsvpStatistics) DeselectSession () error {
 //parameters: SessionHandle
 //AgtRsvpStatistics DeselectSession, m.Object, m.Name);
 return nil
}

func(np *RsvpStatistics) GetSessionAccumulatedValues ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpStatistics GetSessionAccumulatedValues
 return "", nil
}

