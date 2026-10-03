/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	if len(intervals) < 1 {
		return  true
	}

	sort.Slice(intervals, func(i, j int) bool { 
		return intervals[i].start < intervals[j].end
	})

	for i := 0; i < len(intervals); i++ { 
		if i > 0 { 
			if intervals[i].start < intervals[i-1].end { 
				return false
			}
		}
	}
	return true

}
