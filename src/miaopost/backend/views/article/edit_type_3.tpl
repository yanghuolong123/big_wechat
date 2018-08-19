<div class="box box-primary">
            <div class="box-header with-border">
              <h3 class="box-title">编辑文章</h3>
            </div>
            <!-- /.box-header -->
            <!-- form start -->
            <form method="post" role="form" id="edit_form">
              <div class="box-body">  
                <div class="form-group">
                  <label for="content">内容</label>
                  <textarea  id="content"  placeholder="Enter ..." rows="3" class="form-control">{{.article.Content}}</textarea>
                </div> 
                <div class="form-group">
                  <label for="status">状态</label>
                  <select id="status" name="status" class="form-control">
                  	{{range $k,$v := .statusMap}}
                  	<option value="{{$k}}" {{if eq $k $.article.Status}}selected{{end}}>{{$v}}</option>
                  	{{end}}
                  </select>
                </div>
                
              </div>
              <!-- /.box-body -->

              <div class="box-footer">
              	<input type="hidden" name="id" value="{{.article.Id}}">
                <button class="btn btn-primary" id="create">Submit</button>
              </div>
            </form>
          </div>

<!-- CK Editor -->
<script src="/static/plugin/bower_components/ckeditor/ckeditor.js"></script>
<script type="text/javascript">
	$(function(){

		$("form").submit(function(){
			var form = document.getElementById("edit_form");
			var fdata = new FormData(form);
			fdata.append("content",CKEDITOR.instances.content.getData());
			 
			if($.trim(fdata.get("content"))=="") {
				prompt("文章内容为必填不能为空");
				return false;
			}
			 
			
			$.ajax({
				url:"/article/edit",
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

					greeting({
						msg:"编辑成功",
						confirm: function(){
							location.href = "/article/list?type=3";
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