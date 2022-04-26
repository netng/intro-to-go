package maps

type Dictionary map[string]string

func (d Dictionary) search(world string) string {
	return d[world]
}
