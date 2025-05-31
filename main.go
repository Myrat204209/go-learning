package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	str := [3]string{primary, secondary, tertiary}
	var cost [3]int
	cost[0] = len(primary)
	for i := 1; i < 3; i++ {
		cost[i] = cost[i-1] + len(str[i])
	}
	return str, cost

}
