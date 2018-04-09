<div class="box">
            <div class="box-header with-border">
              <h3 class="box-title">文章列表 </h3>
              <a href="/article/create" class="btn btn-info  pull-right">创建文章</a>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <table id="article_list" class="table table-bordered table-hover">
                <thead>
                <tr>
                  <th>ID</th>
                  <th>标题</th>
                  <th>内容</th>
                  <th>排序</th>
                  <th>操作</th>
                </tr>
                </thead>
                <tbody>

                {{range .dataList}}
                <tr>
                  <td>{{.Id}}</td>
                  <td>{{.Title}}</td>
                  <td>{{str2html  (substr .Content 0  100)}}</td>
                  <td> {{.Sort}}</td>
                  <td>
                    <a href="/article/edit?id={{.Id}}"><i class="fa fa-wrench"></i></a>
                    <a href="http://www.miaopost.com/article/view?id={{.Id}}" target="_blank"><i class="fa fa-share"></i></a>
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
  });
</script>