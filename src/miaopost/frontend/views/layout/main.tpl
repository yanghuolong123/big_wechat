<!DOCTYPE HTML>
<html>
        <head>
            {{template "layout/base-style.tpl" .}}
         </head>
         <body>
                <div class="container">
                        <div class="page-header row">
                                <div class="logo col-md-2">
                                       <a href="/">秒Po-UTD</a>
                                </div>
                                <div class="slogan col-md-10">
                                        <span>中国学生极简信息发布平台</span>
                                </div>
                        </div>
                        <div class="content">
                                {{.LayoutContent}}
                        </div>
                </div>

               
                {{template "layout/footer.tpl" .}}
                {{template "layout/base-script.tpl" .}}
         </body>
</html>