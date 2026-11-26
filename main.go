package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.TODO()

		// Configuramos un endpoint personalizado para apuntar a nuestro Floci local
		customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           "http://host.docker.internal:4566", // Usamos host.docker.internal para que el contenedor hable con el host de Windows
				SigningRegion: "us-east-1",
			}, nil
		})

		// Cargamos la configuración de AWS con el resolver personalizado
		cfg, err := config.LoadDefaultConfig(ctx,
			config.WithRegion("us-east-1"),
			config.WithEndpointResolverWithOptions(customResolver),
		)
		if err != nil {
			fmt.Fprintf(w, "Error cargando config de AWS: %v", err)
			return
		}

		// Creamos el cliente de S3
		s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true // Obligatorio para emuladores locales de S3
		})

		// Listamos los buckets disponibles en Floci
		output, err := s3Client.ListBuckets(ctx, &s3.ListBucketsInput{})
		if err != nil {
			fmt.Fprintf(w, "Error listando buckets de S3: %v", err)
			return
		}

		// Mostramos la respuesta en pantalla
		fmt.Fprintf(w, "¡Hola desde mi API en Go conectada a Floci (AWS S3 local)!\n\nBuckets encontrados:\n")
		for _, b := range output.Buckets {
			fmt.Fprintf(w, "- %s\n", *b.Name)
		}
	})

	fmt.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}