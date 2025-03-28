package main

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type string
	MessageCharLimit int
}
func newUser(name string, membershipType string) User {
	msgChartLimit := 100
	if membershipType == "premium" {
		msgChartLimit = 1000
	}
	return User{
		Name: name,
		Membership: Membership{
			Type: membershipType,
			MessageCharLimit: msgChartLimit,
		},
	}
}
