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
		<div class="col-md-8  col-xs-9 title">
			<h4>最新发布</h3>
		</div>
		<div class="col-md-2  col-xs-2">
			<div class="dropdown">
	                                              <button class="btn btn-primary dropdown-toggle" type="button" id="publish-btn" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
	                                                发布    
	                                                <span class="caret"></span>
	                                              </button>
	                                              <ul class="dropdown-menu" aria-labelledby="publish-btn">
	                                              	{{range .cats}}
	                                               	 <li><a href="/info/create?cid={{.Id}}">{{.Name}}</a></li>
	                                                {{end}}
	                                              </ul>
                                           	</div>
			<!--<a href="/info/create" class="btn btn-primary active" role="button">发布</a>-->
		</div>
	</div>
	<div class="row">
		<input type="hidden" id="hasMore"  value="{{.hasMore}}" />
		<input type="hidden" id="page"  value="{{.page}}" />
		<div class="col-md-10 info-list">
		{{range .infos}}
		<div class="info">
			<div class="row">
				<div class="col-md-7 col-xs-3 cat"><span class="label label-warning">{{.Cat.Name}}</span></div>
				<div class="col-md-4 col-xs- 9 meta text-right">
					<span>{{showtime .Info.Create_time}} </span>
					<span>阅读({{.Info.Views}})</span>
					<span><a href="#" onclick="suggestDel({{.Info.Id}});return false;">建议删除</a></span>
				</div>
			</div>
			<div class="row">
				<div class="info-content col-md-11">
					{{substr .Info.Content 0 150}} {{if (.Photos|len)}}<img class="img_tip" src="/static/img/image_s.png">{{end}}...
					<a href="/info/view?id={{.Info.Id}}" class="more">more ››</a>
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
