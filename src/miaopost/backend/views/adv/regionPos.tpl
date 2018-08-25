<div class="box">
            <div class="box-header with-border">
              <h3 class="box-title">广告位价格列表 </h3>
            </div>
            <!-- /.box-header -->
            <div class="box-body">
              <table id="article_list" class="table table-bordered table-hover">
                <thead>
                <tr>
                  <th>区域</th>
                  <th>位置</th>
                  <th>价格</th>
                </tr>
                </thead>
                <tbody>

                {{range .vos}}
                <tr>
                  <td>{{.Region.Shortname}} ({{.Region.Fullname}})</td> 
                  <td>{{.Pos.Name}}</td>
                  <td><input type="hidden" name="adv_re_id" value="{{.AdvRe.Id}}" /><input type="text" class="price" value="{{.AdvRe.Price}}"> </td>
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

    $(".price").change(function(){
        var price = $(this).val();
        var id = $(this).prev().val();
        if(!isMoney(price) || price<=0) {
            prompt("请填写正确金额");
            return;
        }

        $.post("/adv/updatePosPrice",{id:id,price:price}, function(e){
              if(e.code<0) {
                prompt(e.msg);
                return;
              }

              greeting({msg:"广告位价格修改成功！",confirm:function(){
                location.href  = location.href ;
              }});
              
        });

    });

  });
</script>