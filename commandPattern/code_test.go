package commandpattern

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	remoteController := new(RemoteController)
	remoteControllerInvoker := new(remoteControllerInvoker)

	setMusicChannel := NewSetMusicCommand(remoteController)
	remoteControllerInvoker.SetRemoteCommand(setMusicChannel)
	fmt.Println(remoteControllerInvoker.ExecuteRemoteCommand())

	setNewsChannel := NewSetNewsCommand(remoteController)
	remoteControllerInvoker.SetRemoteCommand(setNewsChannel)
	fmt.Println(remoteControllerInvoker.ExecuteRemoteCommand())
}
