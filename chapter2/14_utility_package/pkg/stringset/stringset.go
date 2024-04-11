package stringset

type Set map[string]struct{}

func New(...string) Set {
	// dosomething
	return Set{}
}

func (s Set) Sort() []string {
	// dosomething
	return []string{"hoge"}
}
