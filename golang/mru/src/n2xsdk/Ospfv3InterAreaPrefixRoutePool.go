package n2xsdk

type Ospfv3InterAreaPrefixRoutePool struct {
 Handler string
}

func(np *Ospfv3InterAreaPrefixRoutePool) SetRoutes () error {
 //parameters: RoutePoolHandle FirstRoute PrefixLength NumRoutes Modifier
 //AgtOspfv3InterAreaPrefixRoutePool SetRoutes, m.Object, m.Name);
 return nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) GetRoutes ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfv3InterAreaPrefixRoutePool GetRoutes
 return "", nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) EnableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtOspfv3InterAreaPrefixRoutePool EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) DisableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtOspfv3InterAreaPrefixRoutePool DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) IsTrafficDestinationEnabled () error {
 //parameters: RoutePoolHandle
 //AgtOspfv3InterAreaPrefixRoutePool IsTrafficDestinationEnabled, m.Object, m.Name);
 return nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) SetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtOspfv3InterAreaPrefixRoutePool SetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) UnsetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtOspfv3InterAreaPrefixRoutePool UnsetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) IsRoutePoolFlagSet () error {
 //parameters: RoutePoolHandle Flag
 //AgtOspfv3InterAreaPrefixRoutePool IsRoutePoolFlagSet, m.Object, m.Name);
 return nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) GetRoutePoolStatus ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfv3InterAreaPrefixRoutePool GetRoutePoolStatus
 return "", nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) GetNthRoute ()(string, error) {
 //parameters: RoutePoolHandle Index
 //AgtOspfv3InterAreaPrefixRoutePool GetNthRoute
 return "", nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) SetResourceClasses () error {
 //parameters: RoutePoolHandle Count psaClasses
 //AgtOspfv3InterAreaPrefixRoutePool SetResourceClasses, m.Object, m.Name);
 return nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) GetResourceClasses ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfv3InterAreaPrefixRoutePool GetResourceClasses
 return "", nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) GetConnectedRouter ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfv3InterAreaPrefixRoutePool GetConnectedRouter
 return "", nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) GetPrefixOptions ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfv3InterAreaPrefixRoutePool GetPrefixOptions
 return "", nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) SetPrefixOptions () error {
 //parameters: LsaHandle PrefixOptions
 //AgtOspfv3InterAreaPrefixRoutePool SetPrefixOptions, m.Object, m.Name);
 return nil
}

func(np *Ospfv3InterAreaPrefixRoutePool) GetInterAreaPrefixLsaPool ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfv3InterAreaPrefixRoutePool GetInterAreaPrefixLsaPool
 return "", nil
}

