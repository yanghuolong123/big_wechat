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
	  {{if eq .tid 1}}
	  <div class="form-group">
	    <label for="tag" class="col-sm-2 control-label">广告标签</label>
	    <div class="col-sm-6">
	      <input type="text" class="form-control" name="Tag" placeholder="最多4个汉字，如美食、手机、留学等，不填则默认显示“推广”">
	    </div>
	  </div>
	  {{else }}
	   <div class="form-group">
	    <label for="photos" class="col-sm-2 control-label">广告图片</label>	    
	    <div class="col-sm-10">	
	      <div class="img-up">
                    <div class="img-up-list-1 clearfix">                        
                      
                    </div>
                    <div id="thelist-1" class="uploader-list"></div>
                    <div id="picker-1"><label class="user-img" for="imgs"></label></div>
                  </div>    
	      <input type="hidden" class="form-control" name="Logo" placeholder="">
	    </div>
	  </div>
	  {{end}}
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
	      <input type="hidden" class="form-control" name="Photos" placeholder="">
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
	    	<div class="rows">
	    		<div class="col-sm-2">
	    			<label>投放校区：</label>
	    			<select name="Region_id" class="form-control">
	    				<option value="{{.region.Id}}">{{.region.Shortname}}</option>
	    			</select>
	    		</div>
	    		<div class="col-sm-4">
	    			<label>投放位置：</label>
	    			<select name="Pos" class="form-control" id="adv_pos">
	    				<option value="">请选择</option>
		      		{{range .posList}}
		      			<option value="{{.Pos.Id}}">{{.Pos.Name}} (￥{{.AdvRe.Price}}/千次)</option>
		      		{{end}}
				</select>
				{{range .posList}}
					<input type="hidden" id="pos_price_{{.Pos.Id}}" value="{{.AdvRe.Price}}">
		      		{{end}}
	    		</div>
	    		<div class="col-sm-3">
	    			<label>展示次数(千次)：</label>
	       			<select name="Display_times" class="form-control" id="adv_show">
		       			<option value="1">1</option>
		       			<option value="10">10</option>
		       			<option value="20">20</option>
		       			<option value="30">30</option>
		       			<option value="40">40</option>
		       			<option value="50">50</option>
		       		</select>
	    		</div>
	    		<div class="col-sm-3">
	    			<label>每日上限(千次)：</label>
		       		<select name="Day_limit" class="form-control">
		       			<option value="1">1</option>
		       			<option value="2">2</option>
		       			<option value="3">3</option>
		       			<option value="4">4</option>
		       			<option value="5">5</option>
		       		</select>
	    		</div>
	    	</div>
	      
	      <br/>
	     
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="" class="col-sm-2 control-label">广告费用</label>
	    <div class="col-sm-10">
	      <span class="text-danger" id="total_fee"></span>
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
	.img-up-list-1 .img-li {
	    background-position: center center;
	    background-repeat: no-repeat;
	    background-size: cover;
	    border: 0.02rem solid #efefef;
	    float: left;
	    height: 8.45rem;
	    margin: 0 0 0.2rem 0.2rem;
	    position: relative;
	    width: 12.11rem;
	}
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
	$("#adv_pos,#adv_show").change(function(){
		var times = $("#adv_show").val();
		var pos = $("#adv_pos").val();
		if(parseInt(pos)>0) {
			var price = $("#pos_price_"+pos).val();
			$("#total_fee").text(price*times);
		} else {
			$("#total_fee").text("");
		}
	});

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

		var logo = "";
		for( var s=0; s<$('.img-up-list-1 .img-li-new').length; s++) {
			logo += $('.img-up-list-1 .img-li-new').eq(s).attr('data-url') + ',';
		}
		logo=logo.substring(0,logo.length-1);
		if(logo!="") {
			obj.Logo = logo;
		}

		if(obj.Type==2 && obj.Logo=="") {
			prompt("请先上传广告图片");
			return;
		}

		var photo="";
		for( var s=0; s<$('.img-up-list .img-li-new').length; s++) {
			photo += $('.img-up-list .img-li-new').eq(s).attr('data-url') + ',';
		}
		photo=photo.substring(0,photo.length-1);
		if(photo!="") {
			obj.Photos = photo;
		}
		
		if($.trim(obj.Target)!="" && !isUrl($.trim(obj.Target))) {
			prompt("目标网址url格式不正确");
			return;
		}
		if($.trim(obj.Pos)=="") {
			prompt("请选择投放位置");
			return;
		}


		//var data = $("#adv-form").serialize();
		//alert(data);
		$.post("/adv/create",obj,function(e){
			if(e.code<0) {
                                		prompt(e.msg);
                                		return false;
                        	}

                        	amount = e.data.Total_amount;
                        	var balance=0;
			$.ajax({
		                url:"/pay/balance",
		                async:false,
		                type: "POST",
		                data: {amount:amount, type:3,product_id:e.data.Id},
		                success: function(e){
		                        if(e.code<0) {
					return false;		
				}

				if(e.code==0) {
					balance = -1;
					prompt("广告信息成功！");		
					return false;
				}

				balance = e.data.Amount;
		                }
		       	});

			if(balance<0) {
				window.location = "/info";
				return false;
			}

			amount -= balance;
			amount = amount.toFixed(2);
			var msg = "亲，您的广告发布成功！支付后即可成功展示。钱包余额 "+balance+"元，还需支付："+amount+"元";

			if(isWeiXin()){					
                        		setTimeout(function(){
	                        		window.location.href = "/pay/confirm?product_id="+e.data.Id+"&amount="+amount+"&type=4&msg="+msg;
	                        	}, 2500);
			} else {
				prompt({msg:msg,displayTime:2500});
				setTimeout(function(){
					$("#pay_qr_img").removeClass("qrimg").attr("src", "/static/img/loading.gif");
					$(".pay_amount").html("￥"+amount+"元");
					$('#qrPayModal').modal({backdrop: 'static', keyboard: false});

					$.post("/pay/wxscan", {product_id:e.data.Id, amount:amount, type:4}, function(e){
						if(e.code<0) {
							prompt(e.msg);
							return false;
						}

						$("#pay_qr_img").attr("src", e.data.qrurl).addClass("qrimg");
						orderNo = e.data.order_no;

						var timer = setInterval(function(){
						    $.post('/pay/check', {order_no:orderNo}, function(e){
						            if(e.code < 0) {
						                return false;
						            }
						            
						            clearInterval(timer);
						            $('#qrPayModal').modal("hide");
						            prompt({msg:"广告支付成功！感谢您的支持！",displayTime:3000});
						            setTimeout(function(){
				                        		window.location = "/info";
				                        	}, 2500);
						            
						        });
						}, 1000);

					});
				}, 2500);				
			}



		});
	});

});
</script>
<script type="text/javascript">
$(function() {
    
    $list1 = $('#thelist-1'),
    state = 'pending';
   // uploader;
     
    uploader = WebUploader.create({
        // swf文件路径
        swf: '/static/plugin/webuploader/Uploader.swf',
        // 文件接收服务端。
        server: '/webupload',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker-1',
        // 不压缩image, 默认如果是jpeg，文件上传前会压缩一把再上传！
        resize: false,
        compress: false,
        chunked: true,
        chunkSize:100*1024,
        //sendAsBinary:true,
        accept: {
	    title: 'Images',
	    extensions: 'gif,jpg,jpeg,bmp,png',
	    mimeTypes: 'image/jpg,image/jpeg,image/png,image/gif',
	},
        fileNumLimit: 1,
        auto: true
    });
    
    // 当有文件被添加进队列的时候
    uploader.on( 'fileQueued', function( file ) {
    $list1.append( '<div id="' + file.id + '" class="item">' +
     // '<span class="info">' + file.name + ' </span>' +
    //  '<span class="state"> 等待上传...</span>' +
      //'<span class="text">0%</span>' +
    '</div>' );
    });
    
    // 文件上传过程中创建进度条实时显示。
    uploader.on( 'uploadProgress', function( file, percentage ) {
        var $li = $( '#'+file.id ),
          $percent = $li.find('.progress .progress-bar');

        // 避免重复创建
        if ( !$percent.length ) {
          $percent = $('<div class="progress progress-striped active">' +
           '<div class="progress-bar" role="progressbar" style="width: 0%">' +
           '</div>' +
          '</div>').appendTo( $li ).find('.progress-bar');
        }

        //$li.find('span.state').text('上传中，请等候... ');
        
        $percent.css( 'width', percentage * 100 + '%' );
        // if(percentage==1) {
        //     percentage = 0.99;
        // }
        //$li.find('span.text').text( Math.round( percentage * 100 ) + '%' );
    });
    
    uploader.on( 'uploadSuccess', function( file, obj ) {
        //$( '#'+file.id ).find('span.state').text('上传成功 ');
        //$( '#'+file.id ).find('span.text').text('100% ');        
        	var upImg = obj.data;
	if( obj.code == 0 && upImg!="") {
		$('.img-up-list-1').append('<div class="img-li img-li-new" data-url="' + upImg+ '"  data-big="' + upImg + '" style="background-image:url(' + upImg+ '!200!200)"><i></i></div>');
	}
    });

    uploader.on( 'uploadError', function( file ) {
        $( '#'+file.id ).find('span.state').text('上传出错 ');
    });

    uploader.on( 'uploadComplete', function( file ) {
        $( '#'+file.id ).find('.progress').fadeOut();        	
    });

    // $("#picker").click(function(){
    //     uploader.upload();
    // });
    
});
</script>