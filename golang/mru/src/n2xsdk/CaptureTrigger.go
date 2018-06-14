package n2xsdk

type CaptureTrigger struct {
 Handler string
}

func(np *CaptureTrigger) ClearAllTriggers () error {
 //parameters: PortHandle NumActions psaTriggerActions
 //AgtCaptureTrigger ClearAllTriggers, m.Object, m.Name);
 return nil
}

func(np *CaptureTrigger) AddFrameMatcherTriggers () error {
 //parameters: PortHandle FrameMatcherHandleCount psaFrameMatcherHandles TriggerAction
 //AgtCaptureTrigger AddFrameMatcherTriggers, m.Object, m.Name);
 return nil
}

func(np *CaptureTrigger) RemoveFrameMatcherTriggers () error {
 //parameters: PortHandle FrameMatcherHandleCount psaFrameMatcherHandles TriggerAction
 //AgtCaptureTrigger RemoveFrameMatcherTriggers, m.Object, m.Name);
 return nil
}

func(np *CaptureTrigger) ListFrameMatcherTriggers ()(string, error) {
 //parameters: PortHandle TriggerAction
 //AgtCaptureTrigger ListFrameMatcherTriggers
 return "", nil
}

func(np *CaptureTrigger) AddStatusTriggers () error {
 //parameters: PortHandle StatusCount psaStatusTriggers TriggerAction
 //AgtCaptureTrigger AddStatusTriggers, m.Object, m.Name);
 return nil
}

func(np *CaptureTrigger) RemoveStatusTriggers () error {
 //parameters: PortHandle StatusCount psaStatusTriggers TriggerAction
 //AgtCaptureTrigger RemoveStatusTriggers, m.Object, m.Name);
 return nil
}

func(np *CaptureTrigger) ListStatusTriggers ()(string, error) {
 //parameters: PortHandle TriggerAction
 //AgtCaptureTrigger ListStatusTriggers
 return "", nil
}

