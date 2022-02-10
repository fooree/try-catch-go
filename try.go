package try

type Routine func()

func Try(routine func()) Routine {
	return routine
}

func (routine Routine) Catch(catch func(v interface{})) {
	if routine == nil {
		return
	}
	defer func() {
		if v := recover(); v == nil || catch == nil {
			return
		} else {
			catch(v)
		}
	}()
	routine()
}
