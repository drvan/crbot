package vote

import (
	"fmt"
	"math"
	"time"

	"github.com/jakevoytko/crbot/model"
)

// Returns the full status line. To be reused in the ballot executor.
func StatusLine(clock model.UTCClock, vote *Vote) string {
	// Add the vote totals.
	statusStr := statusString(vote)
	votesFor := len(vote.VotesFor)
	votesAgainst := len(vote.VotesAgainst)
	votesForStr := MsgOneVoteFor
	if votesFor != 1 {
		votesForStr = fmt.Sprintf(MsgVotesFor, votesFor)
	}
	votesAgainstStr := MsgOneVoteAgainst
	if votesAgainst != 1 {
		votesAgainstStr = fmt.Sprintf(MsgVotesAgainst, votesAgainst)
	}

	timeString := TimeString(clock, vote.TimestampEnd)

	return statusStr + ". " + votesForStr + ", " + votesAgainstStr + ". " + timeString
}

func statusString(vote *Vote) string {
	if vote.HasEnoughVotes() {
		switch vote.CalculateActiveStatus() {
		case VoteOutcomePassed:
			return MsgStatusVotePassing

		default:
			return MsgStatusVoteFailing
		}
	}
	return MsgStatusVotesNeeded
}

const (
	MsgNoTimeRemaining       = "No time remaining in vote"
	MsgMinutesRemaining      = "%v minutes remaining"
	MsgSecondsRemaining      = "%v seconds remaining"
	MsgMillisecondsRemaining = "%v milliseconds remaining"
)

func TimeString(clock model.UTCClock, timestampEnd time.Time) string {
	currentTime := clock.Now()
	remaining := timestampEnd.Sub(currentTime)
	timeString := MsgNoTimeRemaining

	// This shows the rounded-up time. Some examples:
	// time.Duration(30) * time.Minute -> 30 minutes
	// time.Duration(30) * time.Minute - time.Nanosecond -> 30 minutes
	// time.Duration(30) * time.Minute - time.Minute + time.Nanosecond -> 30 minutes
	// time.Duration(30) * time.Minute - time.Minute -> 29 minutes
	if remaining > time.Minute {
		minutes := int(math.Ceil(float64(remaining) / float64(time.Minute)))
		timeString = fmt.Sprintf(MsgMinutesRemaining, minutes)
	} else if remaining > time.Second {
		seconds := int(math.Ceil(float64(remaining) / float64(time.Second)))
		timeString = fmt.Sprintf(MsgSecondsRemaining, seconds)
	} else if remaining > time.Millisecond {
		milliseconds := int(math.Ceil(float64(remaining) / float64(time.Millisecond)))
		timeString = fmt.Sprintf(MsgMillisecondsRemaining, milliseconds)
	}

	// If it's less than a millisecond or already over, just go with the "no time
	// remaining" message.
	return timeString
}
