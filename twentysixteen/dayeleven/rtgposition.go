package dayeleven

type RTGPosition struct {
	state RTGState
	steps int
}

func NewRTGPosition(state RTGState, steps int) *RTGPosition {
	return &RTGPosition{state: state, steps: steps}
}

func (p *RTGPosition) IsComplete() bool {
	return p.state.IsComplete()
}

func (p *RTGPosition) GetState() *RTGState {
	return &p.state
}

func (p *RTGPosition) GetSteps() int {
	return p.steps
}

func (p *RTGPosition) GenerateNextPositions(handler func(*RTGPosition)) {
	initialState := p.state
	initialSteps := p.steps

	for dir := -1; dir < 2; dir += 2 {
		currentFloor := initialState.GetElevator()
		newFloor := currentFloor + dir

		if newFloor < 0 || newFloor > 3 {
			continue
		}

		// Don't go down to totally empty floors.
		if currentFloor == 1 && newFloor == 0 && initialState.IsFloorEmpty(0) {
			continue
		} else if currentFloor == 2 &&
			newFloor == 1 &&
			initialState.IsFloorEmpty(0) &&
			initialState.IsFloorEmpty(1) {
			continue
		}

		for itemTwo := -1; itemTwo < 15; itemTwo++ {
			if itemTwo != -1 && !initialState.HasElement(currentFloor, itemTwo) {
				continue
			}
			for itemOne := itemTwo + 1; itemOne < 16; itemOne++ {
				if !initialState.HasElement(currentFloor, itemOne) {
					continue
				}

				if (itemTwo != -1 && itemOne > 7 && itemTwo < 8 && itemOne-8 != itemTwo) ||
					(itemTwo != -1 && itemOne < 8 && itemTwo > 7 && itemTwo-8 != itemOne) {
					continue
				}

				newState := initialState
				newState.MoveElement(currentFloor, newFloor, itemOne)
				if itemTwo != -1 {
					newState.MoveElement(currentFloor, newFloor, itemTwo)
				}
				newState.SetElevator(newFloor)

				if newState.IsValid() {
					handler(&RTGPosition{
						state: newState,
						steps: initialSteps + 1})
				}
			}
		}
	}
}
