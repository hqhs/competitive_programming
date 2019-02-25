package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func main() {
	//STDOUT MUST BE FLUSHED MANUALLY
	defer writer.Flush()

	var n, maxEmotion, maxSeqLen, fMax, sMax int
	var answer uint64
	scanf("%d %d %d\n", &n, &maxEmotion, &maxSeqLen)
	hap := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d ", &hap[i])
		if hap[i] >= fMax {
			sMax = fMax
			fMax = hap[i]
		} else if hap[i] > sMax {
			sMax = hap[i]
		}
	}
	/* find first and second max of hap, use first max maxSeqLen, then use second,
	then again max maxSeqLen and repeat until maxEmotion counter is done */
	// loop obviously don't make threw time restriction
	// emotionCounter := 0
	// for {
	// 	// if emotionCounter >= maxEmotion {
	// 	// 	break
	// 	// }
	// 	i = maxEmotion - emotionCounter // i > 0 always
	// 	printf("i: %d\n", i)
	// 	if i >= maxSeqLen {
	// 		emotionCounter += maxSeqLen + 1
	// 		answer += uint64(fMax)*uint64(maxSeqLen) + uint64(sMax)
	// 	} else {
	// 		// this case could happend only at the end of sequence
	// 		emotionCounter += i
	// 		answer += uint64(fMax) * uint64(i)
	// 		break
	// 	}
	// }
	cycles := maxEmotion / (maxSeqLen + 1)
	answer += (uint64(fMax)*uint64(maxSeqLen) + uint64(sMax)) * uint64(cycles)
	remains := maxEmotion % (maxSeqLen + 1)
	answer += uint64(fMax) * uint64(remains)
	printf("%d\n", answer)
}
