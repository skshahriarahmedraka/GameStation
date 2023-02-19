## Run Docker Container 

```bash
sudo docker build -t shahriarraka/game_station_golangapi:latest .
sudo docker push shahriarraka/game_station_golangapi:latest 
sudo docker run -p 8080:8080 --name mytest1 shahriarraka/game_station_golangapi:latest

```

## List Of Cookies & there use

1. `Auth1` httpOnly cookie for holding 
`Name:user.Name,Email: user.Email,UserID:user.UserID,Accounttype: user.Accounttype,`
2. `noAuth2` httpOnly cookie for holding
`Name:user.Name,Email: user.Email,UserID:user.UserID,Accounttype: user.Accounttype,`
