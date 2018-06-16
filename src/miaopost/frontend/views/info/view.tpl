{{if .user}}
<input type="hidden" id="uid" value="{{.user.Id}}">
{{else}}
<input type="hidden" id="uid" value="0">
{{end}}
<div class="view">
	<div class="row meta">
		<div class="col-md-9 col-xs-9">
			<span class="badge">{{.cat.Name}}</span>
			<span>阅读({{.info.Views}}) </span>
			<span>{{showtime .info.Update_time}}</span>
		</div>
		<div class="col-md-2 col-xs-3">
			<a href="#" onclick="suggestDel({{.info.Id}});return false;">建议删除</a>
		</div>
	</div>
	<div class="row">
		<div class="col-md-8">
			{{str2html .info.Content}}  
			<br>
			<span class="text-info more">联系我时请注明来自秒Po</span>
		</div>
		<div class="col-md-8 photos">
			{{range .photos}}
			<img src="{{.P.Url}}{{cutImgSize .Width .Height}}">
			{{end}}
		</div>
	</div>
	{{if  eq .cat.Type 1}}
		{{template "info/infoMsg.tpl" .}}
	{{end}}
	<div class="adv row">
		{{if .adv}}
		<ul>
		{{range .adv}}
		<li><a href="{{if .Link}}{{.Link}}{{else}}/article/view?id={{.Id}}{{end}}"><img src="{{.Logo}}!800!800" /></a></li>
		{{end}}
		</ul>
		{{end}}
	</div>
</div>