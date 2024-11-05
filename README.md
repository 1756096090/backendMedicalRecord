### Estructura del Backend

1. **Modelos**: La aplicación cuenta con una capa de modelos donde se definen las estructuras de datos que reflejan la base de datos en MongoDB. Los modelos incluyen:
   - **Patient**
   - **Role**
   - **Specialist**
   - **User**

2. **Repositorios**: Cada modelo tiene un repositorio dedicado que se encarga de las consultas a la base de datos:
   - `patient_repository`
   - `role_repository`
   - `specialist_repository`
   - `user_repository`

3. **Controladores**: Los controladores manejan la lógica de negocio y las validaciones, además de mapear datos. Por ejemplo, en `user_controller` se generan tokens para el inicio de sesión.

4. **Rutas**: Cada entidad tiene su propio conjunto de rutas. Las rutas están organizadas de la siguiente manera:
   - `patient_route.go`
   - `role_route.go`
   - `specialist_route.go`
   - `user_route.go`
   - Un archivo de rutas general que permite establecer los endpoints:

   ```go
   package routes

   import (
       "github.com/gorilla/mux"
   )

   func SetupRoutes() *mux.Router {
       router := mux.NewRouter()
       
       userRouter := router.PathPrefix("/user").Subrouter() 
       UserRoutes(userRouter)

       roleRouter := router.PathPrefix("/role").Subrouter()
       RoleRoutes(roleRouter)

       specialistRouter := router.PathPrefix("/specialist").Subrouter()
       SpecialistRoutes(specialistRouter)

       patientRoutes := router.PathPrefix("/patient").Subrouter()
       PatientRoutes(patientRoutes)

       return router
   }

```
5. ### Configuración de la Base de Datos

La configuración de la base de datos se encuentra en `db.go`, donde se establece la conexión con **MongoDB Atlas**. Aquí se incluyen las credenciales necesarias para la conexión:

```go
clientOptions := options.Client().ApplyURI("mongodb+srv://user_test:Ismacs2003@firstproyectwebengineer.b6xlw.mongodb.net/?retryWrites=true&w=majority&appName=FirstProyectWebEngineering")
```
6. Archivo Main
En main.go, se configuran los CORS y se establece el puerto de acceso para la aplicación, que es el `8080`. Este archivo es esencial para iniciar el servidor y manejar las configuraciones iniciales de la aplicación.
7. Ejecución
Para iniciar el servidor, utiliza el siguiente comando en la terminal:
```bash
    go run main.go

