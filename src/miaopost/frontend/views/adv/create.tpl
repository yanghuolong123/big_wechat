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
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="Tag" placeholder="最多4个汉字，如美食、手机、留学等，不填则默认显示“推广”">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="content" class="col-sm-2 control-label">广告内容</label>
	    <div class="col-sm-10">
	      <textarea class="form-control" name="Content" rows="5" placeholder="前300个文字会作为广告概述显示在信息列表。全部文字信息会显示在点击广告后的图文详情页。如希望广告指向已有网页，则只需填写广告概述及目标网页地址，不必上传详情页配图。"></textarea>
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="photos" class="col-sm-2 control-label">详情页配图</label>
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="Photos" placeholder="最多4个汉字，如美食、手机、留学等，不填则默认显示“推广”">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="target" class="col-sm-2 control-label">目标网页地址</label>
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="Target" placeholder="最多4个汉字，如美食、手机、留学等，不填则默认显示“推广”">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="" class="col-sm-2 control-label">投放目标及强度</label>
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="" placeholder="">
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
	    <div class="col-sm-10">
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
		if($.trim(obj.Content)=="") {
			prompt("商户名称不能为空");
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