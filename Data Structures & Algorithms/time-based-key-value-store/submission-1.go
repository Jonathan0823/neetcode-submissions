type TimeMap struct {
	m map[string][]pair
}

type pair struct { 
	timestamp int
	value string
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]pair),
	}

}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	pairs, exist := this.m[key]
	if !exist {
		return ""
	}

	l, r := 0, len(pairs)-1
	ans := ""

	for l <= r {
		mid := l + (r-l)/2

		if pairs[mid].timestamp <= timestamp {
			ans = pairs[mid].value
			l = mid + 1 
		} else {
			r = mid - 1 
		}
	}

	return ans
}
