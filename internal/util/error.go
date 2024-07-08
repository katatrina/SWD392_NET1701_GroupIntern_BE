package util

type MapErrors map[string]string

func (m MapErrors) Add(key, value string) {
	m[key] = value
}
