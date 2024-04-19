# Weather predictor 🌦️

En una galaxia lejana, existen tres civilizaciones: Vulcanos, Ferengis y Betazoides. Cada civilización vive en paz en su respectivo planeta.

Dominan la predicción del clima mediante un complejo sistema informático que se describe a continuación.

## Alcance
### Casos de uso

* Se debe calcular y almacenar en una base de datos la información del clima durante 10 años, a partir de la ejecución del programa.
* Se debe predecir la cantidad de períodos de sequía que habrá.
* Se debe predecir los períodos de lluvia y el día que habrá un pico máximo.
* Se debe predecir los períodos de condiciones óptimas.
* Se debe predecir las condiciones para un día particular.

### Suposiciones

* La unidad de tiempo que determina la velocidad angular es de un día terrestre
* Un año está compuesto por 365 días terrestres
* Un período se entiende como el conjunto de días seguidos en el que el clima se mantiene. Cuando el clima cambia, se cuenta como un período nuevo

## Arquitectura de la solución 🏗️
### Modelo de datos

<p align="center">
<img width="450" alt="Captura de pantalla 2024-04-11 a la(s) 4 52 38 p m" src="https://github.com/dmedinao11/weather-predictor/assets/62181435/70782db5-794a-479c-8cac-b91561dd1bf1">
</p>

### Diagramas de secuencia
#### * Cálculo y almacenamiento en la base de datos de la información del clima.
![Captura de pantalla 2024-04-12 a la(s) 7 52 07 a m](https://github.com/dmedinao11/weather-predictor/assets/62181435/49c7d65e-38db-445c-8459-fce3c1cefc77)
##### Definición del servicio
<img width="956" alt="Captura de pantalla 2024-04-12 a la(s) 8 03 36 a m" src="https://github.com/dmedinao11/weather-predictor/assets/62181435/fe4aa6b3-a521-466e-8ad6-0edeae734aa0">

#### * Lectura de todas las predicciones.
![Captura de pantalla 2024-04-12 a la(s) 7 50 26 a m](https://github.com/dmedinao11/weather-predictor/assets/62181435/f04087cc-ce7d-4068-b842-a96867d17dee)
##### Definición del servicio
<img width="953" alt="Captura de pantalla 2024-04-12 a la(s) 8 04 20 a m" src="https://github.com/dmedinao11/weather-predictor/assets/62181435/cfb8eaa0-65fc-4665-90f7-6c86662f833d">

#### * Lectura de la predicción  para un día.
![Captura de pantalla 2024-04-12 a la(s) 7 46 50 a m](https://github.com/dmedinao11/weather-predictor/assets/62181435/ebac2b4a-1689-452b-9833-47ec0b96ebde)
##### Definición del servicio
![Captura de pantalla 2024-04-12 a la(s) 9 29 04 a m](https://github.com/dmedinao11/weather-predictor/assets/62181435/f1253dbf-2952-43af-832a-4ec181266f83)

## Ejecutando la app 🚀

Estas instrucciones te permitirán obtener una copia del proyecto en funcionamiento en tu
ambiente local para propósitos de desarrollo y pruebas.

### Instalación 🔧

Paso a paso para ejecutar el proyecto desde tu computadora

* Clona el repositorio

```
git clone https://github.com/dmedinao11/weather-predictor
```

* Abre la carpeta e inicia con docker compose

```
cd weather-predictor
docker-compose up  --build
```

* Ejecuta el health check, e inicia la app

 ```
curl http://localhost:8080/ping
```

## Autor ✒️

- **Daniel Medina** - _Desarrollo_ - [dmedinao11](https://github.com/dmedinao11)

---

⌨️ con ❤️ por [dmedinao11](https://github.com/dmedinao11) 😊
