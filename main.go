package main

func main() {
	// First task
	purchase := 3333_33
	percent := 1
	limit := 100

	bonus := purchase / 10000 * percent // вместо нуля ваша формула
	if bonus > limit {
		bonus = limit
	}
	println(bonus)
	//////////////////////////////////////////

	// Second task
	var balance int64 = 15_000_000_00 // 15 миллионов в копейках
	var transfer int64 = 10_000_000_00 // 10 миллионов в копейках
	total := (balance + transfer) / 100 // Вывод в рублях, без копеек
	println(total) //
}
