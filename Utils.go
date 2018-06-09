package router

import (
	
)

var DefaultNotFoundResponse = &DefaultResponse{
		StatusCode: 404,
		Body: `<!Doctype html>
<html>
 <head>
  <meta charset="utf-8" />
 </head>
 <body>
  <h1 align="center">HTTP 404</h1>
  <p align="center">Not Found</p>
  <hr />
  <p><i>Powered by xtlsoft/router (NotFound Module)</i></p>
 </body>
</html>`,
}

var DefaultNotFoundController = func(req Request) Response {
	return DefaultNotFoundResponse
}