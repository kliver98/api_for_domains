# api_for_domains

## Steps to run server in windows:

1. must have installed CockroachDB and golang
2. run on cmd console:
  
 > go get -d github.com/buaazp/fasthttprouter <br>
 > go get -d github.com/lib/pq
 
3. initialize and create cockroach database:
  <br><br>In the same console opened or another one, go to the directory where cockroach is. There execute:<br>
 > cockroach start-single-node --insecure --listen-addr=localhost:26257 --http-addr=localhost:5001 <br>In another cmd console<br>
 > cockroach sql --insecure
 > create database reto_prueba;
 > set database = reto_prueba;

4. creating tables needed into cockroach:
  <br><br>In the last cmd console opened<br>
 > create table history (
			domain STRING(200) NOT NULL,
			searched_at TIMESTAMPTZ NOT NULL	
); <br><br>
 > create table domain (
			name STRING(200) PRIMARY KEY,
			ssl_grade STRING(5) NOT NULL,
			previous_ssl_grade STRING(5) NOT NULL,
			searched_at TIMESTAMPTZ NOT NULL	
);<br>

5. run server:
<br><br>The server port it's found in Main.go file or in the console when server it's runnig<br>

 > go run Main.go Handler.go

## Steps to install android apk:

In the folder ApiForDomains there's a file called App_V1_0.apk<br>That build apk has the ip of my localhost so must not work, you have to change it and re-build apk.
<br>To edit and re-build you need Android Studio to open the project and change the IP, it's on Constants on the folder util.<br>