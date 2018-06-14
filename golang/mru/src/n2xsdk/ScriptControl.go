package n2xsdk

type ScriptControl struct {
 Handler string
}

func(np *ScriptControl) Call () error {
 //parameters: Command
 //AgtScriptControl Call, m.Object, m.Name);
 return nil
}

func(np *ScriptControl) CreateCommandWrappers () error {
 //parameters: WrapperTemplate
 //AgtScriptControl CreateCommandWrappers, m.Object, m.Name);
 return nil
}

