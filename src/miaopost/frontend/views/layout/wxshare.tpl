<script src="/static/js/jweixin-1.0.0.js"></script>
<script>
wx.config({
    debug: true,
    appId: '{{map_get .signPackage "appId"}}',
    timestamp: {{map_get .signPackage "timestamp"}},
    nonceStr: '{{map_get .signPackage "nonceStr"}}',
    signature: '{{map_get .signPackage "signature"}}',
    jsApiList: ['onMenuShareTimeline', 'onMenuShareAppMessage', 'onMenuShareQQ', 'onMenuShareQZone', 'onMenuShareWeibo']
});

wx.ready(function(){	
	
        // 分享到朋友圈
        wx.onMenuShareTimeline({
            title: '{{.wxshare.Title}}', // 分享标题
            link: '{{.wxshare.Link}}', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
            imgUrl: '{{.wxshare.Img}}', // 分享图标
            success: function () { 
                // 用户确认分享后执行的回调函数
            },
            cancel: function () { 
                // 用户取消分享后执行的回调函数
            }
        });

        // 分享给朋友
        wx.onMenuShareAppMessage({
          title: '{{.wxshare.Title}}', // 分享标题
          desc: '{{.wxshare.Desc}}', // 分享描述
          link: '{{.wxshare.Link}}', // 分享链接，该链接域名或路径必须与当前页面对应的公众号JS安全域名一致
          imgUrl: '{{.wxshare.Img}}', // 分享图标
          //type: '', // 分享类型,music、video或link，不填默认为link
          //dataUrl: '', // 如果type是music或video，则要提供数据链接，默认为空
          success: function () { 
              // 用户确认分享后执行的回调函数
          },
          cancel: function () { 
              // 用户取消分享后执行的回调函数
          }

      });  
      
      // 分享到QQ
      wx.onMenuShareQQ({
        title: '{{.wxshare.Title}}', // 分享标题
        desc: '{{.wxshare.Desc}}', // 分享描述
        link: '{{.wxshare.Link}}', // 分享链接
        imgUrl: '{{.wxshare.Img}}', // 分享图标
        success: function () { 
           // 用户确认分享后执行的回调函数
        },
        cancel: function () { 
           // 用户取消分享后执行的回调函数
        }
    });
    
    // 分享到QQ空间
    wx.onMenuShareQZone({
        title: '{{.wxshare.Title}}', // 分享标题
        desc: '{{.wxshare.Desc}}', // 分享描述
        link: '{{.wxshare.Link}}', // 分享链接
        imgUrl: '{{.wxshare.Img}}', // 分享图标
        success: function () { 
           // 用户确认分享后执行的回调函数
        },
        cancel: function () { 
            // 用户取消分享后执行的回调函数
        }
    });
    
    // 分享到腾讯微博
    wx.onMenuShareWeibo({
        title: '{{.wxshare.Title}}', // 分享标题
        desc: '{{.wxshare.Desc}}', // 分享描述
        link: '{{.wxshare.Link}}', // 分享链接
        imgUrl: '{{.wxshare.Img}}', // 分享图标
        success: function () { 
           // 用户确认分享后执行的回调函数
        },
        cancel: function () { 
            // 用户取消分享后执行的回调函数
        }
    });
    
});
</script>
