package lc3

type State struct {
	Memory    []uint16
	Registers []uint16 //r1-r8
	N         uint16
	P         uint16
	Z         uint16
	PC        uint16 //pointer counter
}

func (s State) createComputer() {

}

func (s State) parseCommand(ir uint16) {
	command := getBits(ir, 0, 3)
	switch command {
	case 0:
		n := getBit(ir, 4)
		z := getBit(ir, 5)
		p := getBit(ir, 6)
		pcoff := getBits(ir, 7, 15)
		s.br(n, z, p, pcoff)
	case 1:
		dr := getBits(ir, 4, 6)
		bit5 := getBit(ir, 10)
		var val2 uint16
		sr1 := getBits(ir, 7, 9)
		val1 := s.Registers[sr1]
		if bit5 == 1 {
			val2 = getBits(ir, 11, 15)
			val2 = sext(val2, 4)
		} else {
			sr2 := getBits(ir, 13, 15)
			val2 = s.Registers[sr2]
		}
		result := s.add(val1, val2)
		s.Registers[dr] = result
		s.setCC(result)
	case 2:
		dr := getBits(ir, 4, 6)
		pcoff := getBits(ir, 7, 15)
		s.ld(dr, pcoff)
	case 3:
		sr := getBits(ir, 4, 6)
		pcoff := getBits(ir, 7, 15)
		s.st(sr, pcoff)
	case 5:
		dr := getBits(ir, 4, 6)
		bit5 := getBit(ir, 10)
		var val2 uint16
		sr1 := getBits(ir, 7, 9)
		val1 := s.Registers[sr1]
		if bit5 == 1 {
			val2 = getBits(ir, 11, 15)
			val2 = sext(val2, 4)
		} else {
			sr2 := getBits(ir, 13, 15)
			val2 = s.Registers[sr2]
		}
		result := s.and(val1, val2)
		s.Registers[dr] = result
		s.setCC(result)
	case 6:
		sr := getBits(ir, 4, 6)
		baseR := getBits(ir, 7, 9)
		offset6 := getBits(ir, 10, 15)
		s.ldr(sr, baseR, offset6)
	case 7:
		dr := getBits(ir, 4, 6)
		baseR := getBits(ir, 7, 9)
		offset6 := getBits(ir, 10, 15)
		s.str(dr, baseR, offset6)
	case 9:
		dr := getBits(ir, 4, 6)
		sr := getBits(ir, 7, 9)
		val := s.Registers[sr]
		result := s.not(val)
		s.Registers[dr] = result
	case 10:
		dr := getBits(ir, 4, 6)
		pcoff := getBits(ir, 7, 15)
		s.ldi(dr, pcoff)
	case 11:
		sr := getBits(ir, 4, 6)
		pcoff := getBits(ir, 7, 15)
		s.sti(sr, pcoff)
	case 14:
		dr := getBit(ir, 4, 6)
		pcoff := getBits(ir, 7, 15)
		s.lea(dr, pcoff)
	}

}

// Helpers
func getBits(value uint16, startIndex uint16, endIndex uint16) uint16 {
	maskedValue := value & 0xFFFF
	mask := uint16((1 << (endIndex - startIndex + 1)) - 1)
	return (maskedValue >> startIndex) & mask
}
func getBit(value uint16, index uint16) uint16 {
	maskedValue := value & 0xFFFF
	return (maskedValue >> (15 - index)) & 1
}
func sext(value uint16, signIndex uint16) uint16 {
	if (value>>signIndex)&1 == 0 {
		return value
	}
	mask := uint16(0)
	for i := signIndex; i < 16; i++ {
		mask |= 1 << i
	}
	return mask | value
}
func (s State) setCC(drValue uint16) {
	if drValue < 0 {
		s.N = 1
		s.Z = 0
		s.P = 0
	} else if drValue > 0 {
		s.N = 1
		s.Z = 0
		s.P = 0
	} else {
		s.N = 0
		s.Z = 1
		s.P = 0
	}
}

// OPERATE INSTRUCTIONS
func (s State) add(val1 uint16, val2 uint16) uint16 {
	return val1 + val2
}
func (s State) and(val1 uint16, val2 uint16) uint16 {
	return val1 & val2
}
func (s State) not(val uint16) uint16 {
	return ^val
}

// DATA MOVEMENT INSTRUCTIONS
/*
dr = destination register
PCoffset9 =  9 bit signed 2s compliemnts, sext to 16, added to the incremented pc to form an address --- range -256 to 256
*/
func (s State) ld(dr uint16, pcoffset9 uint16) {
	val := s.Memory[s.PC+sext(pcoffset9, 8)]
	s.Registers[dr] = val
	s.setCC(val)
}
func (s State) ldi(dr uint16, pcoffset9 uint16) {
	val := s.Memory[s.Memory[s.PC+sext(pcoffset9, 8)]]
	s.Registers[dr] = val
	s.setCC(val)
}

/*
dr = destistnation register
baseR = base register one of r0-7 which specs whioch resgister the result of an instruction shoudkl be written to
*/
func (s State) ldr(dr uint16, baseR uint16, offset6 uint16) {
	output := s.Memory[baseR+sext(offset6, 5)]
	s.Registers[dr] = output
	s.setCC(output)
}
func (s State) lea(dr uint16, pcoffset9 uint16) {
	value := s.Memory[s.PC+sext(pcoffset9, 8)]
	s.Registers[dr] = value
	s.setCC(value)
}

func (s State) st(sr uint16, pcoffset9 uint16) {
	val := s.Memory[s.PC+sext(pcoffset9, 8)]
	s.Memory[val] = s.Registers[sr]
}
func (s State) str(sr uint16, baseR uint16, offset6 uint16) {
	s.Memory[baseR+sext(offset6, 5)] = s.Registers[sr]
}
func (s State) sti(sr uint16, pcoffset9 uint16) {
	val := s.Memory[s.PC+sext(pcoffset9, 8)]
	s.Memory[s.Memory[val]] = s.Registers[sr]
}

// CONTROL INSTRUCTIONS
func (s State) br(n uint16, z uint16, p uint16, pcoffset9 uint16) {
	if (n & s.N) | (z & s.Z) | (p & s.P) {
		s.PC = s.PC + sext(pcoffset9, 8)
	}

}
func (s State) jsr() {

}
func (s State) jmp() {

}
func (s State) rti() {

}
func (s State) trap() {

}
