package n2xsdk

type OspfSummaryLsaPool struct {
 Handler string
}

func(np *OspfSummaryLsaPool) SetAdvertisingRouter () error {
 //parameters: LsaHandle RouterId
 //AgtOspfSummaryLsaPool SetAdvertisingRouter, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryLsaPool) GetAdvertisingRouter ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfSummaryLsaPool GetAdvertisingRouter
 return "", nil
}

func(np *OspfSummaryLsaPool) Advertise () error {
 //parameters: LsaHandle
 //AgtOspfSummaryLsaPool Advertise, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryLsaPool) Withdraw () error {
 //parameters: LsaHandle
 //AgtOspfSummaryLsaPool Withdraw, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryLsaPool) GetAdvertiseFlag ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfSummaryLsaPool GetAdvertiseFlag
 return "", nil
}

func(np *OspfSummaryLsaPool) SetRoutes () error {
 //parameters: LsaHandle FirstRoute PrefixLength NumRoutes Modifier
 //AgtOspfSummaryLsaPool SetRoutes, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryLsaPool) GetRoutes ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfSummaryLsaPool GetRoutes
 return "", nil
}

func(np *OspfSummaryLsaPool) SetMetric () error {
 //parameters: LsaHandle Metric
 //AgtOspfSummaryLsaPool SetMetric, m.Object, m.Name);
 return nil
}

func(np *OspfSummaryLsaPool) GetMetric ()(string, error) {
 //parameters: LsaHandle
 //AgtOspfSummaryLsaPool GetMetric
 return "", nil
}

