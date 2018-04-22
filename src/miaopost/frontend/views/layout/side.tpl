{{if  not .isMobile}}
	<div class="side">
	{{if .side_adv}}
	<ul>
	{{range .side_adv}}
	<li><a href="{{if .Link}}{{.Link}}{{else}}/article/view?id={{.Id}}{{end}}"><img src="{{.Logo}}!300!300" /></a></li>
	{{end}}
	</ul>
	{{end}}
	</div>

	<div class="side">
	{{if .side_adv_1}}
	<ul>
	{{range .side_adv_1}}
	<li><a href="{{if .Link}}{{.Link}}{{else}}/article/view?id={{.Id}}{{end}}"><img src="{{.Logo}}!300!300" /></a></li>
	{{end}}
	</ul>
	{{end}}
	</div>
{{end}}