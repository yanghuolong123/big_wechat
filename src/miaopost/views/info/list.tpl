<div class="list">
	<div class="row">		
		<div class="col-md-10 info-list">
			{{range .infos}}
			<div class="info">
				<div class="row">
					<div class="col-md-6 cat"><span class="label label-warning">{{.Cat.Name}}</span></div>
					<div class="col-md-4 meta">
						<span>{{showtime .Info.Create_time}} </span>
						<span>阅读({{.Info.Views}})</span>
						<span><a href="#" onclick="suggestDel({{.Info.Id}});return false;">建议删除</a></span>
					</div>
				</div>
				<div class="row">
					<div class="info-content col-md-10">
						{{.Info.Content}}
						<a href="/info/view?id={{.Info.Id}}" class="more">more>></a>
						<div class="line"></div>
					</div>
				</div>
			</div>
			{{else}}
			<div class="alert alert-warning col-md-10" role="alert">亲，还没有数据哦！</div>
			{{end}}
		</div>
	</div>
</div>