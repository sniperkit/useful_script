package n2xsdk

type RipngRouteList struct {
 Handler string
}

func(np *RipngRouteLis) AddRoutePools () error {
 //parameters: SessionHandle PfxCount saPrefixLengths RouteCount saRoutes
 //AgtRipngRouteList AddRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipngRouteLis) ListRoutePools ()(string, error) {
 //parameters: SessionHandle
 //AgtRipngRouteList ListRoutePools
 return "", nil
}

func(np *RipngRouteLis) RemoveRoutePools () error {
 //parameters: SessionHandle Count saRoutePools
 //AgtRipngRouteList RemoveRoutePools, m.Object, m.Name);
 return nil
}

func(np *RipngRouteLis) ClearRoutePools () error {
 //parameters: SessionHandle
 //AgtRipngRouteList ClearRoutePools, m.Object, m.Name);
 return nil
}

