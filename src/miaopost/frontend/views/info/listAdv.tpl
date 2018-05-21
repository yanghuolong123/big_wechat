<div class="list_adv row">
{{if .adv_list}}
<ul>
{{range .adv_list}}
<li><a href="/article/view?id={{.Id}}"><img src="{{.Logo}}!800!800" /></a></li>
{{end}}
</ul>
{{end}}
</div>
