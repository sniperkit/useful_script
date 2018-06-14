package n2xsdk

type Ospfv2RouterPool struct {
 Handler string
}

func(np *Ospfv2RouterPool) GetState ()(string, error) {
 //parameters: RouterPoolHandle RouterPoolInstance
 //AgtOspfv2RouterPool GetState
 return "", nil
}

func(np *Ospfv2RouterPool) GetNetworkType ()(string, error) {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool GetNetworkType
 return "", nil
}

func(np *Ospfv2RouterPool) SetRouterId () error {
 //parameters: RouterPoolHandle FirstRouterId Increment Repeat
 //AgtOspfv2RouterPool SetRouterId, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) GetRouterId ()(string, error) {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool GetRouterId
 return "", nil
}

func(np *Ospfv2RouterPool) SetAreaId () error {
 //parameters: RouterPoolHandle FirstAreaId Increment Repeat
 //AgtOspfv2RouterPool SetAreaId, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) GetAreaId ()(string, error) {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool GetAreaId
 return "", nil
}

func(np *Ospfv2RouterPool) EnableOption () error {
 //parameters: RouterPoolHandle OptionsFieldBitType
 //AgtOspfv2RouterPool EnableOption, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) DisableOption () error {
 //parameters: RouterPoolHandle OptionsFieldBitType
 //AgtOspfv2RouterPool DisableOption, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) IsOptionEnabled () error {
 //parameters: RouterPoolHandle OptionsFieldBitType
 //AgtOspfv2RouterPool IsOptionEnabled, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) SetInterval () error {
 //parameters: RouterPoolHandle IntervalType Interval
 //AgtOspfv2RouterPool SetInterval, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) GetInterval ()(string, error) {
 //parameters: RouterPoolHandle IntervalType
 //AgtOspfv2RouterPool GetInterval
 return "", nil
}

func(np *Ospfv2RouterPool) EnableChecksumValidation () error {
 //parameters: RouterPoolHandle ChecksumType
 //AgtOspfv2RouterPool EnableChecksumValidation, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) DisableChecksumValidation () error {
 //parameters: RouterPoolHandle ChecksumType
 //AgtOspfv2RouterPool DisableChecksumValidation, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) IsChecksumValidationEnabled () error {
 //parameters: RouterPoolHandle ChecksumType
 //AgtOspfv2RouterPool IsChecksumValidationEnabled, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) SetInterfaceCost () error {
 //parameters: RouterPoolHandle InterfaceCost
 //AgtOspfv2RouterPool SetInterfaceCost, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) GetInterfaceCost ()(string, error) {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool GetInterfaceCost
 return "", nil
}

func(np *Ospfv2RouterPool) EnableUpdateRateControl () error {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool EnableUpdateRateControl, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) DisableUpdateRateControl () error {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool DisableUpdateRateControl, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) IsUpdateRateControlEnabled () error {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool IsUpdateRateControlEnabled, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) SetInterUpdateDelay () error {
 //parameters: RouterPoolHandle InterUpdateDelay
 //AgtOspfv2RouterPool SetInterUpdateDelay, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) GetInterUpdateDelay ()(string, error) {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool GetInterUpdateDelay
 return "", nil
}

func(np *Ospfv2RouterPool) EnableLsaDiscard () error {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool EnableLsaDiscard, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) DisableLsaDiscard () error {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool DisableLsaDiscard, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) IsLsaDiscardEnabled () error {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool IsLsaDiscardEnabled, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) AddLsaProfile () error {
 //parameters: RouterPoolHandle LsaProfileType
 //AgtOspfv2RouterPool AddLsaProfile, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) RemoveLsaProfile () error {
 //parameters: RouterPoolHandle LsaProfileHandle
 //AgtOspfv2RouterPool RemoveLsaProfile, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) RemoveAllLsaProfiles () error {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool RemoveAllLsaProfiles, m.Object, m.Name);
 return nil
}

func(np *Ospfv2RouterPool) ListLsaProfiles ()(string, error) {
 //parameters: RouterPoolHandle
 //AgtOspfv2RouterPool ListLsaProfiles
 return "", nil
}

func(np *Ospfv2RouterPool) ListLsaProfilesByType ()(string, error) {
 //parameters: RouterPoolHandle LsaProfileType
 //AgtOspfv2RouterPool ListLsaProfilesByType
 return "", nil
}

