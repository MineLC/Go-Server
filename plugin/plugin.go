package plugin

type Plugin interface {
	Enable()
	Disable()

	IsReloadeable() bool
	Name() string
}
