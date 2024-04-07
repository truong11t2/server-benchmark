package main

// Bubble sort
func sort(list []Country) (sortedList []Country) {
	for i := 0; i < len(list)-1; i++ {
		for j := len(list) - 1; j > i; j-- {
			if list[j].Id < list[j-1].Id {
				swap(&list[j], &list[j-1])
			}
		}
	}
	return list
}

// Interchange sort
func sort1(list []Country) (sortedList []Country) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[j].Id < list[i].Id {
				swap(&list[j], &list[i])
			}
		}
	}
	return list
}

func swap(country1, country2 *Country) {
	temp := *country1
	*country1 = *country2
	*country2 = temp
}
