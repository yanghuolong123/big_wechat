$(function() {
    
    $list = $('#thelist'),
    state = 'pending';
   // uploader;
     
    uploader = WebUploader.create({
        // swf文件路径
        swf: '/static/plugin/webuploader/Uploader.swf',
        // 文件接收服务端。
        server: '/webupload',
        // 选择文件的按钮。可选。
        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
        pick: '#picker',
        // 不压缩image, 默认如果是jpeg，文件上传前会压缩一把再上传！
        resize: false,
        compress: false,
        chunked: true,
        chunkSize:100*1024,
        sendAsBinary:true,
        accept: {
	    title: 'Images',
	    extensions: 'gif,jpg,jpeg,bmp,png',
	    mimeTypes: 'image/jpg,image/jpeg,image/png,image/gif'
	},
        fileNumLimit: 20,
        auto: true
    });
    
    // 当有文件被添加进队列的时候
    uploader.on( 'fileQueued', function( file ) {
    $list.append( '<div id="' + file.id + '" class="item">' +
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
		$('.img-up-list').append('<div class="img-li img-li-new" data-url="' + upImg+ '"  data-big="' + upImg + '" style="background-image:url(' + upImg+ '!200!200)"><i></i></div>');
	}
    });

    uploader.on( 'uploadError', function( file ) {
        $( '#'+file.id ).find('span.state').text('上传出错 ');
    });

    uploader.on( 'uploadComplete', function( file ) {
        $( '#'+file.id ).find('.progress').fadeOut();        	
    });

    $("#picker").click(function(){
        uploader.upload();
    });
    
});