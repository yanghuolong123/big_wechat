<div class="adv-create">
	<form class="form-horizontal" id="adv-form">
	  <input type="hidden" name="Type" value="{{.tid}}">
	  <input type="hidden" name="Id" value="{{.vo.A.Id}}">
	  <div class="form-group">
	    <label for="merch_name" class="col-sm-2 control-label">商户名称</label>
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="Merch_name" placeholder="商户名称" value="{{.vo.A.Merch_name}}">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="contact" class="col-sm-2 control-label">联系方式</label>
	    <div class="col-sm-10">
	      <input type="text" class="form-control" name="Contact" placeholder="电话/qq/微信" value="{{.vo.A.Contact}}">
	    </div>
	  </div>
	  {{if eq .tid 1}}
	  <div class="form-group">
	    <label for="tag" class="col-sm-2 control-label">广告标签</label>
	    <div class="col-sm-6">
	      <input type="text" class="form-control" name="Tag" value="{{.vo.A.Tag}}" placeholder="最多4个汉字，如美食、手机、留学等，不填则默认显示“推广”">
	    </div>
	  </div>
	  {{else }}
	   <div class="form-group">
	    <label for="photos" class="col-sm-2 control-label">广告图片</label>	    
	    <div class="col-sm-10">	
	      <div class="img-up">
                    <div class="img-up-list-1 clearfix">                        
                      <div class="img-li img-li-new" data-url="{{.vo.Logo}}"  style="background-image:url({{.vo.Logo}}!200!200)"><i></i></div>
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
	      <textarea class="form-control" name="Content" rows="8" placeholder="">{{.vo.A.Content}}</textarea>
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="photos" class="col-sm-2 control-label">详情页配图</label>	    
	    <div class="col-sm-10">	
	      <div class="img-up">
                    <div class="img-up-list clearfix">                        
                      {{range .vo.Photos}}
                        <div class="img-li img-li-new" data-url="{{.}}"  style="background-image:url({{.}}!200!200)"><i></i></div>
                      {{end}}
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
	      <input type="text" class="form-control" name="Target" value="{{.vo.A.Target}}" placeholder="例如：http://www.miaopost.com">
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="" class="col-sm-2 control-label">投放目标及强度</label>
	    <div class="col-sm-10">
	    	<div class="rows">
	    		<div class="col-sm-2">
	    			<label>投放校区：</label>
	    			<select name="Region_id" class="form-control" disabled="disabled">
	    				<option value="{{.region.Id}}">{{.region.Shortname}}</option>
	    			</select>
	    		</div>
	    		<div class="col-sm-4">
	    			<label>投放位置：</label>
	    			<select name="Pos" class="form-control" id="adv_pos" disabled="disabled">
	    				<option value="">请选择</option>
		      		{{range .posList}}
		      			{{if eq $.tid .Pos.Type }}
		      			<option value="{{.Pos.Id}}" {{if eq $.vo.A.Pos .Pos.Id}}selected="selected"{{end}}>{{.Pos.Name}} (￥{{.AdvRe.Price}}/千次)</option>
		      			{{end}}
		      		{{end}}
				</select>
				{{range .posList}}
					{{if eq $.tid .Pos.Type }}
					<input type="hidden" id="pos_price_{{.Pos.Id}}" value="{{.AdvRe.Price}}">
					{{end}}
		      		{{end}}
	    		</div>
	    		<div class="col-sm-3">
	    			<label>展示次数(千次)：</label>
	       			<select name="Display_times" class="form-control" id="adv_show" disabled="disabled">
		       			<option value="1" {{if eq $.vo.A.Display_times 1}}selected="selected"{{end}}>1</option>
		       			<option value="10" {{if eq $.vo.A.Display_times 10}}selected="selected"{{end}}>10</option>
		       			<option value="20" {{if eq $.vo.A.Display_times 20}}selected="selected"{{end}}>20</option>
		       			<option value="30" {{if eq $.vo.A.Display_times 30}}selected="selected"{{end}}>30</option>
		       			<option value="40" {{if eq $.vo.A.Display_times 40}}selected="selected"{{end}}>40</option>
		       			<option value="50" {{if eq $.vo.A.Display_times 50}}selected="selected"{{end}}>50</option>
		       		</select>
	    		</div>
	    		<div class="col-sm-3">
	    			<label>每日上限(千次)：</label>
		       		<select name="Day_limit" class="form-control">
		       			<option value="1" {{if eq $.vo.A.Day_limit 1}}selected="selected"{{end}}>1</option>
		       			<option value="2" {{if eq $.vo.A.Day_limit 2}}selected="selected"{{end}}>2</option>
		       			<option value="3" {{if eq $.vo.A.Day_limit 3}}selected="selected"{{end}}>3</option>
		       			<option value="4" {{if eq $.vo.A.Day_limit 4}}selected="selected"{{end}}>4</option>
		       			<option value="5" {{if eq $.vo.A.Day_limit 5}}selected="selected"{{end}}>5</option>
		       		</select>
	    		</div>
	    	</div>
	      
	      <br/>
	     
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="" class="col-sm-2 control-label">广告费用</label>
	    <div class="col-sm-10">
	      <span class="text-danger" id="total_fee">{{.vo.A.Total_amount}}￥</span>
	    </div>
	  </div>
	  <div class="form-group">
	    <label for="recom_code" class="col-sm-2 control-label">推荐码</label>
	    <div class="col-sm-4">
	      <input type="text" disabled="disabled" class="form-control" name="Recom_code" value="{{.vo.A.Recom_code}}" placeholder="不知道可不填，不影响广告费用">
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
	.img-up-list-1 .img-li i {
	    background: rgba(0, 0, 0, 0) url("/static/img/ico-x.png") no-repeat scroll center center / 100% auto;
	    border-radius: 0.4rem;
	    display: block;
	    height: 1.4rem;
	    position: absolute;
	    right: -0.3rem;
	    top: -0.2rem;
	    width: 1.4rem;
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
		// if($.trim(obj.Pos)=="") {
		// 	prompt("请选择投放位置");
		// 	return;
		// }

		$.post("/adv/edit?id={{.vo.A.Id}}",obj,function(e){
			if(e.code<0) {
                                		prompt(e.msg);
                                		return false;
                        	}  

                        	greeting({msg:"您发布的广告修改成功！",confirm:function(){
		               // location.href  = "/info";
		          	}});

		});
	});

});
</script>
<script type="text/javascript">
$(function() {
	$(".img-up-list-1").on("click", ".img-li i", function(){
		$(this).parent('.img-li').remove();
		$('.user-img').show();
		return false;
	});
    
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