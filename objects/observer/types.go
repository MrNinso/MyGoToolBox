package observer

type OnGetHandle func(this *Observer, obj interface{})

type OnSetHandle func(this *Observer, objOld, objNew interface{})

type Observer struct {
	onGet []*OnGetHandle
	onSet []*OnSetHandle
	value interface{}
	async bool
	wait  bool
}

func (o Observer) Get() interface{} {
	if o.wait {
		o.runGetHandles()
	} else {
		go o.runGetHandles()
	}

	return o.value
}

func (o Observer) Set(v interface{}) {
	if o.wait {
		o.runSetHandles(v)
	} else {
		go o.runSetHandles(v)
	}

	o.value = v
}

func (o Observer) SubscribeGet(h *OnGetHandle) {
	o.onGet = append(o.onGet, h)
}

func (o Observer) SubscribeSet(h *OnSetHandle) {
	o.onSet = append(o.onSet, h)
}

func (o Observer) UnsubscribeGet(h *OnGetHandle) {
	for i := 0; i < len(o.onGet); i++ {
		if o.onGet[i] == h {
			o.onGet[len(o.onGet)-1], o.onGet[i] = o.onGet[i], o.onGet[len(o.onGet)-1]
			o.onGet = o.onGet[:len(o.onGet)-1]
			return
		}
	}
}

func (o Observer) UnsubscribeSet(h *OnSetHandle) {
	for i := 0; i < len(o.onGet); i++ {
		if o.onSet[i] == h {
			o.onSet[len(o.onSet)-1], o.onSet[i] = o.onSet[i], o.onSet[len(o.onSet)-1]
			o.onSet = o.onSet[:len(o.onSet)-1]
			return
		}
	}
}

func (o *Observer) runSetHandles(objNew interface{}) {
	for i := 0; i < len(o.onSet); i++ {
		if o.async {
			go (*o.onSet[i])(o, o.value, objNew)
		} else {
			(*o.onSet[i])(o, o.value, objNew)
		}
	}
}

func (o *Observer) runGetHandles() {
	for i := 0; i < len(o.onSet); i++ {
		if o.async {
			go (*o.onGet[i])(o, o.value)
		} else {
			(*o.onGet[i])(o, o.value)
		}
	}
}
