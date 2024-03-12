package interval

import "errors"

// interval [,]
type Interval struct {
	From uint
	To   uint
}

func (interval *Interval) AddConstraint(constraint string, t uint) error {
	if constraint == ">=" && t >= interval.From {
		interval.From = t
		return nil
	}
	if constraint == "<=" && t <= interval.To {
		interval.To = t
		return nil
	}
	if constraint != "<=" && constraint != ">=" {
		return errors.New("invalid constraint:" + constraint)
	}
	return nil
}

func (interval Interval) GetOptimal() int {
	if interval.From <= interval.To {
		return int(interval.From)
	}
	return -1
}
