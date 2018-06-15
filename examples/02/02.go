package main

import (
    "fmt"
    "strconv"
    router "../.."
)

func main() {

    r := router.New()

    articles := map[int]string{}
    current := 1

    fmt.Println("Server started.")

    r.Group("", func (g *router.Group){
        g.Get("/", func(req router.Request) router.Response {
            var re string

            for k, _ := range articles {
                re += "<li>"

                ks := strconv.Itoa(k)

                re += `<a href="/read/` + ks + `.html">Article` + ks + `</a>`

                re += "</li>"
            }

            return &router.HttpResponse{
                StatusCode: 200,
                Body: fmt.Sprintf(`
                    <!Doctype html>
                    <html>
                     <head>
                      <title>Article List</title>
                      <meta charset="utf-8" />
                     </head>
                     <body>
                      <h1>Articles</h1>
                      <ul>
                       %s
                      </ul>
                      <hr>
                      <a href="/">Home...</a>
                      <a href="/add.html">Add...</a>
                     </body>
                    </html>
                    `, re),
            }
        })
        g.Get("/add.html", func(req router.Request) router.Response {
            return &router.HttpResponse{
                StatusCode: 200,
                Body: fmt.Sprintf(`
                    <!Doctype html>
                    <html>
                     <head>
                      <title>Add article</title>
                      <meta charset="utf-8" />
                     </head>
                     <body>
                      <form method="post" action="/add.do">
                        <input type="number" name="id" placeholder="ID..." value="%d" /><br>
                        <textarea name="value" style="height: 100px; width: 200px;">A quick brown fox jumped over a lazy grey pig.</textarea><br>
                        <input type="submit" />
                      </form>
                     </body>
                    </html>
                `, current),
            }
        })
        g.Get("/read/{@int:id}.html", func (req router.Request) router.Response {
            id, _ := strconv.Atoi( req.(*router.HttpRequest).RouterVariable["id"] )
            return &router.HttpResponse{
                StatusCode: 200,
                Body: fmt.Sprintf(`
                    <!Doctype html>
                    <html>
                     <head>
                      <title>Read article</title>
                      <meta charset="utf-8" />
                     </head>
                     <body>
                      <h1>Article id %d</h1>
                      <p>%s</p>
                      <hr>
                      <a href="/">Home...</a>
                     </body>
                    </html>
                `, id, articles[id]),
            }
        })
        g.Post("/add.do", func(req router.Request) router.Response {
            id, _ := strconv.Atoi( req.(*router.HttpRequest).PostFormValue("id") )
            value := req.(*router.HttpRequest).PostFormValue("value")
            current ++
            articles[id] = value
            return &router.HttpResponse{
                StatusCode: 200,
                Body: fmt.Sprintf(`<html><script>window.location.href="/read/%d.html";</script></html>`, id),
            }
        })
    })
	
    router.HttpServe(r, "127.0.0.1:28080")

}