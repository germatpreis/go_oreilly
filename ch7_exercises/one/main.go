package main

import (
	"io"
	"os"
	"sort"
)

type Ranker interface {
	Ranking() []RankingPosition
}

func RankPrinter(ranker Ranker, writer io.Writer) {
	for _, v := range ranker.Ranking() {
		io.WriteString(writer, v.name)
		writer.Write([]byte("\n"))
	}
}

type Team struct {
	TeamName    string
	MemberNames []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

type RankingPosition struct {
	wins int
	name string
}

func (l *League) MatchResult(teamOneName string, teamOneResult int, teamTwoName string, teamTwoResult int) {
	var winningTeamName string

	if teamOneResult > teamTwoResult {
		winningTeamName = teamOneName
	} else {
		winningTeamName = teamTwoName
	}

	_, ok := l.Wins[winningTeamName]
	if !ok {
		l.Wins[winningTeamName] = 1
	} else {
		l.Wins[winningTeamName] = l.Wins[winningTeamName] + 1
	}
}

func (l *League) Ranking() []RankingPosition {
	ranking := make([]RankingPosition, 0, len(l.Teams))

	for name, wins := range l.Wins {
		ranking = append(ranking, RankingPosition{wins: wins, name: name})
	}

	sort.Slice(ranking, func(i, j int) bool {
		return ranking[i].wins > ranking[j].wins
	})

	return ranking
}

func NewLeague(teams []Team) *League {
	wins := map[string]int{}

	for _, v := range teams {
		wins[v.TeamName] = 0
	}

	return &League{
		Teams: teams,
		Wins:  wins,
	}
}

func main() {
	t1 := Team{
		TeamName:    "cornholio",
		MemberNames: []string{"A", "B", "C"},
	}
	t2 := Team{
		TeamName:    "DukeNukem",
		MemberNames: []string{"C", "D", "E"},
	}
	league := NewLeague([]Team{t1, t2})
	league.MatchResult("cornholio", 3, "DukeNukem", 1)
	league.MatchResult("cornholio", 3, "DukeNukem", 1)
	league.MatchResult("cornholio", 3, "DukeNukem", 5)

	RankPrinter(league, os.Stdout)
}
