package storage

func (t *Team) CanPlayerManage(playerID string) bool {
	if t == nil {
		return false
	}

	for _, player := range t.Players {
		if player.Id == playerID {
			return Team_Player_Roles(player.GetRoles()).HasAny(
				Team_Player_ADMIN,
				Team_Player_MANAGER,
			)
		}
	}

	return false
}

func (t *Team) CanPlayerJoin(playerId string) bool {
	if t == nil {
		return false
	}

	for _, player := range t.Invitations {
		if player.Id == playerId {
			return true
		}
	}

	return false
}

type Team_Player_Roles []Team_Player_Role

func (r Team_Player_Roles) Has(role Team_Player_Role) bool {
	if r == nil {
		return false
	}

	for _, r := range r {
		if r == role {
			return true
		}
	}

	return false
}

func (r Team_Player_Roles) HasAny(roles ...Team_Player_Role) bool {
	if r == nil {
		return false
	}

	for _, role := range roles {
		if r.Has(role) {
			return true
		}
	}

	return false
}

// Check if the player are part of the team.
func (t *Team) ArePlayersIn(playerIDs ...string) bool {
	if t == nil {
		return false
	}

	for _, playerID := range playerIDs {
		found := false
		for _, player := range t.Players {
			if player.Id == playerID {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}
