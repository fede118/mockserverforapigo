# mockserverforapigo
Mock server for Go API

end points:

	"/users/" -> json: Users
	"/countries/" -> json: Countries
	"/sites/" -> json: Sites

	"/users/ping" -> pong
	"/countries/ping" -> pong
	"/sites/ping" -> pong

Para genererar errores en la Api que hace llamados a este mock podemos:
  - cambiar la funcion de cada endpoint por la funcion return500() que devuelve un status code 500
  - aumentar el tiempo de respuesta a > 5000ms
  - no subir el mock server directamente
