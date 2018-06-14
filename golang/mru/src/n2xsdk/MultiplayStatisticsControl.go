package n2xsdk

type MultiplayStatisticsControl struct {
 Handler string
}

func(np *MultiplayStatisticsControl) Enable () error {
 //parameters: NumPorts pPortHandles
 //AgtMultiplayStatisticsControl Enable, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) Disable () error {
 //parameters: NumPorts pPortHandles
 //AgtMultiplayStatisticsControl Disable, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) IsEnabled () error {
 //parameters: PortHandle
 //AgtMultiplayStatisticsControl IsEnabled, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) IsMultiplayStatisticsSupported () error {
 //parameters: PortHandle
 //AgtMultiplayStatisticsControl IsMultiplayStatisticsSupported, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) EnableDualMode () error {
 //parameters: NumPorts pPortHandles
 //AgtMultiplayStatisticsControl EnableDualMode, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) DisableDualMode () error {
 //parameters: NumPorts pPortHandles
 //AgtMultiplayStatisticsControl DisableDualMode, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) IsDualModeEnabled () error {
 //parameters: PortHandle
 //AgtMultiplayStatisticsControl IsDualModeEnabled, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) SetFieldConfiguration () error {
 //parameters: NumPorts pPortHandles FieldConfig
 //AgtMultiplayStatisticsControl SetFieldConfiguration, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) GetFieldConfiguration ()(string, error) {
 //parameters: NumPorts pPortHandles
 //AgtMultiplayStatisticsControl GetFieldConfiguration
 return "", nil
}

func(np *MultiplayStatisticsControl) ListAvailableFields ()(string, error) {
 //parameters: PortHandle
 //AgtMultiplayStatisticsControl ListAvailableFields
 return "", nil
}

func(np *MultiplayStatisticsControl) GetMaximumNumberOfGroups ()(string, error) {
 //parameters: PortHandle
 //AgtMultiplayStatisticsControl GetMaximumNumberOfGroups
 return "", nil
}

func(np *MultiplayStatisticsControl) SetGroupBaseAddress () error {
 //parameters: PortHandle AddressField GroupAddress
 //AgtMultiplayStatisticsControl SetGroupBaseAddress, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) GetGroupAddressRange ()(string, error) {
 //parameters: PortHandle
 //AgtMultiplayStatisticsControl GetGroupAddressRange
 return "", nil
}

func(np *MultiplayStatisticsControl) SetInitialVlanValue () error {
 //parameters: PortHandle VlanField VlanId
 //AgtMultiplayStatisticsControl SetInitialVlanValue, m.Object, m.Name);
 return nil
}

func(np *MultiplayStatisticsControl) GetVlanRange ()(string, error) {
 //parameters: PortHandle
 //AgtMultiplayStatisticsControl GetVlanRange
 return "", nil
}

