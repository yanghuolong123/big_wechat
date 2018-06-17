<div class="row">	
	<div class="col-md-12">
		{{template "user/_menu.tpl" .}} 
	</div>
	<div class="col-md-12">
		<div class="jumbotron account">
		  
		  <p><h4>账户总余额：<span>￥{{.ua.Amount}}</span></h4></p>
		  <p><a class="btn btn-success" href="javascript:;" role="button">提现</a></p>
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