package n2xsdk

type RipngRoutePool struct {
 Handler string
}

func(np *RipngRoutePool) SetRoutes () error {
 //parameters: RoutePoolHandle FirstRoute PrefixLength NumRoutes Modifier
 //AgtRipngRoutePool SetRoutes, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) GetRoutes ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtRipngRoutePool GetRoutes
 return "", nil
}

func(np *RipngRoutePool) EnableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtRipngRoutePool EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) DisableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtRipngRoutePool DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) IsTrafficDestinationEnabled () error {
 //parameters: RoutePoolHandle
 //AgtRipngRoutePool IsTrafficDestinationEnabled, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) SetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtRipngRoutePool SetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) UnsetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtRipngRoutePool UnsetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) IsRoutePoolFlagSet () error {
 //parameters: RoutePoolHandle Flag
 //AgtRipngRoutePool IsRoutePoolFlagSet, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) GetRoutePoolStatus ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtRipngRoutePool GetRoutePoolStatus
 return "", nil
}

func(np *RipngRoutePool) GetNthRoute ()(string, error) {
 //parameters: RoutePoolHandle Index
 //AgtRipngRoutePool GetNthRoute
 return "", nil
}

func(np *RipngRoutePool) SetResourceClasses () error {
 //parameters: RoutePoolHandle Count psaClasses
 //AgtRipngRoutePool SetResourceClasses, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) GetResourceClasses ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtRipngRoutePool GetResourceClasses
 return "", nil
}

func(np *RipngRoutePool) SetMetric () error {
 //parameters: hRoutePool Metric
 //AgtRipngRoutePool SetMetric, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) GetMetric ()(string, error) {
 //parameters: hRoutePool
 //AgtRipngRoutePool GetMetric
 return "", nil
}

func(np *RipngRoutePool) SetRouteTag () error {
 //parameters: hRoutePool RouteTag
 //AgtRipngRoutePool SetRouteTag, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) GetRouteTag ()(string, error) {
 //parameters: hRoutePool
 //AgtRipngRoutePool GetRouteTag
 return "", nil
}

func(np *RipngRoutePool) SetAddressFamily () error {
 //parameters: hRoutePool AddressFamily
 //AgtRipngRoutePool SetAddressFamily, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) GetAddressFamily ()(string, error) {
 //parameters: hRoutePool
 //AgtRipngRoutePool GetAddressFamily
 return "", nil
}

func(np *RipngRoutePool) SetNextHop () error {
 //parameters: hRoutePool NextHop
 //AgtRipngRoutePool SetNextHop, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) GetNextHop ()(string, error) {
 //parameters: hRoutePool
 //AgtRipngRoutePool GetNextHop
 return "", nil
}

func(np *RipngRoutePool) Advertise () error {
 //parameters: hRoutePool
 //AgtRipngRoutePool Advertise, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) Withdraw () error {
 //parameters: hRoutePool
 //AgtRipngRoutePool Withdraw, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) AdvertisePools () error {
 //parameters: Count saRoutePools
 //AgtRipngRoutePool AdvertisePools, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) WithdrawPools () error {
 //parameters: Count saRoutePools
 //AgtRipngRoutePool WithdrawPools, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) GetSessionHandle ()(string, error) {
 //parameters: hRoutePool
 //AgtRipngRoutePool GetSessionHandle
 return "", nil
}

func(np *RipngRoutePool) Enable () error {
 //parameters: hRoutePool
 //AgtRipngRoutePool Enable, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) Disable () error {
 //parameters: hRoutePool
 //AgtRipngRoutePool Disable, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) IsEnabled () error {
 //parameters: hRoutePool
 //AgtRipngRoutePool IsEnabled, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) SelectNextHop () error {
 //parameters: hRoutePool
 //AgtRipngRoutePool SelectNextHop, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) DeselectNextHop () error {
 //parameters: hRoutePool
 //AgtRipngRoutePool DeselectNextHop, m.Object, m.Name);
 return nil
}

func(np *RipngRoutePool) IsNextHopSelected () error {
 //parameters: hRoutePool
 //AgtRipngRoutePool IsNextHopSelected, m.Object, m.Name);
 return nil
}

