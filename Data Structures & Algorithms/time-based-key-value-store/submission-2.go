type pair struct { 
	value string
	timestamp int
}

type TimeMap struct {
	m map[string][]pair

}

func Constructor() TimeMap {
	return TimeMap{ 
		m: make(map[string][]pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], pair{
		value,
		timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exist := this.m[key]; !exist { 
		return ""
	}

	pairs := this.m[key]
	l, r := 0, len(pairs) - 1
	ans := ""
	for l <= r { 
		m := l + (r-l)/2

		if pairs[m].timestamp <= timestamp { 
			ans = pairs[m].value
			l = m+1
		} else { 
			r = m - 1
		}
	}

	return ans

}
