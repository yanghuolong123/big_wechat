<footer class="footer">
            <div class="container">
                <div class="footer_top">
                	{{range  $index,$v := .footer_nav}}
                	<a href="/article/view?id={{.Id}}">{{.Title}}</a>
                	{{if ne $index $.last_footer_nav_index}}
                	<em> · </em>
                	{{end}}
                	{{end}}
                </div>
            </div>
        </footer>