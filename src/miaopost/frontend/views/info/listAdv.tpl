<div class="list_adv row">
{{if .adv_list}}
<ul>
{{range .adv_list}}
<li><a href="{{if .Link}}{{.Link}}{{else}}/article/view?id={{.Id}}{{end}}"><img src="{{.Logo}}!800!800" /></a></li>
{{end}}
</ul>
{{end}}
</div>
