$(function(){
	$("#reback").click(function(){
	            window.history.back();
	});

	$('a.thumbnail').on('click', function() {
	    var $this = $(this);
	    var uploader = new PicUploader({
	        success: function(obj) {
	            $this.find('img').attr('src', obj['data']);
	            $this.next('input').val(obj['data']);
	        }
	    });
	    uploader.start();
	});

});

