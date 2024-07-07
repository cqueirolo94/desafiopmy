# desafiopmy

## Para correr el sistema
1. Abrir una terminal en `./compose-app/sender`, y ejecutar `make level4`
2. Abrir una terminal en `./compose-app` y ejecutar `docker-compose up`
3. Habiendo completado los pasos `1` y `2`, desde `postman`:
   1. `GET http://localhost:9001/order/match/{symbol}` para obtener la cantidad de ordener que hicieron match.
   2. `GET http://localhost:9001/order/pending/{symbol}/prices` para obtener los precios maximos y minimos de las ordenes pendientes, para compra y venta.