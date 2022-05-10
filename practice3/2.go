package practice3

func Sort(list *[]int) {
	for j := 0; j < len(*list); j++ {
		for i := 0; i < len(*list)-j-1; i++ {
			if (*list)[i] > (*list)[i+1] {
				(*list)[i], (*list)[i+1] = (*list)[i+1], (*list)[i]
			}
		}
	}
}
