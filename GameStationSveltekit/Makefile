dev :
	npm run dev
build:
	npm run build
preview:
	npm run preview
test:
	npm run test

docker_test: docker-build_test 
	sudo docker compose up -d 
	# sudo docker compsoe exec -T http go test ./...
	sudo docker compose down 
docker_build:
	sudo docker build -t shahriarraka/game_station_frontend:latest .
	# docker build -t shahriarraka/game_station_frontend:latest .
docker_push:
	sudo docker push shahriarraka/game_station_frontend:latest 
docker_build_test:
	docker build . -t service_test --target=test

docker_run:
	sudo docker stop gamestation_sveltekit || true && sudo docker rm gamestation_sveltekit || true
	sudo docker run --publish 8080:8080 --name gamestation_sveltekit shahriarraka/game_station_frontend:latest  