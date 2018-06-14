package n2xsdk

type LdpvplsBgpAdVrfPool struct {
 Handler string
}

func(np *LdpvplsBgpAdVrfPool) QueryLdpVplsBgpAdVrfEntry () error {
 //parameters: VrfPoolHandle VrfIndex RTType RouteTarget
 //AgtLdpvplsBgpAdVrfPool QueryLdpVplsBgpAdVrfEntry, m.Object, m.Name);
 return nil
}

func(np *LdpvplsBgpAdVrfPool) Enable () error {
 //parameters: VrfPoolHandle
 //AgtLdpvplsBgpAdVrfPool Enable, m.Object, m.Name);
 return nil
}

func(np *LdpvplsBgpAdVrfPool) Disable () error {
 //parameters: VrfPoolHandle
 //AgtLdpvplsBgpAdVrfPool Disable, m.Object, m.Name);
 return nil
}

func(np *LdpvplsBgpAdVrfPool) IsEnabled () error {
 //parameters: VrfPoolHandle
 //AgtLdpvplsBgpAdVrfPool IsEnabled, m.Object, m.Name);
 return nil
}

func(np *LdpvplsBgpAdVrfPool) SetImportRouteTargetIncrementingRange () error {
 //parameters: VrfPoolHandle RouteTargetType FirstRouteTarget RouteTargetIncrement RouteTargetRepeat PercentOverlap
 //AgtLdpvplsBgpAdVrfPool SetImportRouteTargetIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *LdpvplsBgpAdVrfPool) GetImportRouteTargetIncrementingRange ()(string, error) {
 //parameters: VrfPoolHandle RouteTargetType
 //AgtLdpvplsBgpAdVrfPool GetImportRouteTargetIncrementingRange
 return "", nil
}

func(np *LdpvplsBgpAdVrfPool) SetImportRouteTargetType () error {
 //parameters: VrfPoolHandle RouteTargetTypeValue
 //AgtLdpvplsBgpAdVrfPool SetImportRouteTargetType, m.Object, m.Name);
 return nil
}

func(np *LdpvplsBgpAdVrfPool) GetImportRouteTargetType ()(string, error) {
 //parameters: VrfPoolHandle
 //AgtLdpvplsBgpAdVrfPool GetImportRouteTargetType
 return "", nil
}

func(np *LdpvplsBgpAdVrfPool) SetVpnsPerPeer () error {
 //parameters: VrfPoolHandle NumVpnsPerPeer
 //AgtLdpvplsBgpAdVrfPool SetVpnsPerPeer, m.Object, m.Name);
 return nil
}

func(np *LdpvplsBgpAdVrfPool) GetVpnsPerPeer ()(string, error) {
 //parameters: VrfPoolHandle
 //AgtLdpvplsBgpAdVrfPool GetVpnsPerPeer
 return "", nil
}

func(np *LdpvplsBgpAdVrfPool) GetTotalNumberOfVpns ()(string, error) {
 //parameters: VrfPoolHandle
 //AgtLdpvplsBgpAdVrfPool GetTotalNumberOfVpns
 return "", nil
}

