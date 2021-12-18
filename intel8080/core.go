package intel8080

import "fmt"

type CPU_State struct {
	AF, BC, DE, HL [2]uint8
	SP, PC         uint16
	Memory         []uint8
	Cycles         uint64
	InstructionSet []func() (length uint8, cycle uint8)
	Opcode         uint8
}

const stack_start = 0xF000

func (core *CPU_State) Set_AF(val uint16) {
	core.AF[0] = uint8(val >> 8)
	core.AF[1] = uint8(0x00FF & val)
}

func (core *CPU_State) Set_BC(val uint16) {
	core.BC[0] = uint8(val >> 8)
	core.BC[1] = uint8(0x00FF & val)
}

func (core *CPU_State) Set_DE(val uint16) {
	core.DE[0] = uint8(val >> 8)
	core.DE[1] = uint8(0x00FF & val)
}

func (core *CPU_State) Set_HL(val uint16) {
	core.HL[0] = uint8(val >> 8)
	core.HL[1] = uint8(0x00FF & val)
}

func (core *CPU_State) Get_AF() uint16 {
	return uint16(core.AF[0])<<8 | uint16(core.AF[1])
}

func (core *CPU_State) Get_BC() uint16 {
	return uint16(core.BC[0])<<8 | uint16(core.BC[1])
}

func (core *CPU_State) Get_DE() uint16 {
	return uint16(core.DE[0])<<8 | uint16(core.DE[1])
}

func (core *CPU_State) Get_HL() uint16 {
	return uint16(core.HL[0])<<8 | uint16(core.HL[1])
}

func (core *CPU_State) Get_A() uint8 {
	return core.AF[0]
}

func (core *CPU_State) Get_F() uint8 {
	return core.AF[1]
}

func (core *CPU_State) Get_B() uint8 {
	return core.BC[0]
}

func (core *CPU_State) Get_C() uint8 {
	return core.BC[1]
}

func (core *CPU_State) Get_D() uint8 {
	return core.DE[0]
}

func (core *CPU_State) Get_E() uint8 {
	return core.DE[1]
}

func (core *CPU_State) Get_H() uint8 {
	return core.HL[0]
}

func (core *CPU_State) Get_L() uint8 {
	return core.HL[1]
}

func (core *CPU_State) Set_A(val uint8) {
	core.AF[0] = val
}

func (core *CPU_State) Set_F(val uint8) {
	core.AF[1] = val
}

func (core *CPU_State) Set_B(val uint8) {
	core.BC[0] = val
}

func (core *CPU_State) Set_C(val uint8) {
	core.BC[1] = val
}

func (core *CPU_State) Set_D(val uint8) {
	core.DE[0] = val
}

func (core *CPU_State) Set_E(val uint8) {
	core.DE[1] = val
}

func (core *CPU_State) Set_H(val uint8) {
	core.HL[0] = val
}

func (core *CPU_State) Set_L(val uint8) {
	core.HL[1] = val
}

func (core *CPU_State) cycle(count uint8) {
	core.Cycles += uint64(count)
}

func (core *CPU_State) init() {
	core.Memory = make([]uint8, 8*1024)
	core.PC = 0x2000
	core.SP = stack_start

	core.map_instructions()
}

func (core *CPU_State) map_instructions() {
	core.InstructionSet = []func() (length uint8, cycle uint8){0x00: core.NOP, 0x10: core.NOP, 0x20: core.NOP, 0x30: core.NOP, 0x01: core.LXI, 0x11: core.LXI, 0x21: core.LXI, 0x31: core.LXI, 0x02: core.STAX, 0x12: core.STAX}
}

func (core *CPU_State) dump_state() string {
	var state string
	state += fmt.Sprintf("A: %X  B: %X  C: %X  D: %X  E: %X  F: %X  H: %X  L: %X\n", core.Get_A(), core.Get_B(), core.Get_C(), core.Get_D(), core.Get_E(), core.Get_F(), core.Get_H(), core.Get_L())
	state += fmt.Sprintf("AF: %X  BC: %X  DE: %X  HL: %X\n", core.Get_AF(), core.Get_BC(), core.Get_DE(), core.Get_HL())
	state += "F: SZ0A0P1C\n"
	state += fmt.Sprintf("F: %b\n", core.AF[1])
	state += fmt.Sprintf("Cycles: %d\n", core.Cycles)
	return state
}
