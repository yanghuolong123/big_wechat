<div class="box box-primary">
            <div class="box-header with-border">
              <h3 class="box-title">创建导航</h3>
            </div>
            <!-- /.box-header -->
            <!-- form start -->
            <form method="post" role="form" id="create_form">
              <div class="box-body">
              	<div class="form-group">
                  <label for="group_id">分组</label>
                  <select id="group_id" name="group_id" class="form-control">
                  	<option value="">请选择</option>
                  	{{range $k,$v := .groupMap }}
                  	{{if  eq $k $.type}}
                  	{{range $kk,$vv := $v}}                  	
                  	<option value="{{$kk}}">{{$vv}}</option>
                  	{{end}}
                  	{{end}}
                  	{{end}}
                  </select>
                </div>
                <div class="form-group">
                  <label for="title">标题</label>
                  <input type="text" placeholder="文章标题" name="title" id="title" class="form-control">
                </div>
                <div class="form-group">
                  <label for="content">内容</label>
                  <textarea id="content"  placeholder="Enter ..." rows="3" class="form-control"></textarea>
                </div>
                <div class="form-group">
                  <label for="sort">排序</label>
                  <input type="text" name="sort" id="sort" placeholder="排序" class="form-control">
                </div>
                <div class="form-group">
                  <label for="status">状态</label>
                  <select id="status" name="status" class="form-control">
                  	{{range $k,$v := .statusMap}}
                  	<option value="{{$k}}" {{if eq $k 0}}selected{{end}}>{{$v}}</option>
                  	{{end}}
                  </select>
                </div>
                
              </div>
              <!-- /.box-body -->

              <div class="box-footer">
              	<input type="hidden" name="type" value="{{.type}}">
                <button class="btn btn-primary" id="create">Submit</button>
              </div>
            </form>
          </div>

<!-- CK Editor -->
<script src="/static/plugin/bower_components/ckeditor/ckeditor.js"></script>
<script type="text/javascript">
	$(function(){

		$("form").submit(function(){
			var form = document.getElementById("create_form");
			var fdata = new FormData(form);
			fdata.append("content",CKEDITOR.instances.content.getData());
			if(fdata.get("group_id")=="") {
				prompt("请选择分组");
				return false;
			}
			if ($.trim(fdata.get("title")) == "") {
				prompt("文章标题不能为空");
				return false;
			}
			if($.trim(fdata.get("content"))=="") {
				prompt("文章内容为必填不能为空");
				return false;
			}
			if($.trim(fdata.get("sort"))!="" && !isNum(fdata.get("sort"))) {
				prompt("排序请填写数字");
				return false;
			}
			
			$.ajax({
				url:"/article/create",
				type:"post",
				data:fdata,
				processData: false,
				contentType: false,
				dataType: "json",
				cache: false,
				success:function(e){
					if (e.code<0) {
						prompt(e.msg);
						return false;
					}

					var type = e.data.Type;
					greeting({
						msg:"创建成功",
						confirm: function(){
							location.href = "/article/list?type="+type;
						},
					})
				},
			});
			return false;
		});

		CKEDITOR.replace('content',{
			filebrowserImageUploadUrl:"/ckuploadfile",
			image_previewText:" ",
		})

	});	
</script>