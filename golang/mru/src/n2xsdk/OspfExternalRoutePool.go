package n2xsdk

type OspfExternalRoutePool struct {
 Handler string
}

func(np *OspfExternalRoutePool) SetRoutes () error {
 //parameters: RoutePoolHandle FirstRoute PrefixLength NumRoutes Modifier
 //AgtOspfExternalRoutePool SetRoutes, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetRoutes ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetRoutes
 return "", nil
}

func(np *OspfExternalRoutePool) EnableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool EnableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) DisableTrafficDestinations () error {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool DisableTrafficDestinations, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) IsTrafficDestinationEnabled () error {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool IsTrafficDestinationEnabled, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) SetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtOspfExternalRoutePool SetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) UnsetRoutePoolFlag () error {
 //parameters: RoutePoolHandle Flag
 //AgtOspfExternalRoutePool UnsetRoutePoolFlag, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) IsRoutePoolFlagSet () error {
 //parameters: RoutePoolHandle Flag
 //AgtOspfExternalRoutePool IsRoutePoolFlagSet, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetRoutePoolStatus ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetRoutePoolStatus
 return "", nil
}

func(np *OspfExternalRoutePool) GetNthRoute ()(string, error) {
 //parameters: RoutePoolHandle Index
 //AgtOspfExternalRoutePool GetNthRoute
 return "", nil
}

func(np *OspfExternalRoutePool) SetResourceClasses () error {
 //parameters: RoutePoolHandle Count psaClasses
 //AgtOspfExternalRoutePool SetResourceClasses, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetResourceClasses ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetResourceClasses
 return "", nil
}

func(np *OspfExternalRoutePool) GetConnectedRouter ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetConnectedRouter
 return "", nil
}

func(np *OspfExternalRoutePool) SetNssaFlag () error {
 //parameters: RoutePoolHandle NssaFlag
 //AgtOspfExternalRoutePool SetNssaFlag, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetNssaFlag ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetNssaFlag
 return "", nil
}

func(np *OspfExternalRoutePool) SetPropagateLsasFlag () error {
 //parameters: RoutePoolHandle PropagateFlag
 //AgtOspfExternalRoutePool SetPropagateLsasFlag, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetPropagateLsasFlag ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetPropagateLsasFlag
 return "", nil
}

func(np *OspfExternalRoutePool) SetAutoGenerateAsbrSummaryLsaFlag () error {
 //parameters: RoutePoolHandle AsbrFlag
 //AgtOspfExternalRoutePool SetAutoGenerateAsbrSummaryLsaFlag, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetAutoGenerateAsbrSummaryLsaFlag ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetAutoGenerateAsbrSummaryLsaFlag
 return "", nil
}

func(np *OspfExternalRoutePool) SetAdvertisingRouter () error {
 //parameters: RoutePoolHandle RouterId
 //AgtOspfExternalRoutePool SetAdvertisingRouter, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetAdvertisingRouter ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetAdvertisingRouter
 return "", nil
}

func(np *OspfExternalRoutePool) SetMetricType () error {
 //parameters: RoutePoolHandle Type
 //AgtOspfExternalRoutePool SetMetricType, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetMetricType ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetMetricType
 return "", nil
}

func(np *OspfExternalRoutePool) SetMetric () error {
 //parameters: RoutePoolHandle Metric
 //AgtOspfExternalRoutePool SetMetric, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetMetric ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetMetric
 return "", nil
}

func(np *OspfExternalRoutePool) SetForwardingAddress () error {
 //parameters: RoutePoolHandle IpAddress
 //AgtOspfExternalRoutePool SetForwardingAddress, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetForwardingAddress ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetForwardingAddress
 return "", nil
}

func(np *OspfExternalRoutePool) SetExternalRouteTag () error {
 //parameters: RoutePoolHandle RouteTag
 //AgtOspfExternalRoutePool SetExternalRouteTag, m.Object, m.Name);
 return nil
}

func(np *OspfExternalRoutePool) GetExternalRouteTag ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetExternalRouteTag
 return "", nil
}

func(np *OspfExternalRoutePool) GetExternalLsaPool ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetExternalLsaPool
 return "", nil
}

func(np *OspfExternalRoutePool) GetAsbrSummaryLsa ()(string, error) {
 //parameters: RoutePoolHandle
 //AgtOspfExternalRoutePool GetAsbrSummaryLsa
 return "", nil
}

