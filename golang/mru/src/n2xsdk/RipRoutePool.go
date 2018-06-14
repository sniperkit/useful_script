package n2xsdk

type RipRoutePool struct {
 Handler string
}

func(np *RipRoutePool) SetRoutes () error {
 //parameters: RoutePoolHandle FirstRoute PrefixLength NumRoutes Modifier
 //AgtRipRoutePool SetRoutes, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) GetRoutes ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtRipRoutePool GetRoutes
 return "", nil
}

func(np *RipRoutePool) EnableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtRipRoutePool EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) DisableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtRipRoutePool DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) IsTrafficDestinationEnabled () error {
 //parameters: RoutePoolHandle
 //AgtRipRoutePool IsTrafficDestinationEnabled, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) SetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtRipRoutePool SetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) UnsetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtRipRoutePool UnsetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) IsRoutePoolFlagSet () error {
 //parameters: RoutePoolHandle Flag
 //AgtRipRoutePool IsRoutePoolFlagSet, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) GetRoutePoolStatus ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtRipRoutePool GetRoutePoolStatus
 return "", nil
}

func(np *RipRoutePool) GetNthRoute ()(string, error) {
 //parameters: RoutePoolHandle Index
 //AgtRipRoutePool GetNthRoute
 return "", nil
}

func(np *RipRoutePool) SetResourceClasses () error {
 //parameters: RoutePoolHandle Count psaClasses
 //AgtRipRoutePool SetResourceClasses, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) GetResourceClasses ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtRipRoutePool GetResourceClasses
 return "", nil
}

func(np *RipRoutePool) SetMetric () error {
 //parameters: hRoutePool Metric
 //AgtRipRoutePool SetMetric, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) GetMetric ()(string, error) {
 //parameters: hRoutePool
 //AgtRipRoutePool GetMetric
 return "", nil
}

func(np *RipRoutePool) SetRouteTag () error {
 //parameters: hRoutePool RouteTag
 //AgtRipRoutePool SetRouteTag, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) GetRouteTag ()(string, error) {
 //parameters: hRoutePool
 //AgtRipRoutePool GetRouteTag
 return "", nil
}

func(np *RipRoutePool) SetAddressFamily () error {
 //parameters: hRoutePool AddressFamily
 //AgtRipRoutePool SetAddressFamily, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) GetAddressFamily ()(string, error) {
 //parameters: hRoutePool
 //AgtRipRoutePool GetAddressFamily
 return "", nil
}

func(np *RipRoutePool) SetNextHop () error {
 //parameters: hRoutePool NextHop
 //AgtRipRoutePool SetNextHop, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) GetNextHop ()(string, error) {
 //parameters: hRoutePool
 //AgtRipRoutePool GetNextHop
 return "", nil
}

func(np *RipRoutePool) Advertise () error {
 //parameters: hRoutePool
 //AgtRipRoutePool Advertise, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) Withdraw () error {
 //parameters: hRoutePool
 //AgtRipRoutePool Withdraw, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) AdvertisePools () error {
 //parameters: Count saRoutePools
 //AgtRipRoutePool AdvertisePools, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) WithdrawPools () error {
 //parameters: Count saRoutePools
 //AgtRipRoutePool WithdrawPools, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) GetSessionHandle ()(string, error) {
 //parameters: hRoutePool
 //AgtRipRoutePool GetSessionHandle
 return "", nil
}

func(np *RipRoutePool) Enable () error {
 //parameters: hRoutePool
 //AgtRipRoutePool Enable, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) Disable () error {
 //parameters: hRoutePool
 //AgtRipRoutePool Disable, m.Object, m.Name);
 return nil
}

func(np *RipRoutePool) IsEnabled () error {
 //parameters: hRoutePool
 //AgtRipRoutePool IsEnabled, m.Object, m.Name);
 return nil
}

