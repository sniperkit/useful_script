package n2xsdk

type ncpClientPool struct {
 Handler string
}

func(np *ncpClientPool) SetAnSubscriberLinePoolName () error {
 //parameters: DeviceHandle SubscriberPoolHandle SubscriberPoolName
 //AgtAncpClientPool SetAnSubscriberLinePoolName, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) GetAnSubscriberLinePoolName ()(string, error) {
 //parameters: DeviceHandle SubscriberPoolHandle
 //AgtAncpClientPool GetAnSubscriberLinePoolName
 return "", nil
}

func(np *ncpClientPool) AddAnSubscriberLinePool () error {
 //parameters: DeviceHandle SubscriberPoolType
 //AgtAncpClientPool AddAnSubscriberLinePool, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) RemoveAnSubscriberLinePool () error {
 //parameters: DeviceHandle SubscriberPoolHandle
 //AgtAncpClientPool RemoveAnSubscriberLinePool, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) ClearAnSubscriberLinePools () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool ClearAnSubscriberLinePools, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) ListAnSubscriberLinePools ()(string, error) {
 //parameters: DeviceHandle
 //AgtAncpClientPool ListAnSubscriberLinePools
 return "", nil
}

func(np *ncpClientPool) ListAnSubscriberLinePoolsByType ()(string, error) {
 //parameters: DeviceHandle SubscriberPoolType
 //AgtAncpClientPool ListAnSubscriberLinePoolsByType
 return "", nil
}

func(np *ncpClientPool) GetAnSubscriberLinePoolType ()(string, error) {
 //parameters: DeviceHandle SubscriberPoolHandle
 //AgtAncpClientPool GetAnSubscriberLinePoolType
 return "", nil
}

func(np *ncpClientPool) SetAncpStandard () error {
 //parameters: DeviceHandle AncpStandard
 //AgtAncpClientPool SetAncpStandard, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) GetAncpStandard ()(string, error) {
 //parameters: DeviceHandle
 //AgtAncpClientPool GetAncpStandard
 return "", nil
}

func(np *ncpClientPool) SetGsmpStandard () error {
 //parameters: DeviceHandle GsmpV3Standard
 //AgtAncpClientPool SetGsmpStandard, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) GetGsmpStandard ()(string, error) {
 //parameters: DeviceHandle
 //AgtAncpClientPool GetGsmpStandard
 return "", nil
}

func(np *ncpClientPool) SetPartitionId () error {
 //parameters: DeviceHandle PartitionId Repeat Increment
 //AgtAncpClientPool SetPartitionId, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) GetPartitionId ()(string, error) {
 //parameters: DeviceHandle
 //AgtAncpClientPool GetPartitionId
 return "", nil
}

func(np *ncpClientPool) SetKeepAliveTime () error {
 //parameters: DeviceHandle KeepAliveTimeSec
 //AgtAncpClientPool SetKeepAliveTime, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) GetKeepAliveTime ()(string, error) {
 //parameters: DeviceHandle
 //AgtAncpClientPool GetKeepAliveTime
 return "", nil
}

func(np *ncpClientPool) DisableTopologyDiscovery () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool DisableTopologyDiscovery, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) EnableTopologyDiscovery () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool EnableTopologyDiscovery, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) IsTopologyDiscoveryEnabled () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool IsTopologyDiscoveryEnabled, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) DisableTopologyDiscoveryReturnReceipt () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool DisableTopologyDiscoveryReturnReceipt, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) EnableTopologyDiscoveryReturnReceipt () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool EnableTopologyDiscoveryReturnReceipt, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) IsTopologyDiscoveryReturnReceiptEnabled () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool IsTopologyDiscoveryReturnReceiptEnabled, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) DisableTransMulticastCapability () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool DisableTransMulticastCapability, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) EnableTransMulticastCapability () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool EnableTransMulticastCapability, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) IsTransMulticastCapabilityEnabled () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool IsTransMulticastCapabilityEnabled, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) SetTopologyDiscoveryEventsCount () error {
 //parameters: DeviceHandle EventsCount
 //AgtAncpClientPool SetTopologyDiscoveryEventsCount, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) GetTopologyDiscoveryEventsCount ()(string, error) {
 //parameters: DeviceHandle
 //AgtAncpClientPool GetTopologyDiscoveryEventsCount
 return "", nil
}

func(np *ncpClientPool) GetTopologyDiscoveryEventsPerSecond ()(string, error) {
 //parameters: DeviceHandle
 //AgtAncpClientPool GetTopologyDiscoveryEventsPerSecond
 return "", nil
}

func(np *ncpClientPool) SetTopologyDiscoveryEventsInterval () error {
 //parameters: DeviceHandle EventsInterval
 //AgtAncpClientPool SetTopologyDiscoveryEventsInterval, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) GetTopologyDiscoveryEventsInterval ()(string, error) {
 //parameters: DeviceHandle
 //AgtAncpClientPool GetTopologyDiscoveryEventsInterval
 return "", nil
}

func(np *ncpClientPool) SendReset () error {
 //parameters: SessionList
 //AgtAncpClientPool SendReset, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) EnableLineConfiguration () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool EnableLineConfiguration, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) DisableLineConfiguration () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool DisableLineConfiguration, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) IsLineConfigurationEnabled () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool IsLineConfigurationEnabled, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) EnableOam () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool EnableOam, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) DisableOam () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool DisableOam, m.Object, m.Name);
 return nil
}

func(np *ncpClientPool) IsOamEnabled () error {
 //parameters: DeviceHandle
 //AgtAncpClientPool IsOamEnabled, m.Object, m.Name);
 return nil
}

