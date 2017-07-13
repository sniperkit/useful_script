package dsl

import (
	"command"
	"log"
)

type V8500 struct {
	Name string
}

func (v8 V8500) Port(typ, slot, port string) string {
	return " " + typ + " " + slot + "/" + port
}

var V8 = V8500{
	Name: "V8",
}

func (v8 V8500) PortSlotTypeEnable(Port, Slot, Type, Enable string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface " + v8.Port(Type, Slot, Port),
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

func (v8 V8500) PortSlotTypeDisable(Port, Slot, Type, Disable string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})

	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface " + v8.Port(Type, Slot, Port),
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

func (v8 V8500) PortSlotTypeSpeed(Port, Slot, Type, Speed string) []*command.Command {

	var local string
	if Speed == "1000" {
		local = "1g"
	} else if Speed == "100" {
		local = "100m"
	} else if Speed == "10" {
		local = "10m"
	} else {
		log.Printf("Invalid spped set % set port :%s speed to 1g ", Speed, Port)
		local = "1g"
	}

	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface " + v8.Port(Type, Slot, Port),
	})
	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "bandwidth " + local,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res

}

func (v8 V8500) PortSlotTypePvid(Port, Slot, Type, Pvid string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface " + v8.Port(Type, Slot, Port),
	})
	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "pvid " + Pvid,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) VLAN(VLAN string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})

	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "vlan database",
	})

	res = append(res, &command.Command{
		Mode: "config-vlan",
		CMD:  "vlan " + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-vlan",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) NoVLAN(VLAN string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})

	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "vlan database",
	})

	res = append(res, &command.Command{
		Mode: "config-vlan",
		CMD:  "no vlan " + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-vlan",
		CMD:  "exit",
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface " + v8.Port(Type, Slot, Port),
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "switchport mode access",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "switchport access vlan " + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface " + v8.Port(Type, Slot, Port),
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "switchport mode trunk",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "switchport trunk allowed vlan add " + VLAN,
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) VLANDelTypeSlotPort(VLAN, Del, Type, Slot, Port string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface " + v8.Port(Type, Slot, Port),
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no switchport access vlan",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) VLANDelTTypeSlotPort(VLAN, DelT, Type, Slot, Port string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface " + v8.Port(Type, Slot, Port),
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "no switchport access trunk",
	})

	res = append(res, &command.Command{
		Mode: "config-if",
		CMD:  "exit",
	})
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) VLANShutdown(VLAN, Shutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) VLANNoShutdown(VLAN, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) VLANIP(VLAN, IP string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) NoVLANIP(VLAN, IP string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) VLANIP2(VLAN, IP2 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) NoVLANIP2(VLAN, IP2 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) VLANAddTypeSlotPortIP(VLAN, Add, Type, Slot, Port, IP string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v8.VLAN(VLAN)...)
	res = append(res, v8.VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port)...)
	res = append(res, v8.VLANIP(VLAN, IP)...)
	return res
}

func (v8 V8500) VLANAddTypeSlotPortIP2(VLAN, Add, Type, Slot, Port, IP2 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v8.VLAN(VLAN)...)
	res = append(res, v8.VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port)...)
	res = append(res, v8.VLANIP2(VLAN, IP2)...)
	return res
}

func (v8 V8500) VLANAddTypeSlotPortIPNoShutdown(VLAN, Add, Type, Slot, Port, IP, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v8.VLAN(VLAN)...)
	res = append(res, v8.VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port)...)
	res = append(res, v8.VLANIP(VLAN, IP)...)
	res = append(res, v8.VLANNoShutdown(VLAN, NoShutdown)...)
	return res
}

func (v8 V8500) VLANAddTypeSlotPortIP2NoShutdown(VLAN, Add, Type, Slot, Port, IP2, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v8.VLAN(VLAN)...)
	res = append(res, v8.VLANAddTypeSlotPort(VLAN, Add, Type, Slot, Port)...)
	res = append(res, v8.VLANIP2(VLAN, IP2)...)
	res = append(res, v8.VLANNoShutdown(VLAN, NoShutdown)...)
	return res
}

func (v8 V8500) VLANAddTTypeSlotPortIP(VLAN, AddT, Type, Slot, Port, IP string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v8.VLAN(VLAN)...)
	res = append(res, v8.VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port)...)
	res = append(res, v8.VLANIP(VLAN, IP)...)
	return res
}

func (v8 V8500) VLANAddTTypeSlotPortIP2(VLAN, AddT, Type, Slot, Port, IP2 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v8.VLAN(VLAN)...)
	res = append(res, v8.VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port)...)
	res = append(res, v8.VLANIP2(VLAN, IP2)...)
	return res
}

func (v8 V8500) VLANAddTTypeSlotPortIPNoShutdown(VLAN, AddT, Type, Slot, Port, IP, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v8.VLAN(VLAN)...)
	res = append(res, v8.VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port)...)
	res = append(res, v8.VLANIP(VLAN, IP)...)
	res = append(res, v8.VLANNoShutdown(VLAN, NoShutdown)...)
	return res
}

func (v8 V8500) VLANAddTTypeSlotPortIP2NoShutdown(VLAN, AddT, Type, Slot, Port, IP2, NoShutdown string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, v8.VLAN(VLAN)...)
	res = append(res, v8.VLANAddTTypeSlotPort(VLAN, AddT, Type, Slot, Port)...)
	res = append(res, v8.VLANIP2(VLAN, IP2)...)
	res = append(res, v8.VLANNoShutdown(VLAN, NoShutdown)...)
	return res
}

func (v8 V8500) VLANIP6Enable(VLAN, IP6, Enable string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) NoVLANIP6Enable(VLAN, IP6, Enable string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) VLANIP6(VLAN, IP6 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) NoVLANIP6(VLAN, IP6 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) VLANIP6LL(VLAN, IP6LL string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) VLANIP6LLIP6(VLAN, IP6LL, IP6 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, v8.VLANIP6LL(VLAN, IP6LL)...)
	res = append(res, v8.VLANIP6(VLAN, IP6)...)
	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) NoVLANIP6LL(VLAN, IP6LL string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + VLAN,
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

func (v8 V8500) OSPF6(OSPF6 string) []*command.Command {
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

func (v8 V8500) NoOSPF6(OSPF6 string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "no router ipv6 ospf " + OSPF6,
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) OSPF6Rid(OSPF6, Rid string) []*command.Command {
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

func (v8 V8500) NoOSPF6Rid(OSPF6, Rid string) []*command.Command {
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

func (v8 V8500) OSPF6InterfaceArea(OSPF6, Interface, Area string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) NoOSPF6InterfaceArea(OSPF6, Interface, Area string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) OSPF6InterfaceCost(OSPF6, Interface, Cost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) OSPF6InterfaceDeadInterval(OSPF6, Interface, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) OSPF6InterfaceHelloInterval(OSPF6, Interface, HelloInterval string) []*command.Command {

	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) OSPF6InterfaceRetransmitInterval(OSPF6, Interface, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) OSPF6InterfaceTransmitDelay(OSPF6, Interface, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) OSPF6InterfacePriority(OSPF6, Interface, Priority string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) OSPF6InterfaceNetworktype(OSPF6, Interface, Networktype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) NoOSPF6InterfaceCost(OSPF6, Interface, Cost string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) NoOSPF6InterfaceDeadInterval(OSPF6, Interface, DeadInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) NoOSPF6InterfaceHelloInterval(OSPF6, Interface, HelloInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) NoOSPF6InterfaceRetransmitInterval(OSPF6, Interface, RetransmitInterval string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) NoOSPF6InterfaceTransmitDelay(OSPF6, Interface, TransmitDelay string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) NoOSPF6InterfacePriority(OSPF6, Interface, Priority string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) NoOSPF6InterfaceNetworktype(OSPF6, Interface, Networktype string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "interface vlan " + Interface,
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

func (v8 V8500) OSPF6ReferenceBandwidth(OSPF6, ReferenceBandwidth string) []*command.Command {
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

func (v8 V8500) NoOSPF6ReferenceBandwidth(OSPF6, ReferenceBandwidth string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginate(OSPF6, DefaultOriginate string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateRoutemap(OSPF6, DefaultOriginate, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateMetric(OSPF6, DefaultOriginate, Metric string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateMetrictype(OSPF6, DefaultOriginate, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateMetricMetrictype(OSPF6, DefaultOriginate, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateMetricRoutemap(OSPF6, DefaultOriginate, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateMetrictypeRoutemap(OSPF6, DefaultOriginate, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateAlways(OSPF6, DefaultOriginate, Always string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateAlwaysRoutemap(OSPF6, DefaultOriginate, Always, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateAlwaysMetric(OSPF6, DefaultOriginate, Always, Metric string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateAlwaysMetrictype(OSPF6, DefaultOriginate, Always, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateAlwaysMetricMetrictype(OSPF6, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateAlwaysMetricRoutemap(OSPF6, DefaultOriginate, Always, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateAlwaysMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginate(OSPF6, DefaultOriginate string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateRoutemap(OSPF6, DefaultOriginate, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateMetric(OSPF6, DefaultOriginate, Metric string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateMetrictype(OSPF6, DefaultOriginate, Metrictype string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateMetricMetrictype(OSPF6, DefaultOriginate, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateMetricRoutemap(OSPF6, DefaultOriginate, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateMetrictypeRoutemap(OSPF6, DefaultOriginate, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateAlways(OSPF6, DefaultOriginate, Always string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateAlwaysRoutemap(OSPF6, DefaultOriginate, Always, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateAlwaysMetric(OSPF6, DefaultOriginate, Always, Metric string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateAlwaysMetrictype(OSPF6, DefaultOriginate, Always, Metrictype string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateAlwaysMetricMetrictype(OSPF6, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateAlwaysMetricRoutemap(OSPF6, DefaultOriginate, Always, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateAlwaysMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF6, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6Redistribute(OSPF6, Redistribute string) []*command.Command {
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

func (v8 V8500) OSPF6RedistributeMetric(OSPF6, Redistribute, Metric string) []*command.Command {
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

func (v8 V8500) OSPF6RedistributeMetrictype(OSPF6, Redistribute, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPF6RedistributeRoutemap(OSPF6, Redistribute, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6RedistributeMetricMetrictype(OSPF6, Redistribute, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPF6RedistributeMetricRoutemap(OSPF6, Redistribute, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) OSPF6RedistributeMetricMetrictypeRoutemap(OSPF6, Redistribute, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPF6Redistribute(OSPF6, Redistribute string) []*command.Command {
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

func (v8 V8500) NoOSPF6RedistributeMetric(OSPF6, Redistribute, Metric string) []*command.Command {
	return v8.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v8 V8500) NoOSPF6RedistributeMetrictype(OSPF6, Redistribute, Metrictype string) []*command.Command {
	return v8.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v8 V8500) NoOSPF6RedistributeRoutemap(OSPF6, Redistribute, Routemap string) []*command.Command {
	return v8.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v8 V8500) NoOSPF6RedistributeMetricMetrictype(OSPF6, Redistribute, Metric, Metrictype string) []*command.Command {
	return v8.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v8 V8500) NoOSPF6RedistributeMetricRoutemap(OSPF6, Redistribute, Metric, Routemap string) []*command.Command {
	return v8.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v8 V8500) NoOSPF6RedistributeMetricMetrictypeRoutemap(OSPF6, Redistribute, Metric, Metrictype, Routemap string) []*command.Command {
	return v8.NoOSPF6Redistribute(OSPF6, Redistribute)
}

func (v8 V8500) OSPF6Summary(OSPF6, Summary string) []*command.Command {
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

func (v8 V8500) OSPF6SummaryNoAdvertise(OSPF6, Summary, NoAdvertise string) []*command.Command {
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

func (v8 V8500) NoOSPF6Summary(OSPF6, Summary string) []*command.Command {
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

func (v8 V8500) NoOSPF6SummaryNoAdvertise(OSPF6, Summary, NoAdvertise string) []*command.Command {
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

func (v8 V8500) OSPF6DefaultMetric(OSPF6, DefaultMetric string) []*command.Command {
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

func (v8 V8500) NoOSPF6DefaultMetric(OSPF6, DefaultMetric string) []*command.Command {
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

func (v8 V8500) OSPF6Passive(OSPF6, Passive string) []*command.Command {
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

func (v8 V8500) NoOSPF6Passive(OSPF6, Passive string) []*command.Command {
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

func (v8 V8500) OSPF6AdminDistance(OSPF6, AdminDistance string) []*command.Command {
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

func (v8 V8500) NoOSPF6AdminDistance(OSPF6, AdminDistance string) []*command.Command {
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

func (v8 V8500) OSPF6DistanceExternal(OSPF6, Distance, External string) []*command.Command {
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

func (v8 V8500) OSPF6DistanceInter(OSPF6, Distance, Inter string) []*command.Command {
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

func (v8 V8500) OSPF6DistanceIntra(OSPF6, Distance, Intra string) []*command.Command {
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

func (v8 V8500) OSPF6DistanceInterIntra(OSPF6, Distance, Inter, Intra string) []*command.Command {
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

func (v8 V8500) OSPF6DistanceInterExternal(OSPF6, Distance, Inter, External string) []*command.Command {
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

func (v8 V8500) OSPF6DistanceInterIntraExternal(OSPF6, Distance, Inter, Intra, External string) []*command.Command {
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

func (v8 V8500) NoOSPF6DistanceExternal(OSPF6, Distance, External string) []*command.Command {
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

func (v8 V8500) NoOSPF6DistanceInter(OSPF6, Distance, Inter string) []*command.Command {
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

func (v8 V8500) NoOSPF6DistanceIntra(OSPF6, Distance, Intra string) []*command.Command {
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

func (v8 V8500) NoOSPF6DistanceInterIntra(OSPF6, Distance, Inter, Intra string) []*command.Command {
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

func (v8 V8500) NoOSPF6DistanceInterExternal(OSPF6, Distance, Inter, External string) []*command.Command {
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

func (v8 V8500) NoOSPF6DistanceInterIntraExternal(OSPF6, Distance, Inter, Intra, External string) []*command.Command {
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

func (v8 V8500) OSPF6DistributelistIN(OSPF6, Distributelist, IN string) []*command.Command {
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

func (v8 V8500) OSPF6DistributelistOUT(OSPF6, Distributelist, OUT string) []*command.Command {
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

func (v8 V8500) NoOSPF6DistributelistIN(OSPF6, Distributelist, IN string) []*command.Command {
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

func (v8 V8500) NoOSPF6DistributelistOUT(OSPF6, Distributelist, OUT string) []*command.Command {
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

func (v8 V8500) OSPF6AreaDefaultCost(OSPF6, Area, DefaultCost string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaDefaultCost(OSPF6, Area, DefaultCost string) []*command.Command {
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

func (v8 V8500) OSPF6AreaNSSA(OSPF6, Area, NSSA string) []*command.Command {
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

func (v8 V8500) OSPF6AreaNSSADefaultOriginate(OSPF6, Area, NSSA, DefaultOriginate string) []*command.Command {
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

func (v8 V8500) OSPF6AreaNSSANoRedistribution(OSPF6, Area, NSSA, Redistribution string) []*command.Command {
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

func (v8 V8500) OSPF6AreaNSSANoSummary(OSPF6, Area, NSSA, Summary string) []*command.Command {
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

func (v8 V8500) OSPF6AreaNSSAStabilityInterval(OSPF6, Area, NSSA, StabilityInterval string) []*command.Command {
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

func (v8 V8500) OSPF6AreaTranslatorrole(OSPF6, Area, Translatorrole string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaNSSA(OSPF6, Area, NSSA string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaNSSADefaultOriginate(OSPF6, Area, NSSA, DefaultOriginate string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaNSSANoRedistribution(OSPF6, Area, NSSA, Redistribution string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaNSSANoSummary(OSPF6, Area, NSSA, Summary string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaNSSAStabilityInterval(OSPF6, Area, NSSA, StabilityInterval string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaTranslatorrole(OSPF6, Area, Translatorrole string) []*command.Command {
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

func (v8 V8500) OSPF6AreaStub(OSPF6, Area, Stub string) []*command.Command {
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

func (v8 V8500) OSPF6AreaStubNoSummary(OSPF6, Area, Stub, NoSummary string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaStub(OSPF6, Area, Stub string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaStubNoSummary(OSPF6, Area, Stub, NoSummary string) []*command.Command {
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

func (v8 V8500) OSPF6AreaRange(OSPF6, Area, Range string) []*command.Command {
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

func (v8 V8500) OSPF6AreaRangeAdvertise(OSPF6, Area, Range, Advertise string) []*command.Command {
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

func (v8 V8500) OSPF6AreaRangeNoAdvertise(OSPF6, Area, Range, NoAdvertise string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaRange(OSPF6, Area, Range string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaRangeAdvertise(OSPF6, Area, Range, Advertise string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaRangeNoAdvertise(OSPF6, Area, Range, NoAdvertise string) []*command.Command {
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

func (v8 V8500) OSPF6AreaVirtuallink(OSPF6, Area, Virtuallink string) []*command.Command {
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

func (v8 V8500) OSPF6AreaVirtuallinkDeadInterval(OSPF6, Area, Virtuallink, DeadInterval string) []*command.Command {
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

func (v8 V8500) OSPF6AreaVirtuallinkHelloInterval(OSPF6, Area, Virtuallink, HelloInterval string) []*command.Command {
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

func (v8 V8500) OSPF6AreaVirtuallinkInstanceid(OSPF6, Area, Virtuallink, Instanceid string) []*command.Command {
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

func (v8 V8500) OSPF6AreaVirtuallinkRetransmitInterval(OSPF6, Area, Virtuallink, RetransmitInterval string) []*command.Command {
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

func (v8 V8500) OSPF6AreaVirtuallinkTransmitDelay(OSPF6, Area, Virtuallink, TransmitDelay string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaVirtuallink(OSPF6, Area, Virtuallink string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaVirtuallinkDeadInterval(OSPF6, Area, Virtuallink, DeadInterval string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaVirtuallinkHelloInterval(OSPF6, Area, Virtuallink, HelloInterval string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaVirtuallinkInstanceid(OSPF6, Area, Virtuallink, Instanceid string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaVirtuallinkRetransmitInterval(OSPF6, Area, Virtuallink, RetransmitInterval string) []*command.Command {
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

func (v8 V8500) NoOSPF6AreaVirtuallinkTransmitDelay(OSPF6, Area, Virtuallink, TransmitDelay string) []*command.Command {
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
func (v8 V8500) OSPF(OSPF string) []*command.Command {
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

func (v8 V8500) NoOSPF(OSPF string) []*command.Command {
	res := make([]*command.Command, 0, 1)
	res = append(res, &command.Command{Mode: "normal", CMD: "configure terminal"})
	res = append(res, &command.Command{
		Mode: "config",
		CMD:  "no router ospf " + OSPF,
	})

	res = append(res, &command.Command{Mode: "config", CMD: "exit"})
	return res
}

func (v8 V8500) OSPFRid(OSPF, Rid string) []*command.Command {
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

func (v8 V8500) NoOSPFRid(OSPF, Rid string) []*command.Command {
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

func (v8 V8500) OSPFNetworkArea(OSPF, Network, Area string) []*command.Command {
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

func (v8 V8500) NoOSPFNetworkArea(OSPF, Network, Area string) []*command.Command {
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

func (v8 V8500) OSPFInterfaceCost(OSPF, Interface, Cost string) []*command.Command {
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

func (v8 V8500) OSPFInterfaceDeadInterval(OSPF, Interface, DeadInterval string) []*command.Command {
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

func (v8 V8500) OSPFInterfaceHelloInterval(OSPF, Interface, HelloInterval string) []*command.Command {

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

func (v8 V8500) OSPFInterfaceRetransmitInterval(OSPF, Interface, RetransmitInterval string) []*command.Command {
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

func (v8 V8500) OSPFInterfaceTransmitDelay(OSPF, Interface, TransmitDelay string) []*command.Command {
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

func (v8 V8500) OSPFInterfacePriority(OSPF, Interface, Priority string) []*command.Command {
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

func (v8 V8500) OSPFInterfaceNetworktype(OSPF, Interface, Networktype string) []*command.Command {
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

func (v8 V8500) NoOSPFInterfaceCost(OSPF, Interface, Cost string) []*command.Command {
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

func (v8 V8500) NoOSPFInterfaceDeadInterval(OSPF, Interface, DeadInterval string) []*command.Command {
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

func (v8 V8500) NoOSPFInterfaceHelloInterval(OSPF, Interface, HelloInterval string) []*command.Command {
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

func (v8 V8500) NoOSPFInterfaceRetransmitInterval(OSPF, Interface, RetransmitInterval string) []*command.Command {
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

func (v8 V8500) NoOSPFInterfaceTransmitDelay(OSPF, Interface, TransmitDelay string) []*command.Command {
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

func (v8 V8500) NoOSPFInterfacePriority(OSPF, Interface, Priority string) []*command.Command {
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

func (v8 V8500) NoOSPFInterfaceNetworktype(OSPF, Interface, Networktype string) []*command.Command {
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

func (v8 V8500) OSPFReferenceBandwidth(OSPF, ReferenceBandwidth string) []*command.Command {
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

func (v8 V8500) NoOSPFReferenceBandwidth(OSPF, ReferenceBandwidth string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginate(OSPF, DefaultOriginate string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateRoutemap(OSPF, DefaultOriginate, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateMetric(OSPF, DefaultOriginate, Metric string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateMetrictype(OSPF, DefaultOriginate, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateMetricMetrictype(OSPF, DefaultOriginate, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateMetricRoutemap(OSPF, DefaultOriginate, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateMetrictypeRoutemap(OSPF, DefaultOriginate, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateAlways(OSPF, DefaultOriginate, Always string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateAlwaysRoutemap(OSPF, DefaultOriginate, Always, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateAlwaysMetric(OSPF, DefaultOriginate, Always, Metric string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateAlwaysMetrictype(OSPF, DefaultOriginate, Always, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateAlwaysMetricMetrictype(OSPF, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateAlwaysMetricRoutemap(OSPF, DefaultOriginate, Always, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateAlwaysMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginate(OSPF, DefaultOriginate string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateRoutemap(OSPF, DefaultOriginate, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateMetric(OSPF, DefaultOriginate, Metric string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateMetrictype(OSPF, DefaultOriginate, Metrictype string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateMetricMetrictype(OSPF, DefaultOriginate, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateMetricRoutemap(OSPF, DefaultOriginate, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateMetrictypeRoutemap(OSPF, DefaultOriginate, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateAlways(OSPF, DefaultOriginate, Always string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateAlwaysRoutemap(OSPF, DefaultOriginate, Always, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateAlwaysMetric(OSPF, DefaultOriginate, Always, Metric string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateAlwaysMetrictype(OSPF, DefaultOriginate, Always, Metrictype string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateAlwaysMetricMetrictype(OSPF, DefaultOriginate, Always, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateAlwaysMetricRoutemap(OSPF, DefaultOriginate, Always, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateAlwaysMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultOriginateAlwaysMetricMetrictypeRoutemap(OSPF, DefaultOriginate, Always, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFRedistribute(OSPF, Redistribute string) []*command.Command {
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

func (v8 V8500) OSPFRedistributeMetric(OSPF, Redistribute, Metric string) []*command.Command {
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

func (v8 V8500) OSPFRedistributeMetrictype(OSPF, Redistribute, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPFRedistributeRoutemap(OSPF, Redistribute, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFRedistributeMetricMetrictype(OSPF, Redistribute, Metric, Metrictype string) []*command.Command {
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

func (v8 V8500) OSPFRedistributeMetricRoutemap(OSPF, Redistribute, Metric, Routemap string) []*command.Command {
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

func (v8 V8500) OSPFRedistributeMetricMetrictypeRoutemap(OSPF, Redistribute, Metric, Metrictype, Routemap string) []*command.Command {
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

func (v8 V8500) NoOSPFRedistribute(OSPF, Redistribute string) []*command.Command {
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

func (v8 V8500) NoOSPFRedistributeMetric(OSPF, Redistribute, Metric string) []*command.Command {
	return v8.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v8 V8500) NoOSPFRedistributeMetrictype(OSPF, Redistribute, Metrictype string) []*command.Command {
	return v8.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v8 V8500) NoOSPFRedistributeRoutemap(OSPF, Redistribute, Routemap string) []*command.Command {
	return v8.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v8 V8500) NoOSPFRedistributeMetricMetrictype(OSPF, Redistribute, Metric, Metrictype string) []*command.Command {
	return v8.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v8 V8500) NoOSPFRedistributeMetricRoutemap(OSPF, Redistribute, Metric, Routemap string) []*command.Command {
	return v8.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v8 V8500) NoOSPFRedistributeMetricMetrictypeRoutemap(OSPF, Redistribute, Metric, Metrictype, Routemap string) []*command.Command {
	return v8.NoOSPFRedistribute(OSPF, Redistribute)
}

func (v8 V8500) OSPFSummary(OSPF, Summary string) []*command.Command {
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

func (v8 V8500) OSPFSummaryNoAdvertise(OSPF, Summary, NoAdvertise string) []*command.Command {
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

func (v8 V8500) NoOSPFSummary(OSPF, Summary string) []*command.Command {
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

func (v8 V8500) NoOSPFSummaryNoAdvertise(OSPF, Summary, NoAdvertise string) []*command.Command {
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

func (v8 V8500) OSPFDefaultMetric(OSPF, DefaultMetric string) []*command.Command {
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

func (v8 V8500) NoOSPFDefaultMetric(OSPF, DefaultMetric string) []*command.Command {
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

func (v8 V8500) OSPFPassive(OSPF, Passive string) []*command.Command {
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

func (v8 V8500) NoOSPFPassive(OSPF, Passive string) []*command.Command {
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

func (v8 V8500) OSPFAdminDistance(OSPF, AdminDistance string) []*command.Command {
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

func (v8 V8500) NoOSPFAdminDistance(OSPF, AdminDistance string) []*command.Command {
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

func (v8 V8500) OSPFDistanceExternal(OSPF, Distance, External string) []*command.Command {
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

func (v8 V8500) OSPFDistanceInter(OSPF, Distance, Inter string) []*command.Command {
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

func (v8 V8500) OSPFDistanceIntra(OSPF, Distance, Intra string) []*command.Command {
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

func (v8 V8500) OSPFDistanceInterIntra(OSPF, Distance, Inter, Intra string) []*command.Command {
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

func (v8 V8500) OSPFDistanceInterExternal(OSPF, Distance, Inter, External string) []*command.Command {
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

func (v8 V8500) OSPFDistanceInterIntraExternal(OSPF, Distance, Inter, Intra, External string) []*command.Command {
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

func (v8 V8500) NoOSPFDistanceExternal(OSPF, Distance, External string) []*command.Command {
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

func (v8 V8500) NoOSPFDistanceInter(OSPF, Distance, Inter string) []*command.Command {
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

func (v8 V8500) NoOSPFDistanceIntra(OSPF, Distance, Intra string) []*command.Command {
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

func (v8 V8500) NoOSPFDistanceInterIntra(OSPF, Distance, Inter, Intra string) []*command.Command {
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

func (v8 V8500) NoOSPFDistanceInterExternal(OSPF, Distance, Inter, External string) []*command.Command {
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

func (v8 V8500) NoOSPFDistanceInterIntraExternal(OSPF, Distance, Inter, Intra, External string) []*command.Command {
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

func (v8 V8500) OSPFDistributelistIN(OSPF, Distributelist, IN string) []*command.Command {
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

func (v8 V8500) OSPFDistributelistOUT(OSPF, Distributelist, OUT string) []*command.Command {
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

func (v8 V8500) NoOSPFDistributelistIN(OSPF, Distributelist, IN string) []*command.Command {
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

func (v8 V8500) NoOSPFDistributelistOUT(OSPF, Distributelist, OUT string) []*command.Command {
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

func (v8 V8500) OSPFAreaDefaultCost(OSPF, Area, DefaultCost string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaDefaultCost(OSPF, Area, DefaultCost string) []*command.Command {
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

func (v8 V8500) OSPFAreaNSSA(OSPF, Area, NSSA string) []*command.Command {
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

func (v8 V8500) OSPFAreaNSSADefaultOriginate(OSPF, Area, NSSA, DefaultOriginate string) []*command.Command {
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

func (v8 V8500) OSPFAreaNSSANoRedistribution(OSPF, Area, NSSA, Redistribution string) []*command.Command {
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

func (v8 V8500) OSPFAreaNSSANoSummary(OSPF, Area, NSSA, Summary string) []*command.Command {
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

func (v8 V8500) OSPFAreaNSSAStabilityInterval(OSPF, Area, NSSA, StabilityInterval string) []*command.Command {
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

func (v8 V8500) OSPFAreaTranslatorrole(OSPF, Area, Translatorrole string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaNSSA(OSPF, Area, NSSA string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaNSSADefaultOriginate(OSPF, Area, NSSA, DefaultOriginate string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaNSSANoRedistribution(OSPF, Area, NSSA, Redistribution string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaNSSANoSummary(OSPF, Area, NSSA, Summary string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaNSSAStabilityInterval(OSPF, Area, NSSA, StabilityInterval string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaTranslatorrole(OSPF, Area, Translatorrole string) []*command.Command {
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

func (v8 V8500) OSPFAreaStub(OSPF, Area, Stub string) []*command.Command {
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

func (v8 V8500) OSPFAreaStubNoSummary(OSPF, Area, Stub, NoSummary string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaStub(OSPF, Area, Stub string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaStubNoSummary(OSPF, Area, Stub, NoSummary string) []*command.Command {
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

func (v8 V8500) OSPFAreaRange(OSPF, Area, Range string) []*command.Command {
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

func (v8 V8500) OSPFAreaRangeAdvertise(OSPF, Area, Range, Advertise string) []*command.Command {
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

func (v8 V8500) OSPFAreaRangeNoAdvertise(OSPF, Area, Range, NoAdvertise string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaRange(OSPF, Area, Range string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaRangeAdvertise(OSPF, Area, Range, Advertise string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaRangeNoAdvertise(OSPF, Area, Range, NoAdvertise string) []*command.Command {
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

func (v8 V8500) OSPFAreaVirtuallink(OSPF, Area, Virtuallink string) []*command.Command {
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

func (v8 V8500) OSPFAreaVirtuallinkDeadInterval(OSPF, Area, Virtuallink, DeadInterval string) []*command.Command {
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

func (v8 V8500) OSPFAreaVirtuallinkHelloInterval(OSPF, Area, Virtuallink, HelloInterval string) []*command.Command {
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

func (v8 V8500) OSPFAreaVirtuallinkInstanceid(OSPF, Area, Virtuallink, Instanceid string) []*command.Command {
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

func (v8 V8500) OSPFAreaVirtuallinkRetransmitInterval(OSPF, Area, Virtuallink, RetransmitInterval string) []*command.Command {
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

func (v8 V8500) OSPFAreaVirtuallinkTransmitDelay(OSPF, Area, Virtuallink, TransmitDelay string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaVirtuallink(OSPF, Area, Virtuallink string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaVirtuallinkDeadInterval(OSPF, Area, Virtuallink, DeadInterval string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaVirtuallinkHelloInterval(OSPF, Area, Virtuallink, HelloInterval string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaVirtuallinkInstanceid(OSPF, Area, Virtuallink, Instanceid string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaVirtuallinkRetransmitInterval(OSPF, Area, Virtuallink, RetransmitInterval string) []*command.Command {
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

func (v8 V8500) NoOSPFAreaVirtuallinkTransmitDelay(OSPF, Area, Virtuallink, TransmitDelay string) []*command.Command {
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
