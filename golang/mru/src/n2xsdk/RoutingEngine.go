package n2xsdk

type RoutingEngine struct {
 Handler string
}

func(np *RoutingEngine) Start () error {
 //parameters: 
 //AgtRoutingEngine Start, m.Object, m.Name);
 return nil
}

func(np *RoutingEngine) Stop () error {
 //parameters: 
 //AgtRoutingEngine Stop, m.Object, m.Name);
 return nil
}

func(np *RoutingEngine) GetState ()(string, error) {
 //parameters: 
 //AgtRoutingEngine GetState
 return "", nil
}

func(np *RoutingEngine) SetTraceLogLevel () error {
 //parameters: LogLevel
 //AgtRoutingEngine SetTraceLogLevel, m.Object, m.Name);
 return nil
}

func(np *RoutingEngine) GetTraceLogLevel ()(string, error) {
 //parameters: 
 //AgtRoutingEngine GetTraceLogLevel
 return "", nil
}

func(np *RoutingEngine) SetTraceLogFile () error {
 //parameters: FileName
 //AgtRoutingEngine SetTraceLogFile, m.Object, m.Name);
 return nil
}

func(np *RoutingEngine) GetTraceLogFile ()(string, error) {
 //parameters: 
 //AgtRoutingEngine GetTraceLogFile
 return "", nil
}

