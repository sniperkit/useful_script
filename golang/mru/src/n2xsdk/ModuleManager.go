package n2xsdk

type ModuleManager struct {
 Handler string
}

func(np *ModuleManager) GetSystemState ()(string, error) {
 //parameters: 
 //AgtModuleManager GetSystemState
 return "", nil
}

func(np *ModuleManager) RebootAllModules () error {
 //parameters: 
 //AgtModuleManager RebootAllModules, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) UpdateModules () error {
 //parameters: 
 //AgtModuleManager UpdateModules, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) DisableAutoUpdate () error {
 //parameters: 
 //AgtModuleManager DisableAutoUpdate, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) EnableAutoUpdate () error {
 //parameters: 
 //AgtModuleManager EnableAutoUpdate, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) IsAutoUpdateEnabled () error {
 //parameters: 
 //AgtModuleManager IsAutoUpdateEnabled, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) UseSingleModule () error {
 //parameters: SerialNumber
 //AgtModuleManager UseSingleModule, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) ListModules ()(string, error) {
 //parameters: 
 //AgtModuleManager ListModules
 return "", nil
}

func(np *ModuleManager) ListAllModules ()(string, error) {
 //parameters: 
 //AgtModuleManager ListAllModules
 return "", nil
}

func(np *ModuleManager) GetSerialNumber ()(string, error) {
 //parameters: ModuleNumber
 //AgtModuleManager GetSerialNumber
 return "", nil
}

func(np *ModuleManager) GetModuleNumber ()(string, error) {
 //parameters: SerialNumber
 //AgtModuleManager GetModuleNumber
 return "", nil
}

func(np *ModuleManager) GetModuleName ()(string, error) {
 //parameters: SerialNumber
 //AgtModuleManager GetModuleName
 return "", nil
}

func(np *ModuleManager) SaveModuleList ()(string, error) {
 //parameters: 
 //AgtModuleManager SaveModuleList
 return "", nil
}

func(np *ModuleManager) GetSavedModuleList ()(string, error) {
 //parameters: 
 //AgtModuleManager GetSavedModuleList
 return "", nil
}

func(np *ModuleManager) ListNewModules ()(string, error) {
 //parameters: 
 //AgtModuleManager ListNewModules
 return "", nil
}

func(np *ModuleManager) ListMissingModules ()(string, error) {
 //parameters: 
 //AgtModuleManager ListMissingModules
 return "", nil
}

func(np *ModuleManager) GetPortName ()(string, error) {
 //parameters: SerialNumber PortNumber
 //AgtModuleManager GetPortName
 return "", nil
}

func(np *ModuleManager) GetNcpCount ()(string, error) {
 //parameters: SerialNumber
 //AgtModuleManager GetNcpCount
 return "", nil
}

func(np *ModuleManager) GetChassisNumber ()(string, error) {
 //parameters: SerialNumber
 //AgtModuleManager GetChassisNumber
 return "", nil
}

func(np *ModuleManager) GetChassisSlotNumber ()(string, error) {
 //parameters: SerialNumber
 //AgtModuleManager GetChassisSlotNumber
 return "", nil
}

func(np *ModuleManager) SetModuleAnnotation () error {
 //parameters: SerialNumber ModuleAnnotation
 //AgtModuleManager SetModuleAnnotation, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) GetModuleAnnotation ()(string, error) {
 //parameters: SerialNumber
 //AgtModuleManager GetModuleAnnotation
 return "", nil
}

func(np *ModuleManager) IsDummyModule () error {
 //parameters: SerialNumber
 //AgtModuleManager IsDummyModule, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) GetIpAddress ()(string, error) {
 //parameters: SerialNumber NcpIndex
 //AgtModuleManager GetIpAddress
 return "", nil
}

func(np *ModuleManager) GetPrimaryIpAddress ()(string, error) {
 //parameters: SerialNumber
 //AgtModuleManager GetPrimaryIpAddress
 return "", nil
}

func(np *ModuleManager) GetHostIpAddress ()(string, error) {
 //parameters: 
 //AgtModuleManager GetHostIpAddress
 return "", nil
}

func(np *ModuleManager) GetModuleState ()(string, error) {
 //parameters: SerialNumber
 //AgtModuleManager GetModuleState
 return "", nil
}

func(np *ModuleManager) GetModuleLock ()(string, error) {
 //parameters: SerialNumber
 //AgtModuleManager GetModuleLock
 return "", nil
}

func(np *ModuleManager) UnlockModule () error {
 //parameters: SerialNumber
 //AgtModuleManager UnlockModule, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) RebootModule () error {
 //parameters: SerialNumber
 //AgtModuleManager RebootModule, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) FlashModuleLEDs () error {
 //parameters: SerialNumber
 //AgtModuleManager FlashModuleLEDs, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) ShowIpAddresses () error {
 //parameters: 
 //AgtModuleManager ShowIpAddresses, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) GetNumberingErrorMessage ()(string, error) {
 //parameters: 
 //AgtModuleManager GetNumberingErrorMessage
 return "", nil
}

func(np *ModuleManager) GetChassisControllerUpgradeInfo ()(string, error) {
 //parameters: ChassisNumber
 //AgtModuleManager GetChassisControllerUpgradeInfo
 return "", nil
}

func(np *ModuleManager) ListChassisRequiringUpgrade ()(string, error) {
 //parameters: 
 //AgtModuleManager ListChassisRequiringUpgrade
 return "", nil
}

func(np *ModuleManager) RebootPort () error {
 //parameters: SerialNumber PortNumber
 //AgtModuleManager RebootPort, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) IsMultiUserChassisSupported () error {
 //parameters: ChassisNumber
 //AgtModuleManager IsMultiUserChassisSupported, m.Object, m.Name);
 return nil
}

func(np *ModuleManager) IsMultiUserChassisEnabled () error {
 //parameters: ChassisNumber
 //AgtModuleManager IsMultiUserChassisEnabled, m.Object, m.Name);
 return nil
}

