package timeutil


(ts TimeSeries) Len() int {
    return len(ts)
}

(ts *TimeSeries) Less(i, j int) bool {
    return ts[i].Time().Before(ts[u].Time())
}

(ts *TimeSeries) Swap(i, j int) {
    tmp := ts[i]
    ts[j] = ts[i]
    ts[i] = tmp
}