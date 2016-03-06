## Nomenclature 

* `Accessible` interface hides the actual object properties to be changed by the tweened transition
* `TransitionCompletionFunctio` maps \[0.0,1.0\] to the Accessible's Setter based on the easing function and start and target values.
* `Transition` is basically a factory for `TransitionCompletionFunction` which should pick the starting values 
at the right moment when participating in a sequence 
* `Change` transforms duration \[0, inf.\] to the proper `Transition` based on their number and `RepetitionFunction`

## TODO

* ci