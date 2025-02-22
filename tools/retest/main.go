package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v60/github"
)

func main() {
	ctx := context.Background()

	retestComment := "/retest"

	// Use installation transport with client.
	client := github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_TOKEN"))

	//TODO(janisz): handle pagination
	search, _, err := client.Search.Issues(ctx, `repo:stackrox/stackrox label:auto-retest state:open type:pr status:failure`, nil)
	handleError(err)
	log.Printf("Found %d PRs", search.GetTotal())

	for _, pr := range search.Issues {
		prNumber := pr.GetNumber()
		log.Printf("Processing #%d", prNumber)
		comments, _, err := client.Issues.ListComments(ctx, "stackrox", "stackrox", prNumber, nil)
		handleError(err)

		retested := 0
		for _, c := range comments {
			if c.GetBody() == retestComment {
				retested++
			}
		}
		log.Printf("#%d was retested %d times", prNumber, retested)
		if retested > 3 {
			continue
		}

		prDetails, _, err := client.PullRequests.Get(ctx, "stackrox", "stackrox", prNumber)
		handleError(err)

		if isGHACheckFailing(ctx, client, prDetails.GetHead().GetSHA()) {
			continue
		}

		var statuses []Status
		statusRequest, err := http.NewRequest("GET", prDetails.GetStatusesURL(), nil)
		handleError(err)
		_, err = client.Do(ctx, statusRequest, &statuses)
		handleError(err)

		retestComment := github.IssueComment{
			Body: &retestComment,
		}

		for _, status := range statuses {
			log.Printf("#%d %-40s\t%10s", prNumber, status.Context, status.State)
			if status.State != "failure" {
				continue
			}
			comment, _, err := client.Issues.CreateComment(ctx, "stackrox", "stackrox", prNumber, &retestComment)
			handleError(err)
			log.Printf("#%d commented: %s", prNumber, comment.GetURL())
			break
		}
	}
}

func isGHACheckFailing(ctx context.Context, client *github.Client, lastCommit string) bool {
	checks, _, err := client.Checks.ListCheckRunsForRef(ctx, "stackrox", "stackrox", lastCommit, &github.ListCheckRunsOptions{
		Status: ptr("completed"),
		Filter: ptr("latest"),
	})
	handleError(err)
	for _, check := range checks.CheckRuns {
		ghaUrlPrefix := "https://api.github.com/repos/stackrox/stackrox/check-runs/"
		if check.GetConclusion() == "failure" &&
			strings.HasPrefix(check.GetURL(), ghaUrlPrefix) {
			log.Printf("%s has failed: %s", check.GetName(), check.GetHTMLURL())
		}
		return true
	}
	return false
}

type Status struct {
	Context string `json:"context"`
	State   string `json:"state"`
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ptr[T any](in T) *T {
	return &in
}
