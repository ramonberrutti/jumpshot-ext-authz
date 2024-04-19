package storage

import (
	commonv1pb "github.com/ramonberrutti/jumpshot/protogen/common/v1"
)

// Check if the tournament is for players only.
func (t *Tournament) IsForPlayersOnly() bool {
	if t == nil {
		return false
	}

	return t.ParticipantsMode == commonv1pb.Tournament_PLAYERS
}

// Check if the tournament is for teams only.
func (t *Tournament) IsForTeamsOnly() bool {
	if t == nil {
		return false
	}

	return t.ParticipantsMode == commonv1pb.Tournament_TEAMS
}

func (x *Tournament) IsCreated() bool {
	if x == nil {
		return false
	}

	return x.State == commonv1pb.Tournament_CREATED
}

func (x *Tournament) IsOpen() bool {
	if x == nil {
		return false
	}

	return x.State == commonv1pb.Tournament_JOIN_OPEN
}
