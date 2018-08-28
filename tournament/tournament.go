package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

const (
	rtWin  = "win"
	rtDraw = "draw"
	rtLoss = "loss"
)

func validResult(result string) bool {
	return result == rtWin || result == rtDraw || result == rtLoss
}

var rmCommentsRE = regexp.MustCompile("#.*")

var (
	errInvalidLine   = errors.New("Invalid line format")
	errScanInput     = errors.New("Scan input error")
	errInvalidResult = errors.New("Invalid competition result")
)

func Tally(r io.Reader, w io.Writer) error {
	st := newScoreTable()
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		// Trim spaces
		line := strings.TrimSpace(scanner.Text())
		// Remove comments
		line = rmCommentsRE.ReplaceAllString(line, "")
		// Skip newline
		if line == "" {
			continue
		}

		// Parse result
		tokens := strings.Split(line, ";")
		if len(tokens) != 3 {
			return errInvalidLine
		}

		if !validResult(tokens[2]) {
			return errInvalidResult
		}

		// Tally
		st.tally(tokens[0], tokens[1], tokens[2])
	}

	if nil != scanner.Err() {
		return errScanInput
	}

	// Dump into writer
	return st.dump(w)
}

type score struct {
	team   string
	win    int
	lose   int
	draw   int
	points int
}

type scoreTable struct {
	// Mapping of team name to index of slice of score structs
	teams  map[string]int
	scores []*score
}

func newScoreTable() *scoreTable {
	return &scoreTable{
		teams: make(map[string]int),
	}
}

func (st *scoreTable) accumulatePoint(team, result string) {
	var s *score
	i := st.teams[team] - 1

	if i < 0 {
		// Create new score record for a new team
		s = &score{}
	} else {
		s = st.scores[i]
	}

	s.team = team
	switch result {
	case rtWin:
		s.win++
		s.points += 3
	case rtDraw:
		s.draw++
		s.points++
	case rtLoss:
		s.lose++
	}

	if i < 0 {
		// Append new score record
		st.scores = append(st.scores, s)
		st.teams[team] = len(st.scores)
	}
}

func (st *scoreTable) tally(team1, team2, result string) {
	r1, r2 := result, ""
	switch result {
	case rtWin:
		r2 = rtLoss
	case rtLoss:
		r2 = rtWin
	case rtDraw:
		r2 = rtDraw
	}

	st.accumulatePoint(team1, r1)
	st.accumulatePoint(team2, r2)
}

const fmtDump = "%%-31s|%%3%s |%%3%[1]s |%%3%[1]s |%%3%[1]s |%%3%[1]s\n"

var (
	fmtHead = fmt.Sprintf(fmtDump, "s")
	fmtRow  = fmt.Sprintf(fmtDump, "d")
)

func (st *scoreTable) dump(w io.Writer) error {
	fmt.Fprintf(w, fmtHead, "Team", "MP", "W", "D", "L", "P")
	st.sortScores()

	for _, s := range st.scores {
		fmt.Fprintf(w, fmtRow, s.team, s.draw+s.win+s.lose, s.win, s.draw, s.lose, s.points)
	}
	return nil
}

func (st *scoreTable) sortScores() {
	sort.Slice(st.scores, func(i, j int) bool {
		if st.scores[i].points == st.scores[j].points {
			return st.scores[i].team < st.scores[j].team
		}
		return st.scores[i].points > st.scores[j].points
	})
}
