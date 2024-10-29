package documentservice

import "net/http"


var postopt = []string{http.MethodOptions, http.MethodPost}

func (s *Rest) initializeRoutes() {

	api := s.Routes.Group("/api", CORSMiddleware())

	save := api.Group("")
	{
		//сохранение картинки
		save.Match(postopt, "/register", s.RegisterUser)
	}


}