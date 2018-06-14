package n2xsdk

type PimVpnGroupInstanceTable struct {
 Handler string
}

func(np *PimVpnGroupInstanceTable) GetState ()(string, error) {
 //parameters: GroupIdentifiers GroupRangeHandle
 //AgtPimVpnGroupInstanceTable GetState
 return "", nil
}

func(np *PimVpnGroupInstanceTable) AddGroupRange () error {
 //parameters: GroupHandle GroupInstance
 //AgtPimVpnGroupInstanceTable AddGroupRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) RemoveGroupRange () error {
 //parameters: GroupHandle GroupInstance GroupRangeHandle
 //AgtPimVpnGroupInstanceTable RemoveGroupRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) ListGroupRanges ()(string, error) {
 //parameters: GroupHandle GroupInstance
 //AgtPimVpnGroupInstanceTable ListGroupRanges
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetGroupRangeName () error {
 //parameters: GroupHandle GroupInstance GroupRangeHandle GroupRangeName
 //AgtPimVpnGroupInstanceTable SetGroupRangeName, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetGroupRangeName ()(string, error) {
 //parameters: GroupHandle GroupInstance GroupRangeHandle
 //AgtPimVpnGroupInstanceTable GetGroupRangeName
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetGroupAddressIncrementingRange () error {
 //parameters: GroupHandle GroupInstance GroupRangeHandle FirstGroupAddress PrefixLength Increment Count
 //AgtPimVpnGroupInstanceTable SetGroupAddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetGroupAddressIncrementingRange ()(string, error) {
 //parameters: GroupHandle GroupInstance GroupRangeHandle
 //AgtPimVpnGroupInstanceTable GetGroupAddressIncrementingRange
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetGroupAddress () error {
 //parameters: GroupHandle GroupInstance GroupRangeHandle GroupAddress
 //AgtPimVpnGroupInstanceTable SetGroupAddress, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetGroupAddress ()(string, error) {
 //parameters: GroupHandle GroupInstance GroupRangeHandle
 //AgtPimVpnGroupInstanceTable GetGroupAddress
 return "", nil
}

func(np *PimVpnGroupInstanceTable) AddSourceRange () error {
 //parameters: GroupHandle GroupInstance GroupRangeHandle
 //AgtPimVpnGroupInstanceTable AddSourceRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) RemoveSourceRange () error {
 //parameters: GroupHandle GroupInstance GroupRangeHandle SourceRangeHandle
 //AgtPimVpnGroupInstanceTable RemoveSourceRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) ListSourceRanges ()(string, error) {
 //parameters: GroupHandle GroupInstance GroupRangeHandle
 //AgtPimVpnGroupInstanceTable ListSourceRanges
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetSourceAddressIncrementingRange () error {
 //parameters: GroupHandle GroupInstance GroupRangeHandle SourceRangeHandle FirstSourceAddress PrefixLength Increment Count
 //AgtPimVpnGroupInstanceTable SetSourceAddressIncrementingRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetSourceAddressIncrementingRange ()(string, error) {
 //parameters: GroupHandle GroupInstance GroupRangeHandle SourceRangeHandle
 //AgtPimVpnGroupInstanceTable GetSourceAddressIncrementingRange
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetSourceAddress () error {
 //parameters: GroupHandle GroupInstance GroupRangeHandle SourceRangeHandle SourceAddress
 //AgtPimVpnGroupInstanceTable SetSourceAddress, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetSourceAddress ()(string, error) {
 //parameters: GroupHandle GroupInstance GroupRangeHandle SourceRangeHandle
 //AgtPimVpnGroupInstanceTable GetSourceAddress
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetRpAddress () error {
 //parameters: GroupHandle GroupInstance GroupRangeHandle RpAddress
 //AgtPimVpnGroupInstanceTable SetRpAddress, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetRpAddress ()(string, error) {
 //parameters: GroupHandle GroupInstance GroupRangeHandle
 //AgtPimVpnGroupInstanceTable GetRpAddress
 return "", nil
}

func(np *PimVpnGroupInstanceTable) Join () error {
 //parameters: GroupIdentifiers
 //AgtPimVpnGroupInstanceTable Join, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) Prune () error {
 //parameters: GroupIdentifiers
 //AgtPimVpnGroupInstanceTable Prune, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) StartSources () error {
 //parameters: GroupIdentifiers
 //AgtPimVpnGroupInstanceTable StartSources, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) StopSources () error {
 //parameters: GroupIdentifiers
 //AgtPimVpnGroupInstanceTable StopSources, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) JoinGroupRange () error {
 //parameters: GroupIdentifiers GroupRangeHandle
 //AgtPimVpnGroupInstanceTable JoinGroupRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) PruneGroupRange () error {
 //parameters: GroupIdentifiers GroupRangeHandle
 //AgtPimVpnGroupInstanceTable PruneGroupRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) StartSourcesOfGroupRange () error {
 //parameters: GroupIdentifiers GroupRangeHandle
 //AgtPimVpnGroupInstanceTable StartSourcesOfGroupRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) StopSourcesOfGroupRange () error {
 //parameters: GroupIdentifiers GroupRangeHandle
 //AgtPimVpnGroupInstanceTable StopSourcesOfGroupRange, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) Enable () error {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable Enable, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) Disable () error {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable Disable, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) SetAction () error {
 //parameters: GroupHandle ActionType
 //AgtPimVpnGroupInstanceTable SetAction, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetAction ()(string, error) {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable GetAction
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetJoinPruneActionUponEnable () error {
 //parameters: GroupHandle JoinPruneActionUponEnable
 //AgtPimVpnGroupInstanceTable SetJoinPruneActionUponEnable, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetJoinPruneActionUponEnable ()(string, error) {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable GetJoinPruneActionUponEnable
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetRegisterActionUponEnable () error {
 //parameters: GroupHandle RegisterActionUponEnable
 //AgtPimVpnGroupInstanceTable SetRegisterActionUponEnable, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetRegisterActionUponEnable ()(string, error) {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable GetRegisterActionUponEnable
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetNumberOfPdusPerInterval () error {
 //parameters: GroupHandle NumberOfPdusPerInterval
 //AgtPimVpnGroupInstanceTable SetNumberOfPdusPerInterval, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetNumberOfPdusPerInterval ()(string, error) {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable GetNumberOfPdusPerInterval
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetPduInterval () error {
 //parameters: GroupHandle PduInterval
 //AgtPimVpnGroupInstanceTable SetPduInterval, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetPduInterval ()(string, error) {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable GetPduInterval
 return "", nil
}

func(np *PimVpnGroupInstanceTable) SetAggregationFactor () error {
 //parameters: GroupHandle AggregationFactor
 //AgtPimVpnGroupInstanceTable SetAggregationFactor, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetAggregationFactor ()(string, error) {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable GetAggregationFactor
 return "", nil
}

func(np *PimVpnGroupInstanceTable) EnableSourceSpecificJoinPrune () error {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable EnableSourceSpecificJoinPrune, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) DisableSourceSpecificJoinPrune () error {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable DisableSourceSpecificJoinPrune, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) IsSourceSpecificJoinPruneEnabled () error {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable IsSourceSpecificJoinPruneEnabled, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) SetSGDistribution () error {
 //parameters: GroupHandle SGDistribution
 //AgtPimVpnGroupInstanceTable SetSGDistribution, m.Object, m.Name);
 return nil
}

func(np *PimVpnGroupInstanceTable) GetSGDistribution ()(string, error) {
 //parameters: GroupHandle
 //AgtPimVpnGroupInstanceTable GetSGDistribution
 return "", nil
}

