/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i, j int) bool { 
		return intervals[i].start < intervals[j].start
	})

	for i := 1; i < len(intervals); i++ { 
		if (intervals[i-1].end > intervals[i].start) { 
			return false
		}

	}
	return true

}
