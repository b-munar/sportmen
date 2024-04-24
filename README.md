# Instalacion

1. Clonar repositorio

```bash
git clone https://github.com/b-munar/sportmen.git
```

2. Copiar el .env.template y pegarlo en un .env

3. 

```bash
docker compose build
docker compose up
```


El servicio de sportmen en si ofrece un CRUD de los datos de los deportistas, este depende totalmente del microservicio auth, se recomienda instalarlo primero, https://github.com/b-munar/auth.

## 1. Creación de usuarios

Crea un deportista con los datos brindados, el id del usuario debe ser único. Este id debe estar registrado en el servicio auth.

<table>
<tr>
<td> Método </td>
<td> POST </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6250/sportmen</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>N/A</td>
</tr>
<tr>
<td> Cuerpo </td>
<td>

```json
{
  "user": "29a3ad78-6d3c-46e3-9c42-857ca3ec5220",
  "name": "Brahian",
  "last_name": "Munar",
  "age": 26,
  "weight": 63,
  "height": 163,
  "gender": "dog",
  "identification_type":"CC",
  "identification":"314159",
  "country_birth": "Colombia",
  "city_birth": "Cali",
  "country_residence": "Colombia",
  "city_residence": "Elvira",
  "length_residence": 26,
  "sports": [{"sport":"basketball"}]
  }
```
</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 201 </td>
<td>En el caso que el deportista se haya creado con éxito.</td>
<td>

```json
{
  "sportmen": {
    "name": "Brahian",
    "last_name": "Munar",
    "age": 26,
    "weight": 63,
    "height": 163,
    "gender": "dog",
    "identification_type": "CC",
    "identification": "314159",
    "country_birth": "Colombia",
    "city_birth": "Cali",
    "country_residence": "Colombia",
    "city_residence": "Elvira",
    "length_residence": 26,
    "created_at": "2024-04-05T20:08:43.807690403Z",
    "updated_at": "2024-04-05T20:08:43.807690403Z",
    "sports": [
      {
        "sport": "basketball"
      }
    ]
  }
}
```
</td>
</tr>
</tbody>
</table>


## 2. Obtener deportista

Obtener Deportista
<table>
<tr>
<td> Método </td>
<td> GET </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6250/sportmen</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>
 "Authorization": "Bearer eyJ0eXAiOiJKV1QiL..."
</td>
</tr>
<tr>
<td> Cuerpo </td>
<td>
N/A
</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 200 </td>
<td>En el caso de exito.</td>
<td>

```json
{
  "sportmen": {
    "name": "Brahian",
    "last_name": "Munar",
    "age": 26,
    "weight": 63,
    "height": 163,
    "gender": "dog",
    "identification_type": "CC",
    "identification": "314159",
    "country_birth": "Colombia",
    "city_birth": "Cali",
    "country_residence": "Colombia",
    "city_residence": "Elvira",
    "length_residence": 26,
    "created_at": "2024-04-05T20:08:43.80769Z",
    "updated_at": "2024-04-05T20:08:43.80769Z",
    "sports": [
      {
        "sport": "basketball"
      }
    ]
  }
}
```
</td>
</table>

## 3. Obtener deportista por parte de un partner

Obtener Deportista
<table>
<tr>
<td> Método </td>
<td> Post </td>
</tr>
<tr>
<td> Ruta </td>
<td> <strong>localhost:6250/sportmen/info</strong> </td>
</tr>
<tr>
<td> Parámetros </td>
<td> N/A </td>
</tr>
<tr>
<td> Encabezados </td>
<td>
 "Authorization": "Bearer eyJ0eXAiOiJKV1QiL..."
</td>
</tr>
<tr>
<td> Cuerpo </td>
<td>
{
  "user": "9cb63a2e-76c5-4b8e-8434-cd7334a117b0"
  }
</td>
</tr>
</table>

### Respuestas

<table>
<tr>
<th> Código </th>
<th> Descripción </th>
<th> Cuerpo </th>
</tr>
<tbody>
<td> 200 </td>
<td>En el caso de exito.</td>
<td>

```json
{
  "sportmen": {
    "name": "Brahian",
    "last_name": "Munar",
    "age": 26,
    "weight": 63,
    "height": 163,
    "gender": "dog",
    "identification_type": "CC",
    "identification": "314159",
    "country_birth": "Colombia",
    "city_birth": "Cali",
    "country_residence": "Colombia",
    "city_residence": "Elvira",
    "length_residence": 26,
    "created_at": "2024-04-05T20:08:43.80769Z",
    "updated_at": "2024-04-05T20:08:43.80769Z",
    "sports": [
      {
        "sport": "basketball"
      }
    ]
  }
}
```
</td>
</table>