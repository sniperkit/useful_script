package n2xsdk

type BurstProfile struct {
 Handler string
}

func(np *BurstProfile) GetType ()(string, error) {
 //parameters: Handle
 //AgtBurstProfile GetType
 return "", nil
}

func(np *BurstProfile) GetName ()(string, error) {
 //parameters: Handle
 //AgtBurstProfile GetName
 return "", nil
}

func(np *BurstProfile) GetLockCount ()(string, error) {
 //parameters: Handle
 //AgtBurstProfile GetLockCount
 return "", nil
}

func(np *BurstProfile) GetSourcePort ()(string, error) {
 //parameters: ProfileHandle
 //AgtBurstProfile GetSourcePort
 return "", nil
}

func(np *BurstProfile) SetProfileType () error {
 //parameters: ProfileHandle ProfileType
 //AgtBurstProfile SetProfileType, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) GetProfileType ()(string, error) {
 //parameters: ProfileHandle
 //AgtBurstProfile GetProfileType
 return "", nil
}

func(np *BurstProfile) SetAverageLoad () error {
 //parameters: ProfileHandle AverageLoad Units
 //AgtBurstProfile SetAverageLoad, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) GetAverageLoad ()(string, error) {
 //parameters: ProfileHandle Units
 //AgtBurstProfile GetAverageLoad
 return "", nil
}

func(np *BurstProfile) SetMode () error {
 //parameters: ProfileHandle Mode
 //AgtBurstProfile SetMode, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) GetMode ()(string, error) {
 //parameters: ProfileHandle
 //AgtBurstProfile GetMode
 return "", nil
}

func(np *BurstProfile) SetNumberOfPacketsToInject () error {
 //parameters: ProfileHandle NumberOfPackets
 //AgtBurstProfile SetNumberOfPacketsToInject, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) GetNumberOfPacketsToInject ()(string, error) {
 //parameters: ProfileHandle
 //AgtBurstProfile GetNumberOfPacketsToInject
 return "", nil
}

func(np *BurstProfile) SendSingleInjection () error {
 //parameters: ProfileHandle
 //AgtBurstProfile SendSingleInjection, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) IsSendSingleInjectionInProgress () error {
 //parameters: ProfileHandle
 //AgtBurstProfile IsSendSingleInjectionInProgress, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) ListStreamGroups ()(string, error) {
 //parameters: ProfileHandle
 //AgtBurstProfile ListStreamGroups
 return "", nil
}

func(np *BurstProfile) AddStreamGroups () error {
 //parameters: ProfileHandle Count psaStreamGroupHandles
 //AgtBurstProfile AddStreamGroups, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) GetInterdepartureValues ()(string, error) {
 //parameters: ProfileHandle
 //AgtBurstProfile GetInterdepartureValues
 return "", nil
}

func(np *BurstProfile) GetInterdepartureRepeatCounts ()(string, error) {
 //parameters: ProfileHandle
 //AgtBurstProfile GetInterdepartureRepeatCounts
 return "", nil
}

func(np *BurstProfile) GetInterdepartureValueResolution ()(string, error) {
 //parameters: ProfileHandle
 //AgtBurstProfile GetInterdepartureValueResolution
 return "", nil
}

func(np *BurstProfile) Enable () error {
 //parameters: ProfileHandle
 //AgtBurstProfile Enable, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) Disable () error {
 //parameters: ProfileHandle
 //AgtBurstProfile Disable, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) IsEnabled () error {
 //parameters: ProfileHandle
 //AgtBurstProfile IsEnabled, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) SetBurstLoad () error {
 //parameters: ProfileHandle BurstLoad Units
 //AgtBurstProfile SetBurstLoad, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) GetBurstLoad ()(string, error) {
 //parameters: ProfileHandle Units
 //AgtBurstProfile GetBurstLoad
 return "", nil
}

func(np *BurstProfile) SetBurstLength () error {
 //parameters: ProfileHandle BurstLength
 //AgtBurstProfile SetBurstLength, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) GetBurstLength ()(string, error) {
 //parameters: ProfileHandle
 //AgtBurstProfile GetBurstLength
 return "", nil
}

func(np *BurstProfile) IsContiguousBurstsSupported () error {
 //parameters: PortHandle
 //AgtBurstProfile IsContiguousBurstsSupported, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) IsContiguousBurstsEnabled () error {
 //parameters: ProfileHandle
 //AgtBurstProfile IsContiguousBurstsEnabled, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) EnableContiguousBursts () error {
 //parameters: ProfileHandle
 //AgtBurstProfile EnableContiguousBursts, m.Object, m.Name);
 return nil
}

func(np *BurstProfile) DisableContiguousBursts () error {
 //parameters: ProfileHandle
 //AgtBurstProfile DisableContiguousBursts, m.Object, m.Name);
 return nil
}

