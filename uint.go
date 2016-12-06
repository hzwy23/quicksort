package quicksort

// uint
type uintQuickSort struct {
	val []uint
}

func (this *uintQuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]uint)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *uintQuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *uintQuickSort) medianOfThree(low, hight int) (uint, int) {
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

func (this *uintQuickSort) partition(low, hight int, pivotpoint *int) {

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

// uint8
type uint8QuickSort struct {
	val []uint8
}

func (this *uint8QuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]uint8)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *uint8QuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *uint8QuickSort) medianOfThree(low, hight int) (uint8, int) {
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

func (this *uint8QuickSort) partition(low, hight int, pivotpoint *int) {

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
type uint16QuickSort struct {
	val []uint16
}

func (this *uint16QuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]uint16)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *uint16QuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *uint16QuickSort) medianOfThree(low, hight int) (uint16, int) {
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

func (this *uint16QuickSort) partition(low, hight int, pivotpoint *int) {

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
type uint32QuickSort struct {
	val []uint32
}

func (this *uint32QuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]uint32)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *uint32QuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *uint32QuickSort) medianOfThree(low, hight int) (uint32, int) {
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

func (this *uint32QuickSort) partition(low, hight int, pivotpoint *int) {

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
type uint64QuickSort struct {
	val []uint64
}

func (this *uint64QuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]uint64)
	hight := len(this.val) - 1
	this.intsort(low, hight)
}

func (this *uint64QuickSort) intsort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.intsort(low, pivotpoint-1)
		this.intsort(pivotpoint+1, hight)
	}
}

func (this *uint64QuickSort) medianOfThree(low, hight int) (uint64, int) {
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

func (this *uint64QuickSort) partition(low, hight int, pivotpoint *int) {

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
	register("uint", new(uintQuickSort))
	register("uint8", new(uint8QuickSort))
	register("uint16", new(uint16QuickSort))
	register("uint32", new(uint32QuickSort))
	register("uint64", new(uint64QuickSort))
}
