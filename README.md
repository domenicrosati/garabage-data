# What is data and why is it garbage?

**PSA: There is nothing in here about fast fourier transforms**

Lesson Plan:
- The problem
- Brainstorm solutions
- Solve with vectorization

## So whats the problem here?

Data is dirty for different reasons, there is a true signal or pattern here.
We should find a transformation that takes dirty data and spits out the true signal.


## How can we solve it?

So many different ways. But they usually boil down to two things
Come up with an explanation ourselves or come up with a model that will be able to estimate its own explanation (though we will have to provide the parameters of that model). Once we have an approach - we implement it as a transformation.
