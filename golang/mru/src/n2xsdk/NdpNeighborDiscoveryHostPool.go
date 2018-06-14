package n2xsdk

type NdpNeighborDiscoveryHostPool struct {
 Handler string
}

func(np *NdpNeighborDiscoveryHostPool) EnableAutoSendNeighborSolicitations () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool EnableAutoSendNeighborSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) DisableAutoSendNeighborSolicitations () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool DisableAutoSendNeighborSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) IsAutoSendNeighborSolicitationsEnabled () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool IsAutoSendNeighborSolicitationsEnabled, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetCustomState ()(string, error) {
 //parameters: DevicePoolInstance
 //AgtNdpNeighborDiscoveryHostPool GetCustomState
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SendNeighborSolicitations () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool SendNeighborSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) Reset () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool Reset, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) CancelNeighborSolicitations () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool CancelNeighborSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetRemoteMacAddress ()(string, error) {
 //parameters: DevicePoolInstance
 //AgtNdpNeighborDiscoveryHostPool GetRemoteMacAddress
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) EnableStatelessTransmit () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool EnableStatelessTransmit, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) DisableStatelessTransmit () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool DisableStatelessTransmit, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) IsStatelessTransmitEnabled () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool IsStatelessTransmitEnabled, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) SetStatelessTransmitPacketType () error {
 //parameters: DevicePoolHandle PacketType
 //AgtNdpNeighborDiscoveryHostPool SetStatelessTransmitPacketType, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetStatelessTransmitPacketType ()(string, error) {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool GetStatelessTransmitPacketType
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SetStatelessTransmitPacket () error {
 //parameters: DeviceHandle Packet
 //AgtNdpNeighborDiscoveryHostPool SetStatelessTransmitPacket, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetStatelessTransmitPacket ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpNeighborDiscoveryHostPool GetStatelessTransmitPacket
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) StartStatelessTransmit () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool StartStatelessTransmit, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) StopStatelessTransmit () error {
 //parameters: DevicePoolHandle
 //AgtNdpNeighborDiscoveryHostPool StopStatelessTransmit, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) SetStatelessTransmitStartEventState () error {
 //parameters: DeviceHandle StartEventState
 //AgtNdpNeighborDiscoveryHostPool SetStatelessTransmitStartEventState, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetStatelessTransmitStartEventState ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpNeighborDiscoveryHostPool GetStatelessTransmitStartEventState
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SendNeighborSolicitationsByPort () error {
 //parameters: PortHandle
 //AgtNdpNeighborDiscoveryHostPool SendNeighborSolicitationsByPort, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) ResetByPort () error {
 //parameters: PortHandle
 //AgtNdpNeighborDiscoveryHostPool ResetByPort, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) CancelNeighborSolicitationsByPort () error {
 //parameters: PortHandle
 //AgtNdpNeighborDiscoveryHostPool CancelNeighborSolicitationsByPort, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) SendAllNeighborSolicitations () error {
 //parameters: 
 //AgtNdpNeighborDiscoveryHostPool SendAllNeighborSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) ResetAll () error {
 //parameters: 
 //AgtNdpNeighborDiscoveryHostPool ResetAll, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) CancelAllNeighborSolicitations () error {
 //parameters: 
 //AgtNdpNeighborDiscoveryHostPool CancelAllNeighborSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) SetRateControl () error {
 //parameters: PortHandle RateControlMode
 //AgtNdpNeighborDiscoveryHostPool SetRateControl, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetRateControl ()(string, error) {
 //parameters: PortHandle
 //AgtNdpNeighborDiscoveryHostPool GetRateControl
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SetFixedRateProfile () error {
 //parameters: PortHandle FixedRate MaxBurstSize
 //AgtNdpNeighborDiscoveryHostPool SetFixedRateProfile, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetFixedRateProfile ()(string, error) {
 //parameters: PortHandle
 //AgtNdpNeighborDiscoveryHostPool GetFixedRateProfile
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SetSelfPacedProfile () error {
 //parameters: PortHandle MaxPackets MaxRetries RetryTimeout
 //AgtNdpNeighborDiscoveryHostPool SetSelfPacedProfile, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetSelfPacedProfile ()(string, error) {
 //parameters: PortHandle
 //AgtNdpNeighborDiscoveryHostPool GetSelfPacedProfile
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SetStatelessTransmitStartEvent () error {
 //parameters: DeviceHandle StartEvent
 //AgtNdpNeighborDiscoveryHostPool SetStatelessTransmitStartEvent, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetStatelessTransmitStartEvent ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpNeighborDiscoveryHostPool GetStatelessTransmitStartEvent
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SetStatelessTransmitStopEvent () error {
 //parameters: DeviceHandle StopEvent
 //AgtNdpNeighborDiscoveryHostPool SetStatelessTransmitStopEvent, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetStatelessTransmitStopEvent ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpNeighborDiscoveryHostPool GetStatelessTransmitStopEvent
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SetStatelessTransmitStopEventCount () error {
 //parameters: DeviceHandle StopEventCount
 //AgtNdpNeighborDiscoveryHostPool SetStatelessTransmitStopEventCount, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetStatelessTransmitStopEventCount ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpNeighborDiscoveryHostPool GetStatelessTransmitStopEventCount
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SetStatelessTransmitStopEventTime () error {
 //parameters: DeviceHandle TimeUnits Time
 //AgtNdpNeighborDiscoveryHostPool SetStatelessTransmitStopEventTime, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetStatelessTransmitStopEventTime ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpNeighborDiscoveryHostPool GetStatelessTransmitStopEventTime
 return "", nil
}

func(np *NdpNeighborDiscoveryHostPool) SetStatelessTransmitRate () error {
 //parameters: DeviceHandle BurstSize PeriodUnits Period
 //AgtNdpNeighborDiscoveryHostPool SetStatelessTransmitRate, m.Object, m.Name);
 return nil
}

func(np *NdpNeighborDiscoveryHostPool) GetStatelessTransmitRate ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpNeighborDiscoveryHostPool GetStatelessTransmitRate
 return "", nil
}

