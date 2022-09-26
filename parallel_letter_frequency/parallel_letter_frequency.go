package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
    result:= FreqMap{}

	// SEQUENTIAL STARTS //
	// for _, s := range l{
	// 	m := Frequency(s)
	// 	for k,v := range m{
	// 		result[k] += v
	// 	}
	// }
	// SEQUENTIAL ENDS //

  var m chan  FreqMap = make(chan FreqMap)
	for _, s := range l{
		go func(str string){
			m <- Frequency(str)
		}(s)
	}

	for i := 0; i < len(l); i++{
		mapFromChannel :=  <- m
			for k, v := range mapFromChannel{
				result[k] += v
			}
	}

	return result
}
