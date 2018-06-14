package n2xsdk

type LspTraceSession struct {
 Handler string
}

func(np *LspTraceSession) Start () error {
 //parameters: SessionHandle
 //AgtLspTraceSession Start, m.Object, m.Name);
 return nil
}

func(np *LspTraceSession) Reset () error {
 //parameters: SessionHandle
 //AgtLspTraceSession Reset, m.Object, m.Name);
 return nil
}

func(np *LspTraceSession) Renew () error {
 //parameters: SessionHandle
 //AgtLspTraceSession Renew, m.Object, m.Name);
 return nil
}

