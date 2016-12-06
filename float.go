package quicksort

//float32

type floatQuickSort32 struct {
	val []float32
}

func (this *floatQuickSort32) Sort(val interface{}) {
	low := 0
	this.val = val.([]float32)
	hight := len(this.val) - 1
	this.sort(low, hight)
}

func (this *floatQuickSort32) sort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.sort(low, pivotpoint-1)
		this.sort(pivotpoint+1, hight)
	}
}

func (this *floatQuickSort32) medianOfThree(low, hight int) (float32, int) {
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

func (this *floatQuickSort32) partition(low, hight int, pivotpoint *int) {

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

// float64
type floatQuickSort64 struct {
	val []float64
}

func (this *floatQuickSort64) Sort(val interface{}) {
	low := 0
	this.val = val.([]float64)
	hight := len(this.val) - 1
	this.sort(low, hight)
}

func (this *floatQuickSort64) sort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.sort(low, pivotpoint-1)
		this.sort(pivotpoint+1, hight)
	}
}

func (this *floatQuickSort64) medianOfThree(low, hight int) (float64, int) {
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

func (this *floatQuickSort64) partition(low, hight int, pivotpoint *int) {

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
	register("float32", new(floatQuickSort32))
	register("float64", new(floatQuickSort64))
}
