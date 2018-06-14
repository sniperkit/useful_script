package n2xsdk

type OspfSummaryRoutePool struct {
 Handler string
}

func(np *OspfSummaryRoutePool) SetRoutes () error {
 //parameters: RoutePoolHandle FirstRoute PrefixLength NumRoutes Modifier
 //AgtOspfSummaryRoutePool SetRoutes, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryRoutePool) GetRoutes ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfSummaryRoutePool GetRoutes
 return "", nil
}

func(np *OspfSummaryRoutePool) EnableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtOspfSummaryRoutePool EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryRoutePool) DisableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtOspfSummaryRoutePool DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryRoutePool) IsTrafficDestinationEnabled () error {
 //parameters: RoutePoolHandle
 //AgtOspfSummaryRoutePool IsTrafficDestinationEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryRoutePool) SetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtOspfSummaryRoutePool SetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryRoutePool) UnsetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtOspfSummaryRoutePool UnsetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryRoutePool) IsRoutePoolFlagSet () error {
 //parameters: RoutePoolHandle Flag
 //AgtOspfSummaryRoutePool IsRoutePoolFlagSet, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryRoutePool) GetRoutePoolStatus ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfSummaryRoutePool GetRoutePoolStatus
 return "", nil
}

func(np *OspfSummaryRoutePool) GetNthRoute ()(string, error) {
 //parameters: RoutePoolHandle Index
 //AgtOspfSummaryRoutePool GetNthRoute
 return "", nil
}

func(np *OspfSummaryRoutePool) SetResourceClasses () error {
 //parameters: RoutePoolHandle Count psaClasses
 //AgtOspfSummaryRoutePool SetResourceClasses, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryRoutePool) GetResourceClasses ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfSummaryRoutePool GetResourceClasses
 return "", nil
}

func(np *OspfSummaryRoutePool) GetConnectedRouter ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfSummaryRoutePool GetConnectedRouter
 return "", nil
}

func(np *OspfSummaryRoutePool) GetSummaryLsaPool ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfSummaryRoutePool GetSummaryLsaPool
 return "", nil
}

