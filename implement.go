package funcmethod

type work func() (v interface{}, err error)

func (w work) Do() (v interface{}, err error) {
	return w()
}

func dosomething() (v interface{}, err error) {
	return "cool", nil
}

func transform(s string) Worker {
	return work(func() (v interface{}, err error) {
		return s, nil
	})
}
