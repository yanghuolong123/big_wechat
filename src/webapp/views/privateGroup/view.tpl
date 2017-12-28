<div class="view_pg">
	<div class="viewinfo">
		<div class="row">
			<div class="col-md-9">
				<h4>{{.pg.Name}}</h4>
			</div>
			<div class="col-md-3">
				<a class="glink" href="/pg/list"><span class="glyphicon glyphicon-align-left" aria-hidden="true"></span> {{.group.Name}}</a> 
				<button class="btn btn-warning btn-sm">举报</button>
			</div>
		</div>
		<div>
			<span>发布时间：{{.pg.Createtime}}</span>
		</div>
		<p class="desc">
			{{.pg.Introduction}}
		</p>
		<div class="row qrcode">
			 <div class="col-md-2 ">				 
					 <img class="img-rounded" src="/{{.pg.Qrcode}}" alt="群二维码" />					  
					 <p class="qr1">群二维码</p> 
			 </div>
			 <div class="col-md-2">				 
					 <img class="img-rounded" src="/{{.pg.Ower_qrcode}}" alt="群主二维码" />					 
					<p class="qr2">群主二维码/微信号: {{.pg.Wechat_id}}</p>
				  
			 </div>
		</div>		
	</div>
	<div class="comment ">
		<h4 class="text-muted">留言</h4>
		<div class="row">	
			<div class="col-md-7">	 
				<textarea class="form-control"></textarea>
			</div>
		</div>		 
		<div class="comment_btn col-sm-offset-6">
			<button class="btn btn-success">提交</button>
		</div>
		<ul class="commentlist">
			<li>
				<h5>11</h5>
				<p>22</p>
				<p>33</p>
			</li>
		</ul>
	</div>
</div>