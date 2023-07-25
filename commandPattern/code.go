package commandpattern

import "fmt"

type RemoteController struct {
	Volume  string
	Channel string
}

func (r *RemoteController) AdjustVolume(volume string) {
	r.Volume = volume
}

func (r *RemoteController) SwitchChannel(channel string) {
	r.Channel = channel
}

type Command interface {
	Execute() string
}

type musicCommand struct {
	remoteController *RemoteController
}

func NewSetMusicCommand(remoteController *RemoteController) *musicCommand {
	return &musicCommand{remoteController: remoteController}
}

func (m *musicCommand) Execute() string {
	m.remoteController.SwitchChannel("music")
	m.remoteController.AdjustVolume("high")
	fmt.Printf("switch to music channel with %s volume \n", m.remoteController.Volume)
	return "switch successed"
}

type newsCommand struct {
	remoteController *RemoteController
}

func NewSetNewsCommand(remoteController *RemoteController) *newsCommand {
	return &newsCommand{remoteController: remoteController}
}

func (n *newsCommand) Execute() string {
	n.remoteController.SwitchChannel("news")
	n.remoteController.AdjustVolume("normal")
	fmt.Printf("switch to news channel with %s volume \n", n.remoteController.Volume)
	return "switch successed"
}

type remoteControllerInvoker struct {
	remoteCommand Command
}

func (r *remoteControllerInvoker) SetRemoteCommand(remoteCommand Command) {
	r.remoteCommand = remoteCommand
}

func (r *remoteControllerInvoker) ExecuteRemoteCommand() string {
	return r.remoteCommand.Execute()
}
