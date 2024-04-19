package commonv1

import timestamppb "google.golang.org/protobuf/types/known/timestamppb"

func (x *Tournament_PlayersParticipants_Player) SetSeed(seed uint32) {
	if x == nil {
		return
	}
	x.Seed = seed
}

func (x *Tournament_TeamsParticipants_Team) SetSeed(seed uint32) {
	if x == nil {
		return
	}
	x.Seed = seed
}

func (x *Tournament_PlayersParticipants_Player) SetConfirmed(confirmed bool) {
	if x == nil {
		return
	}
	if confirmed {
		x.ConfirmTime = timestamppb.Now()
	} else {
		x.ConfirmTime = nil
	}
}

func (x *Tournament_TeamsParticipants_Team) SetConfirmed(confirmed bool) {
	if x == nil {
		return
	}
	if confirmed {
		x.ConfirmTime = timestamppb.Now()
	} else {
		x.ConfirmTime = nil
	}
}
