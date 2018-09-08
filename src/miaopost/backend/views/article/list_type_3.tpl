<div class="box">
            <div class="box-header with-border">
              <h3 class="box-title">广告发表入口页面配置 </h3>
              {{if eq (.dataList|len) 0}}
              <a href="/article/create?type={{.type}}" class="btn btn-info  pull-right">创建广告</a>
              {{end}}
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <table id="article_list" class="table table-bordered table-hover">
                <thead>
                <tr>
                  <th>内容</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>

                {{range .dataList}}
                <tr>
                  <td>{{str2html .Content}}</td> 
                  <td>
                    <a href="/article/edit?id={{.Id}}&type={{$.type}}"><i class="fa fa-wrench"></i></a> 
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

    $('#article_list').DataTable({
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