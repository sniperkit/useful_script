package dsl

import (
	"command"
)

type V5624G struct {
	Name string
}

func (v5 V5624G) Port(typ, slot, port string) string {
	return port
}

var V5 = V5624G{
	Name: "V5",
}

func (v5 V5624G) PortSlotTypeEnable(Port, Slot, Type, Enable string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})
	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "port enable " + v5.Port(Type, Slot, Port),
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) PortSlotTypeDisable(Port, Slot, Type, Disable string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})
	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "port disable " + v5.Port(Type, Slot, Port),
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) PortSlotTypeSpeed(Port, Slot, Type, Speed string) []*command.Command {

	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})
	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "port speed  " + Port + " " + Speed,
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) PortSlotTypePvid(Port, Slot, Type, Pvid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})
	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "vlan pvid " + Port + " " + Pvid,
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLAN(VLAN string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})

	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "vlan create " + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoVLAN(VLAN string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})

	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "no vlan " + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "vlan add " + VLAN + " " + Port + " untagged",
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "vlan add " + VLAN + " " + Port + " tagged",
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})

	res = append(res, v5.PortSlotTypePvid(Port, Slot, AddT, VLAN)...)
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANDelTypeSlotPort(VLAN, Del, Type, Slot, Port string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "vlan del " + VLAN + " " + Port,
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANDelTTypeSlotPort(VLAN, DelT, Type, Slot, Port string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "bridge",
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "vlan del " + VLAN + " " + Port,
	})

	res = append(res, &command.Command{
		Mode: "bridge",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANShutdown(VLAN, Shutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "shutdown",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANNoShutdown(VLAN, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no shutdown",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANIP(VLAN, IP string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ip address " + IP,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoVLANIP(VLAN, IP string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ip address " + IP,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANIP2(VLAN, IP2 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ip address " + IP2 + " secondary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoVLANIP2(VLAN, IP2 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ip address " + IP2 + " secondary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoInterfaceTypeIfname(Interface, Type, Ifname string) []*command.Command {
	return nil
}

func (v5 V5624G) VLANAddTypeSlotPortIP(VLAN, Add, Type, Slot, Port, IP string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v5.VLAN(VLAN)...)
	res = append(res, v5.VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port)...)
	res = append(res, v5.VLANIP(VLAN, IP)...)
	return res
}

func (v5 V5624G) VLANAddTypeSlotPortIP2(VLAN, Add, Type, Slot, Port, IP2 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v5.VLAN(VLAN)...)
	res = append(res, v5.VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port)...)
	res = append(res, v5.VLANIP2(VLAN, IP2)...)
	return res
}

func (v5 V5624G) VLANAddTypeSlotPortIPNoShutdown(VLAN, Add, Type, Slot, Port, IP, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v5.VLAN(VLAN)...)
	res = append(res, v5.VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port)...)
	res = append(res, v5.VLANIP(VLAN, IP)...)
	res = append(res, v5.VLANNoShutdown(VLAN, NoShutdown)...)
	return res
}

func (v5 V5624G) VLANAddTypeSlotPortIP2NoShutdown(VLAN, Add, Type, Slot, Port, IP2, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v5.VLAN(VLAN)...)
	res = append(res, v5.VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port)...)
	res = append(res, v5.VLANIP2(VLAN, IP2)...)
	res = append(res, v5.VLANNoShutdown(VLAN, NoShutdown)...)
	return res
}

func (v5 V5624G) VLANAddTTypeSlotPortIP(VLAN, AddT, Type, Slot, Port, IP string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v5.VLAN(VLAN)...)
	res = append(res, v5.VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port)...)
	res = append(res, v5.VLANIP(VLAN, IP)...)
	return res
}

func (v5 V5624G) VLANAddTTypeSlotPortIP2(VLAN, AddT, Type, Slot, Port, IP2 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v5.VLAN(VLAN)...)
	res = append(res, v5.VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port)...)
	res = append(res, v5.VLANIP2(VLAN, IP2)...)
	return res
}

func (v5 V5624G) VLANAddTTypeSlotPortIPNoShutdown(VLAN, AddT, Type, Slot, Port, IP, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v5.VLAN(VLAN)...)
	res = append(res, v5.VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port)...)
	res = append(res, v5.VLANIP(VLAN, IP)...)
	res = append(res, v5.VLANNoShutdown(VLAN, NoShutdown)...)
	return res
}

func (v5 V5624G) VLANAddTTypeSlotPortIP2NoShutdown(VLAN, AddT, Type, Slot, Port, IP2, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v5.VLAN(VLAN)...)
	res = append(res, v5.VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port)...)
	res = append(res, v5.VLANIP2(VLAN, IP2)...)
	res = append(res, v5.VLANNoShutdown(VLAN, NoShutdown)...)
	return res
}

func (v5 V5624G) VLANIP6Enable(VLAN, IP6, Enable string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 enable",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoVLANIP6Enable(VLAN, IP6, Enable string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 enable",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANIP6(VLAN, IP6 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 address " + IP6,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoVLANIP6(VLAN, IP6 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 address " + IP6,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANIP6LL(VLAN, IP6LL string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 address link-local " + IP6LL,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) VLANIP6LLIP6(VLAN, IP6LL, IP6 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, v5.VLANIP6LL(VLAN, IP6LL)...)
	res = append(res, v5.VLANIP6(VLAN, IP6)...)
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoVLANIP6LL(VLAN, IP6LL string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 address link-local " + IP6LL,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6(OSPF6 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6(OSPF6 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "no router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6Rid(OSPF6, Rid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "router-id  " + Rid,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6Rid(OSPF6, Rid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no router-id",
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6InterfaceArea(OSPF6, Interface, Area string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 router ospf area " + Area,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6InterfaceArea(OSPF6, Interface, Area string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 router ospf area " + Area,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6InterfaceCost(OSPF6, Interface, Cost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 ospf cost " + Cost,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6InterfaceDeadInterval(OSPF6, Interface, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 ospf dead-interval " + DeadInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6InterfaceHelloInterval(OSPF6, Interface, HelloInterval string) []*command.Command {

	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 ospf hello-interval " + HelloInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6InterfaceRetransmitInterval(OSPF6, Interface, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 ospf retransmit-interval " + RetransmitInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6InterfaceTransmitDelay(OSPF6, Interface, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 ospf transmit-delay " + TransmitDelay,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6InterfacePriority(OSPF6, Interface, Priority string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 ospf priority " + Priority,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6InterfaceNetworktype(OSPF6, Interface, Networktype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ipv6 ospf network " + Networktype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6InterfaceCost(OSPF6, Interface, Cost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 ospf cost",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6InterfaceDeadInterval(OSPF6, Interface, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 ospf dead-interval ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6InterfaceHelloInterval(OSPF6, Interface, HelloInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 ospf hello-interval",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPF6InterfaceRetransmitInterval(OSPF6, Interface, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 ospf retransmit-interval",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6InterfaceTransmitDelay(OSPF6, Interface, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 ospf transmit-delay",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6InterfacePriority(OSPF6, Interface, Priority string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 ospf priority",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6InterfaceNetworktype(OSPF6, Interface, Networktype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface br" + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ipv6 ospf network " + Networktype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6ReferenceBandwidth(OSPF6, ReferenceBandwidth string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "auto-cost reference-bandwidth " + ReferenceBandwidth,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6ReferenceBandwidth(OSPF6, ReferenceBandwidth string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no auto-cost reference-bandwidth",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginate(OSPF6, DefaultOriginate string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateRoutemap(OSPF6, DefaultOriginate, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateMetric(OSPF6, DefaultOriginate, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateMetrictype(OSPF6, DefaultOriginate, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateMetricMetrictype(OSPF6, DefaultOriginate, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateMetricRoutemap(OSPF6, DefaultOriginate, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate route-map " + Routemap + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateMetrictypeRoutemap(OSPF6, DefaultOriginate, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate route-map " + Routemap + " metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate route-map " + Routemap + " metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateAlways(OSPF6, DefaultOriginate, Always string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateAlwaysRoutemap(OSPF6, DefaultOriginate, Always, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateAlwaysMetric(OSPF6, DefaultOriginate, Always, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateAlwaysMetrictype(OSPF6, DefaultOriginate, Always, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateAlwaysMetricMetrictype(OSPF6, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateAlwaysMetricRoutemap(OSPF6, DefaultOriginate, Always, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always route-map " + Routemap + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateAlwaysMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always route-map " + Routemap + " metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always route-map " + Routemap + " metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginate(OSPF6, DefaultOriginate string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginateRoutemap(OSPF6, DefaultOriginate, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginateMetric(OSPF6, DefaultOriginate, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginateMetrictype(OSPF6, DefaultOriginate, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginateMetricMetrictype(OSPF6, DefaultOriginate, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginateMetricRoutemap(OSPF6, DefaultOriginate, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginateMetrictypeRoutemap(OSPF6, DefaultOriginate, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric-type " + Metrictype + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginateMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric-type " + Metrictype + " metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginateAlways(OSPF6, DefaultOriginate, Always string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPF6DefaultOriginateAlwaysRoutemap(OSPF6, DefaultOriginate, Always, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPF6DefaultOriginateAlwaysMetric(OSPF6, DefaultOriginate, Always, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always  metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPF6DefaultOriginateAlwaysMetrictype(OSPF6, DefaultOriginate, Always, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPF6DefaultOriginateAlwaysMetricMetrictype(OSPF6, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPF6DefaultOriginateAlwaysMetricRoutemap(OSPF6, DefaultOriginate, Always, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPF6DefaultOriginateAlwaysMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric-type " + Metrictype + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric-type " + Metrictype + " metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6Redistribute(OSPF6, Redistribute string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6RedistributeMetric(OSPF6, Redistribute, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6RedistributeMetrictype(OSPF6, Redistribute, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6RedistributeRoutemap(OSPF6, Redistribute, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no edistribute " + Redistribute + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6RedistributeMetricMetrictype(OSPF6, Redistribute, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6RedistributeMetricRoutemap(OSPF6, Redistribute, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6RedistributeMetricMetrictypeRoutemap(OSPF6, Redistribute, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric-type " + Metrictype + " metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6Redistribute(OSPF6, Redistribute string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no redistribute " + Redistribute,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6RedistributeMetric(OSPF6, Redistribute, Metric string) []*command.Command {
	return v5.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v5 V5624G) NoOSPF6RedistributeMetrictype(OSPF6, Redistribute, Metrictype string) []*command.Command {
	return v5.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v5 V5624G) NoOSPF6RedistributeRoutemap(OSPF6, Redistribute, Routemap string) []*command.Command {
	return v5.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v5 V5624G) NoOSPF6RedistributeMetricMetrictype(OSPF6, Redistribute, Metric, Metrictype string) []*command.Command {
	return v5.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v5 V5624G) NoOSPF6RedistributeMetricRoutemap(OSPF6, Redistribute, Metric, Routemap string) []*command.Command {
	return v5.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v5 V5624G) NoOSPF6RedistributeMetricMetrictypeRoutemap(OSPF6, Redistribute, Metric, Metrictype, Routemap string) []*command.Command {
	return v5.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v5 V5624G) OSPF6Summary(OSPF6, Summary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "summary-address " + Summary,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6SummaryNoAdvertise(OSPF6, Summary, NoAdvertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "summary-address " + Summary + "no-advertise",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6Summary(OSPF6, Summary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no summary-address " + Summary,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6SummaryNoAdvertise(OSPF6, Summary, NoAdvertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no summary-address " + Summary + " no-advertise ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DefaultMetric(OSPF6, DefaultMetric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-metric " + DefaultMetric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DefaultMetric(OSPF6, DefaultMetric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-metric",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6Passive(OSPF6, Passive string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "passive-interface " + Passive,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6Passive(OSPF6, Passive string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no passive-interfaced " + Passive,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AdminDistance(OSPF6, AdminDistance string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance " + AdminDistance,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AdminDistance(OSPF6, AdminDistance string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance " + AdminDistance,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DistanceExternal(OSPF6, Distance, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospfv3 external " + External,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DistanceInter(OSPF6, Distance, Inter string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospfv3 inter-area " + Inter,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DistanceIntra(OSPF6, Distance, Intra string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospfv3 intra-area " + Intra,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DistanceInterIntra(OSPF6, Distance, Inter, Intra string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospfv3 inter-area " + Inter + " intra-area " + Intra,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DistanceInterExternal(OSPF6, Distance, Inter, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospfv3 inter-area " + Inter + " external " + External,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DistanceInterIntraExternal(OSPF6, Distance, Inter, Intra, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospfv3 inter-area " + Inter + " intra-area " + Intra + " external " + External,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DistanceExternal(OSPF6, Distance, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospfv3",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DistanceInter(OSPF6, Distance, Inter string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospfv3",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DistanceIntra(OSPF6, Distance, Intra string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospfv3",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DistanceInterIntra(OSPF6, Distance, Inter, Intra string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospfv3",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DistanceInterExternal(OSPF6, Distance, Inter, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospfv3",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DistanceInterIntraExternal(OSPF6, Distance, Inter, Intra, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospfv3",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DistributelistIN(OSPF6, Distributelist, IN string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distribute-list " + Distributelist + " in ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6DistributelistOUT(OSPF6, Distributelist, OUT string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distribute-list " + Distributelist + " out ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DistributelistIN(OSPF6, Distributelist, IN string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distribute-list " + Distributelist + " in",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6DistributelistOUT(OSPF6, Distributelist, OUT string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distribute-list " + Distributelist + " out",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaDefaultCost(OSPF6, Area, DefaultCost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " default-cost " + DefaultCost,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaDefaultCost(OSPF6, Area, DefaultCost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " default-cost ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaNSSA(OSPF6, Area, NSSA string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaNSSADefaultOriginate(OSPF6, Area, NSSA, DefaultOriginate string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa default-information-originate",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaNSSANoRedistribution(OSPF6, Area, NSSA, Redistribution string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa no-redistribution",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaNSSANoSummary(OSPF6, Area, NSSA, Summary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa no-summary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaNSSAStabilityInterval(OSPF6, Area, NSSA, StabilityInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa stability-interval " + StabilityInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaTranslatorrole(OSPF6, Area, Translatorrole string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " translator-role " + Translatorrole,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaNSSA(OSPF6, Area, NSSA string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaNSSADefaultOriginate(OSPF6, Area, NSSA, DefaultOriginate string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa  default-information-originate",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaNSSANoRedistribution(OSPF6, Area, NSSA, Redistribution string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa no-redistribution",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaNSSANoSummary(OSPF6, Area, NSSA, Summary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa no-summary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaNSSAStabilityInterval(OSPF6, Area, NSSA, StabilityInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa stability-interval",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaTranslatorrole(OSPF6, Area, Translatorrole string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa translator-role ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaStub(OSPF6, Area, Stub string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " stub ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaStubNoSummary(OSPF6, Area, Stub, NoSummary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " stub no-summary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaStub(OSPF6, Area, Stub string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " stub",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaStubNoSummary(OSPF6, Area, Stub, NoSummary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " stub no-summary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaRange(OSPF6, Area, Range string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " range " + Range,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaRangeAdvertise(OSPF6, Area, Range, Advertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " range " + Range + " advertise ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaRangeNoAdvertise(OSPF6, Area, Range, NoAdvertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " range " + Range + " no-advertise ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaRange(OSPF6, Area, Range string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " range " + Range,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaRangeAdvertise(OSPF6, Area, Range, Advertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " range " + Range,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaRangeNoAdvertise(OSPF6, Area, Range, NoAdvertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " range " + Range,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaVirtuallink(OSPF6, Area, Virtuallink string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaVirtuallinkDeadInterval(OSPF6, Area, Virtuallink, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " dead-interval " + DeadInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaVirtuallinkHelloInterval(OSPF6, Area, Virtuallink, HelloInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " hello-interval " + HelloInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaVirtuallinkInstanceid(OSPF6, Area, Virtuallink, Instanceid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " instance-id " + Instanceid,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaVirtuallinkRetransmitInterval(OSPF6, Area, Virtuallink, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " retransmit-interval " + RetransmitInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPF6AreaVirtuallinkTransmitDelay(OSPF6, Area, Virtuallink, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + "   transmit-delay " + TransmitDelay,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaVirtuallink(OSPF6, Area, Virtuallink string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaVirtuallinkDeadInterval(OSPF6, Area, Virtuallink, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink + " dead-interval " + DeadInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaVirtuallinkHelloInterval(OSPF6, Area, Virtuallink, HelloInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink + " hello-interval " + HelloInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaVirtuallinkInstanceid(OSPF6, Area, Virtuallink, Instanceid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink + " instance-id " + Instanceid,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaVirtuallinkRetransmitInterval(OSPF6, Area, Virtuallink, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " retransmit-interval " + RetransmitInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF6AreaVirtuallinkTransmitDelay(OSPF6, Area, Virtuallink, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink + "   transmit-delay " + TransmitDelay,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

//OSPF
func (v5 V5624G) OSPF(OSPF string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPF(OSPF string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "no router ospf " + OSPF,
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFRid(OSPF, Rid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "router-id  " + Rid,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFRid(OSPF, Rid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no router-id",
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFNetworkArea(OSPF, Network, Area string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "network " + Network + " area " + Area,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFNetworkArea(OSPF, Network, Area string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "network " + Network + " area " + Area,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFInterfaceCost(OSPF, Interface, Cost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ip ospf cost " + Cost,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFInterfaceDeadInterval(OSPF, Interface, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ip ospf dead-interval " + DeadInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFInterfaceHelloInterval(OSPF, Interface, HelloInterval string) []*command.Command {

	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ip ospf hello-interval " + HelloInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFInterfaceRetransmitInterval(OSPF, Interface, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ip ospf retransmit-interval " + RetransmitInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFInterfaceTransmitDelay(OSPF, Interface, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ip ospf transmit-delay " + TransmitDelay,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFInterfacePriority(OSPF, Interface, Priority string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ip ospf priority " + Priority,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFInterfaceNetworktype(OSPF, Interface, Networktype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "ip ospf network " + Networktype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFInterfaceCost(OSPF, Interface, Cost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ip ospf cost",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFInterfaceDeadInterval(OSPF, Interface, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ip ospf dead-interval ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFInterfaceHelloInterval(OSPF, Interface, HelloInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ip ospf hello-interval",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPFInterfaceRetransmitInterval(OSPF, Interface, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ip ospf retransmit-interval",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFInterfaceTransmitDelay(OSPF, Interface, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ip ospf transmit-delay",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFInterfacePriority(OSPF, Interface, Priority string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ip ospf priority",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFInterfaceNetworktype(OSPF, Interface, Networktype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no ip ospf network " + Networktype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFReferenceBandwidth(OSPF, ReferenceBandwidth string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "auto-cost reference-bandwidth " + ReferenceBandwidth,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFReferenceBandwidth(OSPF, ReferenceBandwidth string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no auto-cost reference-bandwidth",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginate(OSPF, DefaultOriginate string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateRoutemap(OSPF, DefaultOriginate, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateMetric(OSPF, DefaultOriginate, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateMetrictype(OSPF, DefaultOriginate, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateMetricMetrictype(OSPF, DefaultOriginate, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateMetricRoutemap(OSPF, DefaultOriginate, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate route-map " + Routemap + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateMetrictypeRoutemap(OSPF, DefaultOriginate, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate route-map " + Routemap + " metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate route-map " + Routemap + " metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateAlways(OSPF, DefaultOriginate, Always string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateAlwaysRoutemap(OSPF, DefaultOriginate, Always, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateAlwaysMetric(OSPF, DefaultOriginate, Always, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateAlwaysMetrictype(OSPF, DefaultOriginate, Always, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateAlwaysMetricMetrictype(OSPF, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateAlwaysMetricRoutemap(OSPF, DefaultOriginate, Always, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always route-map " + Routemap + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateAlwaysMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always route-map " + Routemap + " metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-information originate always route-map " + Routemap + " metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginate(OSPF, DefaultOriginate string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginateRoutemap(OSPF, DefaultOriginate, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginateMetric(OSPF, DefaultOriginate, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginateMetrictype(OSPF, DefaultOriginate, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginateMetricMetrictype(OSPF, DefaultOriginate, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginateMetricRoutemap(OSPF, DefaultOriginate, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginateMetrictypeRoutemap(OSPF, DefaultOriginate, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric-type " + Metrictype + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginateMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate metric-type " + Metrictype + " metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginateAlways(OSPF, DefaultOriginate, Always string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPFDefaultOriginateAlwaysRoutemap(OSPF, DefaultOriginate, Always, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPFDefaultOriginateAlwaysMetric(OSPF, DefaultOriginate, Always, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always  metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPFDefaultOriginateAlwaysMetrictype(OSPF, DefaultOriginate, Always, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPFDefaultOriginateAlwaysMetricMetrictype(OSPF, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPFDefaultOriginateAlwaysMetricRoutemap(OSPF, DefaultOriginate, Always, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v5 V5624G) NoOSPFDefaultOriginateAlwaysMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric-type " + Metrictype + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-information originate always metric-type " + Metrictype + " metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFRedistribute(OSPF, Redistribute string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFRedistributeMetric(OSPF, Redistribute, Metric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFRedistributeMetrictype(OSPF, Redistribute, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric-type " + Metrictype,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFRedistributeRoutemap(OSPF, Redistribute, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no edistribute " + Redistribute + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFRedistributeMetricMetrictype(OSPF, Redistribute, Metric, Metrictype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric-type " + Metrictype + " metric " + Metric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFRedistributeMetricRoutemap(OSPF, Redistribute, Metric, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFRedistributeMetricMetrictypeRoutemap(OSPF, Redistribute, Metric, Metrictype, Routemap string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "redistribute " + Redistribute + " metric-type " + Metrictype + " metric " + Metric + " route-map " + Routemap,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFRedistribute(OSPF, Redistribute string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no redistribute " + Redistribute,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFRedistributeMetric(OSPF, Redistribute, Metric string) []*command.Command {
	return v5.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v5 V5624G) NoOSPFRedistributeMetrictype(OSPF, Redistribute, Metrictype string) []*command.Command {
	return v5.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v5 V5624G) NoOSPFRedistributeRoutemap(OSPF, Redistribute, Routemap string) []*command.Command {
	return v5.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v5 V5624G) NoOSPFRedistributeMetricMetrictype(OSPF, Redistribute, Metric, Metrictype string) []*command.Command {
	return v5.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v5 V5624G) NoOSPFRedistributeMetricRoutemap(OSPF, Redistribute, Metric, Routemap string) []*command.Command {
	return v5.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v5 V5624G) NoOSPFRedistributeMetricMetrictypeRoutemap(OSPF, Redistribute, Metric, Metrictype, Routemap string) []*command.Command {
	return v5.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v5 V5624G) OSPFSummary(OSPF, Summary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "summary-address " + Summary,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFSummaryNoAdvertise(OSPF, Summary, NoAdvertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "summary-address " + Summary + "no-advertise",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFSummary(OSPF, Summary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no summary-address " + Summary,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFSummaryNoAdvertise(OSPF, Summary, NoAdvertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no summary-address " + Summary + " no-advertise ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDefaultMetric(OSPF, DefaultMetric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "default-metric " + DefaultMetric,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDefaultMetric(OSPF, DefaultMetric string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no default-metric",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFPassive(OSPF, Passive string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "passive-interface " + Passive,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFPassive(OSPF, Passive string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no passive-interfaced " + Passive,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAdminDistance(OSPF, AdminDistance string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance " + AdminDistance,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAdminDistance(OSPF, AdminDistance string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance " + AdminDistance,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDistanceExternal(OSPF, Distance, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospf external " + External,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDistanceInter(OSPF, Distance, Inter string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospf inter-area " + Inter,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDistanceIntra(OSPF, Distance, Intra string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospf intra-area " + Intra,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDistanceInterIntra(OSPF, Distance, Inter, Intra string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospf inter-area " + Inter + " intra-area " + Intra,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDistanceInterExternal(OSPF, Distance, Inter, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospf inter-area " + Inter + " external " + External,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDistanceInterIntraExternal(OSPF, Distance, Inter, Intra, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distance ospf inter-area " + Inter + " intra-area " + Intra + " external " + External,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDistanceExternal(OSPF, Distance, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospf",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDistanceInter(OSPF, Distance, Inter string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospf",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDistanceIntra(OSPF, Distance, Intra string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospf",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDistanceInterIntra(OSPF, Distance, Inter, Intra string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospf",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDistanceInterExternal(OSPF, Distance, Inter, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospf",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDistanceInterIntraExternal(OSPF, Distance, Inter, Intra, External string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distance ospf",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDistributelistIN(OSPF, Distributelist, IN string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distribute-list " + Distributelist + " in ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFDistributelistOUT(OSPF, Distributelist, OUT string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "distribute-list " + Distributelist + " out ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDistributelistIN(OSPF, Distributelist, IN string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distribute-list " + Distributelist + " in",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFDistributelistOUT(OSPF, Distributelist, OUT string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no distribute-list " + Distributelist + " out",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaDefaultCost(OSPF, Area, DefaultCost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " default-cost " + DefaultCost,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaDefaultCost(OSPF, Area, DefaultCost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " default-cost ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaNSSA(OSPF, Area, NSSA string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaNSSADefaultOriginate(OSPF, Area, NSSA, DefaultOriginate string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa default-information-originate",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaNSSANoRedistribution(OSPF, Area, NSSA, Redistribution string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa no-redistribution",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaNSSANoSummary(OSPF, Area, NSSA, Summary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa no-summary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaNSSAStabilityInterval(OSPF, Area, NSSA, StabilityInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " nssa stability-interval " + StabilityInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaTranslatorrole(OSPF, Area, Translatorrole string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " translator-role " + Translatorrole,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaNSSA(OSPF, Area, NSSA string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaNSSADefaultOriginate(OSPF, Area, NSSA, DefaultOriginate string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa  default-information-originate",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaNSSANoRedistribution(OSPF, Area, NSSA, Redistribution string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa no-redistribution",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaNSSANoSummary(OSPF, Area, NSSA, Summary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa no-summary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaNSSAStabilityInterval(OSPF, Area, NSSA, StabilityInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa stability-interval",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaTranslatorrole(OSPF, Area, Translatorrole string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " nssa translator-role ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaStub(OSPF, Area, Stub string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " stub ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaStubNoSummary(OSPF, Area, Stub, NoSummary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " stub no-summary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaStub(OSPF, Area, Stub string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " stub",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaStubNoSummary(OSPF, Area, Stub, NoSummary string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " stub no-summary",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaRange(OSPF, Area, Range string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " range " + Range,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaRangeAdvertise(OSPF, Area, Range, Advertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " range " + Range + " advertise ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaRangeNoAdvertise(OSPF, Area, Range, NoAdvertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " range " + Range + " no-advertise ",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaRange(OSPF, Area, Range string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " range " + Range,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaRangeAdvertise(OSPF, Area, Range, Advertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " range " + Range,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaRangeNoAdvertise(OSPF, Area, Range, NoAdvertise string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " range " + Range,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaVirtuallink(OSPF, Area, Virtuallink string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaVirtuallinkDeadInterval(OSPF, Area, Virtuallink, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " dead-interval " + DeadInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaVirtuallinkHelloInterval(OSPF, Area, Virtuallink, HelloInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " hello-interval " + HelloInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaVirtuallinkInstanceid(OSPF, Area, Virtuallink, Instanceid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " instance-id " + Instanceid,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaVirtuallinkRetransmitInterval(OSPF, Area, Virtuallink, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " retransmit-interval " + RetransmitInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) OSPFAreaVirtuallinkTransmitDelay(OSPF, Area, Virtuallink, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + "   transmit-delay " + TransmitDelay,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaVirtuallink(OSPF, Area, Virtuallink string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaVirtuallinkDeadInterval(OSPF, Area, Virtuallink, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink + " dead-interval " + DeadInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaVirtuallinkHelloInterval(OSPF, Area, Virtuallink, HelloInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink + " hello-interval " + HelloInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaVirtuallinkInstanceid(OSPF, Area, Virtuallink, Instanceid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink + " instance-id " + Instanceid,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaVirtuallinkRetransmitInterval(OSPF, Area, Virtuallink, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "area " + Area + " virtual-link " + Virtuallink + " retransmit-interval " + RetransmitInterval,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v5 V5624G) NoOSPFAreaVirtuallinkTransmitDelay(OSPF, Area, Virtuallink, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "router ospf " + OSPF,
	})

	res = append(res, &command.Command{
		Mode: "config-router",
		CMD:  "no area " + Area + " virtual-link " + Virtuallink + "   transmit-delay " + TransmitDelay,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}
