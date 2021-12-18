package intel8080

import "testing"

func TestNOP(t *testing.T) {
	cpuCore := CPU_State{}
	cpuCore.init()

	prevPC := cpuCore.PC
	prevCycles := cpuCore.Cycles

	cpuCore.NOP()

	if cpuCore.PC != prevPC {
		t.Errorf("PC after execution is not correct! Expected PC: %X\n", prevPC)
		t.Log(cpuCore.dump_state())
	}

	if cpuCore.Cycles != prevCycles+4 {
		t.Errorf("Cycles after execution is not correct! Expected Cycles: %d\n", prevCycles+4)
		t.Log(cpuCore.dump_state())
	}
}

func TestSet_GetBC(t *testing.T) {
	cpuCore := CPU_State{}
	cpuCore.init()

	cpuCore.Set_BC(0x2010)

	if cpuCore.Get_B() != cpuCore.BC[0] {
		t.Errorf("Register B val wrong!: %X\n", cpuCore.BC[0])
		t.Log(cpuCore.dump_state())
	}

	if cpuCore.Get_BC() != 0x2010 {
		t.Errorf("RegisterPair BC val wrong!: %X\n", cpuCore.Get_BC())
		t.Log(cpuCore.dump_state())
	}
}

func TestSTAX(t *testing.T) {
	cpuCore := CPU_State{}
	cpuCore.init()

	prevPC := cpuCore.PC
	prevCycles := cpuCore.Cycles
	cpuCore.Set_BC(0x0010)
	cpuCore.Set_A(0x20)
	cpuCore.Opcode = 0x02

	cpuCore.STAX()

	if cpuCore.PC != prevPC {
		t.Errorf("PC after execution is not correct! Expected PC: %X\n", prevPC)
		t.Log(cpuCore.dump_state())
	}

	if cpuCore.Cycles != prevCycles+7 {
		t.Errorf("Cycles after execution is not correct! Expected Cycles: %d\n", prevCycles+7)
		t.Log(cpuCore.dump_state())
	}

	if cpuCore.Memory[0x0010] != 0x20 {
		t.Errorf("STAX didn't execute correctly! Memory: %X\n", cpuCore.Memory[0x0010])
		t.Log(cpuCore.dump_state())
	}
}
