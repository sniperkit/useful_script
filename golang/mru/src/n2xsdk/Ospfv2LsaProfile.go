package n2xsdk

type Ospfv2LsaProfile struct {
 Handler string
}

func(np *Ospfv2LsaProfile) GetType ()(string, error) {
 //parameters: LsaProfileHandle
 //AgtOspfv2LsaProfile GetType
 return "", nil
}

func(np *Ospfv2LsaProfile) SetName () error {
 //parameters: LsaProfileHandle LsaProfileName
 //AgtOspfv2LsaProfile SetName, m.Object, m.Name);
 return nil
}

func(np *Ospfv2LsaProfile) GetName ()(string, error) {
 //parameters: LsaProfileHandle
 //AgtOspfv2LsaProfile GetName
 return "", nil
}

func(np *Ospfv2LsaProfile) Enable () error {
 //parameters: LsaProfileHandle
 //AgtOspfv2LsaProfile Enable, m.Object, m.Name);
 return nil
}

func(np *Ospfv2LsaProfile) Disable () error {
 //parameters: LsaProfileHandle
 //AgtOspfv2LsaProfile Disable, m.Object, m.Name);
 return nil
}

func(np *Ospfv2LsaProfile) IsEnabled () error {
 //parameters: LsaProfileHandle
 //AgtOspfv2LsaProfile IsEnabled, m.Object, m.Name);
 return nil
}

func(np *Ospfv2LsaProfile) EnableTrafficDestination () error {
 //parameters: LsaProfileHandle
 //AgtOspfv2LsaProfile EnableTrafficDestination, m.Object, m.Name);
 return nil
}

func(np *Ospfv2LsaProfile) DisableTrafficDestination () error {
 //parameters: LsaProfileHandle
 //AgtOspfv2LsaProfile DisableTrafficDestination, m.Object, m.Name);
 return nil
}

func(np *Ospfv2LsaProfile) IsTrafficDestinationsEnabled () error {
 //parameters: LsaProfileHandle
 //AgtOspfv2LsaProfile IsTrafficDestinationsEnabled, m.Object, m.Name);
 return nil
}

func(np *Ospfv2LsaProfile) GetNumberOfRouters ()(string, error) {
 //parameters: LsaProfileHandle
 //AgtOspfv2LsaProfile GetNumberOfRouters
 return "", nil
}

