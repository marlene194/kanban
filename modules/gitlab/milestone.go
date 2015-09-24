package gitlab

import (
	"net/http"
	"net/url"
)

type Milestone struct {
	Id    int64  `json:"id"`
	State string `json:"state,omitempty"`
	Title string `json:"title,omitempty"`
}

// Get list milestones for gitlab projects
func (g *GitlabContext) ListMilestones(project_id string, o *ListOptions) ([]*Milestone, error) {
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "milestones"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Milestone
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}
