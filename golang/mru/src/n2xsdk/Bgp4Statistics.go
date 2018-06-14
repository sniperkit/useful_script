package n2xsdk

type Bgp4Statistics struct {
 Handler string
}

func(np *Bgp4Statistics) SelectPorts () error {
 //parameters: Count saPorts
 //AgtBgp4Statistics SelectPorts, m.Object, m.Name);
 return nil
}

func(np *Bgp4Statistics) ListSelectedPorts ()(string, error) {
 //parameters: 
 //AgtBgp4Statistics ListSelectedPorts
 return "", nil
}

func(np *Bgp4Statistics) ListSelectedSessions ()(string, error) {
 //parameters: 
 //AgtBgp4Statistics ListSelectedSessions
 return "", nil
}

func(np *Bgp4Statistics) SelectSessions () error {
 //parameters: Count psaSessions
 //AgtBgp4Statistics SelectSessions, m.Object, m.Name);
 return nil
}

func(np *Bgp4Statistics) DeselectSession () error {
 //parameters: SessionHandle
 //AgtBgp4Statistics DeselectSession, m.Object, m.Name);
 return nil
}

func(np *Bgp4Statistics) GetSessionAccumulatedValues ()(string, error) {
 //parameters: SessionIdCount psaSessionId
 //AgtBgp4Statistics GetSessionAccumulatedValues
 return "", nil
}

