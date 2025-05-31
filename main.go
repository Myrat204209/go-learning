package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	var result = make([]float64,0,len(costs)) 
	for _, c := range costs {
		if c.day == day {
			result = append(result,c.value) 
		}
	}
	return result
}
