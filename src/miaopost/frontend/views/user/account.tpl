<div class="row">	
	<div class="col-md-12">
		{{template "user/_menu.tpl" .}} 
	</div>
	<div class="col-md-12">
		<div class="jumbotron account">
		  
		  <p><h4>账户总余额：<span>￥{{.ua.Amount}}</span></h4></p>
		  <p><a class="btn btn-success" href="#" onclick="withDraw();return false;" role="button">提现</a></p>
		</div>

		<div class="panel panel-success account_detail">
		  <!-- Default panel contents -->
		  <div class="panel-heading">账户明细</div>
		  
		  <!-- Table -->
		  <table class="table">
		   <thead>
		        <tr>
		          <th>入账时间</th>
		          <th>金额</th>
		          <th>说明</th>
		        </tr>
		      </thead>
		      <tbody>
		      {{range .uad}}
		        <tr>
		          <td>{{date .Create_time "Y-m-d H:i:s"}}</td>
		          <td>￥{{.Amount}}</td>
		          <td>{{.Remark}}</td>
		        </tr>
		        {{end}}
		      </tbody>
		  </table>
		</div>

	</div>
</div>

<div id="withDrawModal" class="modal fade" tabindex="-1" role="dialog">
  <div class="modal-dialog modal-sm" role="document">
    <div class="modal-content">
      <div class="modal-header">      	
         <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
         <h4 class="modal-title">提现到微信余额</h4>
      </div>
      <div class="modal-body">
      	<div class="center-block">   
      		<p class="center-block">最多可以提现金额 : <span class="pay_amount">￥{{.ua.Amount}}</span></p>   		
	      	<p class="center-block">
	      		￥<input type="text" id="withdraw_amount">
	      	</p>
	      	<p class="center-block">
	      	<button type="button" class="btn btn-danger with-draw-btn">确认转出</button>
	      	</p>
	        	
      	</div>
      	
      </div>
    </div><!-- /.modal-content -->
  </div><!-- /.modal-dialog -->
</div><!-- /.modal -->