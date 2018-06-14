package n2xsdk

type NdpAaRouterPool struct {
 Handler string
}

func(np *NdpAaRouterPool) SendNdpPacket () error {
 //parameters: DeviceHandle
 //AgtNdpAaRouterPool SendNdpPacket, m.Object, m.Name);
 return nil
}

func(np *NdpAaRouterPool) EnableAutoSendNdpPacket () error {
 //parameters: DeviceHandle
 //AgtNdpAaRouterPool EnableAutoSendNdpPacket, m.Object, m.Name);
 return nil
}

func(np *NdpAaRouterPool) DisableAutoSendNdpPacket () error {
 //parameters: DeviceHandle
 //AgtNdpAaRouterPool DisableAutoSendNdpPacket, m.Object, m.Name);
 return nil
}

func(np *NdpAaRouterPool) IsAutoSendNdpPacketEnabled () error {
 //parameters: DeviceHandle
 //AgtNdpAaRouterPool IsAutoSendNdpPacketEnabled, m.Object, m.Name);
 return nil
}

func(np *NdpAaRouterPool) EnableSendNdpPacketonReceiveIcmpPacket () error {
 //parameters: DeviceHandle
 //AgtNdpAaRouterPool EnableSendNdpPacketonReceiveIcmpPacket, m.Object, m.Name);
 return nil
}

func(np *NdpAaRouterPool) DisableSendNdpPacketonReceiveIcmpPacket () error {
 //parameters: DeviceHandle
 //AgtNdpAaRouterPool DisableSendNdpPacketonReceiveIcmpPacket, m.Object, m.Name);
 return nil
}

func(np *NdpAaRouterPool) IsNdpPacketonReceiveIcmpPacketEnabled () error {
 //parameters: DeviceHandle
 //AgtNdpAaRouterPool IsNdpPacketonReceiveIcmpPacketEnabled, m.Object, m.Name);
 return nil
}

