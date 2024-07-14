package gasstation

func canCompleteStatation(gas []int, cost []int) int {

	start := 0
	total := 0
	tank := 0

	for i := 0; i < len(gas); i++ {

		tank = tank + gas[i] - cost[i]

		if tank < 0 {
			total = total + tank
			tank = 0
			start = i + 1
		}

	}

	if tank+total < 0 {
		return -1
	}

	return start

}
