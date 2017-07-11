package dsl

import (
	"command"
)

type Instruction struct {
	Name   string
	Switch Switch
}

func PortEnable(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["Port"]; !ok {
		return nil
	}
	if _, ok := in["Enable"]; !ok {
		return nil
	}
	return sw.PortEnable(in["Port"], in["Enable"])
}

func PortDisable(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["Port"]; !ok {
		return nil
	}
	if _, ok := in["Disable"]; !ok {
		return nil
	}
	return sw.PortDisable(in["Port"], in["Disable"])
}

func PortPvid(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["Port"]; !ok {
		return nil
	}
	if _, ok := in["Pvid"]; !ok {
		return nil
	}
	return sw.PortPvid(in["Port"], in["Pvid"])
}

func PortSpeed(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["Port"]; !ok {
		return nil
	}
	if _, ok := in["Speed"]; !ok {
		return nil
	}
	return sw.PortSpeed(in["Port"], in["Speed"])
}

func VLAN(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 1 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	return sw.VLAN(in["VLAN"])
}

func NoVLAN(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	return sw.NoVLAN(in["VLAN"])
}

func VLANAdd(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["Add"]; !ok {
		return nil
	}
	return sw.VLANAdd(in["VLAN"], in["Add"])
}

func VLANAddT(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["AddT"]; !ok {
		return nil
	}
	return sw.VLANAddT(in["VLAN"], in["AddT"])
}

func VLANDel(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["Del"]; !ok {
		return nil
	}
	return sw.VLANDel(in["VLAN"], in["Del"])
}

func VLANDelT(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["DelT"]; !ok {
		return nil
	}
	return sw.VLANDelT(in["VLAN"], in["DelT"])
}

func VLANShutdown(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["Shutdown"]; !ok {
		return nil
	}
	return sw.VLANShutdown(in["VLAN"], in["Shutdown"])
}

func VLANNoShutdown(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["NoShutdown"]; !ok {
		return nil
	}
	return sw.VLANNoShutdown(in["VLAN"], in["NoShutdown"])
}

func VLANIP(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP"]; !ok {
		return nil
	}
	return sw.VLANIP(in["VLAN"], in["IP"])
}

func NoVLANIP(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP"]; !ok {
		return nil
	}
	return sw.NoVLANIP(in["VLAN"], in["IP"])
}

func VLANIP2(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP2"]; !ok {
		return nil
	}
	return sw.VLANIP2(in["VLAN"], in["IP2"])
}

func NoVLANIP2(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP2"]; !ok {
		return nil
	}
	return sw.NoVLANIP2(in["VLAN"], in["IP2"])
}

func VLANAddTIP(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["AddT"]; !ok {
		return nil
	}
	if _, ok := in["IP"]; !ok {
		return nil
	}
	return sw.VLANAddTIP(in["VLAN"], in["AddT"], in["IP"])
}

func VLANAddTIP2(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["AddT"]; !ok {
		return nil
	}
	if _, ok := in["IP2"]; !ok {
		return nil
	}
	return sw.VLANAddTIP2(in["VLAN"], in["AddT"], in["IP2"])
}

func VLANAddTIPNoShutdown(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["AddT"]; !ok {
		return nil
	}
	if _, ok := in["IP"]; !ok {
		return nil
	}
	if _, ok := in["NoShutdown"]; !ok {
		return nil
	}
	return sw.VLANAddTIPNoShutdown(in["VLAN"], in["AddT"], in["IP"], in["NoShutdown"])
}

func VLANAddTIP2NoShutdown(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["AddT"]; !ok {
		return nil
	}
	if _, ok := in["IP2"]; !ok {
		return nil
	}
	if _, ok := in["NoShutdown"]; !ok {
		return nil
	}
	return sw.VLANAddTIP2NoShutdown(in["VLAN"], in["AddT"], in["IP2"], in["NoShutdown"])
}

func VLANAddUTIP(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["AddUT"]; !ok {
		return nil
	}
	if _, ok := in["IP"]; !ok {
		return nil
	}
	return sw.VLANAddUTIP(in["VLAN"], in["AddUT"], in["IP"])
}

func VLANAddUTIP2(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["AddUT"]; !ok {
		return nil
	}
	if _, ok := in["IP2"]; !ok {
		return nil
	}
	return sw.VLANAddUTIP2(in["VLAN"], in["AddUT"], in["IP2"])
}

func VLANAddUTIPNoShutdown(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["AddUT"]; !ok {
		return nil
	}
	if _, ok := in["IP"]; !ok {
		return nil
	}
	if _, ok := in["NoShutdown"]; !ok {
		return nil
	}
	return sw.VLANAddUTIPNoShutdown(in["VLAN"], in["AddUT"], in["IP"], in["NoShutdown"])
}

func VLANAddUTIP2NoShutdown(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["AddUT"]; !ok {
		return nil
	}
	if _, ok := in["IP2"]; !ok {
		return nil
	}
	if _, ok := in["NoShutdown"]; !ok {
		return nil
	}
	return sw.VLANAddUTIP2NoShutdown(in["VLAN"], in["AddUT"], in["IP2"], in["NoShutdown"])
}

func VLANIP6Enable(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP6"]; !ok {
		return nil
	}
	if _, ok := in["Enable"]; !ok {
		return nil
	}
	return sw.VLANIP6Enable(in["VLAN"], in["IP6"], in["Enable"])
}

func NoVLANIP6Enable(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP6"]; !ok {
		return nil
	}
	if _, ok := in["Enable"]; !ok {
		return nil
	}
	return sw.NoVLANIP6Enable(in["VLAN"], in["IP6"], in["Enable"])
}

func VLANIP6(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP6"]; !ok {
		return nil
	}
	return sw.VLANIP6(in["VLAN"], in["IP6"])
}

func NoVLANIP6(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP6"]; !ok {
		return nil
	}
	return sw.NoVLANIP6(in["VLAN"], in["IP6"])
}

func VLANIP6LL(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP6LL"]; !ok {
		return nil
	}
	return sw.VLANIP6LL(in["VLAN"], in["IP6LL"])
}

func VLANIP6LLIP6(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP6LL"]; !ok {
		return nil
	}
	if _, ok := in["IP6"]; !ok {
		return nil
	}
	return sw.VLANIP6LLIP6(in["VLAN"], in["IP6LL"], in["IP6"])
}

func NoVLANIP6LL(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["VLAN"]; !ok {
		return nil
	}
	if _, ok := in["IP6LL"]; !ok {
		return nil
	}
	return sw.NoVLANIP6LL(in["VLAN"], in["IP6LL"])
}

func OSPF6(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 1 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	return sw.OSPF6(in["OSPF6"])
}

func NoOSPF6(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	return sw.NoOSPF6(in["OSPF6"])
}

func OSPF6Rid(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Rid"]; !ok {
		return nil
	}
	return sw.OSPF6Rid(in["OSPF6"], in["Rid"])
}

func NoOSPF6Rid(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Rid"]; !ok {
		return nil
	}
	return sw.NoOSPF6Rid(in["OSPF6"], in["Rid"])
}

func OSPF6InterfaceArea(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	return sw.OSPF6InterfaceArea(in["OSPF6"], in["Interface"], in["Area"])
}

func NoOSPF6InterfaceArea(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	return sw.NoOSPF6InterfaceArea(in["OSPF6"], in["Interface"], in["Area"])
}

func OSPF6InterfaceCost(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["Cost"]; !ok {
		return nil
	}
	return sw.OSPF6InterfaceCost(in["OSPF6"], in["Interface"], in["Cost"])
}

func OSPF6InterfaceDeadInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["DeadInterval"]; !ok {
		return nil
	}
	return sw.OSPF6InterfaceDeadInterval(in["OSPF6"], in["Interface"], in["DeadInterval"])
}

func OSPF6InterfaceHelloInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["HelloInterval"]; !ok {
		return nil
	}
	return sw.OSPF6InterfaceHelloInterval(in["OSPF6"], in["Interface"], in["HelloInterval"])
}

func OSPF6InterfaceRetransmitInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["RetransmitInterval"]; !ok {
		return nil
	}
	return sw.OSPF6InterfaceRetransmitInterval(in["OSPF6"], in["Interface"], in["RetransmitInterval"])
}

func OSPF6InterfaceTransmitDelay(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["TransmitDelay"]; !ok {
		return nil
	}
	return sw.OSPF6InterfaceTransmitDelay(in["OSPF6"], in["Interface"], in["TransmitDelay"])
}

func OSPF6InterfacePriority(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["Priority"]; !ok {
		return nil
	}
	return sw.OSPF6InterfacePriority(in["OSPF6"], in["Interface"], in["Priority"])
}

func OSPF6InterfaceNetworktype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["Networktype"]; !ok {
		return nil
	}
	return sw.OSPF6InterfaceNetworktype(in["OSPF6"], in["Interface"], in["Networktype"])
}

func NoOSPF6InterfaceCost(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["Cost"]; !ok {
		return nil
	}
	return sw.NoOSPF6InterfaceCost(in["OSPF6"], in["Interface"], in["Cost"])
}

func NoOSPF6InterfaceDeadInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["DeadInterval"]; !ok {
		return nil
	}
	return sw.NoOSPF6InterfaceDeadInterval(in["OSPF6"], in["Interface"], in["DeadInterval"])
}

func NoOSPF6InterfaceHelloInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["HelloInterval"]; !ok {
		return nil
	}
	return sw.NoOSPF6InterfaceHelloInterval(in["OSPF6"], in["Interface"], in["HelloInterval"])
}

func NoOSPF6InterfaceRetransmitInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["RetransmitInterval"]; !ok {
		return nil
	}
	return sw.NoOSPF6InterfaceRetransmitInterval(in["OSPF6"], in["Interface"], in["RetransmitInterval"])
}

func NoOSPF6InterfaceTransmitDelay(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["TransmitDelay"]; !ok {
		return nil
	}
	return sw.NoOSPF6InterfaceTransmitDelay(in["OSPF6"], in["Interface"], in["TransmitDelay"])
}

func NoOSPF6InterfacePriority(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["Priority"]; !ok {
		return nil
	}
	return sw.NoOSPF6InterfacePriority(in["OSPF6"], in["Interface"], in["Priority"])
}

func NoOSPF6InterfaceNetworktype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Interface"]; !ok {
		return nil
	}
	if _, ok := in["Networktype"]; !ok {
		return nil
	}
	return sw.NoOSPF6InterfaceNetworktype(in["OSPF6"], in["Interface"], in["Networktype"])
}

func OSPF6ReferenceBandwidth(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["ReferenceBandwidth"]; !ok {
		return nil
	}
	return sw.OSPF6ReferenceBandwidth(in["OSPF6"], in["ReferenceBandwidth"])
}

func NoOSPF6ReferenceBandwidth(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["ReferenceBandwidth"]; !ok {
		return nil
	}
	return sw.NoOSPF6ReferenceBandwidth(in["OSPF6"], in["ReferenceBandwidth"])
}

func OSPF6DefaultOriginate(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginate(in["OSPF6"], in["DefaultOriginate"])
}

func OSPF6DefaultOriginateRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Routemap"])
}

func OSPF6DefaultOriginateMetric(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateMetric(in["OSPF6"], in["DefaultOriginate"], in["Metric"])
}

func OSPF6DefaultOriginateMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Metrictype"])
}

func OSPF6DefaultOriginateMetricMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateMetricMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Metrictype"])
}

func OSPF6DefaultOriginateMetricRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateMetricRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Routemap"])
}

func OSPF6DefaultOriginateMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metrictype"], in["Routemap"])
}

func OSPF6DefaultOriginateMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateMetricMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func OSPF6DefaultOriginateAlways(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateAlways(in["OSPF6"], in["DefaultOriginate"], in["Always"])
}

func OSPF6DefaultOriginateAlwaysRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateAlwaysRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Routemap"])
}

func OSPF6DefaultOriginateAlwaysMetric(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateAlwaysMetric(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"])
}

func OSPF6DefaultOriginateAlwaysMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateAlwaysMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metrictype"])
}

func OSPF6DefaultOriginateAlwaysMetricMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateAlwaysMetricMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"])
}

func OSPF6DefaultOriginateAlwaysMetricRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateAlwaysMetricRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Routemap"])
}

func OSPF6DefaultOriginateAlwaysMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateAlwaysMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metrictype"], in["Routemap"])
}

func OSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 6 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6DefaultOriginate(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginate(in["OSPF6"], in["DefaultOriginate"])
}

func NoOSPF6DefaultOriginateRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Routemap"])
}

func NoOSPF6DefaultOriginateMetric(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateMetric(in["OSPF6"], in["DefaultOriginate"], in["Metric"])
}

func NoOSPF6DefaultOriginateMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Metrictype"])
}

func NoOSPF6DefaultOriginateMetricMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateMetricMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Metrictype"])
}

func NoOSPF6DefaultOriginateMetricRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateMetricRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Routemap"])
}

func NoOSPF6DefaultOriginateMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6DefaultOriginateMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 6 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateMetricMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6DefaultOriginateAlways(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateAlways(in["OSPF6"], in["DefaultOriginate"], in["Always"])
}

func NoOSPF6DefaultOriginateAlwaysRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateAlwaysRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Routemap"])
}

func NoOSPF6DefaultOriginateAlwaysMetric(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateAlwaysMetric(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"])
}

func NoOSPF6DefaultOriginateAlwaysMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateAlwaysMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metrictype"])
}

func NoOSPF6DefaultOriginateAlwaysMetricMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 6 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateAlwaysMetricMetrictype(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"])
}

func NoOSPF6DefaultOriginateAlwaysMetricRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 6 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateAlwaysMetricRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Routemap"])
}

func NoOSPF6DefaultOriginateAlwaysMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 6 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateAlwaysMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 7 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	if _, ok := in["Always"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(in["OSPF6"], in["DefaultOriginate"], in["Always"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func OSPF6Redistribute(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	return sw.OSPF6Redistribute(in["OSPF6"], in["Redistribute"])
}

func OSPF6RedistributeMetric(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	return sw.OSPF6RedistributeMetric(in["OSPF6"], in["Redistribute"], in["Metric"])
}

func OSPF6RedistributeMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.OSPF6RedistributeMetrictype(in["OSPF6"], in["Redistribute"], in["Metrictype"])
}

func OSPF6RedistributeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6RedistributeRoutemap(in["OSPF6"], in["Redistribute"], in["Routemap"])
}

func OSPF6RedistributeMetricMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.OSPF6RedistributeMetricMetrictype(in["OSPF6"], in["Redistribute"], in["Metric"], in["Metrictype"])
}

func OSPF6RedistributeMetricRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6RedistributeMetricRoutemap(in["OSPF6"], in["Redistribute"], in["Metric"], in["Routemap"])
}

func OSPF6RedistributeMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.OSPF6RedistributeMetricMetrictypeRoutemap(in["OSPF6"], in["Redistribute"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func NoOSPF6Redistribute(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	return sw.NoOSPF6Redistribute(in["OSPF6"], in["Redistribute"])
}

func NoOSPF6RedistributeMetric(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	return sw.NoOSPF6RedistributeMetric(in["OSPF6"], in["Redistribute"], in["Metric"])
}

func NoOSPF6RedistributeMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.NoOSPF6RedistributeMetrictype(in["OSPF6"], in["Redistribute"], in["Metrictype"])
}

func NoOSPF6RedistributeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6RedistributeRoutemap(in["OSPF6"], in["Redistribute"], in["Routemap"])
}

func NoOSPF6RedistributeMetricMetrictype(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	return sw.NoOSPF6RedistributeMetricMetrictype(in["OSPF6"], in["Redistribute"], in["Metric"], in["Metrictype"])
}

func NoOSPF6RedistributeMetricRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6RedistributeMetricRoutemap(in["OSPF6"], in["Redistribute"], in["Metric"], in["Routemap"])
}

func NoOSPF6RedistributeMetricMetrictypeRoutemap(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 6 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Redistribute"]; !ok {
		return nil
	}
	if _, ok := in["Metric"]; !ok {
		return nil
	}
	if _, ok := in["Metrictype"]; !ok {
		return nil
	}
	if _, ok := in["Routemap"]; !ok {
		return nil
	}
	return sw.NoOSPF6RedistributeMetricMetrictypeRoutemap(in["OSPF6"], in["Redistribute"], in["Metric"], in["Metrictype"], in["Routemap"])
}

func OSPF6Summary(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Summary"]; !ok {
		return nil
	}
	return sw.OSPF6Summary(in["OSPF6"], in["Summary"])
}

func OSPF6SummaryNoAdvertise(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Summary"]; !ok {
		return nil
	}
	if _, ok := in["NoAdvertise"]; !ok {
		return nil
	}
	return sw.OSPF6SummaryNoAdvertise(in["OSPF6"], in["Summary"], in["NoAdvertise"])
}

func NoOSPF6Summary(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Summary"]; !ok {
		return nil
	}
	return sw.NoOSPF6Summary(in["OSPF6"], in["Summary"])
}

func NoOSPF6SummaryNoAdvertise(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Summary"]; !ok {
		return nil
	}
	if _, ok := in["NoAdvertise"]; !ok {
		return nil
	}
	return sw.NoOSPF6SummaryNoAdvertise(in["OSPF6"], in["Summary"], in["NoAdvertise"])
}

func OSPF6DefaultMetric(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultMetric"]; !ok {
		return nil
	}
	return sw.OSPF6DefaultMetric(in["OSPF6"], in["DefaultMetric"])
}

func NoOSPF6DefaultMetric(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["DefaultMetric"]; !ok {
		return nil
	}
	return sw.NoOSPF6DefaultMetric(in["OSPF6"], in["DefaultMetric"])
}

func OSPF6Passive(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Passive"]; !ok {
		return nil
	}
	return sw.OSPF6Passive(in["OSPF6"], in["Passive"])
}

func NoOSPF6Passive(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Passive"]; !ok {
		return nil
	}
	return sw.NoOSPF6Passive(in["OSPF6"], in["Passive"])
}

func OSPF6AdminDistance(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 2 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["AdminDistance"]; !ok {
		return nil
	}
	return sw.OSPF6AdminDistance(in["OSPF6"], in["AdminDistance"])
}

func NoOSPF6AdminDistance(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["AdminDistance"]; !ok {
		return nil
	}
	return sw.NoOSPF6AdminDistance(in["OSPF6"], in["AdminDistance"])
}

func OSPF6DistanceExternal(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["External"]; !ok {
		return nil
	}
	return sw.OSPF6DistanceExternal(in["OSPF6"], in["Distance"], in["External"])
}

func OSPF6DistanceInter(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Inter"]; !ok {
		return nil
	}
	return sw.OSPF6DistanceInter(in["OSPF6"], in["Distance"], in["Inter"])
}

func OSPF6DistanceIntra(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Intra"]; !ok {
		return nil
	}
	return sw.OSPF6DistanceIntra(in["OSPF6"], in["Distance"], in["Intra"])
}

func OSPF6DistanceInterIntra(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Inter"]; !ok {
		return nil
	}
	if _, ok := in["Intra"]; !ok {
		return nil
	}
	return sw.OSPF6DistanceInterIntra(in["OSPF6"], in["Distance"], in["Inter"], in["Intra"])
}

func OSPF6DistanceInterExternal(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Inter"]; !ok {
		return nil
	}
	if _, ok := in["External"]; !ok {
		return nil
	}
	return sw.OSPF6DistanceInterExternal(in["OSPF6"], in["Distance"], in["Inter"], in["External"])
}

func OSPF6DistanceInterIntraExternal(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Inter"]; !ok {
		return nil
	}
	if _, ok := in["Intra"]; !ok {
		return nil
	}
	if _, ok := in["External"]; !ok {
		return nil
	}
	return sw.OSPF6DistanceInterIntraExternal(in["OSPF6"], in["Distance"], in["Inter"], in["Intra"], in["External"])
}

func NoOSPF6DistanceExternal(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["External"]; !ok {
		return nil
	}
	return sw.NoOSPF6DistanceExternal(in["OSPF6"], in["Distance"], in["External"])
}

func NoOSPF6DistanceInter(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Inter"]; !ok {
		return nil
	}
	return sw.NoOSPF6DistanceInter(in["OSPF6"], in["Distance"], in["Inter"])
}

func NoOSPF6DistanceIntra(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Intra"]; !ok {
		return nil
	}
	return sw.NoOSPF6DistanceIntra(in["OSPF6"], in["Distance"], in["Intra"])
}

func NoOSPF6DistanceInterIntra(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Inter"]; !ok {
		return nil
	}
	if _, ok := in["Intra"]; !ok {
		return nil
	}
	return sw.NoOSPF6DistanceInterIntra(in["OSPF6"], in["Distance"], in["Inter"], in["Intra"])
}

func NoOSPF6DistanceInterExternal(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Inter"]; !ok {
		return nil
	}
	if _, ok := in["External"]; !ok {
		return nil
	}
	return sw.NoOSPF6DistanceInterExternal(in["OSPF6"], in["Distance"], in["Inter"], in["External"])
}

func NoOSPF6DistanceInterIntraExternal(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 6 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distance"]; !ok {
		return nil
	}
	if _, ok := in["Inter"]; !ok {
		return nil
	}
	if _, ok := in["Intra"]; !ok {
		return nil
	}
	if _, ok := in["External"]; !ok {
		return nil
	}
	return sw.NoOSPF6DistanceInterIntraExternal(in["OSPF6"], in["Distance"], in["Inter"], in["Intra"], in["External"])
}

func OSPF6DistributelistIN(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distributelist"]; !ok {
		return nil
	}
	if _, ok := in["IN"]; !ok {
		return nil
	}
	return sw.OSPF6DistributelistIN(in["OSPF6"], in["Distributelist"], in["IN"])
}

func OSPF6DistributelistOUT(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distributelist"]; !ok {
		return nil
	}
	if _, ok := in["OUT"]; !ok {
		return nil
	}
	return sw.OSPF6DistributelistOUT(in["OSPF6"], in["Distributelist"], in["OUT"])
}

func NoOSPF6DistributelistIN(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distributelist"]; !ok {
		return nil
	}
	if _, ok := in["IN"]; !ok {
		return nil
	}
	return sw.NoOSPF6DistributelistIN(in["OSPF6"], in["Distributelist"], in["IN"])
}

func NoOSPF6DistributelistOUT(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Distributelist"]; !ok {
		return nil
	}
	if _, ok := in["OUT"]; !ok {
		return nil
	}
	return sw.NoOSPF6DistributelistOUT(in["OSPF6"], in["Distributelist"], in["OUT"])
}

func OSPF6AreaDefaultCost(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["DefaultCost"]; !ok {
		return nil
	}
	return sw.OSPF6AreaDefaultCost(in["OSPF6"], in["Area"], in["DefaultCost"])
}

func NoOSPF6AreaDefaultCost(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["DefaultCost"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaDefaultCost(in["OSPF6"], in["Area"], in["DefaultCost"])
}

func OSPF6AreaNSSA(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	return sw.OSPF6AreaNSSA(in["OSPF6"], in["Area"], in["NSSA"])
}

func OSPF6AreaNSSADefaultOriginate(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	return sw.OSPF6AreaNSSADefaultOriginate(in["OSPF6"], in["Area"], in["NSSA"], in["DefaultOriginate"])
}

func OSPF6AreaNSSANoRedistribution(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	if _, ok := in["NoRedistribution"]; !ok {
		return nil
	}
	return sw.OSPF6AreaNSSANoRedistribution(in["OSPF6"], in["Area"], in["NSSA"], in["NoRedistribution"])
}

func OSPF6AreaNSSANoSummary(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	if _, ok := in["NoSummary"]; !ok {
		return nil
	}
	return sw.OSPF6AreaNSSANoSummary(in["OSPF6"], in["Area"], in["NSSA"], in["NoSummary"])
}

func OSPF6AreaNSSAStabilityInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	if _, ok := in["StabilityInterval"]; !ok {
		return nil
	}
	return sw.OSPF6AreaNSSAStabilityInterval(in["OSPF6"], in["Area"], in["NSSA"], in["StabilityInterval"])
}

func OSPF6AreaTranslatorrole(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Translatorrole"]; !ok {
		return nil
	}
	return sw.OSPF6AreaTranslatorrole(in["OSPF6"], in["Area"], in["Translatorrole"])
}

func NoOSPF6AreaNSSA(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaNSSA(in["OSPF6"], in["Area"], in["NSSA"])
}

func NoOSPF6AreaNSSADefaultOriginate(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	if _, ok := in["DefaultOriginate"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaNSSADefaultOriginate(in["OSPF6"], in["Area"], in["NSSA"], in["DefaultOriginate"])
}

func NoOSPF6AreaNSSANoRedistribution(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	if _, ok := in["NoRedistribution"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaNSSANoRedistribution(in["OSPF6"], in["Area"], in["NSSA"], in["NoRedistribution"])
}

func NoOSPF6AreaNSSANoSummary(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	if _, ok := in["NoSummary"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaNSSANoSummary(in["OSPF6"], in["Area"], in["NSSA"], in["NoSummary"])
}

func NoOSPF6AreaNSSAStabilityInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["NSSA"]; !ok {
		return nil
	}
	if _, ok := in["StabilityInterval"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaNSSAStabilityInterval(in["OSPF6"], in["Area"], in["NSSA"], in["StabilityInterval"])
}

func NoOSPF6AreaTranslatorrole(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Translatorrole"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaTranslatorrole(in["OSPF6"], in["Area"], in["Translatorrole"])
}

func OSPF6AreaStub(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Stub"]; !ok {
		return nil
	}
	return sw.OSPF6AreaStub(in["OSPF6"], in["Area"], in["Stub"])
}

func OSPF6AreaStubNoSummary(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Stub"]; !ok {
		return nil
	}
	if _, ok := in["NoSummary"]; !ok {
		return nil
	}
	return sw.OSPF6AreaStubNoSummary(in["OSPF6"], in["Area"], in["Stub"], in["NoSummary"])
}

func NoOSPF6AreaStub(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Stub"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaStub(in["OSPF6"], in["Area"], in["Stub"])
}

func NoOSPF6AreaStubNoSummary(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Stub"]; !ok {
		return nil
	}
	if _, ok := in["NoSummary"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaStubNoSummary(in["OSPF6"], in["Area"], in["Stub"], in["NoSummary"])
}

func OSPF6AreaRange(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Range"]; !ok {
		return nil
	}
	return sw.OSPF6AreaRange(in["OSPF6"], in["Area"], in["Range"])
}

func OSPF6AreaRangeAdvertise(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Range"]; !ok {
		return nil
	}
	if _, ok := in["Advertise"]; !ok {
		return nil
	}
	return sw.OSPF6AreaRangeAdvertise(in["OSPF6"], in["Area"], in["Range"], in["Advertise"])
}

func OSPF6AreaRangeNoAdvertise(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Range"]; !ok {
		return nil
	}
	if _, ok := in["NoAdvertise"]; !ok {
		return nil
	}
	return sw.OSPF6AreaRangeNoAdvertise(in["OSPF6"], in["Area"], in["Range"], in["NoAdvertise"])
}

func NoOSPF6AreaRange(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Range"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaRange(in["OSPF6"], in["Area"], in["Range"])
}

func NoOSPF6AreaRangeAdvertise(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Range"]; !ok {
		return nil
	}
	if _, ok := in["Advertise"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaRangeAdvertise(in["OSPF6"], in["Area"], in["Range"], in["Advertise"])
}

func NoOSPF6AreaRangeNoAdvertise(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Range"]; !ok {
		return nil
	}
	if _, ok := in["NoAdvertise"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaRangeNoAdvertise(in["OSPF6"], in["Area"], in["Range"], in["NoAdvertise"])
}

func OSPF6AreaVirtuallink(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 3 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	return sw.OSPF6AreaVirtuallink(in["OSPF6"], in["Area"], in["Virtuallink"])
}

func OSPF6AreaVirtuallinkDeadInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["DeadInterval"]; !ok {
		return nil
	}
	return sw.OSPF6AreaVirtuallinkDeadInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["DeadInterval"])
}

func OSPF6AreaVirtuallinkHelloInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["HelloInterval"]; !ok {
		return nil
	}
	return sw.OSPF6AreaVirtuallinkHelloInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["HelloInterval"])
}

func OSPF6AreaVirtuallinkInstanceid(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["Instanceid"]; !ok {
		return nil
	}
	return sw.OSPF6AreaVirtuallinkInstanceid(in["OSPF6"], in["Area"], in["Virtuallink"], in["Instanceid"])
}

func OSPF6AreaVirtuallinkRetransmitInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["RetransmitInterval"]; !ok {
		return nil
	}
	return sw.OSPF6AreaVirtuallinkRetransmitInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["RetransmitInterval"])
}

func OSPF6AreaVirtuallinkTransmitDelay(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["TransmitDelay"]; !ok {
		return nil
	}
	return sw.OSPF6AreaVirtuallinkTransmitDelay(in["OSPF6"], in["Area"], in["Virtuallink"], in["TransmitDelay"])
}

func NoOSPF6AreaVirtuallink(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 4 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaVirtuallink(in["OSPF6"], in["Area"], in["Virtuallink"])
}

func NoOSPF6AreaVirtuallinkDeadInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["DeadInterval"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaVirtuallinkDeadInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["DeadInterval"])
}

func NoOSPF6AreaVirtuallinkHelloInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["HelloInterval"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaVirtuallinkHelloInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["HelloInterval"])
}

func NoOSPF6AreaVirtuallinkInstanceid(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["Instanceid"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaVirtuallinkInstanceid(in["OSPF6"], in["Area"], in["Virtuallink"], in["Instanceid"])
}

func NoOSPF6AreaVirtuallinkRetransmitInterval(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["RetransmitInterval"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaVirtuallinkRetransmitInterval(in["OSPF6"], in["Area"], in["Virtuallink"], in["RetransmitInterval"])
}

func NoOSPF6AreaVirtuallinkTransmitDelay(sw Switch, in map[string]string) []*command.Command {
	if len(in) != 5 {
		return nil
	}
	if _, ok := in["No"]; !ok {
		return nil
	}
	if _, ok := in["OSPF6"]; !ok {
		return nil
	}
	if _, ok := in["Area"]; !ok {
		return nil
	}
	if _, ok := in["Virtuallink"]; !ok {
		return nil
	}
	if _, ok := in["TransmitDelay"]; !ok {
		return nil
	}
	return sw.NoOSPF6AreaVirtuallinkTransmitDelay(in["OSPF6"], in["Area"], in["Virtuallink"], in["TransmitDelay"])
}
