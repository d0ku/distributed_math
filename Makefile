.PHONY: expr adder
HOST=registry.d0ku.org
OWNER=d0ku

info:
	exit

expr:
	docker build -t ${HOST}/${OWNER}/distributed_math/expr .  -f expr.Dockerfile

adder:
	docker build -t ${HOST}/${OWNER}/distributed_math/adder .  -f adder.Dockerfile

push:
	docker push ${HOST}/${OWNER}/distributed_math/expr
