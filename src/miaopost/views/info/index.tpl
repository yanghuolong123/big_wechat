<div class="home">
	<div>
		<div class="input-group">
                          <input type="text" class="form-control" placeholder="搜索, eg.单人间">
                          <span class="input-group-btn">
                            <button class="btn btn-default" type="button"><span class="glyphicon glyphicon-search" aria-hidden="true"></span></button>
                          </span>
                        </div>
	</div>
	<div>
		 {{range .cats}}
                            <a href="/info/list?cid={{.Id}}" class="search_cats">{{.Name}}</a> 
                        {{end}}
	</div>
	<div class="row">
		<div class="col-md-5">
			<h4>最新发布</h3>
		</div>
		<div class="col-md-2">
			<a href="/info/create" class="btn btn-primary active" role="button">发布</a>
		</div>
	</div>
	<div class="new-list">
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
