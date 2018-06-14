package n2xsdk

type IptvTestScenario struct {
 Handler string
}

func(np *IptvTestScenario) Enable () error {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario Enable, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) Disable () error {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario Disable, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) IsEnabled () error {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario IsEnabled, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) IsIncludeInReportEnabled () error {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario IsIncludeInReportEnabled, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) GetStateCount ()(string, error) {
 //parameters: ScenarioHandle ScenarioStateType
 //AgtIptvTestScenario GetStateCount
 return "", nil
}

func(np *IptvTestScenario) GetState ()(string, error) {
 //parameters: ScenarioInstance
 //AgtIptvTestScenario GetState
 return "", nil
}

func(np *IptvTestScenario) GetStatus ()(string, error) {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario GetStatus
 return "", nil
}

func(np *IptvTestScenario) SetStartDelayIncrementingRange () error {
 //parameters: ScenarioHandle TestStartDelay Increment Repeat
 //AgtIptvTestScenario SetStartDelayIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) GetStartDelayIncrementingRange ()(string, error) {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario GetStartDelayIncrementingRange
 return "", nil
}

func(np *IptvTestScenario) SetTestDurationInMs () error {
 //parameters: ScenarioHandle TestDuration
 //AgtIptvTestScenario SetTestDurationInMs, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) GetTestDurationInMs ()(string, error) {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario GetTestDurationInMs
 return "", nil
}

func(np *IptvTestScenario) GetMaxNumberOfChannelsPerSubscriber ()(string, error) {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario GetMaxNumberOfChannelsPerSubscriber
 return "", nil
}

func(np *IptvTestScenario) SetWatchTimeMode () error {
 //parameters: ScenarioHandle WatchTimeMode
 //AgtIptvTestScenario SetWatchTimeMode, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) GetWatchTimeMode ()(string, error) {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario GetWatchTimeMode
 return "", nil
}

func(np *IptvTestScenario) SetRandomSeed () error {
 //parameters: ScenarioHandle RandomSeed
 //AgtIptvTestScenario SetRandomSeed, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) GetRandomSeed ()(string, error) {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario GetRandomSeed
 return "", nil
}

func(np *IptvTestScenario) ListAvailableGroupPoolsForChannels ()(string, error) {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario ListAvailableGroupPoolsForChannels
 return "", nil
}

func(np *IptvTestScenario) AddChannelPools () error {
 //parameters: ScenarioHandle GroupPoolHandles
 //AgtIptvTestScenario AddChannelPools, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) RemoveChannelPools () error {
 //parameters: ScenarioHandle GroupPoolHandles
 //AgtIptvTestScenario RemoveChannelPools, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) RemoveAllChannelPools () error {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario RemoveAllChannelPools, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) ListChannelPools ()(string, error) {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario ListChannelPools
 return "", nil
}

func(np *IptvTestScenario) GetAssignedChannels ()(string, error) {
 //parameters: ScenarioInstance
 //AgtIptvTestScenario GetAssignedChannels
 return "", nil
}

func(np *IptvTestScenario) StartTest () error {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario StartTest, m.Object, m.Name);
 return nil
}

func(np *IptvTestScenario) StopTest () error {
 //parameters: ScenarioHandle
 //AgtIptvTestScenario StopTest, m.Object, m.Name);
 return nil
}

