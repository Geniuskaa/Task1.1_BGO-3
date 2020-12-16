package main

func main() {
	purchase := 3333_33
	percent := 1
	limit := 100

	bonus := purchase / 10000 * percent // вместо нуля ваша формула
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
}
