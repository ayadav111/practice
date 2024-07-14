package gasstation

func canCompleteStatation(gas []int, cost []int) int {

	start := 0
	total := 0
	tank := 0

	for i := 0; i < len(gas); i++ {

		tank = tank + gas[i] - cost[i]
		total = total + tank

		if tank < 0 {
			tank = 0
			start = i + 1
		}

	}

	if total < 0 {
		return -1
	}

	return start

}
