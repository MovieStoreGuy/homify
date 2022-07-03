package device

type (
	List interface {
		Append(devices ...Device)

		Range(fn func(d Device) (more bool))

		Devices() []Device
	}

	list struct {
		devices map[Device]struct{}
	}
)

func NewList(devices ...Device) List {
	dl := &list{
		devices: map[Device]struct{}{},
	}
	dl.Append(devices...)
	return dl
}

func (dl *list) Append(devices ...Device) {
	for _, d := range devices {
		dl.devices[d] = struct{}{}
	}
}

func (dl *list) Range(fn func(d Device) bool) {
	for d := range dl.devices {
		if !fn(d) {
			return
		}
	}
}

func (dl *list) Devices() []Device {
	devices := make([]Device, 0, len(dl.devices))
	dl.Range(func(d Device) bool {
		devices = append(devices, d)
		return true
	})
	return devices
}
