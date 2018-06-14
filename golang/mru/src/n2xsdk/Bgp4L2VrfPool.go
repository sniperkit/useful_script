package n2xsdk

type Bgp4L2VrfPool struct {
 Handler string
}

func(np *Bgp4L2VrfPool) QueryL2VrfEntry () error {
 //parameters: VrfPoolHandle VrfIndex RTType RouteTarget
 //AgtBgp4L2VrfPool QueryL2VrfEntry, m.Object, m.Name);
 return nil
}

func(np *Bgp4L2VrfPool) QueryL2VrfEntryByVeId () error {
 //parameters: VrfPoolHandle VrfIndex RTType RouteTarget VeId
 //AgtBgp4L2VrfPool QueryL2VrfEntryByVeId, m.Object, m.Name);
 return nil
}

func(np *Bgp4L2VrfPool) SetEncapsulationType () error {
 //parameters: VrfPoolHandle EncapsulationType
 //AgtBgp4L2VrfPool SetEncapsulationType, m.Object, m.Name);
 return nil
}

func(np *Bgp4L2VrfPool) GetEncapsulationType ()(string, error) {
 //parameters: VrfPoolHandle
 //AgtBgp4L2VrfPool GetEncapsulationType
 return "", nil
}

func(np *Bgp4L2VrfPool) Enable () error {
 //parameters: VrfPoolHandle
 //AgtBgp4L2VrfPool Enable, m.Object, m.Name);
 return nil
}

func(np *Bgp4L2VrfPool) Disable () error {
 //parameters: VrfPoolHandle
 //AgtBgp4L2VrfPool Disable, m.Object, m.Name);
 return nil
}

func(np *Bgp4L2VrfPool) IsEnabled () error {
 //parameters: VrfPoolHandle
 //AgtBgp4L2VrfPool IsEnabled, m.Object, m.Name);
 return nil
}

func(np *Bgp4L2VrfPool) SetImportRouteTargetIncrementingRange () error {
 //parameters: VrfPoolHandle RouteTargetType FirstRouteTarget RouteTargetIncrement RouteTargetRepeat PercentOverlap
 //AgtBgp4L2VrfPool SetImportRouteTargetIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *Bgp4L2VrfPool) GetImportRouteTargetIncrementingRange ()(string, error) {
 //parameters: VrfPoolHandle RouteTargetType
 //AgtBgp4L2VrfPool GetImportRouteTargetIncrementingRange
 return "", nil
}

func(np *Bgp4L2VrfPool) SetImportRouteTargetType () error {
 //parameters: VrfPoolHandle RouteTargetTypeValue
 //AgtBgp4L2VrfPool SetImportRouteTargetType, m.Object, m.Name);
 return nil
}

func(np *Bgp4L2VrfPool) GetImportRouteTargetType ()(string, error) {
 //parameters: VrfPoolHandle
 //AgtBgp4L2VrfPool GetImportRouteTargetType
 return "", nil
}

func(np *Bgp4L2VrfPool) SetVpnsPerPeer () error {
 //parameters: VrfPoolHandle NumVpnsPerPeer
 //AgtBgp4L2VrfPool SetVpnsPerPeer, m.Object, m.Name);
 return nil
}

func(np *Bgp4L2VrfPool) GetVpnsPerPeer ()(string, error) {
 //parameters: VrfPoolHandle
 //AgtBgp4L2VrfPool GetVpnsPerPeer
 return "", nil
}

func(np *Bgp4L2VrfPool) GetTotalNumberOfVpns ()(string, error) {
 //parameters: VrfPoolHandle
 //AgtBgp4L2VrfPool GetTotalNumberOfVpns
 return "", nil
}

