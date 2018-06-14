package n2xsdk

type RsvpTunnelEgressLis struct {
 Handler string
}

func(np *RsvpTunnelEgressLis) EnableEgress () error {
 //parameters: SessionHandle
 //AgtRsvpTunnelEgressList EnableEgress, m.Object, m.Name);
 return nil
}

func(np *RsvpTunnelEgressLis) EnableAllEgress () error {
 //parameters: 
 //AgtRsvpTunnelEgressList EnableAllEgress, m.Object, m.Name);
 return nil
}

func(np *RsvpTunnelEgressLis) DisableEgress () error {
 //parameters: SessionHandle
 //AgtRsvpTunnelEgressList DisableEgress, m.Object, m.Name);
 return nil
}

func(np *RsvpTunnelEgressLis) DisableAllEgress () error {
 //parameters: 
 //AgtRsvpTunnelEgressList DisableAllEgress, m.Object, m.Name);
 return nil
}

func(np *RsvpTunnelEgressLis) IsEgressEnabled () error {
 //parameters: SessionHandle
 //AgtRsvpTunnelEgressList IsEgressEnabled, m.Object, m.Name);
 return nil
}

func(np *RsvpTunnelEgressLis) ListTunnels ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpTunnelEgressList ListTunnels
 return "", nil
}

func(np *RsvpTunnelEgressLis) GetEnableFlag ()(string, error) {
 //parameters: SessionHandle
 //AgtRsvpTunnelEgressList GetEnableFlag
 return "", nil
}

