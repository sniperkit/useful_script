package n2xsdk

type IsisIpv6RoutePool struct {
 Handler string
}

func(np *IsisIpv6RoutePool) SetRoutes () error {
 //parameters: RoutePoolHandle FirstRoute PrefixLength NumRoutes Modifier
 //AgtIsisIpv6RoutePool SetRoutes, m.Object, m.Name);
 return nil
}

func(np *IsisIpv6RoutePool) GetRoutes ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtIsisIpv6RoutePool GetRoutes
 return "", nil
}

func(np *IsisIpv6RoutePool) EnableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtIsisIpv6RoutePool EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *IsisIpv6RoutePool) DisableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtIsisIpv6RoutePool DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *IsisIpv6RoutePool) IsTrafficDestinationEnabled () error {
 //parameters: RoutePoolHandle
 //AgtIsisIpv6RoutePool IsTrafficDestinationEnabled, m.Object, m.Name);
 return nil
}

func(np *IsisIpv6RoutePool) SetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtIsisIpv6RoutePool SetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *IsisIpv6RoutePool) UnsetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtIsisIpv6RoutePool UnsetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *IsisIpv6RoutePool) IsRoutePoolFlagSet () error {
 //parameters: RoutePoolHandle Flag
 //AgtIsisIpv6RoutePool IsRoutePoolFlagSet, m.Object, m.Name);
 return nil
}

func(np *IsisIpv6RoutePool) GetRoutePoolStatus ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtIsisIpv6RoutePool GetRoutePoolStatus
 return "", nil
}

func(np *IsisIpv6RoutePool) GetNthRoute ()(string, error) {
 //parameters: RoutePoolHandle Index
 //AgtIsisIpv6RoutePool GetNthRoute
 return "", nil
}

func(np *IsisIpv6RoutePool) SetResourceClasses () error {
 //parameters: RoutePoolHandle Count psaClasses
 //AgtIsisIpv6RoutePool SetResourceClasses, m.Object, m.Name);
 return nil
}

func(np *IsisIpv6RoutePool) GetResourceClasses ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtIsisIpv6RoutePool GetResourceClasses
 return "", nil
}

func(np *IsisIpv6RoutePool) GetReachablePrefixCount ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtIsisIpv6RoutePool GetReachablePrefixCount
 return "", nil
}

