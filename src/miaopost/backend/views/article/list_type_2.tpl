<div class="box">
            <div class="box-header with-border">
              <h3 class="box-title">广告列表 </h3>
              <a href="/article/create?type={{.type}}" class="btn btn-info  pull-right">创建广告</a>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <table id="article_list" class="table table-bordered table-hover">
                <thead>
                <tr>                  
                  <th>ID</th>
                  <th>区域</th>
                  <th>分组</th>
                  <th>标题</th>
                  <th>Logo图标</th>
                  <th>外部链接</th>
                  <th>排序</th>
                  <th>状态</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>

                {{range .dataList}}
                <tr>
                  <td>{{.Id}}</td>
                  <td>{{map_get $.regions .Rid}}</td>
                  <td>{{map_get (map_get $.groupMap .Type) .Group_id}}</td>
                  <td>{{.Title}}</td>
                  <td>{{if .Logo}}<img src="{{.Logo}}" style="height: 150px; width: 250px;">{{end}}</td>
                  <td>{{.Link}}</td>
                  <td> {{.Sort}}</td>
                  <td>{{map_get $.statusMap .Status}}</td>
                  <td>
                    <a href="/article/edit?id={{.Id}}&type={{$.type}}"><i class="fa fa-wrench"></i></a>                    
                    <a href="/article/delete?id={{.Id}}&type={{$.type}}" class="delete"><i class="fa fa-times"></i></a>
                    <a href="{{if .Link}}{{.Link}}{{else}}http://www.miaopost.com/article/view?id={{.Id}}{{end}}" target="_blank"><i class="fa fa-external-link"></i></a>
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