## Distributed Math (or how to run microservices on k8s)

I wanted to learn something about stateless microservices that can be easily scaled up, so I made this.
It's my first project that has anything to do with splitting work into smaller chunks and the result is obviously terrible from the performance, quality of code etc. point of view.

### How does it work?

Master node (expr) listens for incoming JSON requests from users that want to calculate value of some expression. After it is received every arithmethic operation, such as addition, subtraction etc. is done on a different node (or at least container). All services share same base code that is defined in (you don't say?...) base directory.

### Is it useful?

Probably not, I just wanted to play with some ideas. It does NOT work properly at the moment (operators precedence is not correct, actually no serious error handling e.g.) and I don't really plan on fixing that, because that's not the point of that repo.

### Can i use it?

Sure, if you really need to. Just set it up (k8s folder contains files that can be used to deploy on Kubernetes, maybe I will write docker-compose setup) and send JSON request of form
```
{"Content": "<yourexpression>"}
```
e.g.
```
{"Content": "2+2*3"}
```
to expr node. Tester program can be used for that (it's included in this repo).
