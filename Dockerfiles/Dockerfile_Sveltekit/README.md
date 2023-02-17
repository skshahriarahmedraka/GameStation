





## Run Docker Container

```bash
sudo docker build -t shahriarraka/game_station_frontend:latest  .
sudo docker push shahriarraka/game_station_frontend:latest 
sudo docker run -p 4173:4173 --name mytest1 shahriarraka/game_station_frontend:latest
```