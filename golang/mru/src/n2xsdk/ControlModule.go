package n2xsdk

type ControlModule struct {
 Handler string
}

func(np *ControlModule) RunScript () error {
 //parameters: filename
 //AgtControlModule RunScript, m.Object, m.Name);
 return nil
}

func(np *ControlModule) SetLogLevel () error {
 //parameters: LogLevel
 //AgtControlModule SetLogLevel, m.Object, m.Name);
 return nil
}

func(np *ControlModule) GetLogLevel ()(string, error) {
 //parameters: 
 //AgtControlModule GetLogLevel
 return "", nil
}

func(np *ControlModule) SetControllerLogFileName () error {
 //parameters: LogName
 //AgtControlModule SetControllerLogFileName, m.Object, m.Name);
 return nil
}

func(np *ControlModule) SetGuiPort () error {
 //parameters: Port
 //AgtControlModule SetGuiPort, m.Object, m.Name);
 return nil
}

func(np *ControlModule) SetGuiPortIpAddress () error {
 //parameters: Port IpAddress
 //AgtControlModule SetGuiPortIpAddress, m.Object, m.Name);
 return nil
}

func(np *ControlModule) EnableSlaveLogging () error {
 //parameters: 
 //AgtControlModule EnableSlaveLogging, m.Object, m.Name);
 return nil
}

func(np *ControlModule) DisableSlaveLogging () error {
 //parameters: 
 //AgtControlModule DisableSlaveLogging, m.Object, m.Name);
 return nil
}

func(np *ControlModule) StopScript () error {
 //parameters: 
 //AgtControlModule StopScript, m.Object, m.Name);
 return nil
}

func(np *ControlModule) GetVersion ()(string, error) {
 //parameters: 
 //AgtControlModule GetVersion
 return "", nil
}

func(np *ControlModule) GetScriptState ()(string, error) {
 //parameters: 
 //AgtControlModule GetScriptState
 return "", nil
}

func(np *ControlModule) IsSlaveLoggingEnabled () error {
 //parameters: 
 //AgtControlModule IsSlaveLoggingEnabled, m.Object, m.Name);
 return nil
}

func(np *ControlModule) SetGlobalScriptVariable () error {
 //parameters: VarName Value
 //AgtControlModule SetGlobalScriptVariable, m.Object, m.Name);
 return nil
}

func(np *ControlModule) SetLocalScriptVariable () error {
 //parameters: ssmId VarName Value
 //AgtControlModule SetLocalScriptVariable, m.Object, m.Name);
 return nil
}

func(np *ControlModule) RunScriptProcedure () error {
 //parameters: ssmId ProcName
 //AgtControlModule RunScriptProcedure, m.Object, m.Name);
 return nil
}

func(np *ControlModule) CompileScript () error {
 //parameters: ScriptName
 //AgtControlModule CompileScript, m.Object, m.Name);
 return nil
}

func(np *ControlModule) CompileSuite () error {
 //parameters: ScriptName
 //AgtControlModule CompileSuite, m.Object, m.Name);
 return nil
}

func(np *ControlModule) GetGlobalVariable ()(string, error) {
 //parameters: VarName
 //AgtControlModule GetGlobalVariable
 return "", nil
}

func(np *ControlModule) GetGlobalArrayMember ()(string, error) {
 //parameters: VarName Index
 //AgtControlModule GetGlobalArrayMember
 return "", nil
}

func(np *ControlModule) SetGlobalVariable () error {
 //parameters: VarName Value
 //AgtControlModule SetGlobalVariable, m.Object, m.Name);
 return nil
}

func(np *ControlModule) SetGlobalArrayMember () error {
 //parameters: VarName Index Value
 //AgtControlModule SetGlobalArrayMember, m.Object, m.Name);
 return nil
}

func(np *ControlModule) SetLocalVariable () error {
 //parameters: ssmId VarName Value
 //AgtControlModule SetLocalVariable, m.Object, m.Name);
 return nil
}

func(np *ControlModule) SetLocalArrayMember () error {
 //parameters: ssmId VarName Index Value
 //AgtControlModule SetLocalArrayMember, m.Object, m.Name);
 return nil
}

func(np *ControlModule) HaltScriptProcedure () error {
 //parameters: ssmId
 //AgtControlModule HaltScriptProcedure, m.Object, m.Name);
 return nil
}

func(np *ControlModule) SetSsmStatsUpdateInterval () error {
 //parameters: ssmId Interval
 //AgtControlModule SetSsmStatsUpdateInterval, m.Object, m.Name);
 return nil
}

func(np *ControlModule) ContinueFromSyncPoint () error {
 //parameters: 
 //AgtControlModule ContinueFromSyncPoint, m.Object, m.Name);
 return nil
}

func(np *ControlModule) OpenRemoteGuiFile () error {
 //parameters: RemoteFile Session
 //AgtControlModule OpenRemoteGuiFile, m.Object, m.Name);
 return nil
}

func(np *ControlModule) WriteRemoteGuiFile () error {
 //parameters: Line
 //AgtControlModule WriteRemoteGuiFile, m.Object, m.Name);
 return nil
}

func(np *ControlModule) CloseRemoteGuiFile () error {
 //parameters: 
 //AgtControlModule CloseRemoteGuiFile, m.Object, m.Name);
 return nil
}

func(np *ControlModule) RunRemoteExpectScript () error {
 //parameters: Session
 //AgtControlModule RunRemoteExpectScript, m.Object, m.Name);
 return nil
}

func(np *ControlModule) GetMctsGuiState ()(string, error) {
 //parameters: 
 //AgtControlModule GetMctsGuiState
 return "", nil
}

func(np *ControlModule) SetMctsGuiRunning () error {
 //parameters: Port IpAddress
 //AgtControlModule SetMctsGuiRunning, m.Object, m.Name);
 return nil
}

func(np *ControlModule) UnSetMctsGuiRunning () error {
 //parameters: 
 //AgtControlModule UnSetMctsGuiRunning, m.Object, m.Name);
 return nil
}

