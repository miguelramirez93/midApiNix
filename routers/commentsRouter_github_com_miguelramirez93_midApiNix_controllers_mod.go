package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"RegistrarPadre",
			`registrar_padre`,
			[]string{"post"},
			nil})

}
