package services

func GetCategories() map[int][]int {
	var categorieslist = map[int][]int{
		10001: {20001, 20002, 20003, 20004, 20005},
		10002: {20019, 20020, 20021, 20022, 20023},
		10003: {20026},
		10004: {20027, 20028, 20029, 20031, 20032, 20033, 20034, 20035},
	}
	return categorieslist
}
