package n2xsdk

type PppEmulation struct {
 Handler string
}

func(np *PppEmulation) Enable () error {
 //parameters: PortHandle
 //AgtPppEmulation Enable, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) Disable () error {
 //parameters: PortHandle
 //AgtPppEmulation Disable, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) EnableNcp () error {
 //parameters: PortHandle NcpProtocol
 //AgtPppEmulation EnableNcp, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) DisableNcp () error {
 //parameters: PortHandle NcpProtocol
 //AgtPppEmulation DisableNcp, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) Open () error {
 //parameters: PortHandle
 //AgtPppEmulation Open, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) Close () error {
 //parameters: PortHandle
 //AgtPppEmulation Close, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) OpenNcp () error {
 //parameters: PortHandle NcpProtocol
 //AgtPppEmulation OpenNcp, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) CloseNcp () error {
 //parameters: PortHandle NcpProtocol
 //AgtPppEmulation CloseNcp, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) GetState ()(string, error) {
 //parameters: PortHandle
 //AgtPppEmulation GetState
 return "", nil
}

func(np *PppEmulation) GetNcpState ()(string, error) {
 //parameters: PortHandle NcpProtocol
 //AgtPppEmulation GetNcpState
 return "", nil
}

func(np *PppEmulation) GetLastError ()(string, error) {
 //parameters: PortHandle
 //AgtPppEmulation GetLastError
 return "", nil
}

func(np *PppEmulation) IsPppSupported () error {
 //parameters: PortHandle
 //AgtPppEmulation IsPppSupported, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) SetDefaultLcpConfigurationOption () error {
 //parameters: PortHandle LcpOption DefaultValue
 //AgtPppEmulation SetDefaultLcpConfigurationOption, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) GetDefaultLcpConfigurationOption ()(string, error) {
 //parameters: PortHandle LcpOption
 //AgtPppEmulation GetDefaultLcpConfigurationOption
 return "", nil
}

func(np *PppEmulation) SetLcpConfigurationOption () error {
 //parameters: PortHandle LcpOption DesiredValue
 //AgtPppEmulation SetLcpConfigurationOption, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) GetLcpConfigurationOption ()(string, error) {
 //parameters: PortHandle LcpOption
 //AgtPppEmulation GetLcpConfigurationOption
 return "", nil
}

func(np *PppEmulation) GetLcpSupportedConfigurationOptions ()(string, error) {
 //parameters: PortHandle LcpOption
 //AgtPppEmulation GetLcpSupportedConfigurationOptions
 return "", nil
}

func(np *PppEmulation) GetLocalLcpConfigurationValue ()(string, error) {
 //parameters: PortHandle LcpOption
 //AgtPppEmulation GetLocalLcpConfigurationValue
 return "", nil
}

func(np *PppEmulation) GetRemoteLcpConfigurationValue ()(string, error) {
 //parameters: PortHandle LcpOption
 //AgtPppEmulation GetRemoteLcpConfigurationValue
 return "", nil
}

func(np *PppEmulation) SetEmulationParameter () error {
 //parameters: PortHandle PppParameter Value
 //AgtPppEmulation SetEmulationParameter, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) GetEmulationParameter ()(string, error) {
 //parameters: PortHandle PppParameter
 //AgtPppEmulation GetEmulationParameter
 return "", nil
}

func(np *PppEmulation) GetRemoteIpcpConfigurationValue ()(string, error) {
 //parameters: PortHandle IpcpOption
 //AgtPppEmulation GetRemoteIpcpConfigurationValue
 return "", nil
}

func(np *PppEmulation) SetDefaultNcpConfigurationOption () error {
 //parameters: PortHandle NcpProtocol NcpOption DefaultValue
 //AgtPppEmulation SetDefaultNcpConfigurationOption, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) GetDefaultNcpConfigurationOption ()(string, error) {
 //parameters: PortHandle NcpProtocol NcpOption
 //AgtPppEmulation GetDefaultNcpConfigurationOption
 return "", nil
}

func(np *PppEmulation) SetNcpConfigurationOption () error {
 //parameters: PortHandle NcpProtocol NcpOption DesiredValue
 //AgtPppEmulation SetNcpConfigurationOption, m.Object, m.Name);
 return nil
}

func(np *PppEmulation) GetNcpConfigurationOption ()(string, error) {
 //parameters: PortHandle NcpProtocol NcpOption
 //AgtPppEmulation GetNcpConfigurationOption
 return "", nil
}

func(np *PppEmulation) GetLocalNcpConfigurationValue ()(string, error) {
 //parameters: PortHandle NcpProtocol NcpOption
 //AgtPppEmulation GetLocalNcpConfigurationValue
 return "", nil
}

func(np *PppEmulation) GetRemoteNcpConfigurationValue ()(string, error) {
 //parameters: PortHandle NcpProtocol NcpOption
 //AgtPppEmulation GetRemoteNcpConfigurationValue
 return "", nil
}

func(np *PppEmulation) GetListOfRejectedProtocols ()(string, error) {
 //parameters: PortHandle
 //AgtPppEmulation GetListOfRejectedProtocols
 return "", nil
}

