package funcmethod

type work func() (v interface{}, err error)

func (w work) Do() (v interface{}, err error) {
	return w()
}

func transform(s string) Worker {
	return work(func() (v interface{}, err error) {
		return s, nil
	})
}
