<div class="box">
            <div class="box-header with-border">
              <h3 class="box-title">{{if eq .advtype 1}}已投放{{else if eq .advtype 2}}待支付{{end}}广告列表 </h3>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <table id="adv_list" class="table table-bordered table-hover">
                <thead>
                <tr>                  
                  <th>ID</th>
                  <th>投放目标</th>
                  <th>广告位置</th>
                  <th>计划投放量</th>
                  <th>已投放量</th>
                  <th>点击量</th>
                  <th>单价(每千次)</th>
                  <th>推荐码</th>
                  <th>支付金</th>
                  <th>投放商家</th>
                  <th>联系方式</th>
                  <th>总部收入</th>
                  <th>运营者收入</th>
                  <th>状态</th>
                  <th>投放日期</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>

                {{range .dataList}}
                <tr>
                  <td>{{.A.Id}}</td>
                  <td>{{.ARvo.Region.Shortname}}</td>
                  <td>{{.ARvo.Pos.Name}}</td>
                  <td>{{.A.Display_times}}</td>
                  <td>{{.A.Display_count}}</td>
                  <td></td>
                  <td>{{.A.Amount}}</td>
                  <td>{{.A.Recom_code}}</td>
                  <td>{{.A.Total_amount}}</td>
                  <td>{{.A.Merch_name}}</td>
                  <td>{{.A.Contact}}</td>
                  <td>{{.A.Head_income}}</td>
                  <td>{{.A.Operator_income}}</td>
                  <td>{{.StatusLabel}}</td>
                  <td>{{date .A.Create_time "Y-m-d"}}</td>
                  <td>
                    {{if eq $.advtype 1}}
                      {{if eq .A.Status 2}}
                      <a href="javascript:;" onclick="updateStatus({{.A.Id}},1)" class="btn btn-success">恢复投放</a>
                      {{else}}
                     <a href="javascript:;" onclick="updateStatus({{.A.Id}},2)" class="btn btn-success">暂停投放</a>
                      {{end}}                      
                    {{ else if eq $.advtype 2}}
                     <a href="javascript:;" onclick="updateStatus({{.A.Id}},1)" class="btn btn-danger">免费投放</a> 
                    {{end}}                                      
                    <a href="javascript:;" onclick="updateStatus({{.A.Id}},-1)" class="btn btn-warning">删除</a>                   
                    </td>
                </tr>
                {{end}}
                
                </tbody>
                
              </table>
            </div>
            <!-- /.box-body -->
          </div>
          <!-- /.box -->
<!-- DataTables -->
<link rel="stylesheet" href="/static/plugin/bower_components/datatables.net-bs/css/dataTables.bootstrap.min.css">
<script src="/static/plugin/bower_components/datatables.net/js/jquery.dataTables.min.js"></script>
<script src="/static/plugin/bower_components/datatables.net-bs/js/dataTables.bootstrap.min.js"></script>
<script type="text/javascript">
  $(function(){

    $('#adv_list').DataTable({
        'paging'      : true,
        'lengthChange': false,
        'searching'   : false,
       'ordering'    : false,
        'info'        : true,
        'autoWidth'   : false
      })

  });

   var updateStatus = function(id, status) {
      $.post("/adv/updateStatus",{id:id,status:status},function(e){
          if(e.code<0) {
            prompt(e.msg);
            return false;
          }

          greeting({msg:"广告位状态修改成功！",confirm:function(){
                location.href  = location.href ;
          }});

      });
    }
</script>