package n2xsdk

type IcmpEmulation struct {
 Handler string
}

func(np *IcmpEmulation) Enable () error {
 //parameters: PortHandle
 //AgtIcmpEmulation Enable, m.Object, m.Name);
 return nil
}

func(np *IcmpEmulation) Disable () error {
 //parameters: PortHandle
 //AgtIcmpEmulation Disable, m.Object, m.Name);
 return nil
}

func(np *IcmpEmulation) IsEnabled () error {
 //parameters: PortHandle
 //AgtIcmpEmulation IsEnabled, m.Object, m.Name);
 return nil
}

func(np *IcmpEmulation) IsIcmpSupported () error {
 //parameters: PortHandle
 //AgtIcmpEmulation IsIcmpSupported, m.Object, m.Name);
 return nil
}

