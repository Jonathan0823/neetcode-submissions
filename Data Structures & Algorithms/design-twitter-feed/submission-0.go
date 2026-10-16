type TweetHeap [][]int

func (h TweetHeap) Len() int { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h TweetHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x any) { 
	*h = append(*h, x.([]int))
}
func (h *TweetHeap) Pop() any { 
	old := *h
	x := old[len(old) - 1]
	*h = old[:len(old) - 1]
	return x
}

type Twitter struct {
	count int
	followMap map[int]map[int]bool
	tweetMap map[int][][2]int
}


func Constructor() Twitter {
	return Twitter{ 
		count: 0,
		followMap: make(map[int]map[int]bool),
		tweetMap: make(map[int][][2]int),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	if _, exist := this.tweetMap[userId]; !exist { 
		this.tweetMap[userId] = make([][2]int, 0, 10)
	}

	this.tweetMap[userId] = append(this.tweetMap[userId], [2]int{this.count, tweetId})
	this.count++     
}


func (this *Twitter) GetNewsFeed(userId int) []int {
	res :=  make([]int, 0, 10)

	if _, ok := this.followMap[userId]; !ok { 
		this.followMap[userId] = make(map[int]bool)
	}
	this.followMap[userId][userId] = true

	var h TweetHeap
	heap.Init(&h)

	for fid := range this.followMap[userId] { 
		if tweets, exist := this.tweetMap[fid]; exist && len(tweets) > 0 { 
			lastIdx := len(tweets) - 1
			// push count, tweetid, followerid, nextidx
			heap.Push(&h, []int{tweets[lastIdx][0], tweets[lastIdx][1], fid, lastIdx-1})
		}
	}

	for h.Len() > 0 && len(res) < 10 { 
		top := heap.Pop(&h).([]int)
		tId := top[1]
		fId := top[2]
		nxtId := top[3]

		res = append(res, tId)

		if nxtId >= 0 { 
			older := this.tweetMap[fId][nxtId]
			heap.Push(&h, []int{older[0], older[1], fId, nxtId-1})
		}
	}
    
	return res
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	if _, ok := this.followMap[followerId]; !ok { 
		this.followMap[followerId] = make(map[int]bool)
	}
	this.followMap[followerId][followeeId] = true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	if _, ok := this.followMap[followerId]; ok { 
		delete(this.followMap[followerId], followeeId)
	}
    
}
