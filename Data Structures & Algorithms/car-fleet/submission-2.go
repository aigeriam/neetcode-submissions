

type Car struct {
	speed int
	pos   int
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}

	cars := make([]Car, n)
	for i := 0; i < n; i++ {
		cars[i] = Car{pos: position[i], speed: speed[i]}
	}

	// 1. Sort descending (closest to target first)
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos 
	})

	// 2. Change the stack to keep track of arrival times instead of Car objects
	fleetTimes := make([]float64, 0)

	for _, car := range cars {
		// Calculate precise standalone time to target
		time := float64(target-car.pos) / float64(car.speed)

		if len(fleetTimes) != 0 {
			timeofleading := fleetTimes[len(fleetTimes)-1]
			
			// 3. If the leading fleet takes LONGER or EQUAL time, this car catches up!
			if timeofleading >= time {
				// It merges into the fleet ahead. We do NOT pop or modify the leader.
				// The leader's arrival time still dictates this whole fleet.
				continue 
			}
		}
		
		// If it can't catch up, it forms a brand new slower fleet leader behind them
		fleetTimes = append(fleetTimes, time)
	}

	// The number of remaining times in our stack is the total number of fleets
	return len(fleetTimes)
}
/// 