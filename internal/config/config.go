package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs"
)

/*set configuration so that once have template set. no need to load again
untill application restarts
this package can be imported by other pkg in application but it doesnt import anything else
from application it will help to avoid import cycle*/

//holds the application config
//any kind of config can add to struct that need for application

/*use cache will help to add any newly add comment in html (about)template while
application is still running i.e during development phase*/
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
