package quicksort

type stringQuickSort struct {
	val []string
}

func (this *stringQuickSort) Sort(val interface{}) {
	low := 0
	this.val = val.([]string)
	hight := len(this.val) - 1
	this.sort(low, hight)
}

func (this *stringQuickSort) sort(low, hight int) {
	// get first element as pivotitem
	var pivotpoint int
	if hight > low {
		this.partition(low, hight, &pivotpoint)
		this.sort(low, pivotpoint-1)
		this.sort(pivotpoint+1, hight)
	}
}

func (this *stringQuickSort) medianOfThree(low, hight int) (string, int) {
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

func (this *stringQuickSort) partition(low, hight int, pivotpoint *int) {

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
	register("string", new(stringQuickSort))
}
