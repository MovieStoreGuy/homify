package device

type (
	// Group extends a list make a named list
	Group interface {
		List
		// GetName returns the anme of the group
		GetName() string
	}

	group struct {
		list
		Name string
	}
)

var (
	_ Group = (*group)(nil)
)

func NewGroup(name string, devices ...Device) Group {
	g := &group{
		Name: name,
	}
	g.list.Append(devices...)

	return g
}

func (g *group) GetName() string {
	return g.Name
}
