package factsphere

import (
	"testing"

	"github.com/jakevoytko/crbot/testutil"
)

func TestKarma(t *testing.T) {
	runner := testutil.NewRunner(t)

	// Correct call format
	runner.SendMessageIgnoringResponse(testutil.MainChannelID, "?factsphere")

}
