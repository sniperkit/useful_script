package dsl

import (
	"command"
)

type Switch interface {
	PortSlotTypeEnable(Port, Slot, Type, Enable string) []*command.Command
	PortSlotTypeDisable(Port, Slot, Type, Disable string) []*command.Command
	PortSlotTypeSpeed(Port, Slot, Type, Speed string) []*command.Command
	PortSlotTypePvid(Port, Slot, Type, Pvid string) []*command.Command
	VLAN(VLAN string) []*command.Command
	NoVLAN(VLAN string) []*command.Command
	VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port string) []*command.Command
	VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port string) []*command.Command
	VLANDelTypeSlotPort(VLAN, Del, Type, Slot, Port string) []*command.Command
	VLANDelTTypeSlotPort(VLAN, DelT, Type, Slot, Port string) []*command.Command
	VLANShutdown(VLAN, Shutdown string) []*command.Command
	VLANNoShutdown(VLAN, NoShutdown string) []*command.Command
	VLANIP(VLAN, IP string) []*command.Command
	NoVLANIP(VLAN, IP string) []*command.Command
	VLANIP2(VLAN, IP2 string) []*command.Command
	NoVLANIP2(VLAN, IP2 string) []*command.Command
	VLANAddTypeSlotPortIP(VLAN, Add, Type, Slot, Port, IP string) []*command.Command
	VLANAddTypeSlotPortIP2(VLAN, Add, Type, Slot, Port, IP2 string) []*command.Command
	VLANAddTypeSlotPortIPNoShutdown(VLAN, Add, Type, Slot, Port, IP, NoShutdown string) []*command.Command
	VLANAddTypeSlotPortIP2NoShutdown(VLAN, Add, Type, Slot, Port, IP2, NoShutdown string) []*command.Command
	VLANAddTTypeSlotPortIP(VLAN, AddT, Type, Slot, Port, IP string) []*command.Command
	VLANAddTTypeSlotPortIP2(VLAN, AddT, Type, Slot, Port, IP2 string) []*command.Command
	VLANAddTTypeSlotPortIPNoShutdown(VLAN, AddT, Type, Slot, Port, IP, NoShutdown string) []*command.Command
	VLANAddTTypeSlotPortIP2NoShutdown(VLAN, AddT, Type, Slot, Port, IP2, NoShutdown string) []*command.Command
	VLANIP6Enable(VLAN, IP6, Enable string) []*command.Command
	NoVLANIP6Enable(VLAN, IP6, Enable string) []*command.Command
	VLANIP6(VLAN, IP6 string) []*command.Command
	NoVLANIP6(VLAN, IP6 string) []*command.Command
	VLANIP6LL(VLAN, IP6LL string) []*command.Command
	VLANIP6LLIP6(VLAN, IP6LL, IP6 string) []*command.Command
	NoVLANIP6LL(VLAN, IP6LL string) []*command.Command
	OSPF6(OSPF6 string) []*command.Command
	NoOSPF6(OSPF6 string) []*command.Command
	OSPF6Rid(OSPF6, Rid string) []*command.Command
	NoOSPF6Rid(OSPF6, Rid string) []*command.Command
	OSPF6InterfaceArea(OSPF6, Interface, Area string) []*command.Command
	NoOSPF6InterfaceArea(OSPF6, Interface, Area string) []*command.Command
	OSPF6InterfaceCost(OSPF6, Interface, Cost string) []*command.Command
	OSPF6InterfaceDeadInterval(OSPF6, Interface, DeadInterval string) []*command.Command
	OSPF6InterfaceHelloInterval(OSPF6, Interface, HelloInterval string) []*command.Command
	OSPF6InterfaceRetransmitInterval(OSPF6, Interface, RetransmitInterval string) []*command.Command
	OSPF6InterfaceTransmitDelay(OSPF6, Interface, TransmitDelay string) []*command.Command
	OSPF6InterfacePriority(OSPF6, Interface, Priority string) []*command.Command
	OSPF6InterfaceNetworktype(OSPF6, Interface, Networktype string) []*command.Command
	NoOSPF6InterfaceCost(OSPF6, Interface, Cost string) []*command.Command
	NoOSPF6InterfaceDeadInterval(OSPF6, Interface, DeadInterval string) []*command.Command
	NoOSPF6InterfaceHelloInterval(OSPF6, Interface, HelloInterval string) []*command.Command
	NoOSPF6InterfaceRetransmitInterval(OSPF6, Interface, RetransmitInterval string) []*command.Command
	NoOSPF6InterfaceTransmitDelay(OSPF6, Interface, TransmitDelay string) []*command.Command
	NoOSPF6InterfacePriority(OSPF6, Interface, Priority string) []*command.Command
	NoOSPF6InterfaceNetworktype(OSPF6, Interface, Networktype string) []*command.Command
	OSPF6ReferenceBandwidth(OSPF6, ReferenceBandwidth string) []*command.Command
	NoOSPF6ReferenceBandwidth(OSPF6, ReferenceBandwidth string) []*command.Command
	OSPF6DefaultOriginate(OSPF6, DefaultOriginate string) []*command.Command
	OSPF6DefaultOriginateRoutemap(OSPF6, DefaultOriginate, Routemap string) []*command.Command
	OSPF6DefaultOriginateMetric(OSPF6, DefaultOriginate, Metric string) []*command.Command
	OSPF6DefaultOriginateMetrictype(OSPF6, DefaultOriginate, Metrictype string) []*command.Command
	OSPF6DefaultOriginateMetricMetrictype(OSPF6, DefaultOriginate, Metric, Metrictype string) []*command.Command
	OSPF6DefaultOriginateMetricRoutemap(OSPF6, DefaultOriginate, Metric, Routemap string) []*command.Command
	OSPF6DefaultOriginateMetrictypeRoutemap(OSPF6, DefaultOriginate, Metrictype, Routemap string) []*command.Command
	OSPF6DefaultOriginateMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command
	OSPF6DefaultOriginateAlways(OSPF6, DefaultOriginate, Always string) []*command.Command
	OSPF6DefaultOriginateAlwaysRoutemap(OSPF6, DefaultOriginate, Always, Routemap string) []*command.Command
	OSPF6DefaultOriginateAlwaysMetric(OSPF6, DefaultOriginate, Always, Metric string) []*command.Command
	OSPF6DefaultOriginateAlwaysMetrictype(OSPF6, DefaultOriginate, Always, Metrictype string) []*command.Command
	OSPF6DefaultOriginateAlwaysMetricMetrictype(OSPF6, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command
	OSPF6DefaultOriginateAlwaysMetricRoutemap(OSPF6, DefaultOriginate, Always, Metric, Routemap string) []*command.Command
	OSPF6DefaultOriginateAlwaysMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command
	OSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command
	NoOSPF6DefaultOriginate(OSPF6, DefaultOriginate string) []*command.Command
	NoOSPF6DefaultOriginateRoutemap(OSPF6, DefaultOriginate, Routemap string) []*command.Command
	NoOSPF6DefaultOriginateMetric(OSPF6, DefaultOriginate, Metric string) []*command.Command
	NoOSPF6DefaultOriginateMetrictype(OSPF6, DefaultOriginate, Metrictype string) []*command.Command
	NoOSPF6DefaultOriginateMetricMetrictype(OSPF6, DefaultOriginate, Metric, Metrictype string) []*command.Command
	NoOSPF6DefaultOriginateMetricRoutemap(OSPF6, DefaultOriginate, Metric, Routemap string) []*command.Command
	NoOSPF6DefaultOriginateMetrictypeRoutemap(OSPF6, DefaultOriginate, Metrictype, Routemap string) []*command.Command
	NoOSPF6DefaultOriginateMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command
	NoOSPF6DefaultOriginateAlways(OSPF6, DefaultOriginate, Always string) []*command.Command
	NoOSPF6DefaultOriginateAlwaysRoutemap(OSPF6, DefaultOriginate, Always, Routemap string) []*command.Command
	NoOSPF6DefaultOriginateAlwaysMetric(OSPF6, DefaultOriginate, Always, Metric string) []*command.Command
	NoOSPF6DefaultOriginateAlwaysMetrictype(OSPF6, DefaultOriginate, Always, Metrictype string) []*command.Command
	NoOSPF6DefaultOriginateAlwaysMetricMetrictype(OSPF6, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command
	NoOSPF6DefaultOriginateAlwaysMetricRoutemap(OSPF6, DefaultOriginate, Always, Metric, Routemap string) []*command.Command
	NoOSPF6DefaultOriginateAlwaysMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command
	NoOSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command
	OSPF6Redistribute(OSPF6, Redistribute string) []*command.Command
	OSPF6RedistributeMetric(OSPF6, Redistribute, Metric string) []*command.Command
	OSPF6RedistributeMetrictype(OSPF6, Redistribute, Metrictype string) []*command.Command
	OSPF6RedistributeRoutemap(OSPF6, Redistribute, Routemap string) []*command.Command
	OSPF6RedistributeMetricMetrictype(OSPF6, Redistribute, Metric, Metrictype string) []*command.Command
	OSPF6RedistributeMetricRoutemap(OSPF6, Redistribute, Metric, Routemap string) []*command.Command
	OSPF6RedistributeMetricMetrictypeRoutemap(OSPF6, Redistribute, Metric, Metrictype, Routemap string) []*command.Command
	NoOSPF6Redistribute(OSPF6, Redistribute string) []*command.Command
	NoOSPF6RedistributeMetric(OSPF6, Redistribute, Metric string) []*command.Command
	NoOSPF6RedistributeMetrictype(OSPF6, Redistribute, Metrictype string) []*command.Command
	NoOSPF6RedistributeRoutemap(OSPF6, Redistribute, Routemap string) []*command.Command
	NoOSPF6RedistributeMetricMetrictype(OSPF6, Redistribute, Metric, Metrictype string) []*command.Command
	NoOSPF6RedistributeMetricRoutemap(OSPF6, Redistribute, Metric, Routemap string) []*command.Command
	NoOSPF6RedistributeMetricMetrictypeRoutemap(OSPF6, Redistribute, Metric, Metrictype, Routemap string) []*command.Command
	OSPF6Summary(OSPF6, Summary string) []*command.Command
	OSPF6SummaryNoAdvertise(OSPF6, Summary, NoAdvertise string) []*command.Command
	NoOSPF6Summary(OSPF6, Summary string) []*command.Command
	NoOSPF6SummaryNoAdvertise(OSPF6, Summary, NoAdvertise string) []*command.Command
	OSPF6DefaultMetric(OSPF6, DefaultMetric string) []*command.Command
	NoOSPF6DefaultMetric(OSPF6, DefaultMetric string) []*command.Command
	OSPF6Passive(OSPF6, Passive string) []*command.Command
	NoOSPF6Passive(OSPF6, Passive string) []*command.Command
	OSPF6AdminDistance(OSPF6, AdminDistance string) []*command.Command
	NoOSPF6AdminDistance(OSPF6, AdminDistance string) []*command.Command
	OSPF6DistanceExternal(OSPF6, Distance, External string) []*command.Command
	OSPF6DistanceInter(OSPF6, Distance, Inter string) []*command.Command
	OSPF6DistanceIntra(OSPF6, Distance, Intra string) []*command.Command
	OSPF6DistanceInterIntra(OSPF6, Distance, Inter, Intra string) []*command.Command
	OSPF6DistanceInterExternal(OSPF6, Distance, Inter, External string) []*command.Command
	OSPF6DistanceInterIntraExternal(OSPF6, Distance, Inter, Intra, External string) []*command.Command
	NoOSPF6DistanceExternal(OSPF6, Distance, External string) []*command.Command
	NoOSPF6DistanceInter(OSPF6, Distance, Inter string) []*command.Command
	NoOSPF6DistanceIntra(OSPF6, Distance, Intra string) []*command.Command
	NoOSPF6DistanceInterIntra(OSPF6, Distance, Inter, Intra string) []*command.Command
	NoOSPF6DistanceInterExternal(OSPF6, Distance, Inter, External string) []*command.Command
	NoOSPF6DistanceInterIntraExternal(OSPF6, Distance, Inter, Intra, External string) []*command.Command
	OSPF6DistributelistIN(OSPF6, Distributelist, IN string) []*command.Command
	OSPF6DistributelistOUT(OSPF6, Distributelist, OUT string) []*command.Command
	NoOSPF6DistributelistIN(OSPF6, Distributelist, IN string) []*command.Command
	NoOSPF6DistributelistOUT(OSPF6, Distributelist, OUT string) []*command.Command
	OSPF6AreaDefaultCost(OSPF6, Area, DefaultCost string) []*command.Command
	NoOSPF6AreaDefaultCost(OSPF6, Area, DefaultCost string) []*command.Command
	OSPF6AreaNSSA(OSPF6, Area, NSSA string) []*command.Command
	OSPF6AreaNSSADefaultOriginate(OSPF6, Area, NSSA, DefaultOriginate string) []*command.Command
	OSPF6AreaNSSANoRedistribution(OSPF6, Area, NSSA, NoRedistribution string) []*command.Command
	OSPF6AreaNSSANoSummary(OSPF6, Area, NSSA, NoSummary string) []*command.Command
	OSPF6AreaNSSAStabilityInterval(OSPF6, Area, NSSA, StabilityInterval string) []*command.Command
	OSPF6AreaTranslatorrole(OSPF6, Area, Translatorrole string) []*command.Command
	NoOSPF6AreaNSSA(OSPF6, Area, NSSA string) []*command.Command
	NoOSPF6AreaNSSADefaultOriginate(OSPF6, Area, NSSA, DefaultOriginate string) []*command.Command
	NoOSPF6AreaNSSANoRedistribution(OSPF6, Area, NSSA, NoRedistribution string) []*command.Command
	NoOSPF6AreaNSSANoSummary(OSPF6, Area, NSSA, NoSummary string) []*command.Command
	NoOSPF6AreaNSSAStabilityInterval(OSPF6, Area, NSSA, StabilityInterval string) []*command.Command
	NoOSPF6AreaTranslatorrole(OSPF6, Area, Translatorrole string) []*command.Command
	OSPF6AreaStub(OSPF6, Area, Stub string) []*command.Command
	OSPF6AreaStubNoSummary(OSPF6, Area, Stub, NoSummary string) []*command.Command
	NoOSPF6AreaStub(OSPF6, Area, Stub string) []*command.Command
	NoOSPF6AreaStubNoSummary(OSPF6, Area, Stub, NoSummary string) []*command.Command
	OSPF6AreaRange(OSPF6, Area, Range string) []*command.Command
	OSPF6AreaRangeAdvertise(OSPF6, Area, Range, Advertise string) []*command.Command
	OSPF6AreaRangeNoAdvertise(OSPF6, Area, Range, NoAdvertise string) []*command.Command
	NoOSPF6AreaRange(OSPF6, Area, Range string) []*command.Command
	NoOSPF6AreaRangeAdvertise(OSPF6, Area, Range, Advertise string) []*command.Command
	NoOSPF6AreaRangeNoAdvertise(OSPF6, Area, Range, NoAdvertise string) []*command.Command
	OSPF6AreaVirtuallink(OSPF6, Area, Virtuallink string) []*command.Command
	OSPF6AreaVirtuallinkDeadInterval(OSPF6, Area, Virtuallink, DeadInterval string) []*command.Command
	OSPF6AreaVirtuallinkHelloInterval(OSPF6, Area, Virtuallink, HelloInterval string) []*command.Command
	OSPF6AreaVirtuallinkInstanceid(OSPF6, Area, Virtuallink, Instanceid string) []*command.Command
	OSPF6AreaVirtuallinkRetransmitInterval(OSPF6, Area, Virtuallink, RetransmitInterval string) []*command.Command
	OSPF6AreaVirtuallinkTransmitDelay(OSPF6, Area, Virtuallink, TransmitDelay string) []*command.Command
	NoOSPF6AreaVirtuallink(OSPF6, Area, Virtuallink string) []*command.Command
	NoOSPF6AreaVirtuallinkDeadInterval(OSPF6, Area, Virtuallink, DeadInterval string) []*command.Command
	NoOSPF6AreaVirtuallinkHelloInterval(OSPF6, Area, Virtuallink, HelloInterval string) []*command.Command
	NoOSPF6AreaVirtuallinkInstanceid(OSPF6, Area, Virtuallink, Instanceid string) []*command.Command
	NoOSPF6AreaVirtuallinkRetransmitInterval(OSPF6, Area, Virtuallink, RetransmitInterval string) []*command.Command
	NoOSPF6AreaVirtuallinkTransmitDelay(OSPF6, Area, Virtuallink, TransmitDelay string) []*command.Command
	OSPF(OSPF string) []*command.Command
	NoOSPF(OSPF string) []*command.Command
	OSPFRid(OSPF, Rid string) []*command.Command
	NoOSPFRid(OSPF, Rid string) []*command.Command
	OSPFNetworkArea(OSPF, Network, Area string) []*command.Command
	NoOSPFNetworkArea(OSPF, Network, Area string) []*command.Command
	OSPFInterfaceCost(OSPF, Interface, Cost string) []*command.Command
	OSPFInterfaceDeadInterval(OSPF, Interface, DeadInterval string) []*command.Command
	OSPFInterfaceHelloInterval(OSPF, Interface, HelloInterval string) []*command.Command
	OSPFInterfaceRetransmitInterval(OSPF, Interface, RetransmitInterval string) []*command.Command
	OSPFInterfaceTransmitDelay(OSPF, Interface, TransmitDelay string) []*command.Command
	OSPFInterfacePriority(OSPF, Interface, Priority string) []*command.Command
	OSPFInterfaceNetworktype(OSPF, Interface, Networktype string) []*command.Command
	NoOSPFInterfaceCost(OSPF, Interface, Cost string) []*command.Command
	NoOSPFInterfaceDeadInterval(OSPF, Interface, DeadInterval string) []*command.Command
	NoOSPFInterfaceHelloInterval(OSPF, Interface, HelloInterval string) []*command.Command
	NoOSPFInterfaceRetransmitInterval(OSPF, Interface, RetransmitInterval string) []*command.Command
	NoOSPFInterfaceTransmitDelay(OSPF, Interface, TransmitDelay string) []*command.Command
	NoOSPFInterfacePriority(OSPF, Interface, Priority string) []*command.Command
	NoOSPFInterfaceNetworktype(OSPF, Interface, Networktype string) []*command.Command
	OSPFReferenceBandwidth(OSPF, ReferenceBandwidth string) []*command.Command
	NoOSPFReferenceBandwidth(OSPF, ReferenceBandwidth string) []*command.Command
	OSPFDefaultOriginate(OSPF, DefaultOriginate string) []*command.Command
	OSPFDefaultOriginateRoutemap(OSPF, DefaultOriginate, Routemap string) []*command.Command
	OSPFDefaultOriginateMetric(OSPF, DefaultOriginate, Metric string) []*command.Command
	OSPFDefaultOriginateMetrictype(OSPF, DefaultOriginate, Metrictype string) []*command.Command
	OSPFDefaultOriginateMetricMetrictype(OSPF, DefaultOriginate, Metric, Metrictype string) []*command.Command
	OSPFDefaultOriginateMetricRoutemap(OSPF, DefaultOriginate, Metric, Routemap string) []*command.Command
	OSPFDefaultOriginateMetrictypeRoutemap(OSPF, DefaultOriginate, Metrictype, Routemap string) []*command.Command
	OSPFDefaultOriginateMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command
	OSPFDefaultOriginateAlways(OSPF, DefaultOriginate, Always string) []*command.Command
	OSPFDefaultOriginateAlwaysRoutemap(OSPF, DefaultOriginate, Always, Routemap string) []*command.Command
	OSPFDefaultOriginateAlwaysMetric(OSPF, DefaultOriginate, Always, Metric string) []*command.Command
	OSPFDefaultOriginateAlwaysMetrictype(OSPF, DefaultOriginate, Always, Metrictype string) []*command.Command
	OSPFDefaultOriginateAlwaysMetricMetrictype(OSPF, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command
	OSPFDefaultOriginateAlwaysMetricRoutemap(OSPF, DefaultOriginate, Always, Metric, Routemap string) []*command.Command
	OSPFDefaultOriginateAlwaysMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command
	OSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command
	NoOSPFDefaultOriginate(OSPF, DefaultOriginate string) []*command.Command
	NoOSPFDefaultOriginateRoutemap(OSPF, DefaultOriginate, Routemap string) []*command.Command
	NoOSPFDefaultOriginateMetric(OSPF, DefaultOriginate, Metric string) []*command.Command
	NoOSPFDefaultOriginateMetrictype(OSPF, DefaultOriginate, Metrictype string) []*command.Command
	NoOSPFDefaultOriginateMetricMetrictype(OSPF, DefaultOriginate, Metric, Metrictype string) []*command.Command
	NoOSPFDefaultOriginateMetricRoutemap(OSPF, DefaultOriginate, Metric, Routemap string) []*command.Command
	NoOSPFDefaultOriginateMetrictypeRoutemap(OSPF, DefaultOriginate, Metrictype, Routemap string) []*command.Command
	NoOSPFDefaultOriginateMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command
	NoOSPFDefaultOriginateAlways(OSPF, DefaultOriginate, Always string) []*command.Command
	NoOSPFDefaultOriginateAlwaysRoutemap(OSPF, DefaultOriginate, Always, Routemap string) []*command.Command
	NoOSPFDefaultOriginateAlwaysMetric(OSPF, DefaultOriginate, Always, Metric string) []*command.Command
	NoOSPFDefaultOriginateAlwaysMetrictype(OSPF, DefaultOriginate, Always, Metrictype string) []*command.Command
	NoOSPFDefaultOriginateAlwaysMetricMetrictype(OSPF, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command
	NoOSPFDefaultOriginateAlwaysMetricRoutemap(OSPF, DefaultOriginate, Always, Metric, Routemap string) []*command.Command
	NoOSPFDefaultOriginateAlwaysMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command
	NoOSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command
	OSPFRedistribute(OSPF, Redistribute string) []*command.Command
	OSPFRedistributeMetric(OSPF, Redistribute, Metric string) []*command.Command
	OSPFRedistributeMetrictype(OSPF, Redistribute, Metrictype string) []*command.Command
	OSPFRedistributeRoutemap(OSPF, Redistribute, Routemap string) []*command.Command
	OSPFRedistributeMetricMetrictype(OSPF, Redistribute, Metric, Metrictype string) []*command.Command
	OSPFRedistributeMetricRoutemap(OSPF, Redistribute, Metric, Routemap string) []*command.Command
	OSPFRedistributeMetricMetrictypeRoutemap(OSPF, Redistribute, Metric, Metrictype, Routemap string) []*command.Command
	NoOSPFRedistribute(OSPF, Redistribute string) []*command.Command
	NoOSPFRedistributeMetric(OSPF, Redistribute, Metric string) []*command.Command
	NoOSPFRedistributeMetrictype(OSPF, Redistribute, Metrictype string) []*command.Command
	NoOSPFRedistributeRoutemap(OSPF, Redistribute, Routemap string) []*command.Command
	NoOSPFRedistributeMetricMetrictype(OSPF, Redistribute, Metric, Metrictype string) []*command.Command
	NoOSPFRedistributeMetricRoutemap(OSPF, Redistribute, Metric, Routemap string) []*command.Command
	NoOSPFRedistributeMetricMetrictypeRoutemap(OSPF, Redistribute, Metric, Metrictype, Routemap string) []*command.Command
	OSPFSummary(OSPF, Summary string) []*command.Command
	OSPFSummaryNoAdvertise(OSPF, Summary, NoAdvertise string) []*command.Command
	NoOSPFSummary(OSPF, Summary string) []*command.Command
	NoOSPFSummaryNoAdvertise(OSPF, Summary, NoAdvertise string) []*command.Command
	OSPFDefaultMetric(OSPF, DefaultMetric string) []*command.Command
	NoOSPFDefaultMetric(OSPF, DefaultMetric string) []*command.Command
	OSPFPassive(OSPF, Passive string) []*command.Command
	NoOSPFPassive(OSPF, Passive string) []*command.Command
	OSPFAdminDistance(OSPF, AdminDistance string) []*command.Command
	NoOSPFAdminDistance(OSPF, AdminDistance string) []*command.Command
	OSPFDistanceExternal(OSPF, Distance, External string) []*command.Command
	OSPFDistanceInter(OSPF, Distance, Inter string) []*command.Command
	OSPFDistanceIntra(OSPF, Distance, Intra string) []*command.Command
	OSPFDistanceInterIntra(OSPF, Distance, Inter, Intra string) []*command.Command
	OSPFDistanceInterExternal(OSPF, Distance, Inter, External string) []*command.Command
	OSPFDistanceInterIntraExternal(OSPF, Distance, Inter, Intra, External string) []*command.Command
	NoOSPFDistanceExternal(OSPF, Distance, External string) []*command.Command
	NoOSPFDistanceInter(OSPF, Distance, Inter string) []*command.Command
	NoOSPFDistanceIntra(OSPF, Distance, Intra string) []*command.Command
	NoOSPFDistanceInterIntra(OSPF, Distance, Inter, Intra string) []*command.Command
	NoOSPFDistanceInterExternal(OSPF, Distance, Inter, External string) []*command.Command
	NoOSPFDistanceInterIntraExternal(OSPF, Distance, Inter, Intra, External string) []*command.Command
	OSPFDistributelistIN(OSPF, Distributelist, IN string) []*command.Command
	OSPFDistributelistOUT(OSPF, Distributelist, OUT string) []*command.Command
	NoOSPFDistributelistIN(OSPF, Distributelist, IN string) []*command.Command
	NoOSPFDistributelistOUT(OSPF, Distributelist, OUT string) []*command.Command
	OSPFAreaDefaultCost(OSPF, Area, DefaultCost string) []*command.Command
	NoOSPFAreaDefaultCost(OSPF, Area, DefaultCost string) []*command.Command
	OSPFAreaNSSA(OSPF, Area, NSSA string) []*command.Command
	OSPFAreaNSSADefaultOriginate(OSPF, Area, NSSA, DefaultOriginate string) []*command.Command
	OSPFAreaNSSANoRedistribution(OSPF, Area, NSSA, NoRedistribution string) []*command.Command
	OSPFAreaNSSANoSummary(OSPF, Area, NSSA, NoSummary string) []*command.Command
	OSPFAreaNSSAStabilityInterval(OSPF, Area, NSSA, StabilityInterval string) []*command.Command
	OSPFAreaTranslatorrole(OSPF, Area, Translatorrole string) []*command.Command
	NoOSPFAreaNSSA(OSPF, Area, NSSA string) []*command.Command
	NoOSPFAreaNSSADefaultOriginate(OSPF, Area, NSSA, DefaultOriginate string) []*command.Command
	NoOSPFAreaNSSANoRedistribution(OSPF, Area, NSSA, NoRedistribution string) []*command.Command
	NoOSPFAreaNSSANoSummary(OSPF, Area, NSSA, NoSummary string) []*command.Command
	NoOSPFAreaNSSAStabilityInterval(OSPF, Area, NSSA, StabilityInterval string) []*command.Command
	NoOSPFAreaTranslatorrole(OSPF, Area, Translatorrole string) []*command.Command
	OSPFAreaStub(OSPF, Area, Stub string) []*command.Command
	OSPFAreaStubNoSummary(OSPF, Area, Stub, NoSummary string) []*command.Command
	NoOSPFAreaStub(OSPF, Area, Stub string) []*command.Command
	NoOSPFAreaStubNoSummary(OSPF, Area, Stub, NoSummary string) []*command.Command
	OSPFAreaRange(OSPF, Area, Range string) []*command.Command
	OSPFAreaRangeAdvertise(OSPF, Area, Range, Advertise string) []*command.Command
	OSPFAreaRangeNoAdvertise(OSPF, Area, Range, NoAdvertise string) []*command.Command
	NoOSPFAreaRange(OSPF, Area, Range string) []*command.Command
	NoOSPFAreaRangeAdvertise(OSPF, Area, Range, Advertise string) []*command.Command
	NoOSPFAreaRangeNoAdvertise(OSPF, Area, Range, NoAdvertise string) []*command.Command
	OSPFAreaVirtuallink(OSPF, Area, Virtuallink string) []*command.Command
	OSPFAreaVirtuallinkDeadInterval(OSPF, Area, Virtuallink, DeadInterval string) []*command.Command
	OSPFAreaVirtuallinkHelloInterval(OSPF, Area, Virtuallink, HelloInterval string) []*command.Command
	OSPFAreaVirtuallinkInstanceid(OSPF, Area, Virtuallink, Instanceid string) []*command.Command
	OSPFAreaVirtuallinkRetransmitInterval(OSPF, Area, Virtuallink, RetransmitInterval string) []*command.Command
	OSPFAreaVirtuallinkTransmitDelay(OSPF, Area, Virtuallink, TransmitDelay string) []*command.Command
	NoOSPFAreaVirtuallink(OSPF, Area, Virtuallink string) []*command.Command
	NoOSPFAreaVirtuallinkDeadInterval(OSPF, Area, Virtuallink, DeadInterval string) []*command.Command
	NoOSPFAreaVirtuallinkHelloInterval(OSPF, Area, Virtuallink, HelloInterval string) []*command.Command
	NoOSPFAreaVirtuallinkInstanceid(OSPF, Area, Virtuallink, Instanceid string) []*command.Command
	NoOSPFAreaVirtuallinkRetransmitInterval(OSPF, Area, Virtuallink, RetransmitInterval string) []*command.Command
	NoOSPFAreaVirtuallinkTransmitDelay(OSPF, Area, Virtuallink, TransmitDelay string) []*command.Command
}
