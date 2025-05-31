package main

func countConnections(groupSize int) int {
	return (groupSize - 1) * groupSize / 2
}
