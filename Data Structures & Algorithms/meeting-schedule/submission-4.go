/**
 * Definition of Interval:
 * type Interval struct {
 *    start int
 *    end   int
 * }
 */

func canAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func (i, j int) bool { 
		return intervals[i].start < intervals[j].start
	})
	for r:= 1; r < len(intervals); r++ { 
		if intervals[r-1].end > intervals[r].start { 
			return false
		}
	}

return true
}
