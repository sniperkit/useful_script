package n2xsdk

type OpticalInterface struct {
 Handler string
}

func(np *OpticalInterface) ListAvailablePortModes ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface ListAvailablePortModes
 return "", nil
}

func(np *OpticalInterface) ListAvailablePortModesByMediaType ()(string, error) {
 //parameters: PortHandle MediaType
 //AgtOpticalInterface ListAvailablePortModesByMediaType
 return "", nil
}

func(np *OpticalInterface) SetPortMode () error {
 //parameters: PortHandle PortMode
 //AgtOpticalInterface SetPortMode, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) GetPortMode ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface GetPortMode
 return "", nil
}

func(np *OpticalInterface) ListAvailableClockSources ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface ListAvailableClockSources
 return "", nil
}

func(np *OpticalInterface) SetClockSource () error {
 //parameters: PortHandle ClockSource
 //AgtOpticalInterface SetClockSource, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) GetClockSource ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface GetClockSource
 return "", nil
}

func(np *OpticalInterface) GetClockState ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface GetClockState
 return "", nil
}

func(np *OpticalInterface) IsPluginMediaSupported () error {
 //parameters: PortHandle
 //AgtOpticalInterface IsPluginMediaSupported, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) IsPluginMediaInserted () error {
 //parameters: PortHandle
 //AgtOpticalInterface IsPluginMediaInserted, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) GetPluginMediaType ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface GetPluginMediaType
 return "", nil
}

func(np *OpticalInterface) ListAvailableMediaTypes ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface ListAvailableMediaTypes
 return "", nil
}

func(np *OpticalInterface) SetMediaType () error {
 //parameters: PortHandle MediaType
 //AgtOpticalInterface SetMediaType, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) GetMediaType ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface GetMediaType
 return "", nil
}

func(np *OpticalInterface) IsMediaTypeSelectionSupported () error {
 //parameters: PortHandle MediaType
 //AgtOpticalInterface IsMediaTypeSelectionSupported, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) ListAvailableExternalClockRates ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface ListAvailableExternalClockRates
 return "", nil
}

func(np *OpticalInterface) SetExternalClockRate () error {
 //parameters: PortHandle Rate
 //AgtOpticalInterface SetExternalClockRate, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) GetExternalClockRate ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface GetExternalClockRate
 return "", nil
}

func(np *OpticalInterface) AllLasersOn () error {
 //parameters: 
 //AgtOpticalInterface AllLasersOn, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) AllLasersOff () error {
 //parameters: 
 //AgtOpticalInterface AllLasersOff, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) LaserOn () error {
 //parameters: PortHandle
 //AgtOpticalInterface LaserOn, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) LaserOff () error {
 //parameters: PortHandle
 //AgtOpticalInterface LaserOff, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) IsLaserOn () error {
 //parameters: PortHandle
 //AgtOpticalInterface IsLaserOn, m.Object, m.Name);
 return nil
}

func(np *OpticalInterface) GetLaserOnTimestamp ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface GetLaserOnTimestamp
 return "", nil
}

func(np *OpticalInterface) GetLaserOffTimestamp ()(string, error) {
 //parameters: PortHandle
 //AgtOpticalInterface GetLaserOffTimestamp
 return "", nil
}

