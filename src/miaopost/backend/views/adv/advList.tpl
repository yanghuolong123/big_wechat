<div class="box">
            <div class="box-header with-border">
              <h3 class="box-title">广告列表 </h3>
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
                  <th>状态</th>
                  <th>投放日期</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>

                {{range .dataList}}
                <tr>
                  <td>{{.A.Id}}</td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td></td>
                  <td>
                    <a href="/"><i class="fa fa-wrench"></i></a>                    
                    <a href="/" class="delete"><i class="fa fa-times"></i></a>
                    <a href="" target="_blank"><i class="fa fa-external-link"></i></a>
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

    $("a.delete").click(function(){
        var ln = $(this);
       actionConfirm({msg:"确定要删除？",confirm:function(){
          location.href = ln.attr("href");
      }});
      return false;
    });

  });
</script>