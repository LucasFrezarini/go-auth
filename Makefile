test:
	-docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
	docker-compose -f docker-compose.test.yml down --volumes

test-db-up:
	docker-compose -f docker-compose.test.yml up --build go-auth-db-test

test-db-down:
	docker-compose -f docker-compose.test.yml down --volumes

run:
	go run server/server.go