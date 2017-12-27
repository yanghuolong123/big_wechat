// addwechat.js 

function uploadFile(obj) {
	var $this = $(obj);
	  //取得前面的DIv的第一个元素，也就是file框。
	  var $pre = $this.prev();	  
	  $pre.click();

	  $pre.on('change',function(){
	  	 var $upload = $(this);
	  	 if ($upload.val() == null || $upload.val() =="") {
	  	 	alert("请选择文件！");
			 return false;
		 } 

		 var formData = new FormData();
		formData.append('file', $upload[0].files[0]);				
		$.ajax({
			url: '/uploadfile',
			type: 'post',
			cache: false,
			data: formData,
			dataType: "json",
			processData: false,
			contentType: false,
			success:function(e,textStatus,jqXHR){
				$this.parent().css( 'background-image', 'url(/' +  e.data + ')' );
				$this.parent().addClass('active');
				$upload .prev().val(e.data);
			}
		});		
	  });
}

$(function(){

	$("#loginBtn").click(function(){
		$.get("/login", function(e){
			$("#modalPage").html(e.data).find('#loginModal').modal({backdrop: 'static', keyboard: false});

			var timer = setInterval(function(){
			    $.post('/login', {"sceneId":$("#sceneId").val()}, function(e){
			            if(e.code < 0) {
			                return false;
			            }
			            
			            clearInterval(timer);
			            window.location = "/";
			            
			        });
			}, 1000);
		});
	});


	$("#create_pg_btn").click(function(){
		$this = $(this);
		var gid = 1;// $("#gid").val();
		var name = $.trim($("#name").val());
		var introduction = $("#introduction").val();
		var qrcode = $("#qrcode").val();
		var ower_qrcode = $("#ower_qrcode").val();
		var wechat_id = $.trim($("#wechat_id").val());

		var flag = true;
		$(".error_tips").html("");
		if (gid=="") {
			$(".error_tips").append("<p>应用范围为必填项。</p>");
			flag = false;
		}
		if (name=="") {
			$(".error_tips").append("<p>群名称为必填项。</p>");
			flag = false;
		}
		if(qrcode == "" && ower_qrcode=="" && wechat_id == "") {
			$(".error_tips").append("<p>上传联系信息至少要填一项。</p>");
			flag = false;
		}

		if (!flag) {
			return false;
		}

		$this.attr("disabled","disabled");
		$.post("/pg/create", {gid:gid, name:name, introduction:introduction,qrcode:qrcode, ower_qrcode:ower_qrcode, wechat_id:wechat_id}, function(e){
			$this.removeAttr("disabled");
			if(e.code<0) {
				$(".error_tips").append(e.msg);
				return false;
			}

			window.location = "/";
		});
	});

});