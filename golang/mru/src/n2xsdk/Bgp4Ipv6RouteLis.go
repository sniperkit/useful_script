package n2xsdk

type Bgp4Ipv6RouteLis struct {
 Handler string
}

func(np *Bgp4Ipv6RouteLis) AddRoutePools () error {
 //parameters: SessionHandle PfxCount saPrefixLengths RouteCount saRoutes
 //AgtBgp4Ipv6RouteList AddRoutePools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6RouteLis) ListRoutePools ()(string, error) {
 //parameters: SessionHandle
 //AgtBgp4Ipv6RouteList ListRoutePools
 return "", nil
}

func(np *Bgp4Ipv6RouteLis) RemoveRoutePools () error {
 //parameters: SessionHandle Count saRoutePools
 //AgtBgp4Ipv6RouteList RemoveRoutePools, m.Object, m.Name);
 return nil
}

func(np *Bgp4Ipv6RouteLis) ClearRoutePools () error {
 //parameters: SessionHandle
 //AgtBgp4Ipv6RouteList ClearRoutePools, m.Object, m.Name);
 return nil
}

