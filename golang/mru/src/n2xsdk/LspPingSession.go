package n2xsdk

type LspPingSession struct {
 Handler string
}

func(np *LspPingSession) Start () error {
 //parameters: SessionHandle
 //AgtLspPingSession Start, m.Object, m.Name);
 return nil
}

func(np *LspPingSession) Reset () error {
 //parameters: SessionHandle
 //AgtLspPingSession Reset, m.Object, m.Name);
 return nil
}

func(np *LspPingSession) Renew () error {
 //parameters: SessionHandle
 //AgtLspPingSession Renew, m.Object, m.Name);
 return nil
}

