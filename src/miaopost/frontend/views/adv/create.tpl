<div class="adv-create">
	<form class="form-horizontal" id="adv-form">
	  <input type="hidden" name="Type" value="{{.tid}}">
	  <div class="form-group">
	    <label for="merch_name" class="col-sm-2 control-label">商户名称</label>
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="Merch_name" placeholder="商户名称">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="contact" class="col-sm-2 control-label">联系方式</label>
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="Contact" placeholder="电话/qq/微信">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="tag" class="col-sm-2 control-label">广告标签</label>
	    <div class="col-sm-6">
	      <input type="text" class="form-control" name="Tag" placeholder="最多4个汉字，如美食、手机、留学等，不填则默认显示“推广”">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="content" class="col-sm-2 control-label">广告内容</label>
	    <div class="col-sm-10">
	      <p class="text-muted">(前300个文字会作为广告概述显示在信息列表。全部文字信息会显示在点击广告后的图文详情页。如希望广告指向已有网页，则只需填写广告概述及目标网页地址，不必上传详情页配图。)</p>
	      <textarea class="form-control" name="Content" rows="8" placeholder=""></textarea>
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="photos" class="col-sm-2 control-label">详情页配图</label>	    
	    <div class="col-sm-10">	
	      <div class="img-up">
                    <div class="img-up-list clearfix">                        
                      
                    </div>
                    <div id="thelist" class="uploader-list"></div>
                    <div id="picker"><label class="user-img" for="imgs"></label></div>
                  </div>    
	      <input type="hidden" class="form-control" name="Photos" placeholder="最多4个汉字，如美食、手机、留学等，不填则默认显示“推广”">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="target" class="col-sm-2 control-label">目标网页地址</label>
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="Target" placeholder="例如：http://www.miaopost.com">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="" class="col-sm-2 control-label">投放目标及强度</label>
	    <div class="col-sm-10">
	      <input type="hidden" name="Region_id" value="{{.region.Id}}">
	      {{.region.Shortname}}
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="" class="col-sm-2 control-label">广告费用</label>
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="Amount" placeholder="">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="recom_code" class="col-sm-2 control-label">推荐码</label>
	    <div class="col-sm-4">
	      <input type="text" class="form-control" name="Recom_code" placeholder="不知道可不填，不影响广告费用">
	    </div>
	  </div>


	  <div class="form-group">
	    <div class="col-sm-offset-2 col-sm-10">
	      <button type="button" class="btn btn-default btn-lg btn-primary btn-adv">发布广告</button>
	    </div>
	  </div>
	</form>
</div>

<style type="text/css">
	form label{color: #777;}
</style>

<link rel="stylesheet" href="/static/plugin/kindeditor/themes/default/default.css" />
<script charset="utf-8" src="/static/plugin/kindeditor/kindeditor-all.modify.js"></script>
<script charset="utf-8" src="/static/plugin/kindeditor/lang/zh-CN.js"></script>
<script type="text/javascript">
	var editor;
	KindEditor.ready(function(K) {
		editor = K.create('textarea[name="Content"]', {
			resizeType : 1,
			allowPreviewEmoticons : false,
			allowImageUpload : true,
			uploadJson:"/kuploadfile",
			afterBlur: function () { this.sync(); },
			//afterFocus: function(){ this.html("");},
			items : [
				//'fontname', 'fontsize', '|', 'forecolor', 'hilitecolor', 'bold', 'italic', 'underline',
				//'removeformat', '|', 'justifyleft', 'justifycenter', 'justifyright', 'insertorderedlist',
				//'insertunorderedlist', '|', 'emoticons', 'image', 'link'
				//'insertorderedlist','bold',  '|', 'emoticons', 'image', 'link'
				'insertorderedlist','bold',  '|', 'emoticons'
				]
		});		
	});

	$(function(){
		$('#msgModal').off('shown.bs.modal').on('shown.bs.modal', function (e) {
		    $(document).off('focusin.modal');//解决编辑器弹出层文本框不能输入的问题
		});
	});
</script>
<script type="text/javascript">
$(function(){

	$(".btn-adv").click(function(){

		var obj = $("#adv-form").serializeObject();
		if($.trim(obj.Merch_name)=="") {
			prompt("商户名称不能为空");
			return;
		}
		if($.trim(obj.Contact)=="") {
			prompt("商户名称不能为空");
			return;
		}
		if($.trim(obj.Tag)!=""&&$.trim(obj.Tag).length>4) {
			prompt("标签最多4个字");
			return;
		}
		if($.trim(obj.Content)=="") {
			prompt("内容概述不内为空");
			return;
		}

		var photo="";
		for( var s=0; s<$('.img-li-new').length; s++) {
			photo += $('.img-li-new').eq(s).attr('data-url') + ',';
		}
		photo=photo.substring(0,photo.length-1);
		if(photo!="") {
			// prompt("请先上传");
			// return;
			obj.Photos = photo;
		}
		
		if($.trim(obj.Target)!="" && !isUrl($.trim(obj.Target))) {
			prompt("目标网址url格式不正确");
			return;
		}
		if($.trim(obj.Amount)=="") {
			prompt("金额不正确");
			return;
		}


		//var data = $("#adv-form").serialize();
		//alert(data);
		$.post("/adv/create",obj,function(e){

		});
	});
});
</script>