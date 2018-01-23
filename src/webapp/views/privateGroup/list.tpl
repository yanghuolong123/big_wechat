<div class="pg_list">
	<h4 class="title">{{.group.Name}}</h4>
	<div class="row">
	            <div class="col-md-6 form-group">
	              <input type="hidden" id="search_group" name="search_group" value="" />
	              <input type="text"  class="form-control" id="search" placeholder="请输入学校关键字或简称" data-provide="typeahead" autocomplete="off">
	            </div>
	            <div class="col-md-1 form-group">
	              <a href="/pg/create" id="publish_pg" class="btn btn-success" role="button">发布群</a>
	            </div>                            
	</div>
	<input type="hidden" id="gid" name="gid" value="{{.group.Id}}">
	{{if .pgs}}
		{{if .isunlock}}
		<button disabled="disabled" class="btn btn-warning unlock_group">已解锁</button>
		{{else}}
		<button id="unlock_group" {{if .user}} {{else}} disabled="disabled" {{end}} class="btn btn-warning unlock_group">解锁</button>
		{{end}}
		<div class="list row">
	                  	{{range .pgs}}
			<div class="col-sm-6 col-md-3">                              
				<div class="thumbnail">						
						{{if $.isunlock}}
							<a href="/pg/view?id={{.Id}}">{{if .Qrcode}}<img src="/{{.Qrcode}}" alt="{{.Name}}">{{else}}<img src="/static/images/default_lack.jpg" alt="{{.Name}}">{{end}}</a>
						{{else}}
							<a href="/tips/pglist"><img src="/static/images/pglist" alt="{{.Name}}"></a>
						{{end}}						
						<div class="caption">
						        <a href="{{if $.isunlock}}/pg/view?id={{.Id}}{{else}}/tips/pglist{{end}}"><h4>{{.Name}}</h4></a>
						         <p class="text-muted">{{.Introduction}}</p>
						</div>
				</div>
			</div> 		
			 {{end}}
		</div>
	{{else}}
		<div class="alert alert-info" role="alert">
			<p>亲，还没有人发布群呢！</p>
		</div>
	{{end}}
</div>

<div id="unlock_pay" class="modal fade" tabindex="-1" role="dialog">
  <div class="modal-dialog modal-sm" role="document">
    <div class="modal-content">
      <div class="modal-header">      	
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
      </div>
      <div class="modal-body">
      	<p class="text-warning">亲, 你已经免费解锁了两所学校，超过两所需要支付</p>
      	<div class="center-block">
      		<h4>微信扫码支付</h4>
	      	<p class="center-block">
	      		<img id="pay_qr_img" src="/static/images/loading.gif" />
	      	</p>
	        	<p class="center-block">支付金额 : <span class="pay_amount">￥2.00</span></p>
      	</div>
      	
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->