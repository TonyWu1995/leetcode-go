package leetcode1396

type UndergroundSystem struct {
	m    map[int]StationTime
	time map[[2]string][2]int
}

type StationTime struct {
	station string
	time    int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		m:    make(map[int]StationTime),
		time: make(map[[2]string][2]int),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.m[id] = StationTime{
		time:    t,
		station: stationName,
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	s := this.m[id]
	delete(this.m, id)
	key := [2]string{s.station, stationName}
	value := this.time[key]
	this.time[key] = [2]int{value[0] + t - s.time, value[1] + 1}
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	key := [2]string{startStation, endStation}
	value := this.time[key]
	return float64(value[0]) / float64(value[1])
}
