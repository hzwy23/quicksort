package quicksort

type intQuickSort struct {
	val []int
}

func (this *intQuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]int)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *intQuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *intQuickSort) medianOfThree(low, hight int) (int, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if this.val[low] > this.val[hight] {
			this.val[low], this.val[hight] = this.val[hight], this.val[low]
		}
		return this.val[low], low
	}
	if this.val[low] > this.val[m] {
		this.val[low], this.val[m] = this.val[m], this.val[low]
	}
	if this.val[low] > this.val[hight] {
		this.val[low], this.val[hight] = this.val[hight], this.val[low]
	}
	if this.val[m] > this.val[hight] {
		this.val[m], this.val[hight] = this.val[hight], this.val[m]
	}
	this.val[low+1], this.val[m] = this.val[m], this.val[low+1]
	return this.val[low+1], low + 1
}

func (this *intQuickSort) partition(low, hight int, pivotpoint *int) {

	pivotitem, j := this.medianOfThree(low, hight)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if this.val[i] < pivotitem {
			this.val[i], this.val[j] = this.val[j], this.val[i]
			j = i
		} else {
			this.val[max], this.val[i] = this.val[i], this.val[max]
			max--
			i--
		}
	}
	*pivotpoint = j

	/*
		var j = low
		//get first element as pivotitem.
		pivotitem := this.val[low]

		for i := low + 1; i <= hight; i++ {
			if this.val[i] < pivotitem {
				this.val[i], this.val[j] = this.val[j], this.val[i]
				j = i
			} else {
				for tj := hight; tj > i; tj-- {
					if this.val[tj] < pivotitem {
						this.val[tj], this.val[i], this.val[j] = this.val[i], this.val[j], this.val[tj]
						hight--
						j = i + 1
						break
					}
					hight--
				}
			}
		}
		*pivotpoint = j
	*/
}

// int8
type int8QuickSort struct {
	val []int8
}

func (this *int8QuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]int8)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *int8QuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *int8QuickSort) medianOfThree(low, hight int) (int8, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if this.val[low] > this.val[hight] {
			this.val[low], this.val[hight] = this.val[hight], this.val[low]
		}
		return this.val[low], low
	}
	if this.val[low] > this.val[m] {
		this.val[low], this.val[m] = this.val[m], this.val[low]
	}
	if this.val[low] > this.val[hight] {
		this.val[low], this.val[hight] = this.val[hight], this.val[low]
	}
	if this.val[m] > this.val[hight] {
		this.val[m], this.val[hight] = this.val[hight], this.val[m]
	}
	this.val[low+1], this.val[m] = this.val[m], this.val[low+1]
	return this.val[low+1], low + 1
}

func (this *int8QuickSort) partition(low, hight int, pivotpoint *int) {

	pivotitem, j := this.medianOfThree(low, hight)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if this.val[i] < pivotitem {
			this.val[i], this.val[j] = this.val[j], this.val[i]
			j = i
		} else {
			this.val[max], this.val[i] = this.val[i], this.val[max]
			max--
			i--
		}
	}
	*pivotpoint = j
}

// int16
type int16QuickSort struct {
	val []int16
}

func (this *int16QuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]int16)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *int16QuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *int16QuickSort) medianOfThree(low, hight int) (int16, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if this.val[low] > this.val[hight] {
			this.val[low], this.val[hight] = this.val[hight], this.val[low]
		}
		return this.val[low], low
	}
	if this.val[low] > this.val[m] {
		this.val[low], this.val[m] = this.val[m], this.val[low]
	}
	if this.val[low] > this.val[hight] {
		this.val[low], this.val[hight] = this.val[hight], this.val[low]
	}
	if this.val[m] > this.val[hight] {
		this.val[m], this.val[hight] = this.val[hight], this.val[m]
	}
	this.val[low+1], this.val[m] = this.val[m], this.val[low+1]
	return this.val[low+1], low + 1
}

func (this *int16QuickSort) partition(low, hight int, pivotpoint *int) {

	pivotitem, j := this.medianOfThree(low, hight)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if this.val[i] < pivotitem {
			this.val[i], this.val[j] = this.val[j], this.val[i]
			j = i
		} else {
			this.val[max], this.val[i] = this.val[i], this.val[max]
			max--
			i--
		}
	}
	*pivotpoint = j
}

// int32
type int32QuickSort struct {
	val []int32
}

func (this *int32QuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]int32)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *int32QuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *int32QuickSort) medianOfThree(low, hight int) (int32, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if this.val[low] > this.val[hight] {
			this.val[low], this.val[hight] = this.val[hight], this.val[low]
		}
		return this.val[low], low
	}
	if this.val[low] > this.val[m] {
		this.val[low], this.val[m] = this.val[m], this.val[low]
	}
	if this.val[low] > this.val[hight] {
		this.val[low], this.val[hight] = this.val[hight], this.val[low]
	}
	if this.val[m] > this.val[hight] {
		this.val[m], this.val[hight] = this.val[hight], this.val[m]
	}
	this.val[low+1], this.val[m] = this.val[m], this.val[low+1]
	return this.val[low+1], low + 1
}

func (this *int32QuickSort) partition(low, hight int, pivotpoint *int) {

	pivotitem, j := this.medianOfThree(low, hight)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if this.val[i] < pivotitem {
			this.val[i], this.val[j] = this.val[j], this.val[i]
			j = i
		} else {
			this.val[max], this.val[i] = this.val[i], this.val[max]
			max--
			i--
		}
	}
	*pivotpoint = j
}

// int64
type int64QuickSort struct {
	val []int64
}

func (this *int64QuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]int64)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *int64QuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *int64QuickSort) medianOfThree(low, hight int) (int64, int) {
	m := (hight + low) / 2
	if hight-low <= 1 {
		if this.val[low] > this.val[hight] {
			this.val[low], this.val[hight] = this.val[hight], this.val[low]
		}
		return this.val[low], low
	}
	if this.val[low] > this.val[m] {
		this.val[low], this.val[m] = this.val[m], this.val[low]
	}
	if this.val[low] > this.val[hight] {
		this.val[low], this.val[hight] = this.val[hight], this.val[low]
	}
	if this.val[m] > this.val[hight] {
		this.val[m], this.val[hight] = this.val[hight], this.val[m]
	}
	this.val[low+1], this.val[m] = this.val[m], this.val[low+1]
	return this.val[low+1], low + 1
}

func (this *int64QuickSort) partition(low, hight int, pivotpoint *int) {

	pivotitem, j := this.medianOfThree(low, hight)
	max := hight - 1
	for i := low + 2; i <= max; i++ {
		if this.val[i] < pivotitem {
			this.val[i], this.val[j] = this.val[j], this.val[i]
			j = i
		} else {
			this.val[max], this.val[i] = this.val[i], this.val[max]
			max--
			i--
		}
	}
	*pivotpoint = j
}

func init() {
	register("int", new(intQuickSort))
	register("int8", new(int8QuickSort))
	register("int16", new(int16QuickSort))
	register("int32", new(int32QuickSort))
	register("int64", new(int64QuickSort))
}
