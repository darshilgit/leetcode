package arraystring

func canCompleteCircuit(gas []int, cost []int) int {

	currentGas, totalGas := 0, 0
	start := 0

	for i := 0; i < len(gas); i++ {
		currentGas += gas[i] - cost[i]
		totalGas += gas[i] - cost[i]
		if currentGas < 0 {
			start = i + 1
			currentGas = 0
		}
	}
	if totalGas < 0 {
		return -1
	}
	return start
}
