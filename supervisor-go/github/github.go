package github

import (
	"net/http"

	"github.com/go-playground/webhooks/v6/github"
	log "github.com/sirupsen/logrus"
	"github.com/yudai/pp"
)

func AddRoutes(secret string) error {
	hook, err := github.New(github.Options.Secret(secret))
	if err != nil {
		return err
	}

	http.HandleFunc("/webhooks", func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r,
			github.CheckRunEvent,
			github.CheckSuiteEvent,
			github.CommitCommentEvent,
			github.CreateEvent,
			github.DeleteEvent,
			github.DeploymentEvent,
			github.DeploymentStatusEvent,
			github.ForkEvent,
			github.GollumEvent,
			github.InstallationEvent,
			github.InstallationRepositoriesEvent,
			github.IntegrationInstallationEvent,
			github.IntegrationInstallationRepositoriesEvent,
			github.IssueCommentEvent,
			github.IssuesEvent,
			github.LabelEvent,
			github.MemberEvent,
			github.MembershipEvent,
			github.MilestoneEvent,
			github.MetaEvent,
			github.OrganizationEvent,
			github.OrgBlockEvent,
			github.PageBuildEvent,
			github.PingEvent,
			github.ProjectCardEvent,
			github.ProjectColumnEvent,
			github.ProjectEvent,
			github.PublicEvent,
			github.PullRequestEvent,
			github.PullRequestReviewEvent,
			github.PullRequestReviewCommentEvent,
			github.PushEvent,
			github.ReleaseEvent,
			github.RepositoryEvent,
			github.RepositoryVulnerabilityAlertEvent,
			github.SecurityAdvisoryEvent,
			github.StatusEvent,
			github.TeamEvent,
			github.TeamAddEvent,
			github.WatchEvent,
		)
		if err != nil {
			log.Error(err)
		}

		switch t := payload.(type) {
		case github.ReleasePayload:
			log.Infof("ReleasePayload: %v\n", pp.Sprint(t))

		case github.PullRequestPayload:
			log.Infof("PullRequestPayload: %v\n", pp.Sprint(t))

		case github.PingPayload:
			log.Infof("PingPayload: %v\n", pp.Sprint(t))

		case github.PushPayload:
			log.Infof("PushPayload: %v\n", pp.Sprint(t))

		default:
			log.Infof("Unknow payload: %v\n", payload)
		}
	})

	return nil
}
