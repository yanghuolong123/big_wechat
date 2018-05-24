<div class="home">
	<div class="row">
		<div class="input-group col-md-5">
	                          <input type="text" class="form-control"  id="search" placeholder="搜索, eg. 室友，moving">
	                          <span class="input-group-btn">
	                            <button class="btn btn-default search-btn" type="button"><span class="glyphicon glyphicon-search" aria-hidden="true"></span></button>
	                          </span>
	                </div>
	</div>
	<div class="row cats">
		<a href="/info" class="search_cats {{if  not .cid}}label label-primary{{end}}">全部</a> 
	        {{range .cats}}
                            <a href="/info/list?cid={{.Id}}" class="search_cats {{if $.cid}}{{if eq .Id $.cid}}label label-primary{{end}}{{end}}">{{.Name}}</a> 
                        {{end}}
	</div>
	<div class="row list-head">
		<div class="col-md-10  col-xs-9 title">
			<h4>一周浏览 <span class="week_pv">({{.weekviews}})</span></h3>
		</div>
		<div class="col-md-2  col-xs-3">
			<div class="dropdown pull-right">
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
		{{if .cid}}
		<input type="hidden" id="cid"  value="{{.cid}}" />
		{{end}}
		<div class="col-md-12 info-list">
		{{template "info/listPage.tpl" .}}
		<a href="javascript:;" class="load-more">加载更多<span class="loading"></span></a>
		{{template "info/listAdv.tpl" .}}
		</div>
	</div>
</div>
