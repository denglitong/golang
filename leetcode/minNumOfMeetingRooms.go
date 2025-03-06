package leetcode

import "slices"

// 253.会议室 II
// 给定一个会议时间安排的数组 intervals ，
// 每个会议时间都会包括开始和结束的时间 intervals[i] = [starti, endi]，
// 为避免会议冲突，同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。

type Meeting struct {
	start, end int
}

type TimeSlot struct {
	time     int
	timeType string // start | end
}

func minNumOfMeetingRooms(meetings []Meeting) int {
	timeSlots := make([]TimeSlot, len(meetings)*2)
	for _, m := range meetings {
		timeSlots = append(timeSlots, TimeSlot{m.start, "start"})
		timeSlots = append(timeSlots, TimeSlot{m.end, "end"})
	}

	slices.SortFunc(timeSlots, func(a, b TimeSlot) int {
		if a.time == b.time {
			if a.timeType == "start" {
				return 1
			}
			return -1
		}
		return b.time - a.time
	})

	meetingRooms, meetingStarted := 0, 0
	for _, t := range timeSlots {
		if t.timeType == "start" {
			meetingStarted++
			meetingRooms++
		} else if t.timeType == "end" {
			meetingRooms--
		}
		if meetingStarted == len(meetings) {
			break
		}
	}
	return meetingRooms
}
