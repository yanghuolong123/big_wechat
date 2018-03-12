<div class="home">
	<div class="row">
		<div class="input-group col-md-5">
	                          <input type="text" class="form-control"  id="search" placeholder="搜索, eg.单人间">
	                          <span class="input-group-btn">
	                            <button class="btn btn-default search-btn" type="button"><span class="glyphicon glyphicon-search" aria-hidden="true"></span></button>
	                          </span>
	                </div>
	</div>
	<div class="row cats">
	        {{range .cats}}
                            <a href="/info/list?cid={{.Id}}" class="search_cats">{{.Name}}</a> 
                        {{end}}
	</div>
	<div class="row list-head">
		<div class="col-md-8 title">
			<h4>最新发布</h3>
		</div>
		<div class="col-md-2">
			<a href="/info/create" class="btn btn-primary active" role="button">发布</a>
		</div>
	</div>
	<div class="info-list row">
		{{range .infos}}
		<div class="info">
			<div class="row">
				<div class="col-md-6 cat"><span class="label label-warning">{{.Cat.Name}}</span></div>
				<div class="col-md-4 meta">
					<span>{{showtime .Info.Create_time}} </span>
					<span>阅读({{.Info.Views}})</span>
					<span><a href="#">建议删除</a></span>
				</div>
			</div>
			<div class="row">
				<div class="info-content col-md-9">
					{{.Info.Content}}
					<a href="/info/view?id={{.Info.Id}}" class="more">more>></a>
					<div class="line"></div>
				</div>
			</div>
		</div>
		{{end}}
	</div>
</div>
