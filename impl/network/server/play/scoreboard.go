package server

import "github.com/minelc/go-server/api/network"

type DISPLAY string
type CRITERIA string
type MODE_OBJECTIVE byte

const (
	INTEGER DISPLAY = "integer"
	HEALTH  DISPLAY = "health"

	LIST      byte = 0
	SIDEBAR   byte = 1
	BELOWNAME byte = 2
	// 18: team specific sidebar, indexed as 3 + team color.

	CREATE       MODE_OBJECTIVE = 0
	DELETE       MODE_OBJECTIVE = 1
	DISPLAY_TEXT MODE_OBJECTIVE = 2 // Sidebar: Display title

	DUMMY           CRITERIA = "dummy"
	TRIGGER         CRITERIA = "trigger"
	DEATH           CRITERIA = "deathCount"
	PLAYER_KILL     CRITERIA = "playerKillCount"
	TOTAL_KILL      CRITERIA = "totalKillCount"
	HEALTH_CRITERIA CRITERIA = "health"
)

/*
Send or remove a sidebar line
Line - Line to remove or add
Objective - Objective name
Score - score of line. Lower = First line, Higher = Last line
Remove - Remove or change the line
*/
type PacketPlayOutScoreboardScore struct {
	Line      string
	Objective string
	Score     int32
	Remove    bool
}

func (p *PacketPlayOutScoreboardScore) UUID() int32 {
	return 60
}

func (p *PacketPlayOutScoreboardScore) Push(writer network.Buffer) {
	writer.PushTxt(p.Line)
	if p.Remove {
		writer.PushVrI(1)
		writer.PushTxt(p.Objective)
		return
	}

	writer.PushVrI(0)
	writer.PushTxt(p.Objective)
	writer.PushVrI(p.Score)
}

/*
Packet to send a objective.
ID - ObjectiveID
Objective - Objective name
ObjectiveDisplayName - Display name (used for sidebar title)
Display = Used for belowname (Integer or health)
*/
type PacketPlayOutScoreboardObjective struct {
	Objective            string
	ObjectiveDisplayName string
	Display              DISPLAY
	Id                   MODE_OBJECTIVE
}

func (p *PacketPlayOutScoreboardObjective) UUID() int32 {
	return 59
}

func (p *PacketPlayOutScoreboardObjective) Push(writer network.Buffer) {
	writer.PushTxt(p.Objective)
	writer.PushByt(byte(p.Id))
	if p.Id == 0 || p.Id == 2 {
		writer.PushTxt(p.ObjectiveDisplayName)
		writer.PushTxt(string(p.Display))
	}
}

/*
Packet to display the objective created.
ID - ObjectiveID
Objective - Objective name
*/
type PacketPlayOutScoreboardDisplayObjective struct {
	Objective string
	Id        byte
}

func (p *PacketPlayOutScoreboardDisplayObjective) UUID() int32 {
	return 61
}

func (p *PacketPlayOutScoreboardDisplayObjective) Push(writer network.Buffer) {
	writer.PushByt(p.Id)
	writer.PushTxt(p.Objective)
}
