## L0

### Nats-streaming server развернут локально
```sh
nats-streaming-server
```
![image](https://user-images.githubusercontent.com/61902513/176861434-0b52c88f-5bac-4f38-b0ef-42a41ecf438d.png)
### Postgres в докер-контейнере
![image](https://user-images.githubusercontent.com/61902513/176861594-f7319bbb-b683-4e69-bf6d-c1744486fbbd.png)

#### WRK тест
```sh
 wrk -t12 -c400 -d30s http://localhost:8080
 ```
 ![image](https://user-images.githubusercontent.com/61902513/176861242-3d667c19-15f1-4a78-a0b3-f14e8b9312ab.png)
