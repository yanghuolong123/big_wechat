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
				//$this.parent().children('a').html('<span>重新上传</span>');
				//alert(e.data);

				//$('.logo').css( 'background-image', 'url(' + rs.data.fileUrl + ')' );
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
});