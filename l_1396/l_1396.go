package leetcode

type stationTime struct {
	sum, count int
}

type checkInTime struct {
	startStation string
	time         int
}

type UndergroundSystem struct {
	allRecords map[string]map[string]*stationTime
	currRecord map[int]*checkInTime
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		make(map[string]map[string]*stationTime),
		make(map[int]*checkInTime),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	if this.currRecord[id] == nil {
		this.currRecord[id] = new(checkInTime)
		this.currRecord[id].startStation = stationName
		this.currRecord[id].time = t
	}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	if currRecord := this.currRecord[id]; currRecord != nil {
		startStation := currRecord.startStation
		deltaTime := t - currRecord.time
		delete(this.currRecord, id)
		if this.allRecords[startStation] == nil {
			this.allRecords[startStation] = make(map[string]*stationTime)
		}
		record := this.allRecords[startStation][stationName]
		if record == nil {
			this.allRecords[startStation][stationName] = new(stationTime)
		}
		record.sum += deltaTime
		record.count += 1
	}
}

func (this *UndergroundSystem) GetAverageTime(startStation, endStation string) float64 {
	if _, ok := this.allRecords[startStation]; !ok {
		return 0
	}
	if _, ok := this.allRecords[startStation][endStation]; !ok {
		return 0
	}
	record := this.allRecords[startStation][endStation]
	return float64(record.sum) / float64(record.count)
}
