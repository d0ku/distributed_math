.PHONY: expr adder subtracter multiplier divider
HOST=""
OWNER=d0ku
TAG=":0.1"

info:
	exit

all: expr adder subtracter multiplier divider

expr:
	docker build -t ${HOST}${OWNER}/distributed_math-expr${TAG} .  -f expr.Dockerfile

adder:
	docker build -t ${HOST}${OWNER}/distributed_math-adder${TAG} .  -f adder.Dockerfile

subtracter:
	docker build -t ${HOST}${OWNER}/distributed_math-subtracter${TAG} .  -f subtracter.Dockerfile

multiplier:
	docker build -t ${HOST}${OWNER}/distributed_math-multiplier${TAG} .  -f multiplier.Dockerfile

divider:
	docker build -t ${HOST}${OWNER}/distributed_math-divider${TAG} .  -f divider.Dockerfile

push:
	docker push ${HOST}${OWNER}/distributed_math-expr${TAG}
	docker push ${HOST}${OWNER}/distributed_math-adder${TAG}
	docker push ${HOST}${OWNER}/distributed_math-subtracter${TAG}
	docker push ${HOST}${OWNER}/distributed_math-multiplier${TAG}
	docker push ${HOST}${OWNER}/distributed_math-divider${TAG}
