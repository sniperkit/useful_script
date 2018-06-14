package n2xsdk

type RsvpTunnelIngressLis struct {
 Handler string
}

func(np *RsvpTunnelIngressLis) AddTunnels () error {
 //parameters: SessionHandle TunnelListSize psaTunnelIds SourceListSize psaSourceAddresses DestListSize psaDestAddresses
 //AgtRsvpTunnelIngressList AddTunnels, m.Object, m.Name);
 return nil
}

func(np *RsvpTunnelIngressLis) RemoveTunnels () error {
 //parameters: SessionHandle Count psaHandles
 //AgtRsvpTunnelIngressList RemoveTunnels, m.Object, m.Name);
 return nil
}

