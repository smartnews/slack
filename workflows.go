package slack

import (
	"context"
	"encoding/json"
)

type WorkflowStep struct {
	WorkflowStepEditId string                       `json:"workflow_step_edit_id"`
	WorkflowId         string                       `json:"workflow_id"`
	StepId             string                       `json:"step_id"`
	Inputs             map[string]WorkflowStepInput `json:"inputs"`
	Output             []WorkflowStepOutput         `json:"outputs"`
}

type WorkflowStepInput struct {
	Value                   string                 `json:"value"`
	SkipVariableReplacement bool                   `json:"skip_variable_replacement"`
	Variables               map[string]interface{} `json:"variables"`
}

type WorkflowStepOutput struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Label string `json:"label"`
}

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
