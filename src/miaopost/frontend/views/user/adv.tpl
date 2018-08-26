<div class="row">	
	<div class="col-md-12">
		{{template "user/_menu.tpl" .}} 
	</div>
	<div class="col-md-12">
		
		<div class="panel panel-success account_detail">
		  <!-- Default panel contents -->
		  <div class="panel-heading">我的广告列表</div>
		  
		  <!-- Table -->
		  <table class="table">
		   <thead>
		        <tr>
		          <th>投放时间</th>
		          <th>投放目标</th>
		          <th>广告位置</th>
		          <th>计划投放量</th>
		          <th>已投放量</th>
		          <th>操作</th>
		        </tr>
		      </thead>
		      <tbody>
		      {{range .advs}}
		        <tr>
		          <td>{{date .A.Create_time "Y-m-d H:i:s"}}</td>
		          <td>{{.ARvo.Region.Shortname}}</td>
		          <td>{{.ARvo.Pos.Name}}</td>
		          <td>{{.A.Display_times}}千次</td>
		          <td>{{.A.Display_count}}</td>
		          <td><a href="/adv/edit?id={{.A.Id}}" class="btn btn-primary">编辑</a></td>
		        </tr>
		        {{end}}
		      </tbody>
		  </table>
		</div>

	</div>
</div>