## Run Docker Container 

```bash
sudo docker build -t shahriarraka/game_station_golangapi:latest  .
sudo docker push shahriarraka/game_station_golangapi:latest 
sudo docker run -p 8080:8080 --name mytest1 shahriarraka/game_station_golangapi:latest

```