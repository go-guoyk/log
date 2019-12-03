package labels

// Labels is a type alias of map[string]interface{}
type Labels map[string]interface{}

// Merge merge another labels into self
func (l Labels) Merge(a Labels) Labels {
	o := l.Clone()
	if len(a) != 0 {
		if o == nil {
			o = make(Labels, len(a))
		}
		for k, v := range a {
			o[k] = v
		}
	}
	return o
}

// Clone a shallow clone of label
func (l Labels) Clone() Labels {
	if len(l) == 0 {
		return nil
	}
	c := make(Labels, len(l))
	for k, v := range l {
		c[k] = v
	}
	return c
}
