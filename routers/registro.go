package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/severuss922/twitterGo/bd"
	"github.com/severuss922/twitterGo/models"
)

func Registro(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi
	r.Status = 400

	fmt.Println("> Entre a registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println("> " + r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el Email"
		fmt.Println("> " + r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "La contraseña debe tener al menos 6 caracteres"
		fmt.Println("> " + r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario registrado con ese email."
		fmt.Println("> " + r.Message)
		return r
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		r.Message = "Ocurrió un error al intentar realizar el registro de usuario: " + err.Error()
		fmt.Println("> " + r.Message)
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el registro del usuario"
		fmt.Println("> " + r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro OK"
	fmt.Println("> " + r.Message)
	return r
}
