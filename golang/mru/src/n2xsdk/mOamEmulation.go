package n2xsdk

type mOamEmulation struct {
 Handler string
}

func(np *mOamEmulation) Enable () error {
 //parameters: PortHandle
 //AgtAtmOamEmulation Enable, m.Object, m.Name);
 return nil
}

func(np *mOamEmulation) Disable () error {
 //parameters: PortHandle
 //AgtAtmOamEmulation Disable, m.Object, m.Name);
 return nil
}

func(np *mOamEmulation) IsEnabled () error {
 //parameters: PortHandle
 //AgtAtmOamEmulation IsEnabled, m.Object, m.Name);
 return nil
}

