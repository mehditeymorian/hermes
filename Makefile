

emq_up:
	docker run -d --name emqx -p 1883:1883 -p 8081:8081 -p 8083:8083 -p 8883:8883 -p 8084:8084 -p 18083:18083 emqx/emqx:v4.0.0

emq_down:
	docker kill emqx
	
up:
	CGO_ENABLED=0 go build
	docker-compose -f docker-compose.yaml up --build

down:
	docker-compose -f docker-compose.yaml down
	docker-clean