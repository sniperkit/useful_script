package n2xsdk

type HdlcInterface struct {
 Handler string
}

func(np *HdlcInterface) SetMtu () error {
 //parameters: PortHandle Mtu
 //AgtHdlcInterface SetMtu, m.Object, m.Name);
 return nil
}

func(np *HdlcInterface) GetMtu ()(string, error) {
 //parameters: PortHandle
 //AgtHdlcInterface GetMtu
 return "", nil
}

