<div class="list">
	<div class="row">
		<div class="col-md-2">
			<h4>{{.cat.Name}}</h4>
		</div>
		<div class="col-md-8">
			{{range .infos}}
			<div>
				<div class="row">
					<div class="col-md-2"></div>
					<div class="col-md-8">
						{{showtime .Create_time}} 
						阅读({{.Views}})
						<a href="#">建议删除</a>
					</div>
				</div>
				<div>
					{{.Content}}
					<a href="/info/view?id={{.Id}}">read more >></a>
				</div>
			</div>
			{{end}}
		</div>
	</div>
</div>