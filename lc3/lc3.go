package lc3

type State struct {
	Memory []int16
	R0     int16
	R1     int16
	R2     int16
	R3     int16
	R4     int16
	R5     int16
	R6     int16
	R7     int16
}

type Register struct {
}

func (s State) parseCommand(ir BitSet) {
	command := ir[0:4]

}

// OPERATE INSTRUCTIONS
func (s State) add() {

}
func (s State) and() {

}
func (s State) not() {

}

// DATA MOVEMENT INSTRUCTIONS
func (s State) ld() {

}
func (s State) ldi() {

}
func (s State) ldr() {

}
func (s State) lea() {

}
func (s State) st() {

}
func (s State) str() {

}
func (s State) sti() {

}

// CONTROL INSTRUCTIONS
func (s State) br() {

}
func (s State) jsr() {

}
func (s State) jmp() {

}
func (s State) rti() {

}
func (s State) trap() {

}
