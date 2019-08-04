.PHONY: expr adder subtracter multiplier divider
HOST=registry.d0ku.org
OWNER=d0ku

info:
	exit

expr:
	docker build -t ${HOST}/${OWNER}/distributed_math/expr .  -f expr.Dockerfile

adder:
	docker build -t ${HOST}/${OWNER}/distributed_math/adder .  -f adder.Dockerfile

subtracter:
	docker build -t ${HOST}/${OWNER}/distributed_math/subtracter .  -f subtracter.Dockerfile

multiplier:
	docker build -t ${HOST}/${OWNER}/distributed_math/multiplier .  -f multiplier.Dockerfile

divider:
	docker build -t ${HOST}/${OWNER}/distributed_math/divider .  -f divider.Dockerfile

push:
	docker push ${HOST}/${OWNER}/distributed_math/expr
	docker push ${HOST}/${OWNER}/distributed_math/adder
	docker push ${HOST}/${OWNER}/distributed_math/subtracter
	docker push ${HOST}/${OWNER}/distributed_math/multiplier
	docker push ${HOST}/${OWNER}/distributed_math/divider
