<div class="list">
	<div class="row">	
		<input type="hidden" id="hasMore"  value="{{.hasMore}}" />
		<input type="hidden" id="page"  value="{{.page}}" />
		<input type="hidden" id="cid"  value="{{.cid}}" />
		<div class="col-md-10 info-list">
			{{template "info/listPage.tpl" .}}
			<a href="javascript:;" class="load-more">加载更多<span class="loading"></span></a>
		</div>
	</div>
</div>