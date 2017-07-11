package main

import (
	"command"
	"dsl"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Printf("%#v\n", dsl.ParserV8)
	file, err := ioutil.ReadFile("instruction.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	ins := make([]string, 0, 10)

	for _, l := range lines {
		fmt.Println(l)
		if l == "" {
			continue
		}
		ins = append(ins, l)
	}

	for _, i := range ins {
		fmt.Println(string(dsl.MapInstructionToKey(i)))
	}

	for _, i := range ins {
		fmt.Printf("%#v\n", dsl.GetInstructionParamsFromCMD(i))
	}

	for _, i := range ins {
		fmt.Println(i)
		cmds, err := dsl.Engine.Parse("V8", &command.Command{CMD: "$" + i})
		if err == nil {
			for _, c := range cmds {
				fmt.Printf("%#v", c)
			}
		} else {
			fmt.Printf("%#v", err)
		}
	}

	cmds, err := dsl.Engine.Parse("V8", &command.Command{
		Mode: "config",
		CMD:  "show running-config",
	})

	if err != nil {
		panic(err)
	}

	for _, c := range cmds {
		fmt.Printf("%#v\n", c)
	}

	TestInstruction("$OSPF6(1234)")
	TestInstruction("$OSPF6(1234).Interface(100)")
	TestInstruction("$OSPF6(1234).Interface(100).Cost(100)")
	TestInstruction("$No().OSPF6(1234).Area(789).Stub()")
	TestInstruction("$No().OSPF6(1234).Area(456).Stub().NoSummary()")

	TestInstruction("$OSPF6(3456).Area(1111).NSSA()")
	TestInstruction("$OSPF6(3456).Area(1111).NSSA().DefaultOriginate()")
	TestInstruction("$OSPF6(3456).Area(1111).NSSA().NoRedistribution()")
	TestInstruction("$OSPF6(3456).Area(1111).NSSA().NoSummary()")
	TestInstruction("$OSPF6(3456).Area(1111).NSSA().StabilityInterval(4321)")
	TestInstruction("$OSPF6(3456).Area(1111).Translatorrole(candidate)")
	TestInstruction("$No().OSPF6(3456).Area(1111).NSSA()")
	TestInstruction("$No().OSPF6(3456).Area(1111).NSSA().DefaultOriginate()")
	TestInstruction("$No().OSPF6(3456).Area(1111).NSSA().NoRedistribution()")
	TestInstruction("$No().OSPF6(3456).Area(1111).NSSA().NoSummary()")
	TestInstruction("$No().OSPF6(3456).Area(1111).NSSA().StabilityInterval(4321)")
	TestInstruction("$No().OSPF6(3456).Area(1111).Translatorrole(always)")

	TestInstruction("$OSPF6(1234).Area(4567).Virtuallink(10.10.10.10)")
	TestInstruction("$OSPF6(1234).Area(4567).Virtuallink(10.1.1.1).DeadInterval(40)")
	TestInstruction("$OSPF6(1234).Distance().External(10000)")
	TestInstruction("$OSPF6(1234).Distance().Inter(10000)")
	TestInstruction("$OSPF6(1234).Distance().Intra(10000)")
	TestInstruction("$OSPF6(1234).Distance().Inter(10000).Intra(10000)")
	TestInstruction("$OSPF6(1234).Distance().Inter(10000).External(10000)")
	TestInstruction("$OSPF6(1234).Distance().Inter(10000).Intra(10000).External(10000)")
	TestInstruction("$No().OSPF6(1234).Distance().External(10000)")
	TestInstruction("$No().OSPF6(1234).Distance().Inter(10000)")
	TestInstruction("$No().OSPF6(1234).Distance().Intra(10000)")
	TestInstruction("$No().OSPF6(1234).Distance().Inter(10000).Intra(10000)")
	TestInstruction("$No().OSPF6(1234).Distance().Inter(10000).External(10000)")
	TestInstruction("$OSPF6(3456).Area(22222).Virtuallink(12.1.1.1)")
	TestInstruction("$OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).DeadInterval(#Interval)")
	TestInstruction("$OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).HelloInterval(4321)")
	TestInstruction("$OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).Instanceid(3456)")
	TestInstruction("$OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).RetransmitInterval(4321)")
	TestInstruction("$OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).TransmitDelay(3333)")
	TestInstruction("$No().OSPF6(3456).Area(22222).Virtuallink(12.1.1.1)")
	TestInstruction("$No().OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).DeadInterval(4321)")
	TestInstruction("$No().OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).HelloInterval(4321)")
	TestInstruction("$No().OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).Instanceid(3456)")
	TestInstruction("$No().OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).RetransmitInterval(4321)")
	TestInstruction("$No().OSPF6(3456).Area(22222).Virtuallink(12.1.1.1).TransmitDelay(3333)")

	TestInstruction("$Port(3).Enable()")
	TestInstruction("$Port(3).Disable()")
	TestInstruction("$Port(3).Pvid()")
	TestInstruction("$Port(3).Speed()")
	TestInstruction("$VLAN(3)")
	TestInstruction("$No().VLAN(3)")
	TestInstruction("$VLAN(124).Add(4)")
	TestInstruction("$VLAN(124).AddT(4)")
	TestInstruction("$VLAN(124).Del(4)")
	TestInstruction("$VLAN(124).DelT(4)")
	TestInstruction("$VLAN(124).Shutdown()")
	TestInstruction("$VLAN(124).NoShutdown()")
	TestInstruction("$VLAN(124).IP(100.1.1.1/24)")
	TestInstruction("$No().VLAN(124).IP(100.1.1.1/24)")
	TestInstruction("$VLAN(124).IP2(200.1.1.1/24)")
	TestInstruction("$No().VLAN(124).IP2(200.1.1.1/24)")
	TestInstruction("$VLAN(124).AddT(4).IP(100.1.1.1/24)")
	TestInstruction("$VLAN(124).AddT(4).IP2(200.1.1.1/24)")
	TestInstruction("$VLAN(124).AddT(4).IP(100.1.1.1/24).NoShutdown()")
	TestInstruction("$VLAN(124).AddT(4).IP2(200.1.1.1/24).NoShutdown()")
	TestInstruction("$VLAN(124).AddUT(4).IP(100.1.1.1/24)")
	TestInstruction("$VLAN(124).AddUT(4).IP2(200.1.1.1/24)")
	TestInstruction("$VLAN(124).AddUT(4).IP(100.1.1.1/24).NoShutdown()")
	TestInstruction("$VLAN(124).AddUT(4).IP2(200.1.1.1/24).NoShutdown()")
	TestInstruction("$VLAN(124).IP6().Enable()")
	TestInstruction("$No().VLAN(124).IP6().Enable()")
	TestInstruction("$VLAN(124).IP6(2001:db8:1000:1000::1/64)")
	TestInstruction("$No().VLAN(124).IP6(2001:db8:1000:1000::1/64)")
	TestInstruction("$VLAN(124).IP6LL(fe80:0001:0002:0003::1)")
	TestInstruction("$VLAN(124).IP6LL(fe80:0001:0002:0003::1).IP6(2001:db8:1000:1000::1/64)")
	TestInstruction("$No().VLAN(124).IP6LL(fe80:0001:0002:0003::1)")

}

func TestInstruction(ins string) {
	cmds, err := dsl.Engine.Parse("V8", &command.Command{
		CMD: ins,
	})

	if err != nil {
		panic(err)
	}

	for _, c := range cmds {
		fmt.Printf("%#v\n", c)
	}
}
