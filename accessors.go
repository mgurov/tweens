package tweens

type Accessible interface {
	//TODO: at the moment we store the start value assuming it won't be mutated afterwards -> check with the best practices
	// and either document or store a copy
	Get() (currentState []float64)
	//the length of the newState will be equal to the one returned by Get at the beginning of the transition
	Set(newState []float64)
}

// the terminal goal of the change, e.g. move from A to B
// maps completion [0,1] -> [A,B]
type TransitionCompletionFunction func(completion float64)

func NoopCompletion(completion float64) {}

// subject.Get()/Set() are expected to return/accept the array of the len(target) (at least).
func Accessor(subject Accessible, target ...float64) TransitionCompletionFunction {
	return FunctionalAccessor(
		func() (currentState []float64) { return subject.Get() },
		func(newState []float64) { subject.Set(newState) },
		target...,
	)
}

// A helper that remembers the initial state of a transformation
// as an array of floating points and transforms the progress of [0, 1] into
// a new Array
func FunctionalAccessor(get func() (currentState []float64), set func(newState []float64), target ...float64) TransitionCompletionFunction {
	start := get()

	delta := make([]float64, len(target))

	for i, targetValue := range target {
		delta[i] = targetValue - start[i]
	}

	return func(progress float64) {
		newState := make([]float64, len(target))

		for i, deltaValue := range delta {
			newState[i] = start[i] + progress*deltaValue
		}

		set(newState)
	}
}

type Transition interface {
	// Transition might depends on the initial state of the subject
	// which in turn might depend on the time, especially in sequences
	// therefore this method is called only when the transition begins
	Start() TransitionCompletionFunction
}

func LazyAccessor(subject Accessible, target ...float64) Transition {
	return lazyAccessor{subject, target}
}

type lazyAccessor struct {
	subject Accessible
	target  []float64
}

func (a lazyAccessor) Start() TransitionCompletionFunction {
	return Accessor(a.subject, a.target...)
}
