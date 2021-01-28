package slack

import (
	"context"
	"encoding/json"
)

type UpdateWorkflowStepParameters struct {
	WorkflowStepEditId string `json:"workflow_step_edit_id"`
}

// OpenView opens a view for a user.
func (api *Client) UpdateWorkflowStep(params *UpdateWorkflowStepParameters) (*SlackResponse, error) {
	return api.UpdateWorkflowStepContext(context.Background(), params)
}

// OpenViewContext opens a view for a user with a custom context.
func (api *Client) UpdateWorkflowStepContext(
	ctx context.Context,
	params *UpdateWorkflowStepParameters,
) (*SlackResponse, error) {
	encoded, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	endpoint := api.endpoint + "workflows.updateStep"
	resp := &SlackResponse{}
	err = postJSON(ctx, api.httpclient, endpoint, api.token, encoded, resp, api)
	if err != nil {
		return nil, err
	}
	return resp, resp.Err()
}
