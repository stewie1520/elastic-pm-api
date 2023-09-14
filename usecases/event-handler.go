package usecases

import "github.com/stewie1520/elasticpmapi/core"

func AddHandlersToHook(app core.App) {
	handleAfterAccountCreated(app)
}

func handleAfterAccountCreated(app core.App) {
	app.OnAfterAccountCreated().Add(func(event *core.AccountCreatedEvent) error {
		command := NewCreateUserCommand(app)
		command.ID = event.ID
		command.TimeJoined = event.TimeJoined
		command.Email = event.Email
		command.ThirdParty = event.ThirdParty
		command.TenantIds = event.TenantIds

		return command.Execute()
	})
}
