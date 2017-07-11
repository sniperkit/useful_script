package dsl

import (
	"command"
)

type Switch interface {
	PortEnable(Port, Enable string) []*command.Command
	PortDisable(Port, Disable string) []*command.Command
	PortPvid(Port, Pvid string) []*command.Command
	PortSpeed(Port, Speed string) []*command.Command
	VLAN(VLAN string) []*command.Command
	NoVLAN(VLAN string) []*command.Command
	VLANAdd(VLAN, Add string) []*command.Command
	VLANAddT(VLAN, AddT string) []*command.Command
	VLANDel(VLAN, Del string) []*command.Command
	VLANDelT(VLAN, DelT string) []*command.Command
	VLANShutdown(VLAN, Shutdown string) []*command.Command
	VLANNoShutdown(VLAN, NoShutdown string) []*command.Command
	VLANIP(VLAN, IP string) []*command.Command
	NoVLANIP(VLAN, IP string) []*command.Command
	VLANIP2(VLAN, IP2 string) []*command.Command
	NoVLANIP2(VLAN, IP2 string) []*command.Command
	VLANAddTIP(VLAN, AddT, IP string) []*command.Command
	VLANAddTIP2(VLAN, AddT, IP2 string) []*command.Command
	VLANAddTIPNoShutdown(VLAN, AddT, IP, NoShutdown string) []*command.Command
	VLANAddTIP2NoShutdown(VLAN, AddT, IP2, NoShutdown string) []*command.Command
	VLANAddUTIP(VLAN, AddUT, IP string) []*command.Command
	VLANAddUTIP2(VLAN, AddUT, IP2 string) []*command.Command
	VLANAddUTIPNoShutdown(VLAN, AddUT, IP, NoShutdown string) []*command.Command
	VLANAddUTIP2NoShutdown(VLAN, AddUT, IP2, NoShutdown string) []*command.Command
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
}
