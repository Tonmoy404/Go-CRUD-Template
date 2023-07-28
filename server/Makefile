start:
	go run main.go

mysql:
	go get -u github.com/go-sql-driver/mysql

runmysql: 
	docker run -p 3306:3306 --name mysql -e  MYSQL_ROOT_PASSWORD=secret -d mysql

startmysql:
	docker start mysql
	
stopmysql:
	docker stop mysql

createdb: 
	docker exec mysql mysql -u root -psecret -e  "CREATE DATABASE Editorial"

dropdb: 
	docker exec mysql mysql -u root -psecret -e "SET FOREIGN_KEY_CHECKS=0; DROP DATABASE IF EXISTS ecommerce; SET FOREIGN_KEY_CHECKS=1;"

showdatabase:
	sudo docker exec mysql mysql -u root -psecret -e "SHOW DATABASES;"

.PHONY: runmysql startmysql stopmysql createdb  dropdb startdb