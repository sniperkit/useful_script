package n2xsdk

type CaptureControl struct {
 Handler string
}

func(np *CaptureControl) SetPortGroup () error {
 //parameters: NumPorts psaPortHandles
 //AgtCaptureControl SetPortGroup, m.Object, m.Name);
 return nil
}

func(np *CaptureControl) ListPortGroup ()(string, error) {
 //parameters: 
 //AgtCaptureControl ListPortGroup
 return "", nil
}

func(np *CaptureControl) SetCaptureMode () error {
 //parameters: CaptureMode
 //AgtCaptureControl SetCaptureMode, m.Object, m.Name);
 return nil
}

func(np *CaptureControl) GetCaptureMode ()(string, error) {
 //parameters: 
 //AgtCaptureControl GetCaptureMode
 return "", nil
}

func(np *CaptureControl) StartCapture () error {
 //parameters: 
 //AgtCaptureControl StartCapture, m.Object, m.Name);
 return nil
}

func(np *CaptureControl) StopCapture () error {
 //parameters: 
 //AgtCaptureControl StopCapture, m.Object, m.Name);
 return nil
}

func(np *CaptureControl) ArmCaptureStart () error {
 //parameters: 
 //AgtCaptureControl ArmCaptureStart, m.Object, m.Name);
 return nil
}

func(np *CaptureControl) GetCaptureState ()(string, error) {
 //parameters: 
 //AgtCaptureControl GetCaptureState
 return "", nil
}

func(np *CaptureControl) GetCaptureTimes ()(string, error) {
 //parameters: PortHandle
 //AgtCaptureControl GetCaptureTimes
 return "", nil
}

func(np *CaptureControl) GetCaptureCount ()(string, error) {
 //parameters: PortHandle
 //AgtCaptureControl GetCaptureCount
 return "", nil
}

func(np *CaptureControl) SetCaptureBufferSize () error {
 //parameters: PortHandle MBytes
 //AgtCaptureControl SetCaptureBufferSize, m.Object, m.Name);
 return nil
}

func(np *CaptureControl) GetCaptureBufferSize ()(string, error) {
 //parameters: PortHandle
 //AgtCaptureControl GetCaptureBufferSize
 return "", nil
}

func(np *CaptureControl) GetMaximumCaptureBufferSize ()(string, error) {
 //parameters: PortHandle
 //AgtCaptureControl GetMaximumCaptureBufferSize
 return "", nil
}

func(np *CaptureControl) SetCaptureCenteringDepth () error {
 //parameters: PortHandle Octets
 //AgtCaptureControl SetCaptureCenteringDepth, m.Object, m.Name);
 return nil
}

func(np *CaptureControl) GetCaptureCenteringDepth ()(string, error) {
 //parameters: PortHandle
 //AgtCaptureControl GetCaptureCenteringDepth
 return "", nil
}

