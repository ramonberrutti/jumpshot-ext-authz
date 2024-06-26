// Code generated by protoc-gen-jumpshot. DO NOT EDIT.
// source: storage/match.proto

package storage

import (
	"context"

	"google.golang.org/grpc"
)

var _ = context.Background()
var _ = grpc.ServiceRegistrar(nil)

func (m *TournamentMatch_Messages) JumpShotIndexClear() {
	// m.Id = ""
	// m.CreateBy = ""
	// m.TeamBy = ""
	// m.Message = ""
}

func (m *TournamentMatch) JumpShotIndexClear() {
	// m.Id = ""
	// m.TournamentId = ""
	// m.PhaseId = ""
	// m.MatchId = ""
	if m, ok := interface{}(m.Participants).(interface{ JumpShotIndexClear() }); ok {
		m.JumpShotIndexClear()
	}

	// m.Version = 0
	// m.Synced = false
	// m.State = 0
	if m, ok := interface{}(m.StartTime).(interface{ JumpShotIndexClear() }); ok {
		m.JumpShotIndexClear()
	}

	if m, ok := interface{}(m.EndTime).(interface{ JumpShotIndexClear() }); ok {
		m.JumpShotIndexClear()
	}

	if m, ok := interface{}(m.Messages).(interface{ JumpShotIndexClear() }); ok {
		m.JumpShotIndexClear()
	}

	// m.WinnersNextMatchIds = nil
	// m.LosersNextMatchIds = nil
	if m, ok := interface{}(m.CreateTime).(interface{ JumpShotIndexClear() }); ok {
		m.JumpShotIndexClear()
	}

	if m, ok := interface{}(m.UpdateTime).(interface{ JumpShotIndexClear() }); ok {
		m.JumpShotIndexClear()
	}

	if m, ok := interface{}(m.MatchConfig).(interface{ JumpShotIndexClear() }); ok {
		m.JumpShotIndexClear()
	}

	// m.ParticipantsMode = 0
}
