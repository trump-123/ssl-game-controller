package statemachine

import "github.com/trump-123/ssl-game-controller/internal/app/state"

func (s *StateMachine) processChangeSetBallPlacementPos(newState *state.State, change *Change_SetBallPlacementPos) (changes []*Change) {
	newState.PlacementPos = change.Pos
	return
}
