package state

import "sync"

type BotState struct {
	state string
	mu    sync.Mutex
}

func (bs *BotState) GetState() string {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	return bs.state
}

func (bs *BotState) SetState(state string) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	bs.state = state
}

func (bs *BotState) Done() {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	bs.state = "*"
}

func NewBotState() *BotState {
	return &BotState{
		state: "*",
	}
}
