package n2xsdk

type Gateway struct {
 Handler string
}

func(np *Gateway) ListEmulations ()(string, error) {
 //parameters: DeviceHandle
 //AgtGateway ListEmulations
 return "", nil
}

func(np *Gateway) ListSupportedEmulationsTypes ()(string, error) {
 //parameters: DeviceHandle
 //AgtGateway ListSupportedEmulationsTypes
 return "", nil
}

func(np *Gateway) ListEmulationsByType ()(string, error) {
 //parameters: DeviceHandle EmulationType
 //AgtGateway ListEmulationsByType
 return "", nil
}

func(np *Gateway) AddEmulation () error {
 //parameters: DeviceHandle EmulationType
 //AgtGateway AddEmulation, m.Object, m.Name);
 return nil
}

func(np *Gateway) RemoveEmulation () error {
 //parameters: EmulationHandles
 //AgtGateway RemoveEmulation, m.Object, m.Name);
 return nil
}

func(np *Gateway) RemoveEmulationsByType () error {
 //parameters: DeviceHandle EmulationType
 //AgtGateway RemoveEmulationsByType, m.Object, m.Name);
 return nil
}

