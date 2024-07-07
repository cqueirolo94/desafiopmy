Se proporciona un docker-compose con 2 imagenes, el servicio publicador de ordenes y una imagen de NATS plano donde se publican las ordenes

El topic de publicación es `pmy.order.{symbol}` por ej `pmy.order.GGAL`
La lista de simbolos para publicación de ordenes es la siguiente

ALUA    
BBAR
BMA
BYMA
CEPU
COME
CRES
EDN
GGAL
IRSA
LOMA
MIRG
PAMP
SUPV
TECO2
TGNO4
TGSU2
TRAN
TXAR
VALO
YPFD

Para la suscripción se pueden utilizar wildcards (https://docs.nats.io/nats-concepts/subjects#wildcards) por ej "pmy.order.*" 
o bien uno a uno según la arquitectura que se desee implementar.

La imagen de NATS tiene expuesto el puerto 4222 para conectarse desde el exterior del compose. Las URL de conexión son
    - Desde el exterior `nats://localhost:4222`
    - Desde la red interna del compose: `nats://pmynats:4222`
    
El makefile posee 4 niveles donde la cantidad de ordenes publicadas es incremental.
1 - 100
2 - 10000
3 - 100000
4 - sin restricción
(valores aproximados)

Para levantar los servicios se debe utilizar el comando `make level1` o el nivel que se desee