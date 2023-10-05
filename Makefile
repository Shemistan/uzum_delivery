test:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

lint:
	golangci-lint run

app_run:
	docker-compose -f docker-compose.yaml down -v
	docker-compose -f docker-compose.yaml up -d db
	sleep 15
	docker-compose -f docker-compose.yaml up -d init
	docker-compose -f docker-compose.yaml up -d migrate
	docker-compose -f docker-compose.yaml up -d uzum_delivery
#	docker-compose -f docker-compose.yaml up -d zookeeper
#	docker-compose -f docker-compose.yaml up -d kafka
#	docker exec -it uzum_delivery-kafka-1 kafka-topics --create --bootstrap-server localhost:29092 --topic my_topic --partitions 3 --replication-factor 1

app_run2:
	docker-compose -f docker-compose.yaml down -v
	docker-compose up -d
#Записать в Топик
#docker exec -it uzum_delivery-kafka-1 kafka-console-producer --broker-list localhost:29092 --topic my_topic