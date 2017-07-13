package dsl
import ( 
 "command"
 )
type InstructionHandler func(Switch, map[string]string) []*command.Command
func (p *Parser) Init() {
p.Ctrie.Insert([]byte("Port.Slot.Type.Enable"), InstructionHandler(PortSlotTypeEnable))
p.Ctrie.Insert([]byte("Port.Slot.Type.Disable"), InstructionHandler(PortSlotTypeDisable))
p.Ctrie.Insert([]byte("Port.Slot.Type.Speed"), InstructionHandler(PortSlotTypeSpeed))
p.Ctrie.Insert([]byte("Port.Slot.Type.Pvid"), InstructionHandler(PortSlotTypePvid))
p.Ctrie.Insert([]byte("VLAN"), InstructionHandler(VLAN))
p.Ctrie.Insert([]byte("No.VLAN"), InstructionHandler(NoVLAN))
p.Ctrie.Insert([]byte("VLAN.Add.Type.Slot.Port"), InstructionHandler(VLANAddTypeSlotPort))
p.Ctrie.Insert([]byte("VLAN.AddT.Type.Slot.Port"), InstructionHandler(VLANAddTTypeSlotPort))
p.Ctrie.Insert([]byte("VLAN.Del.Type.Slot.Port"), InstructionHandler(VLANDelTypeSlotPort))
p.Ctrie.Insert([]byte("VLAN.DelT.Type.Slot.Port"), InstructionHandler(VLANDelTTypeSlotPort))
p.Ctrie.Insert([]byte("VLAN.Shutdown"), InstructionHandler(VLANShutdown))
p.Ctrie.Insert([]byte("VLAN.NoShutdown"), InstructionHandler(VLANNoShutdown))
p.Ctrie.Insert([]byte("VLAN.IP"), InstructionHandler(VLANIP))
p.Ctrie.Insert([]byte("No.VLAN.IP"), InstructionHandler(NoVLANIP))
p.Ctrie.Insert([]byte("VLAN.IP2"), InstructionHandler(VLANIP2))
p.Ctrie.Insert([]byte("No.VLAN.IP2"), InstructionHandler(NoVLANIP2))
p.Ctrie.Insert([]byte("No.Interface.Type.Ifname"), InstructionHandler(NoInterfaceTypeIfname))
p.Ctrie.Insert([]byte("VLAN.Add.Type.Slot.Port.IP"), InstructionHandler(VLANAddTypeSlotPortIP))
p.Ctrie.Insert([]byte("VLAN.Add.Type.Slot.Port.IP2"), InstructionHandler(VLANAddTypeSlotPortIP2))
p.Ctrie.Insert([]byte("VLAN.Add.Type.Slot.Port.IP.NoShutdown"), InstructionHandler(VLANAddTypeSlotPortIPNoShutdown))
p.Ctrie.Insert([]byte("VLAN.Add.Type.Slot.Port.IP2.NoShutdown"), InstructionHandler(VLANAddTypeSlotPortIP2NoShutdown))
p.Ctrie.Insert([]byte("VLAN.AddT.Type.Slot.Port.IP"), InstructionHandler(VLANAddTTypeSlotPortIP))
p.Ctrie.Insert([]byte("VLAN.AddT.Type.Slot.Port.IP2"), InstructionHandler(VLANAddTTypeSlotPortIP2))
p.Ctrie.Insert([]byte("VLAN.AddT.Type.Slot.Port.IP.NoShutdown"), InstructionHandler(VLANAddTTypeSlotPortIPNoShutdown))
p.Ctrie.Insert([]byte("VLAN.AddT.Type.Slot.Port.IP2.NoShutdown"), InstructionHandler(VLANAddTTypeSlotPortIP2NoShutdown))
p.Ctrie.Insert([]byte("VLAN.IP6.Enable"), InstructionHandler(VLANIP6Enable))
p.Ctrie.Insert([]byte("No.VLAN.IP6.Enable"), InstructionHandler(NoVLANIP6Enable))
p.Ctrie.Insert([]byte("VLAN.IP6"), InstructionHandler(VLANIP6))
p.Ctrie.Insert([]byte("No.VLAN.IP6"), InstructionHandler(NoVLANIP6))
p.Ctrie.Insert([]byte("VLAN.IP6LL"), InstructionHandler(VLANIP6LL))
p.Ctrie.Insert([]byte("VLAN.IP6LL.IP6"), InstructionHandler(VLANIP6LLIP6))
p.Ctrie.Insert([]byte("No.VLAN.IP6LL"), InstructionHandler(NoVLANIP6LL))
p.Ctrie.Insert([]byte("OSPF6"), InstructionHandler(OSPF6))
p.Ctrie.Insert([]byte("No.OSPF6"), InstructionHandler(NoOSPF6))
p.Ctrie.Insert([]byte("OSPF6.Rid"), InstructionHandler(OSPF6Rid))
p.Ctrie.Insert([]byte("No.OSPF6.Rid"), InstructionHandler(NoOSPF6Rid))
p.Ctrie.Insert([]byte("OSPF6.Interface.Area"), InstructionHandler(OSPF6InterfaceArea))
p.Ctrie.Insert([]byte("No.OSPF6.Interface.Area"), InstructionHandler(NoOSPF6InterfaceArea))
p.Ctrie.Insert([]byte("OSPF6.Interface.Cost"), InstructionHandler(OSPF6InterfaceCost))
p.Ctrie.Insert([]byte("OSPF6.Interface.DeadInterval"), InstructionHandler(OSPF6InterfaceDeadInterval))
p.Ctrie.Insert([]byte("OSPF6.Interface.HelloInterval"), InstructionHandler(OSPF6InterfaceHelloInterval))
p.Ctrie.Insert([]byte("OSPF6.Interface.RetransmitInterval"), InstructionHandler(OSPF6InterfaceRetransmitInterval))
p.Ctrie.Insert([]byte("OSPF6.Interface.TransmitDelay"), InstructionHandler(OSPF6InterfaceTransmitDelay))
p.Ctrie.Insert([]byte("OSPF6.Interface.Priority"), InstructionHandler(OSPF6InterfacePriority))
p.Ctrie.Insert([]byte("OSPF6.Interface.Networktype"), InstructionHandler(OSPF6InterfaceNetworktype))
p.Ctrie.Insert([]byte("No.OSPF6.Interface.Cost"), InstructionHandler(NoOSPF6InterfaceCost))
p.Ctrie.Insert([]byte("No.OSPF6.Interface.DeadInterval"), InstructionHandler(NoOSPF6InterfaceDeadInterval))
p.Ctrie.Insert([]byte("No.OSPF6.Interface.HelloInterval"), InstructionHandler(NoOSPF6InterfaceHelloInterval))
p.Ctrie.Insert([]byte("No.OSPF6.Interface.RetransmitInterval"), InstructionHandler(NoOSPF6InterfaceRetransmitInterval))
p.Ctrie.Insert([]byte("No.OSPF6.Interface.TransmitDelay"), InstructionHandler(NoOSPF6InterfaceTransmitDelay))
p.Ctrie.Insert([]byte("No.OSPF6.Interface.Priority"), InstructionHandler(NoOSPF6InterfacePriority))
p.Ctrie.Insert([]byte("No.OSPF6.Interface.Networktype"), InstructionHandler(NoOSPF6InterfaceNetworktype))
p.Ctrie.Insert([]byte("OSPF6.ReferenceBandwidth"), InstructionHandler(OSPF6ReferenceBandwidth))
p.Ctrie.Insert([]byte("No.OSPF6.ReferenceBandwidth"), InstructionHandler(NoOSPF6ReferenceBandwidth))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate"), InstructionHandler(OSPF6DefaultOriginate))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Routemap"), InstructionHandler(OSPF6DefaultOriginateRoutemap))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Metric"), InstructionHandler(OSPF6DefaultOriginateMetric))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Metrictype"), InstructionHandler(OSPF6DefaultOriginateMetrictype))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Metric.Metrictype"), InstructionHandler(OSPF6DefaultOriginateMetricMetrictype))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Metric.Routemap"), InstructionHandler(OSPF6DefaultOriginateMetricRoutemap))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Metrictype.Routemap"), InstructionHandler(OSPF6DefaultOriginateMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Metric.Metrictype.Routemap"), InstructionHandler(OSPF6DefaultOriginateMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Always"), InstructionHandler(OSPF6DefaultOriginateAlways))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Always.Routemap"), InstructionHandler(OSPF6DefaultOriginateAlwaysRoutemap))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Always.Metric"), InstructionHandler(OSPF6DefaultOriginateAlwaysMetric))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Always.Metrictype"), InstructionHandler(OSPF6DefaultOriginateAlwaysMetrictype))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Always.Metric.Metrictype"), InstructionHandler(OSPF6DefaultOriginateAlwaysMetricMetrictype))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Always.Metric.Routemap"), InstructionHandler(OSPF6DefaultOriginateAlwaysMetricRoutemap))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Always.Metrictype.Routemap"), InstructionHandler(OSPF6DefaultOriginateAlwaysMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF6.DefaultOriginate.Always.Metric.Metrictype.Routemap"), InstructionHandler(OSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate"), InstructionHandler(NoOSPF6DefaultOriginate))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Routemap"), InstructionHandler(NoOSPF6DefaultOriginateRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Metric"), InstructionHandler(NoOSPF6DefaultOriginateMetric))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Metrictype"), InstructionHandler(NoOSPF6DefaultOriginateMetrictype))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Metric.Metrictype"), InstructionHandler(NoOSPF6DefaultOriginateMetricMetrictype))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Metric.Routemap"), InstructionHandler(NoOSPF6DefaultOriginateMetricRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Metrictype.Routemap"), InstructionHandler(NoOSPF6DefaultOriginateMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Metric.Metrictype.Routemap"), InstructionHandler(NoOSPF6DefaultOriginateMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Always"), InstructionHandler(NoOSPF6DefaultOriginateAlways))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Always.Routemap"), InstructionHandler(NoOSPF6DefaultOriginateAlwaysRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Always.Metric"), InstructionHandler(NoOSPF6DefaultOriginateAlwaysMetric))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Always.Metrictype"), InstructionHandler(NoOSPF6DefaultOriginateAlwaysMetrictype))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Always.Metric.Metrictype"), InstructionHandler(NoOSPF6DefaultOriginateAlwaysMetricMetrictype))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Always.Metric.Routemap"), InstructionHandler(NoOSPF6DefaultOriginateAlwaysMetricRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Always.Metrictype.Routemap"), InstructionHandler(NoOSPF6DefaultOriginateAlwaysMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultOriginate.Always.Metric.Metrictype.Routemap"), InstructionHandler(NoOSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF6.Redistribute"), InstructionHandler(OSPF6Redistribute))
p.Ctrie.Insert([]byte("OSPF6.Redistribute.Metric"), InstructionHandler(OSPF6RedistributeMetric))
p.Ctrie.Insert([]byte("OSPF6.Redistribute.Metrictype"), InstructionHandler(OSPF6RedistributeMetrictype))
p.Ctrie.Insert([]byte("OSPF6.Redistribute.Routemap"), InstructionHandler(OSPF6RedistributeRoutemap))
p.Ctrie.Insert([]byte("OSPF6.Redistribute.Metric.Metrictype"), InstructionHandler(OSPF6RedistributeMetricMetrictype))
p.Ctrie.Insert([]byte("OSPF6.Redistribute.Metric.Routemap"), InstructionHandler(OSPF6RedistributeMetricRoutemap))
p.Ctrie.Insert([]byte("OSPF6.Redistribute.Metric.Metrictype.Routemap"), InstructionHandler(OSPF6RedistributeMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.Redistribute"), InstructionHandler(NoOSPF6Redistribute))
p.Ctrie.Insert([]byte("No.OSPF6.Redistribute.Metric"), InstructionHandler(NoOSPF6RedistributeMetric))
p.Ctrie.Insert([]byte("No.OSPF6.Redistribute.Metrictype"), InstructionHandler(NoOSPF6RedistributeMetrictype))
p.Ctrie.Insert([]byte("No.OSPF6.Redistribute.Routemap"), InstructionHandler(NoOSPF6RedistributeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.Redistribute.Metric.Metrictype"), InstructionHandler(NoOSPF6RedistributeMetricMetrictype))
p.Ctrie.Insert([]byte("No.OSPF6.Redistribute.Metric.Routemap"), InstructionHandler(NoOSPF6RedistributeMetricRoutemap))
p.Ctrie.Insert([]byte("No.OSPF6.Redistribute.Metric.Metrictype.Routemap"), InstructionHandler(NoOSPF6RedistributeMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF6.Summary"), InstructionHandler(OSPF6Summary))
p.Ctrie.Insert([]byte("OSPF6.Summary.NoAdvertise"), InstructionHandler(OSPF6SummaryNoAdvertise))
p.Ctrie.Insert([]byte("No.OSPF6.Summary"), InstructionHandler(NoOSPF6Summary))
p.Ctrie.Insert([]byte("No.OSPF6.Summary.NoAdvertise"), InstructionHandler(NoOSPF6SummaryNoAdvertise))
p.Ctrie.Insert([]byte("OSPF6.DefaultMetric"), InstructionHandler(OSPF6DefaultMetric))
p.Ctrie.Insert([]byte("No.OSPF6.DefaultMetric"), InstructionHandler(NoOSPF6DefaultMetric))
p.Ctrie.Insert([]byte("OSPF6.Passive"), InstructionHandler(OSPF6Passive))
p.Ctrie.Insert([]byte("No.OSPF6.Passive"), InstructionHandler(NoOSPF6Passive))
p.Ctrie.Insert([]byte("OSPF6.AdminDistance"), InstructionHandler(OSPF6AdminDistance))
p.Ctrie.Insert([]byte("No.OSPF6.AdminDistance"), InstructionHandler(NoOSPF6AdminDistance))
p.Ctrie.Insert([]byte("OSPF6.Distance.External"), InstructionHandler(OSPF6DistanceExternal))
p.Ctrie.Insert([]byte("OSPF6.Distance.Inter"), InstructionHandler(OSPF6DistanceInter))
p.Ctrie.Insert([]byte("OSPF6.Distance.Intra"), InstructionHandler(OSPF6DistanceIntra))
p.Ctrie.Insert([]byte("OSPF6.Distance.Inter.Intra"), InstructionHandler(OSPF6DistanceInterIntra))
p.Ctrie.Insert([]byte("OSPF6.Distance.Inter.External"), InstructionHandler(OSPF6DistanceInterExternal))
p.Ctrie.Insert([]byte("OSPF6.Distance.Inter.Intra.External"), InstructionHandler(OSPF6DistanceInterIntraExternal))
p.Ctrie.Insert([]byte("No.OSPF6.Distance.External"), InstructionHandler(NoOSPF6DistanceExternal))
p.Ctrie.Insert([]byte("No.OSPF6.Distance.Inter"), InstructionHandler(NoOSPF6DistanceInter))
p.Ctrie.Insert([]byte("No.OSPF6.Distance.Intra"), InstructionHandler(NoOSPF6DistanceIntra))
p.Ctrie.Insert([]byte("No.OSPF6.Distance.Inter.Intra"), InstructionHandler(NoOSPF6DistanceInterIntra))
p.Ctrie.Insert([]byte("No.OSPF6.Distance.Inter.External"), InstructionHandler(NoOSPF6DistanceInterExternal))
p.Ctrie.Insert([]byte("No.OSPF6.Distance.Inter.Intra.External"), InstructionHandler(NoOSPF6DistanceInterIntraExternal))
p.Ctrie.Insert([]byte("OSPF6.Distributelist.IN"), InstructionHandler(OSPF6DistributelistIN))
p.Ctrie.Insert([]byte("OSPF6.Distributelist.OUT"), InstructionHandler(OSPF6DistributelistOUT))
p.Ctrie.Insert([]byte("No.OSPF6.Distributelist.IN"), InstructionHandler(NoOSPF6DistributelistIN))
p.Ctrie.Insert([]byte("No.OSPF6.Distributelist.OUT"), InstructionHandler(NoOSPF6DistributelistOUT))
p.Ctrie.Insert([]byte("OSPF6.Area.DefaultCost"), InstructionHandler(OSPF6AreaDefaultCost))
p.Ctrie.Insert([]byte("No.OSPF6.Area.DefaultCost"), InstructionHandler(NoOSPF6AreaDefaultCost))
p.Ctrie.Insert([]byte("OSPF6.Area.NSSA"), InstructionHandler(OSPF6AreaNSSA))
p.Ctrie.Insert([]byte("OSPF6.Area.NSSA.DefaultOriginate"), InstructionHandler(OSPF6AreaNSSADefaultOriginate))
p.Ctrie.Insert([]byte("OSPF6.Area.NSSA.NoRedistribution"), InstructionHandler(OSPF6AreaNSSANoRedistribution))
p.Ctrie.Insert([]byte("OSPF6.Area.NSSA.NoSummary"), InstructionHandler(OSPF6AreaNSSANoSummary))
p.Ctrie.Insert([]byte("OSPF6.Area.NSSA.StabilityInterval"), InstructionHandler(OSPF6AreaNSSAStabilityInterval))
p.Ctrie.Insert([]byte("OSPF6.Area.Translatorrole"), InstructionHandler(OSPF6AreaTranslatorrole))
p.Ctrie.Insert([]byte("No.OSPF6.Area.NSSA"), InstructionHandler(NoOSPF6AreaNSSA))
p.Ctrie.Insert([]byte("No.OSPF6.Area.NSSA.DefaultOriginate"), InstructionHandler(NoOSPF6AreaNSSADefaultOriginate))
p.Ctrie.Insert([]byte("No.OSPF6.Area.NSSA.NoRedistribution"), InstructionHandler(NoOSPF6AreaNSSANoRedistribution))
p.Ctrie.Insert([]byte("No.OSPF6.Area.NSSA.NoSummary"), InstructionHandler(NoOSPF6AreaNSSANoSummary))
p.Ctrie.Insert([]byte("No.OSPF6.Area.NSSA.StabilityInterval"), InstructionHandler(NoOSPF6AreaNSSAStabilityInterval))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Translatorrole"), InstructionHandler(NoOSPF6AreaTranslatorrole))
p.Ctrie.Insert([]byte("OSPF6.Area.Stub"), InstructionHandler(OSPF6AreaStub))
p.Ctrie.Insert([]byte("OSPF6.Area.Stub.NoSummary"), InstructionHandler(OSPF6AreaStubNoSummary))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Stub"), InstructionHandler(NoOSPF6AreaStub))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Stub.NoSummary"), InstructionHandler(NoOSPF6AreaStubNoSummary))
p.Ctrie.Insert([]byte("OSPF6.Area.Range"), InstructionHandler(OSPF6AreaRange))
p.Ctrie.Insert([]byte("OSPF6.Area.Range.Advertise"), InstructionHandler(OSPF6AreaRangeAdvertise))
p.Ctrie.Insert([]byte("OSPF6.Area.Range.NoAdvertise"), InstructionHandler(OSPF6AreaRangeNoAdvertise))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Range"), InstructionHandler(NoOSPF6AreaRange))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Range.Advertise"), InstructionHandler(NoOSPF6AreaRangeAdvertise))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Range.NoAdvertise"), InstructionHandler(NoOSPF6AreaRangeNoAdvertise))
p.Ctrie.Insert([]byte("OSPF6.Area.Virtuallink"), InstructionHandler(OSPF6AreaVirtuallink))
p.Ctrie.Insert([]byte("OSPF6.Area.Virtuallink.DeadInterval"), InstructionHandler(OSPF6AreaVirtuallinkDeadInterval))
p.Ctrie.Insert([]byte("OSPF6.Area.Virtuallink.HelloInterval"), InstructionHandler(OSPF6AreaVirtuallinkHelloInterval))
p.Ctrie.Insert([]byte("OSPF6.Area.Virtuallink.Instanceid"), InstructionHandler(OSPF6AreaVirtuallinkInstanceid))
p.Ctrie.Insert([]byte("OSPF6.Area.Virtuallink.RetransmitInterval"), InstructionHandler(OSPF6AreaVirtuallinkRetransmitInterval))
p.Ctrie.Insert([]byte("OSPF6.Area.Virtuallink.TransmitDelay"), InstructionHandler(OSPF6AreaVirtuallinkTransmitDelay))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Virtuallink"), InstructionHandler(NoOSPF6AreaVirtuallink))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Virtuallink.DeadInterval"), InstructionHandler(NoOSPF6AreaVirtuallinkDeadInterval))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Virtuallink.HelloInterval"), InstructionHandler(NoOSPF6AreaVirtuallinkHelloInterval))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Virtuallink.Instanceid"), InstructionHandler(NoOSPF6AreaVirtuallinkInstanceid))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Virtuallink.RetransmitInterval"), InstructionHandler(NoOSPF6AreaVirtuallinkRetransmitInterval))
p.Ctrie.Insert([]byte("No.OSPF6.Area.Virtuallink.TransmitDelay"), InstructionHandler(NoOSPF6AreaVirtuallinkTransmitDelay))
p.Ctrie.Insert([]byte("OSPF"), InstructionHandler(OSPF))
p.Ctrie.Insert([]byte("No.OSPF"), InstructionHandler(NoOSPF))
p.Ctrie.Insert([]byte("OSPF.Rid"), InstructionHandler(OSPFRid))
p.Ctrie.Insert([]byte("No.OSPF.Rid"), InstructionHandler(NoOSPFRid))
p.Ctrie.Insert([]byte("OSPF.Network.Area"), InstructionHandler(OSPFNetworkArea))
p.Ctrie.Insert([]byte("No.OSPF.Network.Area"), InstructionHandler(NoOSPFNetworkArea))
p.Ctrie.Insert([]byte("OSPF.Interface.Cost"), InstructionHandler(OSPFInterfaceCost))
p.Ctrie.Insert([]byte("OSPF.Interface.DeadInterval"), InstructionHandler(OSPFInterfaceDeadInterval))
p.Ctrie.Insert([]byte("OSPF.Interface.HelloInterval"), InstructionHandler(OSPFInterfaceHelloInterval))
p.Ctrie.Insert([]byte("OSPF.Interface.RetransmitInterval"), InstructionHandler(OSPFInterfaceRetransmitInterval))
p.Ctrie.Insert([]byte("OSPF.Interface.TransmitDelay"), InstructionHandler(OSPFInterfaceTransmitDelay))
p.Ctrie.Insert([]byte("OSPF.Interface.Priority"), InstructionHandler(OSPFInterfacePriority))
p.Ctrie.Insert([]byte("OSPF.Interface.Networktype"), InstructionHandler(OSPFInterfaceNetworktype))
p.Ctrie.Insert([]byte("No.OSPF.Interface.Cost"), InstructionHandler(NoOSPFInterfaceCost))
p.Ctrie.Insert([]byte("No.OSPF.Interface.DeadInterval"), InstructionHandler(NoOSPFInterfaceDeadInterval))
p.Ctrie.Insert([]byte("No.OSPF.Interface.HelloInterval"), InstructionHandler(NoOSPFInterfaceHelloInterval))
p.Ctrie.Insert([]byte("No.OSPF.Interface.RetransmitInterval"), InstructionHandler(NoOSPFInterfaceRetransmitInterval))
p.Ctrie.Insert([]byte("No.OSPF.Interface.TransmitDelay"), InstructionHandler(NoOSPFInterfaceTransmitDelay))
p.Ctrie.Insert([]byte("No.OSPF.Interface.Priority"), InstructionHandler(NoOSPFInterfacePriority))
p.Ctrie.Insert([]byte("No.OSPF.Interface.Networktype"), InstructionHandler(NoOSPFInterfaceNetworktype))
p.Ctrie.Insert([]byte("OSPF.ReferenceBandwidth"), InstructionHandler(OSPFReferenceBandwidth))
p.Ctrie.Insert([]byte("No.OSPF.ReferenceBandwidth"), InstructionHandler(NoOSPFReferenceBandwidth))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate"), InstructionHandler(OSPFDefaultOriginate))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Routemap"), InstructionHandler(OSPFDefaultOriginateRoutemap))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Metric"), InstructionHandler(OSPFDefaultOriginateMetric))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Metrictype"), InstructionHandler(OSPFDefaultOriginateMetrictype))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Metric.Metrictype"), InstructionHandler(OSPFDefaultOriginateMetricMetrictype))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Metric.Routemap"), InstructionHandler(OSPFDefaultOriginateMetricRoutemap))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Metrictype.Routemap"), InstructionHandler(OSPFDefaultOriginateMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Metric.Metrictype.Routemap"), InstructionHandler(OSPFDefaultOriginateMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Always"), InstructionHandler(OSPFDefaultOriginateAlways))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Always.Routemap"), InstructionHandler(OSPFDefaultOriginateAlwaysRoutemap))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Always.Metric"), InstructionHandler(OSPFDefaultOriginateAlwaysMetric))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Always.Metrictype"), InstructionHandler(OSPFDefaultOriginateAlwaysMetrictype))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Always.Metric.Metrictype"), InstructionHandler(OSPFDefaultOriginateAlwaysMetricMetrictype))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Always.Metric.Routemap"), InstructionHandler(OSPFDefaultOriginateAlwaysMetricRoutemap))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Always.Metrictype.Routemap"), InstructionHandler(OSPFDefaultOriginateAlwaysMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF.DefaultOriginate.Always.Metric.Metrictype.Routemap"), InstructionHandler(OSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate"), InstructionHandler(NoOSPFDefaultOriginate))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Routemap"), InstructionHandler(NoOSPFDefaultOriginateRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Metric"), InstructionHandler(NoOSPFDefaultOriginateMetric))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Metrictype"), InstructionHandler(NoOSPFDefaultOriginateMetrictype))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Metric.Metrictype"), InstructionHandler(NoOSPFDefaultOriginateMetricMetrictype))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Metric.Routemap"), InstructionHandler(NoOSPFDefaultOriginateMetricRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Metrictype.Routemap"), InstructionHandler(NoOSPFDefaultOriginateMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Metric.Metrictype.Routemap"), InstructionHandler(NoOSPFDefaultOriginateMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Always"), InstructionHandler(NoOSPFDefaultOriginateAlways))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Always.Routemap"), InstructionHandler(NoOSPFDefaultOriginateAlwaysRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Always.Metric"), InstructionHandler(NoOSPFDefaultOriginateAlwaysMetric))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Always.Metrictype"), InstructionHandler(NoOSPFDefaultOriginateAlwaysMetrictype))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Always.Metric.Metrictype"), InstructionHandler(NoOSPFDefaultOriginateAlwaysMetricMetrictype))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Always.Metric.Routemap"), InstructionHandler(NoOSPFDefaultOriginateAlwaysMetricRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Always.Metrictype.Routemap"), InstructionHandler(NoOSPFDefaultOriginateAlwaysMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.DefaultOriginate.Always.Metric.Metrictype.Routemap"), InstructionHandler(NoOSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF.Redistribute"), InstructionHandler(OSPFRedistribute))
p.Ctrie.Insert([]byte("OSPF.Redistribute.Metric"), InstructionHandler(OSPFRedistributeMetric))
p.Ctrie.Insert([]byte("OSPF.Redistribute.Metrictype"), InstructionHandler(OSPFRedistributeMetrictype))
p.Ctrie.Insert([]byte("OSPF.Redistribute.Routemap"), InstructionHandler(OSPFRedistributeRoutemap))
p.Ctrie.Insert([]byte("OSPF.Redistribute.Metric.Metrictype"), InstructionHandler(OSPFRedistributeMetricMetrictype))
p.Ctrie.Insert([]byte("OSPF.Redistribute.Metric.Routemap"), InstructionHandler(OSPFRedistributeMetricRoutemap))
p.Ctrie.Insert([]byte("OSPF.Redistribute.Metric.Metrictype.Routemap"), InstructionHandler(OSPFRedistributeMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.Redistribute"), InstructionHandler(NoOSPFRedistribute))
p.Ctrie.Insert([]byte("No.OSPF.Redistribute.Metric"), InstructionHandler(NoOSPFRedistributeMetric))
p.Ctrie.Insert([]byte("No.OSPF.Redistribute.Metrictype"), InstructionHandler(NoOSPFRedistributeMetrictype))
p.Ctrie.Insert([]byte("No.OSPF.Redistribute.Routemap"), InstructionHandler(NoOSPFRedistributeRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.Redistribute.Metric.Metrictype"), InstructionHandler(NoOSPFRedistributeMetricMetrictype))
p.Ctrie.Insert([]byte("No.OSPF.Redistribute.Metric.Routemap"), InstructionHandler(NoOSPFRedistributeMetricRoutemap))
p.Ctrie.Insert([]byte("No.OSPF.Redistribute.Metric.Metrictype.Routemap"), InstructionHandler(NoOSPFRedistributeMetricMetrictypeRoutemap))
p.Ctrie.Insert([]byte("OSPF.Summary"), InstructionHandler(OSPFSummary))
p.Ctrie.Insert([]byte("OSPF.Summary.NoAdvertise"), InstructionHandler(OSPFSummaryNoAdvertise))
p.Ctrie.Insert([]byte("No.OSPF.Summary"), InstructionHandler(NoOSPFSummary))
p.Ctrie.Insert([]byte("No.OSPF.Summary.NoAdvertise"), InstructionHandler(NoOSPFSummaryNoAdvertise))
p.Ctrie.Insert([]byte("OSPF.DefaultMetric"), InstructionHandler(OSPFDefaultMetric))
p.Ctrie.Insert([]byte("No.OSPF.DefaultMetric"), InstructionHandler(NoOSPFDefaultMetric))
p.Ctrie.Insert([]byte("OSPF.Passive"), InstructionHandler(OSPFPassive))
p.Ctrie.Insert([]byte("No.OSPF.Passive"), InstructionHandler(NoOSPFPassive))
p.Ctrie.Insert([]byte("OSPF.AdminDistance"), InstructionHandler(OSPFAdminDistance))
p.Ctrie.Insert([]byte("No.OSPF.AdminDistance"), InstructionHandler(NoOSPFAdminDistance))
p.Ctrie.Insert([]byte("OSPF.Distance.External"), InstructionHandler(OSPFDistanceExternal))
p.Ctrie.Insert([]byte("OSPF.Distance.Inter"), InstructionHandler(OSPFDistanceInter))
p.Ctrie.Insert([]byte("OSPF.Distance.Intra"), InstructionHandler(OSPFDistanceIntra))
p.Ctrie.Insert([]byte("OSPF.Distance.Inter.Intra"), InstructionHandler(OSPFDistanceInterIntra))
p.Ctrie.Insert([]byte("OSPF.Distance.Inter.External"), InstructionHandler(OSPFDistanceInterExternal))
p.Ctrie.Insert([]byte("OSPF.Distance.Inter.Intra.External"), InstructionHandler(OSPFDistanceInterIntraExternal))
p.Ctrie.Insert([]byte("No.OSPF.Distance.External"), InstructionHandler(NoOSPFDistanceExternal))
p.Ctrie.Insert([]byte("No.OSPF.Distance.Inter"), InstructionHandler(NoOSPFDistanceInter))
p.Ctrie.Insert([]byte("No.OSPF.Distance.Intra"), InstructionHandler(NoOSPFDistanceIntra))
p.Ctrie.Insert([]byte("No.OSPF.Distance.Inter.Intra"), InstructionHandler(NoOSPFDistanceInterIntra))
p.Ctrie.Insert([]byte("No.OSPF.Distance.Inter.External"), InstructionHandler(NoOSPFDistanceInterExternal))
p.Ctrie.Insert([]byte("No.OSPF.Distance.Inter.Intra.External"), InstructionHandler(NoOSPFDistanceInterIntraExternal))
p.Ctrie.Insert([]byte("OSPF.Distributelist.IN"), InstructionHandler(OSPFDistributelistIN))
p.Ctrie.Insert([]byte("OSPF.Distributelist.OUT"), InstructionHandler(OSPFDistributelistOUT))
p.Ctrie.Insert([]byte("No.OSPF.Distributelist.IN"), InstructionHandler(NoOSPFDistributelistIN))
p.Ctrie.Insert([]byte("No.OSPF.Distributelist.OUT"), InstructionHandler(NoOSPFDistributelistOUT))
p.Ctrie.Insert([]byte("OSPF.Area.DefaultCost"), InstructionHandler(OSPFAreaDefaultCost))
p.Ctrie.Insert([]byte("No.OSPF.Area.DefaultCost"), InstructionHandler(NoOSPFAreaDefaultCost))
p.Ctrie.Insert([]byte("OSPF.Area.NSSA"), InstructionHandler(OSPFAreaNSSA))
p.Ctrie.Insert([]byte("OSPF.Area.NSSA.DefaultOriginate"), InstructionHandler(OSPFAreaNSSADefaultOriginate))
p.Ctrie.Insert([]byte("OSPF.Area.NSSA.NoRedistribution"), InstructionHandler(OSPFAreaNSSANoRedistribution))
p.Ctrie.Insert([]byte("OSPF.Area.NSSA.NoSummary"), InstructionHandler(OSPFAreaNSSANoSummary))
p.Ctrie.Insert([]byte("OSPF.Area.NSSA.StabilityInterval"), InstructionHandler(OSPFAreaNSSAStabilityInterval))
p.Ctrie.Insert([]byte("OSPF.Area.Translatorrole"), InstructionHandler(OSPFAreaTranslatorrole))
p.Ctrie.Insert([]byte("No.OSPF.Area.NSSA"), InstructionHandler(NoOSPFAreaNSSA))
p.Ctrie.Insert([]byte("No.OSPF.Area.NSSA.DefaultOriginate"), InstructionHandler(NoOSPFAreaNSSADefaultOriginate))
p.Ctrie.Insert([]byte("No.OSPF.Area.NSSA.NoRedistribution"), InstructionHandler(NoOSPFAreaNSSANoRedistribution))
p.Ctrie.Insert([]byte("No.OSPF.Area.NSSA.NoSummary"), InstructionHandler(NoOSPFAreaNSSANoSummary))
p.Ctrie.Insert([]byte("No.OSPF.Area.NSSA.StabilityInterval"), InstructionHandler(NoOSPFAreaNSSAStabilityInterval))
p.Ctrie.Insert([]byte("No.OSPF.Area.Translatorrole"), InstructionHandler(NoOSPFAreaTranslatorrole))
p.Ctrie.Insert([]byte("OSPF.Area.Stub"), InstructionHandler(OSPFAreaStub))
p.Ctrie.Insert([]byte("OSPF.Area.Stub.NoSummary"), InstructionHandler(OSPFAreaStubNoSummary))
p.Ctrie.Insert([]byte("No.OSPF.Area.Stub"), InstructionHandler(NoOSPFAreaStub))
p.Ctrie.Insert([]byte("No.OSPF.Area.Stub.NoSummary"), InstructionHandler(NoOSPFAreaStubNoSummary))
p.Ctrie.Insert([]byte("OSPF.Area.Range"), InstructionHandler(OSPFAreaRange))
p.Ctrie.Insert([]byte("OSPF.Area.Range.Advertise"), InstructionHandler(OSPFAreaRangeAdvertise))
p.Ctrie.Insert([]byte("OSPF.Area.Range.NoAdvertise"), InstructionHandler(OSPFAreaRangeNoAdvertise))
p.Ctrie.Insert([]byte("No.OSPF.Area.Range"), InstructionHandler(NoOSPFAreaRange))
p.Ctrie.Insert([]byte("No.OSPF.Area.Range.Advertise"), InstructionHandler(NoOSPFAreaRangeAdvertise))
p.Ctrie.Insert([]byte("No.OSPF.Area.Range.NoAdvertise"), InstructionHandler(NoOSPFAreaRangeNoAdvertise))
p.Ctrie.Insert([]byte("OSPF.Area.Virtuallink"), InstructionHandler(OSPFAreaVirtuallink))
p.Ctrie.Insert([]byte("OSPF.Area.Virtuallink.DeadInterval"), InstructionHandler(OSPFAreaVirtuallinkDeadInterval))
p.Ctrie.Insert([]byte("OSPF.Area.Virtuallink.HelloInterval"), InstructionHandler(OSPFAreaVirtuallinkHelloInterval))
p.Ctrie.Insert([]byte("OSPF.Area.Virtuallink.Instanceid"), InstructionHandler(OSPFAreaVirtuallinkInstanceid))
p.Ctrie.Insert([]byte("OSPF.Area.Virtuallink.RetransmitInterval"), InstructionHandler(OSPFAreaVirtuallinkRetransmitInterval))
p.Ctrie.Insert([]byte("OSPF.Area.Virtuallink.TransmitDelay"), InstructionHandler(OSPFAreaVirtuallinkTransmitDelay))
p.Ctrie.Insert([]byte("No.OSPF.Area.Virtuallink"), InstructionHandler(NoOSPFAreaVirtuallink))
p.Ctrie.Insert([]byte("No.OSPF.Area.Virtuallink.DeadInterval"), InstructionHandler(NoOSPFAreaVirtuallinkDeadInterval))
p.Ctrie.Insert([]byte("No.OSPF.Area.Virtuallink.HelloInterval"), InstructionHandler(NoOSPFAreaVirtuallinkHelloInterval))
p.Ctrie.Insert([]byte("No.OSPF.Area.Virtuallink.Instanceid"), InstructionHandler(NoOSPFAreaVirtuallinkInstanceid))
p.Ctrie.Insert([]byte("No.OSPF.Area.Virtuallink.RetransmitInterval"), InstructionHandler(NoOSPFAreaVirtuallinkRetransmitInterval))
p.Ctrie.Insert([]byte("No.OSPF.Area.Virtuallink.TransmitDelay"), InstructionHandler(NoOSPFAreaVirtuallinkTransmitDelay))
}
