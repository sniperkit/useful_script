package n2xsdk

type ncpServerOam struct {
 Handler string
}

func(np *ncpServerOam) SetLoopbackCount () error {
 //parameters: SessionPoolHandle LoopbackCount
 //AgtAncpServerOam SetLoopbackCount, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) GetLoopbackCount ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtAncpServerOam GetLoopbackCount
 return "", nil
}

func(np *ncpServerOam) SetLoopbackTimeout () error {
 //parameters: SessionPoolHandle LoopbackTimeout
 //AgtAncpServerOam SetLoopbackTimeout, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) GetLoopbackTimeout ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtAncpServerOam GetLoopbackTimeout
 return "", nil
}

func(np *ncpServerOam) SetOpaqueData () error {
 //parameters: SessionPoolHandle FirstToken SecondToken
 //AgtAncpServerOam SetOpaqueData, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) GetOpaqueData ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtAncpServerOam GetOpaqueData
 return "", nil
}

func(np *ncpServerOam) AddTlv () error {
 //parameters: SessionPoolHandle TlvType TlvLength TlvValue
 //AgtAncpServerOam AddTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) ListTlvs ()(string, error) {
 //parameters: SessionPoolHandle
 //AgtAncpServerOam ListTlvs
 return "", nil
}

func(np *ncpServerOam) RemoveTlv () error {
 //parameters: SessionPoolHandle TlvType
 //AgtAncpServerOam RemoveTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) RemoveAllTlvs () error {
 //parameters: SessionPoolHandle
 //AgtAncpServerOam RemoveAllTlvs, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) SetTlv () error {
 //parameters: SessionPoolHandle TlvType TlvLength TlvValue
 //AgtAncpServerOam SetTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) GetTlv ()(string, error) {
 //parameters: SessionPoolHandle TlvType
 //AgtAncpServerOam GetTlv
 return "", nil
}

func(np *ncpServerOam) AddSubTlv () error {
 //parameters: SessionPoolHandle TlvType SubTlvType SubTlvLength SubTlvValue
 //AgtAncpServerOam AddSubTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) ListSubTlvs ()(string, error) {
 //parameters: SessionPoolHandle TlvType
 //AgtAncpServerOam ListSubTlvs
 return "", nil
}

func(np *ncpServerOam) RemoveSubTlv () error {
 //parameters: SessionPoolHandle TlvType SubTlvType
 //AgtAncpServerOam RemoveSubTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) RemoveAllSubTlvs () error {
 //parameters: SessionPoolHandle TlvType
 //AgtAncpServerOam RemoveAllSubTlvs, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) SetSubTlv () error {
 //parameters: SessionPoolHandle TlvType SubTlvType SubTlvLength SubTlvValue
 //AgtAncpServerOam SetSubTlv, m.Object, m.Name);
 return nil
}

func(np *ncpServerOam) GetSubTlv ()(string, error) {
 //parameters: SessionPoolHandle TlvType SubTlvType
 //AgtAncpServerOam GetSubTlv
 return "", nil
}

