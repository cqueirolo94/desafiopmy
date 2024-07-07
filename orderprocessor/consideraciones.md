# Desafío Primary
## Consideraciones
* Las ordenes tienen la forma 
  ```json 
  {
    "symbol": "SONIC", 
    "price": 245, 
    "type": "buy"
  }
  ```
  * Ocupa en memoria 13 bytes:
    * `symbol`: tiene como máximo 5 caracteres ASCII, entonces ocupa 5 bytes.
    * `price`: es como mucho un uint32, entonces pesa 32 bits o 4 bytes.
    * `type`: tiene como máximo 4 caracteres ASCII, entonces ocupa 4 bytes.
* Las órdenes se publican a tópicos según el símbolo, y viendo el `readme.md` hay `21` tópicos para escuchar.
* Las órdenes llegan de a una, y hay que guardarlas para poder procesarlas.
* Los endpoints no tienen un formato establecido.
* Las ordenes no necesariamente deben ser procesadas en orden.