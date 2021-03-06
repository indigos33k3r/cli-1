package v6_test

import (
	"code.cloudfoundry.org/cli/actor/actionerror"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/flag"
	. "code.cloudfoundry.org/cli/command/v6"
	"code.cloudfoundry.org/cli/command/v6/v6fakes"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("rename buildpack command", func() {
	var (
		cmd             RenameBuildpackCommand
		fakeActor       *v6fakes.FakeRenameBuildpackActor
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		testUI          *ui.UI
		executeErr      error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v6fakes.FakeRenameBuildpackActor)

		cmd = RenameBuildpackCommand{
			UI:          testUI,
			Actor:       fakeActor,
			Config:      fakeConfig,
			SharedActor: fakeSharedActor,
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	When("checking the target fails", func() {
		var binaryName string

		BeforeEach(func() {
			binaryName = "faceman"
			fakeConfig.BinaryNameReturns(binaryName)
			fakeSharedActor.CheckTargetReturns(actionerror.NotLoggedInError{BinaryName: binaryName})
		})

		It("returns an error", func() {
			Expect(executeErr).To(MatchError(actionerror.NotLoggedInError{BinaryName: "faceman"}))
			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			checkTargetedOrg, checkTargetedSpace := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(checkTargetedOrg).To(BeFalse())
			Expect(checkTargetedSpace).To(BeFalse())
		})
	})

	When("checking the target succeeds", func() {
		var (
			oldName   string
			newName   string
			stackName string
			userName  string
		)

		BeforeEach(func() {
			oldName = "some-old-name"
			newName = "some-new-name"
			stackName = "some-stack"
			userName = "some-user"

			cmd.RequiredArgs = flag.RenameBuildpackArgs{
				OldBuildpackName: oldName,
				NewBuildpackName: newName,
			}

			fakeConfig.CurrentUserReturns(configv3.User{Name: userName}, nil)
		})

		When("the user does not specify a stack", func() {
			When("no error is returned", func() {
				BeforeEach(func() {
					fakeActor.RenameBuildpackReturns(v2action.Warnings{"warning1", "warning2"}, nil)
				})

				It("successfully renames the buildpack and displays any warnings", func() {
					Expect(executeErr).ToNot(HaveOccurred())

					Expect(fakeActor.RenameBuildpackCallCount()).To(Equal(1))

					oldBuildpackName, newBuildpackName, stack := fakeActor.RenameBuildpackArgsForCall(0)
					Expect(oldBuildpackName).To(Equal(oldName))
					Expect(newBuildpackName).To(Equal(newName))
					Expect(stack).To(BeEmpty())

					Expect(testUI.Out).To(Say("Renaming buildpack %s to %s as %s...", oldName, newName, userName))
					Expect(testUI.Err).To(Say("warning1"))
					Expect(testUI.Err).To(Say("warning2"))
				})
			})

			When("an error is returned", func() {
				BeforeEach(func() {
					fakeActor.RenameBuildpackReturns(
						v2action.Warnings{"warning1", "warning2"},
						actionerror.BuildpackNameTakenError{Name: newName})
				})

				It("returns an error and prints warnings", func() {
					Expect(executeErr).To(MatchError(actionerror.BuildpackNameTakenError{Name: newName}))
					Expect(testUI.Err).To(Say("warning1"))
					Expect(testUI.Err).To(Say("warning2"))
				})
			})
		})

		When("the user specifies a stack", func() {
			BeforeEach(func() {
				cmd.Stack = stackName
			})

			When("the actor succeeds", func() {
				BeforeEach(func() {
					fakeActor.RenameBuildpackReturns(v2action.Warnings{"warning1", "warning2"}, nil)
				})

				It("successfully renames the buildpack and displays any warnings", func() {
					Expect(executeErr).ToNot(HaveOccurred())

					Expect(fakeActor.RenameBuildpackCallCount()).To(Equal(1))

					oldBuildpackName, newBuildpackName, stack := fakeActor.RenameBuildpackArgsForCall(0)
					Expect(oldBuildpackName).To(Equal(oldName))
					Expect(newBuildpackName).To(Equal(newName))
					Expect(stack).To(Equal(stackName))

					Expect(testUI.Out).To(Say("Renaming buildpack %s to %s with stack %s as %s...",
						oldName,
						newName,
						stackName,
						userName,
					))
					Expect(testUI.Err).To(Say("warning1"))
					Expect(testUI.Err).To(Say("warning2"))
				})
			})

			When("the actor returns an error", func() {
				When("BuildpackNameTakenError", func() {
					BeforeEach(func() {
						fakeActor.RenameBuildpackReturns(
							v2action.Warnings{"warning1", "warning2"},
							actionerror.BuildpackNameTakenError{Name: newName})
					})

					It("returns an error and prints warnings", func() {
						Expect(executeErr).To(MatchError(actionerror.BuildpackNameTakenError{Name: newName}))
						Expect(testUI.Err).To(Say("warning1"))
						Expect(testUI.Err).To(Say("warning2"))
					})
				})

				When("BuildpackInvalidError: name has invalid characters", func() {
					BeforeEach(func() {
						fakeActor.RenameBuildpackReturns(
							v2action.Warnings{"warning1", "warning2"},
							actionerror.BuildpackInvalidError{Message: "Buildpack is invalid: name can only contain alphanumeric characters"})
					})

					It("returns an error and prints warnings", func() {
						Expect(executeErr).To(MatchError(actionerror.BuildpackInvalidError{Message: "Buildpack is invalid: name can only contain alphanumeric characters"}))
						Expect(testUI.Err).To(Say("warning1"))
						Expect(testUI.Err).To(Say("warning2"))
					})
				})

				When("BuildpackInvalidError: stack non-unique", func() {
					BeforeEach(func() {
						fakeActor.RenameBuildpackReturns(
							v2action.Warnings{"warning1", "warning2"},
							actionerror.BuildpackInvalidError{Message: "Buildpack is invalid: stack unique"})
					})

					It("returns an error and prints warnings", func() {
						Expect(executeErr).To(MatchError("Buildpack some-new-name already exists without a stack"))
						Expect(testUI.Err).To(Say("warning1"))
						Expect(testUI.Err).To(Say("warning2"))
					})
				})
			})
		})
	})
})
