// SequenceSearch I/O 2012 - Go Concurrency Patterns
// https://go.dev/talks/2012/concurrency.slide
// https://www.youtube.com/watch?v=f6kdp27TYZs&ab_channel=GoogleDevelopers
package main

import (
	"github.com/denglitong/golang/concurrency/player"
	"github.com/denglitong/golang/concurrency/software"
)

func main() {
	player.ShowGoroutine()
	player.ShowChannel()
	player.ShowChannelPattern()
	player.ShowChannelAsHandlerOnService()
	player.ShowChannelFanInAndMultiplexing()
	player.ShowRestoringSequences()
	player.ShowChannelFanInAndMultiplexingWithSelect()
	player.ShowBranchTimeoutUsingSelect()
	player.ShowForLoopTimeoutUsingSelect()
	player.ShowChannelQuit()
	player.ShowChanQuitReceiveCleanUp()
	player.ShowDaisyChain()

	software.ShowSequenceSearch()
	software.ShowConcurrentSearch()
	software.ShowConcurrentSearchWithTimeout()
	software.ShowConcurrentSearchAvoidTimeoutReduceTailLatency()

	ShowCommunicateBySharingMemory()
}
