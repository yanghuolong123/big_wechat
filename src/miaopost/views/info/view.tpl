<div class="view">
	<div class="row meta">
		<div class="col-md-offset-1 col-md-7 col-xs-9">
			<span class="badge">{{.cat.Name}}</span>
			<span>阅读({{.info.Views}}) </span>
			<span>{{showtime .info.Create_time}}</span>
		</div>
		<div class="col-md-2 col-xs-3">
			<a href="#" onclick="suggestDel({{.info.Id}});return false;">建议删除</a>
		</div>
	</div>
	<div class="row">
		<div class="col-md-offset-1 col-md-8">
			{{str2html .info.Content}}  
			<br>
			<span class="text-info more">联系我时请注明来自秒Po</span>
		</div>
		<div class="col-md-offset-1 col-md-8 photos">
			{{range .photos}}
			<img src="{{.Url}}">
			{{end}}
		</div>
	</div>
</div>