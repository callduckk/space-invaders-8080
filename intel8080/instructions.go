package intel8080

func (core *CPU_State) fetch() {
	core.Opcode = core.Memory[core.PC]
	core.PC += 1
}

func (core *CPU_State) execute() {
	core.InstructionSet[core.Opcode]()
}

func (core *CPU_State) readLittleEndian() (immediateData uint16, immediateData_Low uint8, immediateData_High uint8) {
	immediateData_Low = core.Memory[core.PC]
	core.PC += 1
	immediateData_High = core.Memory[core.PC]
	core.PC += 1
	immediateData = (uint16(immediateData_High) << 8) & uint16(immediateData_Low)
	return
}

func (core *CPU_State) NOP() (length uint8, cycle uint8) {
	core.cycle(4)
	length = 1
	cycle = 4
	return
}

func (core *CPU_State) LXI() (length uint8, cycle uint8) {
	length = 3
	cycle = 10
	core.cycle(cycle)
	registerPair := core.Opcode >> 4
	immediateData, immediateData_Low, immediateData_High := core.readLittleEndian()
	switch registerPair {
	case 0:
		core.Set_B(immediateData_High)
		core.Set_C(immediateData_Low)
		break
	case 1:
		core.Set_D(immediateData_High)
		core.Set_E(immediateData_Low)
		break
	case 2:
		core.Set_H(immediateData_High)
		core.Set_L(immediateData_Low)
		break
	case 3:
		core.SP = immediateData
		break
	}
	return
}

func (core *CPU_State) STAX() (length uint8, cycle uint8) {
	length = 1
	cycle = 7
	core.cycle(cycle)
	filter := uint8(0b00010000)
	registerPair := (core.Opcode & filter) >> 4
	switch registerPair {
	case 0:
		core.Memory[core.Get_BC()] = core.Get_A()
		break
	case 1:
		core.Memory[core.Get_DE()] = core.Get_A()
		break
	}
	return
}
