<div class="create_pg">
	<div class="row title">
		<div class="col-md-offset-1 col-md-1">
			<h4 class="text-primary">发布群</h4>
		</div>
		<div class="col-md-6">
			<p class="text-muted">建立微信群人数不够，可以添加 <a href="javascript:;">小助手微信</a> 凑够3人</p>
		</div>
	</div>
	<form class="form-horizontal">
	  <div class="form-group">
	    <label for="school" class="col-sm-2 control-label">应用范围</label>
	    <div class="col-sm-6">
	      <input type="hidden" name="gid" id="gid" value="">
	      <input type="text" class="form-control" id="school" placeholder="请输入学校关键字或简称">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="name" class="col-sm-2 control-label">群名称</label>
	    <div class="col-sm-6">
	      <input type="text" class="form-control" id="name" placeholder="">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="introduction" class="col-sm-2 control-label">补充介绍</label>
	    <div class="col-sm-6">	    
	      <textarea class="form-control"  id="introduction" rows="3" placeholder="可补充群名称中未包含的信息、关键字等"></textarea>
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="uploadfile" class="col-sm-4 control-label">上传联系信息 (以下三项至少填一项)</label>	   
	  </div>
	  <div class="form-group">
	    <div class="col-sm-offset-2 col-sm-10">	 
	    	<div class="row">
	      	<div class="uploadimg col-sm-2">
	      		<input type="hidden" value="" name="qrcode">	      
		    	<input type="file" name="qrcode_file" style="display: none" id="qrcode_file">
                                       	<a role="button" href="#" class="upload_btns" onclick="uploadFile(this);return false;" id="file_qrcode_upload">上传群二维码</a>
	      	</div>
	      	<div class="uploadimg col-sm-2">
	      		<input type="hidden" value="" name="ower_qrcode">	      
		    	<input type="file" name="ower_qrcode_file" style="display: none" id="qrcode_file">
                                       	<a role="button" href="#" class="upload_btns" onclick="uploadFile(this);return false;" id="file_ower_qrcode_upload">上传群主二维码</a>	      	
	      	</div>
	      	</div>
	      	<div class="row">
	      		<label for="wechat_id" class="col-sm-2 control-label wechat_id_label">群主微信号</label>
	      		<div class="col-sm-4">
	      			<input type="text" class=" form-control" id="wechat_id" placeholder="">
	      		</div>	      		
	      	</div>
	    </div>
	  </div>	  
	  <div class="form-group">
	    <div class="col-sm-offset-2 col-sm-10">
	      <button type="submit" class="btn btn-success btn-lg">发布</button>
	    </div>
	  </div>
	</form>
</div>