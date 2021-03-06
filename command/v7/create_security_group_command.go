package v7

import (
	"encoding/json"

	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v7action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccerror"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/flag"
	"code.cloudfoundry.org/cli/command/v7/shared"
	"code.cloudfoundry.org/clock"
)

//go:generate counterfeiter . CreateSecurityGroupActor

type CreateSecurityGroupActor interface {
	CreateSecurityGroup(name, filePath string) (v7action.Warnings, error)
}

type CreateSecurityGroupCommand struct {
	RequiredArgs    flag.SecurityGroupArgs `positional-args:"yes"`
	usage           interface{}            `usage:"CF_NAME create-security-group SECURITY_GROUP PATH_TO_JSON_RULES_FILE\n\n   The provided path can be an absolute or relative path to a file. The file should have\n   a single array with JSON objects inside describing the rules. The JSON Base Object is\n   omitted and only the square brackets and associated child object are required in the file.\n\n   Valid json file example:\n   [\n     {\n       \"protocol\": \"tcp\",\n       \"destination\": \"10.0.11.0/24\",\n       \"ports\": \"80,443\",\n       \"description\": \"Allow http and https traffic from ZoneA\"\n     }\n   ]"`
	relatedCommands interface{}            `related_commands:"bind-running-security-group, bind-security-group, bind-staging-security-group, security-groups"`

	UI          command.UI
	Config      command.Config
	Actor       CreateSecurityGroupActor
	SharedActor command.SharedActor
}

func (cmd *CreateSecurityGroupCommand) Setup(config command.Config, ui command.UI) error {
	cmd.UI = ui
	cmd.Config = config
	sharedActor := sharedaction.NewActor(config)
	cmd.SharedActor = sharedActor
	cmd.usage = ``

	ccClient, uaaClient, err := shared.GetNewClientsAndConnectToCF(config, ui, "")
	if err != nil {
		return err
	}
	cmd.Actor = v7action.NewActor(ccClient, config, sharedActor, uaaClient, clock.NewClock())
	return nil
}

func (cmd CreateSecurityGroupCommand) Execute(args []string) error {
	err := cmd.SharedActor.CheckTarget(false, false)
	if err != nil {
		return err
	}

	user, err := cmd.Config.CurrentUser()
	if err != nil {
		return err
	}

	cmd.UI.DisplayTextWithFlavor("Creating security group {{.Name}} as {{.Username}}...", map[string]interface{}{
		"Name":     cmd.RequiredArgs.SecurityGroup,
		"Username": user.Name,
	})
	cmd.UI.DisplayNewline()

	warnings, err := cmd.Actor.CreateSecurityGroup(cmd.RequiredArgs.SecurityGroup, string(cmd.RequiredArgs.PathToJSONRules))
	cmd.UI.DisplayWarnings(warnings)
	if _, ok := err.(*json.SyntaxError); ok {
		return actionerror.SecurityGroupJsonSyntaxError{Path: string(cmd.RequiredArgs.PathToJSONRules)}
	}

	if _, ok := err.(ccerror.SecurityGroupAlreadyExists); ok {
		cmd.UI.DisplayWarning(err.Error())
	} else if err != nil {
		return err
	}

	cmd.UI.DisplayOK()
	return nil
}
