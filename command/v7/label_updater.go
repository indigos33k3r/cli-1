package v7

import (
	"errors"
	"strings"

	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/translatableerror"
	"code.cloudfoundry.org/cli/types"
)

//go:generate counterfeiter . SetLabelActor

type SetLabelActor interface {
	UpdateApplicationLabelsByApplicationName(string, string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateBuildpackLabelsByBuildpackNameAndStack(string, string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateDomainLabelsByDomainName(string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateOrganizationLabelsByOrganizationName(string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateRouteLabels(string, string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateSpaceLabelsBySpaceName(string, string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateStackLabelsByStackName(string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateServiceBrokerLabelsByServiceBrokerName(string, map[string]types.NullString) (v7action.Warnings, error)
	UpdateServiceOfferingLabels(serviceOfferingName string, serviceBrokerName string, labels map[string]types.NullString) (v7action.Warnings, error)
	UpdateServicePlanLabels(servicePlanName string, serviceOfferingName string, serviceBrokerName string, labels map[string]types.NullString) (v7action.Warnings, error)
}

type ActionType string

const (
	Unset ActionType = "Removing"
	Set   ActionType = "Setting"
)

type TargetResource struct {
	ResourceType    string
	ResourceName    string
	BuildpackStack  string
	ServiceBroker   string
	ServiceOffering string
}

type LabelUpdater struct {
	targetResource TargetResource
	labels         map[string]types.NullString

	UI          command.UI
	Config      command.Config
	SharedActor command.SharedActor
	Actor       SetLabelActor

	Username string
	Action   ActionType
}

func (cmd *LabelUpdater) Execute(targetResource TargetResource, labels map[string]types.NullString) error {
	cmd.targetResource = targetResource
	cmd.labels = labels
	cmd.targetResource.ResourceType = strings.ToLower(cmd.targetResource.ResourceType)

	var err error
	cmd.Username, err = cmd.Config.CurrentUserName()
	if err != nil {
		return err
	}

	if err := cmd.validateFlags(); err != nil {
		return err
	}

	if err := cmd.checkTarget(); err != nil {
		return err
	}

	var warnings v7action.Warnings
	switch ResourceType(cmd.targetResource.ResourceType) {
	case App:
		cmd.displayMessageWithOrgAndSpace()
		warnings, err = cmd.Actor.UpdateApplicationLabelsByApplicationName(cmd.targetResource.ResourceName, cmd.Config.TargetedSpace().GUID, cmd.labels)
	case Buildpack:
		cmd.displayMessageWithStack()
		warnings, err = cmd.Actor.UpdateBuildpackLabelsByBuildpackNameAndStack(cmd.targetResource.ResourceName, cmd.targetResource.BuildpackStack, cmd.labels)
	case Domain:
		cmd.displayMessageDefault()
		warnings, err = cmd.Actor.UpdateDomainLabelsByDomainName(cmd.targetResource.ResourceName, cmd.labels)
	case Org:
		cmd.displayMessageDefault()
		warnings, err = cmd.Actor.UpdateOrganizationLabelsByOrganizationName(cmd.targetResource.ResourceName, cmd.labels)
	case Route:
		cmd.displayMessageWithOrgAndSpace()
		warnings, err = cmd.Actor.UpdateRouteLabels(cmd.targetResource.ResourceName, cmd.Config.TargetedSpace().GUID, cmd.labels)
	case ServiceBroker:
		cmd.displayMessageDefault()
		warnings, err = cmd.Actor.UpdateServiceBrokerLabelsByServiceBrokerName(cmd.targetResource.ResourceName, cmd.labels)
	case Space:
		cmd.displayMessageWithOrg()
		warnings, err = cmd.Actor.UpdateSpaceLabelsBySpaceName(cmd.targetResource.ResourceName, cmd.Config.TargetedOrganization().GUID, cmd.labels)
	case Stack:
		cmd.displayMessageDefault()
		warnings, err = cmd.Actor.UpdateStackLabelsByStackName(cmd.targetResource.ResourceName, cmd.labels)
	case ServiceOffering:
		cmd.displayMessageForServiceCommands()
		warnings, err = cmd.Actor.UpdateServiceOfferingLabels(cmd.targetResource.ResourceName, cmd.targetResource.ServiceBroker, cmd.labels)
	case ServicePlan:
		cmd.displayMessageForServiceCommands()
		warnings, err = cmd.Actor.UpdateServicePlanLabels(cmd.targetResource.ResourceName, cmd.targetResource.ServiceOffering, cmd.targetResource.ServiceBroker, cmd.labels)
	}

	cmd.UI.DisplayWarnings(warnings)
	if err != nil {
		return err
	}

	cmd.UI.DisplayOK()
	return nil
}

func (cmd *LabelUpdater) checkTarget() error {
	switch ResourceType(cmd.targetResource.ResourceType) {
	case App, Route:
		return cmd.SharedActor.CheckTarget(true, true)
	case Space:
		return cmd.SharedActor.CheckTarget(true, false)
	default:
		return cmd.SharedActor.CheckTarget(false, false)
	}
}

func (cmd *LabelUpdater) validateFlags() error {
	resourceType := ResourceType(cmd.targetResource.ResourceType)
	switch resourceType {
	case App, Buildpack, Domain, Org, Route, ServiceBroker, Space, Stack, ServiceOffering, ServicePlan:
	default:
		return errors.New(cmd.UI.TranslateText("Unsupported resource type of '{{.ResourceType}}'", map[string]interface{}{"ResourceType": cmd.targetResource.ResourceType}))
	}

	if cmd.targetResource.BuildpackStack != "" && resourceType != Buildpack {
		return translatableerror.ArgumentCombinationError{
			Args: []string{
				cmd.targetResource.ResourceType, "--stack, -s",
			},
		}
	}

	if cmd.targetResource.ServiceBroker != "" && !(resourceType == ServiceOffering || resourceType == ServicePlan) {
		return translatableerror.ArgumentCombinationError{
			Args: []string{
				cmd.targetResource.ResourceType, "--broker, -b",
			},
		}
	}

	if cmd.targetResource.ServiceOffering != "" && resourceType != ServicePlan {
		return translatableerror.ArgumentCombinationError{
			Args: []string{
				cmd.targetResource.ResourceType, "--offering, -o",
			},
		}
	}

	return nil
}

func (cmd *LabelUpdater) displayMessageDefault() {
	cmd.UI.DisplayTextWithFlavor("{{.Action}} label(s) for {{.ResourceType}} {{.ResourceName}} as {{.User}}...", map[string]interface{}{
		"Action":       cmd.Action,
		"ResourceType": cmd.targetResource.ResourceType,
		"ResourceName": cmd.targetResource.ResourceName,
		"User":         cmd.Username,
	})
}

func (cmd *LabelUpdater) displayMessageWithOrgAndSpace() {
	cmd.UI.DisplayTextWithFlavor("{{.Action}} label(s) for {{.ResourceType}} {{.ResourceName}} in org {{.OrgName}} / space {{.SpaceName}} as {{.User}}...", map[string]interface{}{
		"Action":       cmd.Action,
		"ResourceType": cmd.targetResource.ResourceType,
		"ResourceName": cmd.targetResource.ResourceName,
		"OrgName":      cmd.Config.TargetedOrganization().Name,
		"SpaceName":    cmd.Config.TargetedSpace().Name,
		"User":         cmd.Username,
	})
}

func (cmd *LabelUpdater) displayMessageWithStack() {
	var template string
	if cmd.targetResource.BuildpackStack == "" {
		template = "{{.Action}} label(s) for {{.ResourceType}} {{.ResourceName}} as {{.User}}..."
	} else {
		template = "{{.Action}} label(s) for {{.ResourceType}} {{.ResourceName}} with stack {{.StackName}} as {{.User}}..."
	}

	cmd.UI.DisplayTextWithFlavor(template, map[string]interface{}{
		"Action":       cmd.Action,
		"ResourceType": cmd.targetResource.ResourceType,
		"ResourceName": cmd.targetResource.ResourceName,
		"StackName":    cmd.targetResource.BuildpackStack,
		"User":         cmd.Username,
	})
}

func (cmd *LabelUpdater) displayMessageForServiceCommands() {
	var template string
	template = "{{.Action}} label(s) for {{.ResourceType}} {{.ResourceName}}"

	if cmd.targetResource.ServiceOffering != "" || cmd.targetResource.ServiceBroker != "" {
		template += " from"
	}
	if cmd.targetResource.ServiceOffering != "" {
		template += " service offering {{.ServiceOffering}}"
		if cmd.targetResource.ServiceBroker != "" {
			template += " /"
		}
	}

	if cmd.targetResource.ServiceBroker != "" {
		template += " service broker {{.ServiceBroker}}"
	}

	template += " as {{.User}}..."
	cmd.UI.DisplayTextWithFlavor(template, map[string]interface{}{
		"Action":          cmd.Action,
		"ResourceName":    cmd.targetResource.ResourceName,
		"ResourceType":    cmd.targetResource.ResourceType,
		"ServiceBroker":   cmd.targetResource.ServiceBroker,
		"ServiceOffering": cmd.targetResource.ServiceOffering,
		"User":            cmd.Username,
	})
}

func (cmd *LabelUpdater) displayMessageWithOrg() {
	cmd.UI.DisplayTextWithFlavor("{{.Action}} label(s) for {{.ResourceType}} {{.ResourceName}} in org {{.OrgName}} as {{.User}}...", map[string]interface{}{
		"Action":       cmd.Action,
		"ResourceType": cmd.targetResource.ResourceType,
		"ResourceName": cmd.targetResource.ResourceName,
		"OrgName":      cmd.Config.TargetedOrganization().Name,
		"User":         cmd.Username,
	})
}
