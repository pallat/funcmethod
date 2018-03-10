package funcmethod

type Worker interface {
	Do() (v interface{}, err error)
}

func Handler(w Worker) (v interface{}, err error) {
	r, err := w.Do()
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": r,
	}, nil
}
