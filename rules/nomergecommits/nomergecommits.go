package nomergecommits

import (
	"strings"

	"github.com/vbatts/git-validation/git"
	"github.com/vbatts/git-validation/validate"
)

var (
	// NoMergeCommitsRule is the rule being registered
	NoMergeCommitsRule = validate.Rule{
		Name:        "no-merge-commits",
		Description: "commits are not merge commits",
		Run:         ValidateNoMergeCommits,
	}
)

func init() {
	validate.RegisterRule(NoMergeCommitsRule)
}

// ValidateNoMergeCommits checks that the commit is not a merge commit
func ValidateNoMergeCommits(c git.CommitEntry) (vr validate.Result) {
	vr.CommitEntry = c
	if len(strings.Split(c["parent"], " ")) > 1 {
		vr.Pass = false
		vr.Msg = "merge commits are not allowed"
		return vr
	}

	vr.Pass = true
	vr.Msg = "commit is not a merge commit"

	return
}
