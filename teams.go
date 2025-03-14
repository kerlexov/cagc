package cagc

import (
	"context"
	"fmt"
	"net/http"
)

// ListTeams lists all teams
func (c *Client) ListTeams(ctx context.Context) ([]Team, error) {
	var teams []Team
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/teams", nil, &teams)
	return teams, err
}

// GetTeam gets a team by ID
func (c *Client) GetTeam(ctx context.Context, id string) (*Team, error) {
	path := fmt.Sprintf("/api/v1/teams/%s", id)
	var team Team
	err := c.doRequest(ctx, http.MethodGet, path, nil, &team)
	return &team, err
}

// GetTeamMembers gets members by team ID
func (c *Client) GetTeamMembers(ctx context.Context, id string) ([]TeamMember, error) {
	path := fmt.Sprintf("/api/v1/teams/%s/members", id)
	var members []TeamMember
	err := c.doRequest(ctx, http.MethodGet, path, nil, &members)
	return members, err
}

// GetCurrentTeam gets the currently authenticated team
func (c *Client) GetCurrentTeam(ctx context.Context) (*Team, error) {
	var team Team
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/teams/current", nil, &team)
	return &team, err
}

// GetCurrentTeamMembers gets the currently authenticated team members
func (c *Client) GetCurrentTeamMembers(ctx context.Context) ([]TeamMember, error) {
	var members []TeamMember
	err := c.doRequest(ctx, http.MethodGet, "/api/v1/teams/current/members", nil, &members)
	return members, err
}
