package n2xsdk

type RipRouteList struct {
 Handler string
}

func(np *RipRouteLis) AddRoutePools () error {
 //parameters: SessionHandle PfxCount saPrefixLengths RouteCount saRoutes
 //AgtRipRouteList AddRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipRouteLis) ListRoutePools ()(string, error) {
 //parameters: SessionHandle
 //AgtRipRouteList ListRoutePools
 return "", nil
}

func(np *RipRouteLis) RemoveRoutePools () error {
 //parameters: SessionHandle Count saRoutePools
 //AgtRipRouteList RemoveRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipRouteLis) ClearRoutePools () error {
 //parameters: SessionHandle
 //AgtRipRouteList ClearRoutePools, m.Object, m.Name);
 return nil
}

