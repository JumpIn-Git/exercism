package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	var wg sync.WaitGroup
	var mu sync.Mutex

	m := make(FreqMap)
	for _, line := range texts {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()

			for _, r := range line {
				mu.Lock()
				m[r]++
				mu.Unlock()
			}
		}(line)
	}

	wg.Wait()

	return m
}
