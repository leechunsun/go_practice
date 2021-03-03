package main

func SideFind(splits []int, calKey int) []int {
	if len(splits) == 2{
		if splits[0] < calKey{
			return splits[1:]
		} else {
			return splits[:1]
		}
	}
	if len(splits) == 1{
		return splits
	}
	mid := len(splits) / 2
	if splits[mid] > calKey{
		return SideFind(splits[:mid+1], calKey)
	}else if splits[mid] == calKey{
		return []int{calKey}
	}
	return SideFind(splits[mid:], calKey)
}

