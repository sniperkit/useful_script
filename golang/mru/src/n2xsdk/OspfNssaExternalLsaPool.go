package n2xsdk

type OspfNssaExternalLsaPool struct {
 Handler string
}

func(np *OspfNssaExternalLsaPool) SetAdvertisingRouter () error {
 //parameters: LsaHandle RouterId
 //AgtOspfNssaExternalLsaPool SetAdvertisingRouter, m.Object, m.Name);
 return nil
}

func(np *OspfNssaExternalLsaPool) GetAdvertisingRouter ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfNssaExternalLsaPool GetAdvertisingRouter
 return "", nil
}

func(np *OspfNssaExternalLsaPool) Advertise () error {
 //parameters: LsaHandle
 //AgtOspfNssaExternalLsaPool Advertise, m.Object, m.Name);
 return nil
}

func(np *OspfNssaExternalLsaPool) Withdraw () error {
 //parameters: LsaHandle
 //AgtOspfNssaExternalLsaPool Withdraw, m.Object, m.Name);
 return nil
}

func(np *OspfNssaExternalLsaPool) GetAdvertiseFlag ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfNssaExternalLsaPool GetAdvertiseFlag
 return "", nil
}

func(np *OspfNssaExternalLsaPool) SetRoutes () error {
 //parameters: LsaHandle FirstRoute PrefixLength NumRoutes Modifier
 //AgtOspfNssaExternalLsaPool SetRoutes, m.Object, m.Name);
 return nil
}

func(np *OspfNssaExternalLsaPool) GetRoutes ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfNssaExternalLsaPool GetRoutes
 return "", nil
}

func(np *OspfNssaExternalLsaPool) SetMetric () error {
 //parameters: LsaHandle Metric
 //AgtOspfNssaExternalLsaPool SetMetric, m.Object, m.Name);
 return nil
}

func(np *OspfNssaExternalLsaPool) GetMetric ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfNssaExternalLsaPool GetMetric
 return "", nil
}

func(np *OspfNssaExternalLsaPool) SetEbit () error {
 //parameters: hLsaHandle Flag
 //AgtOspfNssaExternalLsaPool SetEbit, m.Object, m.Name);
 return nil
}

func(np *OspfNssaExternalLsaPool) GetEbit ()(string, error) {
 //parameters: hLsaHandle
 //AgtOspfNssaExternalLsaPool GetEbit
 return "", nil
}

func(np *OspfNssaExternalLsaPool) SetForwardingAddress () error {
 //parameters: LsaHandle IpAddress
 //AgtOspfNssaExternalLsaPool SetForwardingAddress, m.Object, m.Name);
 return nil
}

func(np *OspfNssaExternalLsaPool) GetForwardingAddress ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfNssaExternalLsaPool GetForwardingAddress
 return "", nil
}

func(np *OspfNssaExternalLsaPool) SetExternalRouteTag () error {
 //parameters: LsaHandle RouteTag
 //AgtOspfNssaExternalLsaPool SetExternalRouteTag, m.Object, m.Name);
 return nil
}

func(np *OspfNssaExternalLsaPool) GetExternalRouteTag ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfNssaExternalLsaPool GetExternalRouteTag
 return "", nil
}

func(np *OspfNssaExternalLsaPool) SetPbit () error {
 //parameters: hLsaHandle PbitFlag
 //AgtOspfNssaExternalLsaPool SetPbit, m.Object, m.Name);
 return nil
}

func(np *OspfNssaExternalLsaPool) GetPbit ()(string, error) {
 //parameters: hLsaHandle
 //AgtOspfNssaExternalLsaPool GetPbit
 return "", nil
}

