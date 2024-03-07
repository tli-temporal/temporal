// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package api

import (
	"fmt"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/serviceerror"

	enumsspb "go.temporal.io/server/api/enums/v1"
	"go.temporal.io/server/common/payloads"
	"go.temporal.io/server/service/history/consts"
	"go.temporal.io/server/service/history/workflow"
)

// ResolveDuplicateWorkflowID determines how to resolve a workflow ID duplication upon workflow start according
// to the WorkflowIDReusePolicy.
//
// An action (ie "mitigate and allow"), an error (ie "deny") or neither (ie "allow") is returned.
func ResolveDuplicateWorkflowID(
	workflowID,
	newRunID,
	currentRunID string,
	currentState enumsspb.WorkflowExecutionState,
	currentStatus enumspb.WorkflowExecutionStatus,
	currentStartRequestID string,
	wfIDReusePolicy enumspb.WorkflowIdReusePolicy,
) (UpdateWorkflowActionFunc, error) {

	switch currentState {
	// *running* workflow
	case enumsspb.WORKFLOW_EXECUTION_STATE_CREATED, enumsspb.WORKFLOW_EXECUTION_STATE_RUNNING:
		switch wfIDReusePolicy {
		case enumspb.WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING:
			return terminateWorkflowAction(newRunID)
		default:
			msg := "Workflow execution is already running. WorkflowId: %v, RunId: %v."
			return nil, generateWorkflowAlreadyStartedError(msg, currentStartRequestID, workflowID, currentRunID)
		}

	// *completed* workflow
	case enumsspb.WORKFLOW_EXECUTION_STATE_COMPLETED:
		switch wfIDReusePolicy {
		case enumspb.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE, enumspb.WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING:
			// no action or error
		case enumspb.WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY:
			if _, ok := consts.FailedWorkflowStatuses[currentStatus]; !ok {
				msg := "Workflow execution already finished successfully. WorkflowId: %v, RunId: %v. Workflow Id reuse policy: allow duplicate workflow Id if last run failed."
				return nil, generateWorkflowAlreadyStartedError(msg, currentStartRequestID, workflowID, currentRunID)
			}
		case enumspb.WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE:
			msg := "Workflow execution already finished. WorkflowId: %v, RunId: %v. Workflow Id reuse policy: reject duplicate workflow Id."
			return nil, generateWorkflowAlreadyStartedError(msg, currentStartRequestID, workflowID, currentRunID)
		default:
			return nil, serviceerror.NewInternal(fmt.Sprintf("Failed to process start workflow id reuse policy: %v.", wfIDReusePolicy))
		}

	default:
		// persistence.WorkflowStateZombie or unknown type
		return nil, serviceerror.NewInternal(fmt.Sprintf("Failed to process workflow, workflow has invalid state: %v.", currentState))
	}

	return nil, nil
}

func terminateWorkflowAction(
	newRunID string,
) (UpdateWorkflowActionFunc, error) {
	return func(workflowLease WorkflowLease) (*UpdateWorkflowAction, error) {
		mutableState := workflowLease.GetMutableState()
		if !mutableState.IsWorkflowExecutionRunning() {
			return nil, consts.ErrWorkflowCompleted
		}

		return UpdateWorkflowWithoutWorkflowTask, workflow.TerminateWorkflow(
			mutableState,
			"TerminateIfRunning WorkflowIdReusePolicy",
			payloads.EncodeString(
				fmt.Sprintf("terminated by new runID: %s", newRunID),
			),
			consts.IdentityHistoryService,
			false,
		)
	}, nil
}

func generateWorkflowAlreadyStartedError(
	errMsg string,
	createRequestID string,
	workflowID string,
	runID string,
) error {
	return serviceerror.NewWorkflowExecutionAlreadyStarted(
		fmt.Sprintf(errMsg, workflowID, runID),
		createRequestID,
		runID,
	)
}
