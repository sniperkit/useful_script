package n2xsdk

type NdpAddressAutoconfigHostPool struct {
 Handler string
}

func(np *NdpAddressAutoconfigHostPool) EnableAutoSendRouterSolicitations () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool EnableAutoSendRouterSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) DisableAutoSendRouterSolicitations () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool DisableAutoSendRouterSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) IsAutoSendRouterSolicitationsEnabled () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool IsAutoSendRouterSolicitationsEnabled, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) SendRouterSolicitations () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool SendRouterSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) Reset () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool Reset, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) CancelRouterSolicitations () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool CancelRouterSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetCustomState ()(string, error) {
 //parameters: DevicePoolInstance
 //AgtNdpAddressAutoconfigHostPool GetCustomState
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) EnableStatelessTransmit () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool EnableStatelessTransmit, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) DisableStatelessTransmit () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool DisableStatelessTransmit, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) IsStatelessTransmitEnabled () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool IsStatelessTransmitEnabled, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) SetStatelessTransmitPacketType () error {
 //parameters: DevicePoolHandle PacketType
 //AgtNdpAddressAutoconfigHostPool SetStatelessTransmitPacketType, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetStatelessTransmitPacketType ()(string, error) {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool GetStatelessTransmitPacketType
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) SetStatelessTransmitPacket () error {
 //parameters: DeviceHandle Packet
 //AgtNdpAddressAutoconfigHostPool SetStatelessTransmitPacket, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetStatelessTransmitPacket ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpAddressAutoconfigHostPool GetStatelessTransmitPacket
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) StartStatelessTransmit () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool StartStatelessTransmit, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) StopStatelessTransmit () error {
 //parameters: DevicePoolHandle
 //AgtNdpAddressAutoconfigHostPool StopStatelessTransmit, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) SetStatelessTransmitStartEventState () error {
 //parameters: DeviceHandle StartEventState
 //AgtNdpAddressAutoconfigHostPool SetStatelessTransmitStartEventState, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetStatelessTransmitStartEventState ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpAddressAutoconfigHostPool GetStatelessTransmitStartEventState
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) SendRouterSolicitationsByPort () error {
 //parameters: PortHandle
 //AgtNdpAddressAutoconfigHostPool SendRouterSolicitationsByPort, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) ResetByPort () error {
 //parameters: PortHandle
 //AgtNdpAddressAutoconfigHostPool ResetByPort, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) CancelRouterSolicitationsByPort () error {
 //parameters: PortHandle
 //AgtNdpAddressAutoconfigHostPool CancelRouterSolicitationsByPort, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) SendAllRouterSolicitations () error {
 //parameters: 
 //AgtNdpAddressAutoconfigHostPool SendAllRouterSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) ResetAll () error {
 //parameters: 
 //AgtNdpAddressAutoconfigHostPool ResetAll, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) CancelAllRouterSolicitations () error {
 //parameters: 
 //AgtNdpAddressAutoconfigHostPool CancelAllRouterSolicitations, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) SetRateControl () error {
 //parameters: PortHandle RateControlMode
 //AgtNdpAddressAutoconfigHostPool SetRateControl, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetRateControl ()(string, error) {
 //parameters: PortHandle
 //AgtNdpAddressAutoconfigHostPool GetRateControl
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) SetFixedRateProfile () error {
 //parameters: PortHandle FixedRate MaxBurstSize
 //AgtNdpAddressAutoconfigHostPool SetFixedRateProfile, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetFixedRateProfile ()(string, error) {
 //parameters: PortHandle
 //AgtNdpAddressAutoconfigHostPool GetFixedRateProfile
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) SetSelfPacedProfile () error {
 //parameters: PortHandle MaxPackets MaxRetries RetryTimeout
 //AgtNdpAddressAutoconfigHostPool SetSelfPacedProfile, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetSelfPacedProfile ()(string, error) {
 //parameters: PortHandle
 //AgtNdpAddressAutoconfigHostPool GetSelfPacedProfile
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) SetStatelessTransmitStartEvent () error {
 //parameters: DeviceHandle StartEvent
 //AgtNdpAddressAutoconfigHostPool SetStatelessTransmitStartEvent, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetStatelessTransmitStartEvent ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpAddressAutoconfigHostPool GetStatelessTransmitStartEvent
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) SetStatelessTransmitStopEvent () error {
 //parameters: DeviceHandle StopEvent
 //AgtNdpAddressAutoconfigHostPool SetStatelessTransmitStopEvent, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetStatelessTransmitStopEvent ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpAddressAutoconfigHostPool GetStatelessTransmitStopEvent
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) SetStatelessTransmitStopEventCount () error {
 //parameters: DeviceHandle StopEventCount
 //AgtNdpAddressAutoconfigHostPool SetStatelessTransmitStopEventCount, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetStatelessTransmitStopEventCount ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpAddressAutoconfigHostPool GetStatelessTransmitStopEventCount
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) SetStatelessTransmitStopEventTime () error {
 //parameters: DeviceHandle TimeUnits Time
 //AgtNdpAddressAutoconfigHostPool SetStatelessTransmitStopEventTime, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetStatelessTransmitStopEventTime ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpAddressAutoconfigHostPool GetStatelessTransmitStopEventTime
 return "", nil
}

func(np *NdpAddressAutoconfigHostPool) SetStatelessTransmitRate () error {
 //parameters: DeviceHandle BurstSize PeriodUnits Period
 //AgtNdpAddressAutoconfigHostPool SetStatelessTransmitRate, m.Object, m.Name);
 return nil
}

func(np *NdpAddressAutoconfigHostPool) GetStatelessTransmitRate ()(string, error) {
 //parameters: DeviceHandle
 //AgtNdpAddressAutoconfigHostPool GetStatelessTransmitRate
 return "", nil
}

