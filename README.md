# messaging

```markdown
# RabbitMQ Go Wrapper

Este proyecto proporciona un wrapper en Go (Golang) para simplificar la interacción con RabbitMQ, permitiendo la publicación y suscripción de mensajes de manera sencilla y eficiente. Está diseñado para ser utilizado en aplicaciones Go que necesitan comunicarse con RabbitMQ para el procesamiento de mensajes.

## Características

- Conexión simplificada a RabbitMQ.
- Funciones para publicar mensajes en colas.
- Funciones para suscribirse a colas y recibir mensajes.
- Configuración de `prefetch count` para optimizar el procesamiento de mensajes.
- Manejo de cierre de canal y conexión de forma segura.

## Requisitos

Para utilizar este wrapper, necesitas tener Go instalado en tu sistema. Además, debes tener acceso a una instancia de RabbitMQ, ya sea localmente o en la nube.

## Instalación

Puedes importar este wrapper directamente en tu proyecto Go utilizando el comando `go get`:

```bash
go get github.com/<tu_usuario>/<tu_repositorio>
```

## Uso

A continuación se muestra un ejemplo básico de cómo utilizar el wrapper para suscribirse a una cola y recibir mensajes:

```go
package main

import (
	"github.com/<tu_usuario>/<tu_repositorio>/rabbitmq"
	"log"
)

func main() {
	// Inicializa el suscriptor
	rabbitmq.InitSubscriber("mi_cola", "amqp", "usuario", "contraseña", "localhost", "5672", 1, handleEvent)
}

// handleEvent se llama cada vez que se recibe un mensaje.
func handleEvent(delivery *amqp.Delivery) {
	log.Printf("Mensaje recibido: %s", delivery.Body)
}
```

## Configuración

Este wrapper permite algunas configuraciones, como el `prefetch count`, que se puede ajustar para optimizar el rendimiento según tus necesidades específicas.

## Contribuciones

Las contribuciones a este proyecto son bienvenidas. Si tienes sugerencias para mejorar o extender las funcionalidades, no dudes en crear un pull request o abrir un issue.

## Licencia

Este proyecto se distribuye bajo la licencia MIT. Ver el archivo `LICENSE` para más detalles.
```

Este `README` proporciona una visión general básica del proyecto, cómo empezar, y dónde encontrar más información. Asegúrate de personalizar los ejemplos de código y las instrucciones según tu repositorio específico en GitHub y las necesidades de tu proyecto.