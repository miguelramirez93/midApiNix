package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:EntidadController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"] = append(beego.GlobalControllerRouter["github.com/miguelramirez93/midApiNix/controllers:RubroController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil})

}
